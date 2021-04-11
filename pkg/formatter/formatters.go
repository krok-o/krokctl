package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/olekukonko/tablewriter"
)

type kv struct {
	Key   string
	Value string
}

// Formatter defines a formatter for the returned data.
// JSON or Table are supported for now.
type Formatter interface {
	Format(data ...kv) string
}

// JSONFormatter takes a raw data and formats it as JSON data
type JSONFormatter struct{}

// Format formats a key value list as JSON.
func (j *JSONFormatter) Format(data ...kv) string {
	m := make(map[string]interface{}, len(data))
	for _, kv := range data {
		m[kv.Key] = kv.Value
	}
	encoded, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(encoded)
}

// TableFormatter formats data as table entries.
type TableFormatter struct{}

// Format formats a key value list as Table.
func (t *TableFormatter) Format(data ...kv) string {
	row := make([]string, 0)
	header := make([]string, 0)
	for _, v := range data {
		header = append(header, v.Key)
		row = append(row, v.Value)
	}
	d := [][]string{
		row,
	}

	buf := &bytes.Buffer{}
	table := tablewriter.NewWriter(buf)
	table.SetHeader(header)

	for _, v := range d {
		table.Append(v)
	}
	table.Render()
	return buf.String()
}

// NewFormatter creates a formatter based on a set argument.
func NewFormatter(opt string) Formatter {
	if opt == "json" {
		return &JSONFormatter{}
	}

	return &TableFormatter{}
}
