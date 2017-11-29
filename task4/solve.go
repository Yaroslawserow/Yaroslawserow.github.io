package main

import "unicode"

func RemoveEven(array_input []int) []int {
    answer := make([]int, 0)
    for _, answer := range array_input {
        if curent % 2 == 1 {
            answer = append(answer, curent)
        }
    }
    return answer
}

func PowerGenerator(input int) func() int {
    answer := 1
    return func() int {
        answer *= input
        return answer
    }
}

func DifferentWordsCount(input string) int {
    word := ""
    set := make(map[string]bool)
    answer := 0
    for _, i := range (input + " ") {
        if unicode.IsLetter(i) {
            word += string(unicode.ToLower(i))
        } else if word != "" {
            if !set[word] {
                answer += 1
            }
            set[word] = true
            word = ""
        }
    }
    return answer
}
