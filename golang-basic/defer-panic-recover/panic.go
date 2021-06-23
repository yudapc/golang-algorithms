// Panic digunakan untuk menghentikan program
// Biasanya panic digunakan ketika
// program error

package deferpanicrecover

func RunFuncPanic(error bool) {
	if error == true {
		panic("APLIKASI ERROR")
	}
}
