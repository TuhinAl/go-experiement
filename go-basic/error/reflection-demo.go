package error__test

import (
	"context"
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Greet() {
	fmt.Println("Hello,", s.Name)
}

func ReflectionWithComplexType()  {
	piya := Student{"Nahida Akter", 27}
	piyaType := reflect.TypeOf(piya)

	//Field inspect
	for i := 0; i < piyaType.NumField(); i++ {
		field := piyaType.Field(i)
		fmt.Printf("Field Name: %s, Type: %s\n", field.Name, field.Type)
	}

	//Method inspect
	for i := 0; i < piyaType.NumMethod(); i++ {
		method := piyaType.Method(i)
		fmt.Printf("Method Name: %s\n", method.Name)
	}

	x := 10
	v := reflect.ValueOf(&x).Elem() // access via Pointer
	fmt.Println("Before:", x)      // Output: Before: 10

	v.SetInt(42) // change value in Runtime 
	fmt.Println("After:", x)  
}

func MapToStruct(m map[string]interface{}, s interface{}) {
    v := reflect.ValueOf(s).Elem()

    for key, value := range m {
        field := v.FieldByName(key)
        if field.IsValid() && field.CanSet() {
            field.Set(reflect.ValueOf(value))
        }
    }
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Println("UserID:", userID)
	} else {
		fmt.Println("No UserID in context")
	}
}