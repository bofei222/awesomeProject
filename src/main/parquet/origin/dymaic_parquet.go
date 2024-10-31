package origin

//import (
//	"github.com/xitongsys/parquet-go-source/local"
//	"github.com/xitongsys/parquet-go/parquet"
//	"github.com/xitongsys/parquet-go/writer"
//	"log"
//)
//
//type WindTurbineData struct {
//	Timestamp int64
//	Floats    []float32
//	Bools     []bool
//}
//
//func main() {
//	// 浮点数和布尔值字段的映射
//	floatFieldMap := map[int]string{
//		0: "MC0001", 1: "MC0002", //... 继续添加更多映射
//	}
//	boolFieldMap := map[int]string{
//		0: "MA0001", 1: "MA0002", //... 继续添加更多映射
//	}
//
//	// 创建 Parquet 文件
//	fw, err := local.NewLocalFileWriter("wind_turbine_data.parquet")
//	if err != nil {
//		log.Fatalf("Failed to create file: %v", err)
//	}
//	defer fw.Close()
//
//	// 定义 Parquet Writer
//	pw, err := writer.NewParquetWriter(fw, nil, 4)
//	if err != nil {
//		log.Fatalf("Failed to create parquet writer: %v", err)
//	}
//	defer pw.WriteStop()
//
//	pw.RowGroupSize = 128 * 1024 * 1024
//	pw.CompressionType = parquet.CompressionCodec_SNAPPY
//
//	// 设置数据模式
//	schema := &parquet.SchemaElement{Name: "root", NumChildren: 3}
//	schema.Fields = []*parquet.SchemaElement{
//		{Name: "Timestamp", Type: parquet.Type_INT64, RepetitionType: parquet.FieldRepetitionType_OPTIONAL},
//	}
//
//	for i := range floatFieldMap {
//		schema.Fields = append(schema.Fields, &parquet.SchemaElement{
//			Name: floatFieldMap[i], Type: parquet.Type_FLOAT, RepetitionType: parquet.FieldRepetitionType_OPTIONAL,
//		})
//	}
//
//	for i := range boolFieldMap {
//		schema.Fields = append(schema.Fields, &parquet.SchemaElement{
//			Name: boolFieldMap[i], Type: parquet.Type_BOOLEAN, RepetitionType: parquet.FieldRepetitionType_OPTIONAL,
//		})
//	}
//
//	// 编写数据（假设 data 包含 3600 条 WindTurbineData 结构体数据）
//	data := generateData() // 自定义生成 WindTurbineData 数据的方法
//
//	for _, record := range data {
//		recordMap := map[string]interface{}{
//			"Timestamp": record.Timestamp,
//		}
//
//		for i, value := range record.Floats {
//			recordMap[floatFieldMap[i]] = value
//		}
//
//		for i, value := range record.Bools {
//			recordMap[boolFieldMap[i]] = value
//		}
//
//		if err := pw.Write(recordMap); err != nil {
//			log.Fatalf("Failed to write record: %v", err)
//		}
//	}
//}
//
//// 自定义生成 3600 条 WindTurbineData 的方法
//func generateData() []WindTurbineData {
//	var data []WindTurbineData
//	// 填充数据逻辑
//	return data
//}
