package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
)

type App struct {
	Driver gxui.Driver
	Theme  AppTheme
	Window AppWindow
	Client *Client
}

type AppTheme struct {
	Driver gxui.Driver
	theme  gxui.Theme
}

func (at *AppTheme) init(driver gxui.Driver) {
	at.Driver = driver
	at.theme = dark.CreateTheme(driver)
}

func (at *AppTheme) CreateWindow(W int, H int, title string) gxui.Window {
	return at.theme.CreateWindow(W, H, title)
}

func (at *AppTheme) createInputElement(defaultText string) gxui.TextBox {
	input := at.theme.CreateTextBox()
	input.SetText(defaultText)
	return input
}

func (at *AppTheme) CreateLabel(defaultText string) gxui.Label {
	input := at.theme.CreateLabel()
	input.SetText(defaultText)
	return input
}

func (at *AppTheme) CreateButton(defaultText string) gxui.Button {
	input := at.theme.CreateButton()
	input.SetText(defaultText)
	return input
}
func (at *AppTheme) CreateLinearLayout() gxui.LinearLayout {
	input := at.theme.CreateLinearLayout()
	return input
}

func (at *AppTheme) CreateTextBox(defaultText string, maxWidth int) gxui.TextBox {
	input := at.theme.CreateTextBox()
	input.SetDesiredWidth(maxWidth)
	return input
}

type AppWindow struct {
	Theme           AppTheme
	Window          gxui.Window
	ResultArea      ResultForm
	CredentialsForm CredentialsForm
	ExecutionForm   ExecutionForm
	DisconnectForm  DisconnectForm
}

func (aw *AppWindow) updateResultText(newText string) {
	aw.ResultArea.update(newText)
}

func (app *App) ConnectAction(ev gxui.MouseEvent) {

	app.Window.CredentialsForm.layout.SetVisible(false)
	app.Window.ExecutionForm.layout.SetVisible(true)
	app.Window.DisconnectForm.layout.SetVisible(true)
	client := app.Window.CredentialsForm.connect()
	app.Client = &client
	app.Window.updateResultText("Connected")

}
func (app *App) ExecuteAction(ev gxui.MouseEvent) {
	cmd := app.Window.ExecutionForm.command.Text()
	result := app.Client.Execute(cmd)
	app.Window.updateResultText(result)

}
func (app *App) DisconnectAction(ev gxui.MouseEvent) {
	app.Window.CredentialsForm.layout.SetVisible(true)
	app.Window.DisconnectForm.layout.SetVisible(false)
	app.Window.ExecutionForm.layout.SetVisible(false)
	app.Client.Close()
	app.Window.updateResultText("Disconnected")

}

func (app *App) createWindow(W int, H int, title string) {
	window := app.Theme.CreateWindow(W, H, title)

	app.Window = AppWindow{
		app.Theme,
		window,
		ResultForm{},
		CredentialsForm{},
		ExecutionForm{},
		DisconnectForm{},
	}

	app.Window.CredentialsForm.createLayout(app.Theme)
	app.Window.CredentialsForm.button.OnClick(app.ConnectAction)

	app.Window.ExecutionForm.createLayout(app.Theme)
	app.Window.ExecutionForm.button.OnClick(app.ExecuteAction)

	app.Window.DisconnectForm.createLayout(app.Theme)
	app.Window.DisconnectForm.button.OnClick(app.DisconnectAction)

	app.Window.ResultArea.createLayout(app.Theme)
}

func (app *App) buildWindow() {
	label := app.Theme.CreateLabel("Enter command")

	layout := app.Theme.CreateLinearLayout()
	layout.AddChild(label)
	layout.AddChild(app.Window.CredentialsForm.layout)
	layout.AddChild(app.Window.DisconnectForm.layout)
	layout.AddChild(app.Window.ExecutionForm.layout)
	layout.AddChild(app.Window.ResultArea.layout)
	layout.SetHorizontalAlignment(gxui.AlignCenter)
	layout.SetVerticalAlignment(gxui.AlignMiddle)

	app.Window.Window.AddChild(layout)
	app.Window.Window.BorderPen()
	app.Window.Window.SetScale(flags.DefaultScaleFactor)
	app.Window.Window.OnClose(app.terminate)
}

func (app *App) terminate() {
	if app.Client != nil {
		app.Client.Close()
	}
	app.Theme.Driver.Terminate()
}
