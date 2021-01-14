package validations

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"

	"github.com/asaskevich/govalidator"
	"github.com/martinsd3v/gobrvalid"
	valid "github.com/martinsd3v/gobrvalid"
)

//ValidateStruct responsible for validate structs
func ValidateStruct(data interface{}, prefix string, action string) (errorFields []toolsErrors.ErrorFields) {
	vl := reflect.ValueOf(data)
	return validateStruct(vl, prefix, action)
}

func validateStruct(vl reflect.Value, prefix string, action string) (errorFields []toolsErrors.ErrorFields) {

	//must be a pointer
	if vl.Kind() == reflect.Ptr || vl.Kind() == reflect.Struct {

		var elm reflect.Value

		if vl.Kind() == reflect.Ptr {
			elm = vl.Elem()
		} else {
			elm = vl
		}

		for i := 0; i < elm.NumField(); i++ {
			valueField := elm.Field(i)
			typeField := elm.Type().Field(i)
			field := elm.Type().Field(i).Name
			tagValidate := typeField.Tag.Get("validate")
			tagField := field
			prefixField := prefix

			//try to get custom tags
			if tag := elm.Type().Field(i).Tag.Get("json"); tag != "" {
				tagField = tag
			} else if tag := elm.Type().Field(i).Tag.Get("form"); tag != "" {
				tagField = tag
			}

			//Normalise prefix
			if prefixField != "" {
				prefixField = fmt.Sprintf("%s[%s]", prefix, tagField)
			} else {
				prefixField = tagField
			}

			//verify field is public
			if valid.Matches(strings.Split(field, "")[0], "[A-Z]") {
				//If the field type is a struct then validate individual
				if valueField.Kind() == reflect.Struct {
					switch valueField.Interface().(type) {
					case time.Time:
						//Struct time
						validateErrors := validateDate(prefixField, tagValidate, valueField)
						errorFields = append(errorFields, validateErrors...)
					default:
						//Simple struct
						validateErrors := validateStruct(valueField, prefixField, action)
						errorFields = append(errorFields, validateErrors...)
					}
				} else
				//If the type is a slice then loop and validate one by one
				if valueField.Kind() == reflect.Slice {
					//Simple slices
					validateErrors := validateSimpleSlice(prefixField, tagValidate, valueField, action)
					errorFields = append(errorFields, validateErrors...)
				} else {
					//Simple Field
					validateErrors := validateField(prefixField, tagValidate, valueField)
					errorFields = append(errorFields, validateErrors...)
				}
			}
		}
	}

	return
}

func validateSimpleSlice(prefix string, tagValidate string, slice reflect.Value, action string) (errorFields []toolsErrors.ErrorFields) {
	if slice.Kind() == reflect.Slice {
		elementsOfSlice := reflect.ValueOf(slice.Interface())
		quantityElements := elementsOfSlice.Len()
		tagsOfSlice := getTags(tagValidate)

		if quantityElements == 0 {
			for x := range tagsOfSlice {
				if tagsOfSlice[x] == "isRequired" {
					errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: toolsErrors.ErrorRequired})
				}
			}
		} else {
			for i := 0; i < quantityElements; i++ {
				prefixField := fmt.Sprintf("%s[%#v]", prefix, i)
				valueField := elementsOfSlice.Index(i)
				if valueField.Kind() == reflect.Struct {
					validateErrors := validateStruct(valueField, prefixField, action)
					errorFields = append(errorFields, validateErrors...)
				} else {
					validateErrors := validateField(prefixField, tagValidate, valueField)
					errorFields = append(errorFields, validateErrors...)
				}
			}
		}
	}
	return
}

func validateDate(prefix string, tagValidate string, valueField reflect.Value) (errorFields []toolsErrors.ErrorFields) {
	tags := getTags(tagValidate)
	if len(tags) >= 1 {
		//All values to string
		value := fmt.Sprint(valueField)
		//Efetuando looping nas tags e efetuando as validações
		for x := range tags {
			if tags[x] == "isRequired" {
				valid, _ := gobrvalid.IsDate(value)
				fmt.Println(gobrvalid.IsDate(value))
				if !valid {
					errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: toolsErrors.ErrorInvalidDate})
				}
			}
		}
	}
	return
}

func validateField(prefix string, tagValidate string, valueField reflect.Value) (errorFields []toolsErrors.ErrorFields) {
	tags := getTags(tagValidate)
	if len(tags) >= 1 {
		//All values to string
		value := fmt.Sprint(valueField)

		//Efetuando looping nas tags e efetuando as validações
		for x := range tags {
			if tags[x] == "isPassword" && value != "" {
				check := govalidator.IsByteLength(value, 6, 40)
				if !check {
					//Existe e não passou na validação
					errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: toolsErrors.ErrorPasswordLength})
				}
			} else if tags[x] == "isNotZero" {
				check := value != "0" && value != "0.0"
				if !check {
					//Existe e não passou na validação
					errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: toolsErrors.ErrorRequired})
				}
			} else {

				simpleValidations, exists := simpleValidator[tags[x]]

				if exists {
					check := simpleValidations.validator(value)
					if !check {
						//Existe e não passou na validação
						errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: simpleValidations.message})
					}
				}

				validationsAndCleaner, exists := validatorAndCleaner[tags[x]]

				if exists {
					check, _ := validationsAndCleaner.validator(value)
					if !check {
						//Existe e não passou na validação
						errorFields = append(errorFields, toolsErrors.ErrorFields{Field: prefix, Message: validationsAndCleaner.message})
					}
				}
			}
		}
	}
	return
}

func getTags(tagValidate string) (tags []string) {
	if tagValidate != "" {
		//Quebrando tags por |
		tags = strings.Split(tagValidate, "|")

		//Se não quebrar por | então tenta quebrar por ;
		if len(tags) == 1 {
			tags = strings.Split(tags[0], ";")
		}
	}
	return
}
