package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

func main() {
    excelFileName := "/Users/suxiaowen/fmstat/资产负债表/662000_资产负债表_20160731.xlsx"

    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Println(err)
    }
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows[5:] {
            for _, cell := range row.Cells {
                fmt.Print(cell.String())
            }
        }
    }

}
