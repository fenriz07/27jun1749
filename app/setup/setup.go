package setup

import (
	"reflect"
	"strings"

	"github.com/fenriz07/27jun1749/app/config"
	"github.com/go-playground/validator/v10"
	"github.com/gogearbox/gearbox"
)

func Setup() gearbox.Gearbox {

	config.LoadConfiguration()

	gb := gearBox()

	return gb
}

func SetInterface(gb gearbox.Gearbox, listUseCase usesCases) {

	v := validatorsRegistry()

	setEndpoints(gb, v, listUseCase)
}

func validatorsRegistry() *validator.Validate {

	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}
