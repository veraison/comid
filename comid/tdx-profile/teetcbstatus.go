// nolint:dupl
package tdx

import "fmt"

type TeeTcbStatus setType

func NewTeeTcbStatus(val []any) *TeeTcbStatus {
	var ts TeeTcbStatus
	if len(val) == 0 {
		return nil
	}

	for _, v := range val {
		switch t := v.(type) {
		case string:
			ts = append(ts, t)
		default:
			return nil
		}
	}
	return &ts
}

func (o *TeeTcbStatus) AddTeeTcbStatus(val []any) error {
	for _, v := range val {
		switch t := v.(type) {
		case string:
			*o = append(*o, t)
		default:
			return fmt.Errorf("invalid type: %T for tcb status", t)
		}
	}
	return nil
}

func (o TeeTcbStatus) Valid() error {
	if len(o) == 0 {
		return fmt.Errorf("empty tcb status")
	}

	for _, v := range o {
		switch t := v.(type) {
		case string:
			continue
		default:
			return fmt.Errorf("invalid type %T for tcb status", t)
		}
	}
	return nil
}
