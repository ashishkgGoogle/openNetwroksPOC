package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"validator/validator"

	kin "github.com/getkin/kin-openapi/openapi3"
	kinf "github.com/getkin/kin-openapi/openapi3filter"
	mux "github.com/getkin/kin-openapi/routers/gorillamux"

	"github.com/pb33f/libopenapi"
	libV "github.com/pb33f/libopenapi-validator"
)

var (
	jd         = "./testdata/jsonData.json"
	schemaJson = "./testdata/schema.json"
	oas        = "./testdata/openapi.yaml"
	payload    = "./testdata/payload.json"
)

func BenchmarkSanthoshTekuri(b *testing.B) {
	v, err := validator.NewTekuriValidator(oas)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(payload)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := v.Validate(bytes.NewReader(bts))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkKaptinlin(b *testing.B) {
	v, err := validator.NewKaptinlinValidator(schemaJson)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := v.Validate(bytes.NewReader(bts))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkXeipuuv(b *testing.B) {
	v, err := validator.NewXeipuuvValidator(schemaJson)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := v.Validate(bts)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkKinOpenAPI(b *testing.B) {
	// Load OpenAPI document
	swagger, err := kin.NewLoader().LoadFromFile(oas)
	if err != nil {
		b.Fatal(err)
	}

	router, _ := mux.NewRouter(swagger)
	// Load request data
	requestData, err := os.ReadFile(payload)
	if err != nil {
		b.Fatal(err)
	}

	// Create an HTTP request (adjust method and URL as needed)
	req, err := http.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(requestData))
	if err != nil {
		b.Fatal(err)
	}
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		b.Fatal(err)
	}
	input := &kinf.RequestValidationInput{
		Request:    req,
		Route:      route,
		PathParams: pathParams,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := kinf.ValidateRequest(req.Context(), input)
		if err != nil {
			b.Errorf("Validation error: %v", err)
		}
	}
}

func BenchmarkLibOpenAPI(b *testing.B) {
	spec, err := os.ReadFile(oas)

	if err != nil {
		b.Fatal(err)
	}

	// 2. Create a new OpenAPI document using libopenapi
	document, docErrs := libopenapi.NewDocument(spec)
	if docErrs != nil {
		b.Fatal(docErrs)
	}

	highLevelValidator, validatorErrs := libV.NewValidator(document)
	if len(validatorErrs) > 0 {
		b.Fatal("document is bad")
	}
	requestData, err := os.ReadFile(payload)
	if err != nil {
		b.Fatal(err)
	}

	// Create an HTTP request (adjust method and URL as needed)
	req, err := http.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(requestData))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		requestValid, validationErrors := highLevelValidator.ValidateHttpRequest(req)

		if !requestValid {
			for i := range validationErrors {
				b.Error(validationErrors[i].Message) // or something.
			}
		}
	}
}
