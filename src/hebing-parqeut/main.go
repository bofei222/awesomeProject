package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
)

func main() {
	// 输入目录路径和输出文件路径
	inputDir := "./parquet_files"
	outputFile := "./merged.parquet"

	err := mergeParquetFiles(inputDir, outputFile)
	if err != nil {
		log.Fatalf("Failed to merge Parquet files: %v", err)
	}
	fmt.Printf("Parquet files merged successfully to %s\n", outputFile)
}

func mergeParquetFiles(inputDir, outputFile string) error {
	// 获取目录下的所有 Parquet 文件
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var tables []arrow.Table
	mem := memory.DefaultAllocator
	ctx := context.Background()

	// 读取每个 Parquet 文件
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".parquet" {
			continue
		}

		filePath := filepath.Join(inputDir, file.Name())
		fmt.Printf("Reading file: %s\n", filePath)

		// 打开 Parquet 文件
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer f.Close()

		// 使用 pqarrow 读取表
		reader, err := file.NewParquetReader(f)
		if err != nil {
			return fmt.Errorf("failed to create parquet reader: %w", err)
		}
		defer reader.Close()

		arrowReader, err := pqarrow.NewFileReader(reader, pqarrow.ArrowReadProperties{}, mem)
		if err != nil {
			return fmt.Errorf("failed to create arrow reader: %w", err)
		}

		table, err := arrowReader.ReadTable(ctx)
		if err != nil {
			return fmt.Errorf("failed to read table from file %s: %w", filePath, err)
		}
		defer table.Release()

		tables = append(tables, table)
	}

	if len(tables) == 0 {
		return fmt.Errorf("no Parquet files found in directory %s", inputDir)
	}

	// 合并所有表
	mergedTable := mergeArrowTables(tables, mem)
	defer mergedTable.Release()

	// 将合并的表写入 Parquet 文件
	props := parquet.NewWriterProperties(parquet.WithCompression(parquet.CompressionCodec_SNAPPY))
	arrProps := pqarrow.DefaultWriterProps()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputFile, err)
	}
	defer outFile.Close()

	err = pqarrow.WriteTable(mergedTable, outFile, int64(mergedTable.NumRows()), props, arrProps)
	if err != nil {
		return fmt.Errorf("failed to write merged table: %w", err)
	}

	return nil
}

func mergeArrowTables(tables []arrow.Table, mem memory.Allocator) arrow.Table {
	if len(tables) == 1 {
		return tables[0]
	}

	// 合并所有表的列
	schema := tables[0].Schema()
	cols := make([]arrow.Column, 0, len(schema.Fields()))
	for i := 0; i < len(schema.Fields()); i++ {
		var arrays []arrow.Array
		for _, tbl := range tables {
			col := tbl.Column(i)
			for j := 0; j < col.NumChunks(); j++ {
				arrays = append(arrays, col.Chunk(j))
			}
		}

		mergedArray := array.NewChunked(schema.Field(i).Type, arrays)
		cols = append(cols, arrow.NewColumn(schema.Field(i), mergedArray))
	}

	return array.NewTable(schema, cols, tables[0].NumRows())
}
