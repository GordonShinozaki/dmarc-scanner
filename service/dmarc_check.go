package service

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func DmarcChecker(domain string) (map[string]interface{}, error) {
	dmarcTXT, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		fmt.Println("cannot query dmarc. Error: " + err.Error())
		os.Exit(-1)
	}

	dmarc := ExtractDMARC(dmarcTXT)
	fmt.Printf("\n\nDMARC : %s\n", dmarc)

	dmarcSlice := strings.Fields(dmarc)
	dmarcMap := make(map[string]interface{})
	for _, value := range dmarcSlice {
		tmp := strings.Split(value, "=")
		dmarcMap[tmp[0]] = tmp[1]
	}

	return dmarcMap, nil
}
