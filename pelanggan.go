package main

import "fmt"

const NMAX_PELANGGAN int = 20

type Pelanggan struct {
	id           int
	nama         string
	alamat       string
	nomorTelepon string
	alamatEmail  string
}

type tabPelanggan [NMAX_PELANGGAN]Pelanggan

var daftarPelanggan tabPelanggan
var nP int = 0

func CreatePelanggan(nama, alamat, nomorTelepon, alamatEmail string) {
	if nP < NMAX_PELANGGAN {
		daftarPelanggan[nP] = Pelanggan{nP + 1, nama, alamat, nomorTelepon, alamatEmail}
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
			daftarPelanggan[mid].nama = nama
			daftarPelanggan[mid].alamat = alamat
			daftarPelanggan[mid].nomorTelepon = nomorTelepon
			daftarPelanggan[mid].alamatEmail = alamatEmail
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
