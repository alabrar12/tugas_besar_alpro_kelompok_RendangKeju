# Sistem Galang Dana (Crowdfunding)

Program komputer untuk galang dana yang berjalan di layar hitam (command line). Program ini dibuat dengan bahasa Go dan membantu orang membuat kampanye galang dana, memberikan donasi, dan melihat perkembangan dana yang terkumpul dengan tampilan yang mudah digunakan.

## Fitur Utama

### Daftar dan Masuk ke Program
- **Daftar Akun Baru**: Buat akun dengan email yang valid (harus pakai @gmail.com, @yahoo.com, atau @outlook.com)
- **Masuk ke Program**: Login pakai email atau nama pengguna dengan kata sandi
- **Dua Jenis Pengguna**: Ada admin (pengurus) dan donatur (pemberi dana) dengan hak akses yang berbeda
- **Keamanan Kata Sandi**: Kata sandi harus kuat (ada huruf besar, kecil, angka, dan simbol khusus)
- **Keluar Program**: Bisa logout untuk mengakhiri sesi

### Kelola Kampanye Galang Dana
- **Buat Kampanye Baru**: Hanya admin yang bisa membuat kampanye galang dana baru
- **Cari Kampanye**: Bisa mencari kampanye berdasarkan nomor ID atau judul
- **Lihat Detail**: Melihat informasi lengkap kampanye dengan grafik batang progress
- **Update Status**: Otomatis mengubah status menjadi "selesai" kalau target sudah tercapai 100%
- **Urutkan Kampanye**: Bisa mengurutkan kampanye berdasarkan persentase progress

### Sistem Donasi
- **Berdonasi**: Siapa saja bisa memberikan donasi ke kampanye yang masih aktif
- **Perlindungan Donasi**: Kalau donasi berlebihan, kelebihan akan dikembalikan
- **Catat Donatur**: Semua donasi tercatat lengkap dengan nama pemberi dan kampanye tujuan
- **Filter Donasi**: Bisa melihat donasi berdasarkan nama donatur atau kampanye tertentu
- **Urutkan Donasi**: Bisa mengurutkan donasi dari yang terkecil atau terbesar

### Laporan dan Analisis
- **Prediksi Target**: Fitur khusus admin untuk memperkirakan kapan target tercapai
- **Grafik Progress**: Grafik batang sederhana untuk melihat perkembangan kampanye
- **Total Donasi**: Hitung otomatis total dana yang terkumpul
- **Laporan Lengkap**: Statistik donasi per orang dan per kampanye


## Cara Menjalankan Program

1. Pastikan komputer sudah ada program Go (Golang) versi 1.16 atau lebih baru
2. Download atau salin file program ke komputer
3. Buka terminal (command prompt) dan masuk ke folder program
4. Ketik perintah berikut untuk menjalankan:
```bash
go run crowdfunding.go
```
5. Program akan langsung menampilkan menu utama sistem galang dana

## Struktur Data

### Jenis-jenis Data yang Disimpan
```go
type Pengguna struct {
    Email    string  // Alamat email pengguna
    Username string  // Nama pengguna
    Password string  // Kata sandi
    Peran    string  // "admin" atau "donatur"
}

type Kampanye struct {
    Id        int     // Nomor ID kampanye
    Judul     string  // Judul kampanye
    Kategori  string  // Kategori (pendidikan, kesehatan, dll)
    Deskripsi string  // Penjelasan kampanye
    Target    int     // Target dana dalam Rupiah
    Terkumpul int     // Dana yang sudah terkumpul
    Progress  int     // Persentase tercapai (0-100)
    Status    string  // "aktif" atau "selesai"
}

type Donasi struct {
    KampanyeId  int     // ID kampanye tujuan
    NamaDonatur string  // Nama pemberi donasi
    Jumlah      int     // Jumlah donasi dalam Rupiah
}
```

### Batas Maksimum Sistem
```go
const maxPengguna int = 100    // Maksimal 100 pengguna terdaftar
const maxKampanye int = 100    // Maksimal 100 kampanye aktif
const maxDonasi int = 1000     // Maksimal 1000 transaksi donasi
```

## Cara Kerja Program

### Cara Mencari Data
- **Pencarian Cepat**: Mencari ID kampanye pakai cara pencarian biner (lebih cepat)
- **Pencarian Biasa**: Mencari nama donatur dan judul kampanye satu per satu

### Cara Mengurutkan Data
- **Urutkan Kampanye**: 
  - Bisa urutkan dari progress paling kecil ke besar
  - Bisa urutkan dari progress paling besar ke kecil
- **Urutkan Donasi**: 
  - Bisa urutkan dari donasi paling kecil ke besar  
  - Bisa urutkan dari donasi paling besar ke kecil

### Keamanan dan Validasi
- **Cek Email**: Pastikan format email benar dan pakai domain yang diizinkan (@gmail.com, @yahoo.com, @outlook.com)
- **Kata Sandi Kuat**: Minimal 8 karakter dengan campuran huruf besar, kecil, angka, dan simbol khusus
- **Nama Pengguna Unik**: Tidak boleh ada nama pengguna yang sama
- **Cegah Data Ganda**: Email dan nama pengguna harus unik untuk mencegah duplikasi
- **Validasi Input**: Semua input dicek untuk mencegah error dan memastikan data benar

### Fitur Tambahan
- **Grafik Progress**: Membuat grafik batang sederhana pakai karakter ASCII untuk menunjukkan progress
- **Perlindungan Donasi**: Cek donasi berlebihan dan kembalikan otomatis jika melebihi target
- **Update Status**: Otomatis ubah status kampanye berdasarkan pencapaian target

## Menu & Pilihan

### Menu yang Muncul (Berubah Sesuai Status)
Menu yang ditampilkan berubah-ubah tergantung apakah sudah login atau belum dan jenis pengguna:

**Untuk Pengunjung (Belum Login):**
1. **Daftar (1)** - Buat akun baru dengan validasi lengkap
2. **Masuk (2)** - Login ke sistem pakai email/nama pengguna
3. **Lihat Donasi (3)** - Lihat semua transaksi donasi (hanya baca)
4. **Lihat Kampanye (4)** - Lihat dan baca detail kampanye
5. **Berdonasi (5)** - Berikan donasi (harus login dulu)

**Untuk Donatur (Setelah Login):**
3. **Lihat Donasi (3)** - Tampilkan semua transaksi donasi dengan opsi filter
4. **Lihat Kampanye (4)** - Lihat dan baca kampanye
5. **Berdonasi (5)** - Berikan donasi ke kampanye yang masih aktif
7. **Keluar (7)** - Logout dari sistem

**Untuk Admin (Setelah Login):**
3. **Lihat Donasi (3)** - Tampilkan semua transaksi donasi dengan opsi filter
4. **Lihat Kampanye (4)** - Lihat dan baca kampanye
5. **Berdonasi (5)** - Berikan donasi ke kampanye yang masih aktif
6. **Buat Kampanye (6)** - Membuat kampanye baru (khusus admin)
7. **Keluar (7)** - Logout dari sistem
8. **Prediksi Target (8)** - Analisis dan prediksi (khusus admin)

**Pilihan Global:**
- **Ketik '-1'** - Keluar dari program

### Cara Mencari dan Menyaring

#### Cari Kampanye
- **Berdasarkan ID**: Cari kampanye pakai nomor ID dengan pencarian cepat
- **Berdasarkan Judul**: Cari kampanye pakai kata-kata dalam judul (pencarian fleksibel)

#### Saring Donasi
- **Berdasarkan Nama Donatur**: Tampilkan donasi dari orang tertentu dengan opsi urutan
- **Berdasarkan ID Kampanye**: Tampilkan donasi untuk kampanye tertentu dengan batas tampilan

### Pilihan Urutan
- **Kampanye**: Urutkan berdasarkan persentase progress (kecil ke besar/besar ke kecil)
- **Donasi**: Urutkan berdasarkan jumlah donasi (kecil ke besar/besar ke kecil)

## Fitur Analisis & Prediksi

### Prediksi Pencapaian Target (Khusus Admin)
Fitur khusus untuk admin yang menganalisis kemungkinan tercapainya target kampanye:

**Yang Dianalisis:**
- Rata-rata donasi per kampanye berdasarkan riwayat transaksi
- Sisa target yang belum tercapai
- Perkiraan jumlah transaksi yang dibutuhkan untuk mencapai target
- Hanya menganalisis kampanye dengan status "aktif"

**Rumus Perhitungan:**
```
Rata-rata Donasi = Total Donasi Kampanye / Jumlah Transaksi
Sisa Target = Target Kampanye - Dana Terkumpul
Perkiraan Transaksi = pembulatan ke atas(Sisa Target / Rata-rata Donasi)
```

**Hasil Prediksi:**
- Nama kampanye yang dianalisis
- Sisa target dana dalam Rupiah
- Rata-rata donasi per transaksi
- Perkiraan jumlah transaksi untuk mencapai target

### Tampilan Progress Visual
Sistem menggunakan grafik batang sederhana untuk menampilkan pencapaian target:
```
Progress: 75% [███████████████░░░░░]
Progress: 50% [██████████░░░░░░░░░░]
Progress: 100% [████████████████████]
```

Simbol yang digunakan:
- `█` (Kotak penuh) untuk bagian yang sudah tercapai
- `░` (Kotak kosong) untuk bagian yang belum tercapai
- Lebar total: 20 karakter untuk tampilan yang proporsional

## Pengelolaan Data & Contoh Data

### Akun Default untuk Mencoba Program
Sistem sudah dilengkapi dengan 3 akun siap pakai untuk kemudahan mencoba:

1. **Akun Admin**
   - Email: `ammar@gmail.com`
   - Nama Pengguna: `ammar`
   - Kata Sandi: `Ammar1234@`
   - Peran: `admin`

2. **Akun Donatur 1**
   - Email: `abrar@yahoo.com`
   - Nama Pengguna: `abrar`
   - Kata Sandi: `Abrar1234@`
   - Peran: `donatur`

3. **Akun Donatur 2**
   - Email: `annisa@yahoo.com`
   - Nama Pengguna: `annisa`
   - Kata Sandi: `Annisa1234@`
   - Peran: `donatur`

### Contoh Data Kampanye
Program menyediakan 10 kampanye contoh dengan berbagai kategori:

1. **Bantuan Pendidikan Anak Yatim** (ID: 100) - Pendidikan - 50% progress
2. **Renovasi Masjid Al-Ikhlas** (ID: 110) - Religi - 75% progress
3. **Bantuan Korban Bencana Alam** (ID: 120) - Bencana - 75% progress
4. **Kampus Bersih dan Hijau** (ID: 130) - Lingkungan - 100% (selesai)
5. **Beasiswa Mahasiswa Berprestasi** (ID: 140) - Pendidikan - 50% progress
6. **Posyandu Sehat Mandiri** (ID: 150) - Kesehatan - 50% progress
7. **Gerakan Literasi Digital** (ID: 160) - Teknologi - 30% progress
8. **Pembangunan Perpustakaan Desa** (ID: 170) - Pendidikan - 80% progress
9. **Festival Seni Budaya Lokal** (ID: 180) - Budaya - 75% progress
10. **Pemberdayaan UMKM Perempuan** (ID: 190) - Ekonomi - 30% progress

### Contoh Data Donasi
10 transaksi donasi contoh dari berbagai donatur ke kampanye yang tersedia, dengan nilai donasi mulai dari Rp 500.000 hingga Rp 2.500.000.

## Fitur Keamanan & Validasi

### Validasi Input
- **Format Email**: Cek format email dan hanya izinkan domain tertentu (@gmail.com, @yahoo.com, @outlook.com)
- **Nama Pengguna**: Cek keunikan dan format nama pengguna yang tidak boleh ada spasi
- **Persyaratan Kata Sandi**: Minimal 8 karakter dengan campuran:
  - Minimal 1 huruf besar (A-Z)
  - Minimal 1 huruf kecil (a-z)  
  - Minimal 1 angka (0-9)
  - Minimal 1 simbol khusus (@, #, $, %, &, *, dll)

### Kontrol Akses Berdasarkan Peran
- **Akses Admin**: Akses penuh termasuk membuat kampanye dan analisis
- **Akses Donatur**: Bisa berdonasi dan melihat informasi kampanye
- **Akses Pengunjung**: Hanya bisa melihat informasi (tidak bisa edit)

### Perlindungan Data
- **Cegah Data Ganda**: Validasi email dan nama pengguna unik untuk mencegah duplikasi
- **Validasi Kampanye**: Cek status kampanye sebelum menerima donasi
- **Perlindungan Overflow**: Pengembalian otomatis jika donasi melebihi target

### Manajemen Sesi
- **Status Login**: Melacak status login pengguna
- **Fungsi Logout**: Bersihkan data sesi saat logout
- **Kontrol Akses**: Validasi peran pengguna untuk setiap tindakan

## Antarmuka Pengguna & Pengalaman

### Command Line Interaktif
- **Menu Dinamis**: Menu berubah sesuai status login dan peran pengguna
- **Navigasi Jelas**: Petunjuk yang jelas untuk setiap input yang diminta
- **Pesan Error**: Pesan error yang informatif untuk membantu pengguna
- **Konfirmasi**: Konfirmasi untuk tindakan penting seperti donasi dan pembuatan kampanye

### Elemen Visual
- **Grafik Batang**: Visualisasi sederhana untuk progress kampanye
- **Output Terformat**: Tampilan data yang rapi dan mudah dibaca
- **Format Mata Uang**: Format Rupiah (Rp) untuk semua nilai uang
- **Indikator Status**: Indikator jelas untuk status kampanye (aktif/selesai)

### Panduan Pengguna
- **Teks Bantuan**: Petunjuk jelas untuk setiap menu dan input
- **Validasi Input**: Validasi langsung dengan pesan error yang membantu
- **Opsi Keluar**: Berbagai cara untuk kembali ke menu atau keluar aplikasi

## Gambaran Fungsional

### Fungsi Inti (33 Total Fungsi)

#### Autentikasi & Manajemen Pengguna (7 fungsi)
- `daftar()` - Registrasi pengguna dengan validasi komprehensif
- `verikasiEmail()` - Validasi format dan domain email  
- `verikasiUsername()` - Pengecekan keunikan dan format nama pengguna
- `verikasiPassword()` - Validasi kekuatan kata sandi
- `verikasiPeran()` - Validasi pemilihan peran
- `masuk()` - Fungsi login pengguna
- `logOut()` - Pengakhiran sesi

#### Manajemen Kampanye (8 fungsi)
- `buatKampanye()` - Pembuatan kampanye (khusus admin)
- `tampilkanKampanye()` - Tampilkan kampanye dengan pengurutan
- `detailKampanye()` - Detail kampanye dan pencarian
- `tampilkanDetailKampanye()` - Tampilan detail kampanye tunggal
- `findIdKampanye()` - Pencarian biner untuk ID kampanye
- `findJudulKampanye()` - Pencarian fuzzy untuk judul kampanye
- `containsChar()` - Utilitas pencocokan karakter
- `checkKampanyeAktif()` - Validasi status kampanye aktif

#### Sistem Donasi (8 fungsi)
- `tambahDonasi()` - Proses donasi dengan validasi
- `tampilkanDonasi()` - Tampilkan donasi dengan filtering
- `findNamaDonasi()` - Filter donasi berdasarkan nama donatur
- `findIdDonasi()` - Filter donasi berdasarkan ID kampanye
- `totalDonasi()` - Hitung total donasi
- `totalDonasiDonatur()` - Hitung donasi oleh donatur spesifik
- `tampilNamaDonatur()` - Tampilkan informasi spesifik donatur
- `tampilKampanyeDonatur()` - Tampilkan donasi spesifik kampanye

#### Algoritma Pengurutan (4 fungsi)
- `insertionSortAsc()` - Insertion sort naik untuk donasi
- `insertionSortDesc()` - Insertion sort turun untuk donasi
- `sortSelectionAsc()` - Selection sort naik untuk kampanye
- `sortSelectionDesc()` - Selection sort turun untuk kampanye

#### Analisis & Utilitas (6 fungsi)
- `prediksiPencapaianTarget()` - Prediksi pencapaian target (khusus admin)
- `createProgressBar()` - Generator progress bar ASCII
- `findJudul()` - Dapatkan judul kampanye berdasarkan ID
- `menuUtama()` - Tampilan menu utama dinamis
- `dummyDataKampanye()` - Inisialisasi data kampanye contoh
- `dummyDataDonasi()` - Inisialisasi data donasi contoh

#### Program Utama (1 fungsi)
- `main()` - Loop program utama dan inisialisasi

## Spesifikasi Teknis

### Library yang Digunakan
```go
import (
    "fmt"   // Untuk input/output dan format tampilan
    "math"  // Untuk operasi matematika (fungsi pembulatan ke atas)
)
```

### Konstanta
```go
const maxPengguna int = 100    // Maksimal pengguna
const maxKampanye int = 100    // Maksimal kampanye  
const maxDonasi int = 1000     // Maksimal donasi
```

### Jenis Array
```go
type tabPengguna [maxPengguna]Pengguna
type tabKampanye [maxKampanye]Kampanye
type tabDonasi [maxDonasi]Donasi
```

## Keterbatasan & Pertimbangan

### Keterbatasan Saat Ini
- **Penyimpanan Data**: Data tidak tersimpan permanen (hilang setelah program ditutup)
- **Antarmuka**: Hanya antarmuka command line, tidak ada tampilan grafis
- **Bahasa**: Antarmuka menggunakan Bahasa Indonesia
- **Platform**: Bisa jalan di Linux, Windows, macOS dengan runtime Go
- **Concurrency**: Hanya satu pengguna dalam satu waktu

### Pertimbangan Performa
- **Efisiensi Pencarian**: Pencarian biner untuk ID (cepat), pencarian berurutan untuk nama (lambat)
- **Performa Pengurutan**: Selection sort dan Insertion sort (tidak efisien untuk data besar)
- **Penggunaan Memori**: Alokasi array tetap sesuai batas maksimum
- **Skalabilitas**: Terbatas oleh konstanta maksimum yang telah ditentukan

### Saran Perbaikan Masa Depan
- Integrasi database untuk penyimpanan data permanen
- Antarmuka berbasis web untuk pengalaman pengguna yang lebih baik
- Dukungan multi-bahasa (Inggris/Indonesia)
- Notifikasi real-time dan integrasi email
- Fitur analisis dan pelaporan lanjutan
- Dukungan multi-mata uang
---

## Panduan Memulai

### Cara Cepat Memulai
1. Jalankan program: `go run crowdfunding.go`
2. Login sebagai admin: email `ammar@gmail.com`, kata sandi `Ammar1234@`
3. Atau login sebagai donatur: email `abrar@yahoo.com`, kata sandi `Abrar1234@`
4. Jelajahi fitur-fitur yang tersedia sesuai dengan peran pengguna

### Skenario Pengujian
- **Tes Pendaftaran**: Daftar pengguna baru dengan berbagai validasi
- **Tes Donasi**: Lakukan donasi ke berbagai kampanye
- **Tes Analisis**: Login sebagai admin dan gunakan fitur prediksi
- **Tes Pencarian**: Cari kampanye berdasarkan ID atau judul
- **Tes Pengurutan**: Urutkan kampanye dan donasi dengan berbagai kriteria

*Dikembangkan untuk Tugas Besar Algoritma Pemrograman - Kelompok RendangKeju*
