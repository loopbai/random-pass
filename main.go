package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Shuffle(slice []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(slice))
	n := len(slice)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		ret[i] = slice[randIndex]
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}

func main() {
	number := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	capitalLetter := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}
	lowercaseLetter := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "m", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z"}
	symbol := []string{"!", "@", "#", "$", "%", "^", "&", "*"}

	combine := []string{}
	combine = append(combine, number...)
	combine = append(combine, capitalLetter...)
	combine = append(combine, lowercaseLetter...)
	combine = append(combine, symbol...)
	fmt.Println(combine)

	combine = Shuffle(combine)
	fmt.Println(combine)

	result := strings.Join(combine, "")
	fmt.Println(result)

}
