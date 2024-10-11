package screens

import (
	"fmt"
	"heisei/internal/client/api"
	"heisei/internal/common/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

type ThreadList struct {
	*tview.List
	api    *api.Client
	logger *zap.Logger
}

func NewThreadList(api *api.Client, logger *zap.Logger) *ThreadList {
	tl := &ThreadList{
		List:   tview.NewList().ShowSecondaryText(true),
		api:    api,
		logger: logger,
	}
	tl.SetBorder(true).SetTitle("Threads")
	return tl
}

func (tl *ThreadList) LoadThreads(categoryID uint) error {
	threads, err := tl.api.GetThreadsByCategory(categoryID)
	if err != nil {
		tl.logger.Error("Failed to load threads", zap.Error(err), zap.Uint("categoryID", categoryID))
		return err
	}

	tl.Clear()
	for _, thread := range threads {
		tl.AddItem(thread.Title, fmt.Sprintf("Posts: %d", thread.PostCount), 0, func() {
			// TODO: Implement thread selection
		})
	}

	return nil
}

func (tl *ThreadList) SetSelectedFunc(fn func(*models.ThreadDTO)) {
	tl.List.SetSelectedFunc(func(index int, name string, secondaryText string, shortcut rune) {
		threads, _ := tl.api.GetThreadsByCategory(0) // TODO: Store current category ID
		if index < len(threads) {
			fn(&threads[index])
		}
	})
}

func (tl *ThreadList) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) {
	tl.List.SetInputCapture(capture)
}
