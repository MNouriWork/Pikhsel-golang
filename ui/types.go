package ui

import (
	"fyne.io/fyne/v2"
	"github.com/mnouriwork/Pikhsel-golang/apptype"
	"github.com/mnouriwork/Pikhsel-golang/swatch"
)

type AppInit struct {
	PikhselWindow fyne.Window
	State         *apptype.State
	Swatches      []*swatch.Swatch
}
