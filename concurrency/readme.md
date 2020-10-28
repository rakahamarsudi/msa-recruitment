# Concurrency Task Worker

Implementasi sebuah program untuk merangkum data museum di indonesia berdasarkan lokasi kabupaten/kota, simpan dalam file csv nama yang sesuai.

```
Kota Jakarta Pusat.csv
Kota Malang.csv
``` 

* Untuk sumber informasi gunakan API yang disediakan oleh **Open Data Indonesia** khusus data [museum](http://data.go.id/dataset/museum-indonesia)
* Gunakan *net/http* package untuk mengambil data dari API yang disediakan.
* Olah data mentah dari API jika diperlukan.
* Implementasi *concurrent* process untuk process mengambilan data dari API menggunakan *goroutine* . untuk ilustrasi [ilustrasi](https://talks.golang.org/2012/concurrency.slide)
* Batasi jumlah *concurrent* process, dengan mengaplikasikan *Queue* dan *Worker*.

## To Do:
+ **(DONE)** Mengakses informasi dari API.
+ **(DONE)** Menyeleksi (parsing JSON) informasi museum.
+ **(DONE)** Implementasi *concurrent* process menggunakan *goroutine*
+ **(DONE)** Menyeleksi museum berdasarkan kota.
+ **(DONE)** Mengaplikasikan *Worker* (channel buffered)
+ **(DONE)** Menyimpan informasi museum kedalam queue
+ **(DONE)** Menyimpan informasi kedalam file CSV.
+ **(DONE)** Menggunakan Commandline Arguments



## Program
**Mengunakan commandline arguments untuk menentukan jumlah *worker* dan *directory hasil***

`argBuff := flag.Int("concurrent_limit", 2, "an Int")` Menggunakan *flag* untuk mendapatkan argumen jumlah worker (disimpan pada *argBuff*).

`argLoc := flag.String("output", "./museum", "a String directory")` Menggunakan *flag* untuk mendapatkan argumen directory hasil (disimpan pada *argLoc*).

`flag.Parse()` Melakukan parsing Flag untuk menggunakan argumen yang tersedia.

**Mendapatkan informasi melalui API**

`res, err := http.Get( .... )` Mengakses API untuk mendapatkan informasi museum.

`data, _ := ioutil.ReadAll(res.Body)` Mendapatkan informasi (bentuk JSON). 

`err = json.Unmarshal(data, &jsonData)` Parsing informasi (dari bentuk JSON).

**Mengelompokkan data museum berdasarkan lokasi kabupaten**

`func filterData(data jsonDataMuseum) (queues []queueMuseum) { .... }`

```
for _, info := range data.DataMuseum {
    for i, queue := range queues {
        if info.Kabupaten == queue.location {  // ================ (1)
            queues[i].museums = append(queue.museums, info)  // == (2)
            make = false
            break
        } else {
            make = true
            continue
        }
    }
    if make == true {  // ======================================== (3)
        var temp queueMuseum
        temp.location = info.Kabupaten
        temp.museums = append(temp.museums, info)

        queues = append(queues, temp)
    }
}
```

\* perhatikan : type *queueMuseum* struct
1. Memeriksa ketersediaan queue (queueMuseum.location) dengan kabupaten yang sama dengan kabupaten data museum.
2. Bila tersedia, maka data akan ditambahkan pada queue (queueMuseum.museums).
3. Bila tidak tersedia, maka akan dibuat queue baru.

**Menggunakan *goroutine* untuk proses data**

`c := make(chan queueMuseum, *argBuff)` Mendeklarasikan channel dan buffer untuk goroutine.

`var wg sync.WaitGroup` Mendeklarasikan *waitgroup* untuk memastikan tahapan proses berjalan berjalan *berurutan*.

`go queueChannel(queues, c)` Memasukkan data (queue museum) kedalam channel untuk selanjutnya diproses.

`for queue := range c { .... }` Memproses data (queue museum) dari channel untuk diproses

**Membuat dan menulis file CSV**

`csvFile, err := os.Create(string(*argLoc + "/" + queue.location + ".csv"))` Membuat file CSV berdasarkan pengelompokan data museum berdasarkan lokasi.

`writer := csv.NewWriter(csvFile)` Membuat *writter* untuk menuliskan isi file.

`go writing(queue, writer, &wg)` Menulis data (informasi) museum pada file CSV yang telah dibuat.

\* penggunaan `wg.Add(1)` dan `wg.Wait()` berguna untuk memastikan proses penulisan telah selesai dilakukan sebelum proses dilanjutkan lagi.
