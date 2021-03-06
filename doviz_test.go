package tcmb

import (
	"testing"
	"fmt"
	"encoding/xml"
)

func Test_GetXML(t *testing.T) {
	if xmlBytes, err := GetXML("https://www.tcmb.gov.tr/kurlar/today.xml"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	  } else {
		var result tarih_Date
		xml.Unmarshal(xmlBytes, &result)
		fmt.Printf("%+v\n",result)
	  }
	// fmt.Printf("%+v\n",GetXML("https://www.tcmb.gov.tr/kurlar/today.xml"))
}

func Test_GetCurrencies(t *testing.T) {
	currencies:=GetCurrencies()
	fmt.Println(currencies.pretty())

	// fmt.Printf("%+v\n",currencies.Currency[0])

}

func Test_GetUSD(t *testing.T) {
	usd:=GetUSD()
	fmt.Println(usd.pretty())
	// fmt.Printf("%+v\n",currencies.Currency[0])

}