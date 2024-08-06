package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func cleanUp(word string) string {
	var clean string
	for _, char := range word {
		if !unicode.IsPunct(char) {
			clean += strings.ToLower(string(char))
		}
	}
	clean = strings.TrimSpace(clean)
	return clean
}

func wordCount(input string) map[string]int {
	separated := strings.Split(input, " ")
	frequency := make(map[string]int)

	for _, word := range separated {
		word = cleanUp(word)
		if word != "" {
			frequency[word]++
		}
	}
	return frequency

}

func concatenate(input []string) string {
	conc := ""
	for _, val := range input {
		conc += cleanUp(val)
	}
	return conc

}

func checkPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func takeInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v:", prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read Op Failed")
		return ""
	}
	return input
}

func main() {
	input := takeInput("Insert Sentence for count")
	myCount := wordCount(input)
	fmt.Println(myCount)

	input = takeInput("Insert Sentence for palindrome")
	arrayForm := strings.Split(input, " ")
	concat := concatenate(arrayForm)
	answer := checkPalindrome(concat)
	if answer {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

	// fmt.(cleanUp("THIS? is:"))
}
