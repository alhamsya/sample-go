package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	/*Fungsi runtime.GOMAXPROCS(n) digunakan untuk menentukan jumlah core yang diaktifkan untuk eksekusi program.*/
	// OPTIONAL for using syntax this one
	runtime.GOMAXPROCS(10)

	//sync.WaitGroup digunakan untuk menunggu goroutine
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		//inisialisasi variabel string
		var data = fmt.Sprintf("data %d", i)

		//Di tiap perulangan statement wg.Add(1) dipanggil. Kode tersebut akan memberikan informasi
		//kepada wg bahwa jumlah goroutine yang sedang diproses ditambah 1 (karena dipanggil 5 kali
		//maka wg akan sadar bahwa terdapat 5 buah goroutine sedang berjalan).
		wg.Add(1)

		//fungsi doPrint() dieksekusi sebagai goroutine. Didalam fungsi tersebut,
		//sebuah method bernama Done() dipanggil. Method ini digunakan untuk memberikan informasi kepada wg
		//bahwa goroutine dimana method itu dipanggil sudah selesai. Sejumlah 5 buah goroutine dijalankan,
		//maka method tersebut harus dipanggil 5 kali.
		go doPrint(&wg, data)
	}

	//Statement wg.Wait() bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai.
	//Jika Add(1) dipanggil 5 kali, maka Done() juga harus dipanggil 5 kali.
	wg.Wait()
}

func doPrint(wg *sync.WaitGroup, message string) {
	// akan dieksekusi diakhir
	defer wg.Done()

	// print message
	fmt.Println(message)
}
