package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/volatiletech/null"
)

func main() {
	//using null string
	data := null.StringFrom("www.google.com")
	err := validation.Validate(data,
		validation.Required, // not empty
		is.URL,              // is a valid URL
	)
	fmt.Println(err)
}
