package main

import "fmt"

func sequential_search(a *tabdata, n *int64, x, y string) int64 {
	var i, idx int64
	idx = -1
	i = 0
	for i < *n && idx == -1 {
		if a[i].nama == x && a[i].nim == y {
			idx = i
		}
		i++
	}
	return idx
}

func binomial_search(a *tabdata, n *int64, x float64) int64 {
	var left, right, mid, found int64
	left = 0
	right = *n
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < a[mid].nilai {
			right = mid - 1
		} else if x > a[mid].nilai {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func search_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x int64
	fmt.Println("Seacrh Berdasarkan: ")
	fmt.Println("1. Nama & Nim")
	fmt.Println("2. Nilai Mahasiswa")
	fmt.Println("-----------------------------------------------")
	for {
		fmt.Print("Choose 1 atau 2? ")
		fmt.Scan(&x)
		if x == 1 {
			nama_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 2 {
			skor(a, e, n, matkulN, j)
			break
		}
	}
}

func skor(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, o int64
	var x float64
	fmt.Print("masukan nilai: ")
	fmt.Scan(&x)
	i = binomial_search(a, n, x)
	o = 0
	if i == -1 {
		fmt.Println("maaf pencarian tidak ada")
	} else {
		fmt.Println("-----------------------------------------------")
		fmt.Println("                 Result Hasil                  ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Nama :", a[i].nama)
		fmt.Println("Nim :", a[i].nim)
		fmt.Println("Jenis Kelamin :", a[i].kelamin)
		fmt.Println("Tanggal Lahir :", a[i].lahir)
		fmt.Println("Kelas :", a[i].kelas)
		fmt.Println("Semester :", a[i].semester)
		fmt.Println("Jumlah sks :", a[i].sks, "sks")
		fmt.Println("Prodi :", a[i].prodi)
		fmt.Println("Fakultas :", a[i].faculty)
		fmt.Println("mata kuliah diambil :")
		for o < *matkulN {
			if e[i][o].tru {
				fmt.Print("Matkul: ", e[i][o].matkul)
				fmt.Print("Nilai: ", e[i][o].nilai)
				fmt.Print("Grade: ", e[i][o].grade)
				fmt.Println()
			}
			o += 1

		}
	}
}

func nama_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x, y string
	var i, o int64
	fmt.Print("masukan nama: ")
	fmt.Scan(&x)
	fmt.Print("masukan nim: ")
	fmt.Scan(&y)
	i = sequential_search(a, n, x, y)
	o = 0
	if i == -1 {
		fmt.Println("maaf pencarian tidak ada")
	} else {
		fmt.Println("-----------------------------------------------")
		fmt.Println("                 Result Hasil                  ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Nama :", a[i].nama)
		fmt.Println("Nim :", a[i].nim)
		fmt.Println("Jenis Kelamin :", a[i].kelamin)
		fmt.Println("Tanggal Lahir :", a[i].lahir)
		fmt.Println("Kelas :", a[i].kelas)
		fmt.Println("Semester :", a[i].semester)
		fmt.Println("Jumlah sks :", a[i].sks, "sks")
		fmt.Println("Prodi :", a[i].prodi)
		fmt.Println("Fakultas :", a[i].faculty)
		fmt.Println("mata kuliah diambil :")
		for o < *matkulN {
			if e[i][o].tru {
				fmt.Print("Matkul: ", e[i][o].matkul)
				fmt.Print("Nilai: ", e[i][o].nilai)
				fmt.Print("Grade: ", e[i][o].grade)
				fmt.Println()
			}
			o += 1

		}
	}
	admin(a, e, n, matkulN, j)

}
