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
    for i := range join_string {
        if join_string[i] == join_string[len(join_string)- 1 - i] {
            continue
        } else {
            return false
        }
    }
    return true
}


func main () {
	fmt.Println(isPalindrome("race car"))
}