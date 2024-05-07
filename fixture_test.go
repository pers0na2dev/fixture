package fixture

import (
	"testing"
)


type TestStruct struct {
    ExportedField   int
    unexportedField int
}

func TestWithValidField(t *testing.T) {
    fixture := NewFixture[TestStruct]()
    fixture.With("ExportedField", 123)

    if build, _ := fixture.Build(); build.ExportedField != 123 {
        t.Errorf("Expected ExportedField to be 123, got %v", build.ExportedField)
    }
}

func TestWithInvalidField(t *testing.T) {
    fixture := NewFixture[TestStruct]()
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("The code panicked with %v", r)
        }
    }()

    fixture.With("NonExistentField", 123)
}

func TestWithUnsettableField(t *testing.T) {
    fixture := NewFixture[TestStruct]()
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("The code panicked with %v", r)
        }
    }()

    fixture.With("unexportedField", 123)

	if build, _ := fixture.Build(); build.unexportedField != 123 {
		t.Errorf("Expected unexportedField to be 123, got %v", build.unexportedField)
	}	
}

func TestBuildWithErrors(t *testing.T) {
	fixture := NewFixture[TestStruct]()
	fixture.With("NonExistentField", 123)

	_, err := fixture.Build()
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err.Error() != "NonExistentField: Field is not valid" {
		t.Errorf("Expected error message to be 'NonExistentField: Field is not valid', got %v", err.Error())
	}
}