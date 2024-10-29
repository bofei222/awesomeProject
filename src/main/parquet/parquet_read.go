package main

import (
	"fmt"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"log"
)

type Person22 struct {
	Name string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Age  int32  `parquet:"name=age, type=INT32"`
}

func main() {
	fr, err := local.NewLocalFileReader("output.parquet")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(Person22), 4)
	if err != nil {
		log.Fatalf("Failed to create parquet reader: %v", err)
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	persons := make([]Person22, num)
	if err = pr.Read(&persons); err != nil {
		log.Fatalf("Read error: %v", err)
	}

	for _, person := range persons {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}
