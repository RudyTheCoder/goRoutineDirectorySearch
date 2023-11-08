# goRoutineDirectorySearch

goRoutineDirectorySearch is a Go program that demonstrates concurrent directory search using goroutines and channels. It searches for a specified target string in all `.txt` files within a specified directory.

## Table of Contents

- [Requirements](#requirements)
- [Usage](#usage)
- [Customization](#customization)
- [Example Output](#example-output)
- [Contributing](#contributing)

## Requirements

To run this program, you need to have [Go](https://golang.org/dl/) installed on your system.

## Usage

1. Clone or download this repository to your local machine.

2. Open a terminal or command prompt.

3. Navigate to the directory containing the `goRoutineDirectorySearch.go` file.

4. Run the program by executing the following command:

   ```bash
   go run goRoutineDirectorySearch.go

This command will start the directory search, and the program will print lines containing the target string along with the file name and line number.

1. After the program finishes, it will print "All goroutines have finished," and you'll see the search results on your terminal.


## Customization

* You can customize the target string by modifying the target variable in the main function.
* If you want to search in a different directory, change the folderName variable in the main function.
* Ensure that your text files have the .txt extension within the specified directory.


## Example Output

Here's an example of what the output might look like:

file1.txt (3): This is a sample line containing the target string.
file2.txt (10): Another line with the target string
...


## Contributing

If you'd like to contribute to this project, feel free to open issues or pull requests.



