package main

import "fmt"

const NMAX_EKSPEDISI int = 20

type Ekspedisi struct {
	id             int
	idPelanggan    int
	jenisLayanan   string
	biayaEkspedisi string
	status         int
}

type tabEkspedisi [NMAX_EKSPEDISI]Ekspedisi

var daftarEkspedisi tabEkspedisi
var nE int

func createEkspedisi(idPelanggan int, jenisLayanan, biayaEkspedisi string, status int) {
	var val bool
	val = false

	for i := 0; i < nP; i++ {
		if daftarPelanggan[i].id == idPelanggan {
			val = true
			break
		}
	}

	if nE < NMAX_EKSPEDISI && val {
		daftarEkspedisi[nE] = Ekspedisi{nE + 1, idPelanggan, jenisLayanan, biayaEkspedisi, status}
		nE++
		fmt.Println("Ekspedisi berhasil ditambahkan.")
	} else {
		fmt.Println("Daftar ekspedisi penuh atau id pelanggan tidak ditemukan!")
	}
}

func ReadEkspedisi() tabEkspedisi {
	return daftarEkspedisi
}

func UpdateEkspedisi(id int, jenisLayanan, biayaEkspedisi string, status int) {
	var left, right, mid int
	left = 0
	right = nE - 1

	for left <= right {
		mid = (left + right) / 2
		if daftarEkspedisi[mid].id == id {
			if jenisLayanan != "0" {
				daftarEkspedisi[mid].jenisLayanan = jenisLayanan
			}
			if biayaEkspedisi != "0" {
				daftarEkspedisi[mid].biayaEkspedisi = biayaEkspedisi
			}
			if status != 0 {
				daftarEkspedisi[mid].status = status
			}
			fmt.Println("Data ekspedisi berhasil diperbarui.")
			return
		} else if daftarEkspedisi[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Ekspedisi tidak ditemukan!")
}

func DeleteEkspedisi(id int) {
	for i := 0; i < nE; i++ {
		if daftarEkspedisi[i].id == id {
			daftarEkspedisi[i] = daftarEkspedisi[nE-1]
			nE--
			return
		}
	}
	fmt.Println("Ekspedisi tidak ditemukan!")
}
