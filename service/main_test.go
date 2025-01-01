package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"validator/validator"

	// kin "github.com/getkin/kin-openapi/openapi3"
	// kinf "github.com/getkin/kin-openapi/openapi3filter"
	// mux "github.com/getkin/kin-openapi/routers/gorillamux"

	"github.com/pb33f/libopenapi"
	libV "github.com/pb33f/libopenapi-validator"
)

var (
	jd         = "./testdata/jsonData.json"
	schemaJson = "./testdata/schema.json"
	oas        = "./testdata/oas.yaml"
	payload    = "./testdata/payload.json"
	becknSpec  = "./testdata/becknOAS.yaml"
)

func BenchmarkSanthoshTekuri(b *testing.B) {
	v, err := validator.NewTekuriValidator(schemaJson)
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

func BenchmarkSanthoshTekuriParallel(b *testing.B) {
	v, err := validator.NewTekuriValidator(schemaJson)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := v.Validate(bytes.NewReader(bts))
			if err != nil {
				b.Error(err)
			}
		}
	})
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
		err := v.Validate(bts)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkKaptinlinParallel(b *testing.B) {
	v, err := validator.NewKaptinlinValidator(schemaJson)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := v.Validate(bts)
			if err != nil {
				b.Error(err)
			}
		}
	})
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

func BenchmarkXeipuuvParallel(b *testing.B) {
	v, err := validator.NewXeipuuvValidator(schemaJson)
	if err != nil {
		b.Fatal(err)
	}
	bts, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := v.Validate(bts)
			if err != nil {
				b.Error(err)
			}
		}
	})
}

// func BenchmarkKinOpenAPI(b *testing.B) {
// 	// Load OpenAPI document
// 	swagger, err := kin.NewLoader().LoadFromFile(oas)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	router, _ := mux.NewRouter(swagger)
// 	// Load request data
// 	requestData, err := os.ReadFile(payload)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	// Create an HTTP request (adjust method and URL as needed)
// 	req, err := http.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(requestData))
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	route, pathParams, err := router.FindRoute(req)
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	input := &kinf.RequestValidationInput{
// 		Request:    req,
// 		Route:      route,
// 		PathParams: pathParams,
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		err := kinf.ValidateRequest(req.Context(), input)
// 		if err != nil {
// 			b.Errorf("Validation error: %v", err)
// 		}
// 	}
// }

func BenchmarkLibOpenAPI(b *testing.B) {
	spec, err := os.ReadFile(oas)

	if err != nil {
		b.Fatal(err)
	}
	doc, _ := libopenapi.NewDocument(spec)

	v, _ := libV.NewValidator(doc)
	requestData, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	request, _ := http.NewRequest(http.MethodPost, "https://things.com/burgers/createBurger",
		bytes.NewBuffer(requestData))
	request.Header.Set("Content-Type", "application/json")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// requestValid, validationErrors := v.ValidateHttpRequest(request)
		v.ValidateHttpRequest(request)

		// if !requestValid {
		// 	for i := range validationErrors {
		// 		b.Error(validationErrors[i].Message) // or sfmtomething.
		// 	}
		// }
	}
}

func BenchmarkLibOpenAPIParallel(b *testing.B) {
	spec, err := os.ReadFile(oas)
	if err != nil {
		b.Fatal(err)
	}
	doc, _ := libopenapi.NewDocument(spec)

	v, _ := libV.NewValidator(doc)
	requestData, err := os.ReadFile(jd)
	if err != nil {
		b.Fatal(err)
	}
	request, _ := http.NewRequest(http.MethodPost, "https://things.com/burgers/createBurger",
		bytes.NewBuffer(requestData))
	request.Header.Set("Content-Type", "application/json")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v.ValidateHttpRequest(request)
		}
	})
}

func BenchmarkLibOpenAPIBECKNSpec(b *testing.B) {
	highLevelValidator, req := initialse(b)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			highLevelValidator.ValidateHttpRequest(req)
		}
	})
}

func initialse(b *testing.B) (libV.Validator, *http.Request) {
	spec, err := os.ReadFile(becknSpec)

	if err != nil {
		b.Fatal(err)
	}

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

	req, err := http.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(requestData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		b.Fatal(err)
	}
	return highLevelValidator, req
}
