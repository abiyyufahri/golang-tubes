package pelanggan

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

func create_form(p *Pelanggan) {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nama").
				Description("Masukkan nama pelanggan Anda").
				Placeholder("Cth: Yanto Sumargo").
				Value(&p.nama),

			huh.NewInput().
				Title("Alamat").
				Description("Tempat pengiriman ke pelanggan").
				Placeholder("Cth: Jl. Radio Palasari No 34 Gg. ABC").
				Value(&p.alamat),

			huh.NewInput().
				Title("Email").
				Description("Email untuk informasikan ke pengguna").
				Placeholder("Cth: abcd@mantapjiwa.com").
				Value(&p.alamatEmail),
		),

		huh.NewGroup(
			//huh.NewSelect[bool]().
			//	Title("Status").
			//	Description("Bila nonaktif, pengguna tidak bisa dibuatkan ekspedisi").
			//	Options(
			//		huh.NewOption[bool]("Aktif", true),
			//		huh.NewOption[bool]("Nonaktif", false),
			//	).
			//	Value(&p.status),

			huh.NewConfirm().
				Title("Aktifkan Pengguna?").
				Description("Bila tidak aktif, maka pengguna tidak akan dapat dibuatkan ekspedisi").
				Value(&p.status).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("error in pelanggan's create form:", err)
	}
}
