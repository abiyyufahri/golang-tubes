package main

import "fmt"

func main() {
	var opsi int
	opsi = menuPage()

	switch opsi {
	case 1:
		menuPelanggan()
	case 2:
		menuEkspedisi()
	}

}

func menuPage() int {
	var opsi int
	fmt.Println("Pelanggan (1), Ekspedisi (2): ")
	fmt.Scan(&opsi)
	return opsi
}

func menuPelanggan() {
	var opsi int
	fmt.Println("Tambah,Detail,Hapus,Kembali (1/2/3/4): ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		addPelanggan()
	case 2:
		detailPelanggan()
	case 3:

	case 4:

	}

}

func addPelanggan() {
	var nama, alamat, nomorTelepon, alamatEmail string

	for nama != "0" {
		fmt.Scan(&nama, &alamat, &nomorTelepon, &alamatEmail)
		if nama == "0" {
			break
		}
		CreatePelanggan(nama, alamat, nomorTelepon, alamatEmail)
	}
	fmt.Println("Data yang ditambahkan : ")

	// Menampilkan semua pelanggan yang telah ditambahkan
	for i := 0; i < nP; i++ {
		pelanggan := ReadAllPelanggan()
		fmt.Printf("Pelanggan %d: %v\n", i+1, pelanggan[i])
	}

	menuPelanggan()
}

func detailPelanggan() {
	fmt.Println("Data Pelanggan : ")
	var iPelanggan int
	// Menampilkan semua pelanggan yang telah ditambahkan
	for i := 0; i < nP; i++ {
		pelanggan := ReadAllPelanggan()
		fmt.Printf("Pelanggan %d: %v\n", i+1, pelanggan[i])
	}

	fmt.Print("Detail, Update, Delete Pelanggan (1/2/3) ")
	fmt.Scan(&iPelanggan)

	if iPelanggan == 1 {
		fmt.Print("lihat detail pelanggan (index) (kembali : -1): ")
		fmt.Scan(&iPelanggan)

		if iPelanggan == -1 {
			menuPelanggan()
		}
		fmt.Println()
		for i := 0; i < nP; i++ {
			pelanggan := ReadAllPelanggan()
			if iPelanggan-1 == i {
				fmt.Println("Detail Pelanggan", i+1, ":")
				fmt.Println("Nama  :", pelanggan[i].nama)
				fmt.Println("Alamat:", pelanggan[i].alamat)
				fmt.Println("Telp  :", pelanggan[i].nomorTelepon)
				fmt.Println("Email :", pelanggan[i].alamatEmail)
				break
			} else {
				continue
			}
		}
	} else if iPelanggan == 2 {
		var nama, alamat, nomorTelepon, alamatEmail string
		var id int
		fmt.Print("ID Pelanggan : ")
		fmt.Scan(&id)
		fmt.Println()
		fmt.Print("Update (nama,alamat,nomorTelp,Email): ")
		fmt.Scan(&nama, &alamat, &nomorTelepon, &alamatEmail)
		UpdatePelanggan(id, nama, alamat, nomorTelepon, alamatEmail)
	} else if iPelanggan == 3 {
		var id int
		fmt.Print("ID Pelanggan : ")
		fmt.Scan(&id)
		fmt.Println()
		DeletePelanggan(id)
		fmt.Println("Pelanggan berhasil dihapus")
		fmt.Println()
	}

	fmt.Println()
	detailPelanggan()
}

func menuEkspedisi() {
	var opsi int
	fmt.Println("Lihat status,update status,Edit Ekspedisi,Hapus Ekspedisi,Kembali (1/2/3/4/5): ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	}
}
