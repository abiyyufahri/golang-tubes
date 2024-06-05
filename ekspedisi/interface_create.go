package ekspedisi

//
//import (
//	"fmt"
//	"github.com/charmbracelet/huh"
//	p "tubes/pelanggan"
//)
//
//func create_form(e *Ekspedisi, mp p.ModelPelanggan) {
//
//	var daftarPelanggan
//
//	form := huh.NewForm(
//		huh.NewGroup(
//			huh.NewSelect[p.TabPelanggan]().
//				Title("Pilih Pelanggan").
//				Description("Silahkan pilih pelanggan yang akan ada tampilkan").
//				Options(
//					huh.NewOption()
//			)
//
//			huh.NewInput().
//				Title("Alamat").
//				Description("Tempat pengiriman ke pelanggan").
//				Placeholder("Cth: Jl. Radio Palasari No 34 Gg. ABC").
//				Value(&p.alamat),
//
//			huh.NewInput().
//				Title("No. Hp").
//				Description("Nomor telp/hp yang dapat dihubungi").
//				Placeholder("Cth: 08123456789").
//				Value(&p.nomorTelepon),
//
//			huh.NewInput().
//				Title("Email").
//				Description("Email untuk informasikan ke pengguna").
//				Placeholder("Cth: abcd@mantapjiwa.com").
//				Value(&p.alamatEmail),
//		),
//
//		huh.NewGroup(
//			//huh.NewSelect[bool]().
//			//	Title("Status").
//			//	Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
//			//	Options(
//			//		huh.NewOption[bool]("Aktif", true),
//			//		huh.NewOption[bool]("Nonaktif", false),
//			//	).
//			//	Value(&p.status),
//
//			huh.NewConfirm().
//				Title("Aktifkan Pengguna?").
//				Description("Bila tidak aktif, maka pengguna tidak akan dapat dibuatkan ekspedisi").
//				Value(&p.status).
//				Affirmative("Yes!").
//				Negative("No."),
//		),
//	)
//
//	err := form.Run()
//	if err != nil {
//		fmt.Println("error in pelanggan's create form:", err)
//	}
//}
//
//func update_form(p *Pelanggan) {
//
//	var nama, alamat, nomorTelepon, alamatEmail string
//	var confirm bool
//	form := huh.NewForm(
//		huh.NewGroup(
//			huh.NewInput().
//				Title("Nama").
//				Description("Masukkan nama pelanggan Anda").
//				Placeholder(p.nama).
//				Value(&nama),
//
//			huh.NewInput().
//				Title("Alamat").
//				Description("Tempat pengiriman ke pelanggan").
//				Placeholder(p.alamat).
//				Value(&alamat),
//
//			huh.NewInput().
//				Title("No. Hp").
//				Description("Nomor telp/hp yang dapat dihubungi").
//				Placeholder(p.nomorTelepon).
//				Value(&nomorTelepon),
//
//			huh.NewInput().
//				Title("Email").
//				Description("Email untuk informasikan ke pengguna").
//				Placeholder(p.alamatEmail).
//				Value(&alamatEmail),
//		),
//
//		huh.NewGroup(
//			//huh.NewSelect[bool]().
//			//	Title("Status").
//			//	Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
//			//	Options(
//			//		huh.NewOption[bool]("Aktif", true),
//			//		huh.NewOption[bool]("Nonaktif", false),
//			//	).
//			//	Value(&p.status),
//
//			huh.NewConfirm().
//				Title("Konfirmasi?").
//				Description("Perubahan yang terjadi tidak dapat diulang").
//				Value(&confirm).
//				Affirmative("Yes!").
//				Negative("No."),
//		),
//	)
//
//	if confirm {
//		if nama != "" {
//			p.nama = nama
//		}
//
//		if alamat != "" {
//			p.alamat = alamat
//		}
//
//		if alamatEmail != "" {
//			p.alamatEmail = alamatEmail
//		}
//
//		if nomorTelepon != "" {
//			p.nomorTelepon = nomorTelepon
//		}
//	}
//
//	err := form.Run()
//	if err != nil {
//		fmt.Println("error in pelanggan's create form:", err)
//	}
//}
