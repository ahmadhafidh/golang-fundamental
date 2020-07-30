package main

import (
	"fmt"
)

func main() {
	nilai := 10

	if nilai > 7 {
		fmt.Println("Anda Lulus dengan Nilai", nilai)
	} else {
		fmt.Println("Anda Tidak Lulus dengan Nilai", nilai)
	}
}