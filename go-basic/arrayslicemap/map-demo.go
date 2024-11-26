package arrayslicemap

import (
	"fmt"
	"sort"
)

func main5() {
	//map define first way

	/*currencyCode := make(map[string]string)
	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"
	fmt.Println(currencyCode)

	dictionary := map[string]string{
		"Good":    "bhalo",
		"Bad":     "kharap",
		"nothing": "nothing",
	}
	fmt.Println(dictionary["Bad"])*/
	var currencyCode map[string]string
	currencyCode["USD"] = "US Dollar"
}

func mapBsic() {
	teamRanks := make(map[int]string)
	teamRanks[1] = "Real Madrid"
	teamRanks[2] = "Manchester City"
	teamRanks[3] = "Bayern Munich"
	teamRanks[4] = "Liverpool"
	teamRanks[5] = "Barcelona"
	teamRanks[6] = "PSG"

	fmt.Println(teamRanks[5])

	//delete by key
	delete(teamRanks, 5)

	//iterate Map
	for key, value := range teamRanks {
		fmt.Println(key, value)
	}
	fmt.Println(teamRanks[8])

	ages := map[string]int{"Ronaldo": 39, "Toni ross": 34, "Mesut Ozil": 42, "Kylian": 26}

	fmt.Println(ages["Kylian"])
	fmt.Println(ages["Messi"])
	// ages["Messi"] = 36
	ages["Messi"] = ages["Messi"] + 1
	fmt.Println(ages["Messi"])
	ages["Messi"]++
	fmt.Println(ages["Messi"])
}

func MapSortByKey() {
	teamRanks := make(map[string]int)
	teamRanks["Real Madrid"] = 1
	teamRanks["Manchester City"] = 2
	teamRanks["Bayern Munich"] = 3
	teamRanks["Liverpool"] = 4
	teamRanks["Barcelona"] = 5
	teamRanks["PSG"] = 6
	var teams []string
	for team, _ := range teamRanks{
		teams = append(teams, team)
	}
	sort.Strings(teams)
	for _, team := range teams{
		fmt.Println(teamRanks[team])
	}
	fmt.Println("map length: ", len(teamRanks))

	teamName, ok := teamRanks["Arsenal"]
	fmt.Println("team Name: ",teamName," ok: ", ok)
}
