package main

/*import (
	"context"
	"fmt"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/compute"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// max_aggregate_function.go
func main() {
	data := []int64{5, 10, 0, 25, 2, 35, 7, 15}
	bldr := array.NewInt64Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues(data, nil)
	arr := bldr.NewArray()
	defer arr.Release()

	dat, err := compute.Max(context.Background(), compute.NewDatum(arr))

	if err != nil {
		fmt.Println(err)
		return
	}

	ad, ok := dat.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	arr1 := ad.MakeArray()
	fmt.Println(arr1) // [35]
}
*/
