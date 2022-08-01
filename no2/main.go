package main

import "fmt"

func main() {
	string1 := scanInput("input string 1 : ")
	string2 := scanInput("input string 2 : ")

	hasilDiff , totalDiff := helperDiff(string1, string2)
	fmt.Println("hasil pembanding : ", hasilDiff)
	fmt.Println("total Diff  : ", totalDiff)
}

func helperDiff(stringInput string, stringPembanding string) (bool, int) {

	checkDuplicateAlphabet := make(map[string]bool)
	var list_char_not_modified []string

	for _, datas1 := range stringInput {
		dataStringS1 := string(datas1)
		for _, datas2 := range stringPembanding {
			dataStringS2 := string(datas2)
			if (dataStringS1 == dataStringS2) && !checkDuplicateAlphabet[dataStringS1] {
				checkDuplicateAlphabet[dataStringS1] = true
				list_char_not_modified = append(list_char_not_modified, dataStringS1)
				break
			}
		}
	}

	sumCharModified := len(stringInput) - len(list_char_not_modified)

	if sumCharModified > 1 {
		return false, sumCharModified
	} else if len(stringInput) < len(stringPembanding) {
		lenghtStringDiff := len(stringInput) - len(stringPembanding)
		if lenghtStringDiff < 0 {
			lenghtStringDiff = lenghtStringDiff * -1
		}
		sumCharModified = sumCharModified + lenghtStringDiff

		if sumCharModified > 1 {
			return false, sumCharModified
		}
	}

	return true, sumCharModified

}

func scanInput(message string) string {
	var data string

	fmt.Print(message)
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("err :", err)
	}

	return data
}
