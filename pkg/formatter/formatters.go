package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/krok-o/krokctl/cmd"
	"github.com/olekukonko/tablewriter"
)

type kv struct {
	Key   string
	Value string
}

// Formatter defines a formatter for the returned data.
// JSON or Table are supported for now.
type Formatter interface {
	FormatObject(data []kv) string
	FormatList(data [][]kv) string
}

// JSONFormatter takes a raw data and formats it as JSON data
type JSONFormatter struct{}

// FormatObject formats a key value list as JSON.
func (j *JSONFormatter) FormatObject(data []kv) string {
	m := make(map[string]interface{}, len(data))
	for _, kv := range data {
		m[kv.Key] = kv.Value
	}
	encoded, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		cmd.CLILog.Error().Msg("Failed to marshal map.")
		return ""
	}
	return string(encoded)
}

// FormatList formats a list of key value objects as JSON.
func (j *JSONFormatter) FormatList(data [][]kv) string {
	l := make([]map[string]interface{}, len(data))
	for i, kvs := range data {
		m := make(map[string]interface{})
		for _, kv := range kvs {
			m[kv.Key] = kv.Value
		}
		l[i] = m
	}
	encoded, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		cmd.CLILog.Error().Msg("Failed to marshal map.")
		return ""
	}
	return string(encoded)
}

// TableFormatter formats data as table entries.
type TableFormatter struct{}

// FormatObject formats a key value list as Table.
func (t *TableFormatter) FormatObject(data []kv) string {
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
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.Render()

	for _, v := range d {
		table.Append(v)
	}
	table.Render()
	return buf.String()
}

// FormatList formats a list of key value objects as Table.
func (t *TableFormatter) FormatList(data [][]kv) string {
	if len(data) == 0 {
		return ""
	}

	// Gather the headers
	header := make([]string, 0)
	for _, v := range data[0] {
		header = append(header, v.Key)
	}

	// Gather the rows
	d := [][]string{}
	for _, kvs := range data {
		row := make([]string, 0)
		for _, v := range kvs {
			row = append(row, v.Value)
		}
		d = append(d, row)
	}
	buf := &bytes.Buffer{}
	table := tablewriter.NewWriter(buf)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

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
