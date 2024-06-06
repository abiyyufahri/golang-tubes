package ekspedisi

import (
	"github.com/charmbracelet/huh"
	"strconv"
	"strings"
)

func create_form() (bool, Ekspedisi) {

	var e Ekspedisi

	var jenisLayanan Layanan
	var harga, alamatAsal, idPelanggan, alamatTujuan, deskripsi string

	var confirm bool
	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[Layanan]().
				Title("Jenis Layanan").
				Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
				Options(
					huh.NewOption[Layanan]("Reguler", Reguler),
					huh.NewOption[Layanan]("SameDay", SameDay),
					huh.NewOption[Layanan]("Instant", Instant),
					huh.NewOption[Layanan]("Active", Cargo),
				).
				Value(&jenisLayanan),

			huh.NewInput().
				Title("Pelanggan").
				Description("Masukkan id pelanggan.").
				Placeholder("cth. 1").
				Value(&idPelanggan),

			huh.NewText().
				Title("Deskripsi Barang").
				Description("Masukkan penjelasan barang.").
				Placeholder("cth. bedak bayi dan gelas cantik. diantar blabla").
				Value(&deskripsi),

			huh.NewInput().
				Title("Biaya").
				Description("Masukkan bilangan bulat.").
				Placeholder("cth. 100000 atau 100 000").
				Value(&harga),

			huh.NewInput().
				Title("Alamat Asal").
				Description("Tempat paket dijemput.").
				Placeholder("cth. Cth: Jl. Radio Palasari No 34 Gg. ABC ").
				Value(&alamatAsal),

			huh.NewInput().
				Title("Alamat Tujuan").
				Description("Tempat paket diantar.").
				Placeholder("cth. Cth: Jl. Radio Palasari No 34 Gg. BCD ").
				Value(&alamatTujuan),

			huh.NewConfirm().
				Title("Konfirmasi?").
				Description("Anda masih dapat mengubah nya dikemudian. ").
				Value(&confirm).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return false, Ekspedisi{}
	}
	if confirm {

		if jenisLayanan > 0 {
			e.jenisLayanan = jenisLayanan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(harga) != "" {
			hrg := strings.ReplaceAll(harga, " ", "")
			hargaInt, _ := strconv.Atoi(hrg)
			e.biayaEkspedisi = hargaInt
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(idPelanggan) != "" {
			e.idPelanggan = idPelanggan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(alamatAsal) != "" {
			e.alamatAsal = alamatAsal
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(alamatTujuan) != "" {
			e.alamatTujuan = alamatTujuan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(deskripsi) != "" {
			e.deskripsiBarang = deskripsi
		} else {
			e.butuhDilengkapi = true
		}

		return confirm, e
	}

	return false, Ekspedisi{}
}

func search_form() string {

	var resi string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Cari dengan Resi").Value(&resi),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return "cancelled"
	}

	return resi
}

func edit_form(e *Ekspedisi) {

	var jenisLayanan Layanan
	var harga, alamatAsal, idPelanggan, alamatTujuan, deskripsi string

	var confirm bool
	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[Layanan]().
				Title("Jenis Layanan").
				Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
				Options(
					huh.NewOption[Layanan]("Reguler", Reguler).
						Selected(e.jenisLayanan == Reguler),
					huh.NewOption[Layanan]("SameDay", SameDay).
						Selected(e.jenisLayanan == SameDay),
					huh.NewOption[Layanan]("Instant", Instant).
						Selected(e.jenisLayanan == Instant),
					huh.NewOption[Layanan]("Active", Cargo).
						Selected(e.jenisLayanan == Cargo),
				).
				Value(&jenisLayanan),

			huh.NewInput().
				Title("Pelanggan").
				Description("Masukkan id pelanggan.").
				Placeholder(e.idPelanggan).
				Value(&idPelanggan),

			huh.NewText().
				Title("Deskripsi Barang").
				Description("Masukkan penjelasan barang.").
				Placeholder(e.deskripsiBarang).
				Value(&deskripsi),

			huh.NewInput().
				Title("Biaya").
				Description("Masukkan bilangan bulat.").
				Placeholder(strconv.Itoa(e.biayaEkspedisi)).
				Value(&harga),

			huh.NewInput().
				Title("Alamat Asal").
				Description("Tempat paket dijemput.").
				Placeholder(e.alamatAsal).
				Value(&alamatAsal),

			huh.NewInput().
				Title("Alamat Tujuan").
				Description("Tempat paket diantar.").
				Placeholder(e.alamatTujuan).
				Value(&alamatTujuan),

			huh.NewConfirm().
				Title("Konfirmasi?").
				Description("Perubahan tidak dapat dipulihkan. ").
				Value(&confirm).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
	}
	if confirm {

		if jenisLayanan > 0 {
			e.jenisLayanan = jenisLayanan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(harga) != "" {
			hrg := strings.ReplaceAll(harga, " ", "")
			hargaInt, _ := strconv.Atoi(hrg)
			e.biayaEkspedisi = hargaInt
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(idPelanggan) != "" {
			e.idPelanggan = idPelanggan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(alamatAsal) != "" {
			e.alamatAsal = alamatAsal
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(alamatTujuan) != "" {
			e.alamatTujuan = alamatTujuan
		} else {
			e.butuhDilengkapi = true
		}

		if strings.TrimSpace(deskripsi) != "" {
			e.deskripsiBarang = deskripsi
		} else {
			e.butuhDilengkapi = true
		}

	}

}

func confirm_form() bool {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("\n\nKonfirmasi Hapus Ekspedisi?").
				Description("Operasi tidak dapat dipulihkan. ").
				Value(&confirm).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return false
	}

	return confirm
}

func notify_form(resi, milik string, status int) bool {
	var confirm bool

	var statusToString = map[int]string{
		1: "Dikemas",
		2: "Dijemput",
		3: "Diantar",
		4: "Selesai",
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("\n\nBerhasil update Status : " + statusToString[status]).
				Description("\nStatus " + resi + " milih pelanggan " + milik + " Di Update").
				Next(true).NextLabel("Oke!"),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return false
	}

	return confirm
}
