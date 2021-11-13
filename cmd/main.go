package main

import (
	"fmt"
	"github.com/buraksecer/go-bohemian-rhapsody/humanity"
	"time"
)

func main() {
	fmt.Println("Killer Queen!")
	fmt.Println(humanity.HowOldComma(time.Now().AddDate(-28,0,0)))
	fmt.Println(humanity.HowOld(time.Now().AddDate(-28,0,0)))
}
