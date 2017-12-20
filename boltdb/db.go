package main

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"time"
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
)

const (
	WORDS_BUCKET = "words"
)

type RawWord struct {
	Text        string `json:"text"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
}

type Word struct {
	Text        string `json:"text"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Translate   []Def  `json:"translate"`
}

type Translate struct {
	Head Head  `json:"head"`
	Def  []Def `json:"def"`
}

type Head struct{}

type Def struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Ts   string `json:"ts"`
	Tr   []Tr   `json:"tr"`
}

type Tr struct {
	Attr
	Syn  []Syn  `json:"syn"`
	Mean []Mean `json:"mean"`
	Ex   []Ex   `json:"ex"`
}

type Syn struct {
	Attr
}

type Mean struct {
	Attr
}

type Ex struct {
	Attr
	Tr
}

type Attr struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Gen  string `json:"gen"`
}


type DBConnect struct {
	path string
}

func NewDB(path string) *DBConnect {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(WORDS_BUCKET))
		return err
	})
	if err != nil {
		panic(err)
	}

	return &DBConnect{path}
}

func (c *DBConnect) AddWord(word Word) (*Word, error) {
	db, err := open(c)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(WORDS_BUCKET))

		// Marshal user data into bytes.
		buf, err := json.Marshal(&word)
		if err != nil {
			return err
		}

		// Persist bytes to environment bucket.
		return b.Put([]byte(word.Text), buf)
	})
	if err != nil {
		log.Printf("Error on add word to DB. %s\n", err)
		return nil, err
	}

	return &word, nil

}

func (c *DBConnect) CountWords() (int, error) {
	db, err := open(c)
	if err != nil {
		return 0, err
	}

	defer db.Close()

	var count int
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(WORDS_BUCKET))

		count = b.Stats().KeyN

		return nil
	})
	if err != nil {
		log.Printf("Error on get count words in DB. %s\n", err)
		return 0, err
	}

	return count, nil
}

func (c *DBConnect) GetWords(text string) (*Word, error) {
	db, err := open(c)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var data []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(WORDS_BUCKET))
		data = b.Get([]byte(text))
		return nil
	})
	if err != nil {
		return nil, err
	}
	//if not found data by key
	if len(data) == 0 {
		return nil, nil
	}
	//parse byte array
	var env Word
	err = json.Unmarshal(data, &env)
	if err != nil {
		log.Printf("Error on get word from DB. %s\n", err)
		return nil, err
	}

	return &env, nil
}

func open(c *DBConnect) (*bolt.DB, error) {
	db, err := bolt.Open(c.path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Printf("Error on open connections with DB. %s\n", err)
		return nil, err
	}

	return db, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}


func fillBasicEnglishWords(yandex *YDict, db *DBConnect) {
	file, err := os.Open("result.json")
	if err != nil {
		panic(err)
	}

	var rawWords = make([]RawWord, 0)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&rawWords)
	if err != nil {
		panic(err)
	}

	for _, rw := range rawWords {
		tr, err := yandex.translate(rw.Text, "en", "ru")
		if err != nil {
			log.Fatalf("Error on word=%s, with error=%v", rw.Text, err)
		}

		word, err := db.AddWord(Word{rw.Text, rw.Category, rw.Subcategory, tr.Def})
		if err != nil {
			panic(err)
		}

		fmt.Printf("Success translate and save word=%s\n", word.Text)
	}
}

func readRawWords() {
	fInfos, err := ioutil.ReadDir("words")
	if err != nil {
		panic(err)
	}

	var count int
	rawWords := make([]RawWord, 0)
	for _, fInfo := range fInfos {

		words := func(fInfo os.FileInfo) []RawWord {
			result := make([]RawWord, 0)

			file, err := os.Open("words" + "/" + fInfo.Name())
			if err != nil {
				panic(err)
			}

			defer file.Close()

			s := bufio.NewScanner(file)

			for s.Scan() {
				elem := strings.Split(s.Text(), " - ")

				subcategoryTitle := strings.TrimSuffix(fInfo.Name(), ".txt")

				result = append(result, RawWord{elem[0], "Basic English words", subcategoryTitle})
			}

			return result
		}(fInfo)

		for _, w := range words {
			rawWords = append(rawWords, w)
			count++
		}
	}

	b, err := json.MarshalIndent(&rawWords, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("result.json", b, 06444)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result connt %d\n", count)
}

