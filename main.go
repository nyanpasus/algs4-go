package main

import (
	"algs4/chapter5/alphabet"
	"algs4/chapter5/lsd"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := []string{"4PGC938", "2IYE230", "3CI0720", "1ICK750", "1OHV845", "4JZY524", "1ICK750", "3CIO720", "1OHV845", "1OHV845", "2RLA629", "2RLA629", "3ATW723"}
	lsd.Sort(a, 7)
	fmt.Println(a)
}

func Count() {
	alpha, err := alphabet.NewByAlpha(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(alpha)
	r := alpha.R()
	count := make([]int, r)

	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	s := sb.String()
	n := len(s)
	fmt.Println(s)
	// for _, v := range s {
	// 	if alpha.Contains(v) {
	// 		count[alpha.ToIndex(v)]++
	// 	}
	// }
	a := alpha.ToIndices(s)
	for i := 0; i < n; i++ {
		if a[i] != -1 {
			count[a[i]]++
		}
	}

	for c := 0; c < r; c++ {
		fmt.Println(strconv.QuoteRune(alpha.ToRune(c)), count[c])
	}
}
