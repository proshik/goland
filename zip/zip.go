package main

import (
	"bytes"
	"fmt"
	"runtime"
)

//var accounts = make(map[int32]Account)

var buf bytes.Buffer

//var LikesR = make([]Like, 44000010)

func main() {
	//printMemUsage()
	//var i int32
	//for i = 0; i< 44000000; i++ {
	//	LikesR[i]= Like{Ts: 1500393366, ID: 14773}
	//}
	//printMemUsage()
	//
	//runtime.GC()
	//
	//printMemUsage()

	Run()

	//z, err := zip.OpenReader("/Users/proshik/Work/highloadcup_data/elim_accounts_261218/data/data.zip")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer z.Close()
	//
	//files := map[string]*zip.File{}
	//for _, f := range z.File {
	//	if !strings.Contains(f.Name, "accounts") {
	//		continue
	//	}
	//
	//	files[f.Name] = f
	//}
	//
	//var fileNumber int
	//for _, f := range files {
	//	// miss not account files
	//	if !strings.Contains(f.Name, "accounts") {
	//		continue
	//	}
	//	//fmt.Printf("Being reading f: %s\n", f.Name)
	//
	//	jsonFile, err := f.Open()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	buf.Reset()
	//	_, err = buf.ReadFrom(jsonFile)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	_, err = jsonparser.ArrayEach(buf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	//		err = ParseAccount(value, true)
	//		if err != nil {
	//			return
	//		}
	//		//accounts[a.ID] = *a
	//	}, "accounts")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	printMemUsage()
	//
	//	fileNumber++
	//	fmt.Printf("Readed file with number: %d\n", fileNumber)
	//}
	//
	//buf.Reset()
	//
	//runtime.GC()
	//
	//fmt.Printf("After:\n")
	//printMemUsage()
}

func Run() {

	fmt.Printf("Before:\n")
	printMemUsage()

	//var fileNumber int
	for i := 0; i < 130; i++ {

		ToGo()

		//printMemUsage()

		//fileNumber++
		//fmt.Printf("Readed file with number: %d\n", fileNumber)
	}
	printMemUsage()

	runtime.GC()

	fmt.Printf("After:\n")
	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
