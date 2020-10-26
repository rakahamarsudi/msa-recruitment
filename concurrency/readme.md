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
+ Implementasi *concurrent* process menggunakan *goroutine*
+ Menyeleksi museum berdasarkan kota.
+ Mengaplikasikan *Worker*
+ Menyimpan informasi museum kedalam queue
+ **(DONE)** Menyimpan informasi kedalam file CSV.



## Program
**Mendapatkan informasi melalui API**

`res, err := http.Get( .... )` Mengakses API untuk mendapatkan informasi museum.

`data, _ := ioutil.ReadAll(res.Body)` Mendapatkan informasi (bentuk JSON). 

`err = json.Unmarshal(data, &jsonData)` Parsing informasi (dari bentuk JSON).

**Membuat dan menulis file CSV**

`csvFile, err := os.Create("./data.csv")` Membuat file CSV.

`writer := csv.NewWriter(csvFile)` Membuat *writter* untuk menuliskan isi file

`func buildCSV(data jsonInfoMuseum, file *csv.Writer) { .... }` Menulis file CSV berdasarkan informasi museum.