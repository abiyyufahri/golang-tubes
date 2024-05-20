package _interface

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type modelChoiceModule struct {
	choosen int
	choices []string
}

func initialModelMenu() modelChoiceModule {
	return modelChoiceModule{
		// Our menu module
		choosen: 0,
		choices: []string{"Modul Ekspedisi", "Modul Pelanggan"},
	}
}

func (m modelChoiceModule) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m modelChoiceModule) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "q":
			m.choosen = -1
			return m, tea.Quit

		case "1": // jika pilihan nya 1 atau 2
			m.choosen = 1
			return m, tea.Quit
		case "2":
			m.choosen = 2
			return m, tea.Quit
		}
	}

	// Return the updated modelMenu to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m modelChoiceModule) View() string {
	// The header
	s := "\n\nSilakan Pilih Modul yang ingin diakses \n\n"

	// Iterate over our choices
	for i := 0; i < 2; i++ {

		// Render the row
		s += fmt.Sprintf("%d. %s\n", i+1, m.choices[i])
	}

	// The footer
	s += "\n\nTekan 1 atau 2" +
		"\nTekan q untuk keluar.\n"

	// Send the UI for rendering
	return s
}

func GetModuleChoice() int {
	p := tea.NewProgram(initialModelMenu(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Yaa, there's been an error: %v", err)
		os.Exit(1)
	}
	return m.(modelChoiceModule).choosen
}
