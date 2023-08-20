package numeric

import "fmt"

type Int int

func (i Int) String() string {
	return fmt.Sprint(i)
}

func (i Int) Bool() bool {
	result := false

	if i == 1 {
		result = true
	}

	return result
}
