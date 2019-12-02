package main

import (
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard"
)

func (app *upApplication) ShowWaiter(text string, a func(func(string)), b func()) {
	label := frenyard.NewUILabelPtr(frenyard.NewTextTypeChunk("", design.GlobalFont), design.ThemeText, 0, frenyard.Alignment2i{})
	app.slideContainer.TransitionTo(design.LayoutDocument(design.Header{
		Title: text,
	}, label, false), 1.0, false, false)
	go func () {
		a(func (text string) {
			app.upQueued <- func () {
				label.SetText(frenyard.NewTextTypeChunk(text, design.GlobalFont))
			}
		})
		app.upQueued <- b
	}()
}