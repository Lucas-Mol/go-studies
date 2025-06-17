package main

import (
	"fmt"
	"sort"
)

type printer struct {
	pages     int32
	threshold int32
	used      bool
}

func getPages(pages []int32, threshold []int32) int64 {
	n := len(pages)
	printers := make([]printer, n)
	for i := 0; i < n; i++ {
		printers[i] = printer{pages[i], threshold[i], false}
	}

	sort.Slice(printers, func(i, j int) bool {
		if printers[i].pages == printers[j].pages {
			return printers[i].threshold > printers[j].threshold
		}
		return printers[i].pages > printers[j].pages
	})

	var total int64 = 0
	var active []printer

	for _, p := range printers {
		if p.threshold < int32(len(active)) {
			break
		}
		active = append(active, p)
		total += int64(p.pages)
	}

	return total
}

func main() {
	pages := []int32{4, 1, 5, 2, 3}
	threshold := []int32{3, 3, 2, 3, 3}
	fmt.Println("Expected: 14 - Returned:", getPages(pages, threshold))

	pages = []int32{1, 2, 100, 1, 1}
	threshold = []int32{10, 10, 10, 10, 10}
	fmt.Println("Expected: 105 - Returned:", getPages(pages, threshold))

	pages = []int32{3, 3, 3}
	threshold = []int32{5, 5, 5}
	fmt.Println("Expected: 9 - Returned:", getPages(pages, threshold))
}
