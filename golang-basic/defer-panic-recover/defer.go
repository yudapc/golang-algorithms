// Defer function digunakan
// untuk mengeksekusi fungsi ketika selesai dijalankan
// baik fungsi tersebut error atau tidak
package deferpanicrecover

import "fmt"

func EndFuncDefer() {
	fmt.Println("Akhir Aplikasi")
}

func RunFuncDefer(error bool) {
	defer EndFuncDefer()
	if error == true {
		fmt.Println("Aplikasi ERROR")
		panic("ERROR APLIKASI")
	}
	fmt.Println("Aplikasi Berjalan")
}
