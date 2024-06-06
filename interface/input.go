package _interface

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func GetInputString(pertanyaan, placeholder string) string {
	p := tea.NewProgram(initialModelInput(pertanyaan, placeholder), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	return m.(modelInput).textInput.Value()
}

func GetInputInteger(pertanyaan, placeholder string) int {
	p := tea.NewProgram(initialModelInput(pertanyaan, placeholder), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	inputStr := m.(modelInput).textInput.Value()
	inputInt, _ := strconv.Atoi(inputStr) // conversi string ke integer
	if err != nil {                       // Jika gagal
		fmt.Println("Nilai yang tidak valid, harap lakukan pengeditan terlebih dahulu.")
		return 0
	}
	return inputInt
}

type (
	errMsg error
)

type modelInput struct {
	textInput  textinput.Model
	pertanyaan string
	err        error
}

func initialModelInput(pertanyaan, placeholder string) modelInput {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return modelInput{
		textInput:  ti,
		pertanyaan: pertanyaan,
		err:        nil,
	}
}

func (m modelInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m modelInput) View() string {
	return fmt.Sprintf(
		"%s?\n%s\n\n%s",
		m.pertanyaan,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
