package main

import "fmt"

const NMAX_PELANGGAN int = 20

type Pelanggan struct {
	id           int
	nama         string
	alamat       string
	nomorTelepon string
	alamatEmail  string
	status       int
}

type tabPelanggan [NMAX_PELANGGAN]Pelanggan

var daftarPelanggan tabPelanggan
var nP int = 0

func CreatePelanggan(nama, alamat, nomorTelepon, alamatEmail string, status int) {
	if nP < NMAX_PELANGGAN {
		daftarPelanggan[nP] = Pelanggan{nP + 1, nama, alamat, nomorTelepon, alamatEmail, status}
		nP++
	} else {
		fmt.Println("Daftar pelanggan penuh!")
	}
}

func ReadAllPelanggan() tabPelanggan {
	return daftarPelanggan
}

func UpdatePelanggan(id int, nama, alamat, nomorTelepon, alamatEmail string) {
	var left, right, mid int
	left = 0
	right = nP - 1

	for left <= right {
		mid = (left + right) / 2
		if daftarPelanggan[mid].id == id {
			if nama != "0" {
				daftarPelanggan[mid].nama = nama
			}
			if alamat != "0" {
				daftarPelanggan[mid].alamat = alamat
			}
			if nomorTelepon != "0" {
				daftarPelanggan[mid].nomorTelepon = nomorTelepon
			}
			if alamatEmail != "0" {
				daftarPelanggan[mid].alamatEmail = alamatEmail
			}
			fmt.Println("Data pelanggan berhasil diperbarui.")
			return
		} else if daftarPelanggan[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Pelanggan tidak ditemukan.")
}

func UpdateStatusPelanggan(id, status int) {
	var left, right, mid int
	left = 0
	right = nP - 1

	for left <= right {
		mid = (left + right) / 2
		if daftarPelanggan[mid].id == id {
			daftarPelanggan[mid].status = status
			fmt.Println("Status pelanggan berhasil diperbarui.")
			return
		} else if daftarPelanggan[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Pelanggan tidak ditemukan.")
}

func DeletePelanggan(id int) {
	for i := 0; i < nP; i++ {
		if daftarPelanggan[i].id == id {
			for j := i; j < nP; j++ {
				daftarPelanggan[j] = daftarPelanggan[j+1]
			}
			nP--
		}
	}
}

func SortingPelanggan(daftarPelanggan *tabPelanggan, prioritizedStatus int) {
	for i := 1; i < nP; i++ {
		key := daftarPelanggan[i]
		j := i - 1

		for j >= 0 && ((daftarPelanggan[j].status != prioritizedStatus && key.status == prioritizedStatus) ||
			(daftarPelanggan[j].status != prioritizedStatus && key.status < daftarPelanggan[j].status)) {
			daftarPelanggan[j+1] = daftarPelanggan[j]
			j = j - 1
		}
		daftarPelanggan[j+1] = key
	}
}
