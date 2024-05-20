package pelanggan

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type modelTable struct {
	table table.Model
}

func (m modelTable) Init() tea.Cmd { return nil }

func (m modelTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m modelTable) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func viewAllTable(dp ModelPelanggan) {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Nama", Width: 20},
		{Title: "Status", Width: 6},
		{Title: "No. Telp", Width: 14},
		{Title: "Email", Width: 40},
		{Title: "Alamat", Width: 30},
	}

	rows := dp.ToTableRow()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := modelTable{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program on showing pelanggan's table:", err)
		os.Exit(1)
	}
}

func (p *ModelPelanggan) ToTableRow() []table.Row {
	var rows []table.Row
	var pelanggan Pelanggan

	var booleanToString = map[bool]string{
		true:  "Aktif",
		false: "Nonaktif",
	}

	for i := 0; i < p.nPelanggan; i++ {
		pelanggan = p.daftarPelanggan[i]
		row := table.Row{
			strconv.Itoa(pelanggan.id),
			pelanggan.nama,
			booleanToString[pelanggan.status],
			pelanggan.nomorTelepon,
			pelanggan.alamatEmail,
			pelanggan.alamat,
		}
		rows = append(rows, row) // table row tidak dapat digunakan dengan statis
	}
	return rows
}
