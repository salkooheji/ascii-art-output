package main

import (
	functions "Ascii/functions"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//extracting the file name from the args
	var shapeFileContent string
	// check what's the defualt ?
	// if len(os.Args)==4 {
	// 	if len(os.Args) == 2 { // standard is the default
	// 		shapeFileContent = functions.ReadShapeFile("standard.txt")
	// 	} else
			fmt.Println(len(os.Args))
		 if len(os.Args) == 4 {
			if functions.ValidateArugments() {
				fmt.Print("line 22 passed")
				shapeFileContent = functions.ReadShapeFile(os.Args[2])
			} else {
				log.Fatal("Check the input entered the program expect 3 values (Value to be displayed text (string), Font (.txt), the file to show the result in (.txt))")
			}
		} else {
		log.Fatal("Invalid Number of arguments")
	}

	charactersMap := functions.MapFileContentToRunes(os.Args[1],shapeFileContent)
	textToDraw := os.Args[1]
	// This will help in handling \n each index in textToDrawSplit will be considered a line  
	textToDrawSplit  := strings.Split(textToDraw,`\n`)

	for i, line := range textToDrawSplit {
		if line != ""{
			generateShape := functions.GenerateShape(line,charactersMap)
			//fmt.Println(generateShape)
			 os.WriteFile(os.Args[3], []byte(generateShape), 0o644)

		}

		if(i<len(textToDrawSplit)-1){
			fmt.Println()
		}
	}

}
