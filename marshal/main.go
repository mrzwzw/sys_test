package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type S5F3HsmsData struct {
	Alst      string `xml:"ALST"`
	Alcd      string `xml:"ALCD"`
	AlId      string `xml:"ALID"`
	Altx      string `xml:"ALTX"`
	UnitId    string `xml:"UNITID"`
	ChartId   string `xml:"CHARTID"`
	EqpId     string `xml:"EQPID"`
	RuleType  string `xml:"RULETYPE"`
	GraphType string `xml:"GRAPHTYPE"`
	LotId     string `xml:"LOTID"`
	GlsId     string `xml:"GLSID"`
}

func Marshal() {
	p := S5F3HsmsData{

		Alst:      "1",
		Alcd:      "1",
		AlId:      "1010101010",
		Altx:      "1234567890123456789012345678901234567890",
		UnitId:    "1231234",
		ChartId:   "54321543216789067890",
		EqpId:     "11111111111111111111",
		RuleType:  "22222222222222222222",
		GraphType: "33333333333333333333",
		LotId:     "44444444444444444444",
		GlsId:     "55555555555555555555",
	}
	// b, _ := xml.Marshal(p)
	// 无缩进格式
	//b, _ := xml.Marshal(p)
	// 有缩进格式
	b, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("%v\n", string(b))
}

//	func main() {
//		Marshal()
//	}
func main() {
	u := S5F3HsmsData{

		Alst:      "1",
		Alcd:      "1",
		AlId:      "1010101010",
		Altx:      "1234567890123456789012345678901234567890",
		UnitId:    "1231234",
		ChartId:   "54321543216789067890",
		EqpId:     "11111111111111111111",
		RuleType:  "22222222222222222222",
		GraphType: "33333333333333333333",
		LotId:     "44444444444444444444",
		GlsId:     "55555555555555555555",
	}
	structByReflect(u)

}

func structByReflect(u interface{}) {

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%v\n", v.Field(k).Interface())
	}
	fmt.Printf("%v\n", t.NumField())

}

//<S5F3HsmsData><alSt>1234567890</alSt><alCd>1234512345</alCd><alId>1212121212</alId><alTx>22222222222222</alTx><unitId>33333333333</unitId><chartId>44554455</chartId><eqpId>123123123123</eqpId><ruleType>1234567890123</ruleType><graphType>56785678</graphType><lotId>111111111111111</lotId><glsId>22222222222</glsId></S5F3HsmsData>
