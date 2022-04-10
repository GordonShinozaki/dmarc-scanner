package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func ExtractDMARC(records []string) string {
	for _, value := range records {
		if strings.Contains(value, "v=DMARC1") {
			return value
		}
	}
	return "no DMARC found"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <domain to query> \n", os.Args[0])
		os.Exit(0)
	}

	domain := os.Args[1]
	fmt.Printf("Getting %s DNS records \n", domain)

	dnsTXT, err := net.LookupTXT(domain)

	if err != nil {
		fmt.Println("query failed. Error: " + err.Error())
		os.Exit(-1)
	}

	for index, record := range dnsTXT {
		fmt.Printf("[%d] : %s \n", index, record)
	}
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
	fmt.Println("DMARC version : ", dmarcMap["v"])
	fmt.Println("DMARC policy : ", dmarcMap["p"])
	fmt.Println("DMARC subdomain policy : ", dmarcMap["sp"])
	fmt.Println("DMARC bad percentage : ", dmarcMap["pct"])
	fmt.Println("DMARC send reports to : ", dmarcMap["rua"])
}
