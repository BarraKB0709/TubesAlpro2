package main

import "fmt"

const NMAX int = 1000

type Warga struct{
	nama      string
	id        int
	berat     int
	jumSampah int
	jSampah   string
	tanggal   string
}

var tabWarga [NMAX]Warga

func main(){
	var n, idTarget, pilihan int
	var selesai bool
	selesai = false

	fmt.Print("Jumlah warga awal: ")
	fmt.Scan(&n)
	if n > 0{
		inputData(&tabWarga, n)
	}

	fmt.Println("\nAplikasi Waste-Track")
	fmt.Println("1. Tambah Data Warga")
	fmt.Println("2. Ubah Data Warga")
	fmt.Println("3. Hapus Data Warga")
	fmt.Println("4. Pencarian Berdasarkan Nama (Sequential Search)")
	fmt.Println("5. Pencarian Berdasarkan ID (Binary Search)")
	fmt.Println("6. Catat Setoran Sampah & Tanggal (Mingguan)")
	fmt.Println("7. Urutkan Warga berdasarkan Berat Menurun (Selection Sort)")
	fmt.Println("8. Urutkan Warga berdasarkan Berat Menaik (Insertion Sort)")
	fmt.Println("9. Statistik Total Akumulasi Sampah & Min-Max")
	fmt.Println("10. Tampilkan Semua Data")
	fmt.Println("0. Keluar")

	for !selesai{
		fmt.Println("Pilih menu:")
		fmt.Scan(&pilihan)

		if pilihan == 1{
			tambahWarga(&tabWarga, &n)
		} else if pilihan == 2{
			fmt.Print("Masukkan ID Warga yang ingin diubah: ")
			fmt.Scan(&idTarget)
			ubahWarga(&tabWarga, n, idTarget)
		} else if pilihan == 3{
			fmt.Print("Masukkan ID Warga yang ingin dihapus: ")
			fmt.Scan(&idTarget)
			hapusWarga(&tabWarga, &n, idTarget)
		} else if pilihan == 4{
			cariNama(tabWarga, n)
		} else if pilihan == 5{
			fmt.Print("Masukkan ID Warga yang dicari: ")
			fmt.Scan(&idTarget)
			cariID(&tabWarga, n, idTarget)
		} else if pilihan == 6{
			catatSetoran(&tabWarga, n)
		} else if pilihan == 7{
			selectionSortBerat(&tabWarga, n)
			tampilkanData(tabWarga, n)
		} else if pilihan == 8{
			insertionSortBerat(&tabWarga, n)
			tampilkanData(tabWarga, n)
		} else if pilihan == 9{
			statistikMingguan(tabWarga, n)
		} else if pilihan == 10{
			tampilkanData(tabWarga, n)
		} else if pilihan == 0{
			fmt.Println("Program Selesai.")
			selesai = true
		} else{
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputData(tab *[NMAX]Warga, n int){
	var i int
	fmt.Println("Masukkan Data (Nama ID Berat JumlahKantong Jenis Tanggal[DD-MM-YYYY]): ")
	for i = 0; i < n; i++{
		fmt.Scan(&tab[i].nama, &tab[i].id, &tab[i].berat, &tab[i].jumSampah, &tab[i].jSampah, &tab[i].tanggal)
	}
}

func tambahWarga(tab *[NMAX]Warga, n *int){
	if *n < NMAX{
		fmt.Print("Masukkan Data Baru (Nama ID Berat JumlahKantong Jenis Tanggal[DD-MM-YYYY]): ")
		fmt.Scan(&tab[*n].nama, &tab[*n].id, &tab[*n].berat, &tab[*n].jumSampah, &tab[*n].jSampah, &tab[*n].tanggal)
		*n = *n + 1
		fmt.Println("Data berhasil ditambahkan!")
	} else{
		fmt.Println("Kapasitas data sudah penuh!")
	}
}

func ubahWarga(tab *[NMAX]Warga, n int, id int){
	var i int
	var ketemu bool
	var x string

	ketemu = false
	for i = 0; i < n && !ketemu; i++{
		if tab[i].id == id{
			ketemu = true
			fmt.Print("Apa yang ingin diubah? (Nama/ID/Berat/Jenis/Tanggal): ")
			fmt.Scan(&x)

			if x == "Nama"{
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&tab[i].nama)
			} else if x == "ID"{
				fmt.Print("Masukkan ID Baru: ")
				fmt.Scan(&tab[i].id)
			} else if x == "Berat"{
				fmt.Print("Masukkan Berat Baru: ")
				fmt.Scan(&tab[i].berat)
			} else if x == "Tanggal"{
				fmt.Print("Masukkan Tanggal Baru (DD-MM-YYYY): ")
				fmt.Scan(&tab[i].tanggal)
			} else{
				fmt.Print("Masukkan Jenis Sampah Baru: ")
				fmt.Scan(&tab[i].jSampah)
			}
		}
	}

	if ketemu{
		fmt.Println("Data berhasil diubah!")
	} else{
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusWarga(tab *[NMAX]Warga, n *int, id int){
	var i, idxHapus int
	var ketemu bool

	ketemu = false
	idxHapus = -1
	for i = 0; i < *n && !ketemu; i++{
		if tab[i].id == id{
			idxHapus = i
			ketemu = true
		}
	}

	if ketemu{
		for i = idxHapus; i < *n-1; i++{
			tab[i] = tab[i+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus!")
	} else{
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariNama(tab [NMAX]Warga, n int){
	var i int
	var ketemu bool
	var key string

	ketemu = false
	fmt.Print("Nama Dicari: ")
	fmt.Scan(&key)

	for i = 0; i < n; i++{
		if tab[i].nama == key{
			ketemu = true
			fmt.Println("Data Ditemukan!")
			fmt.Println("ID:", tab[i].id, "| Nama:", tab[i].nama, "| Berat:", tab[i].berat, "kg | Jumlah Kantong:", tab[i].jumSampah, "| Jenis Sampah:", tab[i].jSampah, "| Tgl Transaksi:", tab[i].tanggal)
		}
	}

	if !ketemu{
		fmt.Println("Data tidak ditemukan.")
	}
}

func sortingID(tab *[NMAX]Warga, n int){
	var i, x int
	var key Warga

	for i = 1; i < n; i++{
		key = tab[i]
		x = i - 1
		for x >= 0 && tab[x].id > key.id{
			tab[x+1] = tab[x]
			x = x - 1
		}
		tab[x+1] = key
	}
}

func cariID(tab *[NMAX]Warga, n int, key int){
	var left, right, mid int

	sortingID(tab, n)

	left = 0
	right = n - 1
	mid = (left + right) / 2

	for left <= right && tab[mid].id != key{
		if key < tab[mid].id{
			right = mid - 1
		} else{
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if left <= right && tab[mid].id == key{
		fmt.Println("\nData Ditemukan!")
		fmt.Println("ID:", tab[mid].id, "| Nama:", tab[mid].nama, "| Berat:", tab[mid].berat, "kg | Jenis Sampah:", tab[mid].jSampah, "| Tgl Transaksi:", tab[mid].tanggal)
	} else{
		fmt.Println("Data tidak ditemukan.")
	}
}

func catatSetoran(tab *[NMAX]Warga, n int){
	var idTarget, i, tambahanBerat, tambahanKantong int
	var tglBaru string
	var ketemu bool

	fmt.Print("Masukkan ID Warga: ")
	fmt.Scan(&idTarget)

	ketemu = false
	for i = 0; i < n && !ketemu; i++{
		if tab[i].id == idTarget{
			ketemu = true
			fmt.Print("Masukkan tambahan berat sampah (kg): ")
			fmt.Scan(&tambahanBerat)
			fmt.Print("Masukkan tambahan jumlah kantong: ")
			fmt.Scan(&tambahanKantong)
			fmt.Print("Masukkan Tanggal Transaksi Baru (DD-MM-YYYY): ")
			fmt.Scan(&tglBaru)

			tab[i].berat = tab[i].berat + tambahanBerat
			tab[i].jumSampah = tab[i].jumSampah + tambahanKantong
			tab[i].tanggal = tglBaru

			fmt.Println("Setoran sampah berhasil diperbarui!")
		}
	}
	if !ketemu{
		fmt.Println("ID Warga tidak ditemukan.")
	}
}

func selectionSortBerat(tab *[NMAX]Warga, n int){
	var i, x, idxMaks int
	var temp Warga

	for i = 0; i < n-1; i++{
		idxMaks = i
		for x = i + 1; x < n; x++{
			if tab[x].berat > tab[idxMaks].berat{
				idxMaks = x
			}
		}
		temp = tab[i]
		tab[i] = tab[idxMaks]
		tab[idxMaks] = temp
	}
	fmt.Println("Menampilkan berat terurut menurun.")
}

func insertionSortBerat(tab *[NMAX]Warga, n int){
	var i, x int
	var key Warga

	for i = 1; i < n; i++{
		key = tab[i]
		x = i - 1

		for x >= 0 && tab[x].berat > key.berat{
			tab[x+1] = tab[x]
			x = x - 1
		}
		tab[x+1] = key
	}
	fmt.Println("Menampilkan berat terurut menaik.")
}

func findMinSampah(tab [NMAX]Warga, n int) int{
	var i, idxMin int
	idxMin = 0
	for i = 1; i < n; i++{
		if tab[i].jumSampah < tab[idxMin].jumSampah{
			idxMin = i
		}
	}
	return idxMin
}

func findMaxSampah(tab [NMAX]Warga, n int) int{
	var i, idxMax int
	idxMax = 0
	for i = 1; i < n; i++{
		if tab[i].jumSampah > tab[idxMax].jumSampah{
			idxMax = i
		}
	}
	return idxMax
}

func statistikMingguan(tab [NMAX]Warga, n int){
	var i, totalBerat int
	var idxMin, idxMax int

	if n == 0{
		fmt.Println("Belum ada data warga.")
		return
	}

	totalBerat = 0
	for i = 0; i < n; i++{
		totalBerat = totalBerat + tab[i].berat
	}

	idxMin = findMinSampah(tab, n)
	idxMax = findMaxSampah(tab, n)

	fmt.Println("\nSTATISTIK AKUMULASI SAMPAH")
	fmt.Println("Total Akumulasi Berat Sampah:", totalBerat, "kg")
	fmt.Println("Warga Jumlah Sampah terbanyak: Nama:", tab[idxMax].nama, "| Total Kantong:", tab[idxMax].jumSampah, "kantong")
	fmt.Println("Warga Jumlah Sampah terdikit: Nama:", tab[idxMin].nama, "| Total Kantong:", tab[idxMin].jumSampah, "kantong")
}

func tampilkanData(tab [NMAX]Warga, n int){
	var i int
	if n == 0{
		fmt.Println("Data warga kosong.")
		return
	}
	fmt.Println("\nDATA SELURUH WARGA")
	for i = 0; i < n; i++{
		fmt.Printf("ID: %d | Nama: %s | Berat Total: %d kg | Jumlah Kantong: %d | Jenis: %s | Tgl Transaksi: %s\n",
			tab[i].id, tab[i].nama, tab[i].berat, tab[i].jumSampah, tab[i].jSampah, tab[i].tanggal)
	}
}
