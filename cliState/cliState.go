package appState

import (
	"urlcheck/action"

	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Status  []int
	Err     []error
	count   int
	urilist []string
}

func NewModel(urilist []string) *State {

	stati := make([]int, len(urilist))
	errs := []error{}

	initialModel := State{
		Status:  stati,
		Err:     errs,
		count:   0,
		urilist: urilist,
	}
	return &initialModel
}

func (s State) Init() tea.Cmd {
	return action.CheckSomeUrl(s.urilist[s.count])
}

func (s *State) nextOrDie() (tea.Model, tea.Cmd) {
	s.count += 1
	if s.count < len(s.urilist) {
		return s, action.CheckSomeUrl(s.urilist[s.count])
	}
	return s, tea.Quit
}

func (s *State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case action.StatusMsg:
		s.Status[s.count] = int(msg)
		return s.nextOrDie()

	case action.ErrMsg:
		s.Err = append(s.Err, msg)
		if msg.Critical {
			return s, tea.Quit
		}
		return s.nextOrDie()

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return s, tea.Quit
		}
	}

	// If we happen to get any other messages, don't do anything.
	return s, nil
}
