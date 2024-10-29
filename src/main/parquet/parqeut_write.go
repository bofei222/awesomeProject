package main

import (
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"log"
)

type Person struct {
	Name string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Age  int32  `parquet:"name=age, type=INT32"`
}

func main() {
	fw, err := local.NewLocalFileWriter("output.parquet")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer fw.Close()

	pw, err := writer.NewParquetWriter(fw, new(Person), 4)
	if err != nil {
		log.Fatalf("Failed to create parquet writer: %v", err)
	}
	defer pw.WriteStop()

	pw.RowGroupSize = 128 * 1024 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	persons := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for _, person := range persons {
		if err = pw.Write(person); err != nil {
			log.Fatalf("Write error: %v", err)
		}
	}
}
