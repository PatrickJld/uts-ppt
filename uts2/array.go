package main

import (
	"fmt"
)

func main() {

	//declare variabel
	arr1 := [3]string{}
	arr2 := [3]string{}

	//looping untuk proses input data
	//array 1
	fmt.Print("Masukkan data untuk array 1: ")
	for i := 0; i < len(arr1); i++ {
		fmt.Scan(&arr1[i])
	}

	//array 2
	fmt.Print("Masukkan data untuk array 2: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Scan(&arr2[i])
	}

	//logic check dengan menggunakan boolean
	Check := true

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			fmt.Printf("Indeks array ke %d tidak sama", i)
			fmt.Println()

			Check = false
		}
	}

	if Check {
		fmt.Println("Kedua Array memiliki indeks yang sama")
	}
}
