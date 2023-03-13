package service

import "strings"

func ExtractDMARC(records []string) string {
	for _, value := range records {
		if strings.Contains(value, "v=DMARC1") {
			return value
		}
	}
	return "no DMARC found"
}
