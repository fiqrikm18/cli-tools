package models

import (
	"encoding/json"
	"os"
	"strings"
)

type Log struct {
	Content string
}

type jsonFormat struct {
	TimeStamp string
	User      string
	Process   string
	Message   string
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) WriteTextFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(l.Content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (l *Log) WriteJsonFile(fileName string) error {
	var logData []jsonFormat
	var contentLines = strings.Split(l.Content, "\n")

	for _, v := range contentLines {
		if v != "" {
			logData = append(logData, l.parseTextToJson(v))
		}
	}

	jsonData, err := json.MarshalIndent(logData, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (l *Log) parseTextToJson(line string) jsonFormat {
	lineSplit := strings.Split(line, " ")
	return jsonFormat{
		TimeStamp: strings.Join(lineSplit[0:3], " "),
		User:      strings.Join(lineSplit[3:4], " "),
		Process:   strings.Join(lineSplit[4:5], " "),
		Message:   strings.Join(lineSplit[5:], " "),
	}
}
