package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/frame/g"
)

func main() {
	db := g.DB()
	db.SetDebug(true)
	for {
		r, err := db.Table("user").All()
		fmt.Println(err)
		fmt.Println(r)
		time.Sleep(time.Second * 10)
	}
}
