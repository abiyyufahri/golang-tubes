package _interface

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

type modelMenuPelanggan struct {
	choosen int
	choices []string
}

func initialModelMenuPelanggan(name string) modelMenuPelanggan {
	if name != "" {
		return modelMenuPelanggan{
			choosen: 0,
			choices: []string{
				"Tambah Pelanggan",
				"Lihat Daftar Pelanggan",
				"Cari Data Pelanggan",
				"Lihat Detail << " + name + " >> ",
				"Ubah Data << " + name + " >> ",
				"Hapus << " + name + " >> ",
			},
		}
	}

	return modelMenuPelanggan{
		// Our menu module
		choosen: 0,
		choices: []string{
			"Tambah Pelanggan",
			"Lihat Daftar Pelanggan",
			"Cari Data Pelanggan",
			"Lihat Detail Pelanggan",
			"Ubah Data Pelanggan",
			"Hapus Pelanggan",
		},
	}
}

func (m modelMenuPelanggan) Init() tea.Cmd {

	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m modelMenuPelanggan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "0":
			m.choosen = 0
			return m, tea.Quit

		case "1", "2", "3", "4", "5", "6": // jika pilihan nya 1 atau 2 atau seterusnya
			m.choosen, _ = strconv.Atoi(msg.String())
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m modelMenuPelanggan) View() string {
	// The header
	s := "\n\n==++ Modul Ekspedisi ++== \n\n"

	// Iterate over our choices
	for i := 0; i < 6; i++ {
		// Render the row
		s += fmt.Sprintf("%d. %s\n", i+1, m.choices[i])
	}

	// The footer
	s += "\n\nTekan 1,2,3 atau 4 untuk memilih" +
		"\nTekan 0 untuk kembali.\n"

	// Send the UI for rendering
	return s
}

func GetModulSubMenuPelanggan(selectedName string) int {
	p := tea.NewProgram(initialModelMenuPelanggan(selectedName), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Yaa, there's been an error in submenu of ekspedisi: %v", err)
		os.Exit(1)
	}
	return m.(modelMenuPelanggan).choosen
}
