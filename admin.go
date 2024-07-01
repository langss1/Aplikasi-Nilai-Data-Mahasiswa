package main

import "fmt"

func header_admin() {
	//untuk menu tampilan utama admin
	fmt.Println("-----------------------------------------------")
	fmt.Println("Anda terlogin sebagai admin")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Menu Admin: ")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Mata Kuliah")
	fmt.Println("3. Add Nilai")
	fmt.Println("4. Cari Mahasiswa")
	fmt.Println("5. Transkrip Mahasiswa")
	fmt.Println("6. Pengurutan Nilai Mahasiswa")
	fmt.Println("7. Exit")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Pilih opsi 1 / 2 / 3 / 4 / 6 / 7 ? ")
}

func admin(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	//untuk pilihan menu admin
	var x int64
	header_admin()
	for {
		fmt.Scan(&x)
		if x == 1 {
			data_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 2 {
			matkul_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 5 {
			mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 4 {
			search_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 3 {
			add_nilai_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 6 {
			shorting(a, e, n, matkulN, j)
			break
		} else if x == 7 {
			main()
			break
		}
		fmt.Print("Pilih opsi 1 / 2 / 3 / 4 / 5 / 6 / 7 ? ")
	}
}

func header_data_mahasiswa() {
	//untuk menu header data mahasiswa
	fmt.Println("-----------------------------------------------")
	fmt.Println("Pilih Menu: ")
	fmt.Println("1. Add")
	fmt.Println("2. Edit")
	fmt.Println("3. Remove")
	fmt.Println("4. Liat data")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Pilih opsi 1 / 2 3 / 4 / 5? ")
}

func data_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	//untuk memasukan, mengedit, remove data mahasiswa
	var x int64
	header_data_mahasiswa()
	for {
		fmt.Scan(&x)
		if x == 1 {
			add_data_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 2 {
			edit_data_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 3 {
			remove_data_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 4 {
			cetak_data_mahasiswa(a, e, n, matkulN, j)
			break
		} else if x == 5 {
			admin(a, e, n, matkulN, j)
			break
		}
		fmt.Print("Pilih opsi 1 / 2 / 4 / 5? ")
	}
	admin(a, e, n, matkulN, j)
}

func kelamin(a *tabdata, i int64) {
	//untuk memasukkan kelamin mahasiswa
	fmt.Println("Masukan kelamin Mahasiswa")
	fmt.Println("1. Laki - laki")
	fmt.Println("2. Perempuan")
	var x int64
	for {
		fmt.Print("Choose 1 / 2? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].kelamin = "Laki - Laki"
			break
		} else if x == 2 {
			a[i].kelamin = "Perempuan"
			break
		}
	}
}

func prodi_if(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Informatika")
	fmt.Println("2. S1 Teknologi Informasi")
	fmt.Println("3. S1 Rekayasa Perangkat Lunak")
	fmt.Println("4. S1 Data Sains")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Informatika"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Teknologi Informasi"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Rekayasa Perangkat Lunak"
			break
		} else if x == 4 {
			a[i].prodi = "S1 Data Sains"
			break
		}

	}
}

func prodi_te(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Teknik Telekomunikasi")
	fmt.Println("2. S1 Teknik Elektro")
	fmt.Println("3. S1 Teknik Komputer")
	fmt.Println("4. S1 Teknik Biomedis")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Teknik Telekomunikasi"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Teknik Elektro"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Teknik Komputer"
			break
		} else if x == 4 {
			a[i].prodi = "S1 Teknik Biomedis"
			break
		}

	}
}

func prodi_ri(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Teknik Industri")
	fmt.Println("2. S1 Sistem Informasi")
	fmt.Println("3. S1 Teknik Logistik")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Teknik Industri"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Sistem Informasi"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Teknik Logistik"
			break
		}

	}
}

func prodi_eb(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Manajemen")
	fmt.Println("2. S1 Akuntansi")
	fmt.Println("3. S1 Manajemen Bisnis Rekreasi")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Manajemen"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Akuntansi"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Manajemen Bisnis Rekreasi"
			break
		}

	}
}

func prodi_kb(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Adminitrasi Bisnis")
	fmt.Println("2. S1 Ilmu Komunikasi")
	fmt.Println("3. S1 Hubungan Masyarakat Digital")
	fmt.Println("4. S1 Penyiaran Digital")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Adminitrasi Bisnis"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Ilmu Komunikasi"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Hubungan Masyarakat Digital"
			break
		} else if x == 4 {
			a[i].prodi = "S1 Penyiaran Digital"
			break
		}

	}
}

func prodi_ik(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. S1 Desain Komunikasi Visual")
	fmt.Println("2. S1 Desain Produk")
	fmt.Println("3. S1 Desain Interior")
	fmt.Println("4. S1 Seni Rupa")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "S1 Desain Komunikasi Visual"
			break
		} else if x == 2 {
			a[i].prodi = "S1 Desain Produk"
			break
		} else if x == 3 {
			a[i].prodi = "S1 Desain Interior"
			break
		} else if x == 4 {
			a[i].prodi = "S1 Seni Rupa"
			break
		}

	}
}

func prodi_it(a *tabdata, i int64) {
	fmt.Println("Pilih Prodi: ")
	fmt.Println("1. D3 Teknik Informatika")
	fmt.Println("2. D3 Teknologi Komputer")
	fmt.Println("3. D3 Sistem Informasi")
	fmt.Println("4. D3 Perhotelan")
	var x int64
	for {
		fmt.Print("Pilih prodi? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].prodi = "D3 Teknik Informatika"
			break
		} else if x == 2 {
			a[i].prodi = "D3 Teknologi Komputer"
			break
		} else if x == 3 {
			a[i].prodi = "D3 Sistem Informasi"
			break
		} else if x == 4 {
			a[i].prodi = "D3 Perhotelan"
			break
		}

	}
}

func jurusan(a *tabdata, i int64) {
	var x int64
	fmt.Println("Masukkan Fakultas: ")
	fmt.Println("1. Fakultas Informatika")
	fmt.Println("2. Fakultas Teknik elektro")
	fmt.Println("3. Fakultas Rekayasa Industri")
	fmt.Println("4. Fakultas Ekonomi Bisnis")
	fmt.Println("5. Fakultas Komunikasi Bisnis")
	fmt.Println("6. Fakultas Industri Kreatif")
	fmt.Println("7. Fakultas Ilmu Terapan")
	fmt.Println("--------------------------------")
	for {
		fmt.Print("Pilih Fakultas? ")
		fmt.Scan(&x)
		if x == 1 {
			a[i].faculty = "Informatika"
			prodi_if(a, i)
			break
		} else if x == 2 {
			a[i].faculty = "Teknik Elektro"
			prodi_te(a, i)
			break
		} else if x == 3 {
			a[i].faculty = "Rekayasa Indsutri"
			prodi_ri(a, i)
			break
		} else if x == 4 {
			a[i].faculty = "Ekonomi Bisnis"
			prodi_eb(a, i)
			break
		} else if x == 5 {
			a[i].faculty = "Komunikasi Bisnis"
			prodi_kb(a, i)
			break
		} else if x == 6 {
			a[i].faculty = "Industri Kreatif"
			prodi_ik(a, i)
			break
		} else if x == 7 {
			a[i].faculty = "Ilmu Terapan"
			prodi_it(a, i)
			break
		}
	}

}

func add_data_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	//untuk memasukan data mahasiswa
	var i, x, b int64
	var warn bool
	var check string
	if *n == 0 {
		for {
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&a[i].nama)
			warn = true
			for warn {
				fmt.Print("Masukkan NIM: ")
				fmt.Scan(&a[i].nim)
				check = a[i].nim
				b = 0
				warn = false

				if i >= 1 {
					for b < i {
						if a[b].nim == check {
							warn = true
							fmt.Println("NIM sudah digunakan, silakan ulangi.")
							break
						}
						b++
					}
				}
			}
			kelamin(a, i)
			fmt.Print("Masukkan tahun lahir: ")
			fmt.Scan(&a[i].lahir)
			jurusan(a, i)
			fmt.Print("Masukkan kelas: ")
			fmt.Scan(&a[i].kelas)
			fmt.Print("Masukkan semester: ")
			fmt.Scan(&a[i].semester)

			i += 1
			fmt.Print("Ketik 1 jika lanjut, ketik 2 jika udahan: ")
			fmt.Scan(&x)
			if x == 2 {
				break
			}
			fmt.Println("-----------------------------------------------")

		}
	}
	if *n > 0 {
		i = *n
		for {
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&a[i].nama)
			warn = true
			for warn {
				fmt.Print("Masukkan NIM: ")
				fmt.Scan(&a[i].nim)
				check = a[i].nim
				b = 0
				warn = false

				if i >= 1 {
					for b < i {
						if a[b].nim == check {
							warn = true
							fmt.Println("NIM sudah digunakan, silakan ulangi.")
							break
						}
						b++
					}
				}
			}
			kelamin(a, i)
			fmt.Print("Masukkan tahun lahir: ")
			fmt.Scan(&a[i].lahir)
			jurusan(a, i)
			fmt.Print("Masukkan kelas: ")
			fmt.Scan(&a[i].kelas)
			fmt.Print("Masukkan semester: ")
			fmt.Scan(&a[i].semester)
			i += 1

			fmt.Print("Ketik 1 jika lanjut, ketik 2 jika udahan: ")
			fmt.Scan(&x)
			if x == 2 {
				break
			}
			fmt.Println("-----------------------------------------------")
		}

	}
	*n = i
	data_mahasiswa(a, e, n, matkulN, j)
}

func edit_data_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, b int64
	var x, check, c string
	var warn bool
	fmt.Println("-------------------")
	fmt.Println("Fitur edit Mahasiswa")
	fmt.Println("--------------------")

	if *n == 0 {
		fmt.Println("belum ada data sama sekali")
	}

	if *n > 0 {
		fmt.Print("Cari nama mahasiswa: ")
		fmt.Scan(&x)
		fmt.Print("Cari NIM mahasiswa: ")
		fmt.Scan(&c)
		for {
			if a[i].nama == x && a[i].nim == c {
				fmt.Print("Masukkan nama: ")
				fmt.Scan(&a[i].nama)
				warn = true
				for warn {
					fmt.Print("Masukkan NIM: ")
					fmt.Scan(&a[i].nim)
					check = a[i].nim
					b = 0
					warn = false

					for b < *n {
						if b != i {
							if a[b].nim == check {
								warn = true
								fmt.Println("NIM sudah digunakan, silakan ulangi.")
								break
							}
						}
						b++
					}

				}
				kelamin(a, i)
				fmt.Print("Masukkan tahun lahir: ")
				fmt.Scan(&a[i].lahir)
				jurusan(a, i)
				fmt.Print("Masukkan kelas: ")
				fmt.Scan(&a[i].kelas)
				fmt.Print("Masukkan semester: ")
				fmt.Scan(&a[i].semester)
				break
			} else if i == *n {
				fmt.Println("maaf mahasiswa yang dicari tidak ada")
				break

			}
			i += 1
		}
	}
	data_mahasiswa(a, e, n, matkulN, j)
}

func remove_data_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, u, o int64
	var x, y string

	//tambahkan nama dan nim buat penghapusan

	if *n == 0 {
		fmt.Println("belum ada data sama sekali")
	}

	if *n > 0 {
		fmt.Print("Cari nama yang akan dihapus:")
		fmt.Scan(&x)
		fmt.Print("Nim yang akan dihapus: ")
		fmt.Scan(&y)
		for {
			if a[i].nama == x && a[i].nim == y {

				for o < *matkulN {
					u = i
					for u < *n {
						e[u][o] = e[u+1][o]
						u += 1
					}
					o += 1
				}

				for i < *n-1 {
					a[i] = a[i+1]
					i += 1
				}
				*n -= 1

				fmt.Println("Data yang dihapus berhasil")
				break
			} else if i == *n {
				fmt.Println("Nama yang akan dihapus tidak ada")
				break
			}
			i += 1
		}
	}
	data_mahasiswa(a, e, n, matkulN, j)
}

func cetak_data_mahasiswa(a *tabdata, e *tabmatkul, n, matkulN *int64, j *jumlahmatkul) {
	var i, x int64
	if *n == 0 {
		fmt.Println("-----------------------------")
		fmt.Println("Belum ada data sama sekali")
	}
	if *n > 0 {
		for {
			fmt.Println("Data Mahasiswa ke", i+1)
			fmt.Println("Nama:", a[i].nama)
			fmt.Println("NIM:", a[i].nim)
			fmt.Println("Jenis Kelamin:", a[i].kelamin)
			fmt.Println("Tahun lahir:", a[i].lahir)
			fmt.Println("Kelas:", a[i].kelas)
			fmt.Println("Sks sebanyak", a[i].sks)
			fmt.Println("Masukkan semester:", a[i].semester)
			fmt.Println("Program Studi:", a[i].prodi)
			fmt.Println("Fakultas", a[i].faculty)

			fmt.Print("Ketik 1 jika lanjut, ketik 2 jika udahan: ")
			fmt.Scan(&x)
			if x == 2 {
				break
			}
			fmt.Println("-----------------------------------------------")
			if i == *n-1 {
				fmt.Println("data sudah maksimal")
				break
			}
			i += 1
		}
	}
	data_mahasiswa(a, e, n, matkulN, j)

}
