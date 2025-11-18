# Frontend Architecture (feez-react-client)

Aplikasi React berbasis Vite dengan Chakra UI untuk styling.

## Tech Stack
* **Framework:** React 18 (Vite)
* **Styling:** Chakra UI v2
* **Routing:** React Router DOM v6
* **State Management:** React `useState` & `useEffect` (Simple local state).

## Struktur Folder (`client/src`)

* **`components/`**: Komponen UI yang *reusable*.
    * `Navbar.jsx`: Navigasi utama & tombol login/cart.
    * `ProductCard.jsx`: Menampilkan item produk & logika "Add to Cart".
* **`pages/`**: Halaman utama (Views).
    * `HomePage.jsx`: Katalog produk (`GET /products`).
    * `LoginPage.jsx`: Form login & simpan JWT (`POST /login`).
    * `CartPage.jsx`: Keranjang & Checkout (`GET /cart`, `POST /checkout`).
* **`App.jsx`**: Pengatur *Routing* utama.
* **`main.jsx`**: Entry point, setup `ChakraProvider` & `BrowserRouter`.
* **`theme.js`**: Kustomisasi tema Chakra (warna brand Supabase).

## Authentication Flow

1.  **Login:** User login via `LoginPage`.
2.  **Storage:** Token JWT disimpan di browser `localStorage` dengan key `token`.
3.  **Protected Actions:**
    * Saat user klik "Add to Cart" atau buka "Cart Page", aplikasi mengecek keberadaan token.
    * Token dikirim via Header: `Authorization: Bearer <token>`.
    * Jika token tidak ada, user diarahkan ke Login atau diberi notifikasi.