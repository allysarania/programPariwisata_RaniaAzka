package main
import "fmt"

type Fasilitas struct{
	Nama string
	Tersedia bool
}

type TempatWisata struct{
	ID int
	Nama string
	Kategori string
	Jarak float64
	Biaya float64
	Fasilitas [5]Fasilitas /*ini artinya arr fasilitas yg jumlah arrnya 5,
	isi nya sama kaya fasilitas, misal nama: toilet, tersedia: true*/
	JmlFasilitas int
}
// ini knp pake var karena dia tu variabel yang membuat data, sedangkan type itu mendefinisikan struktur data baru
var daftarWisata [100]TempatWisata
var jumlahWisata int = 0 //ini nol karena awalnya blm ada data
/* jadi daftarWisata kan max cuma 100 tuh jadi, jumlahWisata tuh untuk kita tau slot mana aja yg udh keisi. kalau dh diinput data jumlahWisata jadi keisi misal 2*/

func bacaInt() int{
	var s string
	fmt.Scan(&s)
	//ambil hanya bagian digit (dan tanda minus di depan), buang sisanya (misal "23km" -> "23")
	hasil := ""
	i := 0
	for i < len(s) {
		c := s[i]
		if c >= '0' && c <= '9' {
			hasil += string(c)
		} else if c == '-' && i == 0 {
			hasil += string(c)
		} else {
			i = len(s) //berhenti begitu nemu karakter non-digit
		}
		i++
	}
	if hasil == "" || hasil == "-" {
		return 0
	}
	n := 0
	neg := false
	j := 0
	if hasil[0] == '-' {
		neg = true
		j = 1
	}
	for j < len(hasil) {
		n = n*10 + int(hasil[j]-'0')
		j++
	}
	if neg {
		n = -n
	}
	return n
	//ini tuh helper func jadi kalau mw input yg type int tinggal panggil ini
}


func bacaString() string{
	var s string
	fmt.Scan(&s)
	return s
	//ini func helper jg 
}

func bacaFloat() float64{
	var s string
	fmt.Scan(&s)
	//ambil hanya bagian angka & titik desimal, buang sisanya (misal "23km" -> "23", "5,5" -> "5")
	hasil := ""
	i := 0
	titikUdah := false
	for i < len(s) {
		c := s[i]
		if c >= '0' && c <= '9' {
			hasil += string(c)
		} else if c == '.' && !titikUdah {
			hasil += string(c)
			titikUdah = true
		} else if c == '-' && i == 0 {
			hasil += string(c)
		} else {
			i = len(s) //berhenti begitu nemu karakter selain angka/titik
		}
		i++
	}
	var f float64
	fmt.Sscan(hasil, &f)
	return f
	//helper jg
}

func sequentialSearchById(arr [100]TempatWisata, n int,id int) int{
	//search by id 
	idx := -1
	i := 0
	
	for i < n && idx == -1{
		if arr[i].ID == id{ //ID tu id yg mau dicari, kalau bener idnya idx diudah jd nilai i
			idx = i
		}
		i++
	}
	return idx
}

func sequentialSearchByNama(arr [100]TempatWisata, n int, keyword string) int{
	/*pencarian berdasarkan kata kunci, return -1 jika tidak match*/
	idx := -1
	i := 0
	
	for i < n && idx == -1{
		namaMatch := false
		j := 0
		
		for j <= len(arr[i].Nama) - len(keyword) && !namaMatch{
			/*len tu buat ngitung berapa banyak jum lah huruf nya, misal len("pantai") = 6, itu ada !namaMatch artinya bakal berenti kalau dah ketemu*/
			match := true 
			k := 0
			
			for k < len(keyword) && match{
				ca := arr[i].Nama[j+k]
				cb := keyword[k]
				/*itu yg for sampai cb maksudnya untuk
				membandingkan karakter satu persatu antara nama(mulai dari posisi j) 
				dengan keyword j+k = posisi di nama, k posisi utnuk di keyword*/
				// case-insensitive (mengubah huruf kapital menjadi huruf kecil)
				if ca >= 'A' && ca <= 'Z' {
					ca += 32
					/*ini 32 karena selisih ascii antara 'A' dan(65) dan 'a'(97) adalah 3*/
				}
				if cb >= 'A' && cb <= 'Z' {
					cb += 32
				}
				if ca != cb {
					match = false
				}
				k++
			}
			if match {
				namaMatch = true
				//kalau semua karakter cocok maka namaMatch berubah jadi truee 
			}
			j++
		}
		if namaMatch {
			idx = i
			//nah kalau match idx  -1 berubah jadi i
		}
		i++
	}
	return idx
	
	/*contoh cara kerjanya
	Nama    : "Budi Santoso"
	keyword : "sAnt"

	Posisi j=5 → "Sant" vs "sAnt"
	Setelah lowercase → "sant" == "sant" jadi, namaMatch = true*/
}

func binarySearchById(arr [100]TempatWisata, n int, id int) int{
	left := 0
	right := n - 1
	idx := -1
	
	for left <= right && idx == -1 {
		mid := (left + right) / 2
		if arr[mid].ID == id {
			idx = mid
		} else if arr[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

//bagian sorting

func selectionSortJarak(arr *[100]TempatWisata, n int, asc bool){
	i := 0
	for i < n-1{
		idx := i
		j := i + 1
		for j < n{
			pilih := false
			if asc && arr[j].Jarak < arr[idx].Jarak{
				pilih = true
			}else if !asc && arr[j].Jarak > arr[idx].Jarak{
				pilih = true
			}
			if pilih{
				idx = j
			}
			j++
		}
		if idx != i{
			arr[i], arr[idx] = arr[idx], arr[i]
		}
		i++
	}
}

func selectionSortBiaya(arr *[100]TempatWisata, n int, asc bool){
	i := 0
	for i < n-1{
		idx := i
		j := i + 1
		for j < n{
			pilih := false
			if asc && arr[j].Biaya < arr[idx].Biaya{
				pilih = true
			}else if !asc && arr[j].Biaya > arr[idx].Biaya{
				pilih = true
			}
			if pilih{
				idx = j
			}
			j++
		}
		if idx != i{
			arr[i], arr[idx] = arr[idx], arr[i]
		}
		i++
	}
}

func insertionSortNama(arr *[100]TempatWisata, n int, asc bool) {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		for j >= 0 {
			lebih := false
			if asc && arr[j].Nama > key.Nama {
				lebih = true
			} else if !asc && arr[j].Nama < key.Nama {
				lebih = true
			}
			if lebih {
				arr[j+1] = arr[j]
				j--
			} else {
				j = -1 
			}
		}
		arr[j+1] = key
		i++
	}
}

func insertionSortKategori(arr *[100]TempatWisata, n int, asc bool) {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		for j >= 0 {
			lebih := false
			if asc && arr[j].Kategori > key.Kategori {
				lebih = true
			} else if !asc && arr[j].Kategori < key.Kategori {
				lebih = true
			}
			if lebih {
				arr[j+1] = arr[j]
				j--
			} else {
				j = -1
			}
		}
		arr[j+1] = key
		i++
	}
}

//menampilkan data
func tampilkanDetail(tampil TempatWisata) {
	fmt.Printf("  ID        : %d\n", tampil.ID)
	fmt.Printf("  Nama      : %s\n", tampil.Nama)
	fmt.Printf("  Kategori  : %s\n", tampil.Kategori)
	fmt.Printf("  Jarak     : %.1f km\n", tampil.Jarak)
	fmt.Printf("  Biaya     : Rp %.0f.000\n", tampil.Biaya)
	fmt.Printf("  Fasilitas :\n")
	if tampil.JmlFasilitas == 0 {
		fmt.Println("    (tidak ada fasilitas)")
	} else {
		i := 0
		for i < tampil.JmlFasilitas {
			status := "[ada]"
			if !tampil.Fasilitas[i].Tersedia {
				status = "[tidak ada]"
			}
			fmt.Printf("    [%s] %s\n", status, tampil.Fasilitas[i].Nama)
			i++
		}
	}
}

func tampilkanSemua(arr [100]TempatWisata, n int) {
	if n == 0 {
		fmt.Println("  (Belum ada data tempat wisata)")
		return
	}
	i := 0
	for i < n {
		fmt.Printf("\n[%d]\n", i+1) //ini yang nge-print [4], [5] diatas datanya. 
		tampilkanDetail(arr[i])
		i++
	}
}



// menu admin
func generateID() int {
	maxID := 0
	i := 0
	for i < jumlahWisata {
		if daftarWisata[i].ID > maxID {
			maxID = daftarWisata[i].ID
		}
		i++
	}
	return maxID + 1
}

func tambahWisata() {
	if jumlahWisata >= 100 {
		fmt.Println("  Data penuh! Maksimum 100 tempat wisata.")
		return
	}
	var bb TempatWisata
	bb.ID = generateID()

	fmt.Println("  (Catatan: gunakan underscore '_' untuk spasi, contoh: Kawah_Putih)")
	fmt.Print("  Nama tempat wisata        : ")
	bb.Nama = bacaString()

	fmt.Print("  Kategori (Alam/Budaya/Kuliner/Hiburan) : ")
	bb.Kategori = bacaString()

	fmt.Print("  Jarak (km, contoh: 5.5)   : ")
	bb.Jarak = bacaFloat()

	fmt.Print("  Biaya masuk (ribu Rp, contoh: 30) : ")
	bb.Biaya = bacaFloat()

	fmt.Print("  Jumlah fasilitas (maks 5) : ")
	bb.JmlFasilitas = bacaInt()

	if bb.JmlFasilitas > 5 {
		bb.JmlFasilitas = 5
	}
	if bb.JmlFasilitas < 0 {
		bb.JmlFasilitas = 0
	}

	i := 0
	for i < bb.JmlFasilitas {
		fmt.Printf("  Nama fasilitas ke-%d (tanpa spasi) : ", i+1)
		bb.Fasilitas[i].Nama = bacaString()

		fmt.Print("  Tersedia? (1=Ya/0=Tidak)  : ")
		pilih := bacaInt()
		bb.Fasilitas[i].Tersedia = (pilih == 1)
		i++
	}

	daftarWisata[jumlahWisata] = bb
	jumlahWisata++
	fmt.Printf("\n Tempat wisata '%s' berhasil ditambahkan (ID: %d).\n", bb.Nama, bb.ID)
	fmt.Println("=================")
	fmt.Println("  Data yang baru ditambahkan:")
	tampilkanDetail(bb)
	fmt.Println("=================")
}

func ubahWisata(){
	fmt.Print(" massukan id wisata yang ingin diubah: ")
	id := bacaInt()
	//cari data berdasarkan id menggunakan sequential search
	idx := sequentialSearchById(daftarWisata, jumlahWisata, id)
	if idx == -1 {
		fmt.Println("Id tidak ditemukan !")
		return
	}
		fmt.Println("  Data lama:")
		tampilkanDetail(daftarWisata[idx])
		fmt.Println("  Masukkan data baru:")
		fmt.Println("  (Catatan: gunakan underscore '_' untuk spasi, contoh: Kawah_Putih)")
		fmt.Printf("  Nama baru             : ")
		daftarWisata[idx].Nama = bacaString()
		fmt.Printf("  Kategori baru         : ")
		daftarWisata[idx].Kategori = bacaString()
		fmt.Printf("  Jarak baru (km, angka saja)       : ")
		daftarWisata[idx].Jarak = bacaFloat()
		fmt.Printf("  Biaya baru (ribu Rp, angka saja)  : ")
		daftarWisata[idx].Biaya = bacaFloat()
		fmt.Printf("  Jumlah fasilitas baru (maks 5): ")
		daftarWisata[idx].JmlFasilitas = bacaInt()
		if daftarWisata[idx].JmlFasilitas > 5 {
			daftarWisata[idx].JmlFasilitas = 5
		}
		if daftarWisata[idx].JmlFasilitas < 0 {
			daftarWisata[idx].JmlFasilitas = 0
		}
		k := 0
		for k < daftarWisata[idx].JmlFasilitas {
			fmt.Printf("  Nama fasilitas ke-%d   : ", k+1)
			daftarWisata[idx].Fasilitas[k].Nama = bacaString()
			fmt.Printf("  Tersedia? (1=Ya/0=Tidak): ")
			pilih := bacaInt()
			daftarWisata[idx].Fasilitas[k].Tersedia = (pilih == 1)
			k++
		}
		fmt.Println(" Data berhasil diubah.")
		fmt.Println("=================")
		fmt.Println("  Data terbaru:")
		tampilkanDetail(daftarWisata[idx])
		fmt.Println("=================")
}

func hapusWisata() {
	fmt.Printf("  Masukkan ID wisata yang ingin dihapus: ")
	id := bacaInt()
	idx := sequentialSearchById(daftarWisata, jumlahWisata, id)
	if idx == -1 {
		fmt.Println(" ID tidak ditemukan.")
		return
	}
	namaHapus := daftarWisata[idx].Nama
	// Geser elemen ke kiri
	i := idx
	for i < jumlahWisata-1 {
		daftarWisata[i] = daftarWisata[i+1]
		i++
	}
	jumlahWisata--
	fmt.Printf(" Tempat wisata '%s' berhasil dihapus.\n", namaHapus)
}

//menu user
func lihatDaftarWisata(){
	if jumlahWisata == 0{
		fmt.Println("BELUM ADA TEMPAT WISATA")
		return
	}
	
	var temp [100]TempatWisata
	i := 0
	for i < jumlahWisata{
		temp[i] = daftarWisata[i]
		i++
	}
		fmt.Println("  Urutkan berdasarkan:")
		fmt.Println("  1. Jarak")
		fmt.Println("  2. Biaya")
		fmt.Println("  3. Nama")
		fmt.Println("  4. Kategori")
		fmt.Printf("  Pilihan: ")
		kategori := bacaInt()
		fmt.Println("  Urutan:")
		fmt.Println("  1. Ascending (A-Z / Kecil ke Besar)")
		fmt.Println("  2. Descending (Z-A / Besar ke Kecil)")
		fmt.Printf("  Pilihan: ")
		urutan := bacaInt()
		asc := (urutan == 1)
		
		if kategori == 1 {
		selectionSortJarak(&temp, jumlahWisata, asc)
	} else if kategori == 2 {
		selectionSortBiaya(&temp, jumlahWisata, asc)
	} else if kategori == 3 {
		insertionSortNama(&temp, jumlahWisata, asc)
	} else if kategori == 4 {
		insertionSortKategori(&temp, jumlahWisata, asc)
	} else {
		fmt.Println("  [!] Pilihan tidak valid.")
		return
	}
	fmt.Println("=================")
	tampilkanSemua(temp, jumlahWisata)
}

func cariWisata(){
	fmt.Print("Masukkan kata kunci nama (atau '-' untuk skip): ")
	keyword := bacaString()
	fmt.Print("Masukkan kategori filter(atau '-' untuk skip): ")
	filterKategori := bacaString()
	
	fmt.Println()
	fmt.Println("========================")
	fmt.Println("Hasil Pencarian:")
	fmt.Println("========================")
	
	found := false 
	i := 0
	for i < jumlahWisata{
		tw := daftarWisata[i]
		//cek keyword nama
		namaOk := (keyword == "-")
		if !namaOk{
			namaMatch := false
			j := 0
			for j <= len(tw.Nama)-len(keyword) && !namaMatch {
				match := true
				k := 0
				for k < len(keyword) && match {
					ca := tw.Nama[j+k]
					cb := keyword[k]
					if ca >= 'A' && ca <= 'Z' {
						ca += 32
					}
					if cb >= 'A' && cb <= 'Z' {
						cb += 32
					}
					if ca != cb {
						match = false
					}
					k++
				}
				if match {
					namaMatch = true
				}
				j++
			}
			namaOk = namaMatch
		}
		// Cek kategori
		kategoriOk := (filterKategori == "-")
		if !kategoriOk {
			if tw.Kategori == filterKategori {
				kategoriOk = true
			}
		}
		if namaOk && kategoriOk {
			fmt.Printf("\n[Ditemukan]\n")
			tampilkanDetail(tw)
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Tidak ada tempat wisata yang sesuai.")
	}
}
	
func menuAdmin() {
	selesai := false
	for !selesai {
		fmt.Println("========================")
		fmt.Println("MENU ADMIN")
		fmt.Println("=================")
		fmt.Println("  1. Lihat semua data wisata")
		fmt.Println("  2. Tambah tempat wisata")
		fmt.Println("  3. Ubah data wisata")
		fmt.Println("  4. Hapus data wisata")
		fmt.Println("  0. Kembali")
		fmt.Println("=================")
		fmt.Printf("Pilihan: ")
		pilih := bacaInt()
		fmt.Println("========================")
		if pilih == 1 {
			tampilkanSemua(daftarWisata, jumlahWisata)
		} else if pilih == 2 {
			tambahWisata()
		} else if pilih == 3 {
			ubahWisata()
		} else if pilih == 4 {
			hapusWisata()
		} else if pilih == 0 {
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPengguna() {
	selesai := false
	for !selesai {
		fmt.Println("========================")
		fmt.Println("MENU PENGGUNA")
		fmt.Println("=================")
		fmt.Println("  1. Lihat daftar tempat wisata (terurut)")
		fmt.Println("  2. Cari tempat wisata")
		fmt.Println("  0. Kembali")
		fmt.Println("=================")
		fmt.Printf("  Pilihan: ")
		pilih := bacaInt()
		fmt.Println("========================")
		if pilih == 1 {
			lihatDaftarWisata()
		} else if pilih == 2 {
			cariWisata()
		} else if pilih == 0 {
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//inisialisasi data
func inisialisasiData() {
	// Data 1
	daftarWisata[0] = TempatWisata{
		ID: 1, Nama: "Kawah_Putih", Kategori: "Alam",
		Jarak: 46.5, Biaya: 30,
		Fasilitas: [5]Fasilitas{
			{"Parkir", true}, {"Toilet", true}, {"Warung", true}, {"Camping", false}, {"", false},
		},
		JmlFasilitas: 4,
	}
	// Data 2
	daftarWisata[1] = TempatWisata{
		ID: 2, Nama: "Tangkuban_Perahu", Kategori: "Alam",
		Jarak: 30.0, Biaya: 25,
		Fasilitas: [5]Fasilitas{
			{"Parkir", true}, {"Toilet", true}, {"Souvenir", true}, {"", false}, {"", false},
		},
		JmlFasilitas: 3,
	}
	// Data 3
	daftarWisata[2] = TempatWisata{
		ID: 3, Nama: "Saung_Angklung_Udjo", Kategori: "Budaya",
		Jarak: 8.0, Biaya: 60,
		Fasilitas: [5]Fasilitas{
			{"Parkir", true}, {"Toilet", true}, {"Toko", true}, {"Restoran", true}, {"", false},
		},
		JmlFasilitas: 4,
	}
	// Data 4
	daftarWisata[3] = TempatWisata{
		ID: 4, Nama: "Trans_Studio_Bandung", Kategori: "Hiburan",
		Jarak: 5.5, Biaya: 150,
		Fasilitas: [5]Fasilitas{
			{"Parkir", true}, {"Toilet", true}, {"FoodCourt", true}, {"ATM", true}, {"Mushola", true},
		},
		JmlFasilitas: 5,
	}
	// Data 5
	daftarWisata[4] = TempatWisata{
		ID: 5, Nama: "Alun-Alun_Bandung", Kategori: "Hiburan",
		Jarak: 2.0, Biaya: 0,
		Fasilitas: [5]Fasilitas{
			{"Toilet", true}, {"Warung", true}, {"", false}, {"", false}, {"", false},
		},
		JmlFasilitas: 2,
	}
	jumlahWisata = 5
}

func main() {
	inisialisasiData()
	selesai := false
	for !selesai {
		fmt.Println("SELAMAT DATANG!")
		fmt.Println("Login sebagai:")
		fmt.Println("  1. Admin")
		fmt.Println("  2. Pengguna")
		fmt.Println("  0. Keluar")
		fmt.Println("========================")
		fmt.Printf("Pilihan: ")
		pilih := bacaInt()
		fmt.Println("========================")
		if pilih == 1 {
			menuAdmin()
		} else if pilih == 2 {
			menuPengguna()
		} else if pilih == 0 {
			selesai = true
			fmt.Println("Terima kasih! Sampai jumpa.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}