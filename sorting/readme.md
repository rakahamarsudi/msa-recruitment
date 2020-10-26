# Sorting and visualization

1. *Design* dan implementasikan sebuah *program* atau *subprogram* yang akan menampilan visualisasi *data array* sederhana dalam bentuk *vertical barcharts*, dan sebagai tambahan tampilkan setiap nilai data di sumbu *horizontal*.
    
    ```
    INPUT: Numerical array
    [1, 4, 5, 6, 8, 2]

    OUTPUT: Vertical Barcharts

            |   
            |   
          | |  
        | | |   
      | | | |  
      | | | |  
      | | | | |
    | | | | | | 
    1 4 5 6 8 2 

    ```
2. Implementasikan algoritma *insertion sort*, dan gunakan *subprogram* (1) untuk memvisualisasikan setiap langkah/*steps* *sorting* 

    ```
    INPUT: Numerical array

    [1, 4, 5, 6, 8, 2]

    OUTPUT:

    - Sorted array (ascending)
    - Steps visualization

            |   
            |   
          | |  
        | | |   
      | | | |   
      | | | |   
      | | | | | 
    | | | | | | 
    1 4 5 6 8 2 

              | 
              | 
          |   | 
        | |   | 
      | | |   | 
      | | |   | 
      | | | | | 
    | | | | | | 
    1 4 5 6 2 8 

    ... dan seterusnya ...

    ```

3. Modifikasi *subprogram* (2) untuk *reverse sorting* dan lakukan juga visualisasi dengan *subprogram* (1)


## To Do:
+ **(DONE)** Membuat fungsi menulis *vertical graph*.
+ **(DONE)** Membuat fungsi *insertion sort*.
+ **(DONE)** Membuat fungsi *reverses sort*.

## Program
**Fungsi menulis *vertical graph***

`func printGraph(data []int) { .... }`

`for _, val := range data {...}` Mencari tahu nilai tertinggi data untuk menentukan tinggi *vertical graph*.

`for h := 0; h < maxValue; h++ {...}` Menulis *vertical graph* dari data.

**Fungsi *insertion sort* dan *reverse sort***

`func sortInc(data []int) { .... }` *Insertion Sort*.

`func sortDec(data []int) { .... }` *Reverse Sort*.

```
for i := 0; i < len(data)-1; i++ {
```

  `if data[i] > data[i+1] {` ==> (*Insertion Sort*) Memeriksa apabila data selanjutnya memiliki nilai yang lebih kecil dari data sekarang.

  `if data[i] < data[i+1] {` ==> (*Reverse Sort*) Memeriksa apabila data selanjutnya memiliki nilai yang lebih besar dari data sekarang.

```
    temp = data[i]
    data[i] = data[i+1]
    data[i+1] = temp
    done = false
    printGraph(data)
  }
}
```
Merubah posisi data hingga terurut.
