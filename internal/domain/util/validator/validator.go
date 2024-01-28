package validator

import (
	"fmt"
	"strings"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/locale/message"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var errorsParam = make([]internalerrors.ErrorsParam,
		len(validationErrors))

	for i, v := range validationErrors {
		key := fmt.Sprint(v.Tag())
		field := fmt.Sprint(strings.ToLower(v.Field()[0:1]) + v.Field()[1:])
		paramValue := fmt.Sprint(v.Param())

		message, err := message.GetMessage(strings.ToUpper(key), field, paramValue)
		if err != nil {
			return err
		}

		errorsParam[i] = internalerrors.ErrorsParam{Param: field, Message: message}
	}

	return internalerrors.NewErrValidation(errorsParam)
}
