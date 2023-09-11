package zoo

// Private struct, visible inside this package only
type zookeeper struct {
	// Contains different field names with different types
	Name       string
	Surname    string
	Age        int
	Experience int
}

type animal struct {
	Kind           string
	Name           string
	Number_of_legs int
}

type cage struct {
	Size   string
	Color  string
	Number int
}

var Mihalich = zookeeper{"Mihail", "Ivanov", 56, 30}
var Spider = animal{Kind: "spider", Name: "Gora", Number_of_legs: 8}
