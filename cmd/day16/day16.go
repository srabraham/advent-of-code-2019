package main

import (
	"log"
	"strconv"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("starting...")
	//b, err := ioutil.ReadFile("cmd/day16/input16-0.txt")
	//b, err := ioutil.ReadFile("cmd/day16/input16-1.txt")
	// inputStr := "59709511599794439805414014219880358445064269099345553494818286560304063399998657801629526113732466767578373307474609375929817361595469200826872565688108197109235040815426214109531925822745223338550232315662686923864318114370485155264844201080947518854684797571383091421294624331652208294087891792537136754322020911070917298783639755047408644387571604201164859259810557018398847239752708232169701196560341721916475238073458804201344527868552819678854931434638430059601039507016639454054034562680193879342212848230089775870308946301489595646123293699890239353150457214490749319019572887046296522891429720825181513685763060659768372996371503017206185697"
	inputStr := "123412341234123412341234123412341234123412341234"
	var input []int
	for i := range inputStr {
		s, err := strconv.Atoi(inputStr[i:i+1])
		f(err)
		input = append(input, s)
	}
	log.Printf("input str = %v", input)
	for i := 0; i < 100; i++ {
		//log.Printf("pat = %v", patForRound(i))
		var roundResult []int
		//pat := patForRound(i)

		for row := 0; row < len(input); row++ {
			rMult := 0
			var rowNums []int
			pat := patForRow(row+1, len(input))
			for k := 0; k < len(input); k++ {
				newAdd := input[k] * pat[k]
				rMult += newAdd
				rowNums = append(rowNums, newAdd)
			}
			rMult %= 10
			if rMult < 0 {
				rMult = -rMult
			}
			log.Printf("row op %v: %v becomes %v", row, rowNums, rMult)
			roundResult = append(roundResult, rMult)
		}
		log.Printf("roundResult after %v rounds = %v", i+1, roundResult)
		input = roundResult
	}
	log.Printf("final output = %v", input)
}

func doMult(input string, row int) {

}

func patForRow(r int, maxLen int) []int {
	base := []int{0,1,0,-1}
	var result []int
	outer:
		for true {
			for _, b := range base {
				for i := 0; i < r; i++ {
					result = append(result, b)
					if len(result) > maxLen + 10 {
						break outer
					}
				}
			}
		}

	var result2 []int
	for i, r := range result {
		if i != 0 {
			result2 = append(result2, r)
		}
	}
	return result2
}