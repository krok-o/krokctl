package formatter

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/krok-o/krokctl/cmd"
)

type kv struct {
	Key   string
	Value interface{}
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
		cmd.CLILog.Error().Err(err).Msg("Failed to marshal map.")
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
		cmd.CLILog.Error().Err(err).Msg("Failed to marshal map.")
		return ""
	}
	return string(encoded)
}

// TableFormatter formats data as table entries.
type TableFormatter struct{}

// FormatObject formats a key value list as Table.
func (t *TableFormatter) FormatObject(data []kv) string {
	var (
		row, header []string
	)
	for _, v := range data {
		header = append(header, v.Key)
		row = append(row, convertToString(v.Value))
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
	var header []string
	for _, v := range data[0] {
		header = append(header, v.Key)
	}

	// Gather the rows
	var d [][]string
	for _, kvs := range data {
		var row []string
		for _, v := range kvs {
			row = append(row, convertToString(v.Value))
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

func dashIfEmpty(v string) string {
	if v == "" {
		return "-"
	}
	return v
}

// convertToString converts a value to string. Mostly this will be either an int or a string.
// In the future, I'm going to allow slices and format them as sub tables.
func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return dashIfEmpty(v)
	case int:
		return strconv.Itoa(v)
	default:
		cmd.CLILog.Fatal().Interface("value", v).Msg("Unknown formatting type.")
	}
	return ""
}

// NewFormatter creates a formatter based on a set argument.
func NewFormatter(opt string) Formatter {
	if opt == "json" {
		return &JSONFormatter{}
	}

	return &TableFormatter{}
}
