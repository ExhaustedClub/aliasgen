package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"exhausted.club/aliasgen/alias"
	"exhausted.club/aliasgen/resources"
)

func main() {
	var width int = 5
	var height = 20
	tw := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for i := 0; i < height; i++ {
		fmt.Fprintf(tw, namestw(width), names(width)...)
	}
	tw.Flush()
}

func namestw(i int) string {
	s := strings.Repeat("%s\t", i)
	s = s[:len(s)-1] + "\n"
	return s
}

func names(i int) (s []any) {
	s = make([]any, i)
	for si := 0; si < i; si++ {
		s[si] = firstLast()
	}
	return s
}

var fm = map[string]int{}
var lm = map[string]int{}

func firstLast() string {
	f := alias.RandomFromList(resources.GetFirstNameList(), nil)
	l := alias.RandomFromList(resources.GetLastNameList(), nil)

	if fm[f] > 2 {
		panic("First name " + f + " already drew")
	}

	if lm[l] > 2 {
		panic("First name " + l + " already drew")
	}

	fm[f]++
	lm[l]++

	return f + " " + l
}
