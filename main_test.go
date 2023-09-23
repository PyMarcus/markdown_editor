package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
)

func Test_makeUI(t *testing.T) {
	var testConfig config

	a := createApp()
	w := createWindow(a, "Markdown Editor")
	edit, preview := testConfig.makeUI()
	cfg.createMenuItems(w)
	w.SetContent(container.NewHSplit(edit, preview))
	test.Type(edit, "Is a test")
	if preview.String() != "Is a test" {
		t.Error("Failed! Expected: Is a test, received: ", preview.String())
	}
}

func Test_openFunc(t *testing.T) {
	a := createApp()
	window:=createWindow(a, "Markdown Editor")

	z := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {

		if read == nil {
			t.Error("Read is null failt to read")
		}

		defer read.Close()

		if err != nil {
			t.Error(err)
		}

		
	}, window)

	z.Show()
}
