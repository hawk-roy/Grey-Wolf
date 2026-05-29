package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"
)

// 9. 作用域：包级变量
var globalCount = 100

// 12. 结构体
// Person 演示结构体定义和方法。
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("我是%s，今年%d岁", p.Name, p.Age)
}

// 18. 接口
// Speaker 定义行为，具体类型只要实现 Speak() 即可。
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "汪汪，我是" + d.Name
}

// 19. 泛型
func Sum[T int | float64](a, b T) T {
	return a + b
}

// 20. 错误处理
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// 16. 递归
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 24. 继承（Go 中通过组合/嵌入实现）
type Animal struct {
	Name string
}

func (a Animal) Eat() string {
	return a.Name + " 正在吃东西"
}

type Cat struct {
	Animal
}

func (c Cat) Meow() string {
	return c.Name + "：喵喵"
}

func main() {
	fmt.Println("=== 1. 基础语法 ===")
	demoBasicSyntax()

	fmt.Println("\n=== 2. 数据类型 ===")
	demoDataTypes()

	fmt.Println("\n=== 3. 变量 ===")
	demoVariables()

	fmt.Println("\n=== 4. 常量 ===")
	demoConstants()

	fmt.Println("\n=== 5. 运算符 ===")
	demoOperators()

	fmt.Println("\n=== 6. 条件语句 ===")
	demoDecisionMaking(82)

	fmt.Println("\n=== 7. 循环 ===")
	demoLoops()

	fmt.Println("\n=== 8. 函数 ===")
	fmt.Println("3 + 5 =", add(3, 5))

	fmt.Println("\n=== 9. 作用域规则 ===")
	demoScope()

	fmt.Println("\n=== 10. 数组 ===")
	demoArrays()

	fmt.Println("\n=== 11. 指针 ===")
	demoPointers()

	fmt.Println("\n=== 12. 结构体 ===")
	demoStructs()

	fmt.Println("\n=== 13. 切片 ===")
	demoSlices()

	fmt.Println("\n=== 14. range ===")
	demoRange()

	fmt.Println("\n=== 15. map ===")
	demoMap()

	fmt.Println("\n=== 16. 递归 ===")
	fmt.Println("5! =", factorial(5))

	fmt.Println("\n=== 17. 类型转换 ===")
	demoTypeCasting()

	fmt.Println("\n=== 18. 接口 ===")
	demoInterfaces()

	fmt.Println("\n=== 19. 泛型 ===")
	demoGenerics()

	fmt.Println("\n=== 20. 错误处理 ===")
	demoErrorHandling()

	fmt.Println("\n=== 21. 并发 ===")
	demoConcurrency()

	fmt.Println("\n=== 22. 文件处理 ===")
	demoFileHandling()

	fmt.Println("\n=== 23. 正则表达式 ===")
	demoRegex()

	fmt.Println("\n=== 24. 类型断言 ===")
	demoTypeAssertion()

	fmt.Println("\n=== 25. 继承（嵌入） ===")
	demoInheritance()

	fmt.Println("\n=== 26. IDE ===")
	fmt.Println("建议：使用 VS Code + Go 扩展，启用保存时格式化和自动补全")

	fmt.Println("\n=== 27. Modules ===")
	fmt.Println("常用命令：go mod init、go mod tidy、go get")
}

// 1. 基础语法
func demoBasicSyntax() {
	message := "Hello Go"
	fmt.Println(message)
}

// 2. 数据类型
func demoDataTypes() {
	var i int = 42
	var f float64 = 3.14
	var b bool = true
	var s string = "Go"
	fmt.Printf("int=%d float=%.2f bool=%t string=%s\n", i, f, b, s)
}

// 3. 变量
func demoVariables() {
	var name string
	name = "张三"
	age := 20
	fmt.Printf("name=%s age=%d\n", name, age)
}

// 4. 常量
func demoConstants() {
	const Pi = 3.1415926
	fmt.Printf("Pi=%.4f\n", Pi)
}

// 5. 运算符
func demoOperators() {
	a, b := 10, 3
	fmt.Printf("+:%d -:%d *:%d /:%d %%:%d\n", a+b, a-b, a*b, a/b, a%b)
}

// 6. 条件语句
func demoDecisionMaking(score int) {
	if score >= 90 {
		fmt.Println("成绩：优秀")
	} else if score >= 60 {
		fmt.Println("成绩：及格")
	} else {
		fmt.Println("成绩：不及格")
	}
}

// 7. 循环
func demoLoops() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("1~5 累加和:", sum)
}

// 8. 函数
func add(a, b int) int {
	return a + b
}

// 9. 作用域
func demoScope() {
	globalCount := 999
	fmt.Println("局部 globalCount:", globalCount)
	fmt.Println("包级 globalCount:", getPackageGlobalCount())
}

func getPackageGlobalCount() int {
	return globalCount
}

// 10. 数组
func demoArrays() {
	arr := [3]int{1, 2, 3}
	fmt.Println("数组:", arr, "长度:", len(arr))
}

// 11. 指针
func demoPointers() {
	x := 10
	p := &x
	*p = 20
	fmt.Println("x:", x, "p 指向值:", *p)
}

// 12. 结构体
func demoStructs() {
	p := Person{Name: "李四", Age: 30}
	fmt.Println(p.Introduce())
}

// 13. 切片
func demoSlices() {
	nums := []int{1, 2, 3, 4, 5}
	sub := nums[1:4]
	fmt.Println("原切片:", nums, "子切片:", sub)
}

// 14. range
func demoRange() {
	words := []string{"Go", "is", "fun"}
	for i, w := range words {
		fmt.Printf("index=%d value=%s\n", i, w)
	}
}

// 15. map
func demoMap() {
	scores := map[string]int{"张三": 90, "李四": 85}
	scores["王五"] = 88
	fmt.Println("map:", scores)
}

// 17. 类型转换
func demoTypeCasting() {
	i := 9
	f := float64(i)
	fmt.Printf("int=%d -> float64=%.1f\n", i, f)
}

// 18. 接口
func demoInterfaces() {
	var s Speaker = Dog{Name: "旺财"}
	fmt.Println(s.Speak())
}

// 19. 泛型
func demoGenerics() {
	fmt.Println("Sum[int](3,5)=", Sum(3, 5))
	fmt.Println("Sum[float64](1.2,3.4)=", Sum(1.2, 3.4))
}

// 20. 错误处理
func demoErrorHandling() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("发生错误:", err)
		return
	}
	fmt.Println("结果:", result)
}

// 21. 并发
func demoConcurrency() {
	jobs := []int{2, 4, 6}
	ch := make(chan int, len(jobs))
	var wg sync.WaitGroup

	for _, n := range jobs {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			ch <- num * num
		}(n)
	}

	wg.Wait()
	close(ch)

	for v := range ch {
		fmt.Println("平方结果:", v)
	}
}

// 22. 文件处理
func demoFileHandling() {
	f, err := os.CreateTemp("", "go-demo-*.txt")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.WriteString("Go 文件读写示例\n"); err != nil {
		fmt.Println("写文件失败:", err)
		return
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println("读文件失败:", err)
		return
	}
	fmt.Println("文件内容:", string(data))
}

// 23. 正则表达式
func demoRegex() {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString("订单A12, B34, C567", -1)
	fmt.Println("匹配到数字:", matches)
}

// 24. 类型断言
func demoTypeAssertion() {
	var x any = "Go语言"
	s, ok := x.(string)
	fmt.Println("断言结果:", s, "是否成功:", ok)
}

// 25. 继承（嵌入）
func demoInheritance() {
	cat := Cat{Animal: Animal{Name: "咪咪"}}
	fmt.Println(cat.Eat())
	fmt.Println(cat.Meow())
}
