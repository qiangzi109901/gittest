package main

import (
	"fmt"
)

func main() {
	

	m := make(map[string]string)

	m["name"] = "张三"
	m["age"] = "20"

	fmt.Println(m)


	mk := map[string]int{
		"A" : 65,
		"Z" : 91,
	}

	mk["0"] = 48

	fmt.Println(mk)

	ma := map[string]int{}

	ma["JAVA"] = 1
	ma["GO"] = 2
	ma["GROOVY"] = 3

	fmt.Println(ma)



	delete(ma, "GROOVY")
	fmt.Println(ma)

	ma["JAVA"] += 1

	fmt.Println(ma)

	ma["GO"] ++
	fmt.Println(ma)


	for name, loves := range ma {
		fmt.Printf("name : %s\t\tloves : %d \n", name, loves)
	}


	mk2 := make(map[string]int)

	mk2["ajva"] = 1
	//mk2["gis"] = 2

	fmt.Println(len(mk2))


	tt := map[string]interface{}{}

	fmt.Println(tt["java"])


	//ttk["java"] = 1
	//fmt.Println(ttk)

}

func getMapOne(m map[string]interface{}) interface{}{

	if len(m) != 1 {
		return m
	}
	for _,v := range m{
		return v
	}

	return m
}