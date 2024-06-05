package main

import (
	"TUBES_GO/interface"
	p "TUBES_GO/pelanggan"
)

func main() {

	var dataPelanggan p.ModelPelanggan

	dataPelanggan.Init()
	dataPelanggan.GetData()

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
	subPilihan = _interface.GetModulSubMenuPelanggan(dp.GetSelectedName())
	for subPilihan != 0 {

		switch subPilihan {
		case 1:
			dp.Create()
		case 2:
			dp.ReadAll()
		case 3:
			dp.Read()
		case 4:
			dp.Update()
		}

		subPilihan = _interface.GetModulSubMenuPelanggan(dp.GetSelectedName())
	}
}

func subMenuEkspedisi() {
	var subPilihan int
	subPilihan = _interface.GetModulSubMenuEkspedisi()
	for subPilihan != 0 {

		subPilihan = _interface.GetModulSubMenuEkspedisi()
	}
}
