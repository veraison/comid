// Copyright 2021-2023 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package comid

import (
	"fmt"
)

// AttestVerifKey stores an attest-key-triple-record with CBOR and JSON
// serializations.  Note that the CBOR serialization packs the structure into an
// array.  Instead, when serializing to JSON, the structure is converted into an
// object.
type AttestVerifKey struct {
	_           struct{}    `cbor:",toarray"`
	Environment Environment `json:"environment"`
	VerifKeys   CryptoKeys  `json:"verification-keys"`
}

func (o AttestVerifKey) Valid() error {

	if err := o.Environment.Valid(); err != nil {
		return fmt.Errorf("environment validation failed: %w", err)
	}
	if err := o.VerifKeys.Valid(); err != nil {
		return fmt.Errorf("verification keys validation failed: %w", err)
	}
	return nil
}
