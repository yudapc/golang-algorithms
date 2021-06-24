// Error interface adalah sebuah kontrak
// untuk membuat error, nama interface nya adalah error
// kita menggunakan package errors bawaan golang
// caranya dengan errors.New("")
package errorinterface

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	result := nilai / pembagi
	return result, nil
}

func RunSampleError() {
	result, err := Pembagian(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
