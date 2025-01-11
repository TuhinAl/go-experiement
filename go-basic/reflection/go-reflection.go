package reflection

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int                    `json:"id" validate:"required"`
	Name     string                 `json:"name" validate:"required"`
	Email    string                 `json:"email" validate:"required,email"`
	Age      int                    `json:"age" validate:"gte=18,lte=130"`
	IsActive bool                   `json:"is_active"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Contact struct {
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
type Person struct {
	Name    string  `json:"name" validate:"required"`
	Contact Contact `json:"contact"`
}

func Reflection() {
	user := User{ID: 125, Name: "Alauddin Tuhin", Email: "alu@example.com", Age: 29, IsActive: true, Metadata: map[string]interface{}{"book": "golang", "author": "tuhin"}}
	person := Person{
		Name: "Tanvir Hasan",
		Contact: Contact{
			Phone: "01600000000",
			Email: "tanvir@example.com",
		},
	}
	fmt.Println(ValidateReflection(user))
	fmt.Println(getContactNumber(person))
	fmt.Println(generateJSONSchema(reflect.TypeOf(user)))
	fmt.Println(buildResponse(user))
}

func getContactNumber(s interface{}) string {
	personObject := reflect.ValueOf(s)
	return personObject.FieldByIndex([]int{1, 0}).String()
}

// Dynamic field validation
func ValidateReflection(s interface{}) []string {
	var errors []string
	var input int = 120
	v := reflect.ValueOf(s)
	fmt.Println("value of int: ", reflect.ValueOf(input))
	fmt.Println("type of int: ", reflect.TypeOf(input))
	fmt.Println("kind of int: ", reflect.ValueOf(input).Kind())

	fmt.Println("type of struct: ", reflect.TypeOf(s))
	fmt.Println("kind of struct: ", reflect.ValueOf(s).Kind())

	fmt.Println("before: ", v)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		fmt.Println("middle: ", v)
	}
	typeOfValue := v.Type()
	t2 := reflect.TypeOf(s)

	fmt.Println("t2: ", t2)
	fmt.Println("after: ", v)
	fmt.Println("t: ", typeOfValue)
	fmt.Println()
	fmt.Println("Type ======: ", reflect.TypeOf(s).FieldByIndex([]int{0}).Tag.Get("json"))
	fmt.Println()
	fmt.Println("Kind: ", v.Kind())
	fmt.Println("t.NumField(): ", typeOfValue.NumField())
	for i := 0; i < typeOfValue.NumField(); i++ {
		field := typeOfValue.Field(i)
		value := v.Field(i)
		fmt.Println("Filed All Infos: ", field.Name, field.Type, "json: ", field.Tag.Get("json"), "validate: ", field.Tag.Get("validate"))
		// fmt.Println("field2 = : ", field2)
		fmt.Println("value = : ", value)
		fmt.Println("===================")
		fmt.Println("value.Interface(): ", value.Interface())
		fmt.Println("field type interface: ", reflect.Zero(field.Type).Interface())
		fmt.Println()
		fmt.Println()

		valueTag := field.Tag.Get("validate")
		if valueTag != "" {
			if value.Interface() == reflect.Zero(field.Type).Interface() {
				errors = append(errors, fmt.Sprintf("Field %s is required", field.Name))
			}
		}
	}
	if len(errors) == 0 {
		errors = append(errors, "No error found")
	}
	return errors
}


// Example 2: Dynamic JSON Generation
func generateJSONSchema(t reflect.Type) map[string]interface{} {
    schema := make(map[string]interface{})
    
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    
    schema["type"] = "object"
    properties := make(map[string]interface{})
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldSchema := make(map[string]interface{})
        
        // Get JSON tag or use field name
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            jsonTag = field.Name
        }
        
        // Determine field type
        switch field.Type.Kind() {
        case reflect.String:
            fieldSchema["type"] = "string"
        case reflect.Int:
            fieldSchema["type"] = "integer"
        case reflect.Bool:
            fieldSchema["type"] = "boolean"
        case reflect.Map:
            fieldSchema["type"] = "object"
        }
        
        properties[jsonTag] = fieldSchema
    }
    
    schema["properties"] = properties
    return schema
}

func buildResponse(data interface{}) map[string]interface{} {
    v := reflect.ValueOf(data)
    response := make(map[string]interface{})
    
    // If pointer, get the underlying element
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    
    t := v.Type()
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        // Get JSON tag or use field name
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            jsonTag = field.Name
        }
        
        response[jsonTag] = value.Interface()
    }
    
    return response
}