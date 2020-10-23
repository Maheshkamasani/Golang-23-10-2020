package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}
func main() {
   a, b := swap("Madhan", "Mahesh")
   fmt.Println(a, b)
}