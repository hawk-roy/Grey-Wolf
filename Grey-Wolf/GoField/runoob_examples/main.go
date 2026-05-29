package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const AppName = "Go Runoob Examples"

const (
	StatusPending = iota
	StatusRunning
	StatusDone
)

var packageLevelCounter = 100

type Product struct {
	Name  string
	Price float64
}

type Student struct {
	Name  string
	Score int
}

func (s Student) Summary() string {
	return fmt.Sprintf("%s scored %d", s.Name, s.Score)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Number interface {
	~int | ~int64 | ~float64
}

func SumNumbers[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func PrintAny[T any](label string, value T) {
	fmt.Printf("%s: %v\n", label, value)
}

var ErrDivideByZero = errors.New("cannot divide by zero")

type UserError struct {
	Field string
	Msg   string
}

func (e UserError) Error() string {
	return e.Field + ": " + e.Msg
}

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() string {
	return d.Name + " barks"
}

func main() {
	title("01 基础语法")
	basicSyntax()

	title("02 数据类型")
	dataTypes()

	title("03 变量")
	variables()

	title("04 常量")
	constants()

	title("05 运算符")
	operators()

	title("06 条件语句")
	decisionMaking()

	title("07 循环")
	loops()

	title("08 函数")
	functions()

	title("09 作用域")
	scopeRules()

	title("10 数组")
	arrays()

	title("11 指针")
	pointers()

	title("12 结构体")
	structures()

	title("13 切片")
	slices()

	title("14 Range")
	rangeExamples()

	title("15 Map")
	maps()

	title("16 递归")
	recursion()

	title("17 类型转换")
	typeCasting()

	title("18 接口")
	interfaces()

	title("19 泛型")
	generics()

	title("20 错误处理")
	errorHandling()

	title("21 并发")
	concurrency()

	title("22 文件处理")
	fileHandling()

	title("23 正则表达式")
	regexExamples()

	title("24 类型断言")
	typeAssertion()

	title("25 组合模拟继承")
	embeddingInsteadOfInheritance()

	title("26 IDE 与 Modules")
	modulesAndIDE()
}

func title(text string) {
	fmt.Printf("\n========== %s ==========\n", text)
}

func basicSyntax() {
	// Go 程序由 package、import、声明和函数组成。main 包里的 main 函数是程序入口。
	fmt.Println("Hello, Go")

	// 一行通常就是一个语句，Go 会自动插入分号，实际代码中一般不手写分号。
	fmt.Println("Google" + "Runoob")

	stockCode := 123
	endDate := "2026-05-24"
	url := fmt.Sprintf("Code=%d&endDate=%s", stockCode, endDate)
	fmt.Println(url)
}

func dataTypes() {
	var age int = 18
	var price float64 = 19.99
	var active bool = true
	var name string = "Go"
	var letter rune = '中'
	var b byte = 'A'

	fmt.Printf("int=%d, float64=%.2f, bool=%t, string=%s\n", age, price, active, name)
	fmt.Printf("rune=%c, byte=%c\n", letter, b)
}

func variables() {
	var city string
	var score int = 90
	country := "China"

	fmt.Printf("zero value city=%q\n", city)
	fmt.Printf("score=%d, country=%s\n", score, country)

	packageLevelCounter++
	fmt.Println("packageLevelCounter:", packageLevelCounter)
}

func constants() {
	const Pi = 3.14159
	const Language = "Go"

	fmt.Println("Pi:", Pi)
	fmt.Println("Language:", Language)
	fmt.Println("iota:", StatusPending, StatusRunning, StatusDone)
}

func operators() {
	a, b := 21, 10
	fmt.Println("arithmetic:", a+b, a-b, a*b, a/b, a%b)
	fmt.Println("relation:", a == b, a != b, a > b, a <= b)
	fmt.Println("logic:", true && false, true || false, !true)

	var x uint = 60
	var y uint = 13
	fmt.Printf("bit: x&y=%08b, x|y=%08b, x^y=%08b\n", x&y, x|y, x^y)

	a += 5
	fmt.Println("assignment a += 5:", a)
	fmt.Println("address operator &a:", &a)
}

func decisionMaking() {
	score := 86
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	day := "Sunday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
}

func loops() {
	for i := 0; i < 3; i++ {
		fmt.Println("classic for:", i)
	}

	n := 3
	for n > 0 {
		fmt.Println("while style for:", n)
		n--
	}

	for {
		fmt.Println("break infinite loop")
		break
	}
}

func functions() {
	fmt.Println("add:", add(3, 5))

	quotient, remainder := divMod(17, 5)
	fmt.Println("divMod:", quotient, remainder)

	result := apply(10, func(value int) int {
		return value * value
	})
	fmt.Println("function as value:", result)
}

func add(a, b int) int {
	return a + b
}

func divMod(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func apply(value int, fn func(int) int) int {
	return fn(value)
}

func scopeRules() {
	// 局部变量只在函数内可见；包级变量在同一个包的所有文件中可见。
	local := 10
	fmt.Println("local:", local)
	fmt.Println("packageLevelCounter:", packageLevelCounter)

	// 内层变量会遮蔽外层同名变量。
	value := "outer"
	if true {
		value := "inner"
		fmt.Println("inner value:", value)
	}
	fmt.Println("outer value:", value)

	fmt.Println("parameter scope:", double(local))
}

func double(value int) int {
	return value * 2
}

func arrays() {
	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	names := [3]string{"Tom", "Jerry", "Spike"}
	fmt.Println("numbers:", numbers)
	fmt.Println("names:", names)
	fmt.Println("array length:", len(names))
}

func pointers() {
	value := 42
	ptr := &value

	fmt.Println("value:", value)
	fmt.Println("address:", ptr)
	fmt.Println("dereference:", *ptr)

	*ptr = 100
	fmt.Println("changed through pointer:", value)
}

func structures() {
	product := Product{Name: "Book", Price: 39.9}
	student := Student{Name: "Alice", Score: 95}

	fmt.Printf("product=%+v\n", product)
	fmt.Println(student.Summary())
}

func slices() {
	numbers := []int{0, 1, 2, 3, 4}
	fmt.Printf("numbers len=%d cap=%d value=%v\n", len(numbers), cap(numbers), numbers)

	part := numbers[1:4]
	fmt.Println("numbers[1:4]:", part)

	numbers = append(numbers, 5, 6)
	fmt.Printf("after append len=%d cap=%d value=%v\n", len(numbers), cap(numbers), numbers)

	copied := make([]int, len(numbers))
	copy(copied, numbers)
	fmt.Println("copied:", copied)
}

func rangeExamples() {
	colors := []string{"red", "green", "blue"}
	for index, color := range colors {
		fmt.Printf("slice range index=%d value=%s\n", index, color)
	}

	userScores := map[string]int{"Alice": 95, "Bob": 88}
	for name, score := range userScores {
		fmt.Printf("map range key=%s value=%d\n", name, score)
	}

	for index, char := range "Go语言" {
		fmt.Printf("string range byteIndex=%d rune=%c\n", index, char)
	}
}

func maps() {
	ages := map[string]int{
		"Tom":   18,
		"Alice": 20,
	}
	ages["Bob"] = 19

	age, ok := ages["Alice"]
	fmt.Println("Alice:", age, ok)

	delete(ages, "Tom")
	fmt.Println("after delete:", ages)
}

func recursion() {
	fmt.Println("factorial 5:", factorial(5))
	fmt.Println("fibonacci 6:", fibonacci(6))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func typeCasting() {
	sum := 17
	count := 5
	mean := float64(sum) / float64(count)
	fmt.Println("mean:", mean)

	num, err := strconv.Atoi("123")
	fmt.Println("Atoi:", num, err)

	text := strconv.Itoa(456)
	fmt.Println("Itoa:", text)

	floatValue, err := strconv.ParseFloat("3.14", 64)
	fmt.Println("ParseFloat:", floatValue, err)
}

func interfaces() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
	}

	for _, shape := range shapes {
		fmt.Printf("%T area=%.2f\n", shape, shape.Area())
	}
}

func generics() {
	PrintAny("string value", "hello")
	PrintAny("int value", 42)

	fmt.Println("sum int:", SumNumbers([]int{1, 2, 3}))
	fmt.Println("sum float64:", SumNumbers([]float64{1.5, 2.5, 3.5}))
}

func errorHandling() {
	result, err := divide(10, 2)
	fmt.Println("divide 10 / 2:", result, err)

	result, err = divide(10, 0)
	fmt.Println("divide 10 / 0:", result, err)

	err = validateName("")
	fmt.Println("validateName:", err)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func validateName(name string) error {
	if name == "" {
		return UserError{Field: "name", Msg: "must not be empty"}
	}
	return nil
}

func concurrency() {
	values := []int{1, 2, 3, 4}
	results := make(chan int, len(values))

	var wg sync.WaitGroup
	for _, value := range values {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(value)
	}

	wg.Wait()
	close(results)

	var total int
	for result := range results {
		total += result
	}
	fmt.Println("sum of squares:", total)
}

func fileHandling() {
	file, err := os.CreateTemp("", "gofield-*.txt")
	if err != nil {
		fmt.Println("create temp file:", err)
		return
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("Go file example\n")
	if err != nil {
		fmt.Println("write file:", err)
		return
	}

	if err := file.Close(); err != nil {
		fmt.Println("close file:", err)
		return
	}

	data, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("read file:", err)
		return
	}
	fmt.Println("file content:", strings.TrimSpace(string(data)))
}

func regexExamples() {
	text := "Contact us: alice@example.com or bob@example.org"
	emailPattern := regexp.MustCompile(`\b[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[A-Za-z]{2,}\b`)

	fmt.Println("contains email:", emailPattern.MatchString(text))
	fmt.Println("all emails:", emailPattern.FindAllString(text, -1))
}

func typeAssertion() {
	var value any = "hello"

	text, ok := value.(string)
	fmt.Println("assert string:", text, ok)

	number, ok := value.(int)
	fmt.Println("assert int:", number, ok)

	describeAny(99)
	describeAny("Go")
	describeAny(true)
}

func describeAny(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("type switch int:", v)
	case string:
		fmt.Println("type switch string:", v)
	default:
		fmt.Printf("type switch other %T: %v\n", v, v)
	}
}

func embeddingInsteadOfInheritance() {
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	fmt.Println(dog.Speak())
	fmt.Println(dog.Animal.Speak())
	fmt.Println("breed:", dog.Breed)
}

func modulesAndIDE() {
	fmt.Println("go.mod declares this project as module: gofield")
	fmt.Println("In GoLand, click the green triangle beside func main() to run this file.")
	fmt.Println("Terminal command: go run ./runoob_examples")
}
