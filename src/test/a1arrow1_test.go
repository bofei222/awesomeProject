package test

import (
	"fmt"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"testing"
)

// record batch
func Test(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "archer", Type: arrow.BinaryTypes.String},
			{Name: "location", Type: arrow.BinaryTypes.String},
			{Name: "year", Type: arrow.PrimitiveTypes.Int16},
		},
		nil,
	)

	rb := array.NewRecordBuilder(memory.DefaultAllocator, schema)
	defer rb.Release()

	rb.Field(0).(*array.StringBuilder).AppendValues([]string{"tony", "amy", "jim"}, nil)
	rb.Field(1).(*array.StringBuilder).AppendValues([]string{"beijing", "shanghai", "chengdu"}, nil)
	rb.Field(2).(*array.Int16Builder).AppendValues([]int16{1992, 1993, 1994}, nil)

	rec := rb.NewRecord()
	defer rec.Release()

	fmt.Println(rec)

	sl := rec.NewSlice(0, 2)
	fmt.Println(sl)

	cols := sl.Columns()
	a1 := cols[0]

	fmt.Println(a1)
}

// chunked
func Test2(t *testing.T) {
	ib := array.NewInt64Builder(memory.DefaultAllocator)
	defer ib.Release()

	ib.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	i1 := ib.NewInt64Array()
	defer i1.Release()

	ib.AppendValues([]int64{6, 7}, nil)
	i2 := ib.NewInt64Array()
	defer i2.Release()

	ib.AppendValues([]int64{8, 9, 10}, nil)
	i3 := ib.NewInt64Array()
	defer i3.Release()

	c := arrow.NewChunked(
		arrow.PrimitiveTypes.Int64,
		[]arrow.Array{i1, i2, i3},
	)
	defer c.Release()

	for _, arr := range c.Chunks() {
		fmt.Println(arr)
	}

	fmt.Println("chunked length =", c.Len())
	fmt.Println("chunked null count=", c.NullN())
}

// table
func Test3(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.PrimitiveTypes.Int32},
			{Name: "col2", Type: arrow.PrimitiveTypes.Float64},
			{Name: "col3", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	col1 := func() *arrow.Column {
		chunk := func() *arrow.Chunked {
			ib := array.NewInt32Builder(memory.DefaultAllocator)
			defer ib.Release()

			ib.AppendValues([]int32{1, 2, 3}, nil)
			i1 := ib.NewInt32Array()
			defer i1.Release()

			ib.AppendValues([]int32{4, 5, 6, 7, 8, 9, 10}, nil)
			i2 := ib.NewInt32Array()
			defer i2.Release()

			c := arrow.NewChunked(
				arrow.PrimitiveTypes.Int32,
				[]arrow.Array{i1, i2},
			)
			return c
		}()
		defer chunk.Release()

		return arrow.NewColumn(schema.Field(0), chunk)
	}()
	defer col1.Release()

	col2 := func() *arrow.Column {
		chunk := func() *arrow.Chunked {
			fb := array.NewFloat64Builder(memory.DefaultAllocator)
			defer fb.Release()

			fb.AppendValues([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, nil)
			f1 := fb.NewFloat64Array()
			defer f1.Release()

			fb.AppendValues([]float64{6.6, 7.7}, nil)
			f2 := fb.NewFloat64Array()
			defer f2.Release()

			fb.AppendValues([]float64{8.8, 9.9, 10.0}, nil)
			f3 := fb.NewFloat64Array()
			defer f3.Release()

			c := arrow.NewChunked(
				arrow.PrimitiveTypes.Float64,
				[]arrow.Array{f1, f2, f3},
			)
			return c
		}()
		defer chunk.Release()

		return arrow.NewColumn(schema.Field(1), chunk)
	}()
	defer col2.Release()

	col3 := func() *arrow.Column {
		chunk := func() *arrow.Chunked {
			sb := array.NewStringBuilder(memory.DefaultAllocator)
			defer sb.Release()

			sb.AppendValues([]string{"s1", "s2"}, nil)
			s1 := sb.NewStringArray()
			defer s1.Release()

			sb.AppendValues([]string{"s3", "s4"}, nil)
			s2 := sb.NewStringArray()
			defer s2.Release()

			sb.AppendValues([]string{"s5", "s6", "s7", "s8", "s9", "s10"}, nil)
			s3 := sb.NewStringArray()
			defer s3.Release()

			c := arrow.NewChunked(
				arrow.BinaryTypes.String,
				[]arrow.Array{s1, s2, s3},
			)
			return c
		}()
		defer chunk.Release()

		return arrow.NewColumn(schema.Field(2), chunk)
	}()
	defer col3.Release()

	var tbl arrow.Table
	tbl = array.NewTable(schema, []arrow.Column{*col1, *col2, *col3}, -1)
	defer tbl.Release()

	dumpTable(tbl)
}

func dumpTable(tbl arrow.Table) {
	s := tbl.Schema()
	fmt.Println(s)
	fmt.Println("------")

	fmt.Println("the count of table columns=", tbl.NumCols())
	fmt.Println("the count of table rows=", tbl.NumRows())
	fmt.Println("------")

	for i := 0; i < int(tbl.NumCols()); i++ {
		col := tbl.Column(i)
		fmt.Printf("arrays in column(%s):\n", col.Name())
		chunk := col.Data()
		for _, arr := range chunk.Chunks() {
			fmt.Println(arr)
		}
		fmt.Println("------")
	}
}
