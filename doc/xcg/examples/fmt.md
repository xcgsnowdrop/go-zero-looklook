`fmt` 包是 Go 标准库中的一个用于格式化输入和输出的包。它提供了各种函数和方法，用于将数据格式化为字符串以进行输出，以及从字符串中解析数据。以下是 `fmt` 包的一些常见用法：

### 格式化输出

1. `fmt.Print`、`fmt.Println` 和 `fmt.Printf`：这些函数用于将数据输出到标准输出。`Println` 会在输出后自动添加换行符。

```go
a := 42
b := "Hello, Gophers!"
fmt.Print("a is ", a, ", b is ", b) // 输出：a is 42, b is Hello, Gophers!
fmt.Println("a is", a)               // 输出：a is 42（自动换行）
fmt.Printf("a is %d, b is %s\n", a, b) // 输出：a is 42, b is Hello, Gophers!
```

2. `fmt.Sprint`、`fmt.Sprintln` 和 `fmt.Sprintf`：这些函数将数据格式化为字符串，而不是直接输出。它们返回格式化后的字符串。

```go
a := 42
b := "Hello, Gophers!"
result := fmt.Sprint("a is ", a, ", b is ", b) // 返回字符串
fmt.Println(result)
```

### 格式化输入

1. `fmt.Scan`、`fmt.Scanln` 和 `fmt.Scanf`：这些函数用于从标准输入读取并解析数据。与输出函数类似，`Scan` 和 `Scanln` 在读取后会自动根据换行符或空白字符分隔数据。

```go
var a int
var b string
fmt.Print("Enter an integer and a string: ")
fmt.Scan(&a, &b)
fmt.Printf("You entered: %d and %s\n", a, b)
```

2. `fmt.Sscan`、`fmt.Sscanln` 和 `fmt.Sscanf`：这些函数用于从字符串中读取并解析数据。

```go
input := "42 Hello"
var a int
var b string
fmt.Sscanf(input, "%d %s", &a, &b)
fmt.Printf("Parsed values: %d and %s\n", a, b)
```

### 格式化占位符

`fmt` 包使用一些特定的占位符来控制格式化输出和输入。以下是一些常用的占位符：

- `%d`: 整数
- `%s`: 字符串
- `%f`: 浮点数
- `%t`: 布尔值
- `%v`: 默认格式
- `%T`: 类型
- `%p`: 指针
- `%x`: 十六进制格式

```go
a := 42
b := "Hello"
fmt.Printf("a is %d, b is %s\n", a, b)
```

### 格式化修饰符

在占位符之前可以添加修饰符，以改变格式化的行为，例如控制字段宽度、小数点精度等。

```go
x := 3.14159265
fmt.Printf("x is %.2f (2 decimal places)\n", x)
```

### 格式化错误

`fmt` 包还提供了用于格式化错误消息的函数，如 `fmt.Errorf`。

```go
err := fmt.Errorf("An error occurred: %s", "something went wrong")
fmt.Println(err)
```

总之，`fmt` 包是 Go 中用于格式化输入和输出的重要工具，可以帮助你以各种方式显示和解析数据。无论是在控制台应用程序、Web服务还是其他 Go 程序中，`fmt` 包都具有广泛的应用。