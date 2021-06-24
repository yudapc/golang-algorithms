// recover digunakan untuk melanjutkan program
// apabila terjadi panic
// ini mirip seperti throw

package deferpanicrecover

import "fmt"

func EndFuncRecover() {
	message := recover()
	if message != nil {
		fmt.Println("ERROR message: ", message)
	}
}

func RunFuncRecover(error bool) {
	defer EndFuncRecover()
	if error == true {
		panic("ERROR run function")
	}
}
