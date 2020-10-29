package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type queueMobil struct {
	antrian []mobil
}
type chStaff struct {
	id  int
	car mobil
}

type mobil struct {
	nama   string
	jenis  string
	tiket  int
	bensin int
	kargo  string
	rute   []string
}

type ferry struct {
	jenis     string
	kapasitas int
	muatan    []mobil
}

type saldo struct {
	pemasukan float32
	pekerja1  float32
	pekerja2  float32
}

// Biaya : Menghitung pendapatan dari tiap mobil
func Biaya(car mobil, penjualan saldo, staff int) saldo {
	if staff == 1 {
		penjualan.pemasukan += float32(car.tiket) * 0.9
		penjualan.pekerja1 += float32(car.tiket) * 0.1
	} else {
		penjualan.pemasukan += float32(car.tiket) * 0.89
		penjualan.pekerja2 += float32(car.tiket) * 0.11
	}
	return penjualan
}

//RandMobil : Men-generate mobil secara random
func RandMobil(fKecil, fBesar, fEco, varian int) []mobil {
	queueMobil := make([]mobil, fKecil+fBesar+fEco)
	var queue mobil
	var rMobil int
	mKecilFull := 0
	mBesarFull := 0
	mEcoFull := 0

	rand.Intn(varian)
	for i := 0; i < len(queueMobil); i++ {
		for true {
			rMobil = rand.Intn(varian)
			if mKecilFull == fKecil && rMobil <= 1 {
				continue
			} else if mBesarFull == fBesar && rMobil >= 2 && rMobil <= 3 {
				continue
			} else if mEcoFull == fEco && rMobil == 4 {
				continue
			} else {
				break
			}
		}

		switch rMobil {
		case 0:
			{
				mKecilFull++
				queue.jenis = "mobil"
				queue.nama = fmt.Sprintf("%v %v", queue.jenis, mKecilFull)
				queue.tiket = 30000
				queue.bensin = rand.Intn(100)
				queue.kargo = "CLOSE"
				queue.rute = RuteMobil(queue)
				queueMobil[i] = queue
			}
		case 1:
			{
				mKecilFull++
				queue.jenis = "pickup"
				queue.nama = fmt.Sprintf("%v %v", queue.jenis, mKecilFull)
				queue.tiket = 40000
				queue.bensin = rand.Intn(100)
				queue.kargo = "CLOSE"
				queue.rute = RuteMobil(queue)
				queueMobil[i] = queue
			}
		case 2:
			{
				mBesarFull++
				queue.jenis = "bis"
				queue.nama = fmt.Sprintf("%v %v", queue.jenis, mBesarFull)
				queue.tiket = 50000
				queue.bensin = rand.Intn(100)
				queue.kargo = "CLOSE"
				queue.rute = RuteMobil(queue)
				queueMobil[i] = queue
			}
		case 3:
			{
				mBesarFull++
				queue.jenis = "truk"
				queue.nama = fmt.Sprintf("%v %v", queue.jenis, mBesarFull)
				queue.tiket = 60000
				queue.bensin = rand.Intn(100)
				queue.kargo = "CLOSE"
				queue.rute = RuteMobil(queue)
				queueMobil[i] = queue
			}
		case 4:
			{
				mEcoFull++
				queue.jenis = "listrik"
				queue.nama = fmt.Sprintf("%v %v", queue.jenis, mEcoFull)
				queue.tiket = 10000
				queue.bensin = rand.Intn(100)
				queue.kargo = "CLOSE"
				queue.rute = RuteMobil(queue)
				queueMobil[i] = queue
			}
		}
	}
	return queueMobil
}

//RuteMobil : Menentukan rute yang harus dilalui setiap mobil
func RuteMobil(car mobil) (rute []string) {
	rute = append(rute, "A")
	if car.bensin < 10 {
		if car.jenis == "listrik" {
			rute = append(rute, "B")
		} else {
			rute = append(rute, "G")
		}
	}

	if car.jenis == "pickup" || car.jenis == "truk" {
		rute = append(rute, "C")
	}

	if car.jenis == "bis" || car.jenis == "truk" {
		rute = append(rute, "L")
	} else if car.jenis == "listrik" {
		rute = append(rute, "E")
	} else {
		rute = append(rute, "S")
	}
	return rute
}

// PomBensin : Kegiatan yang dilakukan di pom bensin
func (car *mobil) PomBensin(staff int) {
	fmt.Printf("  [POM BENSIN] ===> (staff %v) %10v sedang mengisi bensin \n", staff, car.nama)
	time.Sleep(200 * time.Millisecond)
	car.bensin = 100
	car.rute = car.rute[1:]
}

// PomBensin : Kegiatan yang dilakukan di pom bensin
func (car *mobil) StatBaterai(staff int) {
	fmt.Printf("  [BATERAI] ======> (staff %v) %10v sedang mengisi baterai \n", staff, car.nama)
	time.Sleep(200 * time.Millisecond)
	car.bensin = 100
	car.rute = car.rute[1:]
}

// BeaCukai : Kegiatan yang dilakukan di bea cukai
func (car *mobil) BeaCukai(staff int) {
	car.kargo = "OPEN"
	fmt.Printf("  [BEA CUKAI] ====> (staff %v) %10v sedang diperiksa  (kargo: %v)\n", staff, car.nama, car.kargo)
	time.Sleep(500 * time.Millisecond)
	car.kargo = "CLOSE"
	fmt.Printf("                    (staff %v) %10v selesai diperiksa (kargo: %v)\n", staff, car.nama, car.kargo)
}

//MasukFerry : Kegiatan memasukkan mobil kedalam kapal
func (kapal *ferry) MasukFerry(car mobil, staff int) {
	if len(kapal.muatan) < kapal.kapasitas {
		kapal.muatan = append(kapal.muatan, car)
		fmt.Printf("  [%v] ==> (staff %v) %10v Masuk kapal\n", kapal.jenis, staff, car.nama)
	} else {
		fmt.Printf("  [%v] ==> Ferry Penuh\n", kapal.jenis)
	}
}

// PemindahanMobil : Proses staff melakukan pemindahan mobil di pelabuhan
func (queue *queueMobil) PemindahanMobil(c1, c2 chan mobil) {
	for _, car := range queue.antrian {
		select {
		case c1 <- car:
			{
			}
		case c2 <- car:
			{
			}
		}
	}
	close(c1)
	close(c2)
}

// ManageMoving : Melakukan proses perindahan kendaraan sesuai rute
func (queue *queueMobil) ManageMoving(car mobil, staff int, wg *sync.WaitGroup) {
	defer wg.Done()
loop:
	for len(car.rute) != 0 {
		switch car.rute[0] {
		case "A":
			{
				car.rute = car.rute[1:]
				keuangan = Biaya(car, keuangan, staff)
			}
		case "G":
			{
				car.rute = car.rute[1:]
				car.PomBensin(staff)
			}
		case "B":
			{
				car.rute = car.rute[1:]
				car.StatBaterai(staff)
			}
		case "C":
			{
				car.rute = car.rute[1:]
				car.BeaCukai(staff)
			}
		case "S":
			{
				car.rute = car.rute[1:]
				ferryS.MasukFerry(car, staff)
				queue.antrian = queue.antrian[1:]
				break loop
			}
		case "L":
			{
				car.rute = car.rute[1:]
				ferryL.MasukFerry(car, staff)
				queue.antrian = queue.antrian[1:]
				break loop
			}
		case "E":
			{
				car.rute = car.rute[1:]
				ferryE.MasukFerry(car, staff)
				queue.antrian = queue.antrian[1:]
				break loop
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// PosisiMobil : Menampilkan posisi dan kondisi seluruh mobil
func (queue *queueMobil) PosisiMobil() {
	fmt.Println("\n==== ANTRIAN KENDARAAN ")
	for _, car := range queue.antrian {
		if car.jenis == "pickup" || car.jenis == "truk" {
			fmt.Printf("  Mobil %10v (bensin %3v%%) berada di %v (kargo %v)\n", car.nama, car.bensin, car.rute[0], car.kargo)
		} else {
			fmt.Printf("  Mobil %10v (bensin %3v%%) berada di %v \n", car.nama, car.bensin, car.rute[0])
		}
	}
}

func (kapal *ferry) CekMuatan() {
	fmt.Printf("==== MUATAN %v\n", kapal.jenis)
	for i, car := range kapal.muatan {
		fmt.Printf("  [%v] %10v (bensin %v%%)\n", i, car.nama, car.bensin)
	}
}

func (pendapatan *saldo) CetakKeuangan() {
	fmt.Println("==== PENDAPATAN (hari ini)")
	fmt.Printf("  Hasil Penjualan Tiket  %10v\n", pendapatan.pemasukan+pendapatan.pekerja1+pendapatan.pekerja2)
	fmt.Printf("  Pemasukan Pelabuhan    %10v\n", pendapatan.pemasukan)
	fmt.Printf("  Perolehan Karyawan 01  %10v\n", pendapatan.pekerja1)
	fmt.Printf("  Perolehan Karyawan 02  %10v\n", pendapatan.pekerja2)
}

var (
	antrianMobil queueMobil
	ferryS       ferry
	ferryL       ferry
	ferryE       ferry
	keuangan     saldo
)

func main() {
	keuangan = saldo{0, 0, 0}
	antrianMobil.antrian = RandMobil(8, 6, 8, 5)
	ferryS.jenis = "Ferry Kecil"
	ferryS.kapasitas = 8
	ferryL.jenis = "Ferry Besar"
	ferryL.kapasitas = 6
	ferryE.jenis = "Ferry Eco"
	ferryE.kapasitas = 8

	staff1 := make(chan mobil)
	staff2 := make(chan mobil)
	var wg sync.WaitGroup

	go antrianMobil.PemindahanMobil(staff1, staff2)

	for len(antrianMobil.antrian) != 0 {
		antrianMobil.PosisiMobil()
		fmt.Printf("==== PERPINDAHAN\n")
		car1, ok1 := <-staff1
		if ok1 {
			wg.Add(1)
			antrianMobil.ManageMoving(car1, 1, &wg)
		}
		car2, ok2 := <-staff2
		if ok2 {
			wg.Add(1)
			antrianMobil.ManageMoving(car2, 2, &wg)
		}
		wg.Wait()

		ferryS.CekMuatan()
		ferryL.CekMuatan()
		ferryE.CekMuatan()
		keuangan.CetakKeuangan()
	}

	fmt.Printf("\nEstimasi Pendapatan Pelabuhan: %10.f/bulan\n", 30*keuangan.pemasukan)
	fmt.Printf("Estimasi Gaji Karyawan 1:      %10.f/bulan\n", 30*keuangan.pekerja1)
	fmt.Printf("Estimasi Gaji Karyawan 2:      %10.f/bulan\n", 30*keuangan.pekerja2)
}
