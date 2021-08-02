package main

import (
	"fmt"
	"sort"
)

var (
	wordsDict = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ss [][]string
	for e := range m {
		ss = append(ss, m[e])
	}
	return ss
}

func main() {
	anagrams := groupAnagrams(wordsDict)
	for _, words := range anagrams {
		fmt.Println(words)
	}
}
