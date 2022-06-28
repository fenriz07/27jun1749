package setup

import (
	"github.com/fenriz07/27jun1749/app/config"
	"github.com/gogearbox/gearbox"
)

func Setup() gearbox.Gearbox {

	config.LoadConfiguration()

	gb := gearBox()

	return gb
}

func SetInterface(gb gearbox.Gearbox, listUseCase usesCases) {

	setEndpoints(gb, listUseCase)
}

/*

func ValidatorsRegistry() *validator.Validate {

	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("cmref", ValidateCmRef)

	return validate
}

func ValidateCmRef(fl validator.FieldLevel) bool {

	field := fl.Field()
	var v string

	switch field.Kind() {
	case reflect.String:
		v = field.String()
	default:
		return false
	}

	listCmRefAllowed := config.ListCmRefAllowed()

	for _, cmref := range listCmRefAllowed {
		if cmref == v {
			return true
		}
	}

	return false
}
*/
