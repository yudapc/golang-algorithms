package interfaceimplementation

import (
	"fmt"
	"testing"
)

func TestBangunDatarPersegi(t *testing.T) {
	fmt.Println("=======Persegi=======")
	var bangunDatar hitung = persegi{10.0}
	fmt.Println("Luas = ", bangunDatar.luas())
	fmt.Println("Keliling = ", bangunDatar.keliling())
}

func TestBangunDatarLingkaran(t *testing.T) {
	fmt.Println("=======Lingkaran=======")
	var bangunDatar hitung = lingkaran{40.0}
	fmt.Println("Luas = ", bangunDatar.luas())
	fmt.Println("Keliling = ", bangunDatar.keliling())
}
