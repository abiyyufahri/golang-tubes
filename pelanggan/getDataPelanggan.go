package pelanggan

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const dataFilePelanggan = "a_pelanggan.csv"

func (m *ModelPelanggan) GetData() {

	// Baca file CSV
	file, err := os.Open(dataFilePelanggan)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("File data pelanggan ditutup")
		}
	}(file)

	// Buat pembaca CSV
	reader := csv.NewReader(file)

	// Baca semua baris dari file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error membaca file:", err)
		return
	}

	count := 0
	// Iterasi melalui setiap baris (kecuali header)
	for _, record := range records {

		if count >= NMAXPELANGGAN {
			fmt.Println("Jumlah maksimal pelanggan tercapai")
			break
		}

		// Parsing setiap kolom
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error parsing ID:", err)
			return
		}
		status, err := strconv.ParseBool(record[5])
		if err != nil {
			fmt.Println("Error parsing status:", err)
			return
		}

		m.daftarPelanggan[count] = Pelanggan{
			id:           id,
			nama:         record[1],
			alamat:       record[2],
			nomorTelepon: record[3],
			alamatEmail:  record[4],
			status:       status,
		}

		count++
	}

	m.nPelanggan = count

	// Tampilkan hasil
	//for i := 0; i < count; i++ {
	//	fmt.Printf("%+v\n", pelangganArray[i])
	//}
}
