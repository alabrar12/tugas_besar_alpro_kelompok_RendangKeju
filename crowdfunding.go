package main

import (
	"fmt"
	"math"
)

const maxPengguna int = 100
const maxKampanye int = 100
const maxDonasi int = 1000

type Pengguna struct {
	Email    string
	Username string
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
	Progress  int
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

func daftar(daftarPengguna *tabPengguna, jumlahPengguna *int, penggunaMasuk *Pengguna) {
	var email, username, password, peran string
	var allow bool

	allow = true
	for allow {
		allow = true
		if *jumlahPengguna >= maxPengguna {
			fmt.Println("Jumlah pengguna sudah mencapai batas maksimum!")
			allow = false
		} else {
			fmt.Println("\n=== DAFTAR PENGGUNA BARU (ketik 'keluar') ===")
			verikasiEmail(*daftarPengguna, *jumlahPengguna, &email, &allow)
			verikasiUsername(*daftarPengguna, *jumlahPengguna, &username, &allow)
			verikasiPassword(&password, &allow)
			verikasiPeran(&peran, &allow)
			if !allow {
				fmt.Println()
				fmt.Println("      Pendaftaran dibatalkan.")
			} else {
				daftarPengguna[*jumlahPengguna].Email = email
				daftarPengguna[*jumlahPengguna].Username = username
				daftarPengguna[*jumlahPengguna].Password = password
				daftarPengguna[*jumlahPengguna].Peran = peran
				*jumlahPengguna++

				penggunaMasuk.Email = email
				penggunaMasuk.Username = username
				penggunaMasuk.Password = password
				penggunaMasuk.Peran = peran
				fmt.Println()
				fmt.Println("      Pendaftaran berhasil!")
				allow = false
			}
		}
	}
	fmt.Println()
}

func verikasiEmail(daftarPengguna tabPengguna, jumlahPengguna int, email *string, allow *bool) {
	var i, lenEmail int
	var validDomain, emailExists, valid, validDot bool

	if *allow {
		fmt.Print("Email: ")
		fmt.Scan(&*email)
	}

	lenEmail = len(*email) - 1

	valid = false
	for *allow && !valid && (*email != "keluar") {
		if len(*email) < 12 || len(*email) > 25 {
			fmt.Println("Email harus memiliki panjang antara 12 hingga 25 karakter.")
		} else {
			validDomain = false

			validDot = string((*email)[lenEmail]) == "m" && string((*email)[lenEmail-1]) == "o" && string((*email)[lenEmail-2]) == "c" && string((*email)[lenEmail-3]) == "."

			validDomain = validDomain || string((*email)[lenEmail-4]) == "l" && string((*email)[lenEmail-5]) == "i" && string((*email)[lenEmail-6]) == "a" && string((*email)[lenEmail-7]) == "m" && string((*email)[lenEmail-8]) == "g" && string((*email)[lenEmail-9]) == "@"
			validDomain = validDomain || string((*email)[lenEmail-4]) == "o" && string((*email)[lenEmail-5]) == "o" && string((*email)[lenEmail-6]) == "h" && string((*email)[lenEmail-7]) == "a" && string((*email)[lenEmail-8]) == "y" && string((*email)[lenEmail-9]) == "@"
			validDomain = validDomain || string((*email)[lenEmail-4]) == "k" && string((*email)[lenEmail-5]) == "o" && string((*email)[lenEmail-6]) == "o" && string((*email)[lenEmail-7]) == "l" && string((*email)[lenEmail-8]) == "t" && string((*email)[lenEmail-9]) == "u" && string((*email)[lenEmail-10]) == "o" && string((*email)[lenEmail-11]) == "@"

			if !validDomain || !validDot {
				fmt.Println("Email harus mengandung '@gmail.com', '@yahoo.com' atau '@outlook.com'")
			} else {
				emailExists = false
				for i = 0; i < jumlahPengguna && !emailExists; i++ {
					if daftarPengguna[i].Email == *email {
						emailExists = true
						fmt.Println("Email sudah terdaftar!")
					}
				}

				if !emailExists {
					valid = true
				}
			}
		}

		if !valid && *allow {
			fmt.Print("Masukkan email yang valid: ")
			fmt.Scan(&*email)
		}
	}
	*allow = *allow && valid
}

func verikasiUsername(daftarPengguna tabPengguna, jumlahPengguna int, username *string, allow *bool) {
	var i int
	var valid bool

	if *allow {
		fmt.Print("Username: ")
		fmt.Scan(&*username)
	}

	valid = false
	for *allow && !valid {
		*allow = *username != "keluar"
		if *allow {
			if len(*username) >= 4 && len(*username) <= 25 {
				valid = true
				for i = 0; i < jumlahPengguna; i++ {
					if daftarPengguna[i].Username == *username {
						fmt.Println("Username sudah terdaftar!")
						valid = false
					}
				}
			} else {
				fmt.Println("Username harus memiliki panjang antara 4 hingga 25 karakter.")
				valid = false
			}

			if !valid {
				fmt.Print("Masukkan username yang valid: ")
				fmt.Scan(&*username)
			}
		}
	}
}

func verikasiPassword(pass *string, allow *bool) {
	var i int
	var upper, lower, number, special, passValid bool

	if *allow {
		fmt.Print("Password - Gunakan minimal 8 huruf dengan isi huruf besar, huruf kecil, angka, dan simbol(@,#,$,%,&) : ")
		fmt.Scan(&*pass)
	}

	for !passValid && *allow {
		*allow = *pass != "keluar"
		if *allow {
			upper = false
			lower = false
			number = false
			special = false

			if len(*pass) >= 8 {
				for i = 0; i < len(*pass); i++ {
					if (*pass)[i] >= 'A' && (*pass)[i] <= 'Z' {
						upper = true
					}
					if (*pass)[i] >= 'a' && (*pass)[i] <= 'z' {
						lower = true
					}
					if (*pass)[i] >= '0' && (*pass)[i] <= '9' {
						number = true
					}
					if (*pass)[i] == '@' || (*pass)[i] == '#' || (*pass)[i] == '$' || (*pass)[i] == '%' || (*pass)[i] == '&' {
						special = true
					}
				}
			}

			if len(*pass) < 8 || !upper || !lower || !number || !special {
				fmt.Println("Password harus minimal 8 karakter dan mengandung:")
				fmt.Println("- huruf besar")
				fmt.Println("- huruf kecil")
				fmt.Println("- angka")
				fmt.Println("- simbol (@,#,$,%,&)")
				fmt.Print("Masukkan password yang valid: ")
				fmt.Scan(&*pass)
			} else {
				passValid = true
			}
		}
	}
}

func verikasiPeran(peran *string, allow *bool) {
	var valid bool

	if *allow {
		fmt.Print("Peran (admin/donatur): ")
		fmt.Scan(&*peran)
	}

	valid = false
	for *allow && !valid {
		if *peran == "keluar" {
			*allow = false
			valid = true
		} else if *peran == "admin" || *peran == "donatur" {
			valid = true
		} else {
			fmt.Println("Peran harus 'admin' atau 'donatur'.")
			fmt.Print("Masukkan peran yang valid: ")
			fmt.Scan(&*peran)
		}
	}
}

func masuk(daftarPengguna *tabPengguna, penggunaMasuk *Pengguna, jumlahPengguna int) {
	var user, password string
	var i int
	var loginValid bool

	fmt.Print("Email/Username: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	loginValid = false
	for i = 0; i < jumlahPengguna && !loginValid; i++ {
		if (daftarPengguna[i].Email == user || daftarPengguna[i].Username == user) &&
			(daftarPengguna[i].Password == password) {
			*penggunaMasuk = daftarPengguna[i]
			fmt.Printf("Selamat datang, %s!\n", penggunaMasuk.Username)
			loginValid = true
		}
	}

	if !loginValid {
		fmt.Println()
		fmt.Println("   Email/Username atau Password salah!")
	}
	fmt.Println()
}

func tampilkanDonasi(daftarDonasi *tabDonasi, jumlahDonasi int, daftarKampanye tabKampanye, jumlahKampanye int) {
	var i, pilihan, pilihanId, pilihanMaxIdx, pilihanUrutan int
	var k Donasi
	var tempDaftarDonasi tabDonasi
	var pilihanNama string
	var detailDonatur tabDonasi

	tempDaftarDonasi = *daftarDonasi

	if jumlahDonasi == 0 {
		fmt.Println("    Belum ada donasi yang dilakukan!")
	} else {
		insertionSortDesc(&tempDaftarDonasi, jumlahDonasi)
		fmt.Println("      === CATATAN DONASI ===")

		for i = 0; i < jumlahDonasi; i++ {
			k = tempDaftarDonasi[i]
			fmt.Printf("Kampanye ID: %d\n", k.KampanyeId)
			fmt.Printf("Nama Donatur: %s\n", k.NamaDonatur)
			fmt.Printf("Jumlah Donasi: Rp %d\n", k.Jumlah)
			fmt.Println()
		}

		fmt.Println("===   TOTAL KESELURUHAN DONASI   ===")
		fmt.Println("Jumlah donasi: ", jumlahDonasi)
		fmt.Println("Total donasi: Rp ", totalDonasi(tempDaftarDonasi, jumlahDonasi))
		fmt.Println()
		fmt.Println("= Pilih donasi untuk melihat detail =")
		fmt.Println("1. Berdasarkan Nama Donatur")
		fmt.Println("2. Berdasarkan ID Kampanye")
		fmt.Println("Kembali ke menu utama (0)")

		fmt.Print("Pilih opsi (ketik angka): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("Ketik Nama Donatur: ")
			fmt.Scan(&pilihanNama)
			fmt.Print("Mau diurutkan berdasarkan (1: Ascending, 2: Descending): ")
			fmt.Scan(&pilihanUrutan)

			if pilihanUrutan == 1 {
				insertionSortAsc(&tempDaftarDonasi, jumlahDonasi)
			} else {
				insertionSortDesc(&tempDaftarDonasi, jumlahDonasi)
			}

			detailDonatur = findNamaDonasi(tempDaftarDonasi, jumlahDonasi, pilihanNama)
			if len(detailDonatur) > 0 {

				fmt.Println()
				tampilNamaDonatur(daftarKampanye, jumlahKampanye, detailDonatur)
				fmt.Println("=   Detail Donasi untuk Donatur: ", pilihanNama, "   =")
				fmt.Printf("Total Donasi: Rp %d\n", totalDonasiDonatur(tempDaftarDonasi, jumlahDonasi, pilihanNama))
			}
		}

		if pilihan == 2 {
			fmt.Print("Ketik ID Kampanye: ")
			fmt.Scan(&pilihanId)
			fmt.Print("Maximal Donasi yang ditampilkan: ")
			fmt.Scan(&pilihanMaxIdx)

			detailDonatur = findIdDonasi(tempDaftarDonasi, jumlahDonasi, pilihanId, pilihanMaxIdx)
			if len(detailDonatur) > 0 {
				fmt.Println()
				fmt.Printf("=   Detail Donasi untuk Kampanye ID %d:   =\n", pilihanId)
				tampilKampanyeDonatur(detailDonatur, jumlahDonasi, daftarKampanye, jumlahKampanye, pilihanId)
				fmt.Println("=   Detail Donasi untuk Kampanye ID: ", pilihanId, "   =")
				fmt.Printf("Total Donasi: Rp %d\n", totalDonasiDonatur(detailDonatur, jumlahDonasi, string(pilihanId)))
			}
		}
	}
	fmt.Println()
}

func buatKampanye(daftarKampanye *tabKampanye, jumlahKampanye *int, penggunaMasuk *Pengguna) {
	var judul, deskripsi, kategori string
	var target int

	if penggunaMasuk.Peran != "admin" {
		fmt.Println("Hanya admin yang dapat membuat kampanye!")
		fmt.Println()
	} else if *jumlahKampanye >= maxKampanye {
		fmt.Println("Jumlah kampanye sudah mencapai batas maksimum!")
		fmt.Println()
	} else {
		fmt.Println("\n=== BUAT KAMPANYE BARU ===")
		fmt.Print("Judul: ")
		fmt.Scan(&judul)

		fmt.Print("Kategori: ")
		fmt.Scan(&kategori)

		fmt.Println("Deskripsi: ")
		fmt.Scan(&deskripsi)

		fmt.Print("Target Dana (Rupiah): ")
		fmt.Scan(&target)

		daftarKampanye[*jumlahKampanye].Id = 100 + (*jumlahKampanye * 10)
		daftarKampanye[*jumlahKampanye].Judul = judul
		daftarKampanye[*jumlahKampanye].Kategori = kategori
		daftarKampanye[*jumlahKampanye].Deskripsi = deskripsi
		daftarKampanye[*jumlahKampanye].Target = target
		daftarKampanye[*jumlahKampanye].Status = "aktif"

		*jumlahKampanye = *jumlahKampanye + 1

		fmt.Println()
		fmt.Println("-  Kampanye berhasil dibuat!  -")
		fmt.Println()
	}
}

func findJudulKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanJudul string) int {
	var i, j int
	var judulKampanye string
	var bestMatch, bestIdx int

	bestMatch = 0
	bestIdx = -1

	for i = 0; i < jumlahKampanye; i++ {
		var matchCount int
		matchCount = 0
		judulKampanye = daftarKampanye[i].Judul
		if judulKampanye == pilihanJudul {
			return i
		} else {
			for j = 0; j < len(judulKampanye); j++ {
				if containsChar(pilihanJudul, judulKampanye[j]) {
					matchCount = matchCount + 1
				}
			}
			if matchCount > bestMatch {
				bestMatch = matchCount
				bestIdx = i
			}
		}
	}
	return bestIdx
}

func containsChar(s string, c byte) bool {
	var i int
	var result bool

	result = false
	for i = 0; i < len(s); i++ {
		if s[i] == c || s[i] == c+32 || s[i] == c-32 {
			result = true
		}
	}
	return result
}

func findIdKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanId int) int {
	// Implementasi Binary Search untuk menemukan ID kampanye
	var left, right, mid int

	left = 0
	right = jumlahKampanye - 1

	for left <= right {
		mid = (left + right) / 2

		if daftarKampanye[mid].Id == pilihanId {
			return mid
		} else if daftarKampanye[mid].Id < pilihanId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func tampilkanDetailKampanye(kampanye Kampanye) {

	fmt.Println("\n=== DETAIL KAMPANYE ===")
	fmt.Printf("Judul: %s\n", kampanye.Judul)
	fmt.Printf("Kategori: %s\n", kampanye.Kategori)
	fmt.Printf("Deskripsi: \n%s\n", kampanye.Deskripsi)
	fmt.Printf("Target Dana: Rp %d\n", kampanye.Target)
	fmt.Printf("Terkumpul: Rp %d\n", kampanye.Terkumpul)
	fmt.Printf("Status: %s\n", kampanye.Status)
	fmt.Printf("Progress: %d%% [%s]\n", kampanye.Progress, createProgressBar(kampanye.Progress))
}

func tampilkanKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var i, ascending, progress int
	var tempKampanye tabKampanye
	var k Kampanye

	tempKampanye = *daftarKampanye

	fmt.Print("Urutkan berdasarkan progress (1: ascending, 0: descending): ")
	fmt.Scan(&ascending)

	if ascending == 1 {
		sortSelectionAsc(&tempKampanye, jumlahKampanye)
	} else {
		sortSelectionDesc(&tempKampanye, jumlahKampanye)
	}

	if jumlahKampanye == 0 {
		fmt.Println("=           Belum ada kampanye yang dibuat            =")
	} else {
		fmt.Println("=           Jumlah kampanye yang dibuat: ", jumlahKampanye, "          =")
		for i = 0; i < jumlahKampanye; i++ {
			k = tempKampanye[i]
			progress = (k.Terkumpul * 100) / k.Target

			fmt.Printf("[%d] %s\n", k.Id, k.Judul)
			fmt.Printf("Status: %s\n", k.Status)
			fmt.Printf("Kategori: %s\n", k.Kategori)
			fmt.Printf("Progress: %v%% [%s]\n", progress, createProgressBar(progress))
			fmt.Println()
		}
	}
}

func detailKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var pilihan, pilihanIdx, pilihanId int
	var pilihanJudul, blank string
	var tempKampanye tabKampanye

	tempKampanye = *daftarKampanye

	if jumlahKampanye == 0 {
		fmt.Println("=           Belum ada kampanye yang dibuat            =")
	} else {
		fmt.Println("=              DAFTAR KAMPANYE AKTIF                 =")
		tampilkanKampanye(&tempKampanye, jumlahKampanye)

		fmt.Println("=           Pilih kampanye untuk melihat detail       =")
		fmt.Println("1. ID Kampanye")
		fmt.Println("2. Judul Kampanye")
		fmt.Println("Kembali ke menu utama (0)")
		fmt.Print("Pilih Kampanye untuk melihat detail (ketik Angka): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Ketik ID Kampanye: ")
			fmt.Scan(&pilihanId)
			pilihanIdx = findIdKampanye(tempKampanye, jumlahKampanye, pilihanId)
			if pilihanIdx != -1 {
				tampilkanDetailKampanye(tempKampanye[pilihanIdx])
			} else {
				fmt.Println("Kampanye dengan ID tersebut tidak ditemukan!")
			}
		} else if pilihan == 2 {
			fmt.Print("Ketik Judul Kampanye: ")
			fmt.Scan(&pilihanJudul)
			pilihanIdx = findJudulKampanye(tempKampanye, jumlahKampanye, pilihanJudul)
			if pilihanIdx != -1 {
				tampilkanDetailKampanye(tempKampanye[pilihanIdx])
			} else {
				fmt.Println("Kampanye dengan judul tersebut tidak ditemukan!")
			}
		} else if pilihan == 0 {
			fmt.Println("    Kembali ke menu utama")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}

		if pilihan == 1 || pilihan == 2 {
			fmt.Print("Tekan Enter untuk kembali ke menu utama...")
			fmt.Scan(&blank)
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

func checkKampanyeAktif(daftarKampanye *tabKampanye, jumlahKampanye int, pilihanId int) bool {
	var i int

	for i = 0; i < jumlahKampanye; i++ {
		if daftarKampanye[i].Status == "aktif" && pilihanId == daftarKampanye[i].Id {
			return true
		}
	}
	return false
}

func tambahDonasi(daftarKampanye *tabKampanye, daftarDonasi *tabDonasi, penggunaMasuk *Pengguna, jumlahKampanye int, jumlahDonasi *int) {
	var idx, i, kampanyeId, totalDonasi, jumlah int

	if penggunaMasuk.Peran == "" {
		fmt.Println("    Silakan masuk terlebih dahulu!")
	} else if *jumlahDonasi >= maxDonasi {
		fmt.Println("    Jumlah donasi sudah mencapai batas maksimum!")
	} else if jumlahKampanye == 0 {
		fmt.Println("    Belum ada kampanye yang dibuat!")
	} else {
		fmt.Println("===   DONASI KAMPANYE   ===")
		tampilkanKampanye(daftarKampanye, jumlahKampanye)

		if jumlahKampanye > 0 {
			fmt.Print("ID Kampanye: ")
			fmt.Scan(&kampanyeId)

			if checkKampanyeAktif(daftarKampanye, jumlahKampanye, kampanyeId) {
				fmt.Print("Jumlah Donasi: ")
				fmt.Scan(&jumlah)

				for i = 0; i < *jumlahDonasi; i++ {
					if daftarDonasi[i].KampanyeId == kampanyeId {
						idx = i
					}
				}

				totalDonasi = daftarKampanye[idx].Terkumpul + jumlah

				if totalDonasi > daftarKampanye[idx].Target {
					fmt.Printf("Jumlah donasi melebihi target kampanye!, Donasi dikembalikan sebesar Rp %d\n", totalDonasi-daftarKampanye[idx].Target)
					totalDonasi = daftarKampanye[idx].Target
					jumlah = daftarKampanye[idx].Target
				}

				daftarKampanye[idx].Terkumpul = totalDonasi
				daftarKampanye[idx].Progress = (totalDonasi * 100) / daftarKampanye[idx].Target

				if daftarKampanye[idx].Progress >= 100 {
					daftarKampanye[idx].Status = "selesai"
					daftarKampanye[idx].Progress = 100
					fmt.Printf("=     Kampanye telah selesai, terkumpul Rp %d dari target Rp %d     =\n", daftarKampanye[idx].Terkumpul, daftarKampanye[idx].Target)
				}

				daftarDonasi[*jumlahDonasi].KampanyeId = kampanyeId
				daftarDonasi[*jumlahDonasi].NamaDonatur = penggunaMasuk.Username
				daftarDonasi[*jumlahDonasi].Jumlah = jumlah

				*jumlahDonasi = *jumlahDonasi + 1
				fmt.Println()
				fmt.Println("Donasi berhasil! Terima kasih.")
			} else {
				fmt.Println("Kampanye tidak valid atau tidak aktif!")
			}
		}
	}
	fmt.Println()
}

func insertionSortAsc(daftarDonasi *tabDonasi, jumlahDonasi int) {
	var i, j int
	var k Donasi

	for i = 1; i < jumlahDonasi; i++ {
		k = daftarDonasi[i]
		j = i - 1

		for j >= 0 && daftarDonasi[j].Jumlah > k.Jumlah {
			daftarDonasi[j+1] = daftarDonasi[j]
			j--
		}
		daftarDonasi[j+1] = k
	}
}

func insertionSortDesc(daftarDonasi *tabDonasi, jumlahDonasi int) {
	var i, j int
	var k Donasi

	for i = 1; i < jumlahDonasi; i++ {
		k = daftarDonasi[i]
		j = i - 1

		for j >= 0 && daftarDonasi[j].Jumlah < k.Jumlah {
			daftarDonasi[j+1] = daftarDonasi[j]
			j--
		}
		daftarDonasi[j+1] = k
	}
}

func sortSelectionAsc(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var pass, j, idx int
	var temp Kampanye

	for pass = 0; pass < jumlahKampanye-1; pass++ {
		idx = pass
		for j = pass + 1; j < jumlahKampanye; j++ {
			if daftarKampanye[j].Progress < daftarKampanye[idx].Progress {
				idx = j
			}
		}

		if idx != pass {
			temp = daftarKampanye[pass]
			daftarKampanye[pass] = daftarKampanye[idx]
			daftarKampanye[idx] = temp
		}
	}
}

func sortSelectionDesc(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var pass, j, idx int
	var temp Kampanye

	for pass = 0; pass < jumlahKampanye-1; pass++ {
		idx = pass
		for j = pass + 1; j < jumlahKampanye; j++ {
			if daftarKampanye[j].Progress > daftarKampanye[idx].Progress {
				idx = j
			}
		}

		if idx != pass {
			temp = daftarKampanye[pass]
			daftarKampanye[pass] = daftarKampanye[idx]
			daftarKampanye[idx] = temp
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

func findNamaDonasi(daftarDonasi tabDonasi, jumlahDonasi int, pilihanNama string) tabDonasi {
	var i, idxArr int
	var arrDonasi tabDonasi

	idxArr = 0
	for i = 0; i < jumlahDonasi; i++ {
		if daftarDonasi[i].NamaDonatur == pilihanNama {
			arrDonasi[idxArr] = daftarDonasi[i]
			idxArr++
		}
	}
	return arrDonasi
}

func findIdDonasi(daftarDonasi tabDonasi, jumlahDonasi int, pilihanId int, maxIdx int) tabDonasi {
	var i, idxArr int
	var arrDonasi tabDonasi

	idxArr = 0

	if maxIdx > maxDonasi {
		maxIdx = maxDonasi
	} else if maxIdx <= 0 {
		maxIdx = 5
	}

	for i = 0; i < jumlahDonasi; i++ {
		if daftarDonasi[i].KampanyeId == pilihanId && idxArr < maxIdx {
			arrDonasi[idxArr] = daftarDonasi[i]
			idxArr++
		}
	}
	fmt.Println(idxArr)
	return arrDonasi
}

func findJudul(daftarKampanye tabKampanye, jumlahKampanye int, pilihanId int) string {
	var i int

	for i = 0; i < jumlahKampanye; i++ {
		if daftarKampanye[i].Id == pilihanId {
			return daftarKampanye[i].Judul
		}
	}
	return ""
}

func totalDonasiDonatur(daftarDonasi tabDonasi, jumlahDonasi int, pilihanNama string) int {
	var i int
	var donasi Donasi
	var total int

	for i = 0; i < jumlahDonasi; i++ {
		donasi = daftarDonasi[i]
		if donasi.NamaDonatur == pilihanNama {
			total += donasi.Jumlah
		} else if string(donasi.KampanyeId) == (pilihanNama) {
			total += donasi.Jumlah
		}
	}
	return total
}

func tampilNamaDonatur(daftarKampanye tabKampanye, jumlahKampanye int, arrKampanye tabDonasi) {
	var i, j int
	var nama, blank string
	var printed bool

	nama = arrKampanye[0].NamaDonatur
	fmt.Printf("=     Nama Donatur: %s     =\n", nama)

	for i = 0; i < jumlahKampanye; i++ {
		printed = true
		for j = 0; j < len(arrKampanye); j++ {
			if arrKampanye[j].KampanyeId == daftarKampanye[i].Id && printed {
				fmt.Printf("ID Kampanye: %d\n", daftarKampanye[i].Id)
				fmt.Printf("Judul: %s\n", daftarKampanye[i].Judul)
				fmt.Printf("Kategori: %s\n", daftarKampanye[i].Kategori)
				fmt.Printf("Deskripsi: \n%s\n", daftarKampanye[i].Deskripsi)
				fmt.Printf("Jumlah: %d\n", arrKampanye[j].Jumlah)
				fmt.Println()
				printed = false
			}
		}
	}
	fmt.Scan(&blank)
}

func tampilKampanyeDonatur(daftarDonasi tabDonasi, jumlahDonasi int, daftarKampanye tabKampanye, jumlahKampanye int, pilihanId int) {
	var i int
	var blank string

	insertionSortDesc(&daftarDonasi, jumlahDonasi)

	fmt.Println("Judul Kampanye: ", findJudul(daftarKampanye, jumlahKampanye, pilihanId))
	fmt.Println()
	for i = 0; i < jumlahDonasi; i++ {
		if daftarDonasi[i].KampanyeId == pilihanId {
			fmt.Printf("Nama Donatur: %s\n", daftarDonasi[i].NamaDonatur)
			fmt.Printf("Jumlah Donasi: Rp %d\n", daftarDonasi[i].Jumlah)
			fmt.Println()
		}
	}
	fmt.Scan(&blank)
}

func menuUtama(penggunaMasuk Pengguna) int {
	var pilihan int

	fmt.Println("     ==== SISTEM CROWDFUNDING ====    ")

	if penggunaMasuk.Peran != "" {
		fmt.Printf("Selamat datang, %s (%s)\n", penggunaMasuk.Username, penggunaMasuk.Peran)
	} else {
		fmt.Println("Daftar (1)")
		fmt.Println("Masuk (2)")
	}

	fmt.Println("Lihat Donasi (3)")
	fmt.Println("Lihat Kampanye (4)")
	fmt.Println("Berdonasi (5)")

	if penggunaMasuk.Peran == "admin" {
		fmt.Println("Buat Kampanye (6)")
	}

	if penggunaMasuk.Peran != "" {
		fmt.Println("Log Out (7)")
	}

	if penggunaMasuk.Peran == "admin" {
		fmt.Println("Prediksi Pencapaian Target (8)")
	}

	fmt.Println("Ketik '-1' untuk keluar")
	fmt.Print("Pilih menu (ketik angka tersebut): ")
	fmt.Scan(&pilihan)
	fmt.Println()

	return pilihan
}

func logOut(penggunaMasuk *Pengguna) {
	penggunaMasuk.Peran = ""
	penggunaMasuk.Username = ""
	penggunaMasuk.Email = ""
	penggunaMasuk.Password = ""
	fmt.Println("    Anda telah keluar.")
}

func prediksiPencapaianTarget(daftarKampanye tabKampanye, daftarDonasi tabDonasi, jumlahKampanye int, jumlahDonasi int) {
	var i, j int
	var totalDonasi, jumlahTransaksi int
	var rataRataDonasi float64
	var sisaTarget int
	var estimasiTransaksi float64
	var blank string

	fmt.Println("   === PREDIKSI PENCAPAIAN TARGET ===")

	for i = 0; i < jumlahKampanye; i++ {

		if daftarKampanye[i].Status == "aktif" {
			totalDonasi = 0
			jumlahTransaksi = 0

			for j = 0; j < jumlahDonasi; j++ {
				if daftarDonasi[j].KampanyeId == daftarKampanye[i].Id {
					totalDonasi = totalDonasi + daftarDonasi[j].Jumlah
					jumlahTransaksi = jumlahTransaksi + 1
				}
			}

			if jumlahTransaksi > 0 {
				rataRataDonasi = float64(totalDonasi) / float64(jumlahTransaksi)
				sisaTarget = daftarKampanye[i].Target - daftarKampanye[i].Terkumpul
				estimasiTransaksi = math.Ceil(float64(sisaTarget) / rataRataDonasi)

				fmt.Printf("Kampanye: %s\n", daftarKampanye[i].Judul)
				fmt.Printf("Sisa target: Rp %d\n", sisaTarget)
				fmt.Printf("Rata-rata donasi: Rp %.2f\n", rataRataDonasi)
				fmt.Printf("Estimasi jumlah transaksi untuk mencapai target: %.0f Transaksi\n", estimasiTransaksi)
				fmt.Printf("Progress: %d%% [%s]\n\n", daftarKampanye[i].Progress, createProgressBar(daftarKampanye[i].Progress))
			}
		}
	}
	fmt.Scanln(&blank)
}

func dummyDataKampanye(daftarKampanye *tabKampanye, jumlahKampanye *int) {

	daftarKampanye[0].Id = 100
	daftarKampanye[0].Judul = "Bantuan_Pendidikan_Anak_Yatim"
	daftarKampanye[0].Kategori = "Pendidikan"
	daftarKampanye[0].Deskripsi = "Program_bantuan_biaya_sekolah_untuk_anak-anak_yatim_dan_dhuafa"
	daftarKampanye[0].Target = 5000000
	daftarKampanye[0].Terkumpul = 2500000
	daftarKampanye[0].Progress = 50
	daftarKampanye[0].Status = "aktif"

	daftarKampanye[1].Id = 110
	daftarKampanye[1].Judul = "Renovasi_Masjid_Al-Ikhlas"
	daftarKampanye[1].Kategori = "Religi"
	daftarKampanye[1].Deskripsi = "Perbaikan_atap_dan_lantai_masjid_yang_sudah_rusak"
	daftarKampanye[1].Target = 10000000
	daftarKampanye[1].Terkumpul = 7500000
	daftarKampanye[1].Progress = 75
	daftarKampanye[1].Status = "aktif"

	daftarKampanye[2].Id = 120
	daftarKampanye[2].Judul = "Bantuan_Korban_Bencana_Alam"
	daftarKampanye[2].Kategori = "Bencana"
	daftarKampanye[2].Deskripsi = "Bantuan_logistik_dan_tempat_tinggal_untuk_korban_banjir"
	daftarKampanye[2].Target = 20000000
	daftarKampanye[2].Terkumpul = 15000000
	daftarKampanye[2].Progress = 75
	daftarKampanye[2].Status = "aktif"

	daftarKampanye[3].Id = 130
	daftarKampanye[3].Judul = "Kampus_Bersih_dan_Hijau"
	daftarKampanye[3].Kategori = "Lingkungan"
	daftarKampanye[3].Deskripsi = "Program_penanaman_pohon_dan_pengelolaan_sampah_di_kampus"
	daftarKampanye[3].Target = 3000000
	daftarKampanye[3].Terkumpul = 3000000
	daftarKampanye[3].Progress = 100
	daftarKampanye[3].Status = "selesai"

	daftarKampanye[4].Id = 140
	daftarKampanye[4].Judul = "Beasiswa_Mahasiswa_Berprestasi"
	daftarKampanye[4].Kategori = "Pendidikan"
	daftarKampanye[4].Deskripsi = "Program_beasiswa_untuk_mahasiswa_berprestasi_dari_keluarga_kurang_mampu"
	daftarKampanye[4].Target = 8000000
	daftarKampanye[4].Terkumpul = 4000000
	daftarKampanye[4].Progress = 50
	daftarKampanye[4].Status = "aktif"

	daftarKampanye[5].Id = 150
	daftarKampanye[5].Judul = "Posyandu_Sehat_Mandiri"
	daftarKampanye[5].Kategori = "Kesehatan"
	daftarKampanye[5].Deskripsi = "Pengadaan_alat_kesehatan_dan_vitamin_untuk_posyandu_desa"
	daftarKampanye[5].Target = 4500000
	daftarKampanye[5].Terkumpul = 2250000
	daftarKampanye[5].Progress = 50
	daftarKampanye[5].Status = "aktif"

	daftarKampanye[6].Id = 160
	daftarKampanye[6].Judul = "Gerakan_Literasi_Digital"
	daftarKampanye[6].Kategori = "Teknologi"
	daftarKampanye[6].Deskripsi = "Pelatihan_komputer_dan_internet_untuk_masyarakat_pedesaan"
	daftarKampanye[6].Target = 6000000
	daftarKampanye[6].Terkumpul = 1800000
	daftarKampanye[6].Progress = 30
	daftarKampanye[6].Status = "aktif"

	daftarKampanye[7].Id = 170
	daftarKampanye[7].Judul = "Pembangunan_Perpustakaan_Desa"
	daftarKampanye[7].Kategori = "Pendidikan"
	daftarKampanye[7].Deskripsi = "Membangun_perpustakaan_desa_dengan_koleksi_buku_yang_lengkap"
	daftarKampanye[7].Target = 12000000
	daftarKampanye[7].Terkumpul = 9600000
	daftarKampanye[7].Progress = 80
	daftarKampanye[7].Status = "aktif"

	daftarKampanye[8].Id = 180
	daftarKampanye[8].Judul = "Festival_Seni_Budaya_Lokal"
	daftarKampanye[8].Kategori = "Budaya"
	daftarKampanye[8].Deskripsi = "Penyelenggaraan_festival_untuk_melestarikan_seni_dan_budaya_daerah"
	daftarKampanye[8].Target = 7500000
	daftarKampanye[8].Terkumpul = 5625000
	daftarKampanye[8].Progress = 75
	daftarKampanye[8].Status = "aktif"

	daftarKampanye[9].Id = 190
	daftarKampanye[9].Judul = "Pemberdayaan_UMKM_Perempuan"
	daftarKampanye[9].Kategori = "Ekonomi"
	daftarKampanye[9].Deskripsi = "Program_pelatihan_dan_modal_usaha_untuk_ibu-ibu_rumah_tangga"
	daftarKampanye[9].Target = 9000000
	daftarKampanye[9].Terkumpul = 2700000
	daftarKampanye[9].Progress = 30
	daftarKampanye[9].Status = "aktif"

	*jumlahKampanye = 10
}

func dummyDataDonasi(daftarDonasi *tabDonasi, jumlahDonasi *int) {

	daftarDonasi[0].KampanyeId = 100
	daftarDonasi[0].NamaDonatur = "ammar"
	daftarDonasi[0].Jumlah = 1000000

	daftarDonasi[1].KampanyeId = 110
	daftarDonasi[1].NamaDonatur = "abrar"
	daftarDonasi[1].Jumlah = 1500000

	daftarDonasi[2].KampanyeId = 120
	daftarDonasi[2].NamaDonatur = "annisa"
	daftarDonasi[2].Jumlah = 2000000

	daftarDonasi[3].KampanyeId = 140
	daftarDonasi[3].NamaDonatur = "ammar"
	daftarDonasi[3].Jumlah = 1500000

	daftarDonasi[4].KampanyeId = 150
	daftarDonasi[4].NamaDonatur = "abrar"
	daftarDonasi[4].Jumlah = 800000

	daftarDonasi[5].KampanyeId = 160
	daftarDonasi[5].NamaDonatur = "annisa"
	daftarDonasi[5].Jumlah = 1200000

	daftarDonasi[6].KampanyeId = 170
	daftarDonasi[6].NamaDonatur = "ammar"
	daftarDonasi[6].Jumlah = 900000

	daftarDonasi[7].KampanyeId = 180
	daftarDonasi[7].NamaDonatur = "abrar"
	daftarDonasi[7].Jumlah = 700000

	daftarDonasi[8].KampanyeId = 190
	daftarDonasi[8].NamaDonatur = "annisa"
	daftarDonasi[8].Jumlah = 2500000

	daftarDonasi[9].KampanyeId = 100
	daftarDonasi[9].NamaDonatur = "annisa"
	daftarDonasi[9].Jumlah = 500000

	*jumlahDonasi = 10
}

func main() {
	var pilihan int
	var jumlahPengguna, jumlahKampanye, jumlahDonasi int
	var daftarPengguna tabPengguna
	var daftarKampanye tabKampanye
	var daftarDonasi tabDonasi
	var penggunaMasuk Pengguna

	daftarPengguna[0].Email = "ammar@gmail.com"
	daftarPengguna[0].Username = "ammar"
	daftarPengguna[0].Password = "Ammar1234@"
	daftarPengguna[0].Peran = "admin"

	daftarPengguna[1].Email = "abrar@yahoo.com"
	daftarPengguna[1].Username = "abrar"
	daftarPengguna[1].Password = "Abrar1234@"
	daftarPengguna[1].Peran = "donatur"

	daftarPengguna[2].Email = "annisa@yahoo.com"
	daftarPengguna[2].Username = "annisa"
	daftarPengguna[2].Password = "Annisa1234@"
	daftarPengguna[2].Peran = "donatur"

	jumlahPengguna = 3

	dummyDataKampanye(&daftarKampanye, &jumlahKampanye)
	dummyDataDonasi(&daftarDonasi, &jumlahDonasi)

	for pilihan != -1 {
		pilihan = menuUtama(penggunaMasuk)

		if pilihan == 1 {
			if penggunaMasuk.Peran == "" {
				daftar(&daftarPengguna, &jumlahPengguna, &penggunaMasuk)
			} else {
				fmt.Println("    Anda sudah masuk sebagai", penggunaMasuk.Peran, "!")
				fmt.Println()
			}
		}

		if pilihan == 2 {
			if penggunaMasuk.Peran == "" {
				masuk(&daftarPengguna, &penggunaMasuk, jumlahPengguna)
			} else {
				fmt.Println("    Anda sudah masuk sebagai", penggunaMasuk.Peran, "!")
				fmt.Println()
			}
		}

		if pilihan == 3 {
			tampilkanDonasi(&daftarDonasi, jumlahDonasi, daftarKampanye, jumlahKampanye)
		}

		if pilihan == 4 {
			detailKampanye(&daftarKampanye, jumlahKampanye)
		}

		if pilihan == 5 {
			tambahDonasi(&daftarKampanye, &daftarDonasi, &penggunaMasuk, jumlahKampanye, &jumlahDonasi)
		}

		if pilihan == 6 {
			if penggunaMasuk.Peran == "admin" {
				buatKampanye(&daftarKampanye, &jumlahKampanye, &penggunaMasuk)
			} else {
				fmt.Println("    Hanya admin yang dapat membuat kampanye!")
				fmt.Println()
			}
		}

		if pilihan == 7 {
			if penggunaMasuk.Peran != "" {
				logOut(&penggunaMasuk)
			} else {
				fmt.Println("   Silakan masuk terlebih dahulu!")
				fmt.Println()
			}
		}

		if pilihan == 8 && penggunaMasuk.Peran == "admin" && jumlahKampanye > 0 && jumlahDonasi > 0 {
			prediksiPencapaianTarget(daftarKampanye, daftarDonasi, jumlahKampanye, jumlahDonasi)
		} else if pilihan == 8 && penggunaMasuk.Peran != "admin" {
			fmt.Println("   Hanya admin yang dapat mengakses fitur ini!")
			fmt.Println()
		} else if pilihan == 8 && (jumlahKampanye == 0 || jumlahDonasi == 0) {
			fmt.Println("   Tidak ada kampanye atau donasi yang tersedia untuk prediksi!")
			fmt.Println()
		}

		if pilihan == -1 {
			fmt.Println("============================================")
			fmt.Println("   Terima kasih telah menggunakan aplikasi  ")
			fmt.Println("            SISTEM CROWDFUNDING             ")
			fmt.Println("  Semoga hari-hari KAMUUU menyenangkan WELLL!    ")
			fmt.Println("============================================")
		} else if pilihan < 1 || pilihan > 8 {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}