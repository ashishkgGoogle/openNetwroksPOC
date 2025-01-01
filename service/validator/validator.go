package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	kaptinlin "github.com/kaptinlin/jsonschema"
	tekuri "github.com/santhosh-tekuri/jsonschema/v6"
	xeipuuv "github.com/xeipuuv/gojsonschema"
)

type tekuriValidator struct {
	s *tekuri.Schema
}

// type validator interface {
// 	Validate(d any) error
// }

func NewTekuriValidator(f string) (*tekuriValidator, error) {
	compiler := tekuri.NewCompiler()
	s, err := compiler.Compile(f)
	if err != nil {
		return nil, err
	}
	return &tekuriValidator{s: s}, nil
}

func (v *tekuriValidator) Validate(r io.Reader) error {

	d, err := tekuri.UnmarshalJSON(r)
	if err != nil {
		return err
	}
	return v.s.Validate(d)
}

type kaptinlinValidator struct {
	s *kaptinlin.Schema
}

func NewKaptinlinValidator(f string) (*kaptinlinValidator, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	compiler := kaptinlin.NewCompiler()
	s, err := compiler.Compile(b)
	if err != nil {
		return nil, err
	}
	return &kaptinlinValidator{s: s}, nil
}

func (v *kaptinlinValidator) Validate(b []byte) error {
	d := map[string]any{}

	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	result := v.s.Validate(d)
	if !result.IsValid() {
		details, _ := json.MarshalIndent(result.ToList(), "", "  ")
		return fmt.Errorf(string(details))
	}
	return nil
}

type xeipuuvValidator struct {
	s *xeipuuv.Schema
}

func NewXeipuuvValidator(f string) (*xeipuuvValidator, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	loader := xeipuuv.NewBytesLoader(b)
	s, err := xeipuuv.NewSchema(loader)
	if err != nil {
		return nil, err
	}
	return &xeipuuvValidator{s: s}, nil
}

func (v *xeipuuvValidator) Validate(b []byte) error {
	d := xeipuuv.NewBytesLoader(b)
	result, err := v.s.Validate(d)
	if err != nil {
		return err
	}
	if !result.Valid() {
		errs := make([]string, len(result.Errors()))
		for _, e := range result.Errors() {
			errs = append(errs, e.Description())
		}
		return fmt.Errorf(strings.Join(errs, ","))

	}
	return nil
}
