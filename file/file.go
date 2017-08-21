package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

func main() {

	simpleReadFile()

	readPartFile()

	readFileByLine()

	readFileByLineToEnd()

	simpleCopyFiles()

	copyFilesWithBufio()

	copyFilesWithIOutil()
}

func readFileByLineToEnd() {

	fi, err := os.Open("./file/kipling.txt")
	if err != nil {
		panic(err)
	}

	stat, err := fi.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(stat.Size())

	ret, err := fi.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	_, err = fi.Seek(0, 0)
	if err != nil {
		panic(err)
	}


	fmt.Println(ret)

	r := bufio.NewReader(fi)

	str, err := r.ReadString('\n')

	fmt.Println(str)

	r.Discard(r.Buffered())

	off, err := fi.Seek(0, io.SeekCurrent)

	fmt.Println(off)

	fmt.Println(r.Buffered())

	for {
		if r.Buffered() != 0 {
			str, err := r.ReadString('\n')
			if err != nil{
				panic(err)
			}

			fmt.Println(str)
		}
	}

}

func simpleReadFile() {
	dat, err := ioutil.ReadFile("./file/kipling.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}

func readPartFile() {

	file, err := os.Open("./file/kipling.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	b1 := make([]byte, 50)
	n1, err := file.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

func readFileByLine() {
	fi, err := os.Open("./file/kipling.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	s := bufio.NewScanner(fi)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}

func simpleCopyFiles() {
	// open input file
	fi, err := os.Open("./file/kipling.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	fo, err := os.Create("./file/kipling_s_out.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}
func copyFilesWithBufio() {
	// open input file
	fi, err := os.Open("./file/kipling.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create("./file/kipling_b_out.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

func copyFilesWithIOutil() {
	// read the whole file at once
	b, err := ioutil.ReadFile("./file/kipling.txt")
	if err != nil {
		panic(err)
	}

	// write the whole body at once
	err = ioutil.WriteFile("./file/kipling_i_out.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
