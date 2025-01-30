// Copyright 2025 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package tdx

import "fmt"

type TeeMiscSelect maskType

func NewTeeMiscSelect(val []byte) *TeeMiscSelect {
	var miscSelect TeeMiscSelect
	if val == nil {
		return nil
	}
	miscSelect = TeeMiscSelect(val)
	return &miscSelect
}

func (o TeeMiscSelect) Valid() error {
	if o == nil {
		return fmt.Errorf("nil TeeMiscSelect")
	}
	if len(o) == 0 {
		return fmt.Errorf("zero len TeeMiscSelect")
	}
	return nil
}
