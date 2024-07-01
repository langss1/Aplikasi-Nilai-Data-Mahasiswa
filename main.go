package main

import "fmt"

const NMAX int64 = 10000

const NKUL int64 = 100

type sidekuliah struct {
	jumlah      int64
	nilai, skor float64
	nama        string
}

type kuliah struct {
	matkul, nama, grade   string
	kredit                int64
	tru                   bool
	nilai, quiz, uts, uas float64
}

type data struct {
	grade, nama, nim, kelas, kelamin, lahir, prodi, faculty string
	semester, sks                                           int64
	nilai                                                   float64
}

type tabdata [NMAX]data

type jumlahmatkul [NKUL]sidekuliah

type tabmatkul [NKUL][NKUL]kuliah

func header() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("   Selamat datang di aplikasi nilai akademik   ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Login sebagai : ")
	fmt.Println("1. Layanan akademik")
	fmt.Println("2. Log Out")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Pilih opsi 1 / 2 ? ")
}

func verif(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {

	var pasword string

	//ini dijadiin untuk seleksi admin dlll, berarti pake if
	for {
		fmt.Print("Masukkan Pasword Admin : ")
		fmt.Scan(&pasword)
		if pasword == "apa" {
			break
		}
		fmt.Println("Pasword anda salah")
		fmt.Println("-----------------------------")
	}
	admin(a, e, n, matkulN, j)
}

func main() {
	var a tabdata
	var e tabmatkul
	var j jumlahmatkul
	var n, matkulN int64
	var x int64
	//tampilan login
	header()
	//input x untuk menu
	//tipe kalau tidak bisa maka bikin preload
	for {
		fmt.Scan(&x)
		if x == 1 {
			verif(&a, &e, &n, &matkulN, &j)
			break
		} else if x == 2 {
			break
		}
		fmt.Println("maaf input anda tidak benar, ulangi inputan")
		fmt.Println("-----------------------------------------------")
		fmt.Print("Pilih opsi 1 / 2  ? ")
	}
	//verifikasi

}
