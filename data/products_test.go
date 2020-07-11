package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Saurabh",
		Price: 1.00,
		SKU:   "aaa-bbb-ccc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
