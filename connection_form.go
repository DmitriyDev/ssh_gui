package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
	"strconv"
)

type CredentialsForm struct {
	ip       gxui.TextBox
	user     gxui.TextBox
	password gxui.TextBox
	port     gxui.TextBox
	button   gxui.Button
	layout   gxui.LinearLayout
}

func (c *CredentialsForm) connect() Client {
	connectionPort, _ := strconv.Atoi(c.port.Text())
	return Connect(
		c.ip.Text(),
		c.user.Text(),
		c.password.Text(),
		connectionPort,
	)

}

func (c *CredentialsForm) createLayout(theme AppTheme) {
	c.buildForm(theme)

	btnLayout := theme.CreateLinearLayout()
	btnLayout.SetSize(math.Size{600, 200})
	btnLayout.AddChild(c.ip)
	btnLayout.AddChild(c.user)
	btnLayout.AddChild(c.password)
	btnLayout.AddChild(c.port)
	btnLayout.AddChild(c.button)
	c.layout = btnLayout
}

func (c *CredentialsForm) buildForm(theme AppTheme) {
	c.ip = theme.createInputElement("")
	c.user = theme.createInputElement("")
	c.password = theme.createInputElement("")
	c.port = theme.createInputElement("22")
	c.button = theme.CreateButton("Connect")
}
