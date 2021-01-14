package validations

import (
	"time"

	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"

	"github.com/asaskevich/govalidator"
	valid "github.com/martinsd3v/gobrvalid"
)

var simpleValidator = map[string]struct {
	validator func(str string) bool
	message   string
}{
	"isRequired": {
		validator: govalidator.IsNotNull,
		message:   toolsErrors.ErrorRequired,
	},
	"isRequiredCreated": {
		validator: govalidator.IsNotNull,
		message:   toolsErrors.ErrorRequired,
	},
	"isRequiredUpdated": {
		validator: govalidator.IsNotNull,
		message:   toolsErrors.ErrorRequired,
	},
	"isEmail": {
		validator: govalidator.IsEmail,
		message:   toolsErrors.ErrorInvalidEmail,
	},
	"isPlate": {
		validator: valid.IsVehiclePlate,
		message:   toolsErrors.ErrorInvalidEmail,
	},
}

var validatorAndCleaner = map[string]struct {
	validator  func(str string) (bool, string)
	identifier string
	message    string
}{
	"isCpf": {
		validator: valid.IsCPF,
		message:   toolsErrors.ErrorInvalidCpf,
	},
	"isCnpj": {
		validator: valid.IsCNPJ,
		message:   toolsErrors.ErrorInvalidCnpj,
	},
	"isCpfOrCnpj": {
		validator: valid.IsCPForCNPJ,
		message:   toolsErrors.ErrorInvalidDoc,
	},
	"isCnh": {
		validator: valid.IsCNH,
		message:   toolsErrors.ErrorInvalidCnh,
	},
	"isRenavam": {
		validator: valid.IsVehicleRenavam,
		message:   toolsErrors.ErrorInvalidRenavam,
	},
}

var datesValidator = map[string]struct {
	validator  func(str string) (bool, time.Time)
	identifier string
	message    string
}{
	"isDate": {
		validator: valid.IsDate,
		message:   toolsErrors.ErrorRequired,
	},
}
