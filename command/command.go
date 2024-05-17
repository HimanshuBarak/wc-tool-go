package command 

import (
   "fmt"
   "os"
   "bufio"
   "log"
   "unicode"
   "io"
)


type Result struct{
	totalBytes int64
	totalChars int64
	totalLines int64
	totalWords int64
}


func getFileDetails(args map[string]bool, reader *bufio.Reader, results *Result){
	results.totalBytes = 0
	results.totalChars = 0
	results.totalLines = 0
	results.totalWords = 0
	var prevRune rune
	for {
		runeRead ,runeSize ,err := reader.ReadRune()
      
		if err!=nil {
              if err == io.EOF{
				  break  // when the whole file has been read stop the while loop
			  }else{
				log.Fatal(err.Error())
			  }
		}
		if args["c"] {
			results.totalBytes += int64(runeSize)
		} else if args["l"]  && runeRead == '\n'{
			results.totalLines++;
		} else if args["w"] {
			if unicode.IsSpace(runeRead) && !unicode.IsSpace(prevRune) {
                   // here the logic being if prev word is word and current is whitespace so count a word basically
				   results.totalWords++;
			}
		} else if args["m"] {
			results.totalChars++;
		} else {
			results.totalBytes += int64(runeSize)
			results.totalChars++;
            if runeRead == '\n' {
				results.totalLines++;
			}
			if unicode.IsSpace(runeRead) && !unicode.IsSpace(prevRune) {
				// here the logic being if prev word is word and current is whitespace so count a word basically
				results.totalWords++;
		 	}
		}
		prevRune = runeRead
	}
}


func Run(args map[string]bool,file []string){
    
	results := Result{}
	//check if the user input a file name
    if len(file)==0 {
		// no file name was shared data needs to be readed from the terminal itself
		scanner := bufio.NewReader(os.Stdin)
		getFileDetails(args,scanner,&results)
	}else{
		f, err := os.Open(file[0])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		
		scanner := bufio.NewReader(f)
		getFileDetails(args,scanner,&results)
		defer f.Close()
	}
	display(args,&results)
	if len(file)!=0{
		fmt.Println(file[0])
	}
}


func display(args map[string]bool,results *Result){
	output := ""
		if !(args["c"] || args["l"] || args["w"] || args["m"]){
			output += fmt.Sprintf("%v\t", results.totalBytes)
			output += fmt.Sprintf("%v\t", results.totalLines)
			output += fmt.Sprintf("%v\t", results.totalWords)
			output += fmt.Sprintf("%v\t", results.totalChars)
		}else {
			if args["c"]{
				output += fmt.Sprintf("%v\t", results.totalBytes)
			}
			if args["l"]{
				output += fmt.Sprintf("%v\t", results.totalLines)
			}
			if args["w"]{
				output += fmt.Sprintf("%v\t", results.totalWords)
			}
			if args["m"]{
				output += fmt.Sprintf("%v\t", results.totalChars)	
			}
		}
	fmt.Printf(output)
}