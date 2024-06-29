package main
// package main digunakan untuk mendefinisikan package utama

// import digunakan untuk mengimpor modul-modul yang diperlukan
import (
	"fmt"
	"time"
)

// printNumbers adalah fungsi yang mencetak angka dari 0 hingga 9 dengan interval waktu 1 milidetik
func printNumbers(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
}

// printLetters adalah fungsi yang mencetak huruf dari 'a' hingga 'e' dengan interval waktu 1 milidetik
func printLetters(){
	for i := 'a'; i <= 'e'; i++ {
		fmt.Println("%c\n", i)
		time.Sleep(1 * time.Millisecond)
	}
}

// main adalah fungsi utama yang akan dieksekusi saat program dijalankan
func main(){
	// go printNumbers() digunakan untuk menjalankan fungsi printNumbers secara concurrent
	go printNumbers()
	// go printLetters() digunakan untuk menjalankan fungsi printLetters secara concurrent
	go printLetters()
	
	// time.Sleep(11 * time.Millisecond) digunakan untuk memberhentikan program selama 11 milidetik
	time.Sleep(10 * time.Millisecond)
}
