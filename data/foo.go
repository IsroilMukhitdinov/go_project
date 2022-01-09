package data

import (
	"fmt"

	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/xuri/excelize/v2"
)

func OpenAndWrite(sheetName string, temp *[]entities.Currency) {

	f := excelize.NewFile()

	// if err != nil {
	// 	fmt.Println("here 1")
	// 	return
	// }
	// defer f.Close()

	index := f.NewSheet(sheetName)

	for i, currency := range *temp {
		f.SetCellValue(sheetName, "A"+fmt.Sprint(i+1), currency.CurrencyName)
		f.SetCellValue(sheetName, "B"+fmt.Sprint(i+1), currency.CurrencyValue)
	}

	f.SetActiveSheet(index)
	f.SaveAs(sheetName + ".xlsx")

}
