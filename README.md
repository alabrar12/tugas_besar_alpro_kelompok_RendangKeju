# Sistem Crowdfunding

Program sederhana untuk mengelola kampanye crowdfunding yang dibuat menggunakan bahasa Go.

## Apa itu Program Crowdfunding?

Program ini adalah sistem crowdfunding:
- Seorang **Admin** bisa membuat kampanye untuk mengumpulkan dana
- Seorang **Donatur** bisa memberikan donasi ke kampanye yang ada
- Semua orang bisa melihat daftar kampanye dan donasi

## Cara Menjalankan Program

### 1. Install Go
Pastikan Go sudah terinstall di komputer kamu. Cek dengan mengetik:
```bash
go version
```

### 2. Download Program
Download file `crowdfunding.go` ke komputer kamu.

### 3. Jalankan Program
Buka terminal/command prompt, masuk ke folder tempat file `crowdfunding.go` berada, lalu ketik:
```bash
go run crowdfunding.go
```

## Cara Menggunakan Program

### Menu Utama
Setelah program berjalan, kamu akan melihat menu dengan pilihan angka. Ketik angka sesuai menu yang ingin kamu pilih.

### Akun yang Sudah Ada
Program sudah punya 2 akun untuk testing:

**Admin:**
- Email: `ammar@gmail.com`
- Username: `ammar`
- Password: `Ammar1234@`

**Donatur:**
- Email: `ghifari@yahoo.com`
- Username: `ghifari`
- Password: `Ghifari1234@`

### Fitur-Fitur Program

#### 1. Daftar Akun Baru
- Pilih menu `1`
- Masukkan email (harus pakai @gmail.com, @yahoo.com, atau @outlook.com)
- Masukkan username (4-50 karakter)
- Masukkan password (minimal 8 karakter, harus ada huruf besar, kecil, angka, dan simbol)
- Pilih peran: `admin` atau `donatur`

#### 2. Login
- Pilih menu `2`
- Masukkan email/username dan password
- Jika berhasil, kamu akan masuk ke sistem

#### 3. Lihat Donasi
- Pilih menu `3`
- Bisa dilihat oleh siapa saja (tidak perlu login)
- Menampilkan semua donasi yang sudah dilakukan

#### 4. Lihat Kampanye
- Pilih menu `4`
- Bisa dilihat oleh siapa saja
- Menampilkan daftar kampanye dengan detail lengkap
- Bisa cari berdasarkan ID, judul, atau kategori

#### 5. Berdonasi
- Pilih menu `5`
- Harus login terlebih dahulu
- Pilih kampanye yang masih aktif
- Masukkan jumlah donasi

#### 6. Buat Kampanye (Khusus Admin)
- Pilih menu `6`
- Hanya bisa dilakukan oleh admin
- Masukkan judul, kategori, deskripsi, dan target dana

#### 7. Logout
- Pilih menu `7`
- Keluar dari akun yang sedang login

#### 8. Keluar Program
- Ketik `-1`
- Program akan berhenti

## Aturan-Aturan Penting

### Email
- Harus 12-50 karakter
- Harus pakai domain: @gmail.com, @yahoo.com, atau @outlook.com
- Tidak boleh sama dengan email yang sudah ada

### Username
- Harus 4-50 karakter
- Tidak boleh sama dengan username yang sudah ada

### Password
- Minimal 8 karakter
- Harus ada huruf besar (A-Z)
- Harus ada huruf kecil (a-z)
- Harus ada angka (0-9)
- Harus ada simbol (@, #, $, %, &)

### Donasi
- Hanya bisa donasi ke kampanye yang masih aktif
- Jika donasi melebihi target, kelebihan akan dikembalikan
- Kampanye otomatis selesai jika target tercapai

## Fitur Tambahan

### Progress Bar
Setiap kampanye punya progress bar yang menunjukkan seberapa banyak dana yang sudah terkumpul.

### Sorting (Pengurutan)
- Kampanye bisa diurutkan berdasarkan progress (naik/turun)
- Donasi bisa diurutkan berdasarkan jumlah (besar ke kecil atau sebaliknya)

### Pencarian
- Cari kampanye berdasarkan ID (menggunakan binary search)
- Cari kampanye berdasarkan judul (pencarian fleksibel)
- Cari kampanye berdasarkan kategori

## Contoh Penggunaan

1. **Jalankan program**
2. **Login sebagai admin** (pakai akun ammar)
3. **Buat kampanye baru** (pilih menu 6)
4. **Logout** (pilih menu 7)
5. **Login sebagai donatur** (pakai akun ghifari)
6. **Lihat kampanye** (pilih menu 4)
7. **Berdonasi** (pilih menu 5)
8. **Lihat hasil donasi** (pilih menu 3)

## Batasan Program

- Maksimal 100 pengguna
- Maksimal 100 kampanye
- Maksimal 1000 donasi

## Tips

- Ketik `keluar` saat daftar akun untuk membatalkan
- Semua input harus sesuai aturan yang diberikan
- Pastikan tidak ada spasi di awal atau akhir input
- Program case-sensitive (huruf besar/kecil berpengaruh)

## Troubleshooting

**Program tidak jalan?**
- Pastikan Go sudah terinstall
- Pastikan file `crowdfunding.go` ada di folder yang benar

**Tidak bisa login?**
- Cek email/username dan password
- Pastikan akun sudah terdaftar
- Coba pakai akun testing yang sudah disediakan

**Input tidak diterima?**
- Pastikan input sesuai aturan (panjang karakter, format, dll)
- Jangan pakai spasi di awal atau akhir

---

**Dibuat dengan ❤️ menggunakan Go**
