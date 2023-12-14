
# Transaction Demo

## Deskripsi
Transaction Demo adalah aplikasi demo yang dikembangkan menggunakan bahasa pemrograman Golang. Aplikasi ini menyediakan berbagai fitur, termasuk registrasi pelanggan dan pedagang, sistem otentikasi login yang aman, dan penanganan transaksi.

## Fitur Utama
- Registrasi Pelanggan dan Pedagang
- Sistem Otentikasi Login yang Aman
- Penanganan Transaksi

## Prasyarat
Pastikan Anda telah menginstal Golang dan MySQL di mesin Anda sebelum menjalankan aplikasi ini.

## Instalasi
1. Clone repositori ini: `git clone https://github.com/RaswanSualdi/transaction_demo.git`
2. Pindah ke direktori proyek: `cd transaction_demo`
3. Salin file `.env.example` menjadi `.env` dan konfigurasikan sesuai kebutuhan.
4. Jalankan perintah: `go run main.go` untuk menjalankan aplikasi.

## Kontribusi
Jika Anda ingin berkontribusi pada proyek ini, silakan buat _fork_ dan kirim _pull request_.

## Struktur Proyek
Berikut adalah struktur direktori utama proyek:
```
transaction_demo/
|-- delivery/
|-- models/
|-- payload/
|-- repository/
|-- usecase/
|-- .env.example
|-- main.go
|-- README.md
|-- Transaction_Demo.postman_collection.json
```

## Teknologi yang Digunakan
- Golang
- Gin (Framework Web)
- Gorm (ORM untuk Go)

## Pengujian API dengan Postman

Anda dapat menggunakan Postman untuk menguji API dengan mudah. Silakan unduh koleksi Postman yang telah disiapkan di file `Transaction_Demo.postman_collection.json`.


### Langkah-langkah:
1. Unduh dan instal [Postman](https://www.postman.com/downloads/).
2. Buka Postman.
3. Klik tombol "Run in Postman" di atas atau impor koleksi dengan menggunakan tautan [ini](https://www.getpostman.com/collections).

### Catatan:
Pastikan untuk menyesuaikan tautan "Run in Postman" dan tautan impor koleksi dengan tautan yang sesuai dengan koleksi Postman Anda.

Jika Anda memiliki skenario pengujian atau skrip lain yang ingin Anda sertakan, pastikan untuk menyertakan petunjuk penggunaannya dalam README. Ini akan membantu pengguna dan kontributor untuk lebih memahami fungsionalitas dan cara penggunaan aplikasi Anda.


## Kontak
- **WhatsApp:** [![WhatsApp](https://img.shields.io/badge/-WhatsApp-25D366?logo=whatsapp)](https://wa.me/6282194808091)
- **Instagram:** [![Instagram](https://img.shields.io/badge/-Instagram-E4405F?logo=instagram)](https://instagram.com/raswancodes)
- **GitHub:** [![GitHub](https://img.shields.io/badge/-GitHub-181717?logo=github)](https://github.com/RaswanSualdi/transaction_demo)

## Lisensi
Proyek ini dilisensikan di bawah [nama lisensi] - lihat file [LICENSE.md](LICENSE.md) untuk detailnya.

## Pembuat
Raswan Sualdi - [Kunjungi Profil GitHub](https://github.com/RaswanSualdi)

**Catatan:** Pastikan untuk membaca petunjuk instalasi dan kontribusi yang telah disertakan di README proyek. Proyek ini menunjukkan penggunaan Golang untuk pengembangan aplikasi sederhana dengan fokus pada manajemen transaksi. Silakan berikan masukan atau pertanyaan jika ada.