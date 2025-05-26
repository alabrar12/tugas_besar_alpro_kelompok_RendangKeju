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
	var lenEmail, i int
	var dotCom, domain bool
	var emailCopy string
	// Email harus mengandung '@gmail.com' atau '@yahoo.com' atau '@outlook.com'

	if *allow {
		fmt.Print("Email: ")
		fmt.Scan(&*email)
	}

	for (!dotCom || !domain) && *allow {
		*allow = *email != "keluar"
		lenEmail = len(*email) - 1

		dotCom = false
		domain = false

		if lenEmail > 11 && lenEmail < 50 && *allow {
			emailCopy = *email

			dotCom = string(emailCopy[lenEmail]) == "m" && string(emailCopy[lenEmail-1]) == "o" && string(emailCopy[lenEmail-2]) == "c" && string(emailCopy[lenEmail-3]) == "."

			domain = domain || string(emailCopy[lenEmail-4]) == "l" && string(emailCopy[lenEmail-5]) == "i" && string(emailCopy[lenEmail-6]) == "a" && string(emailCopy[lenEmail-7]) == "m" && string(emailCopy[lenEmail-8]) == "g" && string(emailCopy[lenEmail-9]) == "@"
			domain = domain || string(emailCopy[lenEmail-4]) == "o" && string(emailCopy[lenEmail-5]) == "o" && string(emailCopy[lenEmail-6]) == "h" && string(emailCopy[lenEmail-7]) == "a" && string(emailCopy[lenEmail-8]) == "y" && string(emailCopy[lenEmail-9]) == "@"
			domain = domain || string(emailCopy[lenEmail-4]) == "k" && string(emailCopy[lenEmail-5]) == "o" && string(emailCopy[lenEmail-6]) == "o" && string(emailCopy[lenEmail-7]) == "l" && string(emailCopy[lenEmail-8]) == "t" && string(emailCopy[lenEmail-9]) == "u" && string(emailCopy[lenEmail-10]) == "o" && string(emailCopy[lenEmail-11]) == "@"
		} else if !(lenEmail > 11 && lenEmail < 50) && *allow {
			fmt.Println("Email harus memiliki panjang antara 12 hingga 50 karakter.")
			fmt.Print("Masukkan email yang valid: ")
			fmt.Scan(&*email)
		} 

		if (!dotCom || !domain) && *allow {
			fmt.Println("Email harus mengandung '@gmail.com', '@yahoo.com' atau '@outlook.com'")
			fmt.Print("Masukkan email yang valid: ")
			fmt.Scan(&*email)
		} else if *allow {
			for i = 0; i < jumlahPengguna; i++ {
				if daftarPengguna[i].Email == *email{
					fmt.Println("Email sudah terdaftar!")
					fmt.Print("Masukkan email yang valid: ")
					fmt.Scan(&*email)
				}
			}

		}
	}

}

func verikasiUsername(daftarPengguna tabPengguna, jumlahPengguna int, username *string, allow *bool) {
	var i int
	var usernameValid bool

	usernameValid = false

	if *allow {
		fmt.Print("Username: ")
		fmt.Scan(&*username)
	}

	for !usernameValid && *allow {
		usernameValid = true
		*allow = *username != "keluar"
		if len(*username) < 4 || len(*username) > 50 {
			fmt.Println("Username harus memiliki panjang antara 4 hingga 50 karakter.")
			usernameValid = false
		} else {
			for i = 0; i < jumlahPengguna; i++ {
				if daftarPengguna[i].Username == *username {
					fmt.Println("Username sudah terdaftar!")
					usernameValid = false
				}
			}
			usernameValid = usernameValid && len(*username) >= 4 && len(*username) <= 50
		}
		if !usernameValid {
			fmt.Print("Masukkan username yang valid: ")
			fmt.Scan(&*username)
		}
	}
}

func verikasiPassword(pass *string, allow *bool) {
	var i int
	var upper, lower, number, special, moreEight, passValid bool
	var password string

	if *allow {
		fmt.Print("Password - Gunakan minimal 8 huruf dengan isi huruf besar, angka, dan simbol(@,#,$,%,&) : ")
		fmt.Scan(&*pass)
	}

	for !passValid && *allow {
		*allow = *pass != "keluar"
		password = *pass
		if len(password) < 8 && *allow {
			moreEight = false
			fmt.Println("Password harus memiliki minimal 8 karakter.")
		} else if *allow {
			moreEight = true
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
			if !moreEight || !upper || !lower || !number || !special && *allow {
				fmt.Println("Password harus mengandung minimal 1 huruf besar, 1 huruf kecil, 1 angka, dan 1 simbol (@, #, $, %, &).")
			}
		}

		if !(moreEight && upper && lower && number && special) && *allow {
			fmt.Print("Masukkan password yang valid: ")
			fmt.Scan(&*pass)
		} else if *allow {
			passValid = true
		}
	}
}

func verikasiPeran(peran *string, allow *bool) {
	var peranValid bool
	peranValid = false

	if *allow {
		fmt.Print("Peran (admin/donatur): ")
		fmt.Scan(&*peran)
	}

	for !peranValid && *allow {
		*allow = *peran != "keluar"
		if *peran != "admin" && *peran != "donatur" {
			fmt.Println("Peran harus 'admin' atau 'donatur'.")
			fmt.Print("Masukkan peran yang valid: ")
			fmt.Scan(&*peran)
		} else if *allow {
			peranValid = true
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
	for i = 0; i < jumlahPengguna; i++ {
		if (daftarPengguna[i].Email == user || daftarPengguna[i].Username == user) && (daftarPengguna[i].Password == password) {
			*penggunaMasuk = daftarPengguna[i]
			fmt.Printf("Selamat datang, %s!\n", penggunaMasuk.Username)
			loginValid = true
		} 
	}

	if !loginValid {
		fmt.Println()
		fmt.Println("   Email/Username atau Password salah!")
		fmt.Println()
	}
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
		fmt.Print("Judul: ")
		fmt.Scan(&judul)
		fmt.Print("Kategori: ")
		fmt.Scan(&kategori)
		fmt.Println("Deskripsi: ")
		fmt.Scan(&deskripsi)
		fmt.Print("Target Dana (Rupiah): ")
		fmt.Scan(&target)

		daftarKampanye[*jumlahKampanye].Deskripsi = deskripsi
		daftarKampanye[*jumlahKampanye].Judul = judul
		daftarKampanye[*jumlahKampanye].Target = target
		daftarKampanye[*jumlahKampanye].Kategori = kategori
		daftarKampanye[*jumlahKampanye].Status = "aktif"

		*jumlahKampanye++
		fmt.Println()
		fmt.Println("-  Kampanye berhasil dibuat!  -")
		fmt.Println()
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
	var blank string

	fmt.Println("\n=== DETAIL KAMPANYE ===")
	fmt.Printf("Judul: %s\n", kampanye.Judul)
	fmt.Printf("Kategori: %s\n", kampanye.Kategori)
	fmt.Printf("Deskripsi: \n%s\n", kampanye.Deskripsi)
	fmt.Printf("Target Dana: Rp %d\n", kampanye.Target)
	fmt.Printf("Terkumpul: Rp %d\n", kampanye.Terkumpul)
	fmt.Printf("Status: %s\n", kampanye.Status)
	fmt.Printf("Progress: %d%% [%s]\n", kampanye.Progress, createProgressBar(kampanye.Progress))
	fmt.Scan(&blank)
}

func tampilkanKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var i, progress, ascending int
	var k Kampanye
	var tempKampanye tabKampanye

	tempKampanye = *daftarKampanye

	fmt.Print("Urutkan berdasarkan progress (1: ascending, 0: descending): ")
	fmt.Scan(&ascending)

	if ascending == 1 {
		sortSelectionAsc(&tempKampanye, jumlahKampanye)
	} else if ascending == 0 {
		sortSelectionDesc(&tempKampanye, jumlahKampanye)
	} else {
		fmt.Println("Pilihan tidak valid, menggunakan urutan default (ascending)")
		ascending = 1
	}

	if jumlahKampanye == 0 {
		fmt.Println("=           Belum ada kampanye yang dibuat            =")
	} else {
		fmt.Println("=           Jumlah kampanye yang dibuat: ", jumlahKampanye, "          =")
		for i = 0; i < jumlahKampanye; i++ {
			k = tempKampanye[i]
			progress = (k.Terkumpul * 100) / k.Target

			fmt.Printf("[%d] %s\n", i, k.Judul)
			fmt.Printf("Status: %s\n", k.Status)
			fmt.Printf("Kategori: %s\n", k.Kategori)
			fmt.Printf("Progress: %v%% [%s]\n", progress, createProgressBar(progress))
			// fmt.Printf("Terkumpul: Rp %v dari Rp %v\n", k.Terkumpul, k.Target)
			fmt.Println()
		}
	}
}

func detailKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var pilihan, pilihanIdx, pilihanId int
	var pilihanJudul, pilihanKategori, blank string
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
		fmt.Println("3. Kategori Kampanye")
		fmt.Println("Kembali ke menu utama (0)")
		fmt.Println("Pilih Kampanye untuk melihat detail (ketik Angka): ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 0:
			fmt.Println("    Kembali ke menu utama")
		case 1:
			fmt.Print("Ketik ID Kampanye: ")
			fmt.Scan(&pilihanId)
			pilihanIdx = findIdKampanye(tempKampanye, jumlahKampanye, pilihanId)
		case 2:
			fmt.Print("Ketik Judul Kampanye: ")
			fmt.Scan(&pilihanJudul)
			pilihanIdx = findJudulKampanye(tempKampanye, jumlahKampanye, pilihanJudul)
		case 3:
			fmt.Print("Ketik Kategori Kampanye: ")
			fmt.Scan(&pilihanKategori)
			pilihanIdx = findKategoriKampanye(tempKampanye, jumlahKampanye, pilihanKategori)
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		if pilihan >= 1 && pilihan <= 3 {
			tampilkanDetailKampanye(tempKampanye[pilihanIdx])
		}
	}
	fmt.Scanln(&blank)
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
		fmt.Println("    Silakan masuk terlebih dahulu!")
	} else if *jumlahDonasi >= maxDonasi {
		fmt.Println("    Jumlah donasi sudah mencapai batas maksimum!")
		return
	} else if jumlahKampanye == 0 {
		fmt.Println("    Belum ada kampanye yang dibuat!")
		return
	} else {
		fmt.Println("===   DONASI KAMPANYE   ===")

		tampilkanKampanye(daftarKampanye, jumlahKampanye)
		if jumlahKampanye != 0 {
			fmt.Print("ID Kampanye: ")
			fmt.Scan(&kampanyeId)

			if kampanyeId >= jumlahKampanye || daftarKampanye[kampanyeId].Status != "aktif" || kampanyeId < 0 {
				fmt.Println("Kampanye tidak valid atau tidak aktif!")
			} else {
				fmt.Print("Jumlah Donasi: ")
				fmt.Scan(&jumlah)

				daftarDonasi[*jumlahDonasi].KampanyeId = kampanyeId
				daftarDonasi[*jumlahDonasi].NamaDonatur = penggunaMasuk.Username
				daftarDonasi[*jumlahDonasi].Jumlah = jumlah

				totalDonasi = daftarKampanye[kampanyeId].Terkumpul + jumlah

				if totalDonasi > daftarKampanye[kampanyeId].Target {
					fmt.Println("Jumlah donasi melebihi target kampanye!,", "Donasi dikembalikan sebesar Rp", totalDonasi)
					totalDonasi -= daftarKampanye[kampanyeId].Target
				}

				daftarKampanye[kampanyeId].Terkumpul = totalDonasi
				daftarKampanye[kampanyeId].Progress = (totalDonasi * 100) / daftarKampanye[kampanyeId].Target

				if daftarKampanye[kampanyeId].Progress >= 100 {
					daftarKampanye[kampanyeId].Status = "selesai"
					daftarKampanye[kampanyeId].Progress = 100
					fmt.Println("=     Kampanye telah selesai, terkumpul Rp", daftarKampanye[kampanyeId].Terkumpul, "dari target Rp", daftarKampanye[kampanyeId].Target, "     =")

				}
				*jumlahDonasi = *jumlahDonasi + 1
				fmt.Println()
				fmt.Println("Donasi berhasil! Terima kasih.")
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
	for i = 0; i < jumlahDonasi; i++ {
		if daftarDonasi[i].KampanyeId == pilihanId && idxArr < maxIdx {
			arrDonasi[idxArr] = daftarDonasi[i]
			idxArr++
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
		} else if string(rune(donasi.KampanyeId)) == pilihanNama {
			total += donasi.Jumlah
		}
	}
	return total
}

func tampilNamaDonatur(daftarKampanye tabKampanye, jumlahKampanye int, arrKampanye tabDonasi) {
	var i, j int
	var blank, nama string
	var printed bool

	nama = arrKampanye[0].NamaDonatur
	fmt.Printf("=     Nama Donatur: %s     =\n", nama)
	for i = 0; i < jumlahKampanye; i++ {
		printed = true
		for j = 0; j < len(arrKampanye) && printed; j++ {
			if arrKampanye[j].KampanyeId == daftarKampanye[i].Id {
				fmt.Printf("ID Kampanye: %d\n", daftarKampanye[i].Id)
				fmt.Printf("Judul: %s\n", daftarKampanye[i].Judul)
				fmt.Printf("Kategori: %s\n", daftarKampanye[i].Kategori)
				fmt.Printf("Deskripsi: \n%s\n", daftarKampanye[i].Deskripsi)
				fmt.Printf("Target Dana: Rp %d\n", daftarKampanye[i].Target)
				fmt.Printf("Terkumpul: Rp %d\n", daftarKampanye[i].Terkumpul)
				fmt.Println()
				printed = false
			}
		}
		fmt.Scan(&blank)
	}
}

func tampilKampanyeDonatur(daftarDonasi tabDonasi, jumlahDonasi int) {
	var i int
	var blank string

	insertionSortDesc(&daftarDonasi, jumlahDonasi)

	for i = 0; i < jumlahDonasi; i++ {
		fmt.Printf("Nama Donatur: %s\n", daftarDonasi[i].NamaDonatur)
		fmt.Printf("Jumlah Donasi: Rp %d\n", daftarDonasi[i].Jumlah)
		fmt.Println()
	}
	fmt.Scan(&blank)
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
		switch pilihan {
		case 0:
			fmt.Println("Kembali ke menu utama")
			return
		case 1:
			fmt.Println("Ketik Nama Donatur: ")
			fmt.Scan(&pilihanNama)
			fmt.Print("Mau diurutkan berdasarkan (1: Ascending, 2: Descending): ")
			fmt.Scan(&pilihanUrutan)
			if pilihanUrutan == 1 {
				insertionSortAsc(&tempDaftarDonasi, jumlahDonasi)
			} else if pilihanUrutan == 2 {
				insertionSortDesc(&tempDaftarDonasi, jumlahDonasi)
			} else {
				fmt.Println("Pilihan tidak valid, menggunakan urutan default (descending)")
				insertionSortDesc(&tempDaftarDonasi, jumlahDonasi)
			}
			detailDonatur = findNamaDonasi(tempDaftarDonasi, jumlahDonasi, pilihanNama)
			if len(detailDonatur) > 0 {
				fmt.Printf("Detail Donasi (%v):\n", pilihanNama)
				fmt.Printf("Total Donasi: Rp %d\n", totalDonasiDonatur(tempDaftarDonasi, jumlahDonasi, pilihanNama))
				tampilNamaDonatur(daftarKampanye, jumlahKampanye, detailDonatur)
			}
		case 2:
			fmt.Print("Ketik ID Kampanye: ")
			fmt.Scan(&pilihanId)
			fmt.Print("Maximal Donasi yang ditampilkan: ")
			fmt.Scan(&pilihanMaxIdx)
			detailDonatur = findIdDonasi(tempDaftarDonasi, jumlahDonasi, pilihanId, pilihanMaxIdx)
			if len(detailDonatur) > 0 {
				fmt.Printf("Detail Donasi untuk Kampanye ID %d:\n", pilihanId)
				fmt.Printf("Total Donasi: Rp %d\n", totalDonasiDonatur(detailDonatur, jumlahDonasi, string(rune(pilihanId))))
				tampilKampanyeDonatur(detailDonatur, jumlahDonasi)
			}
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
	fmt.Println()
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

func main() {
	var pilihan int
	var jumlahPengguna, jumlahKampanye, jumlahDonasi int
	var daftarPengguna tabPengguna
	var daftarKampanye tabKampanye
	var daftarDonasi tabDonasi
	var penggunaMasuk Pengguna // Pengguna yang sedang masuk sekarang

	daftarPengguna[0] = Pengguna{
		Email:    "ammar@gmail.com",
		Username: "ammar",
		Password: "Ammar1234@",
		Peran:    "admin",
	}

	daftarPengguna[1] = Pengguna{
		Email:    "ghifari@yahoo.com",
		Username: "ghifari",
		Password: "Ghifari1234@",
		Peran:    "donatur",
	}

	jumlahPengguna = 2 // Jumlah pengguna awal

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
			detailKampanye(&daftarKampanye, jumlahKampanye)
		case 5:
			tambahDonasi(&daftarKampanye, &daftarDonasi, &penggunaMasuk, jumlahKampanye, &jumlahDonasi)
		case 6:
			if penggunaMasuk.Peran == "admin" {
				buatKampanye(&daftarKampanye, &jumlahKampanye, &penggunaMasuk)
			} else {
				fmt.Println("    Hanya admin yang dapat membuat kampanye!")
				fmt.Println()
			}
		case 7:
			if penggunaMasuk.Peran != "" {
				logOut(&penggunaMasuk)
			} else {
				fmt.Println("   Silakan masuk terlebih dahulu!")
				fmt.Println()
			}
		case -1:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}