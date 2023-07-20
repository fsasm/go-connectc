package connectc

import "testing"

func TestAddRange(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test1"
	const rangeSize = 4

	r, err := pool.AddRange(rangeName, rangeSize)
	
	if err != nil {
		t.Fatal("Could not add a variable range")
	}

	if r.Name != rangeName || r.Len() != rangeSize {
		t.Fatal("VariableRange has wrong properties")
	}
}

func TestAddDuplicate(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test2"
	const rangeSize = 4

	_, err := pool.AddRange(rangeName, rangeSize)
	
	if err != nil {
		t.Fatal("Could not add a variable range")
	}

	_, err = pool.AddRange(rangeName, rangeSize + 1)
	if err == nil {
		t.Fatal("Adding duplicates should create an error")
	}
}

func TestGetAllVars(t *testing.T) {
	var pool VariablePool

	const rangeName = "Test3"
	const rangeSize = 4

	r, err := pool.AddRange(rangeName, rangeSize)
	
	if err != nil {
		t.Fatal("Could not add a variable range")
	}

	const startIndex = 0
	for i := 0; i < rangeSize; i++ {
		index, err := r.Var(i)

		if err != nil {
			t.Fatal("Variable index is valid and should not create an error")
		}

		if index != (startIndex + i) {
			t.Fatal("Wrong index for variable")
		}
	}

	// edge cases -1 and rangeSize
	{
		_, err := r.Var(-1)

		if err == nil {
			t.Fatal("Wrong variable index should create an error")
		}
	}
	{
		_, err := r.Var(rangeSize)

		if err == nil {
			t.Fatal("Wrong variable index should create an error")
		}
	}
}

