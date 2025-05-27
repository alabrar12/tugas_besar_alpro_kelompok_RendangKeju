# Sistem Crowdfunding

Aplikasi sistem crowdfunding sederhana yang dibangun menggunakan bahasa pemrograman Go. Sistem ini memungkinkan pengguna untuk membuat kampanye crowdfunding, berdonasi, dan melacak progress pencapaian target dana.

## 🚀 Fitur Utama

### Autentikasi & Manajemen User
- **Registrasi Pengguna**: Validasi email dengan domain tertentu (@gmail.com, @yahoo.com, @outlook.com)
- **Login System**: Login menggunakan email atau username
- **Role-based Access**: Sistem peran admin dan donatur dengan akses yang berbeda
- **Password Security**: Validasi password dengan requirement huruf besar, kecil, angka, dan simbol khusus

### Manajemen Kampanye
- **Pembuatan Kampanye**: Hanya admin yang dapat membuat kampanye baru
- **Pencarian Kampanye**: Fitur pencarian berdasarkan ID, judul, atau kategori
- **Detail Kampanye**: Tampilan detail lengkap dengan progress bar visual
- **Status Tracking**: Otomatis mengubah status kampanye menjadi "selesai" ketika target tercapai

### Sistem Donasi
- **Donasi Real-time**: Pengguna dapat berdonasi ke kampanye aktif
- **Overflow Protection**: Sistem mengembalikan kelebihan donasi jika melebihi target
- **Tracking Donatur**: Pencatatan lengkap semua donasi dengan nama donatur

### Analytics & Reporting
- **Prediksi Target**: Fitur khusus admin untuk memprediksi pencapaian target berdasarkan rata-rata donasi
- **Sorting & Filtering**: Multiple opsi sorting untuk kampanye dan donasi
- **Progress Visualization**: Progress bar ASCII untuk visualisasi pencapaian target


## 🚦 Cara Menjalankan

1. Pastikan Go terinstall di sistem
2. Clone repository
3. Jalankan program:
```bash
go run crowdfunding.go
```

## 🏗️ Struktur Data

### Tipe Bentukan (Struct)
```go
type Pengguna struct {
    Email    string
    Username string
    Password string
    Peran    string  // "admin" atau "donatur"
}

type Kampanye struct {
    Id        int
    Judul     string
    Kategori  string
    Deskripsi string
    Target    int
    Terkumpul int
    Progress  int
    Status    string  // "aktif" atau "selesai"
}

type Donasi struct {
    KampanyeId  int
    NamaDonatur string
    Jumlah      int
}
```

### Batasan Sistem
- Maximum 100 pengguna
- Maximum 100 kampanye
- Maximum 1000 transaksi donasi

## 🔧 Algoritma & Implementasi

### Algoritma Pencarian
- **Binary Search**: Implementasi pada pencarian ID kampanye (fungsi `findIdKampanye`)
- **Sequential Search**: Pencarian berdasarkan kategori dan nama donatur

### Algoritma Sorting
- **Selection Sort**: Untuk sorting kampanye berdasarkan progress (ascending/descending)
- **Insertion Sort**: Untuk sorting donasi berdasarkan jumlah (ascending/descending)

### Validasi & Security
- **Email Validation**: Validasi format dan domain email yang diizinkan
- **Password Strength**: Minimum 8 karakter dengan kombinasi huruf besar, kecil, angka, dan simbol
- **Duplicate Prevention**: Validasi email dan username unik

## 🎮 Menu & Navigasi

### Menu Utama
1. **Daftar** - Registrasi pengguna baru
2. **Masuk** - Login ke sistem
3. **Lihat Donasi** - Tampilkan semua transaksi donasi
4. **Lihat Kampanye** - Browse dan detail kampanye
5. **Berdonasi** - Melakukan donasi ke kampanye
6. **Buat Kampanye** - Membuat kampanye baru (admin only)
7. **Log Out** - Keluar dari sistem
8. **Prediksi Pencapaian Target** - Analytics untuk admin

### Fitur Pencarian
- **ID Kampanye**: Pencarian cepat dengan binary search
- **Judul Kampanye**: Fuzzy matching untuk fleksibilitas
- **Kategori**: Pencarian exact match berdasarkan kategori

## 🔍 Fitur Analytics

### Prediksi Pencapaian Target
Fitur khusus untuk admin yang menganalisis:
- Rata-rata donasi per kampanye
- Estimasi jumlah transaksi yang dibutuhkan untuk mencapai target
- Hanya untuk kampanye dengan status "aktif"

Formula yang digunakan:
```
Estimasi Transaksi = ceil(Sisa Target / Rata-rata Donasi)
```

## 🎨 User Experience

### Progress Visualization
Sistem menggunakan progress bar ASCII untuk menampilkan pencapaian target:
```
Progress: 75% [███████████████░░░░░]
```

### Sorting Options
- **Kampanye**: Sort berdasarkan progress (ascending/descending)
- **Donasi**: Sort berdasarkan jumlah donasi (ascending/descending)

## 💾 Data Management

### Default Users
Sistem dilengkapi dengan 2 akun default untuk testing:
- **Admin**: ammar@gmail.com / ammar (Password: Ammar1234@)
- **Donatur**: ghifari@yahoo.com / ghifari (Password: Ghifari1234@)

## 🔐 Security Features

- **Input Validation**: Validasi komprehensif untuk semua input user
- **Role-based Access Control**: Pembatasan akses berdasarkan peran pengguna
- **Password Requirements**: Enforcement password yang kuat
- **Duplicate Prevention**: Pencegahan data duplikat

## 📋 Limitations

- Data tidak persisten (hilang setelah program ditutup)
- Command-line interface only
- Bahasa Indonesia only untuk UI

---

*Dibuat untuk keperluan Tugas Besar Algoritma Pemrograman*
