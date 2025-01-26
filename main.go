package main

func main() {
	mode, options := GetOption()
	cfg, err := LoadConfig()
	if err != nil {
		PrintErr(err)
		return
	}

	switch mode {
	case Help:
		err := HelpFunc()
		if err != nil {
			PrintErr(err)
			return
		}

	case SimpleSearch:
		err := SearchFunc(cfg, options)
		if err != nil {
			PrintErr(err)
		}

	}
}
