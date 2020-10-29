package main

import (
	"fmt"
	"math/rand"
	"time"
)

type queueMobil struct {
	antrian []mobil
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
	pekerja   float32
}

// Biaya : Menghitung pendapatan dari tiap mobil
func Biaya(car mobil, penjualan saldo) saldo {
	penjualan.pemasukan += float32(car.tiket) * 0.9
	penjualan.pekerja += float32(car.tiket) * 0.1
	return penjualan
}

//RandMobil : Men-generate mobil secara random
func RandMobil(fKecil, fBesar, varian int) []mobil {
	queueMobil := make([]mobil, fKecil+fBesar)
	var queue mobil
	var rMobil int
	mKecilFull := 0
	mBesarFull := 0

	rand.Intn(varian)
	for i := 0; i < len(queueMobil); i++ {
		rMobil = rand.Intn(varian)
		for mKecilFull == fKecil && rMobil <= 1 {
			rMobil = rand.Intn(varian)
		}
		for mBesarFull == fBesar && rMobil >= 2 {
			rMobil = rand.Intn(varian)
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
		}
	}
	return queueMobil
}

//RuteMobil : Menentukan rute yang harus dilalui setiap mobil
func RuteMobil(car mobil) (rute []string) {
	rute = append(rute, "A")
	if car.bensin < 10 {
		rute = append(rute, "G")
	}

	if car.jenis == "pickup" || car.jenis == "truk" {
		rute = append(rute, "C")
	}

	if car.jenis == "bis" || car.jenis == "truk" {
		rute = append(rute, "L")
	} else {
		rute = append(rute, "S")
	}
	return rute
}

// PomBensin : Kegiatan yang dilakukan di pom bensin
func (car *mobil) PomBensin() {
	fmt.Printf("  [POM BENSIN] ===> %8v sedang mengisi bensin \n", car.nama)
	car.bensin = 100
	car.rute = car.rute[1:]
}

// BeaCukai : Kegiatan yang dilakukan di bea cukai
func (car *mobil) BeaCukai() {
	car.kargo = "OPEN"
	fmt.Printf("  [BEA CUKAI] ====> %8v sedang diperiksa  (kargo: %v)\n", car.nama, car.kargo)
	time.Sleep(200 * time.Millisecond)
	car.kargo = "CLOSE"
	fmt.Printf("                    %8v selesai diperiksa (kargo: %v)\n", car.nama, car.kargo)
}

//MasukFerry : Kegiatan memasukkan mobil kedalam kapal
func (kapal *ferry) MasukFerry(car mobil) {
	if len(kapal.muatan) < kapal.kapasitas {
		kapal.muatan = append(kapal.muatan, car)
		fmt.Printf("  [%v] ==> %8v Masuk kapal\n", kapal.jenis, car.nama)
	} else {
		fmt.Printf("  [%v] ==> Ferry Penuh\n", kapal.jenis)
	}
}

// PemindahanMobil : Proses staff melakukan pemindahan mobil di pelabuhan
func PemindahanMobil(queue queueMobil, c chan mobil) {
	for _, car := range queue.antrian {
		c <- car
	}
	close(c)
}

// ManageMoving : Melakukan proses perindahan kendaraan sesuai rute
func (queue *queueMobil) ManageMoving(car mobil) {
	fmt.Printf("==== PERPINDAHAN\n")
loop:
	for len(car.rute) != 0 {
		switch car.rute[0] {
		case "A":
			{
				car.rute = car.rute[1:]
				keuangan = Biaya(car, keuangan)
			}
		case "G":
			{
				car.rute = car.rute[1:]
				car.PomBensin()
			}
		case "C":
			{
				car.rute = car.rute[1:]
				car.BeaCukai()
			}
		case "S":
			{
				car.rute = car.rute[1:]
				ferryS.MasukFerry(car)
				queue.antrian = queue.antrian[1:]
				break loop
			}
		case "L":
			{
				car.rute = car.rute[1:]
				ferryL.MasukFerry(car)
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
			fmt.Printf("  Mobil %8v (bensin %3v%%) berada di %v (kargo %v)\n", car.nama, car.bensin, car.rute[0], car.kargo)
		} else {
			fmt.Printf("  Mobil %8v (bensin %3v%%) berada di %v \n", car.nama, car.bensin, car.rute[0])
		}
	}
}

func (kapal *ferry) CekMuatan() {
	fmt.Printf("==== MUATAN %v\n", kapal.jenis)
	for i, car := range kapal.muatan {
		fmt.Printf("  [%v] %8v (bensin %v%%)\n", i, car.nama, car.bensin)
	}
}

func (pendapatan *saldo) CetakKeuangan() {
	fmt.Println("==== PENDAPATAN (hari ini)")
	fmt.Printf("  Hasil Penjualan Tiket %10v\n", pendapatan.pemasukan+pendapatan.pekerja)
	fmt.Printf("  Pemasukan Pelabuhan   %10v\n", pendapatan.pemasukan)
	fmt.Printf("  Perolehan Karyawan    %10v\n", pendapatan.pekerja)
}

var (
	antrianMobil queueMobil
	ferryS       ferry
	ferryL       ferry
	keuangan     saldo
)

func main() {
	keuangan = saldo{0, 0}
	antrianMobil.antrian = RandMobil(8, 6, 4)
	ferryS.jenis = "Ferry Kecil"
	ferryS.kapasitas = 8
	ferryL.jenis = "Ferry Besar"
	ferryL.kapasitas = 6

	staff := make(chan mobil)

	go PemindahanMobil(antrianMobil, staff)

	for car := range staff {
		antrianMobil.PosisiMobil()
		antrianMobil.ManageMoving(car)
		ferryS.CekMuatan()
		ferryL.CekMuatan()
		keuangan.CetakKeuangan()
	}
	fmt.Printf("\nEstimasi Gaji Karyawan: %10.f/bulan\n", 30*keuangan.pekerja)
}
