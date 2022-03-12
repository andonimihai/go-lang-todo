package lib

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

func HandleBindingError(err error, ctx *gin.Context) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": descriptive(verr)})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

}
