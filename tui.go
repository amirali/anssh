package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/ini.v1"
)

type model struct {
	groups        []string
	hosts         map[string][]string
	cursor        int
	selectedGroup string
}

func initModel(inventory *ini.File) model {
	hosts := make(map[string][]string)
	for _, section := range inventory.SectionStrings() {
		hosts[section] = extarctHost(inventory.Section(section).KeyStrings()...)
	}
	return model{
		groups: inventory.SectionStrings(),
		hosts:  hosts,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			var length int
			if m.selectedGroup != "" {
				length = len(m.hosts[m.selectedGroup])
			} else {
				length = len(m.groups)
			}
			if m.cursor < length-1 {
				m.cursor++
			}

		case "enter":
			if m.selectedGroup == "" {
				m.selectedGroup = m.groups[m.cursor]
				m.cursor = 0
			} else {
				config.targetHost = m.hosts[m.selectedGroup][m.cursor]
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Up/Down and J/K to move, Enter to select, Q and Ctrl+C to quit\n\n"

	switch m.selectedGroup {
	case "":
		for i, group := range m.groups {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, group)
		}
	default:
		s += fmt.Sprintf("[%s]\n\n", m.selectedGroup)
		for i, host := range m.hosts[m.selectedGroup] {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, host)
		}
	}

	return s
}
