// NOTES:
// run `go build`
// run `./cropTracker`
// it will print out the entered cmd default 'add_crop'
// then promt you to enter the name of the crop followed by the amount
// whatever is entered is then added to a csv as crop_record,name amount ( as entered )
// adds each value to data_store.csv as a running list
// run `./cropTracker` -cmd=subtract_amt or -cmd=add_amt
// you will get the cmd entered

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const exitCmd = "exit"

func main() {
	// A string flag.
	cmd := flag.String("cmd", "add_crop", "command to be executed")
	flag.Parse()
	executeCmd(*cmd)
}

func executeCmd(cmd string) {
	fmt.Println("cmd entered:", cmd)
	if cmd == "add_crop" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("name of the crop followed by the amount:")
		for scanner.Scan() {
			enteredInfo := scanner.Text()
			if enteredInfo == exitCmd {
				os.Exit(0)
			}
			fmt.Printf("You Entered - %q", enteredInfo)
			fmt.Println("") // Println will add back the final '\n'
			var currentData = [][]string{{"crop_record", enteredInfo}}
			writeData(currentData)
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}
}

func writeData(data [][]string) {
	file, err := os.OpenFile("data_store.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	checkError("Cannot create file", err)
	writer := csv.NewWriter(file)
	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
