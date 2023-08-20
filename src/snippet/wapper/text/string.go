package text

import (
	"errors"
	"strconv"
	"strings"
)

type String string

func (s String) Int() (i int, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			i = 0
			err = errors.New(rec.(string))
		}
	}()
	intVar, err := strconv.Atoi(string(s))

	if err != nil {
		panic(err)
	}

	return intVar, nil
}

func (s String) Split(sep string) []String {
	result := []String{}

	tmp := strings.Split(string(s), sep)
	for _, v := range tmp {
		result = append(result, String(v))
	}

	return result
}

func (s String) Fields() []String {
	result := []String{}

	tmp := strings.Fields(string(s))
	for _, v := range tmp {
		result = append(result, String(v))
	}

	return result
}
