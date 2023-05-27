package goleveldb

import (
	"testing"

	driver "github.com/weedge/pkg/driver/openkv"
)

func TestLeveldbBatch_Implements(t *testing.T) {
	var i interface{} = &WriteBatch{}
	if _, ok := i.(driver.IWriteBatch); !ok {
		t.Fatalf("does not implement openkv.IWriteBatch")
	}
}
