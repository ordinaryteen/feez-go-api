# feez-go-api

Backend API (Go) untuk *capstone project* e-commerce "feez 2.0".

---

## 🚀 Cara Menjalankan Proyek (Lokal)

### 1. Prasyarat

* Go (versi 1.22+)
* Akun Supabase

### 2. Setup Database

1.  Bikin proyek baru di Supabase.
2.  Pergi ke **SQL Editor** dan "Run" *full script* SQL di `docs/schema.sql` untuk *setup* semua tabel, *trigger*, dan *function*.
    *(Catatan: Kita belum bikin file ini, tapi nanti kita akan pindahin 'script raksasa' kita ke sini).*

### 3. Setup Environment (.env)

1.  Di Supabase, pergi ke **Settings > Database > Connection pooler** (pake "Session" mode).
2.  Bikin *file* `.env` di *root* proyek.
3.  Isi dengan 5 variabel koneksi dari *pooler* dan 1 kunci JWT:

    ```env
    DB_HOST="aws-..."
    DB_PASS="[YOUR_PASSWORD]"
    DB_PORT="5432"
    DB_NAME="postgres"
    DB_USER="postgres.rjg..."

    JWT_SECRET_KEY="[YOUR_SECRET_KEY_YANG_PANJANG_DAN_ACAK]"
    ```

### 4. Instalasi & Run

1.  Clone repo ini.
2.  Install semua *dependencies* (library):
    ```bash
    go mod tidy
    ```
3.  Jalankan server:
    ```bash
    go run ./cmd/api/main.go
    ```
    Server akan otomatis nyambung ke Supabase dan jalan di `http://localhost:8080`.

---

## 📚 Dokumentasi API

Lihat dokumentasi *endpoint* lengkap (Signup, Login, Products, dll) di sini:

**[-> Dokumentasi API v1 (docs/api_v1.md)](./docs/api_v1.md)**