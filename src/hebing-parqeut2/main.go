package main

import (
	"context"
	"fmt"
	"github.com/apache/arrow-go/v18/arrow"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
)

func mergeParquetFiles(inputDir, outputFile string) error {
	// 内存分配器
	allocator := memory.DefaultAllocator
	pool := memory.NewCheckedAllocator(allocator)
	defer pool.AssertSize(nil, 0)

	// 保存所有表的切片
	var tables []arrow.Table

	// 遍历目录获取 Parquet 文件
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".parquet" {
			// 打开 Parquet 文件
			inputFilePath := filepath.Join(inputDir, file.Name())
			reader, err := pqarrow.NewFileReader(inputFilePath, pool, pqarrow.DefaultReaderProps())
			if err != nil {
				return fmt.Errorf("failed to open parquet file %s: %w", inputFilePath, err)
			}
			ctx := context.TODO()
			table, err := reader.ReadTable(ctx)
			if err != nil {
				return fmt.Errorf("failed to read table from file %s: %w", inputFilePath, err)
			}

			tables = append(tables, table)
		}
	}

	if len(tables) == 0 {
		return fmt.Errorf("no parquet files found in directory")
	}

	// 合并 Arrow 表
	mergedTable, err := arrow.TableConcatenate(tables)
	if err != nil {
		return fmt.Errorf("failed to concatenate tables: %w", err)
	}

	// 打开输出文件
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer output.Close()

	// 写入合并后的表
	writerProps := parquet.NewWriterProperties(parquet.WithMaxRowGroupLength(100))
	writer, err := pqarrow.NewFileWriter(mergedTable.Schema(), output, writerProps, pqarrow.DefaultWriterProps())
	if err != nil {
		return fmt.Errorf("failed to create parquet writer: %w", err)
	}

	err = writer.WriteTable(mergedTable)
	if err != nil {
		return fmt.Errorf("failed to write merged table: %w", err)
	}

	return writer.Close()
}

func main() {
	inputDir := "./parquet_files"         // 输入文件目录
	outputFile := "merged_output.parquet" // 合并后的输出文件

	err := mergeParquetFiles(inputDir, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parquet files merged successfully:", outputFile)
	}
}
