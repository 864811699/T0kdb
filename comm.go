package kdbtdx

import (
	"time"
)

const (
	EntrustTab = "request"
	//RequestTab    = "requestv3"
	//CancelTab     = "cancelTab"
	HeartBeat   = "HeartBeat"
	ResponseTab = "response"
	//ResponseTabV3 = "responsev3"
	FuncUpdate = "wsupdv2"
)

var (
	entrustCols = []string{"sym", "qid", "accountname", "time", "entrustno", "stockcode", "askprice", "askvol",
		"bidprice", "bidvol", "withdraw", "status", "note"}
	Tb = &tradeKdb{}
)

type basic struct {
	Sym, Qid string
}

type entrust struct {
	basic
	Accountname              string
	Time                     time.Time
	Entrustno                int32
	Stockcode                string
	Askprice                 float64
	Askvol                   int32
	Bidprice                 float64
	Bidvol, Withdraw, Status int32
	Note                     string
}

type Request struct {
	basic
	Time                                                time.Time
	ParentId, Trader, Fund, Strategy, Broker, Stockcode string
	Stocktype, Side                                     int32
	Askprice                                            float64
	Ordertype, Orderqty                                 int32
	Algorithm, Params                                   string
}

type CancelReq struct {
	basic
	Entrustno int32
}

type Order struct {
	Request
	OrderId           string
	SecurityId        string
	EntrustNo, CumQty int32
	AvgPx             float64
	Withdraw, Status  int32
	Note              string
}

type Cfg struct {
	DbPath string   `json:"dbPath"`
	Host   string   `json:"host"`
	Port   int      `json:"port"`
	Auth   string   `json:"auth"`
	Sym    []string `json:"sym"`
	MaxId  int32    `json:"maxId"`
}

type Hold struct {
	Account    string `json:"account"`
	Ticker     string `json:"ticker"`
	Count      int    `json:"count"`
	FrozenQty  int    `json:"frozenQty"`
	InitialQty int    `json:"initialQty"`
}

type queryResult struct {
	Entrustno, Status, Cumqty int32
}
