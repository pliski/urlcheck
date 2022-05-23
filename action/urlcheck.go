package action

import (
	"net/http"
	"time"

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

type StatusMsg int

type ErrMsg struct {
	err      error
	Critical bool
}

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e ErrMsg) Error() string { return e.err.Error() }
