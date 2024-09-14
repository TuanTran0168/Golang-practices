package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strings"
	"unicode"
)

func main() {
	// randomGuess()
	// quadraticEquation()
	// validateCharacter()
	// sumPrimeNumber()
	matrix()
}

// Exercise 1
func randomGuess() {
	fmt.Println("============== RANDOM NUMBER GUESSING GAME ==============")
	var number uint
	var randomNumber uint
	min := 1
	max := 100

	randomNumber = uint(rand.IntN(max-min) + min)
	for {

		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scan(&number)
		fmt.Println("Your number is: ", number)

		if number < 1 && number > 100 {
			fmt.Println("Number out of range")
			continue
		}

		if number == randomNumber {
			fmt.Println("Correct")
			break
		} else if number < randomNumber {
			fmt.Println("Too low")
		} else if number > randomNumber {
			fmt.Println("Too high")
		}

	}
}

// Exercise 2
func quadraticEquation() {
	fmt.Println("============== QUADRATIC EQUATION SOLVER ==============")
	var n uint
	var array []float64
	var value float64
	var x int64
	var result float64

	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	for i := 0; i <= int(n); i++ {
		fmt.Print("Enter number: ")
		fmt.Scan(&value)
		array = append(array, float64(value))
	}

	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	for i := 0; i < len(array); i++ {
		if uint(i) < n {
			result = result + math.Pow(float64(x), float64(n-uint(i)))*array[n-uint(i)]
		}
		if uint(i) == n {
			result = result + array[n-uint(i)]
		}

	}

	fmt.Println("Total Result: ", result)

}

// Exercise 3
func validateCharacter() {
	fmt.Println("============== VALIDATE CHARACTER ==============")
	validateCharacter_1()
	validateCharacter_2()
	validateCharacter_3()
	validateCharacter_4()
	validateCharacter_5()
}

// Exercise 3 - 1
func validateCharacter_1() {
	fmt.Println("===== VALIDATE CHARACTER 1 =====")
	var character string

	fmt.Print("Enter a character: ")
	fmt.Scan(&character)

	index := strings.Index(character, "@")

	if index == -1 {
		fmt.Println("Invalid character")
		return
	}
	character = character[:index]

	fmt.Println(character)

}

// Exercise 3 - 2
func validateCharacter_2() {
	fmt.Println("===== VALIDATE CHARACTER 2 =====")

	var character *bufio.Reader = bufio.NewReader(os.Stdin)
	character.ReadString('\n')

	fmt.Print("Enter a character: ")
	statement, _ := character.ReadString('\n')
	fmt.Printf("You entered: %s", statement)

	var count int

	for _, element := range statement {
		// rune is an alias for int32, it holds a unicode value
		if unicode.IsUpper(rune(element)) {
			count++
		}
	}
	fmt.Println("Total Upper Case Character: ", count)
}

// Exercise 3 - 3
func validateCharacter_3() {
	fmt.Println("===== VALIDATE CHARACTER 3 =====")
	var newReplace string
	var pathFile = "files/lesson_1_read.txt"
	data := readFile(pathFile)

	fmt.Print("Enter a character to replace: ")
	fmt.Scan(&newReplace)

	data_replace := strings.ReplaceAll(data, "{file}", newReplace)
	fmt.Println(data_replace)

	var pathFile_write = "files/lesson_1_write.txt"

	// Write to file
	var character *bufio.Reader = bufio.NewReader(os.Stdin)
	character.ReadString('\n')
	fmt.Print("Enter a data input file: ")
	data_input, _ := character.ReadString('\n')
	createFile(pathFile_write, data_input)

}

func readFile(pathFile string) string {
	file, err := os.ReadFile(pathFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened", pathFile)
	return string(file)
}

func createFile(pathFile string, data_input string) {
	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("Created", pathFile)

	file.WriteString(data_input)
}

// Exercise 3 - 4
func validateCharacter_4() {
	data := readFile("files/lesson_1_read.txt")
	fmt.Println("===== VALIDATE CHARACTER 4 =====")
	data_split := strings.Split(data, ",")
	var longest string
	for _, str := range data_split {
		if len(str) > len(longest) {
			longest = str
		}
	}

	fmt.Println("Longest string: ", longest)
}

// Exercise 3 - 5
func validateCharacter_5() {
	fmt.Println("===== VALIDATE CHARACTER 5 =====")

	raw_string := "   hELlo   woRLd         tHis is   Go    "
	handled_string := strings.Fields(raw_string)
	fmt.Println(handled_string)
}

// Exercise 4
func sumPrimeNumber() {
	fmt.Println("============== SUM PRIME NUMBER ==============")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum int = 0
	for _, number := range array {
		if isPrimeNumber(number) {
			sum += number
		}
	}

	fmt.Println(sum)
}

func isPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// Exercise 5
func matrix() {
	fmt.Println("============== MATRIX ==============")
	var matrix [][]int
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	fmt.Println(matrix)

	// Print matrix
	fmt.Println("Matrix: ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	// Print matrix
	fmt.Println("Matrix: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Mấy câu kia dễ quá khỏi làm
