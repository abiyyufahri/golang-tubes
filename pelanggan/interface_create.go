package pelanggan

import (
	"github.com/charmbracelet/huh"
	"strings"
)

func create_form() (bool, Pelanggan) {

	var p Pelanggan
	var nama, alamat, nomorTelepon, alamatEmail string
	var confirm, status bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nama").
				Description("Masukkan nama pelanggan Anda").
				Placeholder("Cth: Yanto Sumargo").
				Value(&nama),

			huh.NewInput().
				Title("Alamat").
				Description("Tempat tinggal pelanggan").
				Placeholder("Cth: Jl. Radio Palasari No 34 Gg. ABC").
				Value(&alamat),

			huh.NewInput().
				Title("No. Hp").
				Description("Nomor telp/hp yang dapat dihubungi").
				Placeholder("Cth: 08123456789").
				Value(&nomorTelepon),

			huh.NewInput().
				Title("Email").
				Description("Email untuk informasikan ke pengguna").
				Placeholder("Cth: abcd@mantapjiwa.com").
				Value(&alamatEmail),
		),

		huh.NewGroup(

			huh.NewConfirm().
				Title("Aktifkan Pengguna?").
				Description("Bila tidak aktif, maka pengguna tidak akan dapat dibuatkan ekspedisi").
				Value(&status).
				Affirmative("Aktifkan!").
				Negative("Aktifkan nanti."),

			huh.NewConfirm().
				Title("Konfirmasi?").
				Description("Anda masih dapat mengubah nya dikemudian.").
				Value(&confirm).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return false, Pelanggan{}
	}
	if confirm {

		if strings.TrimSpace(nama) != "" {
			p.nama = nama
		}

		if strings.TrimSpace(alamat) != "" {
			p.alamat = alamat
		}

		if strings.TrimSpace(alamatEmail) != "" {
			p.alamatEmail = alamatEmail
		}

		if strings.TrimSpace(nomorTelepon) != "" {
			p.nomorTelepon = nomorTelepon
		}

		p.status = status

		return confirm, p
	}

	return false, Pelanggan{}
}

func update_form(p *Pelanggan) {

	var nama, alamat, nomorTelepon, alamatEmail string
	var confirm, status bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nama").
				Description("Masukkan nama pelanggan Anda").
				Placeholder(p.nama).
				Value(&nama),

			huh.NewInput().
				Title("Alamat").
				Description("Tempat tinggal pelanggan").
				Placeholder(p.alamat).
				Value(&alamat),

			huh.NewInput().
				Title("No. Hp").
				Description("Nomor telp/hp yang dapat dihubungi").
				Placeholder(p.nomorTelepon).
				Value(&nomorTelepon),

			huh.NewInput().
				Title("Email").
				Description("Email untuk informasikan ke pengguna").
				Placeholder(p.alamatEmail).
				Value(&alamatEmail),
		).
			Title("Ubah data pelanggan").
			Description("Kosongkan bila tidak ingin mengganti"),

		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("Status").
				Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
				Options(
					huh.NewOption[bool]("Aktif", true),
					huh.NewOption[bool]("Nonaktif", false),
				).
				Value(&status),

			huh.NewConfirm().
				Title("Konfirmasi?").
				Description("Perubahan yang terjadi tidak dapat diulang").
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

		if strings.TrimSpace(nama) != "" {
			p.nama = nama
		}

		if strings.TrimSpace(alamat) != "" {
			p.alamat = alamat
		}

		if strings.TrimSpace(alamatEmail) != "" {
			p.alamatEmail = alamatEmail
		}

		if strings.TrimSpace(nomorTelepon) != "" {
			p.nomorTelepon = nomorTelepon
		}

		p.status = status
	}
}

func confirm_form() bool {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("\n\nKonfirmasi Hapus Pelanggan?").
				Description("Data ekspedisi akan tetap tersedia untuk pelanggan ini. ").
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

func search_form() (bool, string) {
	var nama, id string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Cari data pasien").Description("Silahkan isi salah satu, nama lebih diutamakan"),
			huh.NewInput().Title("Cari denang nama").Value(&nama),
			huh.NewInput().Title("Cari denang id").Value(&id),
		),
	)

	err := form.Run()
	if err != nil {
		//fmt.Println("error in pelanggan's create form:", err)
		return true, "cancelled"
	}

	if strings.TrimSpace(nama) != "" {
		return true, strings.TrimSpace(nama)
	} else if strings.TrimSpace(id) != "" {
		return false, strings.TrimSpace(id)
	}

	return true, "cancelled"
}
