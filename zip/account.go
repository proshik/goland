package main

type AccountDataIn struct {
	Account []Account `json:"accounts"`
}

type Account struct {
	ID            int32    `json:"id"`
	Email         []byte   `json:"email"`
	Fname         []byte   `json:"fname"`
	Sname         []byte   `json:"sname"`
	Phone         []byte   `json:"phone"`
	Sex           []byte   `json:"sex"`
	Birth         int64    `json:"birth"`
	Country       []byte   `json:"country"`
	City          []byte   `json:"city"`
	Joined        int64    `json:"joined"`
	Status        []byte   `json:"status"`
	Interests     [][]byte `json:"interests"`
	Premium       Premium  `json:"premium"`
	PremiumSetted bool

	Likes []Like `json:"likes"`
}

type Premium struct {
	Start  int64 `json:"start"`
	Finish int64 `json:"finish"`
}

type Like struct {
	Ts int64 `json:"ts"`
	ID int32 `json:"id"`
}
