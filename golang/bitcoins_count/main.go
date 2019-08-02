package main

import (
	"fmt"
)

func main() {
	total := 0.0
	blockinterval := 21.0		//单位：万	比特币每21万个区块减半
	currentReward := 50.0		//第一个区块50个比特币

	for currentReward > 0 {
		amount1 := blockinterval * currentReward

		currentReward *= 0.5

		total += amount1

	}

	fmt.Printf("比特币总量为:%v 万个\n", total)	//比特币总量为:2100 万个

}


