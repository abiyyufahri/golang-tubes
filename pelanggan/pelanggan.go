package pelanggan

import (
	"strconv"
)

const NMAXPELANGGAN int = 20

type Pelanggan struct {
	id           int
	nama         string
	alamat       string
	nomorTelepon string
	alamatEmail  string
	status       bool
}

type TabPelanggan [NMAXPELANGGAN]Pelanggan

type ModelPelanggan struct {
	selectedId      int // nomor pelanggan yang dipilih
	daftarPelanggan TabPelanggan
	nomorPelanggan  int
	nPelanggan      int
}

func (p *ModelPelanggan) Init() {
	/*
		I.S. model pelanggan, masih kosong
		F.S. membuat inisialisasi state dari model
	*/
	p.selectedId = -1    // tidak ada id yang dipilih
	p.nomorPelanggan = 1 // nomor dimulai dari 1
}

func (p *ModelPelanggan) Create() {
	/*
		Membuat data pelanggan baru, return false bila data telah penuh
	*/

	if p.nPelanggan < NMAXPELANGGAN {

		var i = p.nPelanggan

		var pelanggan Pelanggan
		var isSuccess bool
		isSuccess, pelanggan = create_form()

		if isSuccess {
			p.daftarPelanggan[i] = pelanggan
			p.daftarPelanggan[i].id = p.nomorPelanggan
			p.nPelanggan++
			p.nomorPelanggan++
		}
		return
	}

	content := "Data telah penuh "
	show_pager(content)
}

func (p *ModelPelanggan) ReadAll() {
	/*
		menampilkan seluruh data pelanggan menggunakan table
	*/
	viewAllTable(p)
}

func (p *ModelPelanggan) setSelectedId(selectedId int) {
	/*
		mennganti id yang terpilih
	*/
	p.selectedId = selectedId
}
func (p *ModelPelanggan) GetSelectedName() string {
	/*
		Mengembalikan nama yang terpilih menggunakan
	*/

	if p.selectedId != -1 && p.searchById() != -1 {
		var idx = p.searchById()
		return p.daftarPelanggan[idx].nama
	}
	return ""
}

func (p *ModelPelanggan) Read() {

	var booleanToString = map[bool]string{
		true:  "Aktif",
		false: "Nonaktif",
	}

	var idx = p.searchById()

	var content string
	if idx != -1 {
		var pelanggan = p.daftarPelanggan[idx]
		content = "Nama :" + pelanggan.nama + "\n" +
			"id : " + strconv.Itoa(pelanggan.id) + "\n" +
			"status : " + booleanToString[pelanggan.status] + "\n" +
			"no telp: " + pelanggan.nomorTelepon + "\n" +
			"email  : " + pelanggan.alamatEmail + "\n" +
			"alamat : " + pelanggan.alamat + "\n"
	} else {
		content = " Data tidak ditemukan, harap keep seorang pelanggan dari daftar pelanggan"
	}

	show_pager(content)
}

func (p *ModelPelanggan) Update() {
	/*
		Memperbarui data pelanggan, return false jika data tidak ditemukan
	*/

	var idx int
	if p.selectedId != -1 && p.searchById() != -1 {
		idx = p.searchById()
		update_form(&p.daftarPelanggan[idx])
	} else {
		var content = " Data tidak ditemukan, harap keep seorang pelanggan dari daftar pelanggan"
		show_pager(content)
	}

}

func (p *ModelPelanggan) UpdateStatus(id int, status bool) bool {
	/*
		Memperbarui status, return false jika data tidak ditemukan
	*/

	var idx int
	idx = p.SearchById(id)

	if idx != -1 {
		p.daftarPelanggan[idx].status = status
		return true // berhasil
	}

	return false // gagal

}

func (p *ModelPelanggan) SearchById(id int) int {
	/*
		Mengembalikan index dari id pelanggan, atau -1 bila tidak ditemukan
		Note: Pencarian menggunakan binary search
	*/

	var left, right, mid int
	left = 0
	right = p.nPelanggan - 1

	for left <= right {
		mid = (left + right) / 2
		if p.daftarPelanggan[mid].id == id {
			return mid
		} else if p.daftarPelanggan[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // jika tidak ditemukan
}

func (p *ModelPelanggan) Delete(id int) bool {
	/*
		Menghapus data berdasarkan id, return false bila id tidak ditemukan
	*/
	var idx int
	idx = p.SearchById(id)

	if idx == -1 {
		return false // jika id tidak ditemukan
	}

	for i := id; i < p.nPelanggan-1; i++ {
		p.daftarPelanggan[i] = p.daftarPelanggan[i+1]
	}
	return true
}

// sorting pelanggan
func (p *ModelPelanggan) SortByIdDescending() {
	/*
		Mengurutkan p.daftarPelanggan berdasarkan id pelanggan secara descending
		dengan menggunakan insertion sort
	*/
	for i := 1; i < p.nPelanggan; i++ {
		key := p.daftarPelanggan[i]
		j := i - 1

		// Pindahkan elemen dari p.daftarPelanggan[0..i-1], yang lebih kecil dari key
		// ke satu posisi di depan posisi sekarang mereka
		for j >= 0 && p.daftarPelanggan[j].id < key.id {
			p.daftarPelanggan[j+1] = p.daftarPelanggan[j]
			j = j - 1
		}
		p.daftarPelanggan[j+1] = key
	}
}

func (p *ModelPelanggan) SortByIdAscending() {
	/*
		Mengurutkan p.daftarPelanggan berdasarkan id pelanggan secara descending
		dengan menggunakan insertion sort
	*/
	for i := 0; i < p.nPelanggan-1; i++ {
		minIdx := i
		for j := i + 1; j < p.nPelanggan; j++ {
			if p.daftarPelanggan[j].id < p.daftarPelanggan[minIdx].id {
				minIdx = j
			}
		}

		// Tukar elemen yang ditemukan dengan elemen pertama
		p.daftarPelanggan[i], p.daftarPelanggan[minIdx] = p.daftarPelanggan[minIdx], p.daftarPelanggan[i]
	}
}

func (p *ModelPelanggan) sortByNameDescending() {
	/*
		Mengurutkan p.daftarPelanggan berdasarkan nama pelanggan secara descending
		dengan menggunakan insertion sort
	*/
	for i := 1; i < p.nPelanggan; i++ {
		key := p.daftarPelanggan[i]
		j := i - 1

		// Pindahkan elemen dari p.daftarPelanggan[0..i-1], yang lebih kecil dari key
		// ke satu posisi di depan posisi sekarang mereka
		for j >= 0 && p.daftarPelanggan[j].nama < key.nama {
			p.daftarPelanggan[j+1] = p.daftarPelanggan[j]
			j = j - 1
		}
		p.daftarPelanggan[j+1] = key
	}
}

func (p *ModelPelanggan) SortByNameAscending() {
	/*
		Mengurutkan p.daftarPelanggan berdasarkan nama pelanggan secara descending
		dengan menggunakan insertion sort
	*/
	for i := 0; i < p.nPelanggan-1; i++ {
		minIdx := i
		for j := i + 1; j < p.nPelanggan; j++ {
			if p.daftarPelanggan[j].nama < p.daftarPelanggan[minIdx].nama {
				minIdx = j
			}
		}

		// Tukar elemen yang ditemukan dengan elemen pertama
		p.daftarPelanggan[i], p.daftarPelanggan[minIdx] = p.daftarPelanggan[minIdx], p.daftarPelanggan[i]
	}
}

// filter
func (p *ModelPelanggan) filterByActive() ModelPelanggan {
	/*
		Mengembalikan model pelanggan dengan daftarpelanggan
		yang berstatus aktif saja
	*/
	var pelangganActive ModelPelanggan

	var nActive int
	for i := 0; i < p.nPelanggan; i++ {
		if p.daftarPelanggan[i].status == true {
			pelangganActive.daftarPelanggan[nActive] = p.daftarPelanggan[i]
			nActive++
		}
	}
	pelangganActive.nPelanggan = nActive
	pelangganActive.selectedId = -1
	pelangganActive.nomorPelanggan = p.nomorPelanggan

	return pelangganActive
}

func (p *ModelPelanggan) filterByNonActive() ModelPelanggan {
	/*
		Mengembalikan model pelanggan dengan daftarpelanggan
		yang berstatus non-active saja
	*/
	var pelangganNonActive ModelPelanggan

	var nActive int
	for i := 0; i < p.nPelanggan; i++ {
		if p.daftarPelanggan[i].status == false {
			pelangganNonActive.daftarPelanggan[nActive] = p.daftarPelanggan[i]
			nActive++
		}
	}
	pelangganNonActive.nPelanggan = nActive
	pelangganNonActive.selectedId = -1
	pelangganNonActive.nomorPelanggan = p.nomorPelanggan

	return pelangganNonActive
}

// search
func (p *ModelPelanggan) searchById() int {
	/*
		Mengembalikan nilai index elemen dari p.daftarPelanggan
		saat elemen.id = selectedId
		Pencarian menggunakan binary search
	*/
	low, high := 0, p.nPelanggan-1
	for low <= high {
		mid := (low + high) / 2
		if p.daftarPelanggan[mid].id == p.selectedId {
			return mid
		} else if p.daftarPelanggan[mid].id < p.selectedId {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (p *ModelPelanggan) GetAll() ModelPelanggan {
	return *p
}
