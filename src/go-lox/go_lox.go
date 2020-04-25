// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"go/scanner"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// var hadError bool

// func runFile(path string) {
// 	bytes, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		fmt.Printf("Error reading file %s", path)
// 		os.Exit(1)
// 	}
// 	run(string(bytes))
// 	if hadError {
// 		os.Exit(65)
// 	}
// }

// func runPrompt() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		print("> ")
// 		text, _ := reader.ReadString('\n')
// 		run(text)
// 		hadError = false
// 	}
// }

// func run(source string) {
// 	var s scanner.Scanner
// 	s.Init(strings.NewReader(source))

// 	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
// 		println(tok)
// 	}
// }

// func error(line int, message string) {
// 	report(line, "", message)
// }

// func report(line int, where string, message string) {
// 	fmt.Printf("[line %d] Error %s: %s", line, where, message)
// 	hadError = true
// }
