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
	Nama     string
	Password string
	Peran    string
}

type Kampanye struct {
	Id        int
	Judul     string
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

func daftar(daftarPengguna *tabPengguna, jumlahPengguna *int) {
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
	var judul, deskripsi string
	var target int

	if penggunaMasuk.Peran != "admin" {
		fmt.Println("Hanya admin yang dapat membuat kampanye!")
	} else if *jumlahKampanye >= maxKampanye {
		fmt.Println("Jumlah kampanye sudah mencapai batas maksimum!")
	} else {
		fmt.Println("\n=== BUAT KAMPANYE BARU ===")
		fmt.Println("Judul: ")
		fmt.Scan(&judul)
		fmt.Println("Deskripsi: ")
		fmt.Scan(&deskripsi)
		fmt.Println("Target Dana: ")
		fmt.Scan(&target)

		daftarKampanye[*jumlahKampanye].Deskripsi = deskripsi
		daftarKampanye[*jumlahKampanye].Judul = judul
		daftarKampanye[*jumlahKampanye].Target = target
		daftarKampanye[*jumlahKampanye].Status = "aktif"

		*jumlahKampanye++
		fmt.Println("Kampanye berhasil dibuat!")
	}
}

func tampilkanKampanye(daftarKampanye *tabKampanye, jumlahKampanye int) {
	var i, progress int
	var k Kampanye

	fmt.Println("=              DAFTAR KAMPANYE AKTIF                 =")

	if jumlahKampanye == 0 {
		fmt.Println("=           Belum ada kampanye yang dibuat            =")

	}
	for i = 0; i < jumlahKampanye; i++ {
		k = daftarKampanye[i]
		progress = (k.Terkumpul / k.Target) * 100

		fmt.Printf("[%d] %s\n", i, k.Judul)
		fmt.Printf("Status: %s\n", k.Status)
		fmt.Printf("Progress: %v [%s]\n", progress, createProgressBar(progress))
		fmt.Printf("Terkumpul: Rp %v dari Rp %v\n", k.Terkumpul, k.Target)
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

func main() {
	var pilihan int
	var jumlahPengguna, jumlahKampanye, jumlahDonasi int
	var daftarPengguna tabPengguna
	var daftarKampanye tabKampanye
	var daftarDonasi tabDonasi
	var penggunaMasuk Pengguna // Pengguna yang sedang masuk sekarang

	for pilihan != -1 {
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

		switch pilihan {
		case 1:
			daftar(&daftarPengguna, &jumlahPengguna)
		case 2:
			masuk(&daftarPengguna, &penggunaMasuk, jumlahPengguna)
		case 3:
			tampilkanKampanye(&daftarKampanye, jumlahKampanye)
		case 4:
			tambahDonasi(&daftarKampanye, &daftarDonasi, &penggunaMasuk, jumlahKampanye, &jumlahDonasi)
		case 5:
			if penggunaMasuk.Peran == "admin" {
				buatKampanye(&daftarKampanye, &jumlahKampanye, &penggunaMasuk)
			} else {
				fmt.Println("Hanya admin yang dapat membuat kampanye!")
			}
		case 6:
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
