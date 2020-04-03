package types

import (
	"database/sql"
	"encoding/json"
)

type NullableString sql.NullString

func (s *NullableString) Marshal() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	} else {
		return json.Marshal(nil)
	}
}

func (s *NullableString) Unmarshal(data []byte) error {
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
