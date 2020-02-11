package tcmb
import ( 
	"fmt"
	"encoding/xml"
)

type currency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}

type tarih_Date struct {
	XMLName  xml.Name `xml:"Tarih_Date"`
	Tarih    string   `xml:"Tarih,attr"`
	Date     string   `xml:"Date,attr"`
	BultenNo string   `xml:"Bulten_No,attr"`
	Currency []currency
}
