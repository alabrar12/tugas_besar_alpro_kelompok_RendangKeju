# 💰 Sistem Crowdfunding - Platform Penggalangan Dana Digital

Sebuah aplikasi penggalangan dana digital yang dibuat menggunakan bahasa pemrograman Go. Aplikasi ini memungkinkan pengguna untuk membuat kampanye penggalangan dana dan menerima donasi dari para donatur.

## 📋 Daftar Isi

1. [Tentang Aplikasi](#tentang-aplikasi)
2. [Fitur Utama](#fitur-utama)
3. [Cara Menjalankan](#cara-menjalankan)
4. [Cara Menggunakan](#cara-menggunakan)
5. [Struktur Data](#struktur-data)
6. [Algoritma yang Digunakan](#algoritma-yang-digunakan)
7. [Batasan Sistem](#batasan-sistem)
8. [Contoh Penggunaan](#contoh-penggunaan)

## 🎯 Tentang Aplikasi

Sistem Crowdfunding ini adalah aplikasi berbasis terminal yang memungkinkan:
- **Admin** dapat membuat dan mengelola kampanye penggalangan dana
- **Donatur** dapat memberikan donasi pada kampanye yang mereka pilih
- Semua pengguna dapat melihat daftar kampanye dan riwayat donasi

Aplikasi ini dibangun sebagai tugas besar mata kuliah Algoritma Pemrograman dengan mengimplementasikan berbagai struktur data dan algoritma sorting/searching.

## ✨ Fitur Utama

### 🔐 Sistem Autentikasi
- **Pendaftaran Pengguna Baru**: Validasi email, username, dan password yang kuat
- **Login/Logout**: Sistem keamanan dengan peran pengguna (admin/donatur)
- **Validasi Email**: Mendukung domain @gmail.com, @yahoo.com, dan @outlook.com

### 👥 Manajemen Pengguna
- **Dua Jenis Peran**:
  - **Admin**: Dapat membuat kampanye dan melihat prediksi target
  - **Donatur**: Dapat berdonasi dan melihat riwayat donasi

### 🎯 Manajemen Kampanye
- **Buat Kampanye Baru** (khusus admin)
- **Lihat Daftar Kampanye** dengan sorting berdasarkan progress
- **Detail Kampanye** dengan progress bar visual
- **Status Kampanye**: Aktif atau Selesai
- **Kategori Beragam**: Pendidikan, Kesehatan, Teknologi, dll.

### 💸 Sistem Donasi
- **Berdonasi** pada kampanye aktif
- **Validasi Donasi**: Tidak boleh melebihi target kampanye
- **Auto-Complete**: Kampanye otomatis selesai saat target tercapai
- **Riwayat Donasi** lengkap dengan sorting

### 📊 Laporan dan Analisis
- **Total Donasi** keseluruhan sistem
- **Filter Donasi** berdasarkan nama donatur atau ID kampanye
- **Prediksi Pencapaian Target** menggunakan rata-rata donasi
- **Progress Bar Visual** untuk setiap kampanye

## 🚀 Cara Menjalankan

### Prasyarat
Pastikan Go (Golang) sudah terinstall di komputer Anda. Untuk mengecek:
```bash
go version
```

### Menjalankan Aplikasi
1. Buka terminal/command prompt
2. Navigasi ke folder project:
   ```bash
   cd "/Users/bangjhener/Desktop/Telkom/Semester 2/Algoritma Pemrograman/Tugas Anak Muda/tugas_besar_alpro_kelompok_RendangKeju"
   ```
3. Jalankan aplikasi:
   ```bash
   go run crowdfunding.go
   ```

## 📖 Cara Menggunakan

### 1. Menu Utama
Saat aplikasi dijalankan, Anda akan melihat menu utama:
```
==== SISTEM CROWDFUNDING ====
Daftar (1)
Masuk (2)
Lihat Donasi (3)
Lihat Kampanye (4)
Berdonasi (5)
Ketik '-1' untuk keluar
```

### 2. Pendaftaran Pengguna Baru
- Pilih menu (1) untuk mendaftar
- Masukkan email yang valid (gmail.com, yahoo.com, atau outlook.com)
- Buat username unik (4-25 karakter)
- Password harus mengandung:
  - Minimal 8 karakter
  - Huruf besar dan kecil
  - Angka
  - Simbol (@, #, $, %, &)
- Pilih peran: admin atau donatur

### 3. Login
- Pilih menu (2) untuk masuk
- Masukkan email/username dan password

### 4. Fitur untuk Semua Pengguna

#### Lihat Donasi (3)
- Melihat semua riwayat donasi
- Filter berdasarkan nama donatur atau ID kampanye
- Sorting ascending/descending berdasarkan jumlah donasi

#### Lihat Kampanye (4)
- Daftar semua kampanye dengan progress bar
- Sorting berdasarkan progress kampanye
- Detail kampanye lengkap dengan deskripsi

#### Berdonasi (5)
- Pilih kampanye aktif untuk didonasi
- Masukkan jumlah donasi
- Sistem otomatis validasi tidak melebihi target

### 5. Fitur Khusus Admin

#### Buat Kampanye (6)
- Input judul, kategori, deskripsi, dan target dana
- Kampanye otomatis mendapat ID unik
- Status awal: aktif

#### Prediksi Pencapaian Target (8)
- Analisis berdasarkan rata-rata donasi existing
- Estimasi jumlah transaksi yang dibutuhkan
- Khusus untuk kampanye yang masih aktif

## 🗂️ Struktur Data

### Pengguna (User)
```go
type Pengguna struct {
    Email    string  // Email pengguna
    Username string  // Nama pengguna unik
    Password string  // Password terenkripsi
    Peran    string  // "admin" atau "donatur"
}
```

### Kampanye (Campaign)
```go
type Kampanye struct {
    Id        int     // ID unik kampanye
    Judul     string  // Judul kampanye
    Kategori  string  // Kategori (Pendidikan, Kesehatan, dll)
    Deskripsi string  // Deskripsi detail
    Target    int     // Target dana (Rupiah)
    Terkumpul int     // Dana terkumpul saat ini
    Progress  int     // Persentase progress (0-100%)
    Status    string  // "aktif" atau "selesai"
}
```

### Donasi (Donation)
```go
type Donasi struct {
    KampanyeId  int     // ID kampanye yang didonasi
    NamaDonatur string  // Username donatur
    Jumlah      int     // Jumlah donasi (Rupiah)
}
```

## 🔍 Algoritma yang Digunakan

### 1. Sorting Algorithms

#### Selection Sort (Descending)
Digunakan untuk mengurutkan kampanye berdasarkan progress dari tertinggi ke terendah:
```go
func sortSelectionDesc(daftarKampanye *tabKampanye, jumlahKampanye int)
```

#### Selection Sort (Ascending)
Mengurutkan kampanye dari progress terendah ke tertinggi:
```go
func sortSelectionAsc(daftarKampanye *tabKampanye, jumlahKampanye int)
```

#### Insertion Sort
Mengurutkan donasi berdasarkan jumlah (ascending/descending):
```go
func insertionSortAsc(daftarDonasi *tabDonasi, jumlahDonasi int)
func insertionSortDesc(daftarDonasi *tabDonasi, jumlahDonasi int)
```

### 2. Searching Algorithms

#### Binary Search
Mencari kampanye berdasarkan ID (data harus terurut):
```go
func findIdKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanId int) int
```

#### Linear Search
Mencari kampanye berdasarkan judul:
```go
func findJudulKampanye(daftarKampanye tabKampanye, jumlahKampanye int, pilihanJudul string) int
```

### 3. Filter dan Aggregation

#### Filter Donasi
- `findNamaDonasi()`: Filter donasi berdasarkan nama donatur
- `findIdDonasi()`: Filter donasi berdasarkan ID kampanye

#### Kalkulasi Statistik
- `totalDonasi()`: Menghitung total semua donasi
- `totalDonasiDonatur()`: Total donasi per donatur
- `prediksiPencapaianTarget()`: Prediksi menggunakan rata-rata

## ⚙️ Batasan Sistem

### Kapasitas Maximum
- **Pengguna**: 100 user
- **Kampanye**: 100 kampanye
- **Donasi**: 1,000 transaksi donasi

### Validasi Input
- **Email**: 12-25 karakter, domain terbatas
- **Username**: 4-25 karakter, harus unik
- **Password**: Minimal 8 karakter dengan kompleksitas tinggi
- **Donasi**: Tidak boleh melebihi sisa target kampanye

### Keamanan
- Password validation
- Role-based access control (admin and donatur)
- Session management untuk login/logout

## 💡 Contoh Penggunaan

### Scenario 1: Admin Membuat Kampanye
1. Login sebagai admin
2. Pilih menu "Buat Kampanye (6)"
3. Input data:
   - Judul: "Bantuan_Belajar_Online"
   - Kategori: "Pendidikan"
   - Deskripsi: "Laptop_untuk_siswa_kurang_mampu"
   - Target: 5000000
4. Kampanye otomatis mendapat ID dan status "aktif"

### Scenario 2: Donatur Berdonasi
1. Login sebagai donatur
2. Pilih menu "Lihat Kampanye (4)" untuk melihat daftar
3. Pilih menu "Berdonasi (5)"
4. Input ID kampanye dan jumlah donasi
5. Sistem validasi dan update progress otomatis

### Scenario 3: Melihat Laporan Donasi
1. Pilih menu "Lihat Donasi (3)"
2. Pilih filter berdasarkan nama donatur atau ID kampanye
3. Pilih sorting ascending/descending
4. Lihat total donasi dan detail transaksi

## 🎮 Data Demo

Aplikasi sudah dilengkapi dengan data demo untuk testing:

### Users Default:
- **Admin**: ammar@gmail.com / password: Ammar1234@
- **Donatur 1**: abrar@yahoo.com / password: Abrar1234@
- **Donatur 2**: annisa@yahoo.com / password: Annisa1234@

### Kampanye Demo:
- Bantuan Pendidikan Anak Yatim (50% progress)
- Renovasi Masjid Al-Ikhlas (75% progress)
- Bantuan Korban Bencana Alam (75% progress)
- Dan 7 kampanye lainnya

## 🤝 Kontribusi

Proyek ini merupakan tugas besar kelompok **RendangKeju** untuk mata kuliah Algoritma Pemrograman di Telkom University.

## 📝 Catatan Teknis

- **Bahasa**: Go (Golang)
- **Interface**: Command Line Interface (CLI)
- **Storage**: In-memory (data hilang saat aplikasi ditutup)
- **Algoritma**: Selection Sort, Insertion Sort, Binary Search, Linear Search
- **Pattern**: Procedural Programming dengan struct dan functions

---

*Dibuat dengan ❤️ oleh Kelompok RendangKeju - Telkom University*
