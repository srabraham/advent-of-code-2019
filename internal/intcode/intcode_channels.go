package intcode

import "log"

func RunIntCodeWithChannels(ops []int64, inputCh chan int64, outputCh chan int64) int64 {
	arr := make([]int64, len(ops))
	copy(arr, ops)
	var output int64
	instPtr := 0
	var lastOutput int64
bigLoop:
	for {
		inst := arr[instPtr] % 100
		modeDec := arr[instPtr] / 100
		mode1 := modeDec % 10
		mode2 := (modeDec / 10) % 10
		mode3 := (modeDec / 100) % 10

		//log.Printf("state: %v", arr)
		//log.Printf("running instr %v, %v", arr[instPtr], inst)

		switch inst {
		case 1:
			// log.Printf("doing an add: %v, %v, %v, %v", arr[instPtr], arr[instPtr+1], arr[instPtr+2], arr[instPtr+3])
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			cVal := arr[instPtr+3]
			if mode3 != 0 {
				log.Fatal("unexpected mode 3")
			}
			arr[cVal] = aVal + bVal
			instPtr += 4
		case 2:
			// log.Printf("doing an mult: %v, %v, %v, %v", arr[instPtr], arr[instPtr+1], arr[instPtr+2], arr[instPtr+3])
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			cVal := arr[instPtr+3]
			if mode3 != 0 {
				log.Fatal("unexpected mode 3")
			}
			arr[cVal] = aVal * bVal
			instPtr += 4
		case 3: // input
			arr[arr[instPtr+1]] = <-inputCh
			instPtr += 2
		case 4: // output
			aVal := arr[instPtr+1]
			if mode1 == 0 {
				aVal = arr[aVal]
			}
			log.Printf("OUTPUT %v", aVal)
			output = aVal
			outputCh <- output
			lastOutput = output
			instPtr += 2
		case 5: // jump if true
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			if aVal != 0 {
				instPtr = int(bVal)
			} else {
				instPtr += 3
			}
		case 6: // jump if false
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			if aVal == 0 {
				instPtr = int(bVal)
			} else {
				instPtr += 3
			}
		case 7: // less than
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			if aVal < bVal {
				arr[arr[instPtr+3]] = 1
			} else {
				arr[arr[instPtr+3]] = 0
			}
			instPtr += 4
		case 8: // equals
			aVal := getModeVal(arr, mode1, instPtr+1)
			bVal := getModeVal(arr, mode2, instPtr+2)
			if aVal == bVal {
				arr[arr[instPtr+3]] = 1
			} else {
				arr[arr[instPtr+3]] = 0
			}
			instPtr += 4
		case 99:
			log.Print("got 99. breaking")
			break bigLoop
		default:
			log.Fatalf("bad opcode %v", inst)
		}
	}
	// nothing else to send
	close(outputCh)
	return lastOutput
}
