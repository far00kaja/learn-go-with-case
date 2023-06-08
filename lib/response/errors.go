package response

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BadRequest(ctx *gin.Context, err error, data interface{}) {
	var data1 = []byte(`{
		"key":"value",
		"bool":"true"
		}`)
	ctx.JSON(400,
		ResponseBadRequest{
			Code:    0,
			Message: resolveUnmarshalErr(data1, err),
		})
}

func resolveUnmarshalErr(data []byte, err error) string {
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		// var i int
		// for i = int(e.Offset) - 1; i != -1 && data[i] != '\n' && data[i] != ','; i-- {
		// }
		s := fmt.Sprintf("%s must be a %s", e.Field, e.Type)
		return s
	}
	if _, ok := err.(*json.InvalidUnmarshalError); ok {
		return "invalid unmarshal error"
	}
	if _, ok := err.(*json.SyntaxError); ok {
		return "format json invalid"
	}
	if _, ok := err.(*json.SyntaxError); ok {
		return "format json invalid syntaks"
	}
	return ErrorValidationResponse(err)
}

// func (v *)
func ErrorValidationResponse(err error) string {

	var errors []string
	if err.Error() == "EOF" {
		return "Request Body should be JSON"
	}
	for _, key := range err.(validator.ValidationErrors) {
		errorValidationMsg := fmt.Sprintf("%s %s %s", key.Field(), tagMessages(key.Tag()), key.Param())
		errors = append(errors, errorValidationMsg)
	}
	return errors[0]

}

func tagMessages(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	case "number":
		return "field must be number"
	case "gte", "min":
		return "field must be greater than"
	case "lte", "max":
		return "field must be less than"
	case "oneof":
		return "field must be one of"
	default:
		return "error"
	}
}
