# Unique Queue

Implementasikan sebuah *unique queue* yang mempunyai *fixed size*, *operation signature* sudah dideklarasikan di ```type Queue interface```.

Implementasi harus memenuhi syarat dari *unit testing* yang telah disiapkan.

## To Do:
+ **(DONE)** Membuat *array* queue.
+ **(DONE)** Mendekalrasikan (membentuk) queue mealui fungsi *New*.
+ **(DONE)** Membangun fungsi sesuai *operation signature*

## Program
**Mendeklarasikan *array* queue**

```
type queueBox struct {
	maxItem int
	items   []interface{}
}
```
Mendeklarasikan struct yang berisi informasi *array* queue.


```
func New(size int) queueBox {
	q := queueBox{}
	q.maxItem = size
	return q
}
```
Menginisialisasi queue, dengan parameter *size* digunakan sebagai batas items yang dapat ditampung.

**Fungsi / operasi yang berhubungan dengan queue**

```
func (q *queueBox) Push(key interface{}) {
	if len(q.items) < q.maxItem {  // ================== (1)
		q.items = append(q.items, key)
	} else {  // ======================================= (2)
		q.items = append(q.items[1:], key)
	}
}
```
Fungsi *Push* bertujuan memasukkan data ke dalam queue dengan ketentuan :
1. Bila jumlah item pada queue kurang dari *maxItem*, maka item baru akan langsung ditambahkan ke dalam queue.
2. Bila jumlah item pada queue sudah mencapai *maxItem*, maka item pada queue akan digeser (item terlama / terdepan akan dihapus) dan item baru akan ditambahkan ke dalam queue.

```
func (q *queueBox) Pop() interface{} {
	itemPop := q.items[0]
	q.items = q.items[1:]
	return itemPop
}
```
Fungsi *Pop* bertujuan mengambil data terdepan (terlama) dari queue, kemudian menggeser seluruh data tersisa.

```
func (q *queueBox) Contains(key interface{}) bool {
	var has bool
	for _, val := range q.items {
		if val == key {
			has = true
			break
		} else {
			has = false
		}
	}
	return has
}
```
Fungsi *Contains* bertujuan memeriksa ketersedian data (tertentu) di dalam queue. Bila data terseduia, maka fungsi akan mengembalikan nilai true, dan sebaliknya bila data tidak tersedia.

```
func (q *queueBox) Len() int {
	return len(q.items)
}
```
Fungsi *Len* bertujuan mendapatkan informasi jumlah data yang tersedia di dalam queue.

```
func (q *queueBox) Keys() []interface{} {
	return q.items
}
```
Fungsi *Keys* bertujuan mendapatkan (menampilkan) isi dari seluruh data di dalam queue.