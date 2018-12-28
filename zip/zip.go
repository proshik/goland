package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	"runtime"
	"strings"
	"unsafe"
)

var accounts = make([]Account, 1300010)

func main() {
	z, err := zip.OpenReader("./zip/data_huge.zip")
	if err != nil {
		panic(err)
	}

	defer z.Close()

	//res := make(map[string][]byte, 0)
	var buf bytes.Buffer
	//var accounts
	i := 0
	files := map[string]*zip.File{}

	for _, f := range z.File {
		if !strings.Contains(f.Name, "accounts") {
			continue
		}

		files[f.Name] = f
	}

	for _, f := range files {
		// miss not account files
		if !strings.Contains(f.Name, "accounts") {
			continue
		}

		//timeSave := time.Now()
		fmt.Printf("Being reading f: %s\n", f.Name)
		//func() {
		jsonFile, err := f.Open()
		if err != nil {
			panic(err)
		}

		buf.Reset()
		buf.ReadFrom(jsonFile)

		// easyJson

		//var data AccountDataIn
		//data.UnmarshalJSON(buf.Bytes())
		//for i := range data.Account {
		//	accounts[data.Account[i].ID] = data.Account[i]
		//}

		parseErr := false

		_, err = jsonparser.ArrayEach(buf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			//jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			//	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
			//	return nil
			//}, "premium")

			id, err := jsonparser.GetInt(value, "id")
			if err != nil {
				parseErr = true
				return
			}
			email, _, _, err := jsonparser.Get(value, "email")
			if err != nil {
				parseErr = true
				return
			}

			birth, err := jsonparser.GetInt(value, "birth")
			if err != nil {
				parseErr = true
				return
			}

			accounts[id] = Account{ID: int32(id), Email: email, Birth: birth}

		}, "accounts")
		if err != nil {
			panic(err)
		}

		if parseErr {
			fmt.Println("I believe is error")
		}

		jsonFile.Close()

		//byteValue, err := ioutil.ReadAll(jsonFile)
		//if err != nil {
		//	panic(err)
		//}

		//nextAccountSlice := parseAccount(byteValue)

		//for i := range nextAccountSlice {
		//	l.aMw.Add(&nextAccountSlice[i])
		//}

		//res[f.Name] = append(res[f.Name], buf.Bytes()...)

		//accounts = append(accounts, data.Account...)

		//accounts = append(accounts, data.Account...)
		//}()

		i++
		fmt.Printf("number of file: %d\n", i)
		//fmt.Printf("Loaded data from f for time %v\n", time.Since(timeSave))

		//runtime.GC()

		printMemUsage()
	}

	runtime.GC()

	printMemUsage()

	fmt.Printf("len=%d, size=%v", len(accounts), unsafe.Sizeof(accounts))

}

//func parseAccount(byteValue []byte) []model.Account {
//
//	var data model.AccountDataIn
//	err := json.Unmarshal(byteValue, &data)
//	if err != nil {
//		panic(err)
//	}
//
//	return data.Accounts
//}

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
