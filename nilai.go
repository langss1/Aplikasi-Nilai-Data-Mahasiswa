package main

import "fmt"

func add_nilai_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, o int64
	var z string
	fmt.Print("Masukkan Matakuliah: ")
	fmt.Scan(&z)
	o = 0
	//bikin variabel check, buat menentukan if else jika tidak ada matkul
	for o < *matkulN {
		i = 0
		for i < *n {
			//bikin fungsi untuk melanjutkan input nilai dan edit nilai
			if e[i][o].matkul == z {
				fmt.Println("-----------------------------------------------")
				fmt.Println("Mahasiswa :", a[i].nama)
				fmt.Print("Masukkan Nilai Quiz: ")
				fmt.Scan(&e[i][o].quiz)
				fmt.Print("Masukkan Nilai Uts: ")
				fmt.Scan(&e[i][o].uts)
				fmt.Print("Masukkan Nilai Uas: ")
				fmt.Scan(&e[i][o].uas)
				e[i][o].nilai = (e[i][o].quiz + e[i][o].uts + e[i][o].uas) / 3.0

				fmt.Println("Total Nilai matkul", e[i][o].matkul, "=", e[i][o].nilai)
				fmt.Println("-----------------------------------------------")

				if e[i][o].nilai > 80 {
					e[i][o].grade = "a"
				} else if e[i][o].nilai > 70 && e[i][o].nilai <= 80 {
					e[i][o].grade = "ab"
				} else if e[i][o].nilai > 60 && e[i][o].nilai <= 70 {
					e[i][o].grade = "b"
				} else if e[i][o].nilai > 55 && e[i][o].nilai <= 60 {
					e[i][o].grade = "bc"
				} else if e[i][o].nilai > 50 && e[i][o].nilai <= 55 {
					e[i][o].grade = "c"
				} else if e[i][o].nilai > 45 && e[i][o].nilai <= 50 {
					e[i][o].grade = "d"
				} else if e[i][o].nilai <= 45 {
					e[i][o].grade = "e"
				}

				a[i].nilai += e[i][o].nilai
				j[i].nilai += 1
				a[i].nilai = a[i].nilai / j[i].nilai
			}
			i += 1
		}
		o += 1
	}
	admin(a, e, n, matkulN, j)
}
