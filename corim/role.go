// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package corim

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Role int64

const (
	RoleManifestCreator Role = iota + 1
)

var (
	stringToRole = map[string]Role{
		"manifestCreator": RoleManifestCreator,
	}
	roleToString = map[Role]string{
		RoleManifestCreator: "manifestCreator",
	}
)

type Roles []Role

func NewRoles() *Roles {
	return &Roles{}
}

// Add appends the supplied roles to Roles list.
func (o *Roles) Add(roles ...Role) *Roles {
	if o != nil {
		for _, r := range roles {
			if !isRole(r) {
				return nil
			}
			*o = append(*o, r)
		}
	}
	return o
}

func isRole(r Role) bool {
	return r == RoleManifestCreator
}

// Valid iterates over the range of individual roles to check for validity
func (o Roles) Valid() error {
	if len(o) == 0 {
		return errors.New("empty roles")
	}

	for i, r := range o {
		if !isRole(r) {
			return fmt.Errorf("unknown role %d at index %d", r, i)
		}
	}

	return nil
}

func (o *Roles) UnmarshalJSON(data []byte) error {
	var a []string

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	for _, s := range a {
		r, ok := stringToRole[s]
		if !ok {
			return fmt.Errorf("unknown role %q", s)
		}
		o = o.Add(r)
	}

	return nil
}

func (o Roles) MarshalJSON() ([]byte, error) {
	roles := []string{}

	for _, r := range o {
		s, ok := roleToString[r]
		if !ok {
			return nil, fmt.Errorf("unknown role %d", r)
		}
		roles = append(roles, s)
	}

	return json.Marshal(roles)
}

func (o Roles) ToJSON() ([]byte, error) {
	if err := o.Valid(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("encoding failed: %w", err)
	}

	return data, err
}

func (o *Roles) FromJSON(data []byte) error {
	if err := o.UnmarshalJSON(data); err != nil {
		return fmt.Errorf("decoding failed: %w", err)
	}

	err := o.Valid()
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return err
}
