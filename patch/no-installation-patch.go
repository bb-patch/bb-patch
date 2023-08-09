package patch

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/sys/windows/registry"
)

type langauge uint32

const (
	English langauge = iota
	German
	French
	Italian
	Spanish
)

type noInstallPatchFactory struct {
	lang langauge
}

func (pf *noInstallPatchFactory) Id() string {
	return "NoInstallation"
}

func (pf *noInstallPatchFactory) IsOptional() bool {
	return false
}

func (pf *noInstallPatchFactory) Build() Patch {
	return createNoInstallPatch(pf.lang)
}

func (pf *noInstallPatchFactory) Widget() fyne.CanvasObject {
	options := map[string]langauge{
		"English": English,
		"German":  German,
		"French":  French,
		"Italian": Italian,
		"Spanish": Spanish,
	}

	keys := make([]string, 0, 5)
	for k := range options {
		keys = append(keys, k)
	}

	selection := widget.NewSelect(keys, func(s string) {
		pf.lang = options[s]
	})
	selection.SetSelected("English")

	return container.NewVBox(
		widget.NewLabel("Required. Alters the windows registry so that the game think it has been installed."),
		widget.NewLabel("Chose the language you want the game in:"),
		selection,
	)
}

func createNoInstallPatch(lang langauge) Patch {
	var location, _ = os.Getwd()
	location += "\\"

	keyPath := "SOFTWARE\\Worldweaver Productions\\Beasts and Bumpkins"

	stringKeys := map[string]string{
		"Audio":        location + "audio.box",
		"CD":           location,
		"InstallDir":   location[2:],
		"InstallDrive": location[:2],
		"Link":         "",
		"Misc":         location + "misc.box",
		"Missions":     location + "missions.box",
		"MissionText":  location + "mistext0.box",
		"Speech":       location + "speech0.box",
		"Video":        location + "video.box",
		"Working":      location,
	}

	intKeys := map[string]uint32{
		"Installed": 1,
		"Language":  uint32(lang),
	}

	removeRegistry := func() error {
		err := registry.DeleteKey(registry.CURRENT_USER, keyPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	return simplePatch{
		patchFileFunc: func(_ []byte) error {
			err := removeRegistry()
			if err != nil {
				return err
			}

			key, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, 0666)
			if err != nil {
				return err
			}
			defer key.Close()

			for k, v := range stringKeys {
				err := key.SetStringValue(k, v)
				if err != nil {
					return err
				}
			}

			for k, v := range intKeys {
				err := key.SetDWordValue(k, uint32(v))
				if err != nil {
					return err
				}
			}
			return nil
		},
		removePatchFunc: func(_ []byte) error {
			return removeRegistry()
		},
		descriptionFunc: func() string {
			return "Set the required registry entries"
		},
	}
}
