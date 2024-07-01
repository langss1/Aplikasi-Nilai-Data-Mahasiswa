package main

import "fmt"

func header_matkul_mahasiswa() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("Pilih Menu: ")
	fmt.Println("1. Add mata kuliah baru")
	fmt.Println("2. Edit mata kuliah")
	fmt.Println("3. Remove mata kuliah")
	fmt.Println("4. Liat data")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Pilih opsi 1 / 2 / 3 / 4 / 5 ? ")
}

func matkul_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x int64
	header_matkul_mahasiswa()
	for {
		fmt.Scan(&x)
		if x == 1 {
			add_matkul_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 2 {
			edit_matkul_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 3 {
			remove_matkul_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 4 {
			cetak_matkul_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 5 {
			admin(a, e, n, matkulN, j)
			break
		}
		fmt.Print("Pilih opsi 1 / 2 / 3 / 4 / 5 ? ")
	}
}

func add_matkul_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, x, b, o int64
	var z string
	i = 0
	if *n == 0 {
		fmt.Println("anda belum menginput mahasiswa")
		admin(a, e, n, matkulN, j)
	}
	o = *matkulN
	fmt.Print("Nama Mata Kuliah: ")
	fmt.Scan(&z)
	fmt.Print("Besaran Proporsi SKS: ")
	fmt.Scan(&b)

	for i < *n && x != 3 {
		e[i][o].tru = false
		fmt.Println("Nama Mahasiswa:", a[i].nama, "saat ini memiliki", a[i].sks, "sks")
		fmt.Println("Apakah anda ingin menambahkan matkul", z, "ke mahasiswa ini?")
		fmt.Println("1. setuju")
		fmt.Println("2. tidak")
		fmt.Println("3. Exit")
		fmt.Print("Pilih Opsi 1/2/3: ")
		fmt.Scan(&x)
		if x == 1 {
			if a[i].sks+b <= 24 {
				e[i][o].nama = a[i].nama
				e[i][o].tru = true
				e[i][o].matkul = z
				e[i][o].kredit = b
				e[i][o].grade = "e"
				a[i].sks = a[i].sks + b
				j[o].jumlah += 1
			} else {
				fmt.Println("maaf ananda", a[i].nama, "lebih dari 24 sks")
			}
		}
		i += 1
	}
	*matkulN += 1
	matkul_mahasiswa(a, e, n, matkulN, j)
}

func edit_matkul_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, o, x int64
	var check, warn bool
	var z, y string

	warn = false

	fmt.Print("cari mata kuliah yang akan diedit: ")
	fmt.Scan(&z)

	check = false
	for {
		i = 0
		for i < *n {
			if e[i][o].matkul == z {
				check = true
				break
			}
			i += 1
		}
		o += 1
		if o == *matkulN {
			break
		}
	}

	if check {
		fmt.Print("Nama mata kuliah revisi: ")
		fmt.Scan(&y)
		fmt.Print("masukan sks terbaru: ")
		fmt.Scan(&x)
		o = 0
		for o < *matkulN {
			i = 0
			for i < *n {
				if e[i][o].matkul == z {
					warn = true
					if (a[i].sks-e[i][o].kredit)+x <= 24 {
						a[i].sks -= e[i][o].kredit
						e[i][o].matkul = y
						e[i][o].kredit = x
						//warn = true
					} else {
						//dihapus atau keluar daftar
					}
				}
				i += 1
			}
			o += 1
		}

	} else {
		fmt.Println("mata kuliah yang dicari tidak ada")
	}

	if warn {
		fmt.Println("pengubahan matkul berhasil")
	} else {
		fmt.Println("tidak ditemukan mata kuliah")
	}
	matkul_mahasiswa(a, e, n, matkulN, j)
}

func remove_matkul_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var x string
	var u, o, i int64
	var brik bool
	fmt.Print("cari mata kuliah yang akan dihapus: ")
	fmt.Scan(&x)
	brik = false

	for {
		i = 0
		if o == *matkulN || brik {
			break
		}
		for i < *n {
			if e[i][o].matkul == x {
				u = 0
				j[o].jumlah = 0
				*matkulN -= 1
				brik = true
				for u < *n {
					e[u][o].matkul = e[u][o+1].matkul
					u += 1
				}
			}
			i += 1
		}
		o += 1
	}

	if !brik {
		fmt.Println("tidak ada matkul yang dihapus")
	} else {
		fmt.Println("Penghapusan berhasil")
	}
	matkul_mahasiswa(a, e, n, matkulN, j)
}

func cetak_matkul_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, x, o int64
	var z string
	var check bool
	i = 0
	fmt.Print("Data Matkul: ")
	fmt.Scan(&z)
	fmt.Print("list mengambil Data Matkul: ", z)
	fmt.Println()
	for o < *matkulN {
		i = 0
		for i < *n {
			if e[i][o].matkul == z {
				fmt.Println(e[i][o].nama)
				check = true
			}
			i += 1
		}
		o += 1
	}
	if !check {
		fmt.Println("Tidak ditemukan adanya matkul")
	}
	fmt.Println("----------------------")
	for {
		fmt.Print("ketik 3 untuk kembali ke data matkul: ")
		fmt.Scan(&x)
		if x == 3 {
			matkul_mahasiswa(a, e, n, matkulN, j)
			break
		}
	}
}
