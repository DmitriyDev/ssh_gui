package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type DisconnectForm struct {
	button gxui.Button
	layout gxui.LinearLayout
}

func (df *DisconnectForm) createLayout(theme AppTheme) {

	df.button = theme.CreateButton("Disconnect")

	df.layout = theme.CreateLinearLayout()

	df.layout.SetSize(math.Size{600, 200})
	df.layout.AddChild(df.button)
	df.layout.SetVisible(false)

}
