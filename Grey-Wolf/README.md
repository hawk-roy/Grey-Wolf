# GoField：Go 语言入门知识点文档

> **基于菜鸟教程 Go 目录**，结合实际练习代码，整理成可运行示例 + 知识点讲解 + 代码疑点解析。
>
> 适合有其他编程语言（如 C#、Java、Python）基础，初次学习 Go 的开发者。

---

## 目录

- [仓库结构](#仓库结构)
- [快速运行](#快速运行)
- [知识点地图](#知识点地图)
- [精选示例代码](#精选示例代码)
- [代码疑点详解](#代码疑点详解)
- [代码疑点速查](#代码疑点速查)
- [visibility 示例讲解](#visibility-示例讲解)
- [建议练习](#建议练习)
- [学习路线建议](#学习路线建议)
- [提交规范建议](#提交规范建议)
- [参考资料](#参考资料)

---

这个仓库用于记录 Go 语言入门阶段的学习过程。内容以菜鸟教程 Go 目录为主线，结合 `runoob_examples` 中的可运行代码，把基础语法、数据类型、函数、作用域、数组、切片、map、接口、泛型、错误处理、并发、文件处理、正则表达式、类型断言和结构体嵌入等知识点整理成一份可以持续更新的学习文档。

当前推荐阅读顺序：

1. 先阅读本 README，建立知识点地图。
2. 运行 `runoob_examples/main.go`，观察每个章节的输出。
3. 对照“代码疑点讲解”修改示例代码，验证自己的理解。
4. 阅读 `visibility` 目录，理解 Go 的跨包可见性规则。

## 仓库结构

```text
.
├── README.md
├── go.mod
├── main.go
├── runoob_examples/
│   ├── main.go
│   └── README.md
├── runoob_go_examples/
│   ├── main.go
│   └── README.md
└── visibility/
    ├── scope.go
    ├── scope_test.go
    ├── exported_test.go
    └── private_external_shouldfail_test.go
```

说明：

- `runoob_examples/main.go`：当前主要示例文件，覆盖完整 Go 入门知识点，**推荐从这里开始**。
- `runoob_examples/README.md`：示例目录的简要说明。
- `visibility/`：专门演示大写导出、小写不导出、包级变量和跨包访问规则。
- `main.go`：早期学习笔记式示例，可作为补充阅读。
- `runoob_go_examples/`：另一版按章节组织的示例，可作为对照材料。

**关键概念：** 一个目录 = 一个 Go 包。`gofield`、`gofield/runoob_examples`、`gofield/runoob_go_examples` 即使都写 `package main`，也是三个独立的可执行程序，互不干扰。

## 快速运行

确认 Go 版本：

```bash
go version
```

运行主要示例：

```bash
go run ./runoob_examples
```

运行全部可通过的测试：

```bash
go test ./...
```

验证“小写标识符不能跨包访问”的失败示例：

```bash
go test -tags shouldfail ./visibility
```

上面这条命令预期会编译失败，因为 `visibility.privateConstant` 是小写开头，不能被 `visibility_test` 这个外部测试包访问。

## 知识点地图

| 编号 | 主题 | 对应示例 | 核心要点 | 常见疑点 |
| --- | --- | --- | --- | --- |
| 1 | 基础语法 | `basicSyntax` | `package`、`import`、`func main`、字符串拼接、`fmt.Sprintf` | 左花括号 `{` 不能随意换行 |
| 2 | 数据类型 | `dataTypes` | `int`、`float64`、`bool`、`string`、`rune`、`byte` | `rune` 表示 Unicode 字符，`byte` 表示一个字节 |
| 3 | 变量 | `variables` | `var`、零值、短变量声明、包级变量 | `:=` 只能在函数内部使用 |
| 4 | 常量 | `constants` | `const`、常量分组、`iota` | `iota` 在常量组内从 0 自动递增 |
| 5 | 运算符 | `operators` | 算术、关系、逻辑、位运算、赋值、取地址 | 整数除法会截断小数 |
| 6 | 条件语句 | `decisionMaking` | `if`、`else if`、`else`、`switch` | 条件表达式不需要小括号 |
| 7 | 循环 | `loops` | Go 只有 `for`，可模拟 while 和无限循环 | 无限循环必须设计退出条件 |
| 8 | 函数 | `functions`、`add`、`divMod`、`apply` | 参数、返回值、命名返回值、函数作为值 | 命名返回值可直接 `return`，但不要滥用 |
| 9 | 作用域 | `scopeRules` | 局部变量、包级变量、参数作用域、变量遮蔽 | 内层 `:=` 可能创建新变量而不是修改外层变量 |
| 10 | 数组 | `arrays` | 固定长度集合，长度属于类型的一部分 | `[3]int` 和 `[4]int` 是不同类型 |
| 11 | 指针 | `pointers` | `&` 取地址，`*` 解引用 | Go 有指针，但不支持指针运算 |
| 12 | 结构体 | `structures`、`Product`、`Student` | 字段组合、结构体字面量、方法 | 方法接收者可以是值，也可以是指针 |
| 13 | 切片 | `slices` | 动态长度视图，`len`、`cap`、`append`、`copy` | `append` 可能触发底层数组扩容 |
| 14 | range | `rangeExamples` | 遍历切片、map、字符串 | 遍历字符串时 index 是字节下标，不是字符序号 |
| 15 | map | `maps` | 键值对集合，新增、查询、删除 | 读取 map 建议使用 `value, ok` 判断键是否存在 |
| 16 | 递归 | `recursion`、`factorial`、`fibonacci` | 函数调用自身，必须有终止条件 | 斐波那契的朴素递归会重复计算 |
| 17 | 类型转换 | `typeCasting` | 显式数值转换，`strconv` 字符串转换 | Go 不做大多数隐式数值转换 |
| 18 | 接口 | `interfaces`、`Shape` | 接口定义行为，结构体隐式实现接口 | 不需要写 `implements` |
| 19 | 泛型 | `generics`、`Number`、`SumNumbers` | 类型参数、类型约束、算法复用 | `~int` 表示底层类型为 `int` 的自定义类型也可用 |
| 20 | 错误处理 | `errorHandling`、`divide`、`UserError` | 返回 `error`，调用方显式检查 | `err != nil` 是 Go 常见控制流 |
| 21 | 并发 | `concurrency` | goroutine、channel、`sync.WaitGroup` | 循环变量传入 goroutine 时要作为参数传进去 |
| 22 | 文件处理 | `fileHandling` | 临时文件、写入、关闭、读取、清理 | 读写文件要处理错误，并适合用 `defer` 清理资源 |
| 23 | 正则表达式 | `regexExamples` | `regexp.MustCompile`、`MatchString`、`FindAllString` | `MustCompile` 遇到非法规则会 panic |
| 24 | 类型断言 | `typeAssertion`、`describeAny` | `value.(T)`、`value, ok`、type switch | 不带 `ok` 的断言失败会 panic |
| 25 | 组合模拟继承 | `embeddingInsteadOfInheritance` | 结构体嵌入、方法提升、方法覆盖 | Go 没有 class 继承，推荐组合和接口 |
| 26 | IDE 与 Modules | `modulesAndIDE`、`go.mod` | 模块名、运行入口、GoLand/终端运行 | 包路径和模块名有关 |

## 精选示例代码

### 1. 程序入口和章节组织

`runoob_examples/main.go` 用一个 `main` 函数串联所有知识点，每个章节由 `title` 分隔，便于运行后按输出顺序学习。

```go
func main() {
	title("01 基础语法")
	basicSyntax()

	title("02 数据类型")
	dataTypes()

	title("03 变量")
	variables()
}

func title(text string) {
	fmt.Printf("\n========== %s ==========\n", text)
}
```

这种组织方式适合学习阶段：每个知识点都是一个小函数，后续新增练习时只需要加一个函数，再在 `main` 中调用。

### 2. 变量、零值和包级变量

```go
var packageLevelCounter = 100

func variables() {
	var city string
	var score int = 90
	country := "China"

	fmt.Printf("zero value city=%q\n", city)
	fmt.Printf("score=%d, country=%s\n", score, country)

	packageLevelCounter++
	fmt.Println("packageLevelCounter:", packageLevelCounter)
}
```

要点：

- `city` 没有显式赋值时，字符串零值是空字符串 `""`。
- `country := "China"` 是短变量声明，只能写在函数内部。
- `packageLevelCounter` 在函数外声明，是包级变量，同一个包内的函数都能访问。

### 3. 作用域和变量遮蔽

```go
func scopeRules() {
	local := 10
	fmt.Println("local:", local)
	fmt.Println("packageLevelCounter:", packageLevelCounter)

	value := "outer"
	if true {
		value := "inner"
		fmt.Println("inner value:", value)
	}
	fmt.Println("outer value:", value)
}
```

疑点解释：

- `if` 里面的 `value := "inner"` 创建了一个新的局部变量。
- 它不会修改外层的 `value := "outer"`。
- 如果想修改外层变量，应写成 `value = "inner"`（不带 `:`）。

### 4. 指针修改原变量

```go
func pointers() {
	value := 42
	ptr := &value

	fmt.Println("value:", value)
	fmt.Println("address:", ptr)
	fmt.Println("dereference:", *ptr)

	*ptr = 100
	fmt.Println("changed through pointer:", value)
}
```

要点：

- `&value` 取得变量地址。
- `*ptr` 取得地址里保存的值。
- `*ptr = 100` 会修改 `value` 本身，输出 `100`。

### 5. 切片、长度和容量

```go
func slices() {
	numbers := []int{0, 1, 2, 3, 4}
	fmt.Printf("numbers len=%d cap=%d value=%v\n", len(numbers), cap(numbers), numbers)

	part := numbers[1:4]
	fmt.Println("numbers[1:4]:", part)

	numbers = append(numbers, 5, 6)
	fmt.Printf("after append len=%d cap=%d value=%v\n", len(numbers), cap(numbers), numbers)
}
```

疑点解释：

- `numbers[1:4]` 包含下标 1、2、3，不包含 4（左闭右开）。
- `len` 是当前元素数量，`cap` 是底层数组从切片起点开始还能容纳的数量。
- `append` 后如果容量不够，Go 会分配新的底层数组，`cap` 可能翻倍。

### 6. map 查询和 ok 判断

```go
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
```

要点：

- `ages["Bob"] = 19` 表示新增或更新。
- `age, ok := ages["Alice"]` 中，`ok` 表示键是否存在；如果键不存在，`age` 为零值 `0`。
- `delete(ages, "Tom")` 删除不存在的键也不会报错。

### 7. 接口和多态

```go
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
```

疑点解释：

- `Shape` 只要求有 `Area() float64` 这个行为。
- `Rectangle` 和 `Circle` 只要实现了 `Area` 方法，就自动满足 `Shape` 接口。
- Go 不需要显式写 `implements Shape`，只看方法集合是否匹配。

### 8. 泛型和类型约束

```go
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
```

要点：

- `T Number` 表示类型参数 `T` 必须满足 `Number` 约束。
- `~int` 表示不仅 `int` 可以用，底层类型是 `int` 的自定义类型也可以用（例如 `type Age int`）。
- 因为约束里的类型都支持 `+`，所以函数体里可以写 `total += value`。

### 9. 错误处理

```go
var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
```

调用方式：

```go
result, err := divide(10, 0)
if err != nil {
	fmt.Println("divide failed:", err)
	return
}
fmt.Println(result)
```

Go 的错误处理不是异常机制，而是把错误作为普通返回值交给调用方判断。这样代码更显式，也更容易看出失败路径。

### 10. goroutine、channel 和 WaitGroup

```go
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
	fmt.Println("sum of squares:", total) // 1+4+9+16 = 30
}
```

疑点解释：

- `go func(n int) { ... }(value)` 把循环变量作为参数传入，每个 goroutine 拿到独立的 `n`，不会共享同一个变量。
- `results` 使用**缓冲 channel**，容量等于任务数量，goroutine 发送时不会阻塞。
- `wg.Wait()` 等所有 goroutine 完成后，才执行 `close(results)`；否则可能 panic（向已关闭的 channel 发送数据）。
- `for result := range results` 会一直读取直到 channel 关闭。

### 11. 类型断言和 type switch

```go
func typeAssertion() {
	var value any = "hello"

	text, ok := value.(string)
	fmt.Println("assert string:", text, ok)

	number, ok := value.(int)
	fmt.Println("assert int:", number, ok)
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
```

要点：

- `text, ok := value.(string)` 是**安全写法**，断言失败时 `ok = false`，不会崩溃。
- 如果写成 `text := value.(string)`，断言失败会直接 `panic`。
- 类型分支较多时，用 `type switch` 比多个 `if ok` 更清晰。

### 12. 结构体嵌入不是传统继承

```go
type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

type Dog struct {
	Animal        // 嵌入，不是 extends
	Breed  string
}

func (d Dog) Speak() string {
	return d.Name + " barks" // Dog 覆盖了 Animal 的 Speak
}
```

疑点解释：

- `Dog` 嵌入了 `Animal`，所以可以直接访问 `d.Name`（字段提升）。
- `dog.Speak()` 调用的是 `Dog` 自己的 `Speak`（方法覆盖）。
- 仍然可以通过 `dog.Animal.Speak()` 显式调用被嵌入类型的方法。
- 这不是 class 继承，而是组合。Go 更推荐“小接口 + 组合”的设计。

## visibility 示例讲解

`visibility` 目录用测试演示 Go 的导出规则：

```go
package visibility

const PublicConstant = "I can be accessed from another package."
const privateConstant = "I can only be accessed inside package visibility."

var globalVar = "I am a package-level global variable."
```

规则：

- 大写开头：导出，其他包可以访问，例如 `PublicConstant`。
- 小写开头：未导出，只能在当前包访问，例如 `privateConstant`、`globalVar`。
- “全局变量”在 Go 中通常说成“包级变量”，它属于当前包，不是所有包都能随意访问。

测试文件的区别：

- `scope_test.go` 使用 `package visibility`，属于同一个包，所以能访问 `privateConstant`。
- `exported_test.go` 使用 `package visibility_test`，属于外部测试包，只能访问导出的 `PublicConstant`。
- `private_external_shouldfail_test.go` 带有 `shouldfail` build tag，用来证明外部包访问小写标识符会失败。

## 代码疑点详解

### iota 是什么？

```go
const (
	StatusPending = iota // 0
	StatusRunning        // 1
	StatusDone           // 2
)
```

`iota` 是常量组里的自动递增计数器，从 `0` 开始，每行加 `1`。常用于模拟枚举。

### %08b 格式化含义

```go
fmt.Printf("bit: x&y=%08b\n", x&y)
```

拆解：

| 符号 | 含义 |
| --- | --- |
| `%` | 格式开始 |
| `0` | 不够宽度时用 `0` 填充 |
| `8` | 至少占 8 位宽度 |
| `b` | 按二进制打印 |

例如 `fmt.Printf("%08b\n", 5)` 输出 `00000101`。

### fmt 格式符速查

| 格式符 | 含义 |
| --- | --- |
| `%v` | 默认格式 |
| `%+v` | 结构体带字段名 |
| `%#v` | Go 语法格式 |
| `%T` | 打印类型 |
| `%s` | 字符串 |
| `%d` | 十进制整数 |
| `%f` | 浮点数 |
| `%t` | bool |
| `%c` | 字符 |
| `%b` | 二进制 |
| `%q` | 带引号的字符串 |

### strconv 字符串转换

```go
num, err := strconv.Atoi("123")    // 字符串 -> int
text := strconv.Itoa(456)          // int -> 字符串
f, err := strconv.ParseFloat("3.14", 64) // 字符串 -> float64
```

Go 不做隐式类型转换，字符串和数值之间转换必须使用 `strconv` 包。

### defer 执行时机

```go
func example() {
	defer fmt.Println("最后执行")
	fmt.Println("先执行")
}
// 输出：
// 先执行
// 最后执行
```

`defer` 会在**当前函数返回前**执行。常用场景：关闭文件、解锁、删除临时资源。

多个 `defer` 按**后进先出（LIFO）**顺序执行：

```go
defer fmt.Println("A")
defer fmt.Println("B")
defer fmt.Println("C")
// 输出顺序：C B A
```

### defer 恢复测试中的全局变量

```go
func TestGlobalVarIsPackageLevelVariable(t *testing.T) {
	original := globalVar
	defer func() {
		globalVar = original // 测试结束前恢复原值
	}()

	globalVar = "changed in package visibility"
	if globalVar != "changed in package visibility" {
		t.Fatal("globalVar should be mutable from package scope")
	}
}
```

这是一种常见的测试技巧：测试前保存状态，`defer` 在测试结束时自动恢复，避免影响其他测试。

### build tag shouldfail

`visibility/private_external_shouldfail_test.go` 文件头有：

```go
//go:build shouldfail
// +build shouldfail
```

这表示只有启用 `shouldfail` tag 时，这个文件才参与编译。普通运行 `go test ./...` 时它会被忽略。

运行 `go test -tags shouldfail ./visibility` 时，文件参与编译，因为它尝试访问 `visibility.privateConstant`（小写，未导出），所以会**编译失败**，这正是期望的行为——用来证明小写标识符不能跨包访问。

### 无限循环和 break

```go
for {
	fmt.Println("only once")
	break
}
```

`for {}` 是无限循环，`break` 跳出循环。所以上面的代码只打印一次。

### fmt.Sprintf vs fmt.Printf

```go
url := fmt.Sprintf("Code=%d&endDate=%s", stockCode, endDate) // 返回字符串
fmt.Printf("Code=%d&endDate=%s\n", stockCode, endDate)       // 直接打印
```

- `Sprintf`：按格式生成字符串并**返回**，不打印。
- `Printf`：按格式**打印**到控制台。

### rune 和 byte 的区别

| 类型 | 底层类型 | 表示 |
| --- | --- | --- |
| `byte` | `uint8` | 一个字节（ASCII 字符） |
| `rune` | `int32` | 一个 Unicode 码点（如中文字符） |

遍历字符串时，`range` 按 `rune`（字符）迭代，`index` 是字节下标：

```go
for index, char := range "Go语言" {
	fmt.Printf("byteIndex=%d rune=%c\n", index, char)
}
// 输出中，"语" 的 byteIndex 是 2（不是 2），因为 G=0 o=1 语=2（占3字节）
```

### 1. 为什么 `:=` 不能写在函数外？

`:=` 是短变量声明，只能出现在函数体内部。包级作用域中必须使用 `var`、`const`、`type`、`func` 等声明形式。

正确：

```go
var packageLevelCounter = 100
```

错误（编译失败）：

```go
packageLevelCounter := 100 // syntax error: non-declaration statement outside function body
```

### 2. 为什么大写开头就能跨包访问？

Go 用标识符首字母大小写控制可见性：

- `PublicConstant`：导出，其他包可访问。
- `privateConstant`：未导出，只能本包访问。

这条规则适用于：常量、变量、函数、结构体、接口、方法和结构体字段。

### 3. `rune` 和 `byte` 有什么区别？

`byte` 本质是 `uint8`，表示一个字节。`rune` 本质是 `int32`，表示一个 Unicode 码点。中文字符通常不止一个字节，所以遍历中文字符串时要理解字节下标和字符的区别。

### 4. 数组和切片怎么选？

数组长度固定，长度也是类型的一部分。切片长度可变，更常用于业务代码。

```go
var arr [3]int         // 数组，长度固定
nums := []int{1, 2, 3} // 切片，长度可变
```

学习阶段要理解数组，但实际开发中更多使用切片。

### 5. map 遍历为什么顺序不固定？

Go 没有保证 map 的遍历顺序。每次运行时输出顺序都可能不同。如果需要稳定顺序，应先取出 key，排序后再访问。

### 6. 为什么错误处理总是 `if err != nil`？

Go 鼓励显式处理错误。函数把错误作为返回值交给调用方，调用方根据 `err != nil` 判断是否继续。

这种写法看起来多，但优点是失败路径很清楚，不会被隐藏在异常机制里。

### 7. goroutine 里为什么要传入循环变量？

在循环中启动 goroutine 时，直接捕获循环变量容易造成问题（多个 goroutine 可能读到同一个值）。推荐把当前值作为参数传给匿名函数：

```go
for _, value := range values {
	go func(n int) { // 参数 n 是当前值的副本
		fmt.Println(n)
	}(value) // 立刻传入当前的 value
}
```

### 8. `defer` 什么时候执行？

`defer` 会在当前函数返回前执行，常用于释放资源，例如关闭文件、解锁、删除临时文件。

```go
file, err := os.CreateTemp("", "gofield-*.txt")
if err != nil {
	return
}
defer os.Remove(file.Name()) // 函数返回前自动删除
```

### 9. 接口为什么不用显式实现？

Go 是结构化类型系统。一个类型只要拥有接口要求的方法，就自动满足接口。

```go
type Shape interface {
	Area() float64
}
```

任何有 `Area() float64` 方法的类型，都可以当作 `Shape` 使用。

### 10. 嵌入是不是继承？

不是。嵌入是组合的一种语法便利。它可以提升字段和方法，但不会形成 Java/C# 那种 class 继承体系。Go 中更常见的设计方式是：

- 用结构体保存数据。
- 用方法表达行为。
- 用小接口描述依赖。
- 用嵌入或组合复用已有能力。

## 建议练习

可以在 `runoob_examples/main.go` 后续继续添加这些练习：

1. 在 `typeCasting` 中增加字符串转整数失败的例子，例如 `strconv.Atoi("abc")`。
2. 在 `maps` 中实现一个单词计数器。
3. 在 `slices` 中尝试修改 `part[0]`，观察原切片是否变化。
4. 写一个 `Max[T Number]` 泛型函数。
5. 给 `UserError` 增加更多字段，比如错误码 `Code`。
6. 用 goroutine 并发计算 1 到 10 的平方，并汇总结果。
7. 在 `visibility` 中增加一个导出的函数，用来安全读取未导出的包级变量。

## 学习路线建议

**第一阶段：语法基础**

- `package`、`import`、变量、常量、函数、条件、循环。

**第二阶段：数据结构**

- 数组、切片、map、结构体、指针。

**第三阶段：抽象能力**

- 方法、接口、泛型、组合。

**第四阶段：工程能力**

- error、文件处理、正则、模块、测试、目录结构。

**第五阶段：并发入门**

- goroutine、channel、WaitGroup、关闭 channel 的时机。

**第六阶段（下一步补充）：**

- `defer / panic / recover`
- `context`
- `select`
- `sync.Mutex`、`sync.RWMutex`、`atomic`
- JSON 编解码
- HTTP server/client
- 表驱动测试（Table-driven tests）
- `time`、`filepath`、`log/slog`

## 提交规范建议

常用提交信息：

```text
feat: add slice and map practice examples
docs: update Go learning README
fix: handle divide by zero in error demo
test: add visibility tests
chore: run gofmt
```

建议每次提交只做一类事情。例如只改文档就用 `docs`，只加测试就用 `test`，避免一次提交混入太多无关修改。

## 参考资料

- 菜鸟教程 Go 专栏：https://www.runoob.com/go/go-tutorial.html
- Go 官方文档：https://go.dev/doc/
- Effective Go：https://go.dev/doc/effective_go
- Go by Example：https://gobyexample.com/
- Go 官方泛型教程：https://go.dev/doc/tutorial/generics

## License

这是学习记录项目，可按需选择 MIT License 或仅作为个人学习笔记保存。
