package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# Grimoire
## GitFlow
* git ci - Enter commit message interactively
* git ps - Push the current branch to the remote
* git feat NAME - Switch a new branch from the current branch (feat/NAME)
* git fix NAME - git switch -c fix/NAME
* git hotfix NAME - git switch -c hotfix/NAME
* git docs NAME - git switch -c docs/NAME
* git style NAME - git switch -c style/NAME
* git refactor NAME - git switch -c refactor/NAME
* git chore NAME - git switch -c chore/NAME
* git perf NAME - git switch -c perf/NAME
* git style NAME - git switch -c style/NAME
## .gitignore
* goignore - generate .gitignore
## GitHub
* gh dash
## Type speed
* typer
## Games
* tetris
`

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type example struct {
	viewport viewport.Model
}

func newExample() (*example, error) {
	const width = 100

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &example{
		viewport: vp,
	}, nil
}

func (e example) Init() tea.Cmd {
	return nil
}

func (e example) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return e, tea.Quit
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}

func (e example) View() string {
	return e.viewport.View() + e.helpView()
}

func (e example) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}

func main() {
	model, err := newExample()
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}
