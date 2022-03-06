package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	userPoints := map[int] int {}
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please input users spendings in a list seperated by a comma (','). If you wish to stop inputting, please type 'stop':\t")
		spendings, _ := reader.ReadString('\n')
		if strings.Contains(spendings,"stop"){
			break
		}
		slc := strings.Split(spendings , ",")
		for i:= range slc{
			slc[i] = strings.TrimSpace(slc[i])
			currency,_:= strconv.ParseFloat(slc[i], 64)
			points := math.Floor(currency)
			if _, ok := userPoints[i]; ok {
				userPoints[i] += int(points)
			} else {
				userPoints[i] = int(points)
			}
		}
	}
	fmt.Println(userPoints)
}
