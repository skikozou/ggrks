package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Config struct {
	BaseDir         string
	GoogleSearchAPI string
	SearchEngineID  string
	MaxSearchResult int
}

type Mode int

const (
	Help Mode = iota
	SimpleSearch
	//CustomSearch
)

func GetArgs() []string {
	if len(os.Args) == 1 {
		return nil
	}
	return os.Args[1:]
}

func GetOption() (Mode, []string) {
	options := GetArgs()
	if options == nil {
		return Help, options
	} else {
		return SimpleSearch, options
	}
}

func PrintErr(err error) {
	fmt.Println(color.RedString(err.Error()))
}

func InitConfig() (*Config, error) {
	fmt.Println("Initial settings")
	var config Config
	var input string

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	config.BaseDir = usr.HomeDir + "\\"

	for {
		fmt.Printf("Custom Search JSON API key\n>")
		fmt.Scan(&input)
		if !(len(input) >= 39 && len(input) <= 40) {
			PrintErr(fmt.Errorf("Invalid API Key"))
			continue
		} else {
			config.GoogleSearchAPI = input
			break
		}
	}

	for {
		fmt.Printf("Google Custom Search Engine ID\n>")
		fmt.Scan(&input)
		if !(len(input) == 17) {
			PrintErr(fmt.Errorf("Invalid API Key"))
			continue
		} else {
			config.SearchEngineID = input
			break
		}
	}

	for {
		fmt.Printf("Max search result\n>")
		fmt.Scan(&input)

		val, err := strconv.Atoi(input)
		if err != nil {
			PrintErr(fmt.Errorf("invaid value"))
			continue
		} else if !(val <= 10 && val >= 1) {
			PrintErr(fmt.Errorf("Out of range (must be between 1 and 10)"))
			continue
		} else {
			config.MaxSearchResult = val
			break
		}
	}

	var cfgIni string
	cfgIni += ""

	return &config, nil
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	rawData, err := os.ReadFile(config.BaseDir + "config.ini")

	if err != nil {
		config, err = InitConfig()
		if err != nil {
			return nil, err
		}
	} else {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		config.BaseDir = usr.HomeDir + "\\"

		for i, v := range strings.Split(string(rawData), "\r\n") {
			//blank line check
			if v == "" {
				continue
			}

			key := strings.Split(v, " = ")[0]
			if len(key) == len(v) {
				return nil, fmt.Errorf("Unsupported format: Line %d\n%s", i+1, config.BaseDir+"config.ini")
			}
			value := strings.Split(v, " = ")[1]
			if value == "" {
				return nil, fmt.Errorf("invaid config value: %s\n%s", key, config.BaseDir+"config.ini")
			}

			switch key {
			case "GoogleSearchAPI":
				config.GoogleSearchAPI = value

			case "SearchEngineID":
				config.SearchEngineID = value

			case "MaxSearchResult":
				config.MaxSearchResult, err = strconv.Atoi(value)
				if err != nil {
					return nil, fmt.Errorf("invaid config value: %s\n%s", key, config.BaseDir+"config.ini")
				}
				if config.MaxSearchResult > 10 || config.MaxSearchResult < 0 {
					return nil, fmt.Errorf("MaxSearchResult is out of range (must be between 1 and 10)\n%s", config.BaseDir+"config.ini")
				}

			default:
				//none
			}
		}
	}

	return config, nil
}
