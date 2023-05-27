package goleveldb

import (
	"testing"

	driver "github.com/weedge/pkg/driver/openkv"
)

func TestLeveldbDB_Implements(t *testing.T) {
	var i interface{} = &DB{}
	if _, ok := i.(driver.IDB); !ok {
		t.Fatalf("does not implement openkv.IDB")
	}
}
