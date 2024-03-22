package main

import (
	text_editor_tools "abdyusuf/GoReloded/pkg"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	// Taking the arguments
	FileArgs := os.Args

	// Handling the error of the length of the entered arguments
	if len(FileArgs) < 3 {
		log.Fatal("(NO ENOUGH ARGUMENTS)\nTry to use (go run . sample.txt result.txt)")
		return
	} else if len(FileArgs) > 3 {
		log.Fatal("(TOO MANY ARGUMENT)\nTry to use (go run . sample.txt result.txt)")
		return
	}

	// Saving the files names
	input_file := FileArgs[1]
	output_file := FileArgs[2]

	// Cheging if the file names are corrct or not
	if (!(strings.HasSuffix(input_file, ".txt"))) || (!(strings.HasSuffix(output_file, ".txt"))) {
		log.Fatal("You have enter the wrong file names \nYou must enter the file name to be edited and the file name to be printed.\n(THE FILES NAME HAS TO END WITH (.txt)).")
		return
	}

	// Open the sample file
	S_File, err1 := os.OpenFile(input_file, os.O_RDONLY, 0644)
	if err1 != nil {
		log.Fatal("Error to open the sample.txt file\nThe error is: ", err1)
	}
	defer S_File.Close()

	//Take the contint of the input file
	scanner := bufio.NewScanner(S_File)
	str_of_sample_file := ""

	for scanner.Scan() {
		line := scanner.Text()
		str_of_sample_file = str_of_sample_file + line + "\n"
	}

	// handel the error if the sample file is empty.
	if str_of_sample_file == "" {
		log.Fatal("The sample file is empty.\nEnsert text to sample.txt to edited")
		return
	}

	// handel the error if the sample file is has noly English or not.
	for _, ascii := range str_of_sample_file {
		if ((ascii < 32) || (ascii > 126)) && (ascii != 10) {
			log.Fatal("USE ONLY ENGLISH")
			return
		}
	}

	// Edit the text
	EditedText := text_editor_tools.ModifyingTheText(str_of_sample_file)

	// Open the output file
	R_File, err3 := os.OpenFile(output_file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err3 != nil {
		log.Fatal("Error to open the result.txt file\nThe error is: ", err3)
		return
	}
	defer R_File.Close()

	// Write the final edited text to the result file.
	_, err4 := R_File.WriteString(EditedText)
	if err4 != nil {
		log.Fatal("Error to write to the result.txt file\nThe error is: ", err4)
		return
	}

}
