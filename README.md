```
---------------------------------------------
| ░██████╗░░██████╗░██████╗░██╗░░██╗░██████╗|
| ██╔════╝░██╔════╝░██╔══██╗██║░██╔╝██╔════╝|
| ██║░░██╗░██║░░██╗░██████╔╝█████═╝░╚█████╗░|
| ██║░░╚██╗██║░░╚██╗██╔══██╗██╔═██╗░░╚═══██╗|
| ╚██████╔╝╚██████╔╝██║░░██║██║░╚██╗██████╔╝|
| ░╚═════╝░░╚═════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░|
---------------------------------------------

ggrks: ggrks [-o <number>] [words ...]
    Use Google to search for information based on the provided keywords.

    If no keywords are given, this command will display the help message.

    Options:
    -o <number>   Open the specified search result in your default web browser.
                    The <number> corresponds to the rank of the search result
                    (e.g., 1 for the first result).

    Examples:
    ggrks programming Go
        Search Google for "programming Go".
    ggrks -o 1 Go tutorial
        Search Google for "Go tutorial" and open the first search result in the default web browser.
    ggrks -o 3 Go framework
        Search Google for "Go framework" and open the third search result in the default web browser.

    Exit Status:
    Returns 0 if the search was successfully executed, or a non-zero status if an error occurred.

    Additional Notes:
    - Ensure that your system has internet access to use this command.
    - The `-o` option requires a web browser to be properly configured on your system.
    - If <number> is not provided with `-o`, it defaults to the first search result (1).
```
