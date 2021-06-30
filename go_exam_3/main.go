package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	poArr := []Photo{}

	employeeID := flag.String("employee_id", "", "employee_id")
	idRange := flag.Int("id_range", 0, "id_range")

	if *idRange > 0 && *idRange < 500 {
		// return error
	}

	if *employeeID == "" {
		// return error
	}

	flag.Parse()
	var wg sync.WaitGroup

	for i := 0; i <= *idRange; i++ {
		wg.Add(1)
		// var po Photo
		go func(id int) {
			photo := GetValue(id, &wg)
			photo.RoutineId = id
			poArr = append(poArr, *photo)
			wg.Done()
		}(i)

	}
	wg.Wait()
	writeXlsxPhoto(*employeeID, poArr)
}

func GetValue(id int, wg *sync.WaitGroup) (p *Photo) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	url := fmt.Sprintf("%s%d", "https://jsonplaceholder.typicode.com/photos/", id)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error json unmarshall : ", err)
		return nil
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("error call ", url, " ", err)
		return nil
	}

	defer resp.Body.Close()

	var respData Photo
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		fmt.Println("error photos json decoder2")
		return nil
	}
	return &respData
}

func writeXlsxPhoto(fileName string, photoArr []Photo) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("gobasic_exam3")
	if err != nil {
		fmt.Println("error cannot add sheet name")
		return
	}

	thRow := sheet.AddRow()
	thRow.AddCell().SetValue("EmployeeId")
	thRow.AddCell().SetValue("RoutineId")
	thRow.AddCell().SetValue("ID")
	thRow.AddCell().SetValue("AlbumId")
	thRow.AddCell().SetValue("Title")
	thRow.AddCell().SetValue("Url")
	thRow.AddCell().SetValue("ThumbnailUrl")
	thRow.AddCell().SetValue("Error")

	for _, po := range photoArr {
		row = sheet.AddRow()
		row.AddCell().SetValue(fileName)
		row.AddCell().SetValue(po.RoutineId)
		row.AddCell().SetValue(po.AlbumId)
		row.AddCell().SetValue(po.Id)
		row.AddCell().SetValue(po.Title)
		row.AddCell().SetValue(po.Url)
		row.AddCell().SetValue(po.ThumbnailUrl)
	}

	fullFileName := fmt.Sprintf("%s%s", fileName, ".xlsx")
	err = file.Save(fullFileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
