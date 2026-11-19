# Memahami DDD dan Clean Architecture untuk PayCore

## 1. Sebelum mengenal istilah DDD

Bayangkan kita sedang membuat PayCore, sebuah aplikasi dompet digital.

Pengguna bisa:

- Membuat akun.
- Login.
- Melihat saldo.
- Mengirim uang ke pengguna lain.
- Menerima uang.

Di dalam aplikasi ini ada banyak aturan bisnis. Contohnya:

- Email pengguna tidak boleh sama dengan pengguna lain.
- Pengguna yang dinonaktifkan tidak boleh login.
- Saldo tidak boleh menjadi negatif.
- Setiap transfer harus memiliki catatan uang keluar dan uang masuk.

Aturan-aturan tersebut adalah bagian dari **bisnis PayCore**.

Ketika aplikasi masih kecil, kita mungkin bisa menulis semua kode dalam beberapa file. Namun, ketika fitur semakin banyak, kode akan sulit dikelola.

DDD membantu kita memahami dan memisahkan bagian-bagian bisnis tersebut dengan baik.

---

## 2. Apa itu DDD?

DDD adalah singkatan dari **Domain-Driven Design**.

Mari pecah namanya:

- **Domain**: bidang atau masalah bisnis yang ingin diselesaikan aplikasi.
- **Driven**: sesuatu yang menjadi dasar atau panduan.
- **Design**: cara kita merancang aplikasi.

Jadi, DDD adalah cara merancang aplikasi dengan menjadikan kebutuhan dan aturan bisnis sebagai dasar utama.

Contoh:

PayCore mempunyai aturan bahwa pengguna yang dinonaktifkan tidak boleh melakukan login.

Kita sebaiknya meletakkan aturan tersebut pada bagian kode yang bertanggung jawab terhadap konsep pengguna, bukan menyebarkannya ke berbagai HTTP handler atau query database.

**Intinya:** saat membuat fitur, kita mulai dengan memahami aturan bisnisnya, kemudian menentukan kode yang bertanggung jawab menjalankan aturan tersebut.

---

# 3. Istilah-istilah di dalam DDD

## A. Domain

**Domain adalah masalah atau bidang bisnis yang sedang kita buatkan solusinya.**

Contoh:

Domain PayCore adalah sistem pembayaran dan pengelolaan uang digital.

Di dalamnya terdapat beberapa konsep:

- User
- Wallet
- Transaction
- Ledger
- Payment
- Refund

Masing-masing mempunyai aturan bisnis sendiri.

Contoh aturan pada Wallet:

- Saldo tidak boleh negatif.
- Pengguna hanya boleh menggunakan dompet miliknya.
- Saldo bertambah ketika menerima uang.

Domain bukan nama folder semata. Domain adalah bagian yang berisi konsep dan aturan bisnis aplikasi.

## B. Entity

**Entity adalah objek bisnis yang memiliki identitas.**

Contoh: User.

Misalnya ada dua pengguna:

| ID       | Email       |
| -------- | ----------- |
| user-001 | a@email.com |
| user-002 | b@email.com |

Keduanya merupakan pengguna yang berbeda karena memiliki ID berbeda.

Sekarang bayangkan email user-001 berubah dari `a@email.com` menjadi `new@email.com`.

Apakah pengguna tersebut menjadi pengguna baru?

Tidak. ID-nya masih `user-001`, sehingga dia tetap pengguna yang sama.

Itulah alasan User disebut Entity: identitasnya tetap ada meskipun beberapa datanya berubah.

Contoh kode:

```go
type User struct {
    id           UserID
    email        Email
    passwordHash PasswordHash
    status       UserStatus
}
```

Dalam contoh ini, `User` adalah Entity.

## C. Value Object

**Value Object adalah objek yang nilainya lebih penting daripada identitasnya.**

Contoh: Email.

Misalnya:

```go
email1 := Email("user@email.com")
email2 := Email("user@email.com")
```

Kita tidak perlu membedakan `email1` dan `email2` berdasarkan ID. Keduanya memiliki arti yang sama karena nilainya sama.

Bandingkan dengan User:

- User dikenali berdasarkan ID.
- Email dikenali berdasarkan nilainya.

Contoh lain Value Object:

- Email
- Money
- PhoneNumber
- PasswordHash

Mengapa tidak menggunakan `string` saja?

Karena Value Object bisa memastikan nilai yang digunakan memenuhi aturan tertentu.

Contoh:

```go
email, err := domain.NewEmail("user@email.com")
if err != nil {
    return err
}
```

Kita tidak langsung membuat Email dari sembarang string. Kita menggunakan constructor yang memeriksa validitasnya.

Dengan begitu, aturan tentang format email berada pada satu tempat.

## D. Aggregate

**Aggregate adalah sekumpulan objek bisnis yang harus dijaga agar tetap mengikuti aturan bisnis yang sama.**

Bayangkan sebuah Wallet:

```text
Wallet
├── Wallet ID
├── User ID
├── Balance
└── Currency
```

Wallet memiliki beberapa data yang saling berhubungan.

Misalnya ada aturan:

> Saldo Wallet tidak boleh negatif.

Ketika melakukan pengurangan saldo, kita harus memastikan aturan tersebut tetap terpenuhi.

Aggregate membantu kita mengelompokkan data dan aturan yang harus dijaga bersama.

Contoh sederhananya:

```go
type Wallet struct {
    id      WalletID
    balance Money
}

func (w *Wallet) Withdraw(amount Money) error {
    // Periksa apakah saldo mencukupi.
    // Kurangi saldo jika aturan terpenuhi.
    return nil
}
```

Jadi, aturan pengurangan saldo diletakkan pada objek Wallet, bukan dibiarkan setiap bagian aplikasi mengubah saldo sesuka hati.

## E. Aggregate Root

**Aggregate Root adalah objek utama yang menjadi pintu masuk untuk mengubah isi sebuah Aggregate.**

Contoh:

```text
Wallet (Aggregate Root)
├── Wallet ID
├── Balance
└── Currency
```

Ketika kode lain ingin mengurangi saldo, kode tersebut memanggil method milik Wallet:

```go
wallet.Withdraw(amount)
```

Bukan mengubah saldo secara langsung:

```go
wallet.balance = 0
```

Mengapa?

Karena Wallet harus memastikan aturan saldo tetap terpenuhi.

Untuk saat ini, anggap saja Aggregate Root sebagai **penjaga aturan untuk sekumpulan objek bisnis**.

Catatan: dalam contoh Wallet di atas, Wallet juga merupakan Entity dan Aggregate Root. Satu objek bisa memiliki lebih dari satu peran.

## F. Repository

**Repository adalah bagian yang bertugas menyimpan dan mengambil objek bisnis.**

Bayangkan kita ingin mengambil data User dari PostgreSQL.

Tanpa memisahkan tanggung jawab, kode bisnis mungkin harus mengetahui cara menjalankan SQL atau menggunakan GORM.

Dengan Repository, kode bisnis cukup meminta:

```go
user, err := userRepository.FindByID(ctx, userID)
```

Repository yang bertanggung jawab mengurus bagaimana data tersebut diambil dari database.

Gambaran sederhana:

```text
Application
    |
    | Meminta data User
    v
UserRepository
    |
    | Mengambil data
    v
PostgreSQL
```

Repository bukan database itu sendiri. Repository adalah cara kode aplikasi berkomunikasi dengan penyimpanan data.

## G. Domain Service

**Domain Service adalah tempat untuk aturan bisnis yang tidak cocok dimiliki oleh satu Entity atau Aggregate tertentu.**

Misalnya, PayCore memiliki aturan perhitungan biaya transfer.

Aturan biaya transfer mungkin melibatkan:

- Jumlah uang yang dikirim.
- Jenis transfer.
- Jenis akun.
- Kebijakan biaya.

Jika aturan tersebut tidak secara alami menjadi tanggung jawab satu Entity, kita bisa menggunakan Domain Service.

Contoh:

```go
type TransferFeeService struct {
}

func (s *TransferFeeService) Calculate(amount Money) Money {
    // Hitung biaya transfer berdasarkan aturan bisnis.
    return Money{}
}
```

Jangan menggunakan Domain Service untuk semua fungsi.

Kalau suatu aturan sudah cocok menjadi tanggung jawab Wallet atau User, letakkan aturan itu pada objek tersebut.

## H. Domain Event

**Domain Event adalah pemberitahuan bahwa sesuatu yang penting dalam bisnis sudah terjadi.**

Contoh:

Pengguna berhasil melakukan transfer uang.

Kita bisa menyatakan bahwa sebuah kejadian telah terjadi:

```text
TransferCompleted
```

Kejadian tersebut dapat digunakan oleh bagian lain untuk melakukan pekerjaan lanjutan.

Contoh:

```text
Transfer berhasil
       |
       v
TransferCompleted
       |
       +----> Kirim notifikasi
       |
       +----> Perbarui laporan
       |
       +----> Catat aktivitas
```

Perhatikan perbedaannya:

- `TransferCompleted` adalah kejadian bisnis.
- Kafka adalah teknologi yang bisa digunakan untuk menyebarkan kejadian tersebut ke service lain.

Domain Event tidak sama dengan Kafka. Domain Event adalah konsep bisnis; Kafka adalah salah satu alat untuk mengirimkan pesan.

## I. Bounded Context

**Bounded Context adalah batas yang menentukan bagian bisnis mana yang memiliki istilah, aturan, dan tanggung jawab tertentu.**

Contoh PayCore:

```text
PayCore
├── Identity Context
├── Wallet Context
├── Transaction Context
└── Ledger Context
```

Setiap bagian memiliki tanggung jawab berbeda.

**Identity Context**

Mengurus:

- Registrasi pengguna.
- Login.
- Password.
- Status akun.

**Wallet Context**

Mengurus:

- Dompet pengguna.
- Saldo.
- Penambahan saldo.
- Pengurangan saldo.

**Transaction Context**

Mengurus:

- Permintaan transfer.
- Status transfer.
- Proses transfer.

**Ledger Context**

Mengurus:

- Catatan debit.
- Catatan kredit.
- Keseimbangan pencatatan transaksi.

Mengapa perlu dipisahkan?

Karena aturan dan tanggung jawabnya berbeda. Kita tidak ingin kode untuk mengurus password bercampur dengan kode untuk menghitung saldo.

Bounded Context juga membantu kita menentukan batas service ketika membangun sistem microservices.

---

# 4. Istilah-istilah Clean Architecture

DDD dan Clean Architecture saling melengkapi, tetapi bukan hal yang sama.

- DDD membantu kita memahami dan menyusun model bisnis.
- Clean Architecture membantu kita mengatur struktur kode dan ketergantungan antarbagian aplikasi.

Sekarang kita lihat istilah Clean Architecture yang digunakan pada proyek PayCore.

## A. Application

**Application adalah bagian yang mengatur langkah-langkah untuk menjalankan sebuah kebutuhan pengguna.**

Contoh: pengguna ingin mendaftar.

Apa saja yang harus dilakukan?

1. Memeriksa email.
2. Memeriksa apakah email sudah digunakan.
3. Mengenkripsi password menjadi password hash.
4. Membuat User.
5. Menyimpan User.

Rangkaian langkah tersebut adalah tanggung jawab Application.

Contoh file:

```text
application/
└── auth/
    └── register.go
```

Di dalamnya terdapat proses `RegisterUser`.

Application mengatur urutan pekerjaan, tetapi bukan tempat untuk menulis seluruh aturan bisnis milik User atau Wallet.

## B. Use Case

**Use Case adalah satu tindakan atau kebutuhan yang bisa dilakukan oleh pengguna atau sistem.**

Contoh pada PayCore:

- Register User
- Login User
- Create Wallet
- Get Wallet Balance
- Transfer Money
- Refund Payment

Satu Use Case biasanya mewakili satu tujuan yang jelas.

Contoh:

`RegisterUser` berarti proses untuk mendaftarkan pengguna baru.

File:

```text
application/
└── auth/
    ├── register.go
    └── login.go
```

Jadi, Application adalah bagian yang mengatur proses aplikasi, sedangkan Use Case adalah tindakan spesifik yang dijalankan di dalamnya.

## C. Ports

**Port adalah aturan atau kontrak yang menyatakan kemampuan yang dibutuhkan oleh Application.**

Contoh:

Register User membutuhkan kemampuan untuk:

- Menyimpan User.
- Mengenkripsi password.
- Membuat ID pengguna.

Application tidak perlu tahu teknologi yang digunakan untuk melakukan pekerjaan tersebut.

Kita cukup membuat interface:

```go
type UserRepository interface {
    Save(ctx context.Context, user *domain.User) error
}
```

Artinya:

> Saya membutuhkan sesuatu yang bisa menyimpan User. Bagaimana cara menyimpannya, bukan urusan saya.

Contoh lainnya:

```go
type PasswordHasher interface {
    Hash(ctx context.Context, password string) (domain.PasswordHash, error)
}
```

Artinya:

> Saya membutuhkan sesuatu yang bisa menghasilkan password hash.

Port biasanya dibuat dalam bentuk interface Go.

## D. Adapter

**Adapter adalah kode yang mengimplementasikan Port agar bisa berkomunikasi dengan teknologi atau sistem tertentu.**

Misalnya Application membutuhkan `UserRepository`.

Kita membuat implementasi menggunakan GORM dan PostgreSQL:

```text
Application membutuhkan UserRepository
                 |
                 v
PostgresUserRepository
                 |
                 v
             PostgreSQL
```

Interface `UserRepository` adalah kontrak yang dibutuhkan Application.

`PostgresUserRepository` adalah implementasi yang menggunakan PostgreSQL.

Itulah contoh Adapter.

Adapter tidak harus selalu menggunakan database. Implementasi PasswordHasher menggunakan bcrypt juga bisa dianggap sebagai adapter.

## E. Delivery

**Delivery adalah bagian yang menerima permintaan dari luar aplikasi dan meneruskannya ke Application.**

Contoh:

Pengguna mengirim HTTP request:

```http
POST /register
```

Delivery menerima request, membaca JSON, lalu memanggil Use Case `RegisterUser`.

Gambaran:

```text
HTTP Request
     |
     v
HTTP Handler
     |
     v
RegisterUser Use Case
```

Pada PayCore, Delivery bisa berisi:

- HTTP Handler
- Router
- Middleware
- Request dan Response

Delivery tidak seharusnya menjadi tempat utama untuk menyimpan aturan bisnis.

## F. Infrastructure

**Infrastructure adalah bagian yang berisi implementasi teknis untuk menjalankan aplikasi.**

Contoh teknologi yang digunakan PayCore:

- PostgreSQL
- GORM
- bcrypt
- Redis
- Kafka

Kode untuk berkomunikasi dengan teknologi tersebut termasuk bagian Infrastructure.

Contoh struktur:

```text
infrastructure/
├── persistence/
│   └── postgres/
├── security/
│   └── bcrypt/
└── messaging/
    └── kafka/
```

Contoh tanggung jawab:

- `persistence/postgres`: menyimpan dan mengambil data dari PostgreSQL.
- `security/bcrypt`: menghasilkan password hash.
- `messaging/kafka`: mengirim dan menerima pesan Kafka.

---

# 5. Memahami hubungan semua bagian melalui proses Register User

Sekarang kita gabungkan istilah-istilah tadi.

Pengguna ingin membuat akun baru.

### Langkah 1 — Delivery

HTTP Handler menerima request:

```http
POST /register
```

Body:

```json
{
  "email": "user@email.com",
  "password": "secret123"
}
```

Handler meneruskan data ke Use Case.

### Langkah 2 — Application / Use Case

`RegisterUser` mengatur langkah-langkah registrasi:

1. Periksa email.
2. Periksa apakah email sudah digunakan.
3. Hash password.
4. Buat User.
5. Simpan User.

### Langkah 3 — Domain

Application menggunakan Domain untuk membuat User.

```go
user, err := domain.NewUser(
    userID,
    email,
    passwordHash,
    time.Now(),
)
```

Domain memastikan objek User dibuat sesuai aturan yang berlaku.

### Langkah 4 — Ports dan Infrastructure

Application membutuhkan Repository untuk menyimpan User.

Application memanggil interface:

```go
userRepository.Save(ctx, user)
```

Implementasi PostgreSQL menjalankan proses penyimpanan ke database.

### Langkah 5 — Hasil dikembalikan

Setelah berhasil:

```text
PostgreSQL
    |
    v
Application
    |
    v
Delivery
    |
    v
HTTP Response
```

Pengguna menerima respons bahwa akun berhasil dibuat.

---

# 6. Memahami struktur folder PayCore

Berikut struktur yang kita gunakan:

```text
services/
└── identity/
    ├── cmd/
    │   └── api/
    │       └── main.go
    │
    └── internal/
        ├── domain/
        │
        ├── application/
        │   ├── auth/
        │   └── ports/
        │
        ├── delivery/
        │   └── http/
        │       ├── handler/
        │       ├── middleware/
        │       └── router.go
        │
        └── infrastructure/
            ├── persistence/
            │   └── postgres/
            │
            ├── security/
            │
            └── config/
```

Penjelasan sederhananya:

| Folder              | Tanggung jawab                    |
| ------------------- | --------------------------------- |
| `cmd/api`           | Tempat menjalankan aplikasi       |
| `domain`            | Konsep dan aturan bisnis          |
| `application`       | Mengatur langkah-langkah Use Case |
| `application/ports` | Kontrak kebutuhan Application     |
| `delivery`          | Menerima request dari luar        |
| `infrastructure`    | Implementasi teknologi            |

Contoh lokasi file:

| File                                                     | Fungsinya                               |
| -------------------------------------------------------- | --------------------------------------- |
| `domain/user.go`                                         | Mendefinisikan User dan perilakunya     |
| `domain/value_objects.go`                                | Mendefinisikan Email dan PasswordHash   |
| `application/auth/register.go`                           | Mengatur proses registrasi              |
| `application/ports/user_repository.go`                   | Mendefinisikan kontrak penyimpanan User |
| `delivery/http/handler/register.go`                      | Menerima HTTP request registrasi        |
| `infrastructure/persistence/postgres/user_repository.go` | Menyimpan User ke PostgreSQL            |

---

# 7. Apa itu Dependency Rule?

Ini adalah aturan penting dalam Clean Architecture.

**Bagian inti aplikasi tidak boleh bergantung langsung pada detail teknologi luar.**

Contoh yang kita inginkan:

```text
Delivery
   |
   v
Application
   |
   v
Domain

Infrastructure ----> memenuhi kebutuhan Application
```

Application tidak perlu tahu apakah kita menggunakan:

- PostgreSQL atau MySQL.
- bcrypt atau algoritma hashing lain.
- Kafka atau sistem messaging lain.

Application hanya mengetahui kontrak yang dibutuhkannya.

Misalnya, jika suatu saat kita mengganti PostgreSQL menjadi MySQL, kita bisa mengganti implementasi Repository tanpa harus mengubah aturan bisnis User.

Itulah manfaat pemisahan ini.

---

# 8. Istilah yang perlu dipahami terlebih dahulu

Jangan berusaha menguasai semuanya sekaligus. Untuk tahap awal PayCore, prioritaskan istilah berikut:

1. **Domain** — masalah bisnis yang diselesaikan.
2. **Entity** — objek bisnis yang memiliki identitas.
3. **Value Object** — objek yang dikenali berdasarkan nilainya.
4. **Application** — bagian yang mengatur proses fitur.
5. **Use Case** — tindakan spesifik yang dijalankan aplikasi.
6. **Repository** — cara mengambil dan menyimpan objek bisnis.
7. **Infrastructure** — implementasi teknologi.
8. **Delivery** — pintu masuk request ke aplikasi.

Setelah konsep dasar tersebut dipahami, kita bisa mempelajari Aggregate, Aggregate Root, Domain Service, Domain Event, dan Bounded Context lebih mendalam ketika memang membutuhkannya dalam implementasi.

## Kesimpulan

Bayangkan PayCore sebagai sebuah perusahaan:

- **Domain**: bidang bisnis perusahaan.
- **Entity**: objek penting yang dikelola perusahaan, seperti User.
- **Value Object**: nilai yang digunakan objek tersebut, seperti Email.
- **Application**: mengatur proses pekerjaan.
- **Use Case**: pekerjaan spesifik, seperti mendaftarkan pengguna.
- **Repository**: pihak yang bertanggung jawab mengambil dan menyimpan data.
- **Delivery**: bagian yang menerima permintaan dari luar.
- **Infrastructure**: teknologi yang digunakan untuk menjalankan pekerjaan.

DDD membantu kita memahami dan menyusun aturan bisnis. Clean Architecture membantu kita memisahkan tanggung jawab kode agar bagian bisnis tidak terikat pada teknologi tertentu.

Kita akan mempelajarinya sambil membangun PayCore, bukan dengan menghafalkan semua istilah terlebih dahulu.
