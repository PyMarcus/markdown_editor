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
	// set the content of the window
	w.SetContent(container.NewHSplit(edit, preview))
	// window size
	w.SetFixedSize(true)
	w.Resize(fyne.Size{Width: 480, Height: 600})
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
