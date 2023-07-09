package main

type Company struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	CEO     string `json:"ceo,omitempty"`
	Revenue string `json:"revenue,omitempty"`
}

// data

var Companies = []Company{
	{ID: "1", Name: "DELL", CEO: "CEO DELL", Revenue: "1000"},
	{ID: "2", Name: "LENOVO", CEO: "CEO LONOVO", Revenue: "900"},
	{ID: "3", Name: "ASUS", CEO: "CEO ASUS", Revenue: "800"},
	{ID: "4", Name: "HP", CEO: "CEO HP", Revenue: "900"},
}
