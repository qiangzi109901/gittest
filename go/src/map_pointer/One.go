package map_pointer

var Tpls map[string]*User


type User struct {
	Id int
	Name string
}

func init() {
	Tpls = make(map[string]*User)
}


func put(key string,id int){
	Tpls[key] = &User{id, key}
}

func Run(){
	put("java", 1)
	put("php", 2)
	put("golang", 3)
}