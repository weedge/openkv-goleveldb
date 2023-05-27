package goleveldb

import (
	"testing"

	driver "github.com/weedge/pkg/driver/openkv"
)

func TestLeveldbSnapshot_Implements(t *testing.T) {
	var i interface{} = &Snapshot{}
	if _, ok := i.(driver.ISnapshot); !ok {
		t.Fatalf("does not implement openkv.ISnapshot")
	}
}
