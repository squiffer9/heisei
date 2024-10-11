package screens

import (
	"heisei/internal/client/api"
	"heisei/internal/common/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

type CategoryList struct {
	*tview.List
	api    *api.Client
	logger *zap.Logger
}

func NewCategoryList(api *api.Client, logger *zap.Logger) *CategoryList {
	cl := &CategoryList{
		List:   tview.NewList().ShowSecondaryText(false),
		api:    api,
		logger: logger,
	}
	cl.SetBorder(true).SetTitle("Categories")
	return cl
}

func (cl *CategoryList) LoadCategories() error {
	categories, err := cl.api.GetCategories()
	if err != nil {
		cl.logger.Error("Failed to load categories", zap.Error(err))
		return err
	}

	cl.Clear()
	for _, category := range categories {
		cl.AddItem(category.Name, "", 0, func() {
			// TODO: Implement category selection
		})
	}

	return nil
}

func (cl *CategoryList) SetSelectedFunc(fn func(*models.CategoryDTO)) {
	cl.List.SetSelectedFunc(func(index int, name string, secondaryText string, shortcut rune) {
		categories, _ := cl.api.GetCategories()
		if index < len(categories) {
			fn(&categories[index])
		}
	})
}

func (cl *CategoryList) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) {
	cl.List.SetInputCapture(capture)
}
