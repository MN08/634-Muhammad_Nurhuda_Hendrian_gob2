package main

import (
	"fmt"
	"os"
	"strconv"
)

// struct Biodata
type Biodata struct {
	no         string
	name       string
	address    string
	job        string
	motivation string
}

func main() {

	//inisialisasi data
	var data = []Biodata{
		{no: "1", name: "Albert", address: "Antapani", job: "Manager", motivation: "Untuk bisa Ngoding"},
		{no: "2", name: "Alberto", address: "Bandung", job: "Designer", motivation: "Untuk bisa Ngoding pake bahasa go"},
		{no: "3", name: "Albertus", address: "Batu", job: "Developer", motivation: "untuk bisa Ngoding bahasa go dengan baik"},
		{no: "4", name: "Albertino", address: "Bekasi", job: "Tester", motivation: "untuk bisa Ngoding pake bahasa go dengan baik dan benar"},
	}

	// os.args digunakan untuk menyimpan data berupa array
	var argsRaw = os.Args
	//mengambil 1 data pada array argsRaw
	var args = argsRaw[1]
	no, err := strconv.Atoi(args)
	_ = err

	//memanggil function biodata dengan parameter num
	bio(data, no)
}

func bio(arr []Biodata, num int) {

	//menampilkan data sesuai nomor
	if num-1 < len(arr) {

		fmt.Println("No :", arr[num-1].no)
		fmt.Println("Nama:", arr[num-1].name)
		fmt.Println("Alamat:", arr[num-1].address)
		fmt.Println("Pekerjaan:", arr[num-1].job)
		fmt.Println("Alasan:", arr[num-1].motivation)
		// perhitungan index dimulai dari 0 sehingga untuk menyesuaikan urutan dilakukan -1
		fmt.Println("=============================================")
	} else {
		fmt.Println("Sorry but the number you type too much.")
		fmt.Println("maximum number is : ", len(arr))
	}
}
