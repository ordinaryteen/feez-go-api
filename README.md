# aney-ecommerce

A full-stack e-commerce capstone project.
- **Backend:** Go (Golang), Chi Router, PostgreSQL (Supabase).
- **Frontend:** React (Vite), Chakra UI.
- **Database:** Supabase (Auth, Database, Realtime).

---

## Struktur Proyek (Monorepo)

* **`api/`**: Backend server code (Go).
* **`client/`**: Frontend client code (React).
* **`docs/`**: Dokumentasi teknis & spesifikasi.

---

## Cara Menjalankan (Development)

Anda perlu membuka **2 Terminal** untuk menjalankan Fullstack app ini.

### Terminal 1: Backend (API)

1.  Masuk ke folder API:
    ```bash
    cd api
    ```
2.  Pastikan file `.env` sudah ada di dalam folder `api/` dengan konfigurasi database yang benar.
3.  Jalankan server:
    ```bash
    go run ./cmd/api/main.go
    ```
    *Server akan berjalan di `http://localhost:8080`*

### Terminal 2: Frontend (Client)

1.  Masuk ke folder Client:
    ```bash
    cd client
    ```
2.  Install dependencies (jika belum):
    ```bash
    npm install
    ```
3.  Jalankan server dev:
    ```bash
    npm run dev
    ```
    *Website akan berjalan di `http://localhost:5173`*

---

## Dokumentasi Detail

* **[API Specification (v1)](./docs/api_v1.md)** - Daftar endpoint dan format JSON.
* **[Frontend Architecture](./docs/frontend.md)** - Struktur komponen dan alur UI.