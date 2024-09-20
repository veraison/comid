package tdx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeeMiscSelect_NewTeeMiscSelect_OK(t *testing.T) {
	tA := NewTeeMiscSelect(TestTeeMiscSelect)
	require.NotNil(t, tA)
}

func TestTeeMiscSelect_NewTeeMiscSelect_NOK(t *testing.T) {
	tA := NewTeeMiscSelect(nil)
	require.Nil(t, tA)
}

func TestNewTeeMiscSelect_Valid_OK(t *testing.T) {
	tA := TeeMiscSelect(TestTeeMiscSelect)
	err := tA.Valid()
	require.Nil(t, err)
}

func TestTeeMiscSelect_Valid_NOK(t *testing.T) {
	tA := TeeMiscSelect{}
	expectedErr := "zero len TeeMiscSelect"
	err := tA.Valid()
	assert.EqualError(t, err, expectedErr)
}
