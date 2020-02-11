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