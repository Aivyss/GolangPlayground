package main

import (
	"com.playground/structures"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}

func test1() {
	fmt.Println("====[test1 target: reassign]=====")

	changePriceValue1 := func(i structures.Item) {
		i.Price = 3000
	}
	changePriceValue2 := func(i *structures.Item) {
		i.Price = 3000
	}
	item := structures.NewItem1("itemA", 1000, true)
	copyItem := item

	fmt.Println("====[Price value is changed by directly.]=====")
	copyItem.Price = 2000
	fmt.Println("item =", item.Price, "[not changed]")
	fmt.Println("copyItem =", copyItem.Price, "[changed]")

	fmt.Println("====[Price value is changed by the function1 (value function).]=====")
	changePriceValue1(copyItem)
	fmt.Println("item =", item.Price, "[not changed]")
	fmt.Println("copyItem =", copyItem.Price, "[not changed]")

	fmt.Println("====[Price value is changed by the function2 (pointer function).]=====")
	changePriceValue2(&copyItem)
	fmt.Println("item =", item.Price, "[not changed]")
	fmt.Println("copyItem =", copyItem.Price, "[changed]")
}

func test2() {
	fmt.Println("====[test2 target: reassign by pointer]=====")

	changePriceValue1 := func(i structures.Item) {
		i.Price = 3000
	}
	changePriceValue2 := func(i *structures.Item) {
		i.Price = 3000
	}
	itemP := structures.NewItem2("itemA", 1000, true)
	item := *itemP

	fmt.Println("====[Price value is changed by directly.]=====")
	item.Price = 2000
	fmt.Println("itemP =", itemP.Price, "[not changed]")
	fmt.Println("item =", item.Price, "[changed]")

	fmt.Println("====[Price value is changed by the function1 (value function).]=====")
	changePriceValue1(item)
	fmt.Println("itemP =", itemP.Price, "[not changed]")
	fmt.Println("item =", item.Price, "[not changed]")

	fmt.Println("====[Price value is changed by the function2 (pointer function).]=====")
	changePriceValue2(&item)
	fmt.Println("*itemP =", itemP.Price, "[not changed]")
	fmt.Println("item =", item.Price, "[changed]")
}

func test3() {
	fmt.Println("====[test3 target: pointer]=====")

	changePriceValue1 := func(i structures.Item) {
		i.Price = 3000
	}
	changePriceValue2 := func(i *structures.Item) {
		i.Price = 3000
	}
	itemP := structures.NewItem2("itemA", 1000, true)
	item := *itemP

	fmt.Println("====[Price value is changed by directly.]=====")
	itemP.Price = 2000
	fmt.Println("itemP =", itemP.Price, "[changed]")
	fmt.Println("item =", item.Price, "[not changed]")

	fmt.Println("====[Price value is changed by the function1 (value function).]=====")
	changePriceValue1(*itemP)
	fmt.Println("itemP =", itemP.Price, "[not changed]")
	fmt.Println("item =", item.Price, "[not changed]")

	fmt.Println("====[Price value is changed by the function2 (pointer function).]=====")
	changePriceValue2(itemP)
	fmt.Println("*itemP =", itemP.Price, "[changed]")
	fmt.Println("item =", item.Price, "[not changed]")
}

func test4() {
	fmt.Println("====[test4]=====")
	add := func(slice []int) {
		slice[0] = -1
		fmt.Println("add slice =", slice)
	}

	x := []int{1, 2, 3, 4, 5}
	add(x)
	fmt.Println("main x =", x)
}

func test5() {
	fmt.Println("====[test5 target: sub field]=====")

	id, _ := uuid.NewUUID()
	changeValue1 := func(w structures.Worker) {
		w.Person.Age = 14
	}
	changeValue2 := func(w *structures.Worker) {
		changeValue3 := func(p structures.Person) {
			p.Age = 14
		}

		changeValue3(w.Person)
	}
	changeValue3 := func(w *structures.Worker) {
		w.Person.Age = 14
	}

	worker := structures.Worker{
		Person: structures.Person{
			First: "Han",
			Last:  "Lee",
			Age:   30,
		},
		JobDescription: "Application Engineer",
		Salary:         10000,
		Id:             id.String(),
	}
	changeValue1(worker)
	fmt.Println(worker, "[not changed]")
	changeValue2(&worker)
	fmt.Println(worker, "[not changed]")
	changeValue3(&worker)
	fmt.Println(worker, "[changed]")
}

func test6() {
	fmt.Println("====[test6 target: pointer method]=====")

	itemP := structures.NewItem2("itemA", 1000, true)

	fmt.Println("original id =", itemP.GetId())
	id, _ := uuid.NewUUID()
	itemP.ChangeId1(id)
	fmt.Println(id.String(), "-", itemP.GetId(), "[not changed]")

	id, _ = uuid.NewUUID()
	itemP.ChangeId2(id)
	fmt.Println(id.String(), "-", itemP.GetId(), "[changed]")
}
