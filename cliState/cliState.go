package appState

import (
	"urlcheck/action"
	model "urlcheck/model"

	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Status  []int
	Err     []string
	count   int
	urilist []string
	client  *action.ClientModel[model.TeaLogger]
}

func NewModel(urilist []string) *State {

	stati := make([]int, len(urilist))
	errs := []string{}

	initialModel := State{
		Status:  stati,
		Err:     errs,
		count:   0,
		urilist: urilist,
		client:  action.NewClientTea(&errs),
	}
	return &initialModel
}

func (s State) Init() tea.Cmd {
	// return action.CheckSomeUrl(s.urilist[s.count])
	return action.CheckSomeHeaders(s.urilist[s.count], s.client)
}

func (s *State) nextOrDie() (tea.Model, tea.Cmd) {
	s.count += 1
	if s.count < len(s.urilist) {
		return s, action.CheckSomeHeaders(s.urilist[s.count], s.client)
		// return s, action.CheckSomeUrl(s.urilist[s.count])
	}
	return s, tea.Quit
}

func (s *State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case action.StatusMsg:
		s.Status[s.count] = int(msg)
		return s.nextOrDie()

	// TODO this msg type will disappear
	case action.ErrMsg:
		s.Err = append(s.Err, msg.Error())
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
