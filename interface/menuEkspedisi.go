package _interface

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

type modelMenuEkspedisi struct {
	choosen int
	choices []string
}

func initialModelMenuEkspedisi(packet string) modelMenuEkspedisi {
	if packet != "" {
		return modelMenuEkspedisi{
			choosen: 0,
			choices: []string{
				"Tambah Ekspedisi",
				"Lihat Daftar Ekspedisi",
				"Lacak Data Ekspedisi",
				"Ubah Status " + packet,
				"Lihat Detail " + packet,
				"Ubah Data " + packet,
				"Hapus Data " + packet,
			},
		}
	}

	return modelMenuEkspedisi{
		// Our menu module
		choosen: 0,
		choices: []string{
			"Tambah Ekspedisi",
			"Lihat Daftar Ekspedisi",
			"Lacak Data Ekspedisi",
			"Ubah Status Ekspedisi",
			"Lihat Detail Ekspedisi",
			"Ubah Data Ekspedisi",
			"Hapus Data Ekspedisi",
		},
	}
}

func (m modelMenuEkspedisi) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m modelMenuEkspedisi) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "0":
			m.choosen = 0
			return m, tea.Quit

		case "1", "2", "3", "4", "5", "6", "7": // jika pilihan nya 1 atau 2 atau seterusnya
			m.choosen, _ = strconv.Atoi(msg.String())
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m modelMenuEkspedisi) View() string {
	// The header
	s := "\n\n==++ Modul Ekspedisi ++== \n\n"

	// Iterate over our choices
	for i := 0; i < 7; i++ {
		// Render the row
		s += fmt.Sprintf("%d. %s\n", i+1, m.choices[i])
	}

	// The footer
	s += "\n\nTekan 1,2,3,4,5, 6 atau 7 untuk memilih" +
		"\nTekan 0 untuk kembali.\n"

	// Send the UI for rendering
	return s
}

func GetModulSubMenuEkspedisi(packet string) int {
	p := tea.NewProgram(initialModelMenuEkspedisi(packet), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Yaa, there's been an error in submenu of ekspedisi: %v", err)
		os.Exit(1)
	}
	return m.(modelMenuEkspedisi).choosen
}
