package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Scan(&kelas, &nama, &nim)
	fmt.Print("Perkenalkan saya adalah ", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim)
}
