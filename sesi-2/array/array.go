package main

import "fmt"

func main() {
	fmt.Print("ARRAY")
	fmt.Print("\n------------------------\n")

	//declare array with value
	var arr = [2]int{1, 2}
	var arrs = [2]string{"abc", "bac"}

	fmt.Printf("%#v ", arr)
	fmt.Printf("\n%#v ", arrs)

	//declare array without value
	var arr2 [2]float32
	arr2 = [2]float32{1.2, 2.3}
	fmt.Printf("\n\n%#v \n", arr2)

	//declare array without jumlahdata
	var numbers = [...]int{1, 3, 5, 6, 7, 8, 9, 12, 13}
	fmt.Println("\ndata array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//mengubah value array melalui index
	fmt.Print("\n------------------------\n")
	var lang = [3]string{"GO", "PHP", "PASCAL"}
	lang[2] = "C#"

	fmt.Print("the language: ", lang, "\n\n")

	//menampilkan nilai array dengan looping
	var teams = [...]string{"Arsenal", "MCITY", "Brightons", "Fullham"}
	for i := 0; i < len(teams); i++ {
		fmt.Println("team : ", teams[i])
	}

	//array multidimensional
	fmt.Print("\n------------------------\n")
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//menampilkan nilai array multidimensional dengan looping
	fmt.Print("\n------------------------\n")
	for i := 0; i < len(numbers1); i++ {
		for j := 0; j <= len(numbers1); j++ {
			fmt.Print(numbers1[i][j])
		}
	}
}
