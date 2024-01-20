package ch6

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name     string
	Age      int
	Password string
}
type flyStyles interface {
	fly(Name string) string
}

func (p Person) Fly(Name string) string {
	return p.Name + Name + "起飞"
}

type AirPlane struct {
	Name string
}

func (plane AirPlane) Fly(str string) string {
	return plane.Name + str + "起飞"
}

// Stringer
// fmt 包中定义的 Stringer 是最普遍的接口之一。
//
//	type Stringer interface {
//	   String() string
//	}
//
// Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years,%v )", p.Name, p.Age, p.Password)
}

func (p Person) Error() string {
	return p.Name
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var str string
	for _, b := range ip {
		str = str + strconv.Itoa(int(b)) + "."
	}
	strIp := str[:len(str)-1]
	return strIp
}

func C() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}

func Do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case Person:
		fmt.Printf("%q is %v Person Person\n", v, i)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// error  learning form  strconv.Atoi.go

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, sqrtError("sqrt", strconv.Itoa(int(x)))
	}
	return x, nil
}

type NumberError struct {
	Func string
	Num  string
	Err  error
}

var ErrSyntax = errors.New("negative number")

func (e *NumberError) Error() string {
	return "ch6." + e.Func + ": " + ": " + e.Num + ": " + e.Err.Error()
}

func sqrtError(fn, str string) *NumberError {
	return &NumberError{fn, str, ErrSyntax}
}

//io

func ReaderFromString() {
	r := strings.NewReader("Hello, Reader!")

	buffer := make([]byte, 8)
	for {
		n, err := r.Read(buffer)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}

}

func WriteToFile() {
	// 打开文件，如果文件不存在则创建，以追加模式打开
	filePath := "hello.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 要追加的内容
	content := "\nget a new start."

	// 写入内容到文件末尾
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Content appended to file.")
}

func ReaderFromFile() {
	// 打开文件
	filePath := "hello.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 输出文件内容
	fmt.Println("File Content:")
	fmt.Println(string(data))
}
func GetSomeIMG() {
	// 创建一个200x100的RGBA图像
	img := image.NewRGBA(image.Rect(0, 0, 200, 100))

	// 将图像的每个像素设置为红色
	for y := 0; y < 100; y++ {
		for x := 0; x < 200; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	// 创建一个输出文件
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将图像写入文件
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
