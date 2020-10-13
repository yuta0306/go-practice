package main

import "fmt"

func main()  {
	i := 3
	i32 := '3'
	f := 3.14
	j := 1 + 1i
	b := false
	s := "string"
	u := []byte{'3'}
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("i32 is of type %T\n", i32)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("j is of type %T\n", j)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("s is of type %T\n", s)
	fmt.Printf("u is of type %T\n", u)
}