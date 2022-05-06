package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"os"
	"report-generator/pkg/entity"
	"strconv"
	"strings"
	"time"
)

var dateLayout = "2006/01/02"
var createdAtLayout = "2006/01/02 12:34:56"

var (
	clubName        string
	roomNumber      string
	inputPath       string
	inputExcelPath  string
	outputMdPath    string
	outputExcelPath string
)

func main() {
	now := time.Now()
	flag.StringVar(&clubName, "c", "クラブ名", "club name")
	flag.StringVar(&roomNumber, "r", "0000", "room number")
	flag.StringVar(&inputPath, "i", "file.csv", "input csv")
	flag.StringVar(&inputExcelPath, "t", "月間活動報告書.csv", "input excel path")
	flag.StringVar(&outputMdPath, "om", now.Format("200601")+".md", "output markdown mF")
	flag.StringVar(&outputExcelPath, "oe", now.Format("200601")+".xlsx", "output excel mF")
	flag.Parse()

	cf, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf(errors.Wrap(err, "Failed read csv").Error())
	}

	r := csv.NewReader(cf)
	reports := make([]*entity.Report, 0)
	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(errors.Wrap(err, "Failed line read csv").Error())
		}
		if i == 0 {
			continue
		}

		// 行毎に取り出し Report 型に当てはめる
		date, _ := time.Parse(dateLayout, record[2])
		createdAt, _ := time.Parse(createdAtLayout, record[0])
		student1, _ := strconv.Atoi(record[4])
		student2, _ := strconv.Atoi(record[5])
		student3, _ := strconv.Atoi(record[6])
		student4, _ := strconv.Atoi(record[7])

		reports = append(reports, &entity.Report{
			Date:      date,
			Content:   record[3],
			Students:  []int{student1, student2, student3, student4},
			CreatedAt: createdAt,
		})
	}

	weekdayJa := strings.NewReplacer(
		"Sun", "日",
		"Mon", "月",
		"Tue", "火",
		"Wed", "水",
		"Thu", "木",
		"Fri", "金",
		"Sat", "土",
	)

	mF, _ := os.Create(outputMdPath)
	defer mF.Close()
	t, _ := excelize.OpenFile(inputExcelPath)
	defer t.Close()

	sheetName := "Sheet1"
	t.SetCellValue(sheetName, "D4", clubName)
	t.SetCellValue(sheetName, "P4", now.Format(dateLayout))

	initColLine := 8
	for i, report := range reports {
		students := report.Students[0] + report.Students[1] + report.Students[2] + report.Students[3]

		reportStr := fmt.Sprintf(""+
			"## %s 16:45 ~ 20:00\n\n"+
			"### 活動内容\n\n"+
			"%s\n\n"+
			"### 人数\n\n"+
			"計 %d 名\n\n"+
			"| 学年   | 人数  |\n|:-----|:----|\n| 1 年生 | %d 名 |\n| 2 年生 | %d 名 |\n| 3 年生 | %d 名 |\n| 4 年生 | %d 名 |\n\n",
			report.Date.Format("2006/01/02 (Mon)"),
			report.Content,
			students,
			report.Students[0],
			report.Students[1],
			report.Students[2],
			report.Students[3],
		)

		colNum := strconv.Itoa(initColLine + i)
		mF.WriteString(weekdayJa.Replace(reportStr))
		t.SetCellValue(sheetName, "A"+colNum, report.Date.Format("02"))
		t.SetCellValue(sheetName, "C"+colNum, report.Content)
		t.SetCellValue(sheetName, "K"+colNum, roomNumber)
		t.SetCellValue(sheetName, "N"+colNum, students)
	}

	t.SaveAs(outputExcelPath)
}
