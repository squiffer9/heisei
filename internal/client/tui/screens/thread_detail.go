package screens

import (
	"fmt"
	"heisei/internal/client/api"
	"heisei/internal/common/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

type ThreadDetail struct {
	*tview.Flex
	api           *api.Client
	logger        *zap.Logger
	postsList     *tview.TextView
	inputField    *tview.InputField
	currentThread *models.ThreadDTO
}

func NewThreadDetail(api *api.Client, logger *zap.Logger) *ThreadDetail {
	td := &ThreadDetail{
		Flex:   tview.NewFlex().SetDirection(tview.FlexRow),
		api:    api,
		logger: logger,
	}

	td.postsList = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(true)

	td.inputField = tview.NewInputField().
		SetLabel("New post: ").
		SetFieldWidth(0)

	td.Flex.AddItem(td.postsList, 0, 1, false).
		AddItem(td.inputField, 1, 0, true)

	td.SetBorder(true)

	return td
}

func (td *ThreadDetail) LoadPosts(thread *models.ThreadDTO) error {
	td.currentThread = thread
	td.SetTitle(fmt.Sprintf("Thread: %s", thread.Title))

	posts, err := td.api.GetPostsByThread(thread.ID)
	if err != nil {
		td.logger.Error("Failed to load posts", zap.Error(err), zap.Uint("threadID", thread.ID))
		return err
	}

	td.postsList.Clear()
	for _, post := range posts {
		fmt.Fprintf(td.postsList, "[yellow]%s[white]\n%s\n\n", post.CreatedAt.Format("2006-01-02 15:04:05"), post.Content)
	}

	return nil
}

func (td *ThreadDetail) SetSubmitFunc(fn func(string)) {
	td.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			text := td.inputField.GetText()
			if text != "" {
				fn(text)
				td.inputField.SetText("")
			}
		}
	})
}

func (td *ThreadDetail) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) {
	td.Flex.SetInputCapture(capture)
}

func (td *ThreadDetail) AddPost(post *models.PostDTO) {
	fmt.Fprintf(td.postsList, "[yellow]%s[white]\n%s\n\n", post.CreatedAt.Format("2006-01-02 15:04:05"), post.Content)
	td.postsList.ScrollToEnd()
}
