package main

const NMAX_PELANGGAN int = 20

type Pelanggan struct {
	id           int
	nama         string
	alamat       string
	nomorTelepon string
	alamatEmail  string
	status       bool
}

type tabPelanggan [NMAX_PELANGGAN]Pelanggan

type modelPelanggan struct {
	daftarPelanggan tabPelanggan
	nPelanggan      int
}

func (p modelPelanggan) Create(nama, alamat, nomorTelepon, alamatEmail string, status bool) bool {
	/*
		Membuat data pelanggan baru, return false bila data telah penuh
	*/

	if p.nPelanggan < NMAX_PELANGGAN {
		p.daftarPelanggan[p.nPelanggan] = Pelanggan{p.nPelanggan + 1, nama, alamat, nomorTelepon, alamatEmail, status}
		p.nPelanggan++

		return true
	}

	return false // data pelanggan telah penuh
}

func (p modelPelanggan) ReadAll() (tabPelanggan, int) {
	/*
		Mengembalikan daftar pelanggan dan jumlah pelanggan nya
	*/
	return p.daftarPelanggan, p.nPelanggan
}

func (p modelPelanggan) Update(id int, nama, alamat, nomorTelepon, alamatEmail string) bool {
	/*
		Memperbarui data pelanggan, return false jika data tidak ditemukan
	*/

	var idx int
	idx = p.Search(id)

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

func (p modelPelanggan) UpdateStatus(id int, status bool) bool {
	/*
		Memperbarui status, return false jika data tidak ditemukan
	*/

	var idx int
	idx = p.Search(id)

	if idx != -1 {
		p.daftarPelanggan[idx].status = status
		return true // berhasil
	}

	return false // gagal

}

func (p modelPelanggan) Search(id int) int {
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

func (p modelPelanggan) Delete(id int) bool {
	/*
		Menghapus data berdasarkan id, return false bila id tidak ditemukan
	*/
	var idx int
	idx = p.Search(id)

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
