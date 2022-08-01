package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	listPecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	fmt.Print("Input Jumlah Uang: ")
	var uang int
	_ , err := fmt.Scanln(&uang)
	if err != nil {
		fmt.Println("err :" , err)
	}

	pecahan(uang, listPecahan)
}

func pecahan(uang int, listPecahan []int) {
	mapPecahan := make(map[string]int)

	for _, dataPecahan := range listPecahan {
		indexPecahan, totalPecahan, uangKepakeSamaPecahan := helperPecahan(dataPecahan, uang)
		mapPecahan[indexPecahan] = totalPecahan
		if totalPecahan != 0 {
			uang = uang - uangKepakeSamaPecahan
		}
	}

	fmt.Println(prettyPrint(mapPecahan))
	fmt.Println("sisa uang karena tidak ada pecahannya: " , uang)
}

func helperPecahan(inputPecahan int, inputUang int) (string, int, int) {
	if inputUang >= inputPecahan {
		moduloPecahan := inputUang % inputPecahan
		totalPecahan := (inputUang - moduloPecahan) / inputPecahan
		return "Rp." + strconv.Itoa(inputPecahan), totalPecahan, inputPecahan * totalPecahan
	}

	return "Rp." + strconv.Itoa(inputPecahan), 0, 0
}


func prettyPrint(data interface{}) string {
	JSON, _ := json.MarshalIndent(data, "", "\t")
	//s, _ := json.MarshalIndent(i, "", "\t")

	return string(JSON)
}