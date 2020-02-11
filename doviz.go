package tcmb
import ( 
	"fmt"
	"encoding/xml"
	"net/http"
	"io/ioutil"
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

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	  return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()
  
	if resp.StatusCode != http.StatusOK {
	  return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}
  
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  return []byte{}, fmt.Errorf("Read body: %v", err)
	}
  
	return data, nil
  }

func GetCurrencies() tarih_Date{
	var result tarih_Date
	if xmlBytes, err := GetXML("https://www.tcmb.gov.tr/kurlar/today.xml"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	  } else {
		xml.Unmarshal(xmlBytes, &result)
		// fmt.Printf("%+v\n",result)
	  }
	return result
}

func (tD *tarih_Date) pretty() string{
	out, err:= xml.MarshalIndent(tD, " ", "  ")
	if err!=nil{
		fmt.Printf("error on pretty: %+v\n",err)
	}
	return string(out)
}


