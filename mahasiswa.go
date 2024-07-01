package main

import "fmt"

func mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x, y string
	var i int64
	fmt.Print("Masukkan nama Mahasiswa: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scan(&y)
	i = sequential_search(a, n, x, y)
	if i == -1 {
		fmt.Println("data yang dicari tidak ada")
	} else {
		transkrip(a, e, n, matkulN, j, i)
	}
}

func transkrip(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul, i int64) {
	var o int64
	fmt.Println("-----------------------------------------------")
	fmt.Println("               Transkrip Mahasiswa             ")
	fmt.Println("-----------------------------------------------")
	fmt.Printf("%15s %5s %5s\n", "Mata Kuliah", "Nilai Mutu", "Nilai Index")
	fmt.Println("-----------------------------------------------")
	for o < *matkulN {
		if e[i][o].tru {
			fmt.Printf("%15s %5.1f %5s\n", e[i][o].matkul, e[i][o].nilai, e[i][o].grade)
			fmt.Println("-----------------------------------------------")
		}
		o += 1

	}
}
