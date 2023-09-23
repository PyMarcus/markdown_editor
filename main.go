package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	openMenu := fyne.NewMenuItem("Open", func() {})

	saveMenu := fyne.NewMenuItem("Save", func() {})

	saveAs := fyne.NewMenuItem("Save as", func() {})

	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAs)

	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}
