package myparquet

import (
	"github.com/xitongsys/parquet-go/parquet"
	"testing"
	"time"
)

type example struct {
	ID        int64 `parquet:"name=id, type=INT64"`
	CreatedAt int64 `parquet:"name=created_at,type=TIMESTAMP_MILLIS"`
}

func TestParquet(t *testing.T) {
	parquet.FileMetaData_ColumnOrders_DEFAULT
	ex := example{}
	ex.ID = int64(10)
	ex.CreatedAt = time.Now().Unix()

	fw, err := ParquetFile.NewLocalFileWriter("new.parquet")
	pw, err := ParquetWriter.NewParquetWriter(fw, new(example), 1)
	pw.Write(ex)

}
