package main

import (
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	editWidget    *widget.Entry
	previewWidget *widget.RichText
	currentFile   fyne.URI
	saveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// create a fyne app
	a := createApp()
	// create a windows for the app
	w := createWindow(a, "Markdown Editor")
	// get user interface
	edit, preview := cfg.makeUI()
	// create menu items
	cfg.createMenuItems(w)
	// set the content of the window
	w.SetContent(container.NewHSplit(edit, preview))
	// window size
	w.SetFixedSize(true)
	w.Resize(fyne.Size{Width: 800, Height: 500})
	// open on the center of screen
	w.CenterOnScreen()
	// show window and run app
	w.ShowAndRun()
}

func createApp() fyne.App {
	return app.New()
}

func createWindow(app fyne.App, title string) fyne.Window {
	return app.NewWindow(title)
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.editWidget = edit
	app.previewWidget = preview

	// listener
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}

func (app *config) createMenuItems(window fyne.Window) {
	openMenu := fyne.NewMenuItem("Open", app.openFunc(window))

	saveMenu := fyne.NewMenuItem("Save", app.saveFunc(window))

	app.saveMenuItem = saveMenu
	app.saveMenuItem.Disabled = true

	saveAs := fyne.NewMenuItem("Save as", app.saveAsFunc(window))

	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAs)

	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}

// filter:
var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) openFunc(window fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {

			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)

			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			if read != nil {
				app.editWidget.SetText(string(data))

				app.currentFile = read.URI()
				window.SetTitle(window.Title() + " - " + read.URI().Name())
				app.saveMenuItem.Disabled = false
			}
			return
		}, window)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveFunc(window fyne.Window) func() {
	return func() {
		if app.currentFile != nil {
			write, err := storage.Writer(app.currentFile)

			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			write.Write([]byte(app.editWidget.Text))

			defer write.Close()
		}
	}
}

func (app *config) saveAsFunc(window fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {

			log.Println("User cancel", uc)

			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			// if the user cancel
			if uc == nil {
				return
			}

			defer uc.Close()

			// save
			if uc != nil {

				if !strings.HasSuffix(uc.URI().String(), ".md") {
					dialog.ShowInformation("Ops!",
						"Invalid format. Please, name your file with .md extension",
						window)
				}

				uc.Write([]byte(app.editWidget.Text))
				app.currentFile = uc.URI()

				window.SetTitle(window.Title() + " - " + uc.URI().Name())

				app.saveMenuItem.Disabled = false
			}
			return
		}, window)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
