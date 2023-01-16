package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	mapStrings := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		val := strs[i]
		valS := sortString(val)
		mapStrings[valS] = append(mapStrings[valS], val)
	}

	for _, v := range mapStrings {
		result = append(result, v)
	}

	return result
}

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expect := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	fmt.Println(expect, groupAnagrams(str))
}
