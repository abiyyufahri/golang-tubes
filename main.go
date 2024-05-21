package main

import "fmt"

func main() {
	for {
		opsi := menuPage()
		switch opsi {
		case 1:
			menuPelanggan()
		case 2:
			menuEkspedisi()
		case 3:
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuPage() int {
	var opsi int
	fmt.Println("======== Menu Utama ========")
	fmt.Println("1. Pelanggan")
	fmt.Println("2. Ekspedisi")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih opsi (1/2/3): ")
	fmt.Scan(&opsi)
	return opsi
}

func menuPelanggan() {
	for {
		var opsi int
		fmt.Println("\n======== Menu Pelanggan ========")
		fmt.Println("1. Tambah Pelanggan")
		fmt.Println("2. Detail Pelanggan")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi (1/2/3): ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			addPelanggan()
		case 2:
			detailPelanggan()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func addPelanggan() {
	var nama, alamat, nomorTelepon, alamatEmail string
	var status int
	fmt.Println("\nMasukkan data pelanggan (nama, alamat, nomor telepon, email, status)")
	fmt.Println("Ketik '0' untuk nama untuk kembali ke menu sebelumnya.")
	for {
		fmt.Print("Nama: ")
		fmt.Scan(&nama)
		if nama == "0" {
			break
		}
		fmt.Print("Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print("Nomor Telepon: ")
		fmt.Scan(&nomorTelepon)
		fmt.Print("Email: ")
		fmt.Scan(&alamatEmail)
		fmt.Print("Status (1: Selesai, 2: Dikirim, 3: Pending): ")
		fmt.Scan(&status)
		CreatePelanggan(nama, alamat, nomorTelepon, alamatEmail, status)
		fmt.Println("Pelanggan berhasil ditambahkan.")
	}

	fmt.Println("\nData pelanggan yang ditambahkan:")
	WritePelanggan()
}

func WritePelanggan() {
	var i int
	for i = 0; i < nP; i++ {
		var status string
		pelanggan := ReadAllPelanggan()
		switch pelanggan[i].status {
		case 1:
			status = "Selesai"
		case 2:
			status = "Dikirim"
		case 3:
			status = "Pending"
		default:
			status = "Tidak Ada"
		}
		fmt.Printf("Pelanggan %d: ID: %d, Nama: %s, Alamat: %s, Telp: %s, Email: %s, Status: %s\n",
			i+1, pelanggan[i].id, pelanggan[i].nama, pelanggan[i].alamat, pelanggan[i].nomorTelepon, pelanggan[i].alamatEmail, status)
	}
	if i == 0 {
		fmt.Println("!! Tidak ada data pelanggan. !!")
	}
}

func detailPelanggan() {
	for {
		fmt.Println("\nData Pelanggan:")
		WritePelanggan()
		var opsi int
		fmt.Print("Detail, Update, Delete Pelanggan, Sorting Data (1/2/3/4), Kembali (5): ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			viewDetailPelanggan()
		case 2:
			updatePelanggan()
		case 3:
			deletePelanggan()
		case 4:
			sortPelanggan()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func viewDetailPelanggan() {
	var iPelanggan int
	fmt.Print("Masukkan nomor pelanggan untuk melihat detail (kembali: -1): ")
	fmt.Scan(&iPelanggan)
	if iPelanggan == -1 {
		return
	}
	pelanggan := ReadAllPelanggan()
	if iPelanggan > 0 && iPelanggan <= nP {
		i := iPelanggan - 1
		fmt.Printf("\nDetail Pelanggan %d:\n", iPelanggan)
		fmt.Printf("Nama  : %s\n", pelanggan[i].nama)
		fmt.Printf("Alamat: %s\n", pelanggan[i].alamat)
		fmt.Printf("Telp  : %s\n", pelanggan[i].nomorTelepon)
		fmt.Printf("Email : %s\n", pelanggan[i].alamatEmail)
		fmt.Printf("Status: %d\n", pelanggan[i].status)
	} else {
		fmt.Println("Nomor pelanggan tidak valid.")
	}
}

func updatePelanggan() {
	var opsi int
	var id int
	fmt.Println("\n======== Update Pelanggan ========")
	fmt.Println("1. Update Status Pelanggan ")
	fmt.Println("2. Update Informasi Pelanggan")
	fmt.Print("Pilih opsi (1/2) Kembali (0): ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		var status int
		fmt.Print("Masukkan ID Pelanggan yang akan diupdate: ")
		fmt.Scan(&id)
		fmt.Println("[1] Selesai")
		fmt.Println("[2] Dikirim")
		fmt.Println("[3] Pending")
		fmt.Print("Masukkan status (1/2/3): ")
		fmt.Scan(&status)
		UpdateStatusPelanggan(id, status)
	case 2:
		var nama, alamat, nomorTelepon, alamatEmail string
		fmt.Print("Masukkan ID Pelanggan yang akan diupdate: ")
		fmt.Scan(&id)
		fmt.Println("\n!! Ketik '0' pada masukan untuk skip !!")
		fmt.Println("Masukkan data baru (nama, alamat, nomor telepon, email):")
		fmt.Print("Nama: ")
		fmt.Scan(&nama)
		fmt.Print("Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print("Nomor Telepon: ")
		fmt.Scan(&nomorTelepon)
		fmt.Print("Email: ")
		fmt.Scan(&alamatEmail)
		UpdatePelanggan(id, nama, alamat, nomorTelepon, alamatEmail)
	default:
		menuPelanggan()
	}

}

func deletePelanggan() {
	var id int
	fmt.Print("Masukkan ID Pelanggan yang akan dihapus: ")
	fmt.Scan(&id)
	DeletePelanggan(id)
}

func sortPelanggan() {
	var status int
	fmt.Println("\nStatus yang disortir:")
	fmt.Println("[1] Selesai")
	fmt.Println("[2] Dikirim")
	fmt.Println("[3] Pending")
	fmt.Println("[0] Tidak Ada")
	fmt.Print("Pilih status: ")
	fmt.Scan(&status)
	SortingPelanggan(&daftarPelanggan, status)
}

func menuEkspedisi() {
	for {
		var opsi int
		fmt.Println("\n======== Menu Ekspedisi ========")
		fmt.Println("1. Tambah Ekspedisi")
		fmt.Println("2. Lihat Ekspedisi")
		fmt.Println("3. Update Ekspedisi")
		fmt.Println("4. Hapus Ekspedisi")
		fmt.Println("5. Sortir Ekspedisi")
		fmt.Println("6. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi (1/2/3/4/5/6): ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			addEkspedisi()
		case 2:
			viewEkspedisi()
		case 3:
			updateEkspedisi()
		case 4:
			doDeleteEkspedisi()
		case 5:
			sortEkspedisi()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func addEkspedisi() {
	var idPelanggan, status int
	var jenisLayanan, biayaEkspedisi string
	fmt.Print("Masukkan ID Pelanggan: ")
	fmt.Scan(&idPelanggan)
	fmt.Print("Masukkan Jenis Layanan: ")
	fmt.Scan(&jenisLayanan)
	fmt.Print("Masukkan Biaya Ekspedisi: ")
	fmt.Scan(&biayaEkspedisi)
	fmt.Print("Masukkan Status (1: Selesai, 2: Dikirim, 3: Pending): ")
	fmt.Scan(&status)
	createEkspedisi(idPelanggan, jenisLayanan, biayaEkspedisi, status)
}

func viewEkspedisi() {
	var status int
	var statusStr string
	ekspedisis := ReadEkspedisi()
	pelanggan := ReadAllPelanggan()
	if nE == 0 {
		fmt.Println("Tidak ada data ekspedisi.")
		return
	}
	for i := 0; i < nE; i++ {
		for j := 0; j < nP; j++ {

			status = ekspedisis[i].status
			switch status {
			case 1:
				statusStr = "Selesai"
			case 2:
				statusStr = "Dikirim"
			case 3:
				statusStr = "Pending"
			default:
				statusStr = "Tidak Ada"
			}

			if ekspedisis[i].idPelanggan == pelanggan[j].id {
				fmt.Printf("Ekspedisi %d: ID: %d, ID Pelanggan: %d,Nama Pelanggan: %s, Jenis Layanan: %s, Biaya: %s, Status: %s\n",
					i+1, ekspedisis[i].id, ekspedisis[i].idPelanggan, pelanggan[j].nama, ekspedisis[i].jenisLayanan, ekspedisis[i].biayaEkspedisi, statusStr)
			}
		}
	}
}

func updateEkspedisi() {
	var id, status int
	var jenisLayanan, biayaEkspedisi string
	fmt.Print("Masukkan ID Ekspedisi yang akan diupdate: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Jenis Layanan baru (atau ketik '0' untuk melewati): ")
	fmt.Scan(&jenisLayanan)
	fmt.Print("Masukkan Biaya Ekspedisi baru (atau ketik '0' untuk melewati): ")
	fmt.Scan(&biayaEkspedisi)
	fmt.Print("Masukkan Status baru (1: Selesai, 2: Dikirim, 3: Pending, atau ketik 0 untuk melewati): ")
	fmt.Scan(&status)
	UpdateEkspedisi(id, jenisLayanan, biayaEkspedisi, status)
}

func doDeleteEkspedisi() {
	var id int
	fmt.Print("Masukkan ID Ekspedisi yang akan dihapus: ")
	fmt.Scan(&id)
	DeleteEkspedisi(id)
	fmt.Println("Ekspedisi berhasil dihapus.")
}

func sortEkspedisi() {
	var opsi int
	fmt.Println("\nPilihan pengurutan:")
	fmt.Println("1. Urutkan berdasarkan ID Pelanggan")
	fmt.Println("2. Urutkan berdasarkan Status (prioritaskan status tertentu)")
	fmt.Print("Pilih opsi (1/2): ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		InsertionSortByIDPelanggan(&daftarEkspedisi)
		fmt.Println("Data ekspedisi telah diurutkan berdasarkan ID Pelanggan.")
	case 2:
		var status int
		fmt.Println("\nStatus yang disortir:")
		fmt.Println("[1] Selesai")
		fmt.Println("[2] Dikirim")
		fmt.Println("[3] Pending")
		fmt.Print("Pilih status: ")
		fmt.Scan(&status)
		InsertionSortByStatus(&daftarEkspedisi, status)
		fmt.Println("Data ekspedisi telah diurutkan berdasarkan status.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}
