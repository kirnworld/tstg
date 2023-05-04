package main

import (
	"datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	var names []string
	lines, err := datafile.GetStrings("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name_v := range counts {
		names = append(names, name_v)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}
