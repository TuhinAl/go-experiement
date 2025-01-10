package reflection

import (
	"fmt"
	"reflect"
)

type User struct {
    ID        int    `json:"id" validate:"required"`
    Name      string `json:"name" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Age       int    `json:"age" validate:"gte=0,lte=130"`
    IsActive  bool   `json:"is_active"`
    Metadata  map[string]interface{} `json:"metadata"`
}

func Reflection() {
	user := User{ID: 1, Name: "Tuhin", Email: "alu@example.com", Age: 27, IsActive: true, Metadata: map[string]interface{}{"book": "golang", "author": "tuhin"}}
	ValidateReflection(user)
}

func ValidateReflection(s interface{}) {
	// var errors []string;
	v := reflect.ValueOf(s)
	fmt.Print(v)
	fmt.Println()
	fmt.Print("Type: ", reflect.TypeOf(s))
	fmt.Println()
	fmt.Print("Kind: ", v.Kind())
}
