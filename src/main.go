package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialize() (tea.Model, tea.Cmd) {
	m := model{
		choices:  []string{"multi-cloud", "single-cloud", "elastic-cloud", "elastic-on-prem"},
		selected: make(map[int]struct{}),
	}
	return m, nil
}

func update(msg tea.Msg, mdl tea.Model) (tea.Model, tea.Cmd) {
	m, _ := mdl.(model)
	switch msg := msg.(type) {
	case tead.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func view(mdl tea.Model) string {
	m, _ := mdl.(model)
	s := "Choose your cloud strategy."
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %sn", cursor, checked, choice)
	}
	s += "nPress q to quit.n"
	return s
}

func ain() {
	p := tea.NewProgram(initialize, update, view)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, thereś been an error: %v", err)
		os.Exit(1)
	}
}
