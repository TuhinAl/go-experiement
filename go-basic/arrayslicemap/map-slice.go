package arrayslicemap

import (
	"bytes"
	"fmt"
)

/*
Slice Note:
Slice Components are:
(i) A POINTER: The pointer points to the FIRST_ELEMENT of the array(months), that is reachable through Slice. FIRST_ELEMENT is not necessarily the arrays's ((months)) first element.
(ii) A LENGTH: Number ofthe slice element.
(iii) A CAPACITY: Number of element between "the start of the slice and the end of the underlying array(months)"

Important: "Multiple Slice can share the same underlying array(months) and may refer to overlapping parts of that array(months)"
*/
const (
	DEFAULT_UUID = "00000000-0000-0000-0000-000000000000"
	PI           = 3.141534234
)

func overlappingSlice() {

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "Decemmber"}

	northernSummer := months[6:9]
	quarterTwo := months[4:7]

	fmt.Println("data: len: capacity:", northernSummer, len(northernSummer), cap(northernSummer))
	fmt.Println("data: len:  capacity:", quarterTwo, len(quarterTwo), cap(quarterTwo))

	for _, summer := range northernSummer {
		for _, quarter := range quarterTwo {
			if summer == quarter {
				fmt.Printf("%s appears in both slice.", summer)
				fmt.Println()
			}
		}
	}

	// a := [...]int{0, 1, 2, 3, 4, 5}
	s := []int{0, 1, 2, 3, 4, 5}
	sliceReverse(s[2:])
	sliceReverse(s[:2])
	sliceReverse(s)
	fmt.Println(s)
}

func sliceReverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

}

func comapringTwoStringTypeSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		fmt.Println(slice1[i], " and ", slice2[i])
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func comapareByteSlice(slice1, slice2 []byte) bool {
	return bytes.Equal(slice1, slice2)
}

func checkIndirectnessOfSlice() {

	array := [...]int{10, 20, 30, 40, 50}
	slice := array[:]
	fmt.Printf("slice: %v\n", slice)
	array[1] = 1534
	fmt.Printf("slice: %v\n", slice)
}

func NilAndEmptyCheck() {

	/*
		if you need to test whether a slice is empty, use len(s) == 0, not s == nil. Other than
		comparing equal to nil, a nil slice behaves like any other zero-length slice;
	*/
	var s []int
	// s = []int{}
	// s = []int(nil)
	s = nil
	if len(s) == 0 {
		fmt.Println(s)
	}
}

func appendRune() {
	var runes []rune
	for _, r := range "Banglaesh" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

type Team struct {
	name string
	rank int
}

func ReSlicing() {
	teams := []Team{{name: "Real Madrid", rank: 1}, {name: "Manchester City", rank: 2}, {name: "Bayern Munchen", rank: 3}, {name: "Liverpool", rank: 4}, {name: "Barcelona", rank: 5}, {name: "Manchester United", rank: 6}}

	topThreeTeams := teams[0:3]
	fmt.Println("Original slice len:", len(teams))
	fmt.Println("Original slice capacity:", cap(teams))
	fmt.Println("Sub-slice len:", len(topThreeTeams))
	fmt.Println("Sub-slice capacity:", cap(topThreeTeams))
	teams = append(teams, Team{name: "Chelsea", rank: 7}, Team{name: "Dortmund", rank: 8})
}

func SliceInDetails() {
	var slice1 []string            // this is nil - 'nil'
	slice2 := []string{}           //this is empty - '{}'
	slice3 := make([]string, 5)    // length 5, capacity 5
	slice4 := make([]string, 0, 5) // length 0, capacity 5
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(slice1), cap(slice1), slice1, slice1 == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(slice2), cap(slice2), slice2, slice2 == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(slice3), cap(slice3), slice3, slice3 == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(slice4), cap(slice4), slice4, slice4 == nil)

	a := [...]int{10, 20, 30} //array
	b := a[0:1] //slice
	c := b[0:2:2] //slice

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	// fmt.Println(len(c), cap(c)) // will die: slice bounds out of range [:2] with capacity 1
	//d := c[0:1:1 ] //[i:j:k] len = j - i, cap = k - i
	// fmt.Println(d,len(d), cap(d))
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	c = append(c, 11)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
}
