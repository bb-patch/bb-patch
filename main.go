package main

// go build -ldflags "-s -w -H=windowsgui"

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"main.com/main/patch"
)

var workDir, _ = os.Getwd()
var location = workDir + "\\"
var gamefile = location + "BEASTS.EXE"

var outputConsole = widget.NewTextGrid()
var scrollableTextArea = container.NewVScroll(outputConsole)

func log(format string, a ...any) {
	current := outputConsole.Text()
	s := fmt.Sprintf(format, a...)
	outputConsole.SetText(current + s + "\n")
	scrollableTextArea.ScrollToBottom()
	scrollableTextArea.Refresh()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Beasts and Bumpkins Fan Patch")
	tabs := container.NewAppTabs()

	// Text Area
	scrollableTextArea.SetMinSize(fyne.NewSize(640, 200))

	// Install Container
	installContent := container.NewVBox(
		widget.NewLabel("Install from folder (e.g., a CD Drive)"),
		widget.NewButton("Pick folder", func() {
			diag := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if err != nil {
					log(err.Error())
					return
				}
				err = copyFromFolder(lu.Path() + "\\")
				if err != nil {
					log("Error while copying: ")
					log(err.Error())
				}
				tabs.SelectIndex(1)
			}, myWindow)
			diag.Show()
		}),
	)

	installContainer := container.NewTabItemWithIcon("Install", theme.ContentCopyIcon(), installContent)

	// Patch Container
	innerPatchContainer := container.NewVBox()
	patches := container.New(layout.NewFormLayout())
	checkBoxes := make([]*widget.Check, 0, len(patch.Patches))

	patches.Add(widget.NewSeparator())
	patches.Add(widget.NewSeparator())
	for _, pf := range patch.Patches {
		c := widget.NewCheck(pf.Id(), func(bool) {})
		c.SetChecked(!pf.IsOptional())
		checkBoxes = append(checkBoxes, c)

		patches.Add(c)
		patches.Add(pf.Widget())
		patches.Add(widget.NewSeparator())
		patches.Add(widget.NewSeparator())
	}
	innerPatchContainer.Add(patches)

	actionButtons := container.NewHBox(
		widget.NewButton("Apply selected", func() {
			toApply := make([]patch.Patch, 0)
			for i, c := range checkBoxes {
				if !c.Checked {
					continue
				}
				toApply = append(toApply, patch.Patches[i].Build())
			}
			err := applyPatches(toApply)
			if err != nil {
				log("Error occured while applying patches - aborting:")
				log(err.Error())
			}
		}),
		widget.NewButton("Remove selected", func() {
			toRemove := make([]patch.Patch, 0)
			for i, c := range checkBoxes {
				if !c.Checked {
					continue
				}
				toRemove = append(toRemove, patch.Patches[i].Build())
			}
			err := removePatches(toRemove)
			if err != nil {
				log("Error occured while removing patches - aborting:")
				log(err.Error())
			}
		}),
	)
	innerPatchContainer.Add(actionButtons)
	patchContainer := container.NewTabItemWithIcon("Patch", theme.LoginIcon(), innerPatchContainer)

	// Tabs
	tabs.Append(installContainer)
	tabs.Append(patchContainer)
	tabs.SetTabLocation(container.TabLocationLeading)

	go func() {
		_, err := os.Stat(gamefile)
		if err != nil {
			tabs.DisableItem(patchContainer)
		} else {
			tabs.SelectIndex(1)
		}
		for {
			_, err := os.Stat(gamefile)
			if err != nil {
				tabs.DisableItem(patchContainer)
			} else {
				tabs.EnableItem(patchContainer)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Finalizing
	content := container.NewVBox()
	content.Add(tabs)
	content.Add(scrollableTextArea)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(640, 480))
	myWindow.ShowAndRun()
}

func copyFromFolder(path string) error {
	for i, v := range BeastAndBumpkinsFiles {
		log("Moving file (%d/%d): %s", i+1, len(BeastAndBumpkinsFiles), path+v)
		err := copyFile(path+v, location+v)
		if err != nil {
			return err
		}
	}
	log("Done moving files!")
	return nil
}

func copyFile(input, output string) error {
	r, err := os.Open(input)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := os.Create(output)
	if err != nil {
		return err
	}
	defer w.Close()
	w.ReadFrom(r)

	return nil
}

func removePatches(patches []patch.Patch) error {
	log("Removing patches...")

	file, err := readExeFile()
	if err != nil {
		return err
	}

	for _, p := range patches {
		log("Removing patch: " + p.Description())
		err = p.RemovePatch(file)
		if err != nil {
			return err
		}
	}
	writeExeFile(file)
	log("Done Removing patches...\n")

	return nil
}

func applyPatches(patches []patch.Patch) error {
	log("Applying patches...")

	file, err := readExeFile()
	if err != nil {
		return err
	}

	for _, p := range patches {
		log("Applying patch: " + p.Description())
		err = p.PatchFile(file)
		if err != nil {
			return err
		}
	}
	writeExeFile(file)
	log("Done Applying patches...\n")

	return nil
}

func readExeFile() ([]byte, error) {
	b, err := os.ReadFile(gamefile)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func writeExeFile(input []byte) error {
	err := os.WriteFile(gamefile, input, 0666)
	return err
}
