package utils

import (
	"os"
	"strings"
)

type CSVExporter struct {
	headers map[string]int
	rows    [][]string
}

func NewCSVExporter(path string) *CSVExporter {
	byteContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(byteContent), "\n")
	headersLine := lines[0]
	headers := strings.Split(headersLine, ",")
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[strings.TrimSpace(strings.ToLower(header))] = i
	}
	var rows [][]string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		cells := strings.Split(line, ",")
		for j, cell := range cells {
			cells[j] = strings.TrimSpace(cell)
		}
		rows = append(rows, cells)
	}
	return &CSVExporter{
		headers: headerMap,
		rows:    rows,
	}

}

func (c *CSVExporter) GetHeaderIndex(header string) (int, bool) {
	index, ok := c.headers[strings.TrimSpace(strings.ToLower(header))]
	return index, ok
}

func (c *CSVExporter) GetRow(index int) ([]string, bool) {
	if index < 0 || index >= len(c.rows) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *CSVExporter) GetRowsCount() int {
	return len(c.rows)
}

func (c *CSVExporter) GetCell(collHeader string, rowIndex int) (string, bool) {
	index, ok := c.GetHeaderIndex(collHeader)
	if !ok {
		return "", false
	}
	row, ok := c.GetRow(rowIndex)
	if !ok {
		return "", false
	}
	if index < 0 || index >= len(row) {
		return "", false
	}
	return row[index], true
}
