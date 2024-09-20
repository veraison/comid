// nolint:dupl
package tdx

import "fmt"

type TeeAdvisoryIDs setType

// NewTeeAvisoryIDs create a new TeeAvisoryIDs from the
// supplied interface array and returns a pointer to
// the AdvisoryIDs. In this version only
// Advisory IDs of string type are supported
func NewTeeAvisoryIDs(val []any) *TeeAdvisoryIDs {
	var adv TeeAdvisoryIDs
	if len(val) == 0 {
		return nil
	}

	for _, v := range val {
		switch t := v.(type) {
		case string:
			adv = append(adv, t)
		default:
			return nil
		}
	}
	return &adv
}

// AddTeeAdvisoryIDs add supplied AvisoryIDs to existing AdvisoryIDs
func (o *TeeAdvisoryIDs) AddTeeAdvisoryIDs(val []any) error {
	for _, v := range val {
		switch t := v.(type) {
		case string:
			*o = append(*o, t)
		default:
			return fmt.Errorf("invalid type: %T for AdvisoryIDs", t)
		}
	}
	return nil
}

// Valid checks for validity of TeeAdvisoryIDs and
// returns an error, if invalid
func (o TeeAdvisoryIDs) Valid() error {
	if len(o) == 0 {
		return fmt.Errorf("empty AdvisoryIDs")

	}
	for _, v := range o {
		switch t := v.(type) {
		case string:
			continue
		default:
			return fmt.Errorf("invalid type: %T for AdvisoryIDs", t)
		}
	}
	return nil
}
