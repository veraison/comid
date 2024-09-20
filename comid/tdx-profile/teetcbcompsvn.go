package tdx

import "fmt"

const MaxSVNCount = 16

type TeeTcbCompSvn [MaxSVNCount]TeeSVN

func NewTeeTcbCompSVN(val []uint) *TeeTcbCompSvn {
	if len(val) > MaxSVNCount || len(val) == 0 {
		return nil
	}
	TeeTcbCompSVN := make([]TeeSVN, MaxSVNCount)
	for i, value := range val {
		TeeTcbCompSVN[i] = TeeSVN(value)
	}
	return (*TeeTcbCompSvn)(TeeTcbCompSVN)
}

// nolint:gocritic
func (o TeeTcbCompSvn) Valid() error {
	if len(o) == 0 {
		return fmt.Errorf("empty TeeTcbCompSVN")
	}
	if len(o) > MaxSVNCount {
		return fmt.Errorf("invalid length: %d for TeeTcbCompSVN", len(o))
	}
	return nil
}
