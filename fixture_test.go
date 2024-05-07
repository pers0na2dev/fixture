package fixture

import (
	"testing"
)

type TestStruct struct {
	ExportedField   int
	unexportedField int
}

func TestFixture(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		value         interface{}
		expectedError string
		expectedValue int
		with          []With
	}{
		{
			name: "WithValidField", field: "ExportedField", value: 123,
			expectedValue: 123,
		},
		{
			name: "WithInvalidFieldType", field: "ExportedField", value: true,
			expectedError: "ExportedField: Type mismatch",
		},
		{
			name: "WithInvalidField", field: "NonExistentField", value: 123,
			expectedError: "NonExistentField: Field is not valid",
		},
		{
			name: "WithUnexportedField", field: "unexportedField", value: 123,
		},
		{
			name: "WithContructor", expectedValue: 456,
			with: []With{{Name: "ExportedField", Value: 456}},
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fixture := NewFixture[TestStruct](tt.with...)
			fixture.With(tt.field, tt.value)
			build, err := fixture.Build()

			if tt.expectedError != "" {
				if err == nil {
					t.Error("Expected an error, got nil")
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error message to be '%s', got %v", tt.expectedError, err.Error())
				}
			} else {
				if build.ExportedField != tt.expectedValue {
					t.Errorf("Expected %s to be %d, got %v", tt.field, tt.expectedValue, build.ExportedField)
				}
			}
		})
	}
}
