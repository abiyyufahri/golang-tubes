package ekspedisi

import (
	p "TUBES_GO/pelanggan"
	"fmt"
	"strconv"
	"time"
)

const NMAX_EKSPEDISI int = 20

type Layanan int

const (
	Reguler Layanan = iota + 1
	SameDay
	Instant
	Cargo
)

type Ekspedisi struct {
	resi            string
	idPelanggan     string
	jenisLayanan    Layanan
	deskripsiBarang string
	alamatAsal      string
	alamatTujuan    string

	biayaEkspedisi  int
	status          int
	butuhDilengkapi bool
}

type TabEkspedisi [NMAX_EKSPEDISI]Ekspedisi

type ModelEkspedisi struct {
	selectedResi    string // nomor ekspedisi yang dipilih
	DaftarEkspedisi TabEkspedisi
	nomorEkspedisi  int
	nEkspedisi      int
}

func (e *ModelEkspedisi) Init() {
	/*
		I.S. model ekspedisi, masih kosong
		F.S. membuat inisialisasi state dari model
	*/
	e.selectedResi = "-1" // tidak ada id yang dipilih
	e.nomorEkspedisi = 1  // nomor dimulai dari 1
}

func (e *ModelEkspedisi) Create() {

	/*
		Membuat data pelanggan baru, return false bila data telah penuh
	*/
	if e.nEkspedisi < NMAX_EKSPEDISI {

		var i = e.nEkspedisi
		var pelanggan Ekspedisi
		var isSuccess bool
		isSuccess, pelanggan = create_form()

		if isSuccess {
			e.DaftarEkspedisi[i] = pelanggan
			e.DaftarEkspedisi[i].resi = generateResiCode()
			e.nEkspedisi++
			e.nomorEkspedisi++
			e.DaftarEkspedisi[i].status = 1
		}
		return
	}

	content := "Data Ekspedisi telah penuh "
	show_pager(content)
}

func generateResiCode() string {
	now := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
}

func (e *ModelEkspedisi) ReadAll(dp p.ModelPelanggan) {
	viewAllTable(e, dp)
}

func (e *ModelEkspedisi) Read() {
	var content string
	if e.selectedResi == "-1" {
		content = "Data ekspedisi tidak ditemukan,"
		show_pager(content)
		return
	}

	var statusToString = map[int]string{
		1: "Dikemas",
		2: "Dijemput",
		3: "Diantar",
		4: "Selesai",
	}

	var jenisLayananToString = map[Layanan]string{
		1: "Regular",
		2: "SameDay",
		3: "Instant",
		4: "Cargo",
	}

	var ekspedisi Ekspedisi
	ekspedisi = e.DaftarEkspedisi[e.searchBySelectedResi()]
	content = "Resi :" + ekspedisi.resi + "\n" +
		"Biaya : " + strconv.Itoa(ekspedisi.biayaEkspedisi) + "\n\n" +
		"Pelanggan id : " + ekspedisi.idPelanggan + "\n\n" +
		"Status : " + statusToString[ekspedisi.status] + "\n\n" +
		"Layanan: " + jenisLayananToString[ekspedisi.jenisLayanan] + "\n\n" +
		"Alamat Asal  : " + ekspedisi.alamatAsal + "\n" +
		"Alamat Tujuan: " + ekspedisi.alamatTujuan + "\n\n" +
		"Deskripsi : \n" + ekspedisi.deskripsiBarang
	show_pager(content)
}

func (e *ModelEkspedisi) UpdateStatus() {
	var idx int
	if e.selectedResi != "-1" && e.searchBySelectedResi() != -1 {
		idx = e.searchBySelectedResi()
		if e.DaftarEkspedisi[idx].status < 4 {
			e.DaftarEkspedisi[idx].status += 1
			notify_form(e.selectedResi, e.DaftarEkspedisi[idx].idPelanggan, e.DaftarEkspedisi[idx].status)
		}
	} else {
		var content = " Data tidak ditemukan, harap keep sebuah ekspedisi dari daftar pelanggan"
		show_pager(content)
	}
}

func (e *ModelEkspedisi) Update() {
	/*
		Memperbarui data pelanggan, return false jika data tidak ditemukan
	*/

	var idx int
	if e.selectedResi != "-1" && e.searchBySelectedResi() != -1 {
		idx = e.searchBySelectedResi()
		edit_form(&e.DaftarEkspedisi[idx])
	} else {
		var content = " Data tidak ditemukan, harap keep sebuah ekspedisi dari daftar pelanggan"
		show_pager(content)
	}

}

func (e *ModelEkspedisi) Delete() {
	var idx int
	if e.selectedResi != "-1" && e.searchBySelectedResi() != -1 {
		idx = e.searchBySelectedResi()
		if confirm_form() {

			// Geser semua elemen setelah idx ke kiri
			for i := idx; i < e.nEkspedisi-1; i++ {
				e.DaftarEkspedisi[i] = e.DaftarEkspedisi[i+1]
			}
			// Kosongkan elemen terakhir
			e.DaftarEkspedisi[e.nEkspedisi-1] = Ekspedisi{}

			// Kurangi jumlah elemen
			e.nEkspedisi--
		}
	} else {
		var content = " Data tidak ditemukan, harap keep sebuah ekspedisi dari daftar pelanggan"
		show_pager(content)
	}
}

// search
func (e *ModelEkspedisi) searchBySelectedResi() int {
	/*
		Mengembalikan nilai index elemen dari e.daftarEkspedisi
		saat elemen.resi = selectedResi
		Pencarian menggunakan binary search
	*/

	e.SortByResiAscending() // mengurutkan data sebelum binary search
	low, high := 0, e.nEkspedisi-1
	for low <= high {
		mid := (low + high) / 2
		if e.DaftarEkspedisi[mid].resi == e.selectedResi {
			return mid
		} else if e.DaftarEkspedisi[mid].resi < e.selectedResi {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (e *ModelEkspedisi) searchByResi(resi string) int {
	/*
		Mengembalikan nilai index elemen dari e.daftarEkspedisi
		saat elemen.resi = resi
		Pencarian menggunakan sequential sort
	*/
	for i := 0; i < e.nEkspedisi; i++ {
		if e.DaftarEkspedisi[i].resi == resi {
			return i
		}
	}
	return -1

}

func (e *ModelEkspedisi) SearchResi() {

	var resi string
	resi = search_form()
	if resi != "cancelled" && e.searchByResi(resi) != -1 {
		e.SetSelected(resi)
		e.Read()
	}
}

// setter
func (e *ModelEkspedisi) SetSelected(resi string) {
	e.selectedResi = resi
}

// getter
func (e *ModelEkspedisi) GetSelectedPacket() string {
	if e.selectedResi != "-1" && e.searchBySelectedResi() != -1 {
		var idx = e.searchBySelectedResi()
		return "<< " + e.DaftarEkspedisi[idx].resi + " >>" + " ~ " + e.DaftarEkspedisi[idx].idPelanggan
	}
	return ""
}

// sorting
func (e *ModelEkspedisi) SortByResiAscending() {
	/*
		Mengurutkan e.daftarEkspedisi berdasarkan resi ekspedisi secara descending
		dengan menggunakan insertion sort
	*/
	for i := 0; i < e.nEkspedisi-1; i++ {
		minIdx := i
		for j := i + 1; j < e.nEkspedisi; j++ {
			if e.DaftarEkspedisi[j].resi < e.DaftarEkspedisi[minIdx].resi {
				minIdx = j
			}
		}

		// Tukar elemen yang ditemukan dengan elemen pertama
		e.DaftarEkspedisi[i], e.DaftarEkspedisi[minIdx] = e.DaftarEkspedisi[minIdx], e.DaftarEkspedisi[i]
	}
}

func (e *ModelEkspedisi) SortByResiDescending() {
	/*
		Mengurutkan e.daftarEkspedisi berdasarkan nama pelanggan secara descending
		dengan menggunakan insertion sort
	*/
	for i := 1; i < e.nEkspedisi; i++ {
		key := e.DaftarEkspedisi[i]
		j := i - 1

		// Pindahkan elemen dari e.daftarEkspedisi[0..i-1], yang lebih kecil dari key
		// ke satu posisi di depan posisi sekarang mereka
		for j >= 0 && e.DaftarEkspedisi[j].resi < key.resi {
			e.DaftarEkspedisi[j+1] = e.DaftarEkspedisi[j]
			j = j - 1
		}
		e.DaftarEkspedisi[j+1] = key
	}
}

func (e *ModelEkspedisi) SortByPriorityAscending() {
	/*
		Mengurutkan e.daftarEkspedisi berdasarkan prioritas pengiriman secara ascending
		dengan menggunakan selection sort
	*/
	for i := 0; i < e.nEkspedisi-1; i++ {
		minIdx := i
		for j := i + 1; j < e.nEkspedisi; j++ {
			if (e.DaftarEkspedisi[j].jenisLayanan < e.DaftarEkspedisi[minIdx].jenisLayanan) ||
				(e.DaftarEkspedisi[j].jenisLayanan == e.DaftarEkspedisi[minIdx].jenisLayanan && (e.DaftarEkspedisi[j].status < e.DaftarEkspedisi[minIdx].status)) ||
				(e.DaftarEkspedisi[j].jenisLayanan == e.DaftarEkspedisi[minIdx].jenisLayanan && (e.DaftarEkspedisi[j].status == e.DaftarEkspedisi[minIdx].status) && (e.DaftarEkspedisi[j].resi < e.DaftarEkspedisi[minIdx].resi)) {
				minIdx = j
			}
		}

		// Tukar elemen yang ditemukan dengan elemen pertama
		e.DaftarEkspedisi[i], e.DaftarEkspedisi[minIdx] = e.DaftarEkspedisi[minIdx], e.DaftarEkspedisi[i]
	}
}

func (e *ModelEkspedisi) SortByPriorityDescending() {
	/*
		Mengurutkan e.daftarEkspedisi berdasarkan prioritas pengiriman secara descending
		dengan menggunakan selection sort
	*/
	for i := 0; i < e.nEkspedisi-1; i++ {
		minIdx := i
		for j := i + 1; j < e.nEkspedisi; j++ {
			if (e.DaftarEkspedisi[j].jenisLayanan > e.DaftarEkspedisi[minIdx].jenisLayanan) ||
				(e.DaftarEkspedisi[j].jenisLayanan == e.DaftarEkspedisi[minIdx].jenisLayanan && (e.DaftarEkspedisi[j].status > e.DaftarEkspedisi[minIdx].status)) ||
				(e.DaftarEkspedisi[j].jenisLayanan == e.DaftarEkspedisi[minIdx].jenisLayanan && (e.DaftarEkspedisi[j].status == e.DaftarEkspedisi[minIdx].status) && (e.DaftarEkspedisi[j].resi > e.DaftarEkspedisi[minIdx].resi)) {
				minIdx = j
			}
		}

		// Tukar elemen yang ditemukan dengan elemen pertama
		e.DaftarEkspedisi[i], e.DaftarEkspedisi[minIdx] = e.DaftarEkspedisi[minIdx], e.DaftarEkspedisi[i]
	}
}

// filter
func (e *ModelEkspedisi) FilterByJenisLayanan(filter Layanan) ModelEkspedisi {
	/*
		Mengembalikan model ekspedisi dengan
		jenis layanan tertentu
	*/
	var filtered ModelEkspedisi

	var nActive int
	for i := 0; i < e.nEkspedisi; i++ {
		if e.DaftarEkspedisi[i].jenisLayanan == filter {
			filtered.DaftarEkspedisi[nActive] = e.DaftarEkspedisi[i]
			nActive++
		}
	}
	filtered.nEkspedisi = nActive
	filtered.selectedResi = "-1"
	filtered.nomorEkspedisi = e.nomorEkspedisi

	return filtered
}
