# API Specification v1 (feez-go-api)

Semua endpoint berawal di `/api/v1`.

---

## 1. Authentication (Otentikasi)
Endpoint publik untuk registrasi dan login.

| Fitur | Method | Endpoint | Auth? | Body (JSON Request) | Respons Sukses (200/201) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| User Signup | `POST` | `/signup` | No | `{ "email": "...", "password": "...", "username": "..." }` | `{ "message": "Signup successful" }` |
| User Login | `POST` | `/login` | No | `{ "email": "...", "password": "..." }` | `{ "token": "jwt.token.string" }` |

---

## 2. Products (Produk)
Endpoint publik untuk melihat katalog produk.

| Fitur | Method | Endpoint | Auth? | Body (JSON Request) | Respons Sukses (200) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Get All Products | `GET` | `/products` | No | (N/A) | `[ { "id": 1, "name": "...", "price": 150000, ... }, ... ]` |

---

## 3. Cart (Keranjang)
Endpoint *private* (butuh token) untuk mengelola keranjang belanja user.

| Fitur | Method | Endpoint | Auth? | Body (JSON Request) | Respons Sukses (200/201) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Add to Cart | `POST` | `/cart` | **Yes (Bearer)** | `{ "product_id": 1, "quantity": 2 }` | `{ "message": "Item added to cart" }` |
| Get Cart | `GET` | `/cart` | **Yes (Bearer)** | (N/A) | `[ { "product_id": 1, "product_name": "...", ... }, ... ]` |

---

## 4. Order (Pesanan)
Endpoint *private* (butuh token) untuk checkout.

| Fitur | Method | Endpoint | Auth? | Body (JSON Request) | Respons Sukses (200) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Checkout | `POST` | `/checkout` | **Yes (Bearer)** | (N/A) | `{ "order_id": 1 }` |