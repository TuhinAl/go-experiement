package interface_test

import (
	"fmt"
	"sort"
)

// Explain Why composition is not inheritance
type Team struct {
	name   string
	points int
}

type Teams []Team

func (teams Teams) Len() int           { return len(teams) }
func (teams Teams) Swap(i, j int)      { teams[i], teams[j] = teams[j], teams[i] }

type ByName struct {
	Teams
}

type ByPoint struct {
	Teams
}

func (name ByName) Less(i, j int) bool {
	return name.Teams[i].name < name.Teams[j].name
}

func (point ByPoint) Less(i, j int) bool {
	return point.Teams[i].points < point.Teams[j].points
}

func ShortableInterface() {
	argentina := Team{"Argentina", 1867}
	france := Team{"France", 1859}
	spain := Team{"Spain", 1853}
	england := Team{"England", 1813}
	brazil := Team{"Brazil", 1775}
	teams := []Team{argentina, france, spain, england, brazil}
	sort.Sort(ByName{teams})
	fmt.Println(teams)
}
