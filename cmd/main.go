package main


import (
   "fmt"
   "flag"
   "github.com/HimanshuBarak/ccwc-go/command"
)



func parseArgs()(map[string]bool, []string){
	countBytes := flag.Bool("c", false, "Returns the number of bytes in the provided file")
	countLines := flag.Bool("l", false, "Returns the number of lines in the provided file")
	countWords := flag.Bool("w", false, "Returns the number of words in the provided file")
	countChars := flag.Bool("m", false, "Returns the number of characters in the provided file")
	flag.Parse()
	args := map[string]bool{
		"c": *countBytes,
		"l": *countLines,
		"w": *countWords,
		"m": *countChars,
	}
	//for file name (non-flag arguments)
	files := flag.Args()   // this returns an array
	
	return args, files
}

func main(){

	args, files := parseArgs()
    if len(files)>1 {
		fmt.Println("The tool only accepts at most one file name.")
		return 
	}
	command.Run(args, files) 
	
}



