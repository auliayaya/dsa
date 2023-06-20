package data_structure_algorithm

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string
	ID   int
	Age  int
}

func StudentJSON() {
	student1 := Student{
		Name: "Udin",
		ID:   1,
		Age:  5,
	}
	var dataStudent []byte
	dataStudent, err := json.Marshal(student1)
	if err != nil {
		log.Println(err)
	}
	_ = fmt.Sprintf(string(dataStudent))
}

func StudentsJSON() {
	students := []Student{
		{Name: "Rio",
			ID:  2,
			Age: 11},
		{
			Name: "Udin",
			ID:   1,
			Age:  5,
		},
	}
	for _, v := range students {
		dataStudent, err := json.Marshal(v)
		if err != nil {
			log.Println(err)
		}
		_ = fmt.Sprintf(string(dataStudent))
	}
}
