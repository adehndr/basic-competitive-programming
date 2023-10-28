package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for t := 0; t < T; t++ {
		var M int
		var N int
		fmt.Scanf("%d %d\n", &N, &M)
		var selected string
		fmt.Scanf("%s\n", &selected)
		allResult := make(map[string][]int)
		selectedResult := map[string][]int{selected: {}}
		for n := 0; n < N; n++ {
			var tempInput string
			var result1 int
			var result2 int
			var result3 int
			fmt.Scanf("%s %d %d %d\n", &tempInput, &result1, &result2, &result3)
			tempArr := []int{result1,result2,result3}
			if tempInput == selected {
				selectedResult[selected] = tempArr
			}else {
				allResult[tempInput] = tempArr
			}
		}
		countWinner := 0
		for _, v := range allResult {
			if isOtherWin(selectedResult[selected],v) {
				countWinner += 1
			}
		}

		if countWinner < M {
			fmt.Println("YA")
		}else {
			fmt.Println("TIDAK")
		}
	}

}

func isOtherWin(mC, other []int) bool {
	if other[2] != mC[2] {return other[2] > mC[2]}
	if other[1] != mC[1] {return other[1] > mC[1]}
	return other[0] > mC[0]
}