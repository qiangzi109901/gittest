package main

import (
	"encoding/json"
	"fmt"
	"bytes"
)

type User struct{
	Username string
	Password string
	Age int
	Gender string
}

type Person struct{
	Name string `json:"name"`
	Age int	`json:"age,omitempty"`
	Gender string `json:"gender"`
}


func main() {

	user := User{"张三","123456",20,"男"}

	jsondata,_ := json.Marshal(user)

	fmt.Println(string(jsondata))

	jsondata2,_ := json.MarshalIndent(user, "", "   ")

	fmt.Println(string(jsondata2))
	person := Person{
		Name : "李四",
		Gender : "男",
	}

	persondata,_ := json.MarshalIndent(person, "", "    ")
	fmt.Println(string(persondata))

	persons := make([]Person, 2)
	persons[0] = Person{"A",20,"f"}
	persons[1] = Person{"B",15,"m"}
	persons = append(persons, Person{"C",30,"f"})
	psdata,_ := json.MarshalIndent(persons,"","    ")

	fmt.Println(string(psdata))

	m := make(map[string]interface{})

	m["name"] = "赵六"
	m["age"] = 20
	m["gender"] = true

	mdata,_ := json.MarshalIndent(m, "", "   ")

	fmt.Println(string(mdata))


	sjson := "{\"Username\":\"张三\",\"Password\":\"123456\",\"Age\":20,\"Gender\":\"男\"}"

	bf := bytes.NewBufferString(sjson)

	suser := User{}
	json.Unmarshal(bf.Bytes(), &suser)
	fmt.Println(suser)
}
