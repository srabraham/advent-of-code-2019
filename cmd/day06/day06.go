package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// b, err := ioutil.ReadFile("cmd/day05/input06-0.txt")
	b, err := ioutil.ReadFile("cmd/day06/input06-1.txt")
	f(err)
	rows := strings.Split(string(b), "\n")

	m := make(map[string]string)

	for _, r := range rows {
		split := strings.Split(r, ")")
		if len(split) != 2 {
			log.Fatal("bad")
		}
		if m[split[1]] != "" {
			log.Fatal("bad 2")
		}
		m[split[1]] = split[0]
	}
	var count int
	for k := range m {
		count += getOrbitsCount(k, m)
	}
	log.Printf("count (part 1) = %v", count)

	santaOrbits := getAllOrbitsStr("SAN", m)
	youOrbits := getAllOrbitsStr("YOU", m)

	transfers := make([]string, 0)
	transfers = append(transfers, sliceDiff(santaOrbits, youOrbits)...)
	transfers = append(transfers, sliceDiff(youOrbits, santaOrbits)...)

	// subtract off the entries for SAN and YOU
	log.Printf("found %v transfers (part 2)", len(transfers) - 2)
}

// sliceDiff returns elements of a that are not in b
func sliceDiff(a, b []string) []string {
	bMap := make(map[string]bool)
	for _, e := range b {
		bMap[e] = true
	}
	result := make([]string, 0)
	for _, e := range a {
		if !bMap[e] {
			result = append(result, e)
		}
	}
	return result
}

func getOrbitsCount(start string, m map[string]string) int {
	orb := m[start]
	if orb != "COM" {
		return 1 + getOrbitsCount(orb, m)
	}
	return 1
}

func getAllOrbitsStr(start string, m map[string]string) []string {
	orb := m[start]
	if orb != "COM" {
		s := make([]string, 0)
		s = append(s, getAllOrbitsStr(orb, m)...)
		s = append(s, start)
		return s
	}
	return []string{"COM"}
}