package ch8

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func Func4Reflect(data any) {
	typ := reflect.TypeOf(data)
	fmt.Println("类别：", typ.Kind()) // 类别： struct
	fmt.Println("类型：", typ.Name()) // 类型： User

	val := reflect.ValueOf(data)
	// 获取结构体字段的数量
	numField := val.NumField()

	for i := 0; i < numField; i++ {
		fmt.Println("字段名称：", typ.Field(i).Name)
		fmt.Println("字段类型：", typ.Field(i).Type.Name())
		fmt.Println("字段值：", val.Field(i).Interface())
		fmt.Println("----------------------------")
	}
}

func GetHttp() {
	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	resp, err := client.Get("https://cloud.tencent.com/developer/article/2368120")

	if err != nil {
		return
	}
	str, err1 := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err1 != nil {
		return
	}
	fmt.Println(string(str))

}

// Package traceid 用于在 context 中维护 trace ID

// WithTraceID 往 context 中存入 trace ID
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// TraceID 从 context 中提取 trace ID
func TraceID(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	//v := context.Value(ctx, traceIDKey{})
	id, _ := v.(string)
	return id
}

type traceIDKey struct{}
