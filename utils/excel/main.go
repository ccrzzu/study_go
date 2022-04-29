package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {

	//fmt.Println(strconv.Atoi("0562999"))
	f, err := excelize.OpenFile("./excel/ddc002.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "C5")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)

	//fw := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	//fmt.Println(rows, err)
	if err != nil {
		fmt.Println(rows, err)
		return
	}
	var j int = 1
	var last int = 1
	var flag bool
	for _, row := range rows {
		for i, colCell := range row {
			if i == 0 {
				if colCell == "" {
					flag = true
				}
				colCellInt, _ := strconv.Atoi(colCell)
				fmt.Println("++++++++++++++++++++++++",last,colCellInt,"=====================")
				for m := last; m < colCellInt; m++ {
					mstr := strconv.Itoa(m)
					if !strings.Contains(mstr, "4") {
						f.SetCellValue("Sheet2", "A"+strconv.Itoa(j), m)
						f.SetCellValue("Sheet2", "B"+strconv.Itoa(j), "黄色")
						j++
						fmt.Println("write success:", "A"+strconv.Itoa(j), ":", m)
					}
				}
				last = colCellInt + 1
			}
		}
		if flag {
			break
		}
	}

	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("./excel/ddc002.xlsx"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("save ddc002.xlsx success")
}
