package action

import (
	"net/http"
	"time"
	"urlcheck/model"

	tea "github.com/charmbracelet/bubbletea"
)

func CheckSomeUrl(url string) tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(url)
		if err != nil {
			return ErrMsg{err, false}
		}
		return StatusMsg(res.StatusCode)
	}
}

func CheckSomeHeaders(url string, client *ClientModel[model.TeaLogger]) tea.Cmd {
	return func() tea.Msg {
		status := getStatus(url, time.Duration(10)*time.Second, client)
		if status == 999 {
			status = 0
		}
		return StatusMsg(status)
	}
}

type StatusMsg int

type ErrMsg struct {
	err      error
	Critical bool
}

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e ErrMsg) Error() string { return e.err.Error() }
