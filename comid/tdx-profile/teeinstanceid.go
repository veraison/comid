// nolint:dupl
package tdx

import (
	"encoding/json"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/veraison/corim/encoding"
)

// TeeInstanceID stores an TEE Instance Identifier. The supported formats are uint and variable-length bytes.
type TeeInstanceID struct {
	val interface{}
}

// NewteeInstanceID creates a new InstanceID from the
// supplied interface. The supported types are positive integers and
// byte array
func NewTeeInstanceID(val interface{}) *TeeInstanceID {
	switch t := val.(type) {
	case uint, uint64:
		return &TeeInstanceID{val: t}
	case []byte:
		return &TeeInstanceID{val: t}
	case int:
		if t < 0 {
			return nil
		}
		return &TeeInstanceID{val: t}
	default:
		return nil
	}
}

// SetTeeInstanceID sets the supplied value of Instance ID
func (o *TeeInstanceID) SetTeeInstanceID(val interface{}) error {
	switch t := val.(type) {
	case uint, uint64:
		o.val = val
	case []byte:
		o.val = val
	case int:
		if t < 0 {
			return fmt.Errorf("unsupported negative TeeInstanceID: %d", t)
		}
		o.val = val
	default:
		return fmt.Errorf("unsupported TeeInstanceID type: %T", t)
	}
	return nil
}

// valid checks for validity of TeeInstanceID and
// returns an error if Invalid
func (o TeeInstanceID) Valid() error {
	if o.val == nil {
		return fmt.Errorf("empty TeeInstanceID")
	}
	switch t := o.val.(type) {
	case uint, uint64:
		return nil
	case []byte:
		if len(t) == 0 {
			return fmt.Errorf("empty TeeInstanceID")
		}
	case int:
		if t < 0 {
			return fmt.Errorf("unsupported negative TeeInstanceID: %d", t)
		}
	default:
		return fmt.Errorf("unsupported TeeInstanceID type: %T", t)
	}
	return nil
}

func (o TeeInstanceID) GetUintTeeInstanceID() (uint, error) {
	switch t := o.val.(type) {
	case uint64:
		return uint(t), nil
	case uint:
		return t, nil
	default:
		return 0, fmt.Errorf("TeeInstanceID type is: %T", t)
	}
}

func (o TeeInstanceID) GetBytesTeeInstanceID() ([]byte, error) {
	switch t := o.val.(type) {
	case []byte:
		if len(t) == 0 {
			return nil, fmt.Errorf("TeeInstanceID type is of zero length")
		}
		return t, nil
	default:
		return nil, fmt.Errorf("TeeInstanceID type is: %T", t)
	}
}
func (o TeeInstanceID) IsBytesTeeInstanceID() bool {
	switch o.val.(type) {
	case []byte:
		return true
	default:
		return false
	}
}

func (o TeeInstanceID) IsUintTeeInstanceID() bool {
	switch o.val.(type) {
	case uint64, uint:
		return true
	default:
		return false
	}
}

func (o TeeInstanceID) MarshalJSON() ([]byte, error) {

	if o.Valid() != nil {
		return nil, fmt.Errorf("invalid TeeInstanceID")
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
		return nil, fmt.Errorf("unknown type %T for TeeInstanceID", t)
	}
	return json.Marshal(v)
}

func (o *TeeInstanceID) UnmarshalJSON(data []byte) error {
	var v encoding.TypeAndValue

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case "uint":
		var x uint
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal TeeInstanceID of type uint: %w", err)
		}
		o.val = x
	case "bytes":
		var x []byte
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal TeeInstanceID of type bytes: %w", err)
		}
		o.val = x
	}
	return nil
}
func (o TeeInstanceID) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(o.val)
}

func (o *TeeInstanceID) UnmarshalCBOR(data []byte) error {
	return cbor.Unmarshal(data, &o.val)
}
