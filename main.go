package main

import "fmt"

func arrayDemo1() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "İzmir"
	sehirler[3] = "Bursa"
	sehirler[4] = "Antalya"
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
func main() {
	arrayDemo1()
}
