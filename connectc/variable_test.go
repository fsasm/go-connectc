package connectc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddRange(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test1"
	const rangeSize = 4

	r, err := pool.AddRange(rangeName, rangeSize)

	require.NoError(t, err)
	require.Equal(t, rangeName, r.Name)
	require.Equal(t, rangeSize, r.Len())
}

func TestAddDuplicate(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test2"
	const rangeSize = 4

	_, err := pool.AddRange(rangeName, rangeSize)
	require.NoError(t, err)

	_, err = pool.AddRange(rangeName, rangeSize + 1)
	require.Error(t, err, "Adding duplicates should create an error")
}

func TestGetAllVars(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test3"
	const rangeSize = 4

	r, err := pool.AddRange(rangeName, rangeSize)
	require.NoError(t, err)

	const startIndex = 0
	for i := 0; i < rangeSize; i++ {
		index, err := r.Var(i)

		require.NoError(t, err, "Variable index is valid and should not create an error")
		require.Equal(t, startIndex + i, index, "Wrong index for variable")
	}

	// edge cases -1 and rangeSize
	_, err = r.Var(-1)
	require.Error(t, err, "Wrong variable index should create an error")
	_, err = r.Var(rangeSize)
	require.Error(t, err, "Wrong variable index should create an error")
}

