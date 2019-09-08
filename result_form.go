package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type ResultForm struct {
	result gxui.TextBox
	layout gxui.LinearLayout
}

func (c *ResultForm) update(newText string) {
	text := c.result.Text()
	c.result.SetText(text + "\r\n " + newText)
}

func (rf *ResultForm) createLayout(theme AppTheme) {
	rf.result = theme.CreateTextBox("", math.MaxSize.W)
	rf.layout = theme.CreateLinearLayout()
	rf.layout.SetSize(math.Size{600, 200})
	rf.layout.AddChild(rf.result)
}
