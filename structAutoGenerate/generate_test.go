package structAutoGenerate

import (
	"fmt"
	"testing"
)

func TestStructAutoGenerate_Generate(t *testing.T) {
	type OrmArgs struct {
		table    string
		fields   []string        // fields
		where    [][]interface{} // where
		order    string          // order
		limit    int             // limit
		offset   int             // offset
		join     [][]interface{} // join
		distinct bool            // distinct
		union    string          // sum/count/avg/max/min
		group    string          // group
		having   string          // having
		data     interface{}     // data
	}
	//var a tt
	err := New(&Option{
		Obj:         OrmArgs{},
		PackageName: "structAutoGenerate",
	}).Generate()

	fmt.Println(err)
}
