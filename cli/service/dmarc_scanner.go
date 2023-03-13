package service

import (
	"fmt"
	"main/service"

	"github.com/spf13/cobra"
)

func DmarcService(cmd *cobra.Command, args []string) {
	domain := args[0]
	dmarc_records, err := service.DmarcChecker(domain)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DMARC version : ", dmarc_records["v"])
		fmt.Println("DMARC policy : ", dmarc_records["p"])
		fmt.Println("DMARC subdomain policy : ", dmarc_records["sp"])
		fmt.Println("DMARC bad percentage : ", dmarc_records["pct"])
		fmt.Println("DMARC send reports to : ", dmarc_records["rua"])
	}
}
