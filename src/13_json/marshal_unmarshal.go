package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id       string `json:"id"`
	Pw       string `json:"pw"`
	Nickname string `json:"nickname"`
}

func main() {
	/*
		* marshal process
		Marshal function can access only public field (the first  spell is capital)
		Marshal function can convert struct value and slice
	*/
	fmt.Println("=====[Marshal]=====")
	users := []user{
		{
			Id:       "testId",
			Pw:       "testPw",
			Nickname: "Aivyss",
		},
		{
			Id:       "testId2",
			Pw:       "testPw2",
			Nickname: "Aivyss2",
		},
	}
	marshal, err := json.Marshal(users)

	if err != nil {
		fmt.Println("error occured:", err)
	}

	jsonStr := string(marshal)
	fmt.Println(jsonStr)

	/*
		* unmarshal process
		Unmarshal function needs two arguments (byte slice and pointer)
		Unmarshal function can convert byte slice(json string) to struct or slice
	*/
	fmt.Println("=====[Unmarshal]=====")
	var slice []user
	var p user
	json.Unmarshal([]byte(jsonStr), &slice)
	json.Unmarshal([]byte(`{"id":"testId","pw":"testPw","nickname":"Aivyss"}`), &p)

	fmt.Println(slice)
	fmt.Println(p)
}
