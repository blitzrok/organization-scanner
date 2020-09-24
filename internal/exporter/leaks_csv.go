package exporter

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	gitleaks "github.com/zricethezav/gitleaks/src"
	"os"
)

func LeaksToCSV(leaks []gitleaks.Leak, filename string) {
	csvFile, err := os.Create(filename)
	if err != nil {
		logrus.Error("Error creating CSV file", err)
	}
	writer := csv.NewWriter(csvFile)
	writer.Write(writeHeaders())
	for _, leak := range leaks {
		writer.Write(writeRow(leak))
	}
	writer.Flush()
}

func writeHeaders() []string {
	var headers []string
	headers = append(headers, "Repository")
	headers = append(headers, "Message")
	headers = append(headers, "Offender")
	headers = append(headers, "Author")
	headers = append(headers, "Type")
	headers = append(headers, "Commit")
	headers = append(headers, "Email")
	headers = append(headers, "File")
	headers = append(headers, "Line")
	headers = append(headers, "Date")
	return headers
}

func writeRow(leak gitleaks.Leak) []string {
	var row []string
	row = append(row, leak.Repo)
	row = append(row, leak.Message)
	row = append(row, leak.Offender)
	row = append(row, leak.Author)
	row = append(row, leak.Type)
	row = append(row, leak.Commit)
	row = append(row, leak.Email)
	row = append(row, leak.File)
	row = append(row, leak.Line)
	row = append(row, leak.Date.String())
	return row
}
