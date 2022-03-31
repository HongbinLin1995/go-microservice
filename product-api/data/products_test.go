package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Hongbin",
		Price: 123.00,
		SKU:   "abs-asdas-asdasd",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
