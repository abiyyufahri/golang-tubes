package ekspedisi

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func (e *ModelEkspedisi) GetData() {
	err := e.readEkspedisiFromCSV("a_ekspedisi.csv")
	if err != nil {
		return
	}
}

func (e *ModelEkspedisi) readEkspedisiFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			// Skip header
			continue
		}
		if e.nEkspedisi >= NMAX_EKSPEDISI {
			fmt.Println("Jumlah maksimal ekspedisi tercapai")
			break
		}

		// Parsing fields
		idPelanggan := record[1]

		jenisLayanan, err := strconv.Atoi(record[2])
		if err != nil {
			return fmt.Errorf("error parsing jenisLayanan: %v", err)
		}

		biayaEkspedisi, err := strconv.Atoi(record[6])
		if err != nil {
			return fmt.Errorf("error parsing biayaEkspedisi: %v", err)
		}

		status, err := strconv.Atoi(record[7])
		if err != nil {
			return fmt.Errorf("error parsing status: %v", err)
		}

		ekspedisi := Ekspedisi{
			resi:            record[0],
			idPelanggan:     idPelanggan,
			jenisLayanan:    Layanan(jenisLayanan),
			deskripsiBarang: record[3],
			alamatAsal:      record[4],
			alamatTujuan:    record[5],
			biayaEkspedisi:  biayaEkspedisi,
			status:          status,
		}

		e.DaftarEkspedisi[e.nEkspedisi] = ekspedisi
		e.nEkspedisi++
	}

	return nil
}
