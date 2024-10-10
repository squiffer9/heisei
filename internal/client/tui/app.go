package tui

import (
	"heisei/internal/client/api"
	"heisei/internal/client/config"

	"github.com/rivo/tview"
	"go.uber.org/zap"
)

type App struct {
	*tview.Application
	Config    *config.Config
	APIClient *api.Client
	Logger    *zap.Logger

	// Main layout
	mainFlex *tview.Flex
}

func NewApp(cfg *config.Config, apiClient *api.Client, logger *zap.Logger) (*App, error) {
	app := &App{
		Application: tview.NewApplication(),
		Config:      cfg,
		APIClient:   apiClient,
		Logger:      logger,
	}

	if err := app.initUI(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initUI() error {
	a.mainFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	// Add a placeholder text view
	textView := tview.NewTextView().
		SetText("Welcome to Heisei BBS").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	a.mainFlex.AddItem(textView, 0, 1, true)
	a.SetRoot(a.mainFlex, true)

	return nil
}

func (a *App) Run() error {
	return a.Application.Run()
}
