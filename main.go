package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices []string // items on the to-do list
	cursor  int      // which to-do list item our cursor is pointing at
	current string
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Modul Ekspedisi", "Modul Pelanggan"},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			fmt.Println("Terimakasih, sampai jumpa lagi")
			fmt.Println(m.current)
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			pilihan := m.choices[m.cursor]
			m.choices = nil

			fmt.Print(m.current)
			if pilihan == "Modul Ekspedisi" {
				m.choices = []string{"Tambah Ekspedisi", "Lihat Status", "Update Status", "Edit Ekpedisi",
					"Hapus Ekspedisi", "Kembali"}
			} else if pilihan == "Modul Pelanggan" {
				m.choices = []string{"Tambah Pelanggan", "Detail Pelanggan", "Hapus Detail Pelanggan", "Kembali"}
			}

			if pilihan == "Tambah Ekspedisi" {
				m.choices = []string{"Tambah Pelanggan", "Detail Pelanggan", "Hapus Detail Pelanggan", "Kembali"}
			}

			if pilihan == "Kembali" && (m.current == "Modul Ekspedisi" || m.current == "Modul Pelanggan") {
				m.choices = []string{"Modul Ekspedisi", "Modul Pelanggan"}
			}
			m.cursor = 0
			m.current = pilihan
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		//checked := " " // not selected
		//if _, ok := m.selected[i]; ok {
		//	checked = "x" // selected!
		//}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
