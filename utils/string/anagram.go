package utils

import (
	"sort"
	"strings"
)

func Anagram(listString ...string) [][]string {
	//store data anagram
	dataAnagram := map[string][]string{}

	//sorting alphabet
	var sortTemp string
	for _, str := range listString {
		sortTemp = str
		//convert  word to character
		sortingWord := strings.Split(sortTemp, "")

		//sorting character
		sort.Strings(sortingWord)

		//joining the character
		sortTemp = strings.Join(sortingWord, "")

		//append data by sort temp word
		dataAnagram[sortTemp] = append(dataAnagram[sortTemp], str)
	}

	//initialize result anagram
	var anagram [][]string

	//looping data from data anagram
	for e := range dataAnagram {
		//add data string from data anagram by sort temp word
		anagram = append(anagram, dataAnagram[e])
	}

	return anagram
}