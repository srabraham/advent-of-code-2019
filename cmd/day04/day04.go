package main

import (
	"github.com/srabraham/advent-of-code-2019/internal/seanmath"
	"log"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//b, err := ioutil.ReadFile("cmd/day04/input04-1.txt")
	//f(err)
	//lines := strings.Split(string(b), "\n")
	nums := make([]int64, 0)
	for i := int64(171309); i <= 643603; i++ {

		s := seanmath.ToString(i)
		doubleNum := (s[0] == s[1]) || (s[1] == s[2]) ||(s[2] == s[3]) ||(s[3] == s[4]) ||(s[4] == s[5])
		nonDec := (s[0] <= s[1]) && (s[1] <= s[2]) &&(s[2] <= s[3]) &&(s[3] <= s[4]) &&(s[4] <= s[5])

		pt3 := (s[0] == s[1] && s[1] != s[2]) || (s[0] != s[1] && s[1] == s[2] && s[2] != s[3]) ||(s[1] != s[2] && s[2] == s[3] && s[3] != s[4]) ||(s[2] != s[3] && s[3] == s[4] && s[4] != s[5]) ||(s[3] != s[4] && s[4] == s[5])

		if doubleNum && nonDec && pt3 {
			nums = append(nums, i)
		}

	}
	log.Printf("len = %v, nums = %v", len(nums), nums)
}