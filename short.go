package main

import "fmt"

func header_shorting() {
	fmt.Println("1. Berdasarkan nilai Mahasiswa")
	fmt.Println("2. Berdasarkan Jumlah SKS Mahasiswa")
	fmt.Println("-----------------------------------------------")
}

func shorting(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	header_shorting()
	var x int64
	for {
		fmt.Print("Choose 1 atau 2? ")
		fmt.Scan(&x)
		if x == 1 {
			shorting_nilai(a, e, n, matkulN, j)
			break
		} else if x == 2 {
			shorting_sks(a, e, n, matkulN, j)
			break
		}
	}
}

func shorting_sks(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x int64
	fmt.Println("1. Berdasarkan jumlah SKS terendah ke terbesar")
	fmt.Println("2. Berdasarkan jumlah SKS terbesar ke terkecil")
	fmt.Println("-----------------------------------------------")
	for {
		fmt.Print("Choose 1 atau 2?")
		fmt.Scan(&x)
		if x == 2 {
			descending_sks(a, e, n, matkulN, j)
			break
		} else if x == 1 {
			ascending_sks(a, e, n, matkulN, j)
			break
		}
	}

}

func ascending_sks(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var pass, i, o int64
	var temp data
	var tam kuliah
	pass = 1
	for pass <= *n-1 {
		i = pass
		temp = a[pass]
		tam = e[pass][o]
		for i > 0 && temp.sks < a[i-1].sks {
			a[i] = a[i-1]
			for o < *matkulN {
				e[i][o] = e[i-1][o]
				o += 1
			}
			i -= 1
		}
		a[i] = temp
		for o < *matkulN {
			e[i][o] = tam
			o += 1
		}
		pass += 1
	}

	fmt.Println("------------------------------------------------------")
	fmt.Println("Mahasiswa berdasarkan jumlah SKS terbesar ke terkecil")
	i = 0
	for i < *n {
		fmt.Println(i+1, a[i].nama, "=", a[i].sks)
		i++
	}
}

func descending_sks(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var pass, idx, i, o int64
	var temp data
	var tam kuliah
	pass = 1
	for pass <= *n-1 {
		o = 0
		idx = pass - 1
		i = pass
		for i < *n {
			if a[idx].sks < a[i].sks {
				idx = i
			}
			i += 1
		}
		temp = a[pass-1]
		a[pass-1] = a[idx]
		a[idx] = temp

		for o < *matkulN {
			tam = e[pass-1][o]
			e[pass-1][o] = e[idx][o]
			e[idx][o] = tam
			o += 1
		}
		pass += 1

		fmt.Println("------------------------------------------------------")
		fmt.Println("Mahasiswa berdasarkan sks terbesar ke terkecil")
		i = 0
		for i < *n {
			fmt.Println(i+1, a[i].nama, "=", a[i].sks)
			i++
		}
	}
}

func shorting_nilai(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	fmt.Println("1. Berdasarkan terendah ke terbesar")
	fmt.Println("2. Berdasarkan terbesar ke terkecil")
	var x int64
	for {
		fmt.Print("Choose 1 atau 2?")
		fmt.Scan(&x)
		if x == 2 {
			descending_nilai(a, e, n, matkulN, j)
			break
		} else if x == 1 {
			ascending_nilai(a, e, n, matkulN, j)
			break
		}
	}
}

func ascending_nilai(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var pass, i, o int64
	var temp data
	var tam kuliah
	pass = 1
	for pass <= *n-1 {
		i = pass
		temp = a[pass]
		tam = e[pass][o]
		for i > 0 && temp.nilai < a[i-1].nilai {
			a[i] = a[i-1]
			for o < *matkulN {
				e[i][o] = e[i-1][o]
				o += 1
			}
			i -= 1
		}
		a[i] = temp
		for o < *matkulN {
			e[i][o] = tam
			o += 1
		}
		pass += 1
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println("Mahasiswa berdasarkan nilai terkecil ke terbesar")
	i = 0
	for i < *n {
		fmt.Println(i+1, a[i].nama, "=", a[i].nilai)
		i++
	}
}

func descending_nilai(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var pass, idx, i, o int64
	var temp data
	var tam kuliah
	pass = 1
	for pass <= *n-1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if a[idx].nilai < a[i].nilai {
				idx = i
			}
			i += 1
		}
		temp = a[pass-1]
		a[pass-1] = a[idx]
		a[idx] = temp
		for o < *matkulN {
			tam = e[pass-1][o]
			e[pass-1][o] = e[idx][o]
			e[idx][o] = tam
			o += 1
		}
		pass += 1
	}

	fmt.Println("------------------------------------------------------")
	fmt.Println("Mahasiswa berdasarkan nilai terbesar ke terkecil")
	i = 0
	for i < *n {
		fmt.Println(i+1, a[i].nama, "=", a[i].nilai)
		i++
	}
}
