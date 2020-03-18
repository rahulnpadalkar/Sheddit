package requesthandler

import (
	"reflect"
	"sheddit/types"
	"strings"

	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v9"
)

func addcustomvalidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(ProviderExclusiveFields, types.ScheduleRequest{})
	}
}

// ProviderExclusiveFields : Validator for mandatory fileds for various providers
func ProviderExclusiveFields(structLevel validator.StructLevel) {
	sreq := structLevel.Current().Interface().(types.ScheduleRequest)
	if strings.EqualFold(sreq.Provider, "Reddit") && (sreq.Subreddits == "" || sreq.Title == "") && (sreq.Link != "" || sreq.Text != "") {
		structLevel.ReportError(reflect.ValueOf(sreq.Provider), "provider", "Provider", "exclusivefields", "")
	} else if strings.EqualFold(sreq.Provider, "Twitter") && (sreq.Text == "") {
		structLevel.ReportError(reflect.ValueOf(sreq.Provider), "provider", "Provider", "exclusivefields", "")
	}
}
