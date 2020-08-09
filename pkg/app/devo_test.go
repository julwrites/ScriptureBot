package app

import "testing"

func TestGetMCheyneReferences(t *testing.T) {
	options := GetMCheyneReferences()

	if len(options) == 0 {
		t.Errorf("Failed TestGetMCheyneReferences, no options retrieved")
	}
}
