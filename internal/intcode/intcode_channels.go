package intcode

import (
	"log"
	"math"
)

func getModeTarget(arr []int64, mode int64, pos int, relativeBase int64) int64 {
	var targ int64
	aVal := arr[pos]
	if mode == 0 {
		targ = aVal
	}
	if mode == 2 {
		targ = aVal + relativeBase
	}
	return targ
}

func RunIntCodeWithChannels(ops []int64, inputCh chan int64, outputCh chan int64) int64 {
	var commandRunCount int64
	arr := make([]int64, math.MaxInt32)
	copy(arr, ops)
	var output int64
	instPtr := 0
	var lastOutput int64
	var relativeBase int64
bigLoop:
	for {
		commandRunCount++
		inst := arr[instPtr] % 100
		modeDec := arr[instPtr] / 100
		mode1 := modeDec % 10
		mode2 := (modeDec / 10) % 10
		mode3 := (modeDec / 100) % 10

		//log.Printf("state: %v", arr)
		//log.Printf("running instr %v, %v", arr[instPtr], inst)

		switch inst {
		case 1:
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(arr, mode3, instPtr+3, relativeBase)
			arr[cTarg] = aVal + bVal
			instPtr += 4
		case 2:
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(arr, mode3, instPtr+3, relativeBase)
			arr[cTarg] = aVal * bVal
			instPtr += 4
		case 3: // input
			aTarg := getModeTarget(arr, mode1, instPtr+1, relativeBase)
			arr[aTarg] = <-inputCh
			instPtr += 2
		case 4: // output
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			log.Printf("OUTPUT %v", aVal)
			output = aVal
			outputCh <- output
			lastOutput = output
			instPtr += 2
		case 5: // jump if true
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			if aVal != 0 {
				instPtr = int(bVal)
			} else {
				instPtr += 3
			}
		case 6: // jump if false
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			if aVal == 0 {
				instPtr = int(bVal)
			} else {
				instPtr += 3
			}
		case 7: // less than
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(arr, mode3, instPtr+3, relativeBase)
			if aVal < bVal {
				arr[cTarg] = 1
			} else {
				arr[cTarg] = 0
			}
			instPtr += 4
		case 8: // equals
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(arr, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(arr, mode3, instPtr+3, relativeBase)
			if aVal == bVal {
				arr[cTarg] = 1
			} else {
				arr[cTarg] = 0
			}
			instPtr += 4
		case 9:
			aVal := getModeVal(arr, mode1, instPtr+1, relativeBase)
			relativeBase += aVal
			instPtr += 2
		case 99:
			log.Print("got 99. breaking")
			break bigLoop
		default:
			log.Fatalf("bad opcode %v", inst)
		}
	}
	// nothing else to send
	close(outputCh)
	log.Printf("ran %v commands", commandRunCount)
	return lastOutput
}
