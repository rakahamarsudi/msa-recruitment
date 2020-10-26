# Compare Folder

### TASK #1. Implementasikan sebuah program yang membandingkan isi dari dua direktori melalui parameter.
1. Jika file ada di source dan target, abaikan
2. Jika file ada di source tapi tidak ada di target berikan keterangan NEW
3. Jika file tidak ada di source tapi ada di target berikan keterangan DELETED
### TASK #2. Modifikasi program #1 untuk compare file content untuk rule (1), jika ada perbedaan beri keterangan MODIFIED

## To Do:
+ **(DONE)** Mendapatkan informasi *directory tree*.
+ **(DONE)** Menyimpan informasi ke dalam *maps*.
+ **(DONE)** Membandingkan isi *directory tree*.
+ **(DONE)** Memberikan keterangan terhadap kondisi *directory*.

## Program
**Mendapatkan argumen penjalanan program sebagai identifikasi directory yang hendak dibandingkan**

`srcDir := os.Args[1]` directory asal.

`dstDir := os.Args[2]` directory target.

**Fungsi untuk mendapatkan informasi *directory tree***

`func lookDir(dirAddress string) (m map[string]int) { .... }`

```
err := filepath.Walk(dirAddress,
    func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        m[path] = int(info.Size())
        return nil
    })
```

`filepath.Walk` mendapatkan setiap directory dalam *directory tree*.

`func(path string, info os.FileInfo, err error)` mendapatkan informasi directory.

`m[path] = int(info.Size())` menyimpan informasi directory kedalam *map* (untuk proses pembandingan directory).

**Fungsi untuk membandingkan isi *directory tree***

`func cmprDir(src, dst map[string]int) { .... }`

```
for dir, size := range src {
    val, ok := dst[dir]
    if !ok {
        fmt.Println(dir, "NEW")
    } else if size != val {
        fmt.Println(dir, "MODIFIED")
    }
}
```

`val, ok := dst[dir]` memeriksa ketersediaan directory yang sama dari *asal* pada *target*.

`if !ok {...} else if size != val {...}` memberikan keterangan berdasarkan ketersediaan dan persamaan kondisi directory berdasarkan besar memori direcory.

```
for dir := range dst {
    _, ok := src[dir]
    if !ok {
        fmt.Println(dir, "DELETED")
    }
}
```
memeriksa directory dari *target* pada *asal* yang tidak tersedia, untuk diberikan keterangan DELETED.