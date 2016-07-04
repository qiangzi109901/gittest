package main

import (
	"time"
	"strconv"
	"strings"
	"fmt"
	"sync"
)


var allnum = 0
var downnum = 0
var lock sync.Mutex
func main() {

	test1("http://www.baidu.com")

}


func test1(url string){

	isClosed := false

	//采用广度优先
	//page1  n个link —> page2  n*m个link  -> page3 n*m*k个link
	worklist := make(chan []string)

	go func(){
		worklist <- crawl(url)
	}()

	//缓存已下载的url
	seen := make(map[string]bool)

	for list := range worklist{
		fmt.Println("receive ...")
		for _,murl := range list{
			if !seen[murl] {
				seen[murl] = true
				if isClosed {
					break
				}
				go func(url string) {
					fmt.Println("下载:" + url)
					ns := crawl(url)
					if ns == nil {
						close(worklist)
						fmt.Println("closed")
						isClosed = true
					}
					if ns != nil && !isClosed{
						fmt.Println("send....", isClosed)
						worklist <- ns

						fmt.Println("send over...")
					} else{
						fmt.Println("endend" + url)
					}
				}(murl)
			}
		}
	}
}


//根据一个url获取该页面下的所有url
func crawl(url string) (ks []string){
	if strings.Count(url, "_") == 3{
		return nil
	}
	time.Sleep(500 * time.Millisecond)
	for x:=0;x<10;x++ {
		ks = append(ks, url + "_" + strconv.Itoa(x))
	}
	return ks
}
