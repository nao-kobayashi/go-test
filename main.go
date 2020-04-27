package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println(math.Floor(2.1))
	fmt.Println(strings.Title("head first go"))

	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf(999999999999))
	fmt.Println(reflect.TypeOf(2.1))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("aaa"))
	n := 1
	fmt.Println(n)

	var n1 = &n
	fmt.Println(n1)
	fmt.Println(*n1)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(input)

	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(grade)

	fileinfo, err := os.Stat("./main.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("name => ", fileinfo.Name(), " size => ", fileinfo.Size())

	notes := [4]string{"a", "b", "c", "d"}
	for index, note := range notes {
		fmt.Println(index, " ", note)
	}

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	if file == nil {
		log.Fatal("file is nil")
		return
	}
	ShowFile(file)

	floats, err := GetFloats("./data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	var sum = 0.0
	for _, f := range floats {
		fmt.Println("read from data.txt", f)
		sum += f
	}
	fmt.Printf("Sum: %0.2f\n", sum)
	fmt.Printf("SliceCount: %d\n", len(floats))

	fmt.Println(MaxNumber(71.8, 56.2, 89.5))
	fmt.Println(MaxNumber(71.8, 56.2, 89.5, 98.5, 90.7))

	s := "a"
	fmt.Println(s)
	s += "b"
	fmt.Println(s)

	var num myNumber = 11
	num.Double()
	num.Display()
}

type myNumber int

func (m *myNumber) Double() {
	*m *= 2
}

func (m *myNumber) Display() {
	fmt.Println(m)
	fmt.Println(*m)
}

// MaxNumber 与えられた可変の浮動小数点から最大値を返す。
func MaxNumber(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// ShowFile ファイルの中身をコンソールに出力しファイルを閉じる。
func ShowFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := file.Close()
	if err != nil {
		log.Fatal("file is nil")
	}
}

// GetFloats data.txtファイルから値を読み取り浮動小数点の配列を返す。
func GetFloats(fileName string) ([]float64, error) {
	numbers := make([]float64, 0, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, num)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
