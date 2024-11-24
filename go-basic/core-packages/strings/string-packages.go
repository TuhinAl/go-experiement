package strings

import (
	"fmt"
	"strings"
)

func StringDemo() {
	// Clone
	team := "Real Madrid"
	clonedName := strings.Clone(team)
	fmt.Println(team)
	fmt.Println(clonedName)

	//Contains
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}
