package pelanggan

import "strconv"

const NMAXPELANGGAN int = 20

type Pelanggan struct {
	id           int
	nama         string
	alamat       string
	nomorTelepon string
	alamatEmail  string
	status       bool
}

type tabPelanggan [NMAXPELANGGAN]Pelanggan

type ModelPelanggan struct {
	daftarPelanggan tabPelanggan
	nPelanggan      int
}

func (p *ModelPelanggan) Create() bool {
	/*
		Membuat data pelanggan baru, return false bila data telah penuh
	*/

	if p.nPelanggan < NMAXPELANGGAN {

		var i = p.nPelanggan

		p.daftarPelanggan[i].id = p.nPelanggan + 1
		create_form(&p.daftarPelanggan[i])
		p.nPelanggan++

		return true
	}

	return false // data pelanggan telah penuh
}

func (p *ModelPelanggan) ReadAll() {
	/*
		menampilkan seluruh data pelanggan menggunakan table
	*/
	viewAllTable(*p)
}

func (p *ModelPelanggan) Read(id int) {

	var booleanToString = map[bool]string{
		true:  "Aktif",
		false: "Nonaktif",
	}

	var pelanggan Pelanggan
	pelanggan = p.daftarPelanggan[id]

	var content string
	content = "Nama :" + "\n" +
		"id : " + strconv.Itoa(pelanggan.id) + "\n" +
		"status : " + booleanToString[pelanggan.status] + "\n" +
		"no telp: " + pelanggan.nomorTelepon + "\n" +
		"email  : " + pelanggan.alamatEmail + "\n" +
		"alamat : " + pelanggan.alamat + "\n"

	show_pager(content)
}

func (p *ModelPelanggan) Update(id int, nama, alamat, nomorTelepon, alamatEmail string) bool {
	/*
		Memperbarui data pelanggan, return false jika data tidak ditemukan
	*/

	var idx int
	idx = p.SearchById(id)

	if idx != -1 {
		if nama != "0" {
			p.daftarPelanggan[idx].nama = nama
		}
		if alamat != "0" {
			p.daftarPelanggan[idx].alamat = alamat
		}
		if nomorTelepon != "0" {
			p.daftarPelanggan[idx].nomorTelepon = nomorTelepon
		}
		if alamatEmail != "0" {
			p.daftarPelanggan[idx].alamatEmail = alamatEmail
		}
		return true // sukses
	}

	return false // gagal
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

//func SortPelanggan(daftarPelanggan *tabPelanggan, prioritizedStatus int) {
//	for i := 1; i < nP; i++ {
//		key := daftarPelanggan[i]
//		j := i - 1
//
//		for j >= 0 && ((daftarPelanggan[j].status != prioritizedStatus && key.status == prioritizedStatus) ||
//			(daftarPelanggan[j].status != prioritizedStatus && key.status < daftarPelanggan[j].status)) {
//			daftarPelanggan[j+1] = daftarPelanggan[j]
//			j = j - 1
//		}
//		daftarPelanggan[j+1] = key
//	}
//}
