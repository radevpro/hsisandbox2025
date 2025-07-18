package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	// Variables
	var readerSuhu io.Reader = os.Stdin
	bufReader := bufio.NewReader(readerSuhu)
	var suhu_reamur float64
	var suhu_fahrenheit float64

	// Constants
	const konversi_reamur = 4.0 / 5.0
	const konversi_fahrenheit_a = 9.0 / 5.0
	const konversi_fahrenheit_b = 32

	fmt.Println("--- Konverter Suhu ---")
	fmt.Println("Masukkan Suhu dalam Celcius:")

	// Baca input sebagai string
	input, err := bufReader.ReadString('\n')
	if err != nil {
		fmt.Println("Gagal membaca input:", err)
		return
	}

	// Konversi string ke float64
	suhu_celcius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid! Harap masukkan angka.")
		return
	}

	//Konversi ke Reamur
	suhu_reamur = konversi_reamur * suhu_celcius
	fmt.Printf("Suhu dalam Reamur = %.2f\n", suhu_reamur)

	//Konversi ke Fahrenheit
	suhu_fahrenheit = (suhu_celcius * konversi_fahrenheit_a) + konversi_fahrenheit_b
	fmt.Printf("Suhu dalam Fahrenheit= %.2f\n", suhu_fahrenheit)

}
