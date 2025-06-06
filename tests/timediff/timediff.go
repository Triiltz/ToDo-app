package main

import (
	"fmt"
	"time"

	"github.com/mergestat/timediff"
)

func main() {
	then := time.Now().Add(time.Hour * -1)

	humanTime := timediff.TimeDiff(then)
	fmt.Println(humanTime)
}
