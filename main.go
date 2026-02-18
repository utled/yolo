package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Options struct {
	Programs map[string]Program `json:"programs"`
	Configs map[string]Config `json:"configs"`
}

type Program struct {
	Description string `json:"description"`
	RunCommand string `json:"runcommand"`
}

type Config struct {
	Description string `json:"description"`
	FullPath string `json:"fullpath"`
}

func createNewOptionsFile(homedir string) error {
	someProgram := Program{Description: "program description", RunCommand: "'./path/to/program' or similar"}
	someConfig := Config{Description: "config file description", FullPath: "path/to/file/config/file"}

	thePrograms := make(map[string]Program)
	thePrograms["programDisplayName"] = someProgram

	theConfigs := make(map[string]Config)
	theConfigs["configDisplayName"] = someConfig

	theOptions := Options{Programs: thePrograms, Configs: theConfigs}
	jsonOptions, _ := json.MarshalIndent(theOptions, "", "  ")
	
	dirPath := filepath.Join(homedir, ".yolo/")
	if _, err := os.Stat(dirPath); err != nil {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create .yolo dir:\n%v", err)
		}
	}

	filepath := filepath.Join(dirPath, "optionFile.json")
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create new file:\n%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(jsonOptions))
	if err != nil {
		return fmt.Errorf("failed to write json to file\n%v", err)
	}
	file.Sync()

	return nil
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to read user home dir:\n", err)
	}
	filepath := filepath.Join(homeDir, "/.yolo/", "optionFile.json")
	_, err = os.Stat(filepath)
	if err != nil {
		err := createNewOptionsFile(homeDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	optionFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("failed to read option file:\n", err)
	}
	var theOptions Options
	json.Unmarshal(optionFile, &theOptions)

	var programs []string
	var configs []string
	
	for program := range theOptions.Programs {
		programs = append(programs, program)
	}
	for config := range theOptions.Configs {
		configs = append(configs, config)
	}

	fmt.Println(theOptions.Configs[configs[0]].Description)
	fmt.Println(theOptions.Programs[programs[0]].Description)
}
