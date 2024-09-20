package tdx

import (
	"fmt"
)

type TeePCEID string

func NewTeePCEID(val string) *TeePCEID {
	var pceID TeePCEID
	if val == "" {
		return nil
	}
	pceID = TeePCEID(val)
	return &pceID
}

func (o TeePCEID) Valid() error {
	if o == "" {
		return fmt.Errorf("nil TeePCEID")
	}
	return nil
}
