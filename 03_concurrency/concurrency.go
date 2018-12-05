// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	FileSum("input.txt", "output.txt")
	inputFile := OpenFile("input.txt")
	defer inputFile.Close()
	outputFile := OpenFile("output.txt")
	defer outputFile.Close()
	IOSum(inputFile, outputFile)
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	inputFile := OpenFile(input)
	defer inputFile.Close()
	outputFile := OpenFile(output)
	defer outputFile.Close()

	sum := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Failed to convert string to int", err)
			os.Exit(1)
		}
		sum += i
	}

	_, err := outputFile.WriteString(fmt.Sprintf("%d", sum))
	if err != nil {
		fmt.Println("Failed to write to output file", err)
	}
}

func OpenFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Failed to open file", filename)
		os.Exit(1)
	}
	return file
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	buffer := make([]byte, 2)
	sum := 0

	for {
		bytesRead, err := input.Read(buffer)
		if bytesRead > 0 {
			char := buffer[0]
			i, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error converting input to int", err)
				os.Exit(1)
			}
			sum += i
		}

		if err == io.EOF {
			// success - write output to file
			outputBuffer := []byte(fmt.Sprintf("%d", sum))
			_, err := output.Write(outputBuffer)
			if err != nil {
				fmt.Println("Error writing file", err)
				os.Exit(1)
			}
			break
		}
		if err != nil {
			fmt.Println("Error reading file", err)
			os.Exit(1)
		}
	}
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	// TODO
	directory map[int]string
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	// TODO
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	// TODO
	return ""
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	// TODO
}
