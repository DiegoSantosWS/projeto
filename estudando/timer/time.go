package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("t é uma variavel do tipo %T\r\n", t)
	fmt.Println((t.Format(time.RFC3339)))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	dateHoursAM := t.Format("2006-01-02 15:04:05") //FORMAT AMERICAN
	fmt.Println("Data/hora atual AM: ", dateHoursAM)
	dateHoursBR := t.Format("02/01/2006 15:04:05") //FORMAT BRASIL
	fmt.Println("Data/hora atual BR: ", dateHoursBR)
}
