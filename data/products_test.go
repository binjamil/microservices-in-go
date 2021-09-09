package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 1.0,
		SKU:   "toh-phir-ao",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
