package main

import "fmt"

const NMAX int = 100

type Setoran struct {
	tanggal string
	berat   int
}
type dataWarga struct {
	warga  [NMAX]Warga
	nWarga int
}

type Warga struct {
	nama       string
	id         int
	jumSampah  int
	jSampah    string
	log        [NMAX]Setoran
	nLog       int
	totalBerat int
}


var tabWarga [NMAX]Warga

func main() {
	var ubah, hapus, tambah string
	var n int
	var idTarget int
	var cari, cariid, sortingMenurun, sortingMenaik string
	var catat, statistik string

	var pilihan int
	selesai := false

	// input awal
	fmt.Print("Jumlah warga awal: ")
	fmt.Scan(&n)
	if n > 0 {
		inputData(&tabWarga, n)
	}
	// startup
	for !selesai {
		fmt.Println("\n=== Aplikasi Waste-Track ===")
		fmt.Println("1. Tambah Data Warga")
		fmt.Println("2. Ubah Data Warga")
		fmt.Println("3. Hapus Data Warga")
		fmt.Println("4. Pecarian nama berdasarkan nama (Sequential Search)")
		fmt.Println("5. Pecarian nama berdasarkan ID (Binary Search)")
		fmt.Println("6. Catat Setoran Sampah (Mingguan)")
		fmt.Println("7. Urutkan Warga berdasarkan Berat Menurun(Selection Sort)")
		fmt.Println("8. Urutkan Warga berdasarkan Berat Menaik (Insertion Sort)")
		fmt.Println("9. Statistik Total Akumulasi Sampah")
		fmt.Println("10. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("\nApakah anda ingin menambah data? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&tambah)
			if tambah == "Ya" {
				tambahWarga(&tabWarga, &n)
			}

		} else if pilihan == 2 {
			fmt.Print("Apakah anda ingin mengubah data? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&ubah)
			if ubah == "Ya" {
				fmt.Print("Masukkan ID Warga yang ingin diubah: ")
				fmt.Scan(&idTarget)
				ubahWarga(&tabWarga, n, idTarget)
			}
		} else if pilihan == 3 {
			fmt.Print("Apakah anda ingin menghapus data? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&hapus)
			if hapus == "Ya" {
				fmt.Print("Masukkan ID yang ingin dihapus: ")
				fmt.Scan(&idTarget)
				hapusWarga(&tabWarga, &n, idTarget)
			}
		} else if pilihan == 4 {
			fmt.Print("Apakah anda ingin melakukan pencarian nama? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&cari)
			if cari == "Ya" {
				cariNama(tabWarga, n)
			}
		} else if pilihan == 5 {
			fmt.Print("Apakah anda ingin melakukan pencarian ID? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&cariid)
			if cariid == "Ya" {
				fmt.Print("Masukkan ID Warga yang dicari: ")
				fmt.Scan(&idTarget)
				cariID(&tabWarga, n, idTarget) 
			}
		} else if pilihan == 6 {
			fmt.Print("Apakah anda ingin mencatat setoran mingguan? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&catat)
			if catat == "Ya" {
				catatSetoran(&tabWarga, n)
			}
		} else if pilihan == 7 {
			fmt.Print("Apakah anda ingin melakukan Pengurutan Menurun/Selection Sort? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&sortingMenurun)
			if sortingMenurun == "Ya" {
				selectionSortBerat(&tabWarga, n)
				tampilkanData(tabWarga, n)
			}
		} else if pilihan == 8 {
			fmt.Print("Apakah anda ingin melakukan Pengurutan Menaik/Insertion Sort? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&sortingMenaik)
			if sortingMenaik == "Ya" {
				insertionSortBerat(&tabWarga, n)
				tampilkanData(tabWarga, n)
			}
		} else if pilihan == 9 {
			fmt.Print("Apakah anda ingin melihat Statistik Mingguan? (Ketik 'Ya'/'Tidak'): ")
			fmt.Scan(&statistik)
			if statistik == "Ya" {
				statistikMingguan(tabWarga, n)
			}

		} else if pilihan == 10 {
			tampilkanData(tabWarga, n)
		} else if pilihan == 0 {
			fmt.Println("Program Selesai")
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}


func inputData(tab *[NMAX]Warga, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("Masukkan Data Warga (Nama ID Jumlah-Sampah Jenis-Sampah): ")
		fmt.Scan(&tab[i].nama, &tab[i].id, &tab[i].jumSampah, &tab[i].jSampah)

	}
}

func ubahWarga(tab *[NMAX]Warga, n int, id int) {
	var i int
	var ketemu bool
	var x string

	ketemu = false

	for i = 0; i < n && !ketemu; i++ {
		if tab[i].id == id {
			ketemu = true

			fmt.Print("Apa yang ingin diubah? (Ketik 'Nama'/'ID'/'Berat'/'Jenis Sampah'/'Jumlah Sampah'): ")
			fmt.Scan(&x)

			if x == "Nama" {
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&tab[i].nama)
			} else if x == "ID" {
				fmt.Print("Masukkan ID Baru: ")
				fmt.Scan(&tab[i].id)
			} else if x == "Berat" {
				fmt.Print("Masukkan Berat Baru: ")
				fmt.Scan(&tab[i].totalBerat)
			} else if x == "Jenis Sampah" {
				fmt.Print("Masukkan Jenis Sampah Baru: ")
				fmt.Scan(&tab[i].jSampah)
			} else if x == "Jumlah Sampah" {
				fmt.Print("Masukkan Jumlah Sampah Baru: ")
				fmt.Scan(&tab[i].jumSampah)
			}
		}
	}

	if ketemu {
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusWarga(tab *[NMAX]Warga, n *int, id int) {
	var i, idxHapus int
	var ketemu bool

	ketemu = false
	idxHapus = -1

	for i = 0; i < *n && !ketemu; i++ {
		if tab[i].id == id {
			idxHapus = i
			ketemu = true
		}
	}

	if ketemu {
		for i = idxHapus; i < *n-1; i++ {
			tab[i] = tab[i+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func tambahWarga(tab *[NMAX]Warga, n *int) {
	if *n < NMAX {
		fmt.Print("Masukkan Data Baru (Nama ID Jumlah-Sampah Jenis-Sampah): ")
		fmt.Scan(&tab[*n].nama, &tab[*n].id, &tab[*n].jumSampah, &tab[*n].jSampah)

		*n = *n + 1
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas data sudah penuh!")
	}
}

func cariNama(tab [NMAX]Warga, n int) {
	var i int
	var ketemu bool
	var key string

	ketemu = false
	
	fmt.Print("Nama Dicari: ")
	fmt.Scan(&key)

	for i = 0; i < n && !ketemu; i++ {
		if tab[i].nama == key {
			ketemu = true
			fmt.Println("\nData Ditemukan!")
			fmt.Println("ID:", tab[i].id, ", Nama:", tab[i].nama, ", Berat:", tab[i].totalBerat, " ,Jumlah Sampah", tab[i].jumSampah, ", Jenis Sampah:", tab[i].jSampah)
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariID(tab *[NMAX]Warga, n int, key int) {
	var left, right, mid int

	
	sortingID(tab, n)

	left = 0
	right = n - 1
	mid = (left + right) / 2

	for left <= right && tab[mid].id != key {
		if key < tab[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if mid >= 0 && tab[mid].id == key {
		fmt.Println("\nData Ditemukan!")
		fmt.Println("ID:", tab[mid].id, ", Nama:", tab[mid].nama, ", Berat:", tab[mid].totalBerat, ", Jumlah Sampah:", tab[mid].jumSampah, ", Jenis Sampah:", tab[mid].jSampah)

	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func sortingID(tab *[NMAX]Warga, n int) {
	var i, x int
	var key Warga

	for i = 1; i < n; i++ {
		key = tab[i]
		x = i - 1

		for x >= 0 && tab[x].id > key.id {
			tab[x+1] = tab[x]
			x = x - 1
		}
		tab[x+1] = key
	}
}


func selectionSortBerat(tab *[NMAX]Warga, n int) {
	var i, x, idxMaks int
	var temp Warga

	for i = 0; i < n-1; i++ {
		idxMaks = i
		for x = i + 1; x < n; x++ {
			
			if tab[x].totalBerat > tab[idxMaks].totalBerat {
				idxMaks = x
			}
		}
		temp = tab[i]
		tab[i] = tab[idxMaks]
		tab[idxMaks] = temp
	}
	fmt.Println("Menampilkan berat terurut menurun")
}


func insertionSortBerat(tab *[NMAX]Warga, n int	) {
	var i, x int
	var key Warga
	for i = 1; i < n; i++ {
		key = tab[i]
		x = i - 1

		
		for x >= 0 && tab[x].totalBerat > key.totalBerat {
			tab[x+1] = tab[x]
			x = x - 1
		}
		tab[x+1] = key
	}
	fmt.Println("Menampilkan berat menaik")
}

func catatSetoran(tab *[NMAX]Warga, n int) {
	var idTarget, indexWarga, i int
	var ketemu bool

	fmt.Print("Masukkan ID Warga yang menyetor: ")
	fmt.Scan(&idTarget)

	ketemu = false
	indexWarga = -1
	for i = 0; i < n && !ketemu; i++ {
		if tab[i].id == idTarget {
			ketemu = true
			indexWarga = i
		}
	}

	if ketemu {
		if tab[indexWarga].nLog < NMAX {
			fmt.Print("Masukkan Tanggal Setoran (cth: 07-09-2026): ")
			fmt.Scan(&tab[indexWarga].log[tab[indexWarga].nLog].tanggal)

			fmt.Print("Masukkan Berat Setoran: ")
			fmt.Scan(&tab[indexWarga].log[tab[indexWarga].nLog].berat)

			tab[indexWarga].totalBerat = tab[indexWarga].totalBerat + tab[indexWarga].log[tab[indexWarga].nLog].berat
			tab[indexWarga].nLog = tab[indexWarga].nLog + 1

			fmt.Println("Setoran berhasil dicatat!")
		} else {
			fmt.Println("Riwayat setoran sudah penuh!")
		}
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}


func statistikMingguan(tab [NMAX]Warga, n int) {
	var cariTanggal string
	var i, j, total int

	fmt.Print("Masukkan Tanggal/Minggu yang ingin dihitung: ")
	fmt.Scan(&cariTanggal)

	total = 0
	for i = 0; i < n; i++ {
		for j = 0; j < tab[i].nLog; j++ {
			if tab[i].log[j].tanggal == cariTanggal {
				total = total + tab[i].log[j].berat
			}
		}
	}
	fmt.Printf("\n--- Statistik ---\n")
	fmt.Printf("Total Sampah Terkumpul pada %s adalah %d kg\n", cariTanggal, total)
}


func tampilkanData(tab [NMAX]Warga, n int) {
	var i int
	fmt.Println("\n--- Data Seluruh Warga ---")
	if n == 0 {
		fmt.Println("Data Kosong.")
	} else {
		for i = 0; i < n; i++ {
			fmt.Printf("ID: %d | Nama: %s | Total Berat: %d kg | Jenis: %s | Jumlah: %d\n",
				tab[i].id, tab[i].nama, tab[i].totalBerat, tab[i].jSampah, tab[i].jumSampah)
		}
	}
	fmt.Println("--------------------------")
}
