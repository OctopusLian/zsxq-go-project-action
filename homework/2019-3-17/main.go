package main

import (
	"fmt"
)

func main() {
	text := "Go中文网"
	// text := "中国"
	encoderx(text)
	encoderu(text)
	encoderU(text)
}

func encoderx(text string) {
	var etext string
	for _, s := range []byte(text) {
		etext += fmt.Sprintf("\\x%x", s)
	}
	fmt.Println(etext)
}

func encoderu(text string) {
	var etext string

	for _, s := range []rune(text) {
		etext += fmt.Sprintf("\\u%x", s)
	}
	fmt.Println(etext)
}

func encoderU(text string) {
	var etext string
	for _, s := range []rune(text) {
		etext += fmt.Sprintf("\\U%08x", s)
	}
	fmt.Println(etext)
}