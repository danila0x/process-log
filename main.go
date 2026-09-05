package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Timestamp struct {
	LogDate time.Time `json:"timestamp"`
}

type ProcessLog struct {
	ID    int       `json:"id"`
	Date  Timestamp `json:"timestamp"`
	Value float64   `json:"value"`
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	lTime := t.LogDate.UTC().Format(time.RFC3339)
	result := fmt.Sprintf(`"%s"`, lTime)
	return []byte(result), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	normalized := strings.ReplaceAll(dateStr, "/", "-")
	layout := "2006-01-02-15:04:05-0700"
	parsedTime, err := time.Parse(layout, normalized)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}
	t.LogDate = parsedTime
	return nil
}

func main() {
	start := time.Date(2026, 9, 6, 10, 0, 0, 0, time.UTC)
	end := time.Date(2026, 9, 6, 11, 0, 0, 0, time.UTC)
	err := ProcessLogs("input.txt", "output.txt", start, end)
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessLogs(inputPath, outputPath string, start, end time.Time) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	for scanner.Scan() {
		data := []byte(scanner.Text())
		var processLog ProcessLog
		err = json.Unmarshal(data, &processLog)
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(processLog)
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))
		if (processLog.Date.LogDate.Before(end) || processLog.Date.LogDate.Equal(end)) && (processLog.Date.LogDate.After(start) || processLog.Date.LogDate.Equal(start)) {
			writer.WriteString(string(jsonData) + "\n")
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
