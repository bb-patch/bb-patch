package patch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type developerConsolePatchFactory struct {
}

func (pf *developerConsolePatchFactory) Id() string {
	return "DeveloperConsole"
}

func (pf *developerConsolePatchFactory) Build() Patch {
	return createDeveloperConsolePatch()
}

func (pf *developerConsolePatchFactory) Widget() fyne.CanvasObject {
	return widget.NewLabel("Optional. Adds the developer console F11 to the game (allows cheats).")
}

func (pf *developerConsolePatchFactory) IsOptional() bool {
	return true
}

func createDeveloperConsolePatch() Patch {
	// 24 = 0x0048427c + 39
	p := singleAddressPatch(534691, 0, 1).(simplePatch)
	p.descriptionFunc = func() string {
		return "Enabling developer console"
	}

	return p
}
