package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instanceA := GetInstance()
	instanceB := GetInstance()

	if instanceA.Get() != instanceB.Get() {
		t.Errorf("A : %v want B : %v", instanceA.Get(), instanceB.Get())
	}
}
