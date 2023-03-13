package service

import (
	"fmt"
	"net"
)

func GetDnsRecords(domain string) ([]string, error) {
	fmt.Printf("Getting %s DNS records \n", domain)

	dnsTxt, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err
	}

	return dnsTxt, nil
}
