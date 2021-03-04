package models

type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}
