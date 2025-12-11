package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    // pointer1 := 0
    // pointer2 := len(s) - 1
    split_string := strings.Fields(s)
    fmt.Println(split_string)
    join_string := strings.Join(split_string, "")
    fmt.Println(join_string)
        cleaned := make([]rune, 0, len(s))
        for _, r := range s {
            if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') {
                if 'A' <= r && r <= 'Z' {
                    r = r + ('a' - 'A') // convert to lowercase
                }
                cleaned = append(cleaned, r)
            }
        }
        left, right := 0, len(cleaned)-1
        for left < right {
            if cleaned[left] != cleaned[right] {
                return false
            }
            left++
            right--
        }
        return true
}


func main () {
	fmt.Println(isPalindrome("race car"))
}