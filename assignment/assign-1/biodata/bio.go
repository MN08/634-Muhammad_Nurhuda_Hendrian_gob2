package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("error : index student needed")
		return
	}
	//mengambil 1 data pada array argsRaw
	args := os.Args[1]
	argsi, err := strconv.Atoi(args)
	_ = err
	searchBiodataNumber, err := getBioByIndex(argsi)
	if err != nil {
		// show the error if from getStudentByIndex return error
		fmt.Println("error : ", err.Error())
		return
	} else {
		fmt.Println("No :", searchBiodataNumber.no)
		fmt.Println("Nama:", searchBiodataNumber.name)
		fmt.Println("Alamat:", searchBiodataNumber.address)
		fmt.Println("Pekerjaan:", searchBiodataNumber.job)
		fmt.Println("Alasan:", searchBiodataNumber.motivation)
		// perhitungan index dimulai dari 0 sehingga untuk menyesuaikan urutan dilakukan -1
		fmt.Println("=============================================")
	}

}

// create struct Biodata
type Biodata struct {
	no         string
	name       string
	address    string
	job        string
	motivation string
}

// construct data and use slice with pointer data to easily get the null / exist data
func getBioData() []*Biodata {
	return []*Biodata{
		{
			no:         "1",
			name:       "Albert",
			address:    "Antapani",
			job:        "Manager",
			motivation: "Untuk bisa Ngoding",
		},
		{
			no:         "2",
			name:       "Alberto",
			address:    "Bandung",
			job:        "Designer",
			motivation: "Untuk bisa Ngoding pake bahasa go"},
		{
			no:         "3",
			name:       "Albertus",
			address:    "Batu",
			job:        "Developer",
			motivation: "untuk bisa Ngoding bahasa go dengan baik",
		},
		{
			no:         "4",
			name:       "Albertino",
			address:    "Bekasi",
			job:        "Tester",
			motivation: "untuk bisa Ngoding pake bahasa go dengan baik dan benar",
		},
	}
}

func getBioByIndex(num int) (*Biodata, error) {

	bioData := getBioData()
	// fmt.Println("data type : ", reflect.ValueOf(bioData).Kind())
	totalData := len(bioData)
	if num < totalData {
		searchBiodataNumber := bioData[num-1]

		if searchBiodataNumber == nil {
			return nil, errors.New("Data Not Found")
		} else {
			return searchBiodataNumber, nil
		}
	} else {
		return nil, errors.New("Sorry but the number you type too much.")

	}

}
