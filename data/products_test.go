package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "CokkKK",
		Price: 1.0,
		SKU:   "asc-sd-sd",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
