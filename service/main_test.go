package main

import (
	"bytes"
	"os"
	"testing"

	"validator/validator"
)

var (
	schemaFile   = "./testdata/personSchema.json"
	instanceFile = "./testdata/personData.json"
	jd           = "./testdata/jsonData.json"
	schemaJson   = "./testdata/schema.json"
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
	bts, err := os.ReadFile(instanceFile)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err := v.Validate(bts)
		if err != nil {
			b.Error(err)
		}
	}
}
