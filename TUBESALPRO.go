package main

import "fmt"

const MAX = 100

type Transaksi struct {
	tanggal string
	berat   float64
}

type Warga struct {
	id          int
	nama        string
	jenisSampah string
	log         [MAX]Transaksi
	nLog        int
	totalBerat  float64
}

type DataWarga struct {
	warga  [MAX]Warga
	nWarga int
}

func main() {
	var D DataWarga
	var pilihan int
	selesai := false

	for !selesai {
		fmt.Println("\n=== Aplikasi Waste-Track ===")
		fmt.Println("1. Tambah Data Warga")
		fmt.Println("2. Ubah Data Warga")
		fmt.Println("3. Hapus Data Warga")
		fmt.Println("4. Catat Setoran Sampah (Mingguan)")
		fmt.Println("5. Cari Warga berdasarkan Nama (Sequential Search)")
		fmt.Println("6. Cari Warga berdasarkan ID (Binary Search)")
		fmt.Println("7. Urutkan Warga berdasarkan Berat (Selection Sort)")
		fmt.Println("8. Urutkan Warga berdasarkan Berat (Insertion Sort)")
		fmt.Println("9. Statistik Total Akumulasi Sampah")
		fmt.Println("10. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahWarga(&D)
		} else if pilihan == 2 {
			ubahWarga(&D)
		} else if pilihan == 3 {
			hapusWarga(&D)
		} else if pilihan == 4 {
			catatSetoran(&D)
		} else if pilihan == 5 {
			cariNama(D)
		} else if pilihan == 6 {
			cariID(&D) 
		} else if pilihan == 7 {
			selectionSortBerat(&D)
			tampilkanData(D)
		} else if pilihan == 8 {
			insertionSortBerat(&D)
			tampilkanData(D)
		} else if pilihan == 9 {
			statistikMingguan(D)
		} else if pilihan == 10 {
			tampilkanData(D)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Waste-Track!")
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahWarga(D *DataWarga) {
	if D.nWarga < MAX {
		fmt.Println("\n-- Tambah Data Warga --")
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&D.warga[D.nWarga].id)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&D.warga[D.nWarga].nama)
		fmt.Print("Masukkan Jenis Sampah Utama: ")
		fmt.Scan(&D.warga[D.nWarga].jenisSampah)
		D.nWarga++
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

func ubahWarga(D *DataWarga) {
	var idCari int
	fmt.Print("\nMasukkan ID Warga yang ingin diubah: ")
	fmt.Scan(&idCari)

	idx := -1
	
	for i := 0; i < D.nWarga && idx == -1; i++ {
		if D.warga[i].id == idCari {
			idx = i
		}
	}

	if idx != -1 {
		fmt.Print("Masukkan Nama Baru: ")
		fmt.Scan(&D.warga[idx].nama)
		fmt.Print("Masukkan Jenis Sampah Baru: ")
		fmt.Scan(&D.warga[idx].jenisSampah)
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}

func hapusWarga(D *DataWarga) {
	var idCari int
	fmt.Print("\nMasukkan ID Warga yang ingin dihapus: ")
	fmt.Scan(&idCari)

	idx := -1
	for i := 0; i < D.nWarga && idx == -1; i++ {
		if D.warga[i].id == idCari {
			idx = i
		}
	}

	if idx != -1 {
		
		for i := idx; i < D.nWarga-1; i++ {
			D.warga[i] = D.warga[i+1]
		}
		D.nWarga--
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}


func catatSetoran(D *DataWarga) {
	var idCari int
	fmt.Print("\nMasukkan ID Warga yang menyetor: ")
	fmt.Scan(&idCari)

	idx := -1
	for i := 0; i < D.nWarga && idx == -1; i++ {
		if D.warga[i].id == idCari {
			idx = i
		}
	}

	if idx != -1 {
		w := &D.warga[idx]
		if w.nLog < MAX {
			fmt.Print("Masukkan Tanggal (cth: 12-05-2024): ")
			fmt.Scan(&w.log[w.nLog].tanggal)
			fmt.Print("Masukkan Berat Sampah (kg): ")
			fmt.Scan(&w.log[w.nLog].berat)

			w.totalBerat += w.log[w.nLog].berat
			w.nLog++
			fmt.Println("Setoran berhasil dicatat!")
		} else {
			fmt.Println("Log setoran untuk warga ini sudah penuh!")
		}
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}


func cariNama(D DataWarga) {
	var namaCari string
	fmt.Print("\nMasukkan Nama Warga yang dicari: ")
	fmt.Scan(&namaCari)

	ditemukan := false
	for i := 0; i < D.nWarga && !ditemukan; i++ {
		if D.warga[i].nama == namaCari {
			fmt.Printf("Ditemukan: ID: %d | Nama: %s | Jenis: %s | Total Berat: %.2f kg\n",
				D.warga[i].id, D.warga[i].nama, D.warga[i].jenisSampah, D.warga[i].totalBerat)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Warga dengan nama tersebut tidak ditemukan.")
	}
}


func sortIDAsc(D *DataWarga) {
	for i := 1; i < D.nWarga; i++ {
		key := D.warga[i]
		j := i - 1
		for j >= 0 && D.warga[j].id > key.id {
			D.warga[j+1] = D.warga[j]
			j--
		}
		D.warga[j+1] = key
	}
}

func cariID(D *DataWarga) {

	sortIDAsc(D)

	var idCari int
	fmt.Print("\nMasukkan ID Warga yang dicari: ")
	fmt.Scan(&idCari)

	kiri := 0
	kanan := D.nWarga - 1
	idxDitemukan := -1

	
	for kiri <= kanan && idxDitemukan == -1 {
		tengah := (kiri + kanan) / 2
		if D.warga[tengah].id == idCari {
			idxDitemukan = tengah
		} else if D.warga[tengah].id < idCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if idxDitemukan != -1 {
		w := D.warga[idxDitemukan]
		fmt.Printf("Ditemukan: ID: %d | Nama: %s | Jenis: %s | Total Berat: %.2f kg\n",
			w.id, w.nama, w.jenisSampah, w.totalBerat)
	} else {
		fmt.Println("Warga dengan ID tersebut tidak ditemukan.")
	}
}

func selectionSortBerat(D *DataWarga) {
	for i := 0; i < D.nWarga-1; i++ {
		idxMaks := i
		for j := i + 1; j < D.nWarga; j++ {
			if D.warga[j].totalBerat > D.warga[idxMaks].totalBerat {
				idxMaks = j
			}
		}
		temp := D.warga[i]
		D.warga[i] = D.warga[idxMaks]
		D.warga[idxMaks] = temp
	}
	fmt.Println("\nBerhasil diurutkan berdasarkan berat (Selection Sort).")
}


func insertionSortBerat(D *DataWarga) {
	for i := 1; i < D.nWarga; i++ {
		key := D.warga[i]
		j := i - 1
	
		for j >= 0 && D.warga[j].totalBerat < key.totalBerat {
			D.warga[j+1] = D.warga[j]
			j--
		}
		D.warga[j+1] = key
	}
	fmt.Println("\nBerhasil diurutkan berdasarkan berat (Insertion Sort).")
}


func statistikMingguan(D DataWarga) {
	var totalSemua float64 = 0
	for i := 0; i < D.nWarga; i++ {
		totalSemua += D.warga[i].totalBerat
	}
	fmt.Printf("\n--- Statistik Mingguan ---\n")
	fmt.Printf("Total Warga yang menyetor  : %d orang\n", D.nWarga)
	fmt.Printf("Total Sampah Terkumpul     : %.2f kg\n", totalSemua)
	fmt.Println("--------------------------")
}

func tampilkanData(D DataWarga) {
	fmt.Println("\n--- Data Seluruh Warga ---")
	if D.nWarga == 0 {
		fmt.Println("Belum ada data warga.")
	} else {
		for i := 0; i < D.nWarga; i++ {
			fmt.Printf("%d. ID: %d | Nama: %s | Jenis Sampah: %s | Total Setoran: %.2f kg\n",
				i+1, D.warga[i].id, D.warga[i].nama, D.warga[i].jenisSampah, D.warga[i].totalBerat)
		}
	}
}
