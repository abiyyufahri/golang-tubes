package main

import (
	e "TUBES_GO/ekspedisi"
	"TUBES_GO/interface"
	p "TUBES_GO/pelanggan"
)

func main() {

	var dataPelanggan p.ModelPelanggan
	var dataEkspedisi e.ModelEkspedisi

	dataPelanggan.Init()
	dataEkspedisi.Init()

	dataPelanggan.GetData()
	dataEkspedisi.GetData()

	var pilihan int
	pilihan = _interface.GetModuleChoice()
	for pilihan != -1 {
		switch pilihan {
		case 1:
			subMenuEkspedisi(&dataEkspedisi, dataPelanggan)
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
			dp.Search()
		case 4:
			dp.Read()
		case 5:
			dp.Update()
		case 6:
			dp.Delete()
		}

		subPilihan = _interface.GetModulSubMenuPelanggan(dp.GetSelectedName())
	}
}

func subMenuEkspedisi(de *e.ModelEkspedisi, dp p.ModelPelanggan) {
	var subPilihan int
	subPilihan = _interface.GetModulSubMenuEkspedisi(de.GetSelectedPacket())
	for subPilihan != 0 {

		switch subPilihan {
		case 1:
			de.Create()
		case 2:
			de.ReadAll(dp)
		case 3:
			de.SearchResi()
		case 4:
			de.UpdateStatus()
		case 5:
			de.Read()
		case 6:
			de.Update()
		case 7:
			de.Delete()
		}

		subPilihan = _interface.GetModulSubMenuEkspedisi(de.GetSelectedPacket())
	}
}
