package patch

import (
	"errors"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type setResolutionPatchFactory struct {
	width  uint32
	height uint32
}

func (pf *setResolutionPatchFactory) Id() string {
	return "SetResolution"
}

func (pf *setResolutionPatchFactory) Build() Patch {
	return createSetResolutionPatch(pf.width, pf.height)
}

func (pf *setResolutionPatchFactory) IsOptional() bool {
	return true
}

func (pf *setResolutionPatchFactory) Widget() fyne.CanvasObject {
	validator := func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if i < 100 {
			return errors.New("must be greater than 100")
		}
		return nil
	}

	width := widget.NewEntry()
	width.Text = strconv.Itoa(int(pf.width))
	width.Validator = validator

	height := widget.NewEntry()
	height.Text = strconv.Itoa(int(pf.height))
	height.Validator = validator

	width.OnChanged = func(s string) {
		i, _ := strconv.Atoi(s)
		pf.width = uint32(i)
	}

	height.OnChanged = func(s string) {
		i, _ := strconv.Atoi(s)
		pf.height = uint32(i)
	}

	return container.NewVBox(
		widget.NewLabel("Optional (but nice). Videos and menus are not upscaled. Game is very playable, but there are graphical glitches."),
		widget.NewLabel("Recommended resolution is 1280x720. Patch has only been tested on this."),
		widget.NewLabel("Width: "),
		width,
		widget.NewLabel("Height: "),
		height,
	)
}

func createSetResolutionPatch(width, height uint32) Patch {
	indexOfScreenWidth := 0x82A44  // Address 0x00484444
	indexOfScreenHeight := 0x82A48 // Address 0x00484448
	indexOfPortraitX := 0x29714    // Address 0x0042A310
	indexOfPortraitY := 0x2972D    // Address 0x0042A32A

	// Flip function in gameloop: Address 0040C57E

	defaultWidth := 640
	defaultHeight := 480

	return simplePatch{
		patchFileFunc: func(input []byte) error {
			writeDWORD(input, indexOfPortraitX, width-100)
			writeDWORD(input, indexOfPortraitY, height-100)
			writeDWORD(input, indexOfScreenWidth, width)
			writeDWORD(input, indexOfScreenHeight, height)
			return nil
		},
		removePatchFunc: func(input []byte) error {
			writeDWORD(input, indexOfPortraitX, uint32(defaultWidth)-100)
			writeDWORD(input, indexOfPortraitY, uint32(defaultHeight)-100)
			writeDWORD(input, indexOfScreenWidth, uint32(defaultWidth))
			writeDWORD(input, indexOfScreenHeight, uint32(defaultHeight))
			return nil
		},
		descriptionFunc: func() string {
			return fmt.Sprintf("Set screen resolution to %dx%d", width, height)
		},
	}
}
