package patch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type noCdPatchFactory struct {
}

func (pf *noCdPatchFactory) Id() string {
	return "NoCD"
}

func (pf *noCdPatchFactory) Build() Patch {
	return createNoCdPatch()
}

func (pf *noCdPatchFactory) Widget() fyne.CanvasObject {
	return widget.NewLabel("Required. Alters the game so that a CD is not required.")
}

func (pf *noCdPatchFactory) IsOptional() bool {
	return false
}

func createNoCdPatch() Patch {
	indexOfJump := 0x2DE39 // Address 0x0042ea39
	originalOpCodes := []byte{0xE9, 0x62, 0xFD, 0xFF, 0xFF}

	return simplePatch{
		patchFileFunc: func(input []byte) error {
			input[indexOfJump] = 0x90
			input[indexOfJump+1] = 0x90
			input[indexOfJump+2] = 0x90
			input[indexOfJump+3] = 0x90
			input[indexOfJump+4] = 0x90
			return nil
		},
		removePatchFunc: func(input []byte) error {
			copy(input[indexOfJump:], originalOpCodes)
			return nil
		},
		descriptionFunc: func() string {
			return "No CD required"
		},
	}
}
