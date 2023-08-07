package structures

type Person struct {
	First string
	Last  string
	Age   int
}

type Student struct {
	Person
	StudentId string
}
