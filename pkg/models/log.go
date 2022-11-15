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
	Ip        string
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

func (l *Log) WriteJsonFile(fileName string, nginxLog bool) error {
	var logData []jsonFormat
	var contentLines = strings.Split(l.Content, "\n")

	for _, v := range contentLines {
		if v != "" {
			logData = append(logData, l.parseTextToJson(v, nginxLog))
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

func (l *Log) parseTextToJson(line string, nginxLog bool) jsonFormat {
	lineSplit := strings.Split(line, " ")

	if nginxLog {
		return jsonFormat{
			Ip:        lineSplit[0],
			TimeStamp: strings.Join(lineSplit[3:5], " "),
			Message:   strings.Join(lineSplit[5:], " "),
		}
	} else {
		return jsonFormat{
			TimeStamp: strings.Join(lineSplit[0:3], " "),
			User:      strings.Join(lineSplit[3:4], " "),
			Process:   strings.Join(lineSplit[4:5], " "),
			Message:   strings.Join(lineSplit[5:], " "),
		}
	}
}
