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
	inputStr := "59709511599794439805414014219880358445064269099345553494818286560304063399998657801629526113732466767578373307474609375929817361595469200826872565688108197109235040815426214109531925822745223338550232315662686923864318114370485155264844201080947518854684797571383091421294624331652208294087891792537136754322020911070917298783639755047408644387571604201164859259810557018398847239752708232169701196560341721916475238073458804201344527868552819678854931434638430059601039507016639454054034562680193879342212848230089775870308946301489595646123293699890239353150457214490749319019572887046296522891429720825181513685763060659768372996371503017206185697"
	//inputStr := "03036732577212944063491565474664" //12345678123456781234567812345678"
	multiplier := 10000

	var input []int
	for n := 0; n < multiplier; n++ {
		for i := range inputStr {
			s, err := strconv.Atoi(inputStr[i : i+1])
			f(err)
			input = append(input, s)
		}
	}
	//origLen := len(inputStr)
	offset, err := strconv.Atoi(inputStr[0:7])
	input = input[offset:]
	f(err)
	log.Printf("input str with len %v = %v...", len(input), input[:20])
	log.Printf("offset = %v", offset)

	//var sum int
	//for i := range inputStr {
	//	s, err := strconv.Atoi(inputStr[i : i+1])
	//	f(err)
	//	sum += s
	//}


	for i := 0; i < 100; i++ {
		//log.Printf("pat = %v", patForRound(i))
		//var roundResult []int
		roundResult := make([]int, 0)
		//pat := patForRound(i)

		var lastRowSum int
		var lastRowTerm int
		for row := 0; row < len(input); row++ {
			//rMult := 0
			//var rowNums []int
			//pat := patForRow(row+1, len(input))
			//rowNums := make([]int, offset, len(input))


			if row == 0 {
				for col := 0; col < len(input); col++ {
					newAdd :=  doMult(input, col, col+offset, row+offset)
					lastRowSum += newAdd // doMult(input, col, row)
					//rMult %= 10
					//rowNums = append(rowNums, newAdd)
				}
				lastRowTerm = doMult(input, 0, offset, offset)
				firstRowCp := lastRowSum % 10
				if firstRowCp < 0 {
					firstRowCp = -firstRowCp
				}
				roundResult = append(roundResult, firstRowCp)
			} else {
				newRow := lastRowSum - lastRowTerm
				lastRowTerm = doMult(input, row, row+offset, row+offset)
				lastRowSum = newRow
				newRow = newRow % 10
				if newRow < 0 {
					newRow = -newRow
				}
				roundResult = append(roundResult, newRow)
			}

			// ANSWER 48776785
			//
			//for col := row; col < len(input); col++ {
			//	newAdd :=  doMult(input, col, col+offset, row+offset)
			//	rMult += newAdd // doMult(input, col, row)
			//	//rMult %= 10
			//	//rowNums = append(rowNums, newAdd)
			//}
			//
			//firstRow :=  doMult(input, row, row+offset, row+offset)
			//firstRowCp := firstRow % 10
			//if firstRow < 0 {
			//	firstRowCp = -firstRowCp
			//}
			//roundResult = append(roundResult, rMult)

			//for col := row; col < (len(input)-row)%origLen; col++ {
			//	//newAdd :=  doMult(input, col, col+offset, row+offset)
			//	rMult += input[col] // doMult(input, col, row)
			//	//rMult %= 10
			//	//rowNums = append(rowNums, newAdd)
			//}
			//rMult += ((len(input)-row) / origLen) * sum


			//rMult %= 10
			//if rMult < 0 {
			//	rMult = -rMult
			//}
			////log.Printf("row op %v: %v becomes %v", row, rowNums, rMult)
			//roundResult = append(roundResult, rMult)
			//log.Printf("for round %v, done row %v, = %v", i, row, rMult)
		}
		log.Printf("roundResult len %v after %v rounds = %v...", len(roundResult), i+1, roundResult[:20])
		input = roundResult
	}
	log.Printf("final output = %v...", input[:20])
	//var outputStr
	//f(ioutil.WriteFile("cmd/day16/bigoutput.txt", []byte(input), 0644))
}

var (
	base = []int{0,1,0,-1}
)

func doMult(input []int, modifiedCol, col, row int) int {
	ch := input[modifiedCol]
	b := base[((col+1)/(row+1)) % 4]
	return ch * b
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