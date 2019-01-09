package main

import (
	"errors"
	"github.com/buger/jsonparser"
	"strconv"
	"sync"
	"unicode/utf8"
)

type AccountData struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID           int32    `json:"id"`
	Email        []byte   `json:"email"`
	Fname        []byte   `json:"fname"`
	Sname        []byte   `json:"sname"`
	Phone        []byte   `json:"phone"`
	Sex          byte     `json:"sex"`
	Birth        int32    `json:"birth"`
	Country      []byte   `json:"country"`
	City         []byte   `json:"city"`
	Joined       int32    `json:"joined"`
	Status       []byte   `json:"status"`
	Interests    [][]byte `json:"interests"`
	Premium      Premium  `json:"premium"`
	PremiumAdded bool
	Likes        []Like `json:"likes"`
}

type Premium struct {
	Start  int32 `json:"start"`
	Finish int32 `json:"finish"`
}

type Like struct {
	Ts int32 `json:"ts"`
	ID int32 `json:"id"`
}

var (
	bytesPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 4094)
		},
	}
)

//var AccountsRes = make([]Account, 1301000)

var AccountsRes = make([]Account, 1301000)

//var LikeRes = make(map[int32][]Like)

//API server listening at: 127.0.0.1:65430
//Before:
//Alloc = 347 MiB	TotalAlloc = 347 MiB	Sys = 398 MiB	NumGC = 1
//Alloc = 1374 MiB	TotalAlloc = 1517 MiB	Sys = 1487 MiB	NumGC = 4
//After:
//Alloc = 1260 MiB	TotalAlloc = 1517 MiB	Sys = 1489 MiB	NumGC = 5

func ParseAccount(value []byte, fallEmptyValues bool) error {
	// required field
	ID, err := jsonparser.GetInt(value, "id")
	if (err != nil && err != jsonparser.KeyPathNotFoundError) || (err != nil && err == jsonparser.KeyPathNotFoundError && fallEmptyValues) {
		return err
	}

	// required field
	email, _, _, err := jsonparser.Get(value, "email")
	if err != nil && fallEmptyValues {
		return err
	}
	// required field
	sex, _, _, err := jsonparser.Get(value, "sex")
	if err != nil && fallEmptyValues {
		return err
	}

	fname, _, _, _ := jsonparser.Get(value, "fname")
	sname, _, _, _ := jsonparser.Get(value, "sname")
	phone, _, _, _ := jsonparser.Get(value, "phone")
	// required field
	birth, err := jsonparser.GetInt(value, "birth")
	if (err != nil && err != jsonparser.KeyPathNotFoundError) || (err != nil && err == jsonparser.KeyPathNotFoundError && fallEmptyValues) {
		return err
	}
	country, _, _, _ := jsonparser.Get(value, "country")
	city, _, _, _ := jsonparser.Get(value, "city")
	// required field
	joined, err := jsonparser.GetInt(value, "joined")
	if (err != nil && err != jsonparser.KeyPathNotFoundError) || (err != nil && err == jsonparser.KeyPathNotFoundError && fallEmptyValues) {
		return err
	}
	// required field
	status, _, _, err := jsonparser.Get(value, "status")
	if err != nil && fallEmptyValues {
		return err
	}

	// validate params
	if fallEmptyValues && (ID == 0 || len(email) == 0 || len(sex) == 0 || birth == 0 || joined == 0 || len(status) == 0) {
		return errors.New("required field has empty value")
	}

	var interests [][]byte
	_, err = jsonparser.ArrayEach(value, func(bInterest []byte, dtInterests jsonparser.ValueType, oInterests int, err error) {
		if len(bInterest) > 0 {
			interests = append(interests, bInterest)
		}
	}, "interests")
	if err != nil && err != jsonparser.KeyPathNotFoundError && fallEmptyValues {
		return err
	}

	var internalErr bool
	var likes []Like
	_, err = jsonparser.ArrayEach(value, func(bLike []byte, dtLikes jsonparser.ValueType, oLikes int, err error) {
		if len(bLike) > 0 {
			likeTS, err := jsonparser.GetInt(bLike, "ts")
			if err != nil || likeTS == 0 {
				internalErr = true
				return
			}
			likeID, err := jsonparser.GetInt(bLike, "id")
			if err != nil || likeID == 0 {
				internalErr = true
				return
			}
			likes = append(likes, Like{int32(likeTS), int32(likeID)})
		}
	}, "likes")
	if err != nil && err != jsonparser.KeyPathNotFoundError && !fallEmptyValues {
		return err
	}
	if internalErr == true {
		return errors.New("internal objects error")
	}

	var pStart int32
	var pFinish int32
	err = jsonparser.ObjectEach(value, func(kPremium []byte, vPremium []byte, dtPremium jsonparser.ValueType, oPremium int) error {
		if len(kPremium) > 0 && len(vPremium) > 0 {
			pV, err := strconv.Atoi(string(vPremium))
			if err != nil {
				//internalErr = true
				return err
			}
			if string(kPremium) == "start" {
				pStart = int32(pV)
			}
			if string(kPremium) == "finish" {
				pFinish = int32(pV)
			}
		}
		return nil
	}, "premium")
	if err != nil && err != jsonparser.KeyPathNotFoundError && !fallEmptyValues {
		return err
	}

	var PremiumAdded bool
	var premium Premium
	if pStart != 0 && pFinish != 0 {
		premium = Premium{pStart, pFinish}
		PremiumAdded = true
	}

	var sexByte byte
	if len(sex) > 0 {
		sexByte = sex[0]
	} else {
		sexByte = 0
	}

	fID := int32(ID)

	AccountsRes[fID] = Account{ID: fID, Email: email, Fname: Utf8Unescaped(fname), Sname: Utf8Unescaped(sname),
		Phone: phone, Sex: sexByte, Birth: int32(birth), Country: Utf8Unescaped(country), City: Utf8Unescaped(city),
		Joined: int32(joined), Status: Utf8Unescaped(status), Interests: interests, Premium: premium,
		PremiumAdded: PremiumAdded, Likes: likes}

	//return &Account{ID: int32(ID), Email: email, Fname: fname,
	//	Sname: sname, Phone: phone, Sex: sexByte, Birth: int32(birth),
	//	Country: country, City: city, Joined: int32(joined),
	//	Status: status, Interests: interests, Premium: premium, Likes: likes}, nil

	return nil
}

var incID = 0

func ToGo() {
	for i := 0; i <= 10000; i++ {
		incID++

		fID := int32(incID)
		email := []byte("ewheten@icloud.com")
		fname := []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")
		sname := []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")
		phone := []byte("ewheten@icloud.com")
		sexByte := byte('m')
		birth := 1398124800
		country := []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")
		city := []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430 \u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")
		joined := 1398124800
		status := []byte("ewheten@icloud.com")
		interests := [][]byte{[]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"), []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"), []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"),
			[]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"), []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")}
		premium := Premium{Start: 1465481670, Finish: 1465481690}
		PremiumAdded := true
		likes := []Like{
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}}

		//LikeRes[fID] = likes

		AccountsRes[fID] = Account{ID: fID, Email: email, Fname: Utf8Unescaped(fname),
			Sname: Utf8Unescaped(sname), Phone: phone, Sex: sexByte, Birth: int32(birth), Country: Utf8Unescaped(country),
			City: Utf8Unescaped(city), Joined: int32(joined), Status: Utf8Unescaped(status), Interests: interests,
			Premium: premium, PremiumAdded: PremiumAdded, Likes: likes}
	}
}

func ToGo2() {
	for i := 0; i <= 10000; i++ {
		incID++

		var a Account

		a.ID = int32(incID)
		a.Email = []byte("ewheten@icloud.com")
		a.Fname = Utf8Unescaped([]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"))
		a.Sname = Utf8Unescaped([]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"))
		a.Phone = []byte("ewheten@icloud.com")
		a.Sex = byte('m')
		a.Birth = 1398124800
		a.Country = Utf8Unescaped([]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"))
		a.City = Utf8Unescaped([]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"))
		a.Joined = 1398124800
		a.Status = Utf8Unescaped([]byte("ewheten@icloud.com"))
		a.Interests = [][]byte{[]byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430"), []byte("\u0410\u043d\u0436\u0435\u043b\u0438\u043a\u0430")}
		a.Premium = Premium{Start: 1465481670, Finish: 1465481690}
		a.PremiumAdded = true
		a.Likes = []Like{{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773},
			{Ts: 1500393366, ID: 14773}, {Ts: 1500393366, ID: 14773}}

		AccountsRes[a.ID] = a
	}
}

// хак для перевода экранированных строк вида "\u1234\u5678" в нормальный юникод
// выделяет память под ответ
func Utf8Unescaped(input []byte) []byte {
	buf := bytesPool.Get().([]byte)[:0]

	var tmp [4]byte

	i, l := 0, len(input)
	for i < l {
		ch := input[i]

		// \u1234

		// в случае любых ошибок просто пропускаем один байт
		if ch != '\\' {
		} else if (i >= l-4) || input[i+1] != 'u' {
		} else if r, err := strconv.ParseUint(string(input[i+2:i+6]), 16, 64); err != nil {
		} else {
			n := utf8.EncodeRune(tmp[:], rune(r))
			buf = append(buf, tmp[:n]...)
			i += 6
			continue
		}

		buf = append(buf, ch)
		i++
	}

	res := append([]byte{}, buf...)

	bytesPool.Put(buf)

	return res
}
