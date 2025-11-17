# API Specification v1 (feez-go-api)

Semua endpoint berawal di `/api/v1`.

## 1. Authentication

Endpoint untuk mendaftarkan dan memvalidasi user.

| Fitur | Method | Endpoint | Body (JSON Request) | Respons Sukses (200/201) | Respons Gagal (400/401) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Signup** | `POST` | `/signup` | `{ "email": "...", "password": "...", "username": "..." }` | `{ "message": "Signup successful" }` | `{ "error": "Email already exists" }` |
| **Login** | `POST` | `/login` | `{ "email": "...", "password": "..." }` | `{ "token": "jwt.token.string" }` | `{ "error": "Invalid credentials" }` |