package tdx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initAdvisoryIDs() []any {
	s := make([]any, len(TestAdvisoryIDs))
	for i := range TestAdvisoryIDs {
		s[i] = TestAdvisoryIDs[i]
	}
	return s
}

func TestAdvisoryIDs_NewTeeAvisoryIDs_OK(t *testing.T) {
	a := initAdvisoryIDs()
	adv := NewTeeAvisoryIDs(a)
	require.NotNil(t, adv)
}

func TestAdvisoryIDs_NewTeeAvisoryIDs_NOK(t *testing.T) {
	a := make([]any, len(TestAdvisoryIDs))
	for i := range TestAdvisoryIDs {
		a[i] = i
	}
	adv := NewTeeAvisoryIDs(a)
	require.Nil(t, adv)
}

func TestAdvisoryIDs_AddAdvisoryIDs_OK(t *testing.T) {
	a := initAdvisoryIDs()
	adv := TeeAdvisoryIDs{}
	err := adv.AddTeeAdvisoryIDs(a)
	require.NoError(t, err)
}

func TestAdvisoryIDs_AddAdvisoryIDs_NOK(t *testing.T) {
	expectedErr := "invalid type: float64 for AdvisoryIDs at index: 0"
	s := make([]any, len(TestInvalidAdvisoryIDs))
	for i := range TestInvalidAdvisoryIDs {
		s[i] = TestInvalidAdvisoryIDs[i]
	}
	adv := TeeAdvisoryIDs{}
	err := adv.AddTeeAdvisoryIDs(s)
	assert.EqualError(t, err, expectedErr)
}

func TestAdvisoryIDs_Valid_OK(t *testing.T) {
	a := initAdvisoryIDs()
	adv := NewTeeAvisoryIDs(a)
	err := adv.Valid()
	require.NoError(t, err)
}

func TestAdvisoryIDs_Valid_NOK(t *testing.T) {
	expectedErr := "empty AdvisoryIDs"
	adv := TeeAdvisoryIDs{}
	err := adv.Valid()
	assert.EqualError(t, err, expectedErr)

	expectedErr = "invalid type: float64 for AdvisoryIDs at index: 0"
	s := make([]any, len(TestInvalidAdvisoryIDs))
	for i := range TestInvalidAdvisoryIDs {
		s[i] = TestInvalidAdvisoryIDs[i]
	}
	adv = TeeAdvisoryIDs(s)
	err = adv.Valid()
	assert.EqualError(t, err, expectedErr)

}
