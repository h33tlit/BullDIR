package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	fmt.Print(string(colorGreen), " ▄▄▄▄    █    ██  ██▓     ██▓    ▓█████▄  ██▓ ██▀███  \n▓█████▄  ██  ▓██▒▓██▒    ▓██▒    ▒██▀ ██▌▓██▒▓██ ▒ ██▒\n▒██▒ ▄██▓██  ▒██░▒██░    ▒██░    ░██   █▌▒██▒▓██ ░▄█ ▒\n▒██░█▀  ▓▓█  ░██░▒██░    ▒██░    ░▓█▄   ▌░██░▒██▀▀█▄  \n░▓█  ▀█▓▒▒█████▓ ░██████▒░██████▒░▒████▓ ░██░░██▓ ▒██▒\n░▒▓███▀▒░▒▓▒ ▒ ▒ ░ ▒░▓  ░░ ▒░▓  ░ ▒▒▓  ▒ ░▓  ░ ▒▓ ░▒▓░\n▒░▒   ░ ░░▒░ ░ ░ ░ ░ ▒  ░░ ░ ▒  ░ ░ ▒  ▒  ▒ ░  ░▒ ░ ▒░\n ░    ░  ░░░ ░ ░   ░ ░     ░ ░    ░ ░  ░  ▒ ░  ░░   ░ \n ░         ░         ░  ░    ░  ░   ░     ░     ░     \n      ░                           ░               \n")
	fmt.Print(string(colorRed), "Fast hidden directory/file scanner made with GO!\n\nAuthor: Jubaer alnazi, Version: 1.0, Website: https://jubaeralnazi.com\n\n\n")

	// Declaring input variable for target.
	var targetinput string

	// Taking input for target.
	fmt.Print(string(colorGreen), "Enter your target (example: https://jubaeralnazi.com) ==> ")
	fmt.Scan(&targetinput)

	// Declaring input variable for wordlist directory
	var dirinput string

	// Taking input for absolute path
	fmt.Print(string(colorGreen), "Enter your wordlist directory path ==> ")
	fmt.Scan(&dirinput)

	// Opening the file
	content, _ := os.Open(dirinput)

	// Reading the file line by line
	fileread := bufio.NewScanner(content)

	// Storing the lines in array
	var listofall []string
	for fileread.Scan() {
		var line string
		line = fileread.Text()
		fullurl := fmt.Sprintf("%s/%s", targetinput, line)
		listofall = append(listofall, fullurl)
	}

	content.Close()

	testresp, err := http.Get(targetinput)
	if err != nil {
		fmt.Println(string(colorRed), targetinput, "is down !!!")
		os.Exit(1)
	}
	testresp.Body.Close()

	for _, url1 := range listofall {

		// Call the function check
		scanUrl(url1)
	}

	fmt.Scanln()
	fmt.Println("Scan finished!")
}

// Function to scan each URL
func scanUrl(url1 string) {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorWhite := "\033[37m"

	resp, err := http.Get(url1)
	if err != nil {
		fmt.Println(string(colorRed), url1, "is down !!!")
		os.Exit(1)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println(string(colorWhite), url1, "==> Status:", string(colorGreen), resp.StatusCode)
			return
		} else {
			fmt.Println(string(colorWhite), url1, "==> Status:", string(colorRed), resp.StatusCode)
			return
		}

	}

}
