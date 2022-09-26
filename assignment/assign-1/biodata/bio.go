package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//check if command line arguments available < 2 (only 1 argument)
	if len(os.Args) < 2 {
		fmt.Println("error : index student needed")
		return
	}
	//mengambil 1 argument
	args := os.Args[1]
	// karna os.Args hanya menyimpan array string dan input berupa int, maka kita convert no to int
	//strconv membutuhkan 2 parameter, variable untuk menyimpan hasil, serta error value, disini err diset menjadi nil
	argsi, errs := strconv.Atoi(args)
	_ = errs

	// send argsi to getBioByIndex function
	searchBiodataNumber, err := getBioByIndex(argsi)
	if err != nil {
		// show the error if from getBioByIndex return error
		fmt.Println("error : ", err.Error())
		return
	} else {
		// if not error show the data
		fmt.Println("No :", searchBiodataNumber.no)
		fmt.Println("Nama:", searchBiodataNumber.name)
		fmt.Println("Alamat:", searchBiodataNumber.address)
		fmt.Println("Pekerjaan:", searchBiodataNumber.job)
		fmt.Println("Alasan:", searchBiodataNumber.motivation)
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
// use of pointer because cannot return nil from a function that returns a string
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

// func getBioByIndex to find the bio
func getBioByIndex(num int) (*Biodata, error) {
	// get all data in GetBioData
	bioData := getBioData()

	totalData := len(bioData)

	if num < totalData {
		// find biodata we want to found
		searchBiodataNumber := bioData[num-1]

		// if data not found
		if searchBiodataNumber == nil {
			return nil, errors.New("Data Not Found")
		} else { //if data found
			return searchBiodataNumber, nil
		}
	} else {
		return nil, errors.New("Sorry but the number you type too much.")

	}

}
