package main

import (
	"EmailGO/internal/campaign"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		fmt.Println("nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			// fmt.Printf("Error on field '%s': Tag '%s', Parameter '%s'\n", v.Field(), v.Tag(), v.Param())
			fmt.Println(v.Error())
		}
	}
}
