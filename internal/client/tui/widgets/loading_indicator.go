package widgets

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type LoadingIndicator struct {
	*tview.TextView
	stopChan chan struct{}
}

func NewLoadingIndicator() *LoadingIndicator {
	li := &LoadingIndicator{
		TextView: tview.NewTextView().
			SetDynamicColors(true).
			SetTextAlign(tview.AlignCenter).
			SetText("Loading..."),
		stopChan: make(chan struct{}),
	}
	li.SetBorder(true)
	li.SetTitle("Please wait")
	li.SetBackgroundColor(tcell.ColorBlack)
	li.SetTextColor(tcell.ColorWhite)
	return li
}

func (li *LoadingIndicator) Start(app *tview.Application) {
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		for {
			select {
			case <-li.stopChan:
				return
			default:
				for _, frame := range frames {
					li.SetText(fmt.Sprintf("%s Loading...", frame))
					app.Draw()
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()
}

func (li *LoadingIndicator) Stop() {
	close(li.stopChan)
}

func (li *LoadingIndicator) Show(app *tview.Application, pages *tview.Pages) {
	pages.AddPage("loading", li, true, true)
	app.SetFocus(li)
	li.Start(app)
}

func (li *LoadingIndicator) Hide(pages *tview.Pages) {
	li.Stop()
	pages.RemovePage("loading")
}
