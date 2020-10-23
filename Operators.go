package main

import "fmt"

func main() {

   var a int = 21
   var b int = 10
   var i,j,k,l,m int

   i = a + b
   fmt.Printf("Line 1 - Value of c is %d\n", i )
   
   j = a - b
   fmt.Printf("Line 2 - Value of c is %d\n", j )
   
   k = a * b
   fmt.Printf("Line 3 - Value of c is %d\n", k )
   
   l = a / b
   fmt.Printf("Line 4 - Value of c is %d\n", l )
   
   m = a % b
   fmt.Printf("Line 5 - Value of c is %d\n", m )
   
   a++
   fmt.Printf("Line 6 - Value of a is %d\n", a )
   
   a--
   fmt.Printf("Line 7 - Value of a is %d\n", a )
}