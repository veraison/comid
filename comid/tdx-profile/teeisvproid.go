// nolint:dupl
package tdx

import (
	"encoding/json"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/veraison/corim/encoding"
)

// TeeISVProdID stores an ISV Product Identifier. The supported formats are uint and variable-length bytes.
type TeeISVProdID struct {
	val interface{}
}

// NewTeeISVProdID creates a new TeeISVProdID from the
// supplied interface and return a pointer to TeeISVProdID
// Supported values are positive integers and byte array
func NewTeeISVProdID(val interface{}) *TeeISVProdID {
	switch t := val.(type) {
	case uint, uint64:
		return &TeeISVProdID{val: t}
	case []byte:
		return &TeeISVProdID{val: t}
	case int:
		if t < 0 {
			return nil
		}
		return &TeeISVProdID{val: t}
	default:
		return nil
	}
}

// SetTeeISVProdID sets the supplied value of TeeISVProdID from the interface
// Supported values are either positive integers or byte array
func (o *TeeISVProdID) SetTeeISVProdID(val interface{}) error {
	switch t := val.(type) {
	case uint, uint64:
		o.val = val
	case []byte:
		o.val = val
	case int:
		if t < 0 {
			return fmt.Errorf("unsupported negative TeeISVProdID: %d", t)
		}
		o.val = val
	default:
		return fmt.Errorf("unsupported TeeISVProdID type: %T", t)
	}
	return nil
}

// Valid checks for validity of TeeISVProdID and returns an error if Invalid
func (o TeeISVProdID) Valid() error {
	if o.val == nil {
		return fmt.Errorf("empty TeeISVProdID")
	}
	switch t := o.val.(type) {
	case uint, uint64:
		return nil
	case []byte:
		if len(t) == 0 {
			return fmt.Errorf("empty TeeISVProdID")
		}
	case int:
		if t < 0 {
			return fmt.Errorf("unsupported negative TeeISVProdID: %d", t)
		}
	default:
		return fmt.Errorf("unsupported TeeISVProdID type: %T", t)
	}
	return nil
}

// GetUint returns a uint TeeISVProdID
func (o TeeISVProdID) GetUint() (uint, error) {
	switch t := o.val.(type) {
	case uint64:
		return uint(t), nil
	case uint:
		return t, nil
	default:
		return 0, fmt.Errorf("TeeISVProdID type is: %T", t)
	}
}

// GetBytes returns a []byte TeeISVProdID
func (o TeeISVProdID) GetBytes() ([]byte, error) {
	switch t := o.val.(type) {
	case []byte:
		if len(t) == 0 {
			return nil, fmt.Errorf("TeeISVProdID type is of zero length")
		}
		return t, nil
	default:
		return nil, fmt.Errorf("TeeIsvProdID type is: %T", t)
	}
}

// IsBytes returns true if TeeISVProdID is a byte array
func (o TeeISVProdID) IsBytes() bool {
	switch o.val.(type) {
	case []byte:
		return true
	default:
		return false
	}
}

// IsUint returns true if TeeISVProdID is a positive integer
func (o TeeISVProdID) IsUint() bool {
	switch t := o.val.(type) {
	case uint64, uint:
		return true
	case int:
		if t < 0 {
			return false
		}
		return true
	default:
		return false
	}
}

// MarshalJSON Marshals TeeISVProdID to JSON
func (o TeeISVProdID) MarshalJSON() ([]byte, error) {
	if o.Valid() != nil {
		return nil, fmt.Errorf("invalid TeeISVProdID")
	}
	var (
		v   encoding.TypeAndValue
		b   []byte
		err error
	)
	switch t := o.val.(type) {
	case uint, uint64, int:
		b, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}
		v = encoding.TypeAndValue{Type: "uint", Value: b}
	case []byte:
		b, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}
		v = encoding.TypeAndValue{Type: "bytes", Value: b}
	default:
		return nil, fmt.Errorf("unknown type %T for TeeISVProdID", t)
	}
	return json.Marshal(v)
}

// UnmarshalJSON UnMarshals supplied JSON buffer to TeeISVProdID
func (o *TeeISVProdID) UnmarshalJSON(data []byte) error {
	var v encoding.TypeAndValue

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case "uint":
		var x uint
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal TeeISVProdID of type uint: %w", err)
		}
		o.val = x
	case "bytes":
		var x []byte
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal TeeISVProdID of type bytes: %w", err)
		}
		o.val = x
	}
	return nil
}

// MarshalCBOR Marshals TeeISVProdID to CBOR bytes
func (o TeeISVProdID) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(o.val)
}

// UnmarshalCBOR UnMarshals supplied CBOR bytes to TeeISVProdID
func (o *TeeISVProdID) UnmarshalCBOR(data []byte) error {
	return cbor.Unmarshal(data, &o.val)
}
