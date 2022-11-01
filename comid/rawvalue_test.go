package comid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawValue_NewRawValue_ok(t *testing.T) {
	tv := NewRawValue()
	require.NotNil(t, tv)
}

func TestRawValue_Set_Get_Bytes_ok(t *testing.T) {
	tv := RawValue{}
	expected := []byte{0x01, 0x02, 0x03}
	rv := tv.SetBytes([]byte{0x01, 0x02, 0x03})
	require.NotNil(t, rv)
	rval, err := rv.GetBytes()
	assert.NoError(t, err)
	assert.Equal(t, expected, rval)
}

func TestRawValue_Get_Bytes_nok(t *testing.T) {
	rv := RawValue{}
	expectedErr := "raw value is not set"
	_, err := rv.GetBytes()
	assert.EqualError(t, err, expectedErr)
	rv = RawValue{"testraw"}
	expectedErr = "unknown type string for $raw-value-type-choice"
	_, err = rv.GetBytes()
	assert.EqualError(t, err, expectedErr)
}
