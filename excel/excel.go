package excel

import (
	"github.com/xuri/excelize/v2"
	"jira-get-tickets/structs"
	"log"
	"strconv"
	"time"
)

const sheetName = "tickets"

func ConvertToExcel(data *structs.TicketsData) error {
	f := excelize.NewFile()

	f.NewSheet(sheetName)
	i, _ := f.GetSheetIndex(sheetName)
	f.SetActiveSheet(i)
	f.DeleteSheet("Sheet1")

	if err := setGeneralStyle(f, data.Total); err != nil {
		log.Fatalf(err.Error())
	}

	if err := fillData(f, data); err != nil {
		log.Fatalf(err.Error())
	}

	if err := f.SaveAs(GetSpreadsheetName()); err != nil {
		return err
	}
	return nil
}

func setGeneralStyle(f *excelize.File, i int) error {
	blueStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Font: &excelize.Font{Size: 12},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#DCE6F1"},
			Pattern: 1,
		},
	})

	whiteStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Font: &excelize.Font{Size: 12},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"},
			Pattern: 1,
		},
	})

	if err != nil {
		return err
	}

	for row := 0; row <= i+1; row++ {

		if row%2 == 0 {
			f.SetRowStyle(sheetName, row, row, whiteStyle)
		} else {
			f.SetRowStyle(sheetName, row, row, blueStyle)
		}
	}

	f.SetColWidth(sheetName, "A", "F", 20)
	f.SetColWidth(sheetName, "F", "F", 34)
	f.SetColWidth(sheetName, "G", "I", 20)
	f.SetColWidth(sheetName, "I", "J", 34)

	f.SetRowHeight(sheetName, 2, 100)
	f.SetRowHeight(sheetName, 3, 100)

	return nil
}

func fillData(f *excelize.File, data *structs.TicketsData) error {

	//columns := 8
	rows := data.Total

	f.SetCellValue(sheetName, "A1", "Key")
	f.SetCellValue(sheetName, "B1", "Issue Type")
	f.SetCellValue(sheetName, "C1", "Priority")
	f.SetCellValue(sheetName, "D1", "Severity")
	f.SetCellValue(sheetName, "E1", "Component")
	f.SetCellValue(sheetName, "F1", "Summary")
	f.SetCellValue(sheetName, "G1", "Assignee")
	f.SetCellValue(sheetName, "H1", "Status")
	f.SetCellValue(sheetName, "I1", "Created")
	f.SetCellValue(sheetName, "J1", "Updated")

	for i := 0; i < rows; i++ {

		istr := strconv.Itoa(i + 2)

		f.SetCellValue(sheetName, "A"+istr, data.Issues[i].Key)
		f.SetCellValue(sheetName, "B"+istr, data.Issues[i].Fields.Issuetype.Name)
		f.SetCellValue(sheetName, "C"+istr, data.Issues[i].Fields.Priority.Name)
		f.SetCellValue(sheetName, "D"+istr, data.Issues[i].Fields.Severity.Value)
		f.SetCellValue(sheetName, "E"+istr, data.Issues[i].Fields.Components[0].Name)
		f.SetCellValue(sheetName, "F"+istr, data.Issues[i].Fields.Summary)
		f.SetCellValue(sheetName, "G"+istr, data.Issues[i].Fields.Assignee.DisplayName)
		f.SetCellValue(sheetName, "H"+istr, data.Issues[i].Fields.Status.Name)
		f.SetCellValue(sheetName, "I"+istr, data.Issues[i].Fields.Created)
		f.SetCellValue(sheetName, "J"+istr, data.Issues[i].Fields.Updated)

	}

	return nil

}

func GetSpreadsheetName() string {
	return "Tickets " + time.Now().Format("2006-01-02") + ".xlsx"
}
