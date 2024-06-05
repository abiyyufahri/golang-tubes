package pelanggan

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type modelTable struct {
	table         table.Model
	dataPelanggan *ModelPelanggan
}

func (m *modelTable) Init() tea.Cmd {
	m.table.SetRows(m.dataPelanggan.ToTableRow())
	return nil
}

func (m *modelTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1": // urut nama menaik
			m.dataPelanggan.SortByNameAscending()
			m.table.SetRows(m.dataPelanggan.ToTableRow())
			return m, nil
		case "2": // urut nama menurun
			m.dataPelanggan.sortByNameDescending()
			m.table.SetRows(m.dataPelanggan.ToTableRow())
			return m, nil
		case "3": // urut id menaik
			m.dataPelanggan.SortByIdAscending()
			m.table.SetRows(m.dataPelanggan.ToTableRow())
			return m, nil
		case "4": // urut id menurun
			m.dataPelanggan.SortByIdDescending()
			m.table.SetRows(m.dataPelanggan.ToTableRow())
			return m, nil
		case "5":
			actives := m.dataPelanggan.filterByActive()
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "6":
			actives := m.dataPelanggan.filterByNonActive()
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "7":
			nonActives := m.dataPelanggan.GetAll()
			m.table.SetRows(nonActives.ToTableRow())
			return m, nil
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "enter":
			selected := m.table.SelectedRow()
			selectedId, _ := strconv.Atoi(selected[0])
			m.dataPelanggan.setSelectedId(selectedId)
			fmt.Println("\n\n\n >>> Kamu memilih", selected[1], "                       \n tunggu beberapa saat")
			time.AfterFunc(3*time.Second, func() {
				// do nothing

			})
			time.Sleep(3 * time.Second)
			return m, tea.Quit
		case "q", "ctrl+c", "0":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *modelTable) View() string {
	return baseStyle.Render(m.table.View()) + "\n" +
		"Pilih pelanggan untuk di-keep\n" +
		"Menu sorting/filter (klik keyboard): \n" +
		"1. Urut nama menaik \n" +
		"2. Urut nama menurun\n" +
		"3. Urut id menaik\n" +
		"4. Urut id menurun\n" +
		"5. Tampilkan pelanggan aktif saja\n" +
		"6. Tampilkan pelanggan non-aktif saja\n" +
		"7. Tampilkan pelanggan semua status"
}

func viewAllTable(dp *ModelPelanggan) {
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

	m := modelTable{t, dp}
	if _, err := tea.NewProgram(&m, tea.WithAltScreen()).Run(); err != nil {
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
