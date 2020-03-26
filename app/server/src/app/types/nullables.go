package types

import (
	"database/sql"
	"encoding/json"
)

type NullableString sql.NullString

func (s *NullableString) Scan(value interface{}) error {
	if value != nil {
		s.String, s.Valid = string(value.([]uint8)), true
	} else {
		s.String, s.Valid = "", false
	}
	return nil
}

func (s *NullableString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	} else {
		return json.Marshal(nil)
	}
}

func (s *NullableString) UnmarshalJSON(data []byte) error {
	var tmp *string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp != nil {
		s.String, s.Valid = *tmp, true
	} else {
		s.String, s.Valid = "", false
	}
	return nil
}

func ToJsonNullString(s interface{}) NullableString {
	return NullableString{String: s.(string), Valid: s == nil}
}
