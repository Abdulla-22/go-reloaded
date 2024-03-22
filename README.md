# Text Completion/Editing/Auto-Correction Tool

This is the README file for the text completion/editing/auto-correction tool project. The objective of this project is to create a simple tool in Go that can modify text based on specific rules and requirements. The tool will receive input from a file, perform the necessary modifications, and output the modified text to another file.

## Project Objectives

- Utilize existing functions from your old repository.
- Write the project code in Go, following good coding practices.
- Include unit tests by creating test files.
- Implement the following modifications to the input text:

### Text Modifications

1. Replace every instance of `(hex)` with the decimal version of the preceding word (assuming the preceding word is always a hexadecimal number). For example:
   - Input: "1E (hex) files were added"
   - Output: "30 files were added"

2. Replace every instance of `(bin)` with the decimal version of the preceding word (assuming the preceding word is always a binary number). For example:
   - Input: "It has been 10 (bin) years"
   - Output: "It has been 2 years"

3. Convert every instance of `(up)` to the uppercase version of the preceding word. For example:
   - Input: "Ready, set, go (up) !"
   - Output: "Ready, set, GO !"

4. Convert every instance of `(low)` to the lowercase version of the preceding word. For example:
   - Input: "I should stop SHOUTING (low)"
   - Output: "I should stop shouting"

5. Convert every instance of `(cap)` to the capitalized version of the preceding word. For example:
   - Input: "Welcome to the Brooklyn bridge (cap)"
   - Output: "Welcome to the Brooklyn Bridge"

6. If `(low)`, `(up)`, or `(cap)` is followed by a number, apply the specified modification (lowercase, uppercase, or capitalized) to the specified number of words. For example:
   - Input: "This is so exciting (up, 2)"
   - Output: "This is SO EXCITING"

7. Ensure that punctuation marks `.`, `,`, `!`, `?`, `:`, and `;` are placed close to the preceding word and have a space before the next word. However, if there are groups of punctuation marks like `...` or `!?`, the text should be formatted accordingly. For example:
   - Input: "I was sitting over there ,and then BAMM !!"
   - Output: "I was sitting over there, and then BAMM!!"

8. Handle single quotation marks `'` by placing them to the right and left of the word in the middle without any spaces. If there are more than one word between the two `'` marks, place the marks next to the corresponding words. For example:
   - Input: "I am exactly how they describe me: ' awesome '"
   - Output: "I am exactly how they describe me: 'awesome'"

9. Replace every instance of `a` with `an` if the next word begins with a vowel (a, e, i, o, u) or an `h`. For example:
   - Input: "There it was. A amazing rock!"
   - Output: "There it was. An amazing rock!"

## Usage

To use the text completion/editing/auto-correction tool, follow the instructions below:

1. Create a text file containing the text that needs modification (input file).
2. Run the tool by executing the following command in the terminal:
   ```
   go run main.go <input_file> <output_file>
   ```
   Replace `<input_file>` with the name of the input file and `<output_file>` with the name of the file where the modified text should be saved.
3. Check the output file to view the modified text.

### Example

Let's assume we have a file named `sample.txt` with the following content:

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

To execute the tool and save the modified text to `result.txt`, run the following command:

```
go run main.go sample.txt result.txt
```

The modified text will be saved in the `result.txt` file:

```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it wasthe season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

Repeat the above steps with different input files and modifications to test the tool's functionality.

## Conclusion

Congratulations! You now have a text completion/editing/auto-correction tool implemented in Go. The tool can modify text based on specific rules and requirements, allowing you to automate various text editing tasks. Feel free to explore and expand upon this project further to enhance its capabilities. Happy coding!