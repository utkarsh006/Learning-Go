package main

import (
	"fmt"
	"strings"
)

func main() {
   str := "Hello Guyzz"
   fmt.Println("Length:", len(str))

   fmt.Println("contains Guyzzz? ", strings.Contains(str, "Guyzz"))

   fmt.Println("Count of o is:", strings.Count(str, "o"))

   //making array of strings
   words := []string{"Hello", "Go", "World"}
   joined := strings.Join(words, "-")
   fmt.Println("Joined string:", joined)

   // user@gmail.com

   email := "user123@gmail.com"
   var index = strings.Index(email, "@")
   fmt.Println(email[:index])
   fmt.Println(email[index+1:])
}  