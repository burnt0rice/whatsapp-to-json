package main

import (
	"encoding/json"
	"flag"
	"github.com/theckman/yacspin"
	"io/ioutil"
	"regexp"
	"time"
)

type ChatMessage struct {
	Date    string `json:"date"`
	Time    string `json:"time"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type ChatFile struct {
	Data []ChatMessage `json:"data"`
}

func main() {
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[14],
		SuffixAutoColon:   true,
		StopCharacter:     "✓",
		StopFailCharacter: "x",
		StopColors:        []string{"fgGreen"},
		StopFailColors:    []string{"fgRed"},
	}

	spinnerFile, err := yacspin.New(cfg)

	spinnerFile.Start()

	var path string

	flag.StringVar(&path, "p", "", "Specify the chat file path.")

	flag.Parse()

	if path == "" {
		spinnerFile.StopFailMessage(" missing chat path")
		spinnerFile.StopFail()
	}

	spinnerFile.Suffix(" reading txt file from path " + path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		spinnerFile.StopFail()
		return
	}

	fileContent := string(file)

	regexLine := regexp.MustCompile(`.+\b`)
	chatLines := regexLine.FindAll([]byte(fileContent), -1)

	regexMessage := regexp.MustCompile(`^\[(?P<Date>\d+.\d+.\d+)\,\s(?P<Time>\d+:\d+:\d+)\]\s(?P<Sender>.+):\s(?P<Message>.+)\b$`)
	var chat []ChatMessage
	var chatJSON ChatFile

	spinnerFile.Stop()
	spinnerBuild, err := yacspin.New(cfg)
	spinnerBuild.Suffix(" building json data")
	spinnerBuild.Start()

	for _, lines := range chatLines {
		line := string(lines)
		groupNames := regexMessage.SubexpNames()

		var chatMessage ChatMessage

		for _, match := range regexMessage.FindAllStringSubmatch(line, -1) {
			for groupIdx, group := range match {
				name := groupNames[groupIdx]

				if name == "" {
					name = "*"
				}

				switch name {
				case "Date":
					chatMessage.Date = group
				case "Time":
					chatMessage.Time = group
				case "Sender":
					chatMessage.Sender = group
				case "Message":
					chatMessage.Message = group
				}
			}
		}

		if chatMessage.Message != "" {
			chat = append(chat, chatMessage)
		}
	}

	chatJSON.Data = chat

	spinnerBuild.Stop()
	spinnerSave, err := yacspin.New(cfg)
	spinnerSave.Suffix(" saving json file")
	spinnerSave.Start()
	jsonFile, err := json.MarshalIndent(chatJSON, "", " ")

	if err != nil {
		spinnerSave.StopFail()
		return
	}

	err = ioutil.WriteFile("chat.json", jsonFile, 0644)

	if err != nil {
		spinnerSave.StopFail()
		return
	}

	spinnerSave.Stop()
}
