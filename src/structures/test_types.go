package structures

type Person struct {
	First string
	Last  string
	Age   int
}

type Student struct {
	Person    // no field name => embedded structure
	StudentId string
}

type Worker struct {
	Person         Person
	JobDescription string
	Salary         int
	Id             string
}
