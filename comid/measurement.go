// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package comid

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/veraison/eat"
	"github.com/veraison/swid"
)

const MaxUint64 = ^uint64(0)

// Measurement stores a measurement-map with CBOR and JSON serializations.
type Measurement struct {
	Key *Mkey `cbor:"0,keyasint,omitempty" json:"key,omitempty"`
	Val Mval  `cbor:"1,keyasint" json:"value"`
}

// Mkey stores a $measured-element-type-choice.
// The supported types are UUID, PSA refval-id, CCA platform-config-id and unsigned integer
// TO DO Add tagged OID: see https://github.com/veraison/corim/issues/35
type Mkey struct {
	val interface{}
}

func (o Mkey) IsSet() bool {
	return o.val != nil
}

func (o Mkey) Valid() error {
	switch t := o.val.(type) {
	case TaggedUUID:
		if UUID(t).Empty() {
			return fmt.Errorf("empty UUID")
		}
		return nil
	case TaggedPSARefValID:
		return PSARefValID(t).Valid()
	case TaggedCCAPlatformConfigID:
		if CCAPlatformConfigID(t).Empty() {
			return fmt.Errorf("empty CCAPlatformConfigID")
		}
	case uint64:
		if o.val == nil {
			return fmt.Errorf("empty uint Mkey")
		}
		return nil
	default:
		return fmt.Errorf("unknown measurement key type: %T", t)
	}
	return nil
}

func (o Mkey) IsPSARefValID() bool {
	_, ok := o.val.(TaggedPSARefValID)
	return ok
}

func (o Mkey) IsCCAPlatformConfigID() bool {
	_, ok := o.val.(TaggedCCAPlatformConfigID)
	return ok
}

func (o Mkey) GetPSARefValID() (PSARefValID, error) {
	switch t := o.val.(type) {
	case TaggedPSARefValID:
		return PSARefValID(t), nil
	default:
		return PSARefValID{}, fmt.Errorf("measurement-key type is: %T", t)
	}
}

func (o Mkey) GetCCAPlatformConfigID() (CCAPlatformConfigID, error) {
	switch t := o.val.(type) {
	case TaggedCCAPlatformConfigID:
		return CCAPlatformConfigID(t), nil
	default:
		return CCAPlatformConfigID(""), fmt.Errorf("measurement-key type is: %T", t)
	}
}

func (o Mkey) GetKeyUint() (uint64, error) {
	switch t := o.val.(type) {
	case uint64:
		return t, nil
	default:
		return MaxUint64, fmt.Errorf("measurement-key type is: %T", t)
	}
}

// UnmarshalJSON deserializes the type'n'value JSON object into the target Mkey
func (o *Mkey) UnmarshalJSON(data []byte) error {
	var v tnv

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case "uuid":
		var x UUID
		if err := x.UnmarshalJSON(v.Value); err != nil {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type UUID: %w",
				err,
			)
		}
		o.val = TaggedUUID(x)
	case "psa.refval-id":
		var x PSARefValID
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type PSARefValID: %w",
				err,
			)
		}
		if err := x.Valid(); err != nil {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type PSARefValID: %w",
				err,
			)
		}
		o.val = TaggedPSARefValID(x)
	case "cca.platform-config-id":
		var x CCAPlatformConfigID
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type CCAPlatformConfigID: %w",
				err,
			)
		}
		if x.Empty() {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type CCAPlatformConfigID: empty label",
			)
		}
		o.val = TaggedCCAPlatformConfigID(x)
	case "uint":
		var x uint64
		if err := json.Unmarshal(v.Value, &x); err != nil {
			return fmt.Errorf(
				"cannot unmarshal $measured-element-type-choice of type uint: %w",
				err,
			)
		}
		o.val = x
	default:
		return fmt.Errorf("unknown type %s for $measured-element-type-choice", v.Type)
	}

	return nil
}

// MarshalJSON serializes the target Mkey into the type'n'value JSON object
// Supported types are: uuid, psa.refval-id and unsigned integer
func (o Mkey) MarshalJSON() ([]byte, error) {
	var (
		v   tnv
		b   []byte
		err error
	)

	switch t := o.val.(type) {
	case TaggedUUID:
		uuidString := UUID(t).String()
		b, err = json.Marshal(uuidString)
		if err != nil {
			return nil, err
		}
		v = tnv{Type: "uuid", Value: b}
	case TaggedPSARefValID:
		b, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}
		v = tnv{Type: "psa.refval-id", Value: b}
	case TaggedCCAPlatformConfigID:
		b, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}
		v = tnv{Type: "cca.platform-config-id", Value: b}

	case uint64:
		b, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}
		v = tnv{Type: "uint", Value: b}

	default:
		return nil, fmt.Errorf("unknown type %T for mkey", t)
	}

	return json.Marshal(v)
}

func (o Mkey) MarshalCBOR() ([]byte, error) {
	return em.Marshal(o.val)
}

func (o *Mkey) UnmarshalCBOR(data []byte) error {
	return dm.Unmarshal(data, &o.val)
}

// Mval stores a measurement-values-map with JSON and CBOR serializations.
type Mval struct {
	Ver          *Version  `cbor:"0,keyasint,omitempty" json:"version,omitempty"`
	SVN          *SVN      `cbor:"1,keyasint,omitempty" json:"svn,omitempty"`
	Digests      *Digests  `cbor:"2,keyasint,omitempty" json:"digests,omitempty"`
	OpFlags      *OpFlags  `cbor:"3,keyasint,omitempty" json:"op-flags,omitempty"`
	RawValue     *RawValue `cbor:"4,keyasint,omitempty" json:"raw-value,omitempty"`
	RawValueMask *[]byte   `cbor:"5,keyasint,omitempty" json:"raw-value-mask,omitempty"`
	MACAddr      *MACaddr  `cbor:"6,keyasint,omitempty" json:"mac-addr,omitempty"`
	IPAddr       *net.IP   `cbor:"7,keyasint,omitempty" json:"ip-addr,omitempty"`
	SerialNumber *string   `cbor:"8,keyasint,omitempty" json:"serial-number,omitempty"`
	UEID         *eat.UEID `cbor:"9,keyasint,omitempty" json:"ueid,omitempty"`
	UUID         *UUID     `cbor:"10,keyasint,omitempty" json:"uuid,omitempty"`
}

func (o Mval) Valid() error {
	if o.Ver == nil &&
		o.SVN == nil &&
		o.Digests == nil &&
		o.OpFlags == nil &&
		o.RawValue == nil &&
		o.RawValueMask == nil &&
		o.MACAddr == nil &&
		o.IPAddr == nil &&
		o.SerialNumber == nil &&
		o.UEID == nil &&
		o.UUID == nil {
		return fmt.Errorf("no measurement value set")
	}

	if o.Ver != nil {
		if err := o.Ver.Valid(); err != nil {
			return err
		}
	}

	if o.Digests != nil {
		if err := o.Digests.Valid(); err != nil {
			return err
		}
	}

	if o.OpFlags != nil {
		if err := o.OpFlags.Valid(); err != nil {
			return err
		}
	}

	// raw value and mask have no specific semantics

	// TODO(tho) MAC addr & friends (see https://github.com/veraison/corim/issues/18)

	return nil
}

// Version stores a version-map with JSON and CBOR serializations.
type Version struct {
	Version string             `cbor:"0,keyasint" json:"value"`
	Scheme  swid.VersionScheme `cbor:"1,keyasint" json:"scheme"`
}

func NewVersion() *Version {
	return &Version{}
}

func (o *Version) SetVersion(v string) *Version {
	if o != nil {
		o.Version = v
	}
	return o
}

func (o *Version) SetScheme(v int64) *Version {
	if o != nil {
		if o.Scheme.SetCode(v) != nil {
			return nil
		}
	}
	return o
}

func (o Version) Valid() error {
	if o.Version == "" {
		return fmt.Errorf("empty version")
	}
	return nil
}

// NewMeasurement instantiates an empty measurement
func NewMeasurement() *Measurement {
	return &Measurement{}
}

// SetKeyPSARefValID sets the key of the target measurement-map to the supplied
// PSA refval-id
func (o *Measurement) SetKeyPSARefValID(psaRefValID PSARefValID) *Measurement {
	if o != nil {
		if psaRefValID.Valid() != nil {
			return nil
		}
		o.Key = &Mkey{
			val: TaggedPSARefValID(psaRefValID),
		}
	}
	return o
}

// SetKeyCCAPlatformConfigID sets the key of the target measurement-map to the supplied
// CCA platform-config-id
func (o *Measurement) SetKeyCCAPlatformConfigID(ccaPlatformConfigID CCAPlatformConfigID) *Measurement {
	if o != nil {
		if ccaPlatformConfigID.Empty() {
			return nil
		}
		o.Key = &Mkey{
			val: TaggedCCAPlatformConfigID(ccaPlatformConfigID),
		}
	}
	return o
}

// SetKeyKeyUUID sets the key of the target measurement-map to the supplied
// UUID
func (o *Measurement) SetKeyUUID(u UUID) *Measurement {
	if o != nil {
		if u.Empty() {
			return nil
		}

		if u.Valid() != nil {
			return nil
		}

		o.Key = &Mkey{
			val: TaggedUUID(u),
		}
	}
	return o
}

// SetKeyUint sets the key of the target measurement-map to the supplied
// unsigned integer
func (o *Measurement) SetKeyUint(u uint64) *Measurement {
	if o != nil {
		o.Key = &Mkey{
			val: u,
		}
	}
	return o
}

// NewPSAMeasurement instantiates a new measurement-map with the key set to the
// supplied PSA refval-id
func NewPSAMeasurement(psaRefValID PSARefValID) *Measurement {
	m := &Measurement{}
	return m.SetKeyPSARefValID(psaRefValID)
}

// NewCCAPlatCfgMeasurement instantiates a new measurement-map with the key set to the
// supplied CCA platform-config-id
func NewCCAPlatCfgMeasurement(ccaPlatformConfigID CCAPlatformConfigID) *Measurement {
	m := &Measurement{}
	return m.SetKeyCCAPlatformConfigID(ccaPlatformConfigID)
}

// NewUUIDMeasurement instantiates a new measurement-map with the key set to the
// supplied UUID
func NewUUIDMeasurement(uuid UUID) *Measurement {
	m := &Measurement{}
	return m.SetKeyUUID(uuid)
}

// NewUintMeasurement instantiates a new measurement-map with the key set to the
// supplied Uint
func NewUintMeasurement(mkey uint64) *Measurement {
	m := &Measurement{}
	return m.SetKeyUint(mkey)
}

func (o *Measurement) SetVersion(ver string, scheme int64) *Measurement {
	if o != nil {
		v := NewVersion().SetVersion(ver).SetScheme(scheme)
		if v == nil {
			return nil
		}

		o.Val.Ver = v
	}
	return o
}

// SetRawValueBytes sets the supplied raw-value and its mask in the
// measurement-values-map of the target measurement
func (o *Measurement) SetRawValueBytes(rawValue, rawValueMask []byte) *Measurement {
	if o != nil {
		o.Val.RawValue = NewRawValue().SetBytes(rawValue)
		if len(rawValueMask) != 0 {
			o.Val.RawValueMask = &rawValueMask
		}
	}
	return o
}

// SetSVN sets the supplied svn in the measurement-values-map of the target
// measurement
func (o *Measurement) SetSVN(svn uint64) *Measurement {
	if o != nil {
		s := SVN{}
		if s.SetSVN(svn) == nil {
			return nil
		}
		o.Val.SVN = &s
	}
	return o
}

// SetMinSVN sets the supplied min-svn in the measurement-values-map of the
// target measurement
func (o *Measurement) SetMinSVN(svn uint64) *Measurement {
	if o != nil {
		s := SVN{}
		if s.SetMinSVN(svn) == nil {
			return nil
		}
		o.Val.SVN = &s
	}
	return o
}

// AddDigest add the supplied digest - comprising the digest itself together
// with the hash algorithm used to obtain it - to the measurement-values-map of
// the target measurement
func (o *Measurement) AddDigest(algID uint64, digest []byte) *Measurement {
	if o != nil {
		ds := o.Val.Digests
		if ds == nil {
			ds = NewDigests()
		}
		if ds.AddDigest(algID, digest) == nil {
			return nil
		}
		o.Val.Digests = ds
	}
	return o
}

// SetOpFlags sets the supplied operational flags in the measurement-values-map
// of the target measurement
func (o *Measurement) SetOpFlags(flags ...OpFlags) *Measurement {
	if o != nil {
		o.Val.OpFlags = NewOpFlags()
		o.Val.OpFlags.SetOpFlags(flags...)
	}
	return o
}

// SetIPaddr sets the supplied IP (v4 or v6) address in the
// measurement-values-map of the target measurement
func (o *Measurement) SetIPaddr(a net.IP) *Measurement {
	if o != nil {
		o.Val.IPAddr = &a
	}
	return o
}

// SetMACaddr sets the supplied MAC address in the measurement-values-map of the
// target measurement
func (o *Measurement) SetMACaddr(a MACaddr) *Measurement {
	if o != nil {
		o.Val.MACAddr = &a
	}
	return o
}

// SetSerialNumber sets the supplied serial number in the measurement-values-map
// of the target measurement
func (o *Measurement) SetSerialNumber(sn string) *Measurement {
	if o != nil {
		o.Val.SerialNumber = &sn
	}
	return o
}

// SetUEID sets the supplied ueid in the measurement-values-map
// of the target measurement
func (o *Measurement) SetUEID(ueid eat.UEID) *Measurement {
	if o != nil {
		if ueid.Validate() != nil {
			return nil
		}
		o.Val.UEID = &ueid
	}
	return o
}

// SetUUID sets the supplied uuid in the measurement-values-map
// of the target measurement
func (o *Measurement) SetUUID(u UUID) *Measurement {
	if o != nil {
		if u.Valid() != nil {
			return nil
		}
		o.Val.UUID = &u
	}
	return o
}

func (o Measurement) Valid() error {
	if o.Key != nil && o.Key.IsSet() {
		if err := o.Key.Valid(); err != nil {
			return err
		}
	}

	// check non-empty<> condition on measurement-values-map
	if err := o.Val.Valid(); err != nil {
		return err
	}

	return nil
}
