package main

import "fmt"

func isPalindrome(s string) bool {
    // pointer1 := 0
    // pointer2 := len(s) - 1
    for i := range s {
        if s[i] == s[-i - 1] {
            if s[i] == s[-i - 1] {
                return true
            }
        } else {
            return false
        }
    }
}


func main () {
	fmt.Println("tabacat")
}