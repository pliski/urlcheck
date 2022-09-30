package appState

import (
	"fmt"
	"net/http"
)

const (
	OkColor      = "\033[1;32m%s\033[0m"
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func (s *State) View() string {

	// Tell the user we're doing something.
	msg := "Connecting ... \n"
	status := 0
	var premsg string

	var color string

	for i := 0; i < s.count; i++ {
		status = s.Status[i]
		if status == 0 {
			color = ErrorColor
		} else if status < 204 {
			color = OkColor
		} else if status < 400 {
			color = WarningColor
		} else {
			color = ErrorColor
		}

		premsg = fmt.Sprintf("[%3d %2s]\t%s\n", status, http.StatusText(status), s.urilist[i])
		msg += fmt.Sprintf(color, premsg)
	}

	if len(s.Err) > 0 {
		msg += "\n\033[0;36mWe had some trouble: \n"
		for n := 0; n < len(s.Err); n++ {
			msg += fmt.Sprintf("%v\n", s.Err[n])
		}
		msg += "\033[0m"
	}

	// Send off whatever we came up with above for rendering.
	return "\n" + msg + "\n\n"
}
