package ekspedisi

import (
	p "TUBES_GO/pelanggan"
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
	dataEkspedisi *ModelEkspedisi
}

func (m *modelTable) Init() tea.Cmd {
	return nil
}

func (m *modelTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1": // urut resi terlama
			m.dataEkspedisi.SortByResiAscending()
			m.table.SetRows(m.dataEkspedisi.ToTableRow())
			return m, nil
		case "2": // urut resi terbaru
			m.dataEkspedisi.SortByResiDescending()
			m.table.SetRows(m.dataEkspedisi.ToTableRow())
			return m, nil
		case "3": // urut prioritas tertinggi
			m.dataEkspedisi.SortByPriorityDescending()
			m.table.SetRows(m.dataEkspedisi.ToTableRow())
			return m, nil
		case "4": // urut prioritas terendah
			m.dataEkspedisi.SortByPriorityAscending()
			m.table.SetRows(m.dataEkspedisi.ToTableRow())
			return m, nil
		case "5": // instant
			actives := m.dataEkspedisi.FilterByJenisLayanan(Instant)
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "6": // reguler
			actives := m.dataEkspedisi.FilterByJenisLayanan(Reguler)
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "7": // cargo
			actives := m.dataEkspedisi.FilterByJenisLayanan(Cargo)
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "8": // sameday
			actives := m.dataEkspedisi.FilterByJenisLayanan(SameDay)
			m.table.SetRows(actives.ToTableRow())
			return m, nil
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "enter":
			selected := m.table.SelectedRow()
			m.dataEkspedisi.selectedResi = selected[0]
			fmt.Println("\n\n\n >>> Kamu memilih", selected[0], " "+
				" milik pelanggan"+selected[1]+
				"                        \n tunggu beberapa saat")
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
		"1. Urut dari yang terlama \n" +
		"2. Urut dari yang terbaru\n" +
		"3. Urut prioritas tertinggi\n" +
		"4. Urut prioritas terendah\n" +
		"5. Tampilkan ekspedisi instant\n" +
		"6. Tampilkan ekspedisi reguler\n" +
		"7. Tampilkan ekspedisi cargo\n" +
		"8. Tampilkan ekspedisi sameday\n" +
		"0. Untuk kembali"
}

func viewAllTable(de *ModelEkspedisi, dp p.ModelPelanggan) {
	columns := []table.Column{
		{Title: "Resi", Width: 20},
		{Title: "Pelanggan", Width: 10},
		{Title: "Status", Width: 10},
		{Title: "Layanan", Width: 9},
		{Title: "Alamat Asal", Width: 20},
		{Title: "Alamat Tujuan", Width: 20},
	}

	rows := de.ToTableRow()

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

	m := modelTable{t, de}
	if _, err := tea.NewProgram(&m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program on showing pelanggan's table:", err)
		os.Exit(1)
	}
}

func (e *ModelEkspedisi) ToTableRow( /*dp p.ModelPelanggan*/ ) []table.Row {
	var rows []table.Row
	var ekspedisi Ekspedisi

	var statusToString = map[int]string{
		1: "DiKemas",
		2: "DiJemput",
		3: "DiAntar",
		4: "Selesai",
	}

	var jenisLayananToString = map[Layanan]string{
		1: "Regular",
		2: "SameDay",
		3: "Instant",
		4: "Cargo",
	}

	idInt, _ := strconv.ParseInt(ekspedisi.idPelanggan, 10, 64)
	for i := 0; i < e.nEkspedisi; i++ {
		ekspedisi = e.DaftarEkspedisi[i]
		if getNamaPelanggan( /*dp,*/ int(idInt)) == "Not Found" {
			ekspedisi.butuhDilengkapi = true
		}
		row := table.Row{
			ekspedisi.resi,
			ekspedisi.idPelanggan,
			statusToString[ekspedisi.status],
			jenisLayananToString[ekspedisi.jenisLayanan],
			ekspedisi.alamatAsal,
			ekspedisi.alamatTujuan,
		}
		rows = append(rows, row) // table row tidak dapat digunakan dengan statis
	}
	return rows
}

func getNamaPelanggan( /*dp p.ModelPelanggan,*/ id int) string {
	//return dp.GetNamaById(id)
	return ""
}
