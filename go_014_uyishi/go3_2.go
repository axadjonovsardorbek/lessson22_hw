package main

import (
	"fmt"
	"time"
)

func Hour(t1, t2 int){
	fmt.Println(t2-t1, "soat")
}
func Minut(t1, t2 int){
	fmt.Println(t2-t1, "minut")
}
func Second(t1, t2 int){
	fmt.Println(t2-t1, "sikund")
}

func main(){

	old := time.Date(2000, 1, 1, 1, 3, 3, 0, time.UTC)
	now := time.Now()

	fmt.Println(old)
	fmt.Println(now)

	Hour(old.Hour(), now.Hour())
	Minut(old.Minute(), now.Minute())
	Second(old.Second(), now.Second())
}