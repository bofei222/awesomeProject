package main

import (
	"context"
	"fmt"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/flight"
	"log"
)

type flightServer struct {
	flight.BaseFlightServer
}

func (f *flightServer) DoPut(ctx context.Context, stream flight.FlightService_DoPutServer) error {
	desc, err := stream.RecvDescriptor()
	if err != nil {
		return err
	}
	fmt.Printf("Received descriptor: %s\n", desc)

	for {
		record, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Received record: %v\n", record)
		record.Release() // Remember to release record after processing
	}
	return nil
}

func (f *flightServer) DoGet(ctx context.Context, ticket *flight.Ticket, stream flight.FlightService_DoGetServer) error {
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "example_field", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	}, nil)

	mem := arrow.NewAllocator()
	builder := arrow.NewInt32Builder(mem)
	defer builder.Release()

	builder.AppendValues([]int32{1, 2, 3, 4}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := arrow.NewRecord(schema, []arrow.Array{arr}, 4)
	defer record.Release()

	if err := stream.Send(record); err != nil {
		return err
	}

	return nil
}

func main() {
	server := flight.NewFlightServer()
	server.Init("localhost:0")
	server.RegisterFlightService(&flightServer{})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal(err)
		}
	}()
	defer server.Shutdown()

	fmt.Println("Server listening on:", server.Addr().String())
	select {}
}
