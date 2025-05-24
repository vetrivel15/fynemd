package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

func TestMakeUI(t *testing.T) {

	var testConfig config
	edit, preview := testConfig.makeUI()

	test.Type(edit, "Hello universe")

	if preview.String() != "Hello universe" {
		t.Errorf("Failed did not find expected value in preview")
	}
}

func RunApp(t *testing.T) {

	var testConfig config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test Markdown")

	edit, preview := testConfig.makeUI()

	testConfig.createMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))
	testApp.Run()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {
		t.Errorf("Failed did not find the expected in preview")
	}
}
