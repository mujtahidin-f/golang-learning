package main

// 1 project min package main, package main akan di eksekusi pertama kali

import (
	"fmt"
	// "math/bits"
)

// func main() {
// 	fmt.Println("Hello", "World")
// 	fmt.Print("Hello World\n")
// 	fmt.Printf("Hello World")
// }

//  func main merupakan fungsi yg pertama kali di eksekusi

// func main() {
// 	var nama_depan string = "John"

// 	var nama_belakang string
// 	nama_belakang = "Lennon"

// 	fmt.Println("halo ", nama_depan, nama_belakang, "!")
// }

// ---------------------------------
//       Kalkulator sederhana
// ---------------------------------
// func main() {
// 	var nominal1 int
// 	fmt.Println("Angka 1 : ")
// 	fmt.Scanln(&nominal1)

// 	var nominal2 int
// 	fmt.Println("Angka 2 : ")
// 	fmt.Scanln(&nominal2)

// 	var jumlah int = nominal1 + nominal2
// 	var pengurangan int = nominal1 - nominal2
// 	var perkalian int = nominal1 * nominal2
// 	var pembagian int = nominal1 / nominal2
// 	fmt.Println("Hasil Penjumlahan : ", jumlah)
// 	fmt.Println("Hasil Pengurangan : ", pengurangan)
// 	fmt.Println("Hasil Perkalian   : ", perkalian)
// 	fmt.Println("Hasil Pembagian   : ", pembagian)
// }

// ---------------------------------
//         Kalkulator Modif
// ---------------------------------
func main() {
	fmt.Println("---------------------------------")
	fmt.Println("               MENU              ")
	fmt.Println("---------------------------------")
	var list int
	fmt.Println("1. Kalkulator")
	fmt.Println("2. Luas Lingkaran")
	fmt.Println("3. Luas Segitiga")
	fmt.Scanln(&list)

	switch list {
	case 1:
		fmt.Println("---------------------------------")
		fmt.Println("            Kalkulator           ")
		fmt.Println("---------------------------------")
		var pilih int
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Exit")
		fmt.Println("Pilih operator matematika :")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Println("---------------------------------")
			fmt.Println("       Operator Penjumlahan      ")
			fmt.Println("---------------------------------")
			var a int
			var b int
			fmt.Println("Angka Pertama : ")
			fmt.Scanln(&a)

			fmt.Println("Angka Kedua   : ")
			fmt.Scanln(&b)

			var jumlah int = a + b
			fmt.Println("Hasil Penjumlahan : ", jumlah)

		case 2:
			fmt.Println("---------------------------------")
			fmt.Println("       Operator Pengurangan      ")
			fmt.Println("---------------------------------")
			var a int
			var b int

			fmt.Println("Angka Pertama : ")
			fmt.Scanln(&a)

			fmt.Println("Angka Kedua   : ")
			fmt.Scanln(&b)

			var kurang int = a - b
			fmt.Println("Hasil Pengurangan : ", kurang)

		case 3:
			fmt.Println("---------------------------------")
			fmt.Println("        Operator Perkalian       ")
			fmt.Println("---------------------------------")
			var a int
			var b int

			fmt.Println("Angka Pertama : ")
			fmt.Scanln(&a)

			fmt.Println("Angka Kedua   : ")
			fmt.Scanln(&b)

			var kali int = a * b
			fmt.Println("Hasil Perkalian : ", kali)

		case 4:
			fmt.Println("---------------------------------")
			fmt.Println("        Operator Pembagian       ")
			fmt.Println("---------------------------------")
			var a int
			var b int

			fmt.Println("Angka Pertama : ")
			fmt.Scanln(&a)

			fmt.Println("Angka Kedua   : ")
			fmt.Scanln(&b)

			var bagi int = a / b
			fmt.Println("Hasil Perkalian : ", bagi)

		case 5:
			fmt.Println("---------------------------------")
			fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")

		default:
			fmt.Println("Keluar Aplikasi")
		}
	case 2:
		fmt.Println("---------------------------------")
		fmt.Println("          Luas Lingkaran         ")
		fmt.Println("---------------------------------")
		fmt.Println("Masukkan jari-jari lingkaran :")
		var jarijari float32
		fmt.Scanln(&jarijari)
		var p float32 = 3.143

		var luaslingkaran float32 = p * jarijari * jarijari
		fmt.Println("Luas Lingkaran : ", luaslingkaran)

	case 3:
		fmt.Println("---------------------------------")
		fmt.Println("          Luas Segitiga          ")
		fmt.Println("---------------------------------")
		fmt.Println("Masukkan alas   : ")
		fmt.Println("Masukkan Tinggi : ")

		var alassegitigas float32
		fmt.Scanln(&alassegitigas)
		var tinggisegitiga float32
		fmt.Scanln(&tinggisegitiga)

		var luassegitiga float32 = 0.5 * alassegitigas * tinggisegitiga
		fmt.Println("Luas segitiga : ", luassegitiga)
	}
}

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	var positiveNumber uint8 = 89
// 	var negativeNumber = -1243423644
// 	fmt.Printf("bilangan positif: %d\n", positiveNumber)
// 	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
// }
