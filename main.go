package main

import (
	"tubes/interface"
	p "tubes/pelanggan"
)

func main() {

	var dataPelanggan p.ModelPelanggan

	var pilihan int
	pilihan = _interface.GetModuleChoice()
	for pilihan != -1 {
		switch pilihan {
		case 1:
			subMenuEkspedisi()
		case 2:
			subMenuPelanggan(&dataPelanggan)
		}

		pilihan = _interface.GetModuleChoice()
	}

	quit()
}

func subMenuPelanggan(dp *p.ModelPelanggan) {
	var subPilihan int
	subPilihan = _interface.GetModulSubMenuPelanggan()
	for subPilihan != 0 {

		subPilihan = _interface.GetModulSubMenuPelanggan()
	}
}

func subMenuEkspedisi() {
	var subPilihan int
	subPilihan = _interface.GetModulSubMenuEkspedisi()
	for subPilihan != 0 {

		subPilihan = _interface.GetModulSubMenuEkspedisi()
	}
}
