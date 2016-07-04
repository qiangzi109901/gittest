package main

import (
	"gatis/src/gatis"
	"time"
	"fmt"
)

func main() {

	t := time.Now()


	gatis.Ps(t.Second(),t.Minute(),t.Hour(),t.Day(),t.Month(),t.Year(),t.UTC(),t.UnixNano(),t.Unix())

	gatis.Ps(t.String())

	gatis.Ps(t.Local())

	gatis.Ps(t.Second())

	t.Weekday()

	un := t.Unix()
	tm := time.Unix(un - 100, 0)


	gatis.Ps(tm.Format("2006-01-02"))

	t2,_ := time.Parse("2006-01-02", "2015-03-22")
	fmt.Println(t2.Format("2006/01/02"))

}
