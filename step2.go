package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	type user struct {
		spending float64
		history []float64
		todays_spending float64
	}

	

	userPoints:=make(map[int]user)
	var prev_transaction []float64
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please input users spendings in a list seperated by a comma (','). If you wish to stop inputting, please type 'stop':\t")
		spendings, _ := reader.ReadString('\n')
		if strings.Contains(spendings, "stop") {
			break
		}

		var todays_history []float64
		var counter int
		slc := strings.Split(spendings , ",")
		for i:= range slc{
			slc[i] = strings.TrimSpace(slc[i])
			currency,_:= strconv.ParseFloat(slc[i], 64)
			points := math.Floor(currency)
			todays_history = append(todays_history, float64(int(points)))
			counter += int(points)
		}
		var cumulative_spending float64 = 0
		if len(prev_transaction) > 0 {
			for i := range prev_transaction{
				cumulative_spending += float64(prev_transaction[i])
			}
		}
		var pointTally []float64
		for i := range todays_history{
			out := todays_history[i]
			if entry, ok := userPoints[i]; ok {

				history_index := len(entry.history) - 1
				ratio := float64(counter)*(float64(entry.history[history_index]))/(cumulative_spending)
				output := float64(int(out)) + float64(entry.history[history_index]) +ratio
				output = math.Floor(output*100)/100
				entry.spending = entry.spending+(output)
				entry.history = append(entry.history, (output))
				entry.todays_spending = out
				userPoints[i] = entry
				pointTally = append(pointTally, (output))
			} else {
				entry := user{spending:float64(int(out)), todays_spending:out, history:[]float64{}}
				entry.history = append(entry.history, float64(int(out)))
				entry.spending = float64(int(out))
				userPoints[i] = entry
				pointTally = append(pointTally, float64(int(out)))

			}
		}
		prev_transaction = pointTally
		
	}
	for i := range userPoints{
		fmt.Println(i, userPoints[i].history[len(userPoints[i].history)-1])
	}
}