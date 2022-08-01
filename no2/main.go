package main

import (
	"fmt"
)

func main() {

	string1 := "telkom"
	string2 := "telecom"

	helperDiff(string1, string2)
	//fmt.Println(helperDiff2(strings.Split(string1, ""), strings.Split(string2, "")))
}

func helperDiff(stringInput string, stringPembanding string) {

	checkDuplicateAlphabet := make(map[string]bool)

	countDiff := 0
	var diff []string

	for _, datas1 := range stringInput {
		dataStringS1 := string(datas1)
		fmt.Println("datas1 :", dataStringS1)

		for _, datas2 := range stringPembanding {
			dataStringS2 := string(datas2)

			fmt.Println("datas2 :", dataStringS2)

			if (dataStringS1 == dataStringS2) && !checkDuplicateAlphabet[dataStringS1] {
				checkDuplicateAlphabet[dataStringS1] = true
			} else {
				if (dataStringS1 == dataStringS2) && checkDuplicateAlphabet[dataStringS2] {
					diff = append(diff, dataStringS2)
				}

				//if !checkDuplicateAlphabet[dataStringS2]{
				//	fmt.Println("checkkk : " , dataStringS2)
				//}

				//if (dataStringS1 != dataStringS2) && !checkDuplicateAlphabet[dataStringS2]{
				//	fmt.Println("checkkk : " , dataStringS2)
				//}
				//if (dataStringS1 != dataStringS2) && !checkDuplicateAlphabet[dataStringS2] {
				//	fmt.Println("datas1diff :", dataStringS1)
				//	fmt.Println("datas2diff :", dataStringS2)
				//
				//	diff = append(diff, dataStringS2)
				//}
			}
		}

	}
	fmt.Println("check duplicate : " , checkDuplicateAlphabet)

	fmt.Println(countDiff)
	fmt.Println(diff)

}

//func helperDiff2(slice1 []string, slice2 []string) bool {
//	var diff []string
//
//	for i := 0; i < 2; i++ {
//		for _, s1 := range slice1 {
//			found := false
//			for _, s2 := range slice2 {
//				if s1 == s2 {
//					found = true
//					break
//				}
//			}
//			if !found {
//				diff = append(diff, s1)
//			}
//		}
//		if i == 0 {
//			slice1, slice2 = slice2, slice1
//		}
//	}
//
//	fmt.Println(diff)
//	if len(diff) > 1 {
//		return false
//	}
//	return true
//}
