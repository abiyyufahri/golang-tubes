package main

import "fmt"

const NMAX_EKSPEDISI int = 20

type Ekspedisi struct {
	idEkspedisi   int
	idPelanggan   int
	namaEkspedisi string
	jenisLayanan  string
	status        int
}

type tabEkspedisi [NMAX_EKSPEDISI]Ekspedisi

var daftarEkspedisi tabEkspedisi
var nE int

func createEkspedisi(idPelanggan int, namaEkspedisi, jenisLayanan string, status int) {
	if nE < NMAX_EKSPEDISI {
		daftarEkspedisi[nE] = Ekspedisi{nE + 1, idPelanggan, namaEkspedisi, jenisLayanan, status}
		nE++
	} else {
		fmt.Println("Daftar ekspedisi penuh!")
	}
}

func readEkspedisi() tabEkspedisi {
	return daftarEkspedisi
}

func updateEkspedisi(id int, namaEkspedisi, jenisLayanan string, status int) {
	for i := 0; i < nE; i++ {
		if daftarEkspedisi[i].idEkspedisi == id {
			daftarEkspedisi[i].namaEkspedisi = namaEkspedisi
			daftarEkspedisi[i].jenisLayanan = jenisLayanan
			daftarEkspedisi[i].status = status
			return
		}
	}
	fmt.Println("Ekspedisi tidak ditemukan!")
}

func deleteEkspedisi(id int) {
	for i := 0; i < nE; i++ {
		if daftarEkspedisi[i].idEkspedisi == id {
			daftarEkspedisi[i] = daftarEkspedisi[nE-1]
			nE--
			return
		}
	}
	fmt.Println("Ekspedisi tidak ditemukan!")
}
