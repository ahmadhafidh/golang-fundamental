package main

import (
	"fmt"
)

func main() {

	nilai := 70

	switch nilai {
	case 50:
		fmt.Println("Nilai Anda C")

	case 70:
		fmt.Println("Nilai Anda B")
		fallthrough
	case 100:
		fmt.Println("Nilai Anda A")

	default:
		fmt.Println("Tidak ada kategori")
	}
}
