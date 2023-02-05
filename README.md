## Obscene Vocabulary Checker

A simple console program written in Go that replaces certain words with `******`.

### Objectives

1. Read the name of the file that contains the taboo words from the user input.
2. Read the file and store the taboo words in a list.
3. Read sentences from the input. If the sentence contains taboo words, replace them with `****` and print the result. If the sentence does not contain any taboo words, print the sentence.
4. Repeat this process until the `exit` command is received.
5. When the program receives the exit command, print `Bye!` and exit the program.

### Getting Started
To run the project, you will need to have Go installed on your machine.
1. Clone the repository: git clone `git@github.com:rugpanov/go-obscene-vocabulary-checker.git`
2. Navigate to the project directory: `cd go-obscene-vocabulary-checker`
3. Generate a file with a list of taboo words
4. `Run the project: go run main.go`
5. Provide the name of the file containing taboo words
6. Provide the sentences to check and censor
7. To `exit` the program, type exit in the console

### Example

The `forbidden_words.txt` contents:
* disgusting
* unpleasant
* ugly
* bad

```text
> forbidden_words.txt
> It is a Bad and ugly sentence.
It is a *** and **** sentence.
> It is a Badge.
It is a Badge.
> exit
Bye!
```