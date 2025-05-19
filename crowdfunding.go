package main

import (
	"fmt"
	"math"
)

const maxDonasiDonatur int = 10
const maxPengguna int = 100
const maxKampanye int = 100
const maxDonasi int = 1000

type Pengguna struct {
	Email    string
	Nama     string
	Password string
	Peran    string
}

type Kampanye struct {
	Id        int
	Judul     string
	Kategori  string
	Deskripsi string
	Target    int
	Terkumpul int
	Status    string
}

type Donasi struct {
	KampanyeId  int
	NamaDonatur string
	Jumlah      int
}

type tabPengguna [maxPengguna]Pengguna
type tabKampanye [maxKampanye]Kampanye
type tabDonasi [maxDonasi]Donasi
type tabDonasiDonatur [maxDonasiDonatur]int

func daftar(daftarPengguna *tabPengguna, jumlahPengguna *int, penggunaMasuk *Pengguna) {
	var email, nama, password, peran string
	var allow bool
	var i int

	for allow == false {
		allow = true
		if *jumlahPengguna >= maxPengguna {
			fmt.Println("Jumlah pengguna sudah mencapai batas maksimum!")
			allow = false
		} else {
			fmt.Println("\n=== DAFTAR PENGGUNA BARU ===")
			fmt.Print("Email: ")
			fmt.Scan(&email)
			if !verikasiEmail(email) {
				fmt.Println("Format email tidak valid!")
				allow = false
			} else {
				for i = 0; i < *jumlahPengguna; i++ {
					if daftarPengguna[i].Email == email {
						fmt.Println("Email sudah terdaftar!")
						allow = false
					}
				}
			}
			if allow {
				fmt.Print("Nama: ")
				fmt.Scan(&nama)
			}
			if allow {
				fmt.Print("Password: ")
				fmt.Scan(&password)
				if !verikasiPassword(password) {
					fmt.Println("Password tidak valid! Harus mengandung huruf besar, huruf kecil, angka, dan karakter khusus (@#$%&).")
					allow = false
				}
			}
			if allow {
				fmt.Print("Peran (admin/donatur): ")
				fmt.Scan(&peran)
				if peran != "admin" && peran != "donatur" {
					fmt.Println("Peran tidak valid!")
					allow = false
				}
			}

		}
		if allow {
			daftarPengguna[*jumlahPengguna].Email = email
			daftarPengguna[*jumlahPengguna].Nama = nama
			daftarPengguna[*jumlahPengguna].Password = password
			daftarPengguna[*jumlahPengguna].Peran = peran
			*jumlahPengguna++

			penggunaMasuk.Email = email
			penggunaMasuk.Nama = nama
			penggunaMasuk.Password = password
			penggunaMasuk.Peran = peran
			fmt.Println("Pendaftaran berhasil!")
		}
	}
}

func masuk(daftarPengguna *tabPengguna, penggunaMasuk *Pengguna, jumlahPengguna int) {
	var user, password string
	var i int

	fmt.Print("Email/Username: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i = 0; i < jumlahPengguna; i++ {
		if (daftarPengguna[i].Email == user || daftarPengguna[i].Nama == user) && (daftarPengguna[i].Password == password) {
			*penggunaMasuk = daftarPengguna[i]
			fmt.Printf("Selamat datang, %s!\n", penggunaMasuk.Nama)
		} else if i == jumlahPengguna-1 {
			fmt.Println("Email atau password salah!")
		}
	}
}

func verikasiPassword(password string) bool {
	var i int
	var upper, lower, number, special bool

	if len(password) < 8 {
		return false
	}

	for i = 0; i < len(password); i++ {
		if password[i] >= 'A' && password[i] <= 'Z' {
			upper = true
		} else if password[i] >= 'a' && password[i] <= 'z' {
			lower = true
		} else if password[i] >= '0' && password[i] <= '9' {
			number = true
		} else if password[i] == '@' || password[i] == '#' || password[i] == '$' || password[i] == '%' || password[i] == '&' {
			special = true
		}
	}

	return upper && lower && number && special
}

func verikasiEmail(email string) bool {
	var atSymbol bool
	var dotSymbol bool
	var i int

	for i = 0; i < len(email); i++ {
		if email[i] == '@' {
			atSymbol = true
		} else if email[i] == '.' {
			dotSymbol = true
		}
	}
	return atSymbol && dotSymbol
}

func buatKampanye(daftarKampanye *tabKampanye, jumlahKampanye *int, penggunaMasuk *Pengguna) {
	var judul, deskripsi, kategori string
	var target int

	if penggunaMasuk.Peran != "admin" {
		fmt.Println("Hanya admin yang dapat membuat kampanye!")
	} else if *jumlahKampanye >= maxKampanye {
		fmt.Println("Jumlah kampanye sudah mencapai batas maksimum!")
	} else {
		fmt.Println("\n=== BUAT KAMPANYE BARU ===")
		fmt.Println("Judul: ")
		fmt.Scan(&judul)
		fmt.Println("Kategori: ")
		fmt.Scan(&kategori)
		fmt.Println("Deskripsi: ")
		fmt.Scan(&deskripsi)
		fmt.Println("Target Dana: ")
		fmt.Scan(&target)

		daftarKampanye[*jumlahKampanye].Deskripsi = deskripsi
		daftarKampanye[*jumlahKampanye].Judul = judul
		daftarKampanye[*jumlahKampanye].Target = target
		daftarKampanye[*jumlahKampanye].Kategori = kategori
		daftarKampanye[*jumlahKampanye].Status = "aktif"

		*jumlahKampanye++
		fmt.Println("Kampanye berhasil dibuat!")
	}
}

func findJudulKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanJudul string) int {
	var i, k, j, jumKata, nearJudul, bestIdx int
	var judulKampanye string
	nearJudul = 0
	bestIdx = -1
	for i = 0; i < jumlahKampanye; i++ {
		judulKampanye = daftarKampanye[i].Judul
		jumKata = 0
		for j = 0; j < len(judulKampanye); j++ {
			for k = 0; k < len(pilihanJudul); k++ {
				if judulKampanye[j] == pilihanJudul[k] {
					jumKata++
				}
			}
		}
		if jumKata > nearJudul {
			nearJudul = jumKata
			bestIdx = i
		}
	}
	return bestIdx
}

func findKategoriKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanKategori string) int {
	var idx int
	var kategoriKampanye string
	for idx = 0; idx < jumlahKampanye; idx++ {
		kategoriKampanye = daftarKampanye[idx].Kategori
		if kategoriKampanye == pilihanKategori {
			return idx
		}
	}
	return -1
}

func findIdKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanId int) int {
	var left, right, mid, idKampanye int

	left = 0
	right = jumlahKampanye - 1

	for left <= right {
		mid = (left + right) / 2

		if idKampanye == pilihanId {
			return mid
		} else if idKampanye < pilihanId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func tampilkanDetailKampanye(kampanye Kampanye) {
	var progress int
	progress = (kampanye.Terkumpul * 100) / kampanye.Target

	fmt.Println("\n=== DETAIL KAMPANYE ===")
	fmt.Printf("Judul: %s\n", kampanye.Judul)
	fmt.Printf("Kategori: %s\n", kampanye.Kategori)
	fmt.Printf("Deskripsi: %s\n", kampanye.Deskripsi)
	fmt.Printf("Target Dana: Rp %d\n", kampanye.Target)
	fmt.Printf("Terkumpul: Rp %d\n", kampanye.Terkumpul)
	fmt.Printf("Status: %s\n", kampanye.Status)
	fmt.Printf("Progress: %d%% [%s]\n", progress, createProgressBar(progress))
}

func tampilkanKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var i, progress, ascending, pilihan, pilihanId, pilihanIdx int
	var k Kampanye
	var tempKampanye tabKampanye
	var pilihanJudul, pilihanKategori string

	tempKampanye = *daftarKampanye

	fmt.Print("Urutkan berdasarkan progress (1: ascending, 0: descending): ")
	fmt.Scan(&ascending)

	sortSelection(&tempKampanye, jumlahKampanye, ascending)

	fmt.Println("=              DAFTAR KAMPANYE AKTIF                 =")

	if jumlahKampanye == 0 {
		fmt.Println("=           Belum ada kampanye yang dibuat            =")
	} else {
		fmt.Println("=           Jumlah kampanye yang dibuat: ", jumlahKampanye, "          =")
		for i = 0; i < jumlahKampanye; i++ {
			k = tempKampanye[i]
			progress = (k.Terkumpul * 100) / k.Target

			fmt.Printf("[%d] %s\n", i, k.Judul)
			fmt.Printf("Status: %s\n", k.Status)
			fmt.Printf("Progress: %v%% [%s]\n", progress, createProgressBar(progress))
			fmt.Printf("Terkumpul: Rp %v dari Rp %v\n", k.Terkumpul, k.Target)
			fmt.Println()
		}
		fmt.Println("=           Pilih kampanye untuk melihat detail       =")
		fmt.Println("1. ID Kampanye")
		fmt.Println("2. Judul Kampanye")
		fmt.Println("3. Kategori Kampanye")
		fmt.Println("Kembali ke menu utama (0)")
		fmt.Println("Pilih Kampanye untuk melihat detail (ketik Angka): ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 0:
			fmt.Println("Kembali ke menu utama")
		case 1:
			fmt.Println("Ketik ID Kampanye: ")
			fmt.Scan(&pilihanId)
			pilihanIdx = findIdKampanye(tempKampanye, jumlahKampanye, pilihanId)
		case 2:
			fmt.Println("Ketik Judul Kampanye: ")
			fmt.Scan(&pilihanJudul)
			pilihanIdx = findJudulKampanye(tempKampanye, jumlahKampanye, pilihanJudul)
		case 3:
			fmt.Println("Ketik Kategori Kampanye: ")
			fmt.Scan(&pilihanKategori)
			pilihanIdx = findKategoriKampanye(tempKampanye, jumlahKampanye, pilihanKategori)
		case -1:
			fmt.Println("Pilihan tidak ditemukan!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		if pilihan >= 1 && pilihan <= 3 {
			tampilkanDetailKampanye(tempKampanye[pilihanIdx])
		}
	}
}

func createProgressBar(percent int) string {
	var width, filled int
	var bar string

	width = 20
	filled = int(math.Round(float64(percent) / 100 * float64(width)))
	if filled > width {
		filled = width
	}
	bar = ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func tambahDonasi(daftarKampanye *tabKampanye, daftarDonasi *tabDonasi, penggunaMasuk *Pengguna, jumlahKampanye int, jumlahDonasi *int) {
	var kampanyeId, totalDonasi, jumlah int

	if penggunaMasuk.Peran == "" {
		fmt.Println("Silakan masuk terlebih dahulu!")
	} else if *jumlahDonasi >= maxDonasi {
		fmt.Println("Jumlah donasi sudah mencapai batas maksimum!")
		return
	} else if jumlahKampanye == 0 {
		fmt.Println("Belum ada kampanye yang dibuat!")
		return
	} else {
		fmt.Println("\n=== DONASI KAMPANYE ===")

		tampilkanKampanye(daftarKampanye, jumlahKampanye)
		fmt.Print("ID Kampanye: ")
		fmt.Scan(&kampanyeId)

		if kampanyeId >= jumlahKampanye || daftarKampanye[kampanyeId].Status != "aktif" || kampanyeId < 0 {
			fmt.Println("Kampanye tidak valid atau tidak aktif!")
		} else {
			fmt.Print("Jumlah Donasi: ")
			fmt.Scan(&jumlah)

			daftarDonasi[*jumlahDonasi].KampanyeId = kampanyeId
			daftarDonasi[*jumlahDonasi].NamaDonatur = penggunaMasuk.Nama
			daftarDonasi[*jumlahDonasi].Jumlah = jumlah
			totalDonasi = daftarKampanye[kampanyeId].Terkumpul + jumlah

			if totalDonasi > daftarKampanye[kampanyeId].Target {
				fmt.Println("Jumlah donasi melebihi target kampanye!")
				totalDonasi -= daftarKampanye[kampanyeId].Target
				fmt.Println("Donasi dikembalikan sebesar Rp", totalDonasi)

			}
			daftarKampanye[kampanyeId].Terkumpul = totalDonasi
			*jumlahDonasi = *jumlahDonasi + 1
			fmt.Println("Donasi berhasil! Terima kasih.")
		}
	}

}

func logOut(penggunaMasuk *Pengguna) {
	penggunaMasuk.Peran = ""
	penggunaMasuk.Nama = ""
	penggunaMasuk.Email = ""
	penggunaMasuk.Password = ""
	fmt.Println("Anda telah keluar.")
}

func insertionSort(daftarDonasi *tabDonasi, jumlahDonasi int, ascending bool) {
	var i, j int
	var k Donasi

	for i = 1; i < jumlahDonasi; i++ {
		k = daftarDonasi[i]
		j = i - 1

		if ascending {
			for j >= 0 && daftarDonasi[j].Jumlah > k.Jumlah {
				daftarDonasi[j+1] = daftarDonasi[j]
				j--
			}
		} else {
			for j >= 0 && daftarDonasi[j].Jumlah < k.Jumlah {
				daftarDonasi[j+1] = daftarDonasi[j]
				j--
			}
		}
		daftarDonasi[j+1] = k
	}
}

func sortSelection(daftarKampanye *tabKampanye, jumlahKampanye int, ascending int) {
	var i, j, minIdx, progressA, progressB int
	var k Kampanye

	for i = 0; i < jumlahKampanye-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlahKampanye; j++ {
			progressA = (daftarKampanye[j].Terkumpul * 100) / daftarKampanye[j].Target
			progressB = (daftarKampanye[minIdx].Terkumpul * 100) / daftarKampanye[minIdx].Target
			if ascending == 1 {
				if progressA < progressB {
					minIdx = j
				}
			} else {
				if progressA > progressB {
					minIdx = j
				}
			}
		}

		if minIdx != i {
			k = daftarKampanye[i]
			daftarKampanye[i] = daftarKampanye[minIdx]
			daftarKampanye[minIdx] = k
		}
	}
}

func totalDonasi(daftarDonasi tabDonasi, jumlahDonasi int) int {
	var i int
	var total int

	for i = 0; i < jumlahDonasi; i++ {
		total += daftarDonasi[i].Jumlah
	}
	return total
}

func findNamaDonasi(daftarDonasi tabDonasi, jumlahDonasi int, pilihanNama string) tabDonasiDonatur {
	var i int
	var donasi Donasi
	var arrDonasi tabDonasiDonatur

	for i = 0; i < jumlahDonasi; i++ {
		donasi = daftarDonasi[i]
		if donasi.NamaDonatur == pilihanNama {
			arrDonasi[i] = donasi.Jumlah
		}
	}
	return arrDonasi
}

func totalDonasiDonatur(daftarDonasi tabDonasi, jumlahDonasi int, pilihanNama string) int {
	var i int
	var donasi Donasi
	var total int

	for i = 0; i < jumlahDonasi; i++ {
		donasi = daftarDonasi[i]
		if donasi.NamaDonatur == pilihanNama {
			total += donasi.Jumlah
		}
	}
	return total
}

func tampilKampanyeDonatur(daftarKampanye tabKampanye, jumlahKampanye int, arrKampanye tabDonasiDonatur) {
	var i int
	var kampanye Kampanye

	for i = 0; i < len(arrKampanye); i++ {
		kampanye = daftarKampanye[i]
		fmt.Printf("ID Kampanye: %d\n", kampanye.Id)
		fmt.Printf("Judul: %s\n", kampanye.Judul)
		fmt.Printf("Kategori: %s\n", kampanye.Kategori)
		fmt.Printf("Deskripsi: %s\n", kampanye.Deskripsi)
		fmt.Printf("Target Dana: Rp %d\n", kampanye.Target)
		fmt.Printf("Terkumpul: Rp %d\n", kampanye.Terkumpul)
		fmt.Println()
	}
}

func tampilkanDonasi(daftarDonasi *tabDonasi, jumlahDonasi int, daftarKampanye tabKampanye, jumlahKampanye int) {
	var i int
	var k Donasi
	var tempDaftarDonasi tabDonasi
	var pilihanNama string
	var detailDonatur tabDonasiDonatur
	tempDaftarDonasi = *daftarDonasi

	if jumlahDonasi == 0 {
		fmt.Println("Belum ada donasi yang dilakukan!")
	} else {
		insertionSort(&tempDaftarDonasi, jumlahDonasi, true)
		fmt.Println("\n=== DAFTAR DONASI ===")
		for i = 0; i < jumlahDonasi; i++ {
			k = tempDaftarDonasi[i]
			fmt.Printf("Kampanye ID: %d\n", k.KampanyeId)
			fmt.Printf("Nama Donatur: %s\n", k.NamaDonatur)
			fmt.Printf("Jumlah Donasi: Rp %d\n", k.Jumlah)
			fmt.Println()
		}

		fmt.Println("Jumlah donasi: ", jumlahDonasi)
		fmt.Println("Total donasi: Rp ", totalDonasi(tempDaftarDonasi, jumlahDonasi))

		fmt.Println("Pilih donasi untuk melihat detail (ketik Nama): ")
		fmt.Scan(&pilihanNama)
		detailDonatur = findNamaDonasi(tempDaftarDonasi, jumlahDonasi, pilihanNama)
		if len(detailDonatur) > 0 {
			fmt.Printf("Detail Donasi (%v):\n", pilihanNama)
			fmt.Printf("Total Donasi: Rp %d\n", totalDonasiDonatur(tempDaftarDonasi, jumlahDonasi, pilihanNama))
			tampilKampanyeDonatur(daftarKampanye, jumlahKampanye, detailDonatur)
		} else {
			fmt.Println("ID donasi tidak ditemukan!")
		}

	}
}

func menuUtama(penggunaMasuk Pengguna) int {
	var pilihan int

	fmt.Println("\n=== SISTEM CROWDFUNDING ===")
	if penggunaMasuk.Peran != "" {
		fmt.Printf("Selamat datang, %s (%s)\n", penggunaMasuk.Nama, penggunaMasuk.Peran)
	} else {
		fmt.Println("Daftar (1)")
		fmt.Println("Masuk (2)")
	}
	fmt.Println("Lihat Kampanye (3)")
	fmt.Println("Donasi (4)")
	if penggunaMasuk.Peran == "admin" {
		fmt.Println("Buat Kampanye (5)")
	}
	if penggunaMasuk.Peran != "" {
		fmt.Println("Log Out (6)")
	}
	fmt.Println("Ketik '-1' untuk keluar")

	fmt.Print("Pilih menu (ketik angka tersebut): ")
	fmt.Scan(&pilihan)
	fmt.Println()

	return pilihan
}

func main() {
	var pilihan int
	var jumlahPengguna, jumlahKampanye, jumlahDonasi int
	var daftarPengguna tabPengguna
	var daftarKampanye tabKampanye
	var daftarDonasi tabDonasi
	var penggunaMasuk Pengguna // Pengguna yang sedang masuk sekarang

	for pilihan != -1 {
		pilihan = menuUtama(penggunaMasuk)

		switch pilihan {
		case 1:
			daftar(&daftarPengguna, &jumlahPengguna, &penggunaMasuk)
		case 2:
			masuk(&daftarPengguna, &penggunaMasuk, jumlahPengguna)
		case 3:
			tampilkanDonasi(&daftarDonasi, jumlahDonasi, daftarKampanye, jumlahKampanye)
		case 4:
			tampilkanKampanye(&daftarKampanye, jumlahKampanye)
		case 5:
			tambahDonasi(&daftarKampanye, &daftarDonasi, &penggunaMasuk, jumlahKampanye, &jumlahDonasi)
		case 6:
			if penggunaMasuk.Peran == "admin" {
				buatKampanye(&daftarKampanye, &jumlahKampanye, &penggunaMasuk)
			} else {
				fmt.Println("Hanya admin yang dapat membuat kampanye!")
			}
		case 7:
			if penggunaMasuk.Peran != "" {
				logOut(&penggunaMasuk)
			} else {
				fmt.Println("Silakan masuk terlebih dahulu!")
			}
		case -1:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println(penggunaMasuk.Peran)
	}
}
