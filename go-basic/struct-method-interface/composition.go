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

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
return fmt.Sprintf("error is: %s: %s", e.path, e.err)
}

func XYZ(i int) *errFoo{
	return nil
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

	var err error = XYZ(1)
	err2 := XYZ(1)
	fmt.Println("err is: ",err)
	fmt.Println("err2 is: ",err2)
	if err2 == nil {
		fmt.Println("err2 is nill")
	}else{
		fmt.Println("err2 is NOT nill")
	}
	if err == nil {
		fmt.Println("err is nill")
	}else{
		fmt.Println("err is NOT nill")
	}
}
