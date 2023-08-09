package patch

import (
	"encoding/binary"
	"fmt"

	"fyne.io/fyne/v2"
)

type Patch interface {
	PatchFile([]byte) error
	RemovePatch([]byte) error
	Description() string
}

type PatchFactory interface {
	Build() Patch
	Id() string
	Widget() fyne.CanvasObject
	IsOptional() bool
}

type simplePatch struct {
	patchFileFunc   func(input []byte) error
	removePatchFunc func(input []byte) error
	descriptionFunc func() string
}

func (p simplePatch) PatchFile(input []byte) error {
	return p.patchFileFunc(input)
}

func (p simplePatch) RemovePatch(input []byte) error {
	return p.removePatchFunc(input)
}

func (p simplePatch) Description() string {
	return p.descriptionFunc()
}

func singleAddressPatch(index int, defaultValue uint32, newValue uint32) Patch {
	return simplePatch{
		patchFileFunc: func(input []byte) error {
			writeDWORD(input, index, newValue)
			return nil
		},
		removePatchFunc: func(input []byte) error {
			writeDWORD(input, index, defaultValue)
			return nil
		},
		descriptionFunc: func() string {
			return fmt.Sprintf("Address Patch: %X - %d", index, newValue)
		},
	}
}

func toDWORD(number uint32) []byte {
	output := make([]byte, 4)
	binary.LittleEndian.PutUint32(output, number)
	return output
}

func writeDWORD(b []byte, index int, number uint32) {
	copy(b[index:], toDWORD(number))
}
