package text_editor_tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ModifyingTheText(str_to_edit string) string {

	str_to_edit = Remove_extra_spaces(str_to_edit)

	str_to_edit = Rep_A_with_An(str_to_edit)

	str_to_edit = Remove_extra_spaces(str_to_edit)

	str_to_edit = Punct2(str_to_edit)

	///////////////////////////////////////

	str_to_edit = RepHex(str_to_edit)

	str_to_edit = RepBin(str_to_edit)

	///////////////////////////////////////

	str_to_edit = ToUp(str_to_edit)

	str_to_edit = ToLow(str_to_edit)

	str_to_edit = ToCap(str_to_edit)

	///////////////////////////////////////

	str_to_edit = New_up_num(str_to_edit)

	str_to_edit = New_low_num(str_to_edit)

	str_to_edit = New_cap_num(str_to_edit)

	///////////////////////////////////////

	str_to_edit = Remove_extra_spaces(str_to_edit)

	str_to_edit = Punct2(str_to_edit)

	str_to_edit = Remove_extra_spaces(str_to_edit)

	str_to_edit = SingleQuotation(str_to_edit)

	return str_to_edit
}

func RepHex(input string) string {
	re := regexp.MustCompile(`([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*)(\w+)(\s*)([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*\(+\s*((h)|(H))\s*((e)|(E))\s*((x)|(X))\s*\))+`)
	the_str := re.ReplaceAllStringFunc(input, func(s string) string {

		re1 := regexp.MustCompile(`\s*\(+\s*((h)|(H))\s*((e)|(E))\s*((x)|(X))\s*\)+`)
		the_hex := re1.ReplaceAllString(s, "")

		re2 := regexp.MustCompile(`((0x)*)((\w+)+)`)

		edit := re2.ReplaceAllStringFunc(the_hex, func(match string) string {
			re3 := regexp.MustCompile(`((0x)*)`)
			okok := re3.ReplaceAllString(match, "")

			dec, err := strconv.ParseInt(okok, 16, 64)
			TheNmber := strconv.FormatInt(dec, 10)
			if err != nil {
				fmt.Println("Enter a valed hex number syntax\nThe error is: ", err)
				TheNmber = the_hex
			}
			return TheNmber
		})

		return edit
	})

	re2 := regexp.MustCompile(`\s*\(+\s*((h)|(H))\s*((e)|(E))\s*((x)|(X))\s*\)+`)
	Last_V := re2.ReplaceAllString(the_str, "")
	return Last_V
}

func RepBin(s string) string {

	re := regexp.MustCompile(`([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*)(\w+)(\s*)([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*\(+\s*((b)|(B))\s*((i)|(I))\s*((n)|(N))\s*\)+)`)
	str_ok := re.ReplaceAllStringFunc(s, func(m string) string {

		re1 := regexp.MustCompile(`\s*\(+\s*((b)|(B))\s*((i)|(I))\s*((n)|(N))\s*\)+`)
		TheBin := re1.ReplaceAllString(m, "")

		re3 := regexp.MustCompile(`(\w+)`)

		edit := re3.ReplaceAllStringFunc(TheBin, func(match string) string {
			num, err_bin := strconv.ParseInt(match, 2, 64)
			TheNmber := strconv.FormatInt(num, 10)
			if err_bin != nil {
				fmt.Println("Enter a valed binary number syntax\nThe error is: ", err_bin)
				TheNmber = match
			}
			return TheNmber
		})
		return edit
	})

	re2 := regexp.MustCompile(`\s*\(+\s*((b)|(B))\s*((i)|(I))\s*((n)|(N))\s*\)+`)
	Last_V := re2.ReplaceAllString(str_ok, "")
	return Last_V
}

func ToUp(str string) string {
	re := regexp.MustCompile(`([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*)(\w+)(\s*)([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*\(+\s*((u)|(U))\s*((p)|(P))\s*\)+)`)
	modfied := re.ReplaceAllStringFunc(str, func(match string) string {

		re1 := regexp.MustCompile(`\s*\(+\s*((u)|(U))\s*((p)|(P))\s*\)+`)
		THENUM := re1.ReplaceAllString(match, "")

		return strings.ToUpper(THENUM)
	})
	re2 := regexp.MustCompile(`\s*\(+((u)|(U))\s*((p)|(P))\s*\)+`) //The (+) and (*) in the first \s
	Last_V := re2.ReplaceAllString(modfied, "")
	return Last_V
}

func ToLow(str string) string {
	re := regexp.MustCompile(`([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*)(\w+)(\s*)([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s*\(+\s*((l)|(L))\s*((o)|(O))\s*((w)|(W))\s*\)+)`)
	modfied := re.ReplaceAllStringFunc(str, func(match string) string {

		re1 := regexp.MustCompile(`\s*\(+\s*((l)|(L))\s*((o)|(O))\s*((w)|(W))\s*\)+`)
		TheWord := re1.ReplaceAllString(match, "")

		return strings.ToLower(TheWord)
	})
	re2 := regexp.MustCompile(`\s*\(+((l)|(L))\s*((o)|(O))\s*((w)|(W))\s*\)+`)
	Last_V := re2.ReplaceAllString(modfied, "")
	return Last_V
}

func ToCap(str string) string {
	re := regexp.MustCompile(`([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\w+)([\+\[\]{}=~@#$%^&*()'",./\?!\-_ ]*)(\s+\(+\s*((c)|(C))\s*((a)|(A))\s*((p)|(P))\s*\)+)`)
	modified := re.ReplaceAllStringFunc(str, func(match string) string {

		re1 := regexp.MustCompile(`\s*\(+\s*((c)|(C))\s*((a)|(A))\s*((p)|(P))\s*\)+`)
		The_Word := re1.ReplaceAllString(match, "")

		re2 := regexp.MustCompile(`(\w+)`)
		edit := re2.ReplaceAllStringFunc(The_Word, func(s string) string {
			return strings.ToUpper(string(s[0])) + s[1:]
		})

		return edit
	})

	re2 := regexp.MustCompile(`\s*\(+((c)|(C))\s*((a)|(A))\s*((p)|(P))\s*\)+`)
	Last_V := re2.ReplaceAllString(modified, "")
	return Last_V
}

func New_up_num(str string) string {
	re := regexp.MustCompile(`\s*\(+\)*((u)|(U))+((p)|(P))+,+([\s\w\-]*)\(*\)*`)
	re_num := regexp.MustCompile(`\s*\-*\w*\s*\)+`)
	only_num := regexp.MustCompile(`\-*\d*`)

	S_str := strings.Split(str, " ")
	New := ""
	NewStr := ""

	for i, word := range S_str {
		//Find the (up,
		if re.MatchString(word) {
			Location := i

			//Look if the number is like (up,num)
			if re_num.MatchString(word) {
				The_Number_Of_Word := only_num.FindString(word)

				N_up, err_up_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_up_num != nil {
							fmt.Println("Enter a valid number syntax for the (up, num) function\nThe error is: ", err_up_num)
							break
						}
					}
				}

				if N_up < 0 {
					fmt.Println("Enter a valid positive number syntax for the (up, num) function")

				}

				Length := len(S_str) - 1

				if (N_up) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_up); i-- {
						S_str[i] = strings.ToUpper(S_str[i])
					}
				} else if (N_up) > ((Length) - (Length - Location)) {
					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (up, num) function")
				}

				//Look for the number like (up, num)
			} else {
				n := S_str[i+1]
				The_Number_Of_Word := only_num.FindString(n)

				N_up, err_up_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_up_num != nil {
							fmt.Println("Enter a valid number syntax for the (up, num) function\nThe error is: ", err_up_num, i)
							break
						}
					}
				}

				if N_up < 0 {
					fmt.Println("Enter a valid positive number syntax for the (up, num) function")

				}

				Length := len(S_str) - 1

				if (N_up) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_up); i-- {
						S_str[i] = strings.ToUpper(S_str[i])
					}
				} else if (N_up) > ((Length) - (Length - Location)) {

					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (up, num) function")
				}

			}

		}

		New = strings.Join(S_str, " ")
		re_remove_all := regexp.MustCompile(`\s*\(+\)*\s*((u)|(U))\s*((p)|(P))\s*,+([\s\w\-]*)\(*\)\(*`)
		NewStr = re_remove_all.ReplaceAllString(New, "")
	}
	return NewStr
}

func New_low_num(str string) string {
	re := regexp.MustCompile(`\s*\(+\)*((l)|(L))+\s*((o)|(O))+\s*((w)|(W))+\s*,+([\s\w\-]*)\(*\)*`)
	re_num := regexp.MustCompile(`\s*\-*\w*\s*\)+`)
	only_num := regexp.MustCompile(`\-*\d*`)

	S_str := strings.Split(str, " ")
	New := ""
	NewStr := ""

	for i, word := range S_str {
		//Find the (up,
		if re.MatchString(word) {
			Location := i

			//Look if the number is like (up,num)
			if re_num.MatchString(word) {
				The_Number_Of_Word := only_num.FindString(word)

				N_low, err_low_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_low_num != nil {
							fmt.Println("Enter a valid number syntax for the (low, num) function\nThe error is: ", err_low_num)
							break
						}
					}
				}

				if N_low < 0 {
					fmt.Println("Enter a valid positive number syntax for the (low, num) function")

				}

				Length := len(S_str) - 1

				if (N_low) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_low); i-- {
						S_str[i] = strings.ToLower(S_str[i])
					}
				} else if (N_low) > ((Length) - (Length - Location)) {

					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (low, num) function")

				}

				//Look for the number like (up, num)
			} else {
				n := S_str[i+1]
				The_Number_Of_Word := only_num.FindString(n)

				N_low, err_up_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_up_num != nil {
							fmt.Println("Enter a valid number syntax for the (low, num) function\nThe error is: ", err_up_num, i)
							break
						}
					}
				}

				if N_low < 0 {
					fmt.Println("Enter a valid positive number syntax for the (low, num) function")

				}

				Length := len(S_str) - 1

				if (N_low) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_low); i-- {
						S_str[i] = strings.ToLower(S_str[i])
					}
				} else if (N_low) > ((Length) - (Length - Location)) {

					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (low, num) function")

				}

			}

		}

		New = strings.Join(S_str, " ")
		re_remove_all := regexp.MustCompile(`\s*\(+\)*\s*((l)|(L))+\s*((o)|(O))+\s*((w)|(W))+\s*,+([\s\w\-]*)\(*\)`)
		NewStr = re_remove_all.ReplaceAllString(New, "")
	}
	return NewStr

}

func New_cap_num(str string) string {
	re := regexp.MustCompile(`\s*\(+\)*((c)|(C))+\s*((a)|(A))+\s*((p)|(P))+\s*,+([\s\w\-]*)\(*\)*`)
	re_num := regexp.MustCompile(`\s*\-*\s*\w*\)+`)
	only_num := regexp.MustCompile(`-*\d*`)

	S_str := strings.Split(str, " ")
	New := ""
	NewStr := ""

	for i, word := range S_str {
		//Find the (up,
		if re.MatchString(word) {
			Location := i

			//Look if the number is like (up,num)
			if re_num.MatchString(word) {
				The_Number_Of_Word := only_num.FindString(word)

				N_cap, err_cap_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_cap_num != nil {
							fmt.Println("Enter a valid number syntax for the (cap, num) function\nThe error is: ", err_cap_num)
							break
						}
					}
				}

				if N_cap < 0 {
					fmt.Println("Enter a valid positive number syntax for the (cap, num) function")

				}

				Length := len(S_str) - 1
				a_word := ""
				S_OfAword := ""
				//Do the edting work
				if (N_cap) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_cap); i-- {

						a_word = S_str[i]
						a_word = string(a_word)
						S_OfAword = strings.ToUpper(string(a_word[0])) + a_word[1:]
						S_str[i] = S_OfAword
					}

				} else if (N_cap) > ((Length) - (Length - Location)) {
					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (cap, num) function")

				}

				//Look for the number like (up, num)
			} else {
				n := S_str[i+1]
				The_Number_Of_Word := only_num.FindString(n)

				N_cap, err_up_num := strconv.Atoi(The_Number_Of_Word)

				// Handl the error
				for _, s := range S_str {
					if re.MatchString(s) {
						if err_up_num != nil {
							fmt.Println("Enter a valid number syntax for the (cap, num) function\nThe error is: ", err_up_num, i)
							break
						}
					}
				}

				if N_cap < 0 {
					fmt.Println("Enter a valid positive number syntax for the (cap, num) function")

				}

				Length := len(S_str) - 1
				a_word := ""
				S_OfAword := ""
				//Do the edting work
				if (N_cap) <= ((Length) - (Length - Location)) {
					for i := Location - 1; i >= (Location - N_cap); i-- {

						a_word = S_str[i]
						S_OfAword = strings.ToUpper(string(a_word[0])) + a_word[1:]
						S_str[i] = S_OfAword
					}

				} else if (N_cap) > ((Length) - (Length - Location)) {
					fmt.Println("You have entered a bigger number than expected.\nEnter a valid positive number syntax for the (cap, num) function")

				}

			}

		}

		New = strings.Join(S_str, " ")
		re_remove_all := regexp.MustCompile(`\s*\(+\s*\)*\s*((c)|(C))+\s*((a)|(A))+\s*((p)|(P))+\s*,+([\s\w\-]*)\(*\)`)
		NewStr = re_remove_all.ReplaceAllString(New, "")
	}
	return NewStr

}

func Punct2(str string) string {
	re := regexp.MustCompile(`(.)(\s*)([.,!?;:]+)(.*?)`)
	punctuationRegex := regexp.MustCompile(`([.,!?;:])`)

	addSpace := punctuationRegex.ReplaceAllString(str, "$1 ")

	MatchText := re.ReplaceAllStringFunc(addSpace, func(match string) string {
		re2 := regexp.MustCompile(`(\s*)`)
		FinalStep := re2.ReplaceAllString(match, "")
		return FinalStep
	})
	return MatchText
}

func Rep_A_with_An(str string) string {
	vowels := []string{"a", "e", "h", "i", "o", "u", "A", "E", "H", "I", "O", "U"}

	not_vowels := []string{"B", "C", "D", "F", "G", "J", "K", "L", "M",
		"N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z",
		"b", "c", "d", "f", "g", "j", "k", "l", "m", "n",
		"p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	words := strings.Split(str, " ")
	for i, word := range words {
		if i > 0 {
			for _, vowel := range vowels {
				if strings.HasPrefix(word, vowel) && (words[i-1] == "a") {
					words[i-1] = "an"
				}

				if strings.HasPrefix(word, vowel) && (words[i-1] == "A") {
					words[i-1] = "An"
				}
			}
		}
	}

	for j, word_an := range words {
		if j > 0 {
			for _, NotVowel := range not_vowels {
				if (strings.HasPrefix(word_an, NotVowel)) && (words[j-1] == "an") {
					words[j-1] = "a"
				}

				if (strings.HasPrefix(word_an, NotVowel)) && (words[j-1] == "An") {
					words[j-1] = "A"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func Remove_extra_spaces(str string) string {
	re_func_word := regexp.MustCompile(`((\s*\(+\)*((u)|(U))+((p)|(P))+,+([\s\w\-]*)\(*\)*\s*)|(\s*\(+\)*((l)|(L))+\s*((o)|(O))+\s*((w)|(W))+\s*,+([\s\w\-]*)\(*\)*\s*)|(\s*\(+\)*((c)|(C))+\s*((a)|(A))+\s*((p)|(P))+\s*,+([\s\w\-]*)\(*\)*\s*))`)
	re_whitespace := regexp.MustCompile(`[^\S\n]+`)

	ok := re_func_word.ReplaceAllStringFunc(str, func(match string) string {
		fine := re_whitespace.ReplaceAllString(match, "")
		return fine
	})

	e := re_func_word.ReplaceAllString(ok, " $1 ")

	Remove_S := re_whitespace.ReplaceAllString(e, " ")
	Remove_S = strings.TrimSpace(Remove_S)

	return Remove_S
}

func SingleQuotation(str string) string {
	re := regexp.MustCompile(`'(\s*)(.*?)(\s*)'`)
	Replace_TheQ := re.ReplaceAllString(str, "'$2'")
	return Replace_TheQ
}
