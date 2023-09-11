package main

import (
	"fmt"
	"structs/zoo"
)

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

func main() {
	fmt.Printf("On duty today: %s, with %d years of experience\n", zoo.Mihalich.Name, zoo.Mihalich.Experience)
	fmt.Printf("In the program there is a %s named: %s. It has %d lags", zoo.Spider.Kind, zoo.Spider.Name, zoo.Spider.Number_of_legs)

}
