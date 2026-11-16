# Roadmap Pembelajaran Go

**Dokumen progres pembelajaran**\
**Pembaruan:** 16 November 2026

Dokumen ini merangkum materi Go yang telah dipelajari berdasarkan sesi
pembelajaran sejauh ini. Status dibedakan menjadi:

-   **Sudah dipelajari**: konsep dan sintaks telah dibahas.
-   **Sudah dipraktikkan**: konsep telah digunakan dalam latihan atau
    contoh kode.
-   **Perlu diperdalam**: dasar konsep telah dibahas, tetapi penerapan
    dan variasi kasusnya masih perlu dilatih.
-   **Belum dipelajari**: belum menjadi materi pembelajaran utama.

> Status "sudah dipelajari" tidak otomatis berarti materi sudah dikuasai
> sepenuhnya. Penguasaan perlu diuji kembali melalui latihan mandiri dan
> proyek.

------------------------------------------------------------------------

# 1. Ringkasan Progres

Pembelajaran sejauh ini mencakup fondasi bahasa Go, pemrograman berbasis
struct dan method, concurrency, komunikasi antargoroutine, sinkronisasi,
serta pengelolaan alur program menggunakan `defer`, `panic`, dan
`recover`.

Fokus yang paling banyak dibahas adalah **concurrency di Go**, terutama:

-   Goroutine
-   `sync.WaitGroup`
-   Channel
-   Buffered dan unbuffered channel
-   Channel direction
-   `select`
-   `sync.Mutex`
-   Race condition
-   Dasar `context`
-   Timer melalui `time.After`

Materi berikutnya dapat diarahkan pada pendalaman concurrency, interface
dan composition, lalu pembangunan aplikasi backend.

------------------------------------------------------------------------

# 2. Fondasi Bahasa Go

Bagian ini menjadi dasar untuk membaca dan menulis program Go. Materi
fondasional telah digunakan dalam latihan-latihan berikutnya.

  -----------------------------------------------------------------------
  Materi                  Penjelasan              Status
  ----------------------- ----------------------- -----------------------
  Struktur program Go     Susunan package,        Sudah dipelajari
                          import, dan fungsi      
                          `main`                  

  Package dan import      Mengorganisasi kode dan Sudah dipelajari
                          menggunakan package     
                          standar                 

  Function                Membuat fungsi,         Sudah dipelajari
                          parameter, dan return   
                          value                   

  Variabel dan konstanta  Deklarasi serta         Fondasi
                          penggunaan nilai        

  Tipe data               Menggunakan tipe        Fondasi
                          seperti `int`,          
                          `string`, `bool`        

  Control flow            Percabangan dan         Fondasi
                          perulangan              
  -----------------------------------------------------------------------

Bagian ini tetap perlu digunakan secara konsisten saat mengerjakan
latihan dan proyek. Daftar ini merupakan fondasi yang mendukung materi
lanjutan, bukan klaim bahwa setiap detail bahasa telah dituntaskan.

------------------------------------------------------------------------

# 3. Struct, Methods, Receivers, dan Pointer

## 3.1 Struct

`struct` mengelompokkan beberapa field ke dalam satu tipe data.

Contoh:

``` go
type Task struct {
    ID     int
    Name   string
    Status string
}

func main() {
    task := Task{
        ID:     1,
        Name:   "Download File",
        Status: "Pending",
    }

    fmt.Println(task.Name)
}
```

Hal yang perlu dipahami:

-   Mendefinisikan tipe struct.
-   Menentukan field dan tipe datanya.
-   Membuat nilai struct.
-   Mengakses field menggunakan tanda titik (`.`).
-   Mengubah nilai field.

**Status:** Sudah dipelajari dan digunakan dalam contoh.

## 3.2 Methods

Method adalah fungsi yang memiliki receiver dan terhubung dengan suatu
tipe.

``` go
type Task struct {
    Name string
}

func (task Task) Run() {
    fmt.Println(task.Name, "is running")
}
```

Pemanggilan:

``` go
task := Task{Name: "Download File"}
task.Run()
```

Hal yang perlu dipahami:

-   Perbedaan function biasa dan method.
-   Cara mendeklarasikan receiver.
-   Cara memanggil method melalui nilai struct.

**Status:** Sudah dipelajari.

## 3.3 Value Receiver

Value receiver menerima salinan nilai receiver.

``` go
func (task Task) Rename(name string) {
    task.Name = name
}
```

Perubahan pada salinan tersebut tidak mengubah struct asli.

## 3.4 Pointer Receiver

Pointer receiver menerima pointer ke nilai aslinya.

``` go
func (task *Task) Rename(name string) {
    task.Name = name
}
```

Perubahan melalui pointer receiver dapat mengubah struct asli.

  -----------------------------------------------------------------------
  Value Receiver                      Pointer Receiver
  ----------------------------------- -----------------------------------
  `func (t Task)`                     `func (t *Task)`

  Menerima salinan nilai              Menerima pointer ke nilai

  Perubahan field tidak mengubah      Dapat mengubah nilai asli
  nilai asli                          

  Cocok untuk operasi yang tidak      Cocok ketika method perlu
  perlu memodifikasi receiver         memodifikasi receiver
  -----------------------------------------------------------------------

**Status:** Sudah dipelajari.

## 3.5 Pointer

Pointer menyimpan alamat memori suatu nilai.

``` go
number := 100

pointer := &number

fmt.Println(number)   // 100
fmt.Println(pointer) // alamat memori
fmt.Println(*pointer) // 100
```

Operator utama:

  Operator   Kegunaan
  ---------- ----------------------------------------
  `&x`       Mengambil alamat dari `x`
  `*p`       Mengakses nilai yang ditunjuk oleh `p`
  `*T`       Menyatakan tipe pointer ke `T`

**Status:** Sudah dipelajari dan digunakan bersama pointer receiver.

------------------------------------------------------------------------

# 4. Goroutine dan Concurrency

Concurrency memungkinkan beberapa pekerjaan mengalami kemajuan secara
tumpang tindih. Go menyediakan goroutine sebagai salah satu
mekanismenya.

## 4.1 Goroutine

Goroutine dijalankan dengan menambahkan keyword `go` sebelum pemanggilan
fungsi.

``` go
func process(name string) {
    fmt.Println(name)
}

func main() {
    go process("Task 1")
    go process("Task 2")
}
```

Hal yang perlu dipahami:

-   `go function()` memulai goroutine.
-   Goroutine berjalan secara concurrent dengan goroutine pemanggil.
-   Program utama tidak otomatis menunggu goroutine selesai.
-   Urutan eksekusi dan output tidak selalu deterministik.

**Status:** Sudah dipelajari dan dipraktikkan.

## 4.2 Anonymous Goroutine

Goroutine juga dapat menjalankan anonymous function.

``` go
go func() {
    fmt.Println("Running in goroutine")
}()
```

**Status:** Sudah dipelajari.

## 4.3 `time.Sleep`

`time.Sleep` menunda goroutine yang memanggilnya selama durasi tertentu.

``` go
time.Sleep(2 * time.Second)
```

Sleep berguna untuk demonstrasi atau jeda waktu, tetapi bukan mekanisme
sinkronisasi yang tepat untuk memastikan pekerjaan goroutine selesai.

**Status:** Sudah digunakan dalam contoh concurrency.

------------------------------------------------------------------------

# 5. WaitGroup

`sync.WaitGroup` digunakan untuk menunggu sejumlah goroutine
menyelesaikan pekerjaan.

Import:

``` go
import "sync"
```

Tiga method utama:

  Method        Fungsi
  ------------- ------------------------------------------------
  `wg.Add(n)`   Menambah counter pekerjaan yang harus ditunggu
  `wg.Done()`   Mengurangi counter sebanyak satu
  `wg.Wait()`   Menunggu sampai counter kembali menjadi nol

Contoh:

``` go
func process(name string, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Println(name, "started")
    time.Sleep(time.Second)
    fmt.Println(name, "finished")
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2)

    go process("Task 1", &wg)
    go process("Task 2", &wg)

    wg.Wait()

    fmt.Println("All tasks completed")
}
```

Hal penting:

-   Panggil `Add` sebelum goroutine yang akan dihitung mulai berjalan.
-   Setiap pekerjaan yang dihitung harus memanggil `Done`.
-   `defer wg.Done()` membantu memastikan counter dikurangi ketika
    fungsi selesai.
-   Jangan menyalin `WaitGroup` setelah mulai digunakan; teruskan
    pointer ke goroutine.
-   `WaitGroup` menunggu penyelesaian, bukan mengirimkan hasil
    pekerjaan.

**Status:** Sudah dipelajari dan dipraktikkan.

------------------------------------------------------------------------

# 6. Channel

Channel merupakan sarana komunikasi untuk mengirim dan menerima nilai
antar-goroutine.

## 6.1 Membuat Channel

``` go
ch := make(chan int)
```

Channel di atas adalah unbuffered channel.

## 6.2 Send dan Receive

``` go
ch <- 100 // mengirim nilai

value := <-ch // menerima nilai
```

Operasi channel dapat memblokir goroutine sampai operasi pasangannya
siap.

## 6.3 Unbuffered Channel

Unbuffered channel tidak memiliki kapasitas penyimpanan. Pengiriman dan
penerimaan harus bertemu agar operasi dapat berlangsung.

``` go
ch := make(chan int)

go func() {
    ch <- 100
}()

value := <-ch
fmt.Println(value)
```

## 6.4 Buffered Channel

Buffered channel memiliki kapasitas penyimpanan.

``` go
ch := make(chan int, 2)

ch <- 10
ch <- 20

fmt.Println(len(ch)) // 2
fmt.Println(cap(ch)) // 2
```

Pengiriman dapat berlangsung tanpa penerima langsung selama buffer belum
penuh.

  Fungsi              Kegunaan
  ------------------- -----------------------------------------------
  `make(chan T)`      Membuat unbuffered channel
  `make(chan T, n)`   Membuat buffered channel berkapasitas `n`
  `len(ch)`           Jumlah elemen yang sedang tersimpan di buffer
  `cap(ch)`           Kapasitas channel

## 6.5 Menutup Channel

``` go
close(ch)
```

Channel biasanya ditutup oleh pihak yang mengirim data ketika tidak ada
lagi nilai yang akan dikirim.

Menutup channel bukan berarti menghapus data yang masih berada di
buffer. Data yang sudah dikirim tetap dapat diterima.

Jangan mengirim data ke channel yang sudah ditutup karena akan
menyebabkan panic. Menutup channel yang sudah ditutup juga menyebabkan
panic.

## 6.6 Range pada Channel

``` go
for value := range ch {
    fmt.Println(value)
}
```

`range` menerima nilai sampai channel ditutup dan seluruh nilai yang
sudah dikirim selesai diterima.

## 6.7 Memeriksa Status Channel

``` go
value, ok := <-ch
```

-   `value` berisi nilai yang diterima.
-   `ok` bernilai `true` jika penerimaan mendapatkan nilai dari channel
    yang masih terbuka.
-   `ok` bernilai `false` jika channel sudah ditutup dan tidak ada nilai
    tersisa.

## 6.8 Channel Direction

Direction membatasi operasi yang boleh dilakukan terhadap channel.

``` go
func sendData(ch chan<- int) {
    ch <- 100
}

func receiveData(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}
```

  Tipe           Hak operasi
  -------------- -----------------------
  `chan int`     Mengirim dan menerima
  `chan<- int`   Hanya mengirim
  `<-chan int`   Hanya menerima

**Status:** Sudah dipelajari dan dipraktikkan, termasuk
buffered/unbuffered channel, send/receive, close, range, `len`, `cap`,
dan channel direction.

------------------------------------------------------------------------

# 7. Select

`select` memilih operasi komunikasi channel yang siap dijalankan.
Sintaksnya mirip `switch`, tetapi case-nya berupa operasi channel.

## 7.1 Multiple Receive

``` go
select {
case value := <-ch1:
    fmt.Println("From ch1:", value)

case value := <-ch2:
    fmt.Println("From ch2:", value)
}
```

Jika beberapa operasi siap pada saat yang sama, Go memilih salah satunya
secara pseudo-random.

Jika tidak ada operasi yang siap dan tidak ada `default`, `select`
menunggu.

## 7.2 Default Case

``` go
select {
case value := <-ch:
    fmt.Println(value)

default:
    fmt.Println("No channel operation is ready")
}
```

`default` dijalankan ketika tidak ada case komunikasi yang siap. Ini
membuat pola tersebut non-blocking.

## 7.3 Send di dalam Select

``` go
ch := make(chan int, 1)

select {
case ch <- 100:
    fmt.Println("Data sent")
default:
    fmt.Println("Cannot send now")
}
```

## 7.4 Select di dalam For Loop

`select` dapat digunakan di dalam loop untuk terus memantau beberapa
channel.

``` go
for {
    select {
    case value := <-ch1:
        fmt.Println(value)

    case value := <-ch2:
        fmt.Println(value)
    }
}
```

Loop perlu memiliki kondisi berhenti yang sesuai. Jika semua channel
yang dipantau ditutup, case receive dapat terus siap dan menghasilkan
zero value. Salah satu pola yang dapat digunakan adalah memeriksa `ok`
dan menonaktifkan channel dengan menetapkannya menjadi `nil`.

``` go
for ch1 != nil || ch2 != nil {
    select {
    case value, ok := <-ch1:
        if !ok {
            ch1 = nil
            continue
        }
        fmt.Println(value)

    case value, ok := <-ch2:
        if !ok {
            ch2 = nil
            continue
        }
        fmt.Println(value)
    }
}
```

Channel `nil` tidak pernah siap untuk operasi send atau receive,
sehingga case terkait tidak dipilih.

**Status:** Sudah dipelajari. Pembahasan mencakup multiple case,
default, send, receive, loop, pemeriksaan `ok`, dan channel `nil`.

------------------------------------------------------------------------

# 8. Mutex dan Race Condition

## 8.1 Race Condition

Race condition dapat terjadi ketika beberapa goroutine mengakses data
yang sama secara bersamaan, setidaknya salah satunya melakukan
penulisan, tanpa sinkronisasi yang tepat.

Contoh situasi: beberapa goroutine menambah nilai `Stock` pada objek
Inventory secara bersamaan.

## 8.2 Mutex

`sync.Mutex` digunakan untuk mengatur akses eksklusif terhadap bagian
kode yang mengakses shared data.

``` go
type Inventory struct {
    Stock int
    mu    sync.Mutex
}

func (inventory *Inventory) AddStock(amount int) {
    inventory.mu.Lock()
    defer inventory.mu.Unlock()

    inventory.Stock += amount
}
```

Method utama:

  Method          Fungsi
  --------------- ------------------
  `mu.Lock()`     Mengunci mutex
  `mu.Unlock()`   Melepaskan kunci

`defer mu.Unlock()` membantu memastikan mutex dilepas ketika fungsi
selesai.

## 8.3 Latihan Inventory

Kita pernah menggunakan pola Inventory dengan lima goroutine. Setiap
goroutine menambahkan 10 stock.

Hasil yang diharapkan:

``` text
Final stock: 50
```

Mutex melindungi operasi penambahan agar tidak terjadi race pada field
`Stock`.

**Status:** Sudah dipelajari dan dipraktikkan melalui latihan Inventory.

------------------------------------------------------------------------

# 9. Context dan Timer

Context digunakan untuk meneruskan sinyal pembatalan, deadline, dan
batas waktu pada operasi yang berjalan.

## 9.1 `context.Background()`

Membuat context kosong yang umum digunakan sebagai titik awal.

``` go
ctx := context.Background()
```

## 9.2 `context.WithCancel()`

Membuat context turunan yang dapat dibatalkan.

``` go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

Memanggil `cancel()` menutup channel yang dikembalikan oleh `ctx.Done()`
dan memberi sinyal kepada operasi yang menggunakan context tersebut.

## 9.3 `ctx.Done()`

`ctx.Done()` menghasilkan receive-only channel.

``` go
select {
case <-ctx.Done():
    fmt.Println("Worker stopped")
}
```

Operasi receive menjadi siap ketika context dibatalkan atau deadline-nya
tercapai.

## 9.4 `time.After()`

`time.After` mengembalikan channel yang akan menerima nilai waktu
setelah durasi tertentu.

``` go
select {
case <-ctx.Done():
    fmt.Println("Worker stopped")

case <-time.After(3 * time.Second):
    fmt.Println("Worker completed")
}
```

Perhatikan bahwa `time.After(...)` harus digunakan sebagai operasi
receive di dalam `select`, sehingga sintaksnya menggunakan `<-`.

## 9.5 Memahami Tipe Return

  Ekspresi          Hasil
  ----------------- ------------------------------------------
  `ctx.Done()`      Receive-only channel bertipe `struct{}`
  `time.After(d)`   Receive-only channel bertipe `time.Time`
  `timer.C`         Channel untuk menerima event timer

**Status:** Dasar `Background`, `WithCancel`, `cancel`, `Done`, dan
penggunaan `time.After` bersama `select` sudah dibahas. `WithTimeout`,
`WithDeadline`, serta pola cancellation yang lebih kompleks masih perlu
diperdalam.

------------------------------------------------------------------------

# 10. Defer, Panic, dan Recover

## 10.1 Defer

`defer` menjadwalkan pemanggilan fungsi agar dijalankan ketika fungsi
yang memuatnya akan selesai.

``` go
func process() {
    defer fmt.Println("Cleanup")

    fmt.Println("Processing")
}
```

Output:

``` text
Processing
Cleanup
```

## 10.2 Multiple Defer dan LIFO

Jika terdapat beberapa pernyataan `defer`, pemanggilannya dilakukan
dengan urutan LIFO (*Last In, First Out*).

``` go
func process() {
    defer fmt.Println("Cleanup 1")
    defer fmt.Println("Cleanup 2")
    defer fmt.Println("Cleanup 3")

    fmt.Println("Processing")
}
```

Output:

``` text
Processing
Cleanup 3
Cleanup 2
Cleanup 1
```

## 10.3 Panic

`panic` menghentikan alur normal fungsi saat ini dan memulai proses
unwinding: fungsi-fungsi yang sedang ditinggalkan menjalankan `defer`
masing-masing.

``` go
func process() {
    fmt.Println("Task started")
    panic("task failed")
}
```

## 10.4 Recover

`recover` dapat menangkap panic yang sedang berlangsung jika dipanggil
dari fungsi deferred pada goroutine yang sama.

``` go
func process() {
    defer func() {
        result := recover()

        if result != nil {
            fmt.Println("Recovered:", result)
        }
    }()

    panic("task failed")
}
```

Setelah panic berhasil dipulihkan, eksekusi dapat berlanjut setelah
pemanggilan fungsi yang mengalami panic.

## 10.5 Latihan Defer, Panic, Recover

Kita sudah membahas contoh `processTask()` yang:

1.  Mendaftarkan beberapa `defer`.
2.  Memanggil `panic("task failed")`.
3.  Menangkap panic menggunakan `recover()`.
4.  Menampilkan pesan cleanup.
5.  Membiarkan fungsi `main` melanjutkan eksekusi setelah pemulihan.

**Status:** Sudah dipelajari dan dipraktikkan. Pendalaman penanganan
panic pada batas goroutine dan strategi error handling aplikasi masih
diperlukan.

------------------------------------------------------------------------

# 11. Daftar Materi yang Sudah Dibahas

  Area              Materi                             Status
  ----------------- ---------------------------------- ------------------------
  Struct            Struct dan fields                  Sudah dipelajari
  Methods           Method dan receiver                Sudah dipelajari
  Pointer           Value/pointer receiver, `&`, `*`   Sudah dipelajari
  Concurrency       Goroutine                          Sudah dipraktikkan
  Concurrency       Anonymous goroutine                Sudah dibahas
  Concurrency       WaitGroup: Add, Done, Wait         Sudah dipraktikkan
  Channel           Unbuffered dan buffered            Sudah dipraktikkan
  Channel           Send dan receive                   Sudah dipraktikkan
  Channel           Close dan range                    Sudah dipelajari
  Channel           Channel direction                  Sudah dipelajari
  Channel           `len` dan `cap`                    Sudah dipelajari
  Select            Multiple case dan default          Sudah dipelajari
  Select            Send dan receive                   Sudah dipelajari
  Select            Loop, `ok`, dan channel nil        Sudah dibahas
  Synchronization   Mutex, Lock, Unlock                Sudah dipraktikkan
  Concurrency       Race condition                     Sudah dipraktikkan
  Context           Background, WithCancel, Done       Dasar sudah dipelajari
  Timer             Sleep dan After                    Sudah dipelajari
  Error flow        Defer dan LIFO                     Sudah dipraktikkan
  Error flow        Panic dan Recover                  Sudah dipraktikkan

------------------------------------------------------------------------

# 12. Materi yang Perlu Diperdalam atau Dipelajari

Materi di bawah ini merupakan rekomendasi lanjutan. Jangan menganggapnya
sudah dikuasai hanya karena berkaitan dengan materi yang pernah dibahas.

## 12.1 Interface dan Composition

-   [ ] Interface dasar
-   [ ] Implicit interface implementation
-   [ ] Interface embedding
-   [ ] Composition
-   [ ] Type assertion
-   [ ] Type switch
-   [ ] Empty interface (`any`)
-   [ ] Pointer dan interface
-   [ ] Memahami typed nil pada interface

## 12.2 Advanced Concurrency

-   [ ] `sync.RWMutex`
-   [ ] `sync.Once`
-   [ ] `sync.Cond`
-   [ ] `sync.Map`
-   [ ] Atomic operations (`sync/atomic`)
-   [ ] Worker Pool
-   [ ] Fan-in
-   [ ] Fan-out
-   [ ] Pipeline pattern
-   [ ] Cancellation dan graceful shutdown
-   [ ] Deadlock dan strategi menghindarinya
-   [ ] Race detector (`go test -race`)

## 12.3 Context Lanjutan

-   [ ] `context.WithTimeout`
-   [ ] `context.WithDeadline`
-   [ ] `context.WithValue` dan penggunaan yang tepat
-   [ ] Cancellation propagation
-   [ ] Menggunakan context pada HTTP request
-   [ ] Graceful shutdown server

## 12.4 Error Handling

-   [ ] Membuat custom error
-   [ ] `errors.New`
-   [ ] `fmt.Errorf`
-   [ ] Error wrapping dengan `%w`
-   [ ] `errors.Is`
-   [ ] `errors.As`
-   [ ] Strategi menangani panic di batas aplikasi

## 12.5 Testing

-   [ ] Unit testing
-   [ ] Table-driven tests
-   [ ] Subtests
-   [ ] Testify atau assertion library
-   [ ] Mocking dan dependency injection
-   [ ] Benchmark
-   [ ] Fuzz testing

------------------------------------------------------------------------

# 13. Roadmap Menuju Backend Development

Setelah fondasi bahasa dan concurrency, pembelajaran dapat berlanjut
menuju pembuatan aplikasi backend.

## Stage 1 --- Go Fundamentals

-   [x] Dasar struktur program dan function
-   [x] Fondasi tipe data dan control flow

## Stage 2 --- Intermediate Go

-   [x] Struct dan Methods
-   [x] Receivers dan Pointer
-   [x] Goroutines
-   [x] WaitGroup
-   [x] Channels
-   [x] Select
-   [x] Mutex dan Race Condition
-   [x] Dasar Context
-   [x] Defer, Panic, dan Recover

## Stage 3 --- Go Language Deepening

-   [ ] Interfaces
-   [ ] Composition
-   [ ] Advanced Concurrency
-   [ ] Advanced Context
-   [ ] Advanced Error Handling
-   [ ] Testing
-   [ ] Generics
-   [ ] Package Design

## Stage 4 --- Backend Fundamentals

-   [ ] Cara kerja HTTP
-   [ ] HTTP methods
-   [ ] Status codes
-   [ ] Headers
-   [ ] Request dan Response
-   [ ] JSON encoding dan decoding
-   [ ] HTTP server di Go
-   [ ] Routing
-   [ ] Middleware
-   [ ] REST API design

## Stage 5 --- Database dan Persistence

-   [ ] Relational database
-   [ ] SQL dan query
-   [ ] Database driver Go
-   [ ] `database/sql`
-   [ ] Connection pooling
-   [ ] Transactions
-   [ ] Repository pattern
-   [ ] Database migrations

## Stage 6 --- Backend Application Architecture

-   [ ] Layered architecture
-   [ ] Handler dan service
-   [ ] Repository
-   [ ] Dependency injection
-   [ ] Configuration management
-   [ ] Logging
-   [ ] Validation
-   [ ] Authentication
-   [ ] Authorization

## Stage 7 --- Production Readiness

-   [ ] Integration testing
-   [ ] Docker
-   [ ] Environment configuration
-   [ ] Graceful shutdown
-   [ ] Observability
-   [ ] Metrics dan tracing
-   [ ] Deployment
-   [ ] Security fundamentals

------------------------------------------------------------------------

# 14. Rekomendasi Urutan Belajar Berikutnya

Urutan yang disarankan berdasarkan materi yang sudah dibahas:

1.  **Interfaces** --- memahami kontrak perilaku antar-tipe.
2.  **Composition** --- menyusun perilaku dan struktur program tanpa
    bergantung pada inheritance.
3.  **Type Assertion dan Type Switch** --- bekerja dengan nilai
    interface secara lebih spesifik.
4.  **Context lanjutan** --- memahami timeout, deadline, dan
    cancellation propagation.
5.  **Advanced Concurrency** --- mempelajari worker pool, fan-in,
    fan-out, dan atomic operations.
6.  **Error Handling** --- membangun penanganan error yang konsisten.
7.  **Testing** --- menguji fungsi, komponen, dan perilaku concurrent.
8.  **HTTP Server** --- mulai membangun aplikasi backend.

------------------------------------------------------------------------

# 15. Cara Menggunakan Dokumen Ini

Gunakan checklist untuk memperbarui progres belajar.

-   Tandai materi sebagai selesai setelah memahami konsep dan
    menyelesaikan latihan.
-   Buat catatan khusus untuk konsep yang masih sering tertukar.
-   Ulangi materi concurrency dengan latihan yang melibatkan beberapa
    goroutine.
-   Buat proyek kecil yang menggabungkan struct, interface, channel,
    context, dan error handling.
-   Jangan hanya menghafal syntax; pastikan memahami alasan dan kondisi
    kapan suatu fitur digunakan.

**Target utama:** mampu merancang dan membangun aplikasi backend Go yang
terstruktur, aman, dapat diuji, serta mampu menangani concurrency dengan
benar.
