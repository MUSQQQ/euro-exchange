package src

import (
	"github.com/xeipuuv/gojsonschema"
)

var schema = `{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
		  "table": {
			"type": "string"
		  },
		  "currency": {
			"type": "string"
		  },
		  "code": {
			"type": "string"
		  },
		  "rates": {
			"type": "array",
			"items": [
			  {
				"type": "object",
				"properties": {
				  "no": {
					"type": "string"
				  },
				  "effectiveDate": {
					"type": "string"
				  },
				  "mid": {
					"type": "number"
				  }
				},
				"required": [
				  "no",
				  "effectiveDate",
				  "mid"
				]
			  }
			]
		  }
		},
		"required": [
		  "table",
		  "currency",
		  "code",
		  "rates"
		]
	  }
}`

type validator struct {
	schema gojsonschema.JSONLoader
}

func newValidator() *validator {
	return &validator{
		schema: gojsonschema.NewStringLoader(schema),
	}
}

func (v *validator) validate(body map[string]interface{}) bool {
	documentLoader := gojsonschema.NewGoLoader(body)

	result, err := gojsonschema.Validate(v.schema, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	return result.Valid()
}
