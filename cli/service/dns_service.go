package service

import (
	"fmt"
	"main/service"

	"github.com/spf13/cobra"
)

func DnsService(cmd *cobra.Command, args []string) {
	domain := args[0]
	dns_records, err := service.GetDnsRecords(domain)
	if err != nil {
		fmt.Println(err)
	} else {
		for index, record := range dns_records {
			fmt.Printf("[%d] : %s \n", index, record)
		}
	}
}
