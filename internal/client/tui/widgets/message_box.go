package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MessageBox struct {
	*tview.Modal
}

func NewMessageBox() *MessageBox {
	m := &MessageBox{
		Modal: tview.NewModal(),
	}
	m.SetBackgroundColor(tcell.ColorBlack)
	m.SetTextColor(tcell.ColorWhite)
	m.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 0 {
			m.Hide()
		}
	})
	return m
}

func (m *MessageBox) SetMessage(message string) *MessageBox {
	m.SetText(message)
	return m
}

func (m *MessageBox) SetTitle(title string) *MessageBox {
	m.Modal.SetTitle(title)
	return m
}

func (m *MessageBox) Show(app *tview.Application, pages *tview.Pages) {
	m.AddButtons([]string{"OK"})
	pages.AddPage("message", m, true, true)
	app.SetFocus(m)
}

func (m *MessageBox) Hide() {
	m.Modal.ClearButtons()
}

func ShowError(app *tview.Application, pages *tview.Pages, message string) {
	NewMessageBox().
		SetTitle("Error").
		SetMessage(message).
		Show(app, pages)
}

func ShowInfo(app *tview.Application, pages *tview.Pages, message string) {
	NewMessageBox().
		SetTitle("Information").
		SetMessage(message).
		Show(app, pages)
}
