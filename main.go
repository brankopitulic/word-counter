package main

import (
	"os"
)

func main() {
	// open txt file
	// Open file and create scanner on top of it
	//file, err := os.Open("test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//scanner := bufio.NewScanner(file)
	//
	//// Default scanner is bufio.ScanLines. Lets use ScanWords.
	//// Could also use a custom function of SplitFunc type
	//scanner.Split(bufio.ScanWords)
	//
	//// Scan for next token.
	//success := scanner.Scan()
	//if success == false {
	//	// False on error or EOF. Check error
	//	err = scanner.Err()
	//	if err == nil {
	//		log.Println("Scan completed and reached EOF")
	//	} else {
	//		log.Fatal(err)
	//	}
	//}
	//
	//// Get data from scan with Bytes() or Text()
	//fmt.Println("First word found:", scanner.Text())

	// Open file and create a buffered reader on top
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		return
	}

	// Call scanner.Scan() again to find next token
	// read file
	// given result are words, characters and lines
}
