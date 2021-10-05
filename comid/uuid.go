// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package comid

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUID represents an Universally Unique Identifier (UUID, see RFC4122)
type UUID uuid.UUID

// TaggedUUID is an alias to allow automatic tagging of a UUID type
type TaggedUUID UUID

// ParseUUID parses the supplied string into a UUID
func ParseUUID(s string) (UUID, error) {
	v, err := uuid.Parse(s)

	return UUID(v), err
}

// String returns a string representation of the binary UUID
func (o UUID) String() string {
	return uuid.UUID(o).String()
}

func (o UUID) Empty() bool {
	return o == (UUID{})
}

// Valid checks that the target UUID is formatted as per RFC4122
func (o UUID) Valid() error {
	if variant := uuid.UUID(o).Variant(); variant != uuid.RFC4122 {
		return fmt.Errorf("expecting RFC4122 UUID, got %s instead", variant)
	}
	return nil
}

// UnmarshalJSON deserializes the supplied string into the UUID target
// The UUID string in expected to be in canonical 8-4-4-4-12 format
func (o *UUID) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	u, err := ParseUUID(s)
	if err != nil {
		return fmt.Errorf("bad UUID: %w", err)
	}

	*o = u

	return nil
}

// MarshalJSON serialize the target UUID to a JSON string in canonical
// 8-4-4-4-12 format
func (o UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}
