package main

import (
	"log"
	"math/rand"
	"plugin"
	"time"
)

// go build -buildmode=plugin plugs/dynamic.go
// go run main.go dynamic.so

// LoadPlugin 负责从指定的 so 文件中加载两个函数
// LoadPlugin 函数的返回参数是 so 文件中两个函数的签名
func LoadPlugin(filename string) (func() string, func(rand.Source) int) {
	p, err := plugin.Open(filename)
	if err != nil {
		log.Fatalf("load plugin %s fail, %v", filename, err)
	}
	symbolGreeting, err := p.Lookup("Greeting")
	if err != nil {
		log.Fatalf("load Greeting func fail, %v", err)
	}
	greeting := symbolGreeting.(func() string)

	symbolTimes, err := p.Lookup("Times")
	if err != nil {
		log.Fatalf("load Times func fail, %v", err)
	}
	times := symbolTimes.(func(rand.Source) int)

	return greeting, times
}

func main() {
	greeting, times := LoadPlugin("dynamic.so")
	newSource := rand.NewSource(time.Now().UnixNano())
	greetingTimes := times(newSource)
	greetingContent := greeting()

	for i := 0; i < greetingTimes; i++ {
		log.Println(greetingContent)
	}
}
