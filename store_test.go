package goleveldb

import (
	"testing"

	driver "github.com/weedge/pkg/driver/openkv"
)

func TestLeveldbStore_Implements(t *testing.T) {
	var i interface{} = &Store{}
	if _, ok := i.(driver.IStore); !ok {
		t.Fatalf("does not implement openkv.IStore")
	}
}
