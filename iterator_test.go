package goleveldb

import (
	"testing"

	driver "github.com/weedge/pkg/driver/openkv"
)

func TestLeveldbIter_Implements(t *testing.T) {
	var i interface{} = &Iterator{}
	if _, ok := i.(driver.IIterator); !ok {
		t.Fatalf("does not implement openkv.IIterator")
	}
}
