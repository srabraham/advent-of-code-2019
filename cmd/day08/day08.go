package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type LayerVal struct {
	vals [][]int
}

func (lv LayerVal) String() string {
	var str string
	for i := 0; i < len(lv.vals); i++ {
		for j := 0; j < len(lv.vals[i]); j++ {
			switch lv.vals[i][j] {
			case 0:
				str += "  "
			case 1:
				str += "X "
			case 2:
				str += "  "
			}
		}
		str += "\n"
	}
	return str
}

func (lv LayerVal) countVal(v int) int {
	var count int
	for i := 0; i < len(lv.vals); i++ {
		for j := 0; j < len(lv.vals[i]); j++ {
			if lv.vals[i][j] == v {
				count++
			}
		}
	}
	return count
}

func MergeLayers(lv []LayerVal) LayerVal {
	outLayer := LayerVal{}
	for i := 0; i < len(lv[0].vals); i++ {
		var outRow []int
		for j := 0; j < len(lv[0].vals[i]); j++ {
			outVal := 2
		layerLoop:
			for l := 0; l < len(lv); l++ {
				switch lv[l].vals[i][j] {
				case 0:
					outVal = 0
					break layerLoop
				case 1:
					outVal = 1
					break layerLoop
				case 2:
					continue layerLoop
				}
			}
			outRow = append(outRow, outVal)
		}
		outLayer.vals = append(outLayer.vals, outRow)
	}
	return outLayer
}

func LayerFromString(str string, width, height int) LayerVal {
	lv := LayerVal{}
	for j := 0; j < height; j++ {
		heightCont := str[j*width : (j+1)*width]
		var widthCont []int
		for _, c := range heightCont {
			conv, err := strconv.Atoi(string(c))
			f(err)
			widthCont = append(widthCont, conv)
		}
		lv.vals = append(lv.vals, widthCont)
	}
	//log.Print(lv.vals)
	return lv
}

func Part1(filename string, width, height int) int {
	b, err := ioutil.ReadFile(filename)
	f(err)
	layerLength := width * height
	layerCount := len(b) / layerLength
	log.Printf("layerCount = %v", layerCount)

	layers := readLayers(layerCount, b, layerLength, width, height)

	minZero := -1
	minZeroIndex := -1
	for i := 0; i < layerCount; i++ {
		countZero := layers[i].countVal(0)
		if i == 0 || countZero < minZero {
			minZero = countZero
			minZeroIndex = i
		}
	}
	minMult := layers[minZeroIndex].countVal(1) * layers[minZeroIndex].countVal(2)
	log.Printf("minMult2 = %v", minMult)
	return minMult
}

func readLayers(layerCount int, b []byte, layerLength int, width int, height int) []LayerVal {
	var layers []LayerVal
	for i := 0; i < layerCount; i++ {
		contents := string(b)[i*layerLength : (i+1)*layerLength]
		layers = append(layers, LayerFromString(contents, width, height))
	}
	return layers
}

func Part2(filename string, width, height int) string {
	b, err := ioutil.ReadFile(filename)
	f(err)
	layerLength := width * height

	layerCount := len(b) / layerLength
	log.Printf("layerCount = %v", layerCount)

	layers := readLayers(layerCount, b, layerLength, width, height)

	result := MergeLayers(layers).String()
	log.Printf("New and improved:\n%v", result)
	return result
}

// 1330
func main() {
	Part1("cmd/day08/input08-1.txt", 25, 6)
	Part2("cmd/day08/input08-1.txt", 25, 6)
}
