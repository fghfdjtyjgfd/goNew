package entities

import "encoding/json"

type BeerType int

const (
	WhiteBeer BeerType = iota + 1
	PaleAle
	IndianAle
	AmberAle
	Stout
	Porter
)

// Beer beer
type Beer struct {
	Model
	Name        string   `json:"name"`
	Type        BeerType `json:"-"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
}

// MarshalJSON custom AgeRange json
func (b Beer) MarshalJSON() ([]byte, error) {
	type Alias Beer
	m := &struct {
		*Alias
		TypeName string `json:"type_name"`
	}{
		Alias:    (*Alias)(&b),
		TypeName: b.Type.Name(),
	}
	return json.Marshal(m)
}

// Name name
func (s BeerType) Name() string {
	name := "-"
	switch s {
	case WhiteBeer:
		name = "Wheat Beer (วีทเบียร์)"

	case PaleAle:
		name = "Pale Ale (เพลเอล)"

	case IndianAle:
		name = "India Pale Ale (อินเดียเพลเอล)"

	case AmberAle:
		name = "Amber Ale (อัมเบอร์เอล)"

	case Stout:
		name = "Stout (สเตาท์)"

	case Porter:
		name = "Porter (พอร์เตอร์)"
	}

	return name
}
