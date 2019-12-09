package intcode

import (
	"log"
)

func getModeTarget(vals map[int64]int64, mode int64, pos int64, relativeBase int64) int64 {
	var targ int64
	aVal := vals[pos]
	if mode == 0 {
		targ = aVal
	}
	if mode == 2 {
		targ = aVal + relativeBase
	}
	return targ
}

func getModeVal(vals map[int64]int64, mode int64, pos int64, relativeBase int64) int64 {
	aVal := vals[pos]
	if mode == 0 {
		aVal = vals[aVal]
	}
	if mode == 2 {
		aVal = vals[aVal+relativeBase]
	}
	return aVal
}

func RunIntCodeWithChannels(ops []int64, inputCh chan int64, outputCh chan int64) int64 {
	var commandRunCount int64
	vals := make(map[int64]int64)
	for i, op := range ops {
		vals[int64(i)] = op
	}
	var output int64
	var instPtr int64
	var lastOutput int64
	var relativeBase int64
bigLoop:
	for {
		commandRunCount++
		inst := vals[instPtr] % 100
		modeDec := vals[instPtr] / 100
		mode1 := modeDec % 10
		mode2 := (modeDec / 10) % 10
		mode3 := (modeDec / 100) % 10

		//log.Printf("state: %v", arr)
		//log.Printf("running instr %v, %v", arr[instPtr], inst)

		switch inst {
		case 1:
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(vals, mode3, instPtr+3, relativeBase)
			vals[cTarg] = aVal + bVal
			instPtr += 4
		case 2:
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(vals, mode3, instPtr+3, relativeBase)
			vals[cTarg] = aVal * bVal
			instPtr += 4
		case 3: // input
			aTarg := getModeTarget(vals, mode1, instPtr+1, relativeBase)
			vals[aTarg] = <-inputCh
			instPtr += 2
		case 4: // output
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			log.Printf("OUTPUT %v", aVal)
			output = aVal
			outputCh <- output
			lastOutput = output
			instPtr += 2
		case 5: // jump if true
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			if aVal != 0 {
				instPtr = bVal
			} else {
				instPtr += 3
			}
		case 6: // jump if false
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			if aVal == 0 {
				instPtr = bVal
			} else {
				instPtr += 3
			}
		case 7: // less than
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(vals, mode3, instPtr+3, relativeBase)
			if aVal < bVal {
				vals[cTarg] = 1
			} else {
				vals[cTarg] = 0
			}
			instPtr += 4
		case 8: // equals
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
			bVal := getModeVal(vals, mode2, instPtr+2, relativeBase)
			cTarg := getModeTarget(vals, mode3, instPtr+3, relativeBase)
			if aVal == bVal {
				vals[cTarg] = 1
			} else {
				vals[cTarg] = 0
			}
			instPtr += 4
		case 9:
			aVal := getModeVal(vals, mode1, instPtr+1, relativeBase)
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
