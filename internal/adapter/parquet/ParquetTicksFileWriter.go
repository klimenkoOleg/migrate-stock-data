package parquet

import (
	"fmt"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type TickParquetWriter struct {
	fw source.ParquetFile
	pw *writer.ParquetWriter
}

func NewTickParquetWriter(fileName string) (*TickParquetWriter, error) {
	var err error
	fw, err := local.NewLocalFileWriter(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't create local file, err: %w", err)
	}

	pw, err := writer.NewParquetWriter(fw, new(dto.Tick1Day), 4)
	if err != nil {
		return nil, fmt.Errorf("can't create parquet writer, err: %w", err)
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.PageSize = 8 * 1024              //8K
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	return &TickParquetWriter{fw: fw, pw: pw}, nil
}

func (p *TickParquetWriter) Write(data any) error {
	if err := p.pw.Write(data); err != nil {
		return fmt.Errorf("write error, %w", err)
	}

	return nil
}

func (p *TickParquetWriter) Close() error {
	if err := p.pw.WriteStop(); err != nil {
		return fmt.Errorf("writeStop error, %w", err)
	}

	err := p.fw.Close()
	if err != nil {
		return fmt.Errorf("close error, %w", err)
	}

	return nil
}
