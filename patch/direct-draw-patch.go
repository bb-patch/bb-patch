package patch

import (
	_ "embed"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type directDrawFixFactory struct {
}

func (pf *directDrawFixFactory) Id() string {
	return "DirectDrawFix"
}

func (pf *directDrawFixFactory) Build() Patch {
	return directDrawFix()
}

func (pf *directDrawFixFactory) IsOptional() bool {
	return false
}

func (pf *directDrawFixFactory) Widget() fyne.CanvasObject {
	url, _ := url.Parse("https://github.com/elishacloud/dxwrapper")

	return container.NewVBox(
		widget.NewLabel("Required. Fixes a bug in the standard DirectDraw library that makes the game unplayably slow on new computers."),
		container.NewHBox(
			widget.NewLabel("Uses DxWrapper, see"),
			widget.NewHyperlink("https://github.com/elishacloud/dxwrapper", url),
		),
	)
}

//go:embed resources/ddraw.dll
var ddrawDLL []byte

//go:embed resources/dxwrapper.dll
var dxwrapperDLL []byte

//go:embed resources/dxwrapper.ini
var dxwrapperINI []byte

func directDrawFix() Patch {
	var location, _ = os.Getwd()
	location += "\\"

	ddrawPath := location + "ddraw.dll"
	dxwrapperDLLPath := location + "dxwrapper.dll"
	dxwrapperINIPath := location + "dxwrapper.ini"

	removeFile := func(path string) error {
		err := os.Remove(path)
		if err != nil && os.IsExist(err) {
			return err
		}
		return nil
	}

	return simplePatch{
		patchFileFunc: func(_ []byte) error {
			err := os.WriteFile(ddrawPath, ddrawDLL, 0666)
			if err != nil {
				return err
			}

			err = os.WriteFile(dxwrapperDLLPath, dxwrapperDLL, 0666)
			if err != nil {
				return err
			}

			err = os.WriteFile(dxwrapperINIPath, dxwrapperINI, 0666)
			if err != nil {
				return err
			}

			return nil
		},
		removePatchFunc: func(b []byte) error {
			err := removeFile(ddrawPath)
			if err != nil {
				return err
			}

			err = removeFile(dxwrapperDLLPath)
			if err != nil {
				return err
			}

			err = removeFile(dxwrapperINIPath)
			if err != nil {
				return err
			}

			return nil
		},
		descriptionFunc: func() string {
			return "Fix laggy gameplay"
		},
	}
}
