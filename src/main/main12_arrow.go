// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/flight"
	"github.com/apache/arrow/go/v15/arrow/ipc"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"github.com/apache/arrow/go/v15/parquet"
	"github.com/apache/arrow/go/v15/parquet/pqarrow"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
)

type flightServer struct {
	mem memory.Allocator
	flight.BaseFlightServer
	Records        map[string][]arrow.Record
	uploadedChunks map[string]integrationDataSet
}

func (f *flightServer) getmem() memory.Allocator {
	if f.mem == nil {
		f.mem = memory.NewGoAllocator()
	}

	return f.mem
}

var schema = arrow.NewSchema(
	[]arrow.Field{
		{Name: "MC001", Type: arrow.FixedWidthTypes.Boolean, Nullable: false},
		{Name: "MA001", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
		{Name: "MB001", Type: arrow.PrimitiveTypes.Float32, Nullable: false},
		{Name: "timestamp", Type: arrow.PrimitiveTypes.Uint64, Nullable: false},
		{Name: "turbine", Type: arrow.BinaryTypes.String, Nullable: false},
	}, nil,
)

type integrationDataSet struct {
	schema *arrow.Schema
	chunks []arrow.Record
}

func (f *flightServer) ListFlights(c *flight.Criteria, fs flight.FlightService_ListFlightsServer) error {

	auth := ""
	authVal := flight.AuthFromContext(fs.Context())
	if authVal != nil {
		auth = authVal.(string)
	}

	recs := f.Records["scada"]
	totalRows := int64(0)
	for _, r := range recs {
		totalRows += r.NumRows()
	}

	fs.Send(&flight.FlightInfo{
		Schema: flight.SerializeSchema(schema, f.getmem()),
		FlightDescriptor: &flight.FlightDescriptor{
			Type: flight.DescriptorPATH,
			Path: []string{"scada", auth},
		},
		TotalRecords: totalRows,
		TotalBytes:   -1,
	})

	return nil
}

func (f *flightServer) DoGet(tkt *flight.Ticket, fs flight.FlightService_DoGetServer) error {
	recs, ok := f.Records[string(tkt.GetTicket())]
	if !ok {
		return status.Error(codes.NotFound, "flight not found")
	}

	w := flight.NewRecordWriter(fs, ipc.WithSchema(schema))
	for _, r := range recs {
		w.Write(r)
	}

	return nil
}

func (f *flightServer) DoPut(stream flight.FlightService_DoPutServer) error {
	rdr, err := flight.NewRecordReader(stream)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	var (
		key     string
		dataset integrationDataSet
	)

	// creating the reader should have gotten the first message which would
	// have the schema, which should have a populated flight descriptor
	desc := rdr.LatestFlightDescriptor()
	if desc.Type != flight.DescriptorPATH || len(desc.Path) < 1 {
		return status.Error(codes.InvalidArgument, "must specify a path")
	}

	key = desc.Path[0]
	dataset.schema = rdr.Schema()
	dataset.chunks = make([]arrow.Record, 0)
	for rdr.Next() {
		rec := rdr.Record()
		rec.Retain()

		dataset.chunks = append(dataset.chunks, rec)
		if len(rdr.LatestAppMetadata()) > 0 {
			stream.Send(&flight.PutResult{AppMetadata: rdr.LatestAppMetadata()})
		}
	}
	f.uploadedChunks[key] = dataset
	return nil
}

type servAuth struct{}

func (a *servAuth) Authenticate(c flight.AuthConn) error {
	tok, err := c.Read()
	if errors.Is(err, io.EOF) {
		return nil
	}

	if string(tok) != "foobar" {
		return errors.New("novalid")
	}

	if err != nil {
		return err
	}

	return c.Send([]byte("baz"))
}

func (a *servAuth) IsValid(token string) (interface{}, error) {
	if token == "baz" {
		return "bar", nil
	}
	return "", errors.New("novalid")
}

type ctxauth struct{}

type clientAuth struct{}

func (a *clientAuth) Authenticate(ctx context.Context, c flight.AuthConn) error {
	if err := c.Send(ctx.Value(ctxauth{}).([]byte)); err != nil {
		return err
	}

	_, err := c.Read()
	return err
}

func (a *clientAuth) GetToken(ctx context.Context) (string, error) {
	return ctx.Value(ctxauth{}).(string), nil
}

func main() {
	f := &flightServer{Records: make(map[string][]arrow.Record)}

	mem := memory.NewGoAllocator()

	chunks := [][]arrow.Array{
		{
			arrayOf(mem, []bool{true, false, true, false, true}),
			arrayOf(mem, []int32{10, 11, 12, 13, 14}),
			arrayOf(mem, []float32{11.11, 22.22, 33.33, 44.44, 55.55}),
			arrayOf(mem, []uint64{1000, 10002, 10003, 10004, 10005}),
			arrayOf(mem, []string{"001", "001", "001", "001", "001"}),
		},
		{
			arrayOf(mem, []bool{true, false, true, false, true}),
			arrayOf(mem, []int32{10, 11, 12, 13, 14}),
			arrayOf(mem, []float32{11.11, 22.22, 33.33, 44.44, 55.55}),
			arrayOf(mem, []uint64{1000, 10002, 10003, 10004, 10005}),
			arrayOf(mem, []string{"002", "002", "002", "002", "002"}),
		},
		{
			arrayOf(mem, []bool{true, false, true, false, true}),
			arrayOf(mem, []int32{10, 11, 12, 13, 14}),
			arrayOf(mem, []float32{11.11, 22.22, 33.33, 44.44, 55.55}),
			arrayOf(mem, []uint64{1000, 10002, 10003, 10004, 10005}),
			arrayOf(mem, []string{"002", "002", "002", "002", "002"}),
		},
	}

	defer func() {
		for _, chunk := range chunks {
			for _, col := range chunk {
				col.Release()
			}
		}
	}()

	recs := make([]arrow.Record, len(chunks))
	for i, chunk := range chunks {
		recs[i] = array.NewRecord(schema, chunk, -1)
	}

	f.Records["scada"] = recs

	ff, err := os.Create("111.parquet")
	if err != nil {
		panic(err)
	}

	props := parquet.NewWriterProperties()
	writer, err := pqarrow.NewFileWriter(schema, ff, props,
		pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}

	for _, rec := range recs {
		if err := writer.Write(rec); err != nil {
			panic(err)
		}
		rec.Release()
	}
	writer.Close()

	f.SetAuthHandler(&servAuth{})

	s := flight.NewFlightServer()
	//s.Init("localhost:0")
	s.Init("localhost:8120")
	s.RegisterFlightService(f)

	s.Serve()
	defer s.Shutdown()

}

//type flightMetadataWriterServer struct {
//  flight.BaseFlightServer
//}
//
//func (f *flightMetadataWriterServer) DoGet(tkt *flight.Ticket, fs flight.FlightService_DoGetServer) error {
//  recs := f.[string(tkt.GetTicket())]
//
//  w := flight.NewRecordWriter(fs, ipc.WithSchema(schema))
//  defer w.Close()
//  for idx, r := range recs {
//     w.WriteWithAppMetadata(r, []byte(fmt.Sprintf("%d_%s", idx, string(tkt.GetTicket()))) /*metadata*/)
//  }
//  return nil
//}

func arrayOf(mem memory.Allocator, a interface{}) arrow.Array {
	if mem == nil {
		mem = memory.NewGoAllocator()
	}

	switch a := a.(type) {

	case []bool:
		bldr := array.NewBooleanBuilder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewBooleanArray()

	case []int8:
		bldr := array.NewInt8Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewInt8Array()

	case []int16:
		bldr := array.NewInt16Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewInt16Array()

	case []int32:
		bldr := array.NewInt32Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewInt32Array()

	case []int64:
		bldr := array.NewInt64Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewInt64Array()

	case []uint8:
		bldr := array.NewUint8Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewUint8Array()

	case []uint16:
		bldr := array.NewUint16Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewUint16Array()

	case []uint32:
		bldr := array.NewUint32Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewUint32Array()

	case []uint64:
		bldr := array.NewUint64Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewUint64Array()

	case []float32:
		bldr := array.NewFloat32Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewFloat32Array()

	case []float64:
		bldr := array.NewFloat64Builder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewFloat64Array()

	case []string:
		bldr := array.NewStringBuilder(mem)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewStringArray()

	case [][]byte:
		bldr := array.NewBinaryBuilder(mem, arrow.BinaryTypes.Binary)
		defer bldr.Release()

		bldr.AppendValues(a, nil)
		return bldr.NewBinaryArray()

	//case []timestamp_s:
	// bldr := array.NewTimestampBuilder(mem, arrow.FixedWidthTypes.Timestamp_s.(*arrow.TimestampType))
	// defer bldr.Release()
	//
	// vs := make([]arrow.Timestamp, len(a))
	// for i, v := range a {
	//    vs[i] = arrow.Timestamp(v)
	// }
	// bldr.AppendValues(vs, nil)
	// return bldr.NewArray()
	//
	//case []timestamp_ms:
	// bldr := array.NewTimestampBuilder(mem, arrow.FixedWidthTypes.Timestamp_ms.(*arrow.TimestampType))
	// defer bldr.Release()
	//
	// vs := make([]arrow.Timestamp, len(a))
	// for i, v := range a {
	//    vs[i] = arrow.Timestamp(v)
	// }
	// bldr.AppendValues(vs, nil)
	// return bldr.NewArray()

	default:
		panic(fmt.Errorf("arrdata: invalid data slice type %T", a))
	}
}
