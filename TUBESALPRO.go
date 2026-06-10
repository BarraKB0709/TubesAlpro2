package main
import "fmt"

const NMAX int = 100
type Warga struct{
	nama string
	id int
	berat int
	jSampah string
}
var tabWarga [NMAX]Warga





func main() {
	var nama, jSampah, ubah, hapus string
	var id, i, n, berat int

	fmt.Print("Jumlah warga:")
	fmt.Scan(&n)
	
	for i = 0; i < n; i++{
		fmt.Scan(&tabWarga[i].nama, &tabWarga[i].id, &tabWarga[i].jSampah, &tabWarga[i].berat)
	}
	
	fmt.Println("Ubah Warga:")
	fmt.Scan(&ubah)
	if ubah == "ubah"{
		fmt.Scan(&idxUbah)
		ubahWarga(tab *[NMAX]Warga, n int, id int, namaBaru string)
	
	}

	fmt.Scan(&hapus)
	if hapus == "hapus"{
		fmt.Scan(&idxHapus)
		hapusWarga(tab *[NMAX]Warga, n *int, id int)
	}




func ubahWarga(tab *[NMAX]Warga, n int, id int, namaBaru string) {
	var i int
	for i = 0; i < n; i++ {
		if tab[i].id == id {
			tab[i].nama = namaBaru
			fmt.Printf("Data ID %d berhasil diubah menjadi: %s\n", id, namaBaru)
			return
		}
	}
	fmt.Println("Warga tidak ditemukan.")
}

func hapusWarga(tab *[NMAX]Warga, n *int, id int) {
	var iHapus int = -1
	var i int
	for i = 0; i < *n && indexHapus == -1; i++ {
		if tab[i].id == id {
			indexHapus = i 
	}
}