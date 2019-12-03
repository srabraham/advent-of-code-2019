package intcode

import (
	"log"
)

// RunIntCode implements https://adventofcode.com/2019/day/2
func RunIntCode(in []int64) []int64 {
	arr := make([]int64, len(in))
	copy(arr, in)
	instPtr := 0
	for {
		switch arr[instPtr] {
		case 1:
			// log.Printf("doing an add: %v, %v, %v, %v", arr[instPtr], arr[instPtr+1], arr[instPtr+2], arr[instPtr+3])
			aVal := arr[instPtr+1]
			bVal := arr[instPtr+2]
			cVal := arr[instPtr+3]
			arr[cVal] = arr[aVal] + arr[bVal]
			instPtr += 4
		case 2:
			// log.Printf("doing an mult: %v, %v, %v, %v", arr[instPtr], arr[instPtr+1], arr[instPtr+2], arr[instPtr+3])
			aVal := arr[instPtr+1]
			bVal := arr[instPtr+2]
			cVal := arr[instPtr+3]
			arr[cVal] = arr[aVal] * arr[bVal]
			instPtr += 4
		case 99:
			return arr
		default:
			log.Fatalf("bad opcode %v", arr[instPtr])
		}
	}
}
