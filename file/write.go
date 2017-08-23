package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	writeByteArrayToFile()

	writeToOpeningFile()

	writeThrowBuffer()

	writePrettyJson()

}
func writeByteArrayToFile() {

	data := []byte("{" +
		"\"value\": 13" +
		"}")

	err := ioutil.WriteFile("/tmp/file_write_simple.txt", data, 0644)
	if err != nil {
		panic(err)
	}

}

func writeToOpeningFile() {

	f, err := os.Create("/tmp/file_write_simple.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil{
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
}

func writeThrowBuffer() {
	f, err := os.Create("/tmp/file_write_thr_buf.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func writePrettyJson() {

	data := Message{"Alice", "Hello", 1294706395881547000}
	//or may get []byte
	//b, err := json.Marshal(m)

	b, err := json.MarshalIndent(&data, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("./file/message.json", b, 06444)
	if err != nil {
		panic(err)
	}

}
