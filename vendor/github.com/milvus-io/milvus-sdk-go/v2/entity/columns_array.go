package entity

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
)

// ColumnVarCharArray generated columns type for VarChar
type ColumnVarCharArray struct {
	ColumnBase
	name   string
	values [][][]byte
}

// Name returns column name
func (c *ColumnVarCharArray) Name() string {
	return c.name
}

// Type returns column FieldType
func (c *ColumnVarCharArray) Type() FieldType {
	return FieldTypeArray
}

// Len returns column values length
func (c *ColumnVarCharArray) Len() int {
	return len(c.values)
}

// Get returns value at index as interface{}.
func (c *ColumnVarCharArray) Get(idx int) (interface{}, error) {
	var r []string // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnVarCharArray) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Array,
		FieldName: c.name,
	}

	data := make([]*schemapb.ScalarField, 0, c.Len())
	for _, arr := range c.values {
		converted := make([]string, 0, c.Len())
		for i := 0; i < len(arr); i++ {
			converted = append(converted, string(arr[i]))
		}
		data = append(data, &schemapb.ScalarField{
			Data: &schemapb.ScalarField_StringData{
				StringData: &schemapb.StringArray{
					Data: converted,
				},
			},
		})
	}
	fd.Field = &schemapb.FieldData_Scalars{
		Scalars: &schemapb.ScalarField{
			Data: &schemapb.ScalarField_ArrayData{
				ArrayData: &schemapb.ArrayArray{
					Data:        data,
					ElementType: schemapb.DataType_VarChar,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnVarCharArray) ValueByIdx(idx int) ([][]byte, error) {
	var r [][]byte // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnVarCharArray) AppendValue(i interface{}) error {
	v, ok := i.([][]byte)
	if !ok {
		return fmt.Errorf("invalid type, expected []string, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnVarCharArray) Data() [][][]byte {
	return c.values
}

// NewColumnVarChar auto generated constructor
func NewColumnVarCharArray(name string, values [][][]byte) *ColumnVarCharArray {
	return &ColumnVarCharArray{
		name:   name,
		values: values,
	}
}
