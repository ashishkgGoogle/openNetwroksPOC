package validator

import (
	"os"
	"testing"
)

var (
	schemaFile   = "./../testdata/personSchema.json"
	instanceFile = "./../testdata/personData.json"
)

func BenchmarkSanthoshTekuri(b *testing.B) {
	v, err := NewTekuriValidator(schemaFile)
	f, err := os.Open(instanceFile)
	defer f.Close()
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		err := v.Validate(f)
		if err != nil {
			b.Error(err)
		}
	}
}
