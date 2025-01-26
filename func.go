package main

import (
	"fmt"
	"strconv"

	"github.com/skratchdot/open-golang/open"
)

func HelpFunc() error {
	fmt.Println("---------------------------------------------                                                          ")
	fmt.Println("| ░██████╗░░██████╗░██████╗░██╗░░██╗░██████╗|                                                          ")
	fmt.Println("| ██╔════╝░██╔════╝░██╔══██╗██║░██╔╝██╔════╝|                                                          ")
	fmt.Println("| ██║░░██╗░██║░░██╗░██████╔╝█████═╝░╚█████╗░|                                                          ")
	fmt.Println("| ██║░░╚██╗██║░░╚██╗██╔══██╗██╔═██╗░░╚═══██╗|                                                          ")
	fmt.Println("| ╚██████╔╝╚██████╔╝██║░░██║██║░╚██╗██████╔╝|                                                          ")
	fmt.Println("| ░╚═════╝░░╚═════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░|                                                          ")
	fmt.Println("---------------------------------------------                                                          ")
	fmt.Println("                                                                                                       ")
	fmt.Println("ggrks: ggrks [-o <number>] [words ...]                                                                 ")
	fmt.Println("    Use Google to search for information based on the provided keywords.                               ")
	fmt.Println("                                                                                                       ")
	fmt.Println("    If no keywords are given, this command will display the help message.                              ")
	fmt.Println("                                                                                                       ")
	fmt.Println("    Options:                                                                                           ")
	fmt.Println("    -o <number>   Open the specified search result in your default web browser.                        ")
	fmt.Println("                    The <number> corresponds to the rank of the search result                          ")
	fmt.Println("                    (e.g., 1 for the first result).                                                    ")
	fmt.Println("                                                                                                       ")
	fmt.Println("    Examples:                                                                                          ")
	fmt.Println("    ggrks programming Go                                                                               ")
	fmt.Println("        Search Google for \"programming Go\".                                                          ")
	fmt.Println("    ggrks -o 1 Go tutorial                                                                             ")
	fmt.Println("        Search Google for \"Go tutorial\" and open the first search result in the default web browser. ")
	fmt.Println("    ggrks -o 3 Go framework                                                                            ")
	fmt.Println("        Search Google for \"Go framework\" and open the third search result in the default web browser.")
	fmt.Println("                                                                                                       ")
	fmt.Println("    Exit Status:                                                                                       ")
	fmt.Println("    Returns 0 if the search was successfully executed, or a non-zero status if an error occurred.      ")
	fmt.Println("                                                                                                       ")
	fmt.Println("    Additional Notes:                                                                                  ")
	fmt.Println("    - Ensure that your system has internet access to use this command.                                 ")
	fmt.Println("    - The `-o` option requires a web browser to be properly configured on your system.                 ")
	fmt.Println("    - If <number> is not provided with `-o`, it defaults to the first search result (1).               ")

	return nil
}

func SearchFunc(config *Config, options []string) error {
	if options[0] == "-o" {
		if len(options) != 1 {
			num, err := strconv.Atoi(options[1])
			if err != nil {
				results, err := SearchOnGoogle(config, options[1:])
				if err != nil {
					return err
				}

				fmt.Printf("- %s\n%s\n", results.Items[0].Title, results.Items[0].Link)
				open.Run(results.Items[0].Link)
				return nil
			} else {
				results, err := SearchOnGoogle(config, options[2:])
				if err != nil {
					return err
				}

				fmt.Printf("- %s\n%s\n", results.Items[num].Title, results.Items[num].Link)
				open.Run(results.Items[num].Link)
				return nil
			}
		} else {
			return HelpFunc()
		}
	} else {
		results, err := SearchOnGoogle(config, options)
		if err != nil {
			return err
		}

		for i, v := range results.Items {
			if i >= config.MaxSearchResult {
				break
			}

			fmt.Printf("- %s\n%s", v.Title, v.Link)
			if i != len(results.Items)-1 {
				fmt.Printf("\n\n")
			}
		}

		return nil
	}
}
