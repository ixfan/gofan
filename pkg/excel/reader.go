package excel

import (
	"github.com/xuri/excelize/v2"
	"io"
)

func Read(r io.Reader) ([][]string, error) {
	rows := make([][]string, 0)
	file, err := excelize.OpenReader(r)
	if err != nil {
		return rows, err
	}
	index := file.GetActiveSheetIndex()
	return file.GetRows(file.GetSheetName(index))
}
