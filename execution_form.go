package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type ExecutionForm struct {
	command gxui.TextBox
	button  gxui.Button
	layout  gxui.LinearLayout
}

func (ef *ExecutionForm) createLayout(theme AppTheme) {

	ef.command = theme.CreateTextBox("", math.MaxSize.W)
	ef.button = theme.CreateButton("Execute")

	executeLayout := theme.CreateLinearLayout()

	executeLayout.SetSize(math.Size{600, 200})
	executeLayout.AddChild(ef.command)
	executeLayout.AddChild(ef.button)
	executeLayout.SetVisible(false)

	ef.layout = executeLayout

}
