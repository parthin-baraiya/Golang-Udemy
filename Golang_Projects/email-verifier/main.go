package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	for Scanner.Scan() {
		checkDomain(Scanner.Text())
	}

	if err := Scanner.Err(); err != nil {
		log.Printf("Error in input scanner: %v\n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(domain)

	for _, record := range spfRecords {
		if strings.HasPrefix(record, "v=spfi") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("domain:= %v\nhasMX:= %v\nhasSPF:= %v\nspfRecord:= %v\nhasDMARC:= %v\ndmarcRecord:= %v\n\n\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
