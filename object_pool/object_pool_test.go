package objectpool

import (
	"testing"

	"github.com/google/uuid"
)

type testObj struct {
	id string
}

func (o *testObj) GetID() string {
	return o.id
}

type testObjFactory struct {
}

func (f *testObjFactory) Create() IPoolObject {
	return &testObj{
		id: uuid.New().String(),
	}
}

func TestIPool(t *testing.T) {
	tof := &testObjFactory{}
	testPool := NewPool(2, tof)
	obj1, err := testPool.Get()
	if err != nil {
		t.Errorf("Computer.Print() got error = %v ", err)
	}
	obj2, err := testPool.Get()
	if err != nil {
		t.Errorf("Computer.Print() got error = %v ", err)
	}

	if obj1.GetID() == obj2.GetID() {
		t.Errorf("Got same object ID")
	}

	obj3, got := testPool.Get()
	want := NoIdle
	if got != want {
		t.Errorf("Computer.Print() got = %v, want %v", got, want)
	}

	if obj3 != nil {
		t.Errorf("Computer.Print() obj3 = %v ", obj3)
	}

}
