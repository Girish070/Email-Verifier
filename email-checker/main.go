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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSpf, sprRecord, hasDmarc, dmrcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read input: %v\n", err)
	}
}

func checkDomain(domain string) {

	var hasMX, hasSpf, hasDmarc bool
	var sprRecord, dmrcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

    txtRecords, err := net.LookupTXT(domain)
    if err != nil {
        log.Printf("Error: %v\n", err)
    }

    for _, record := range txtRecords {
        if strings.HasPrefix(record, "v=spf1"){
            hasSpf = true
            sprRecord = record
            break
        }
    }

	dmarcRecords, err := net.LookupTXT("_dmarc."+ domain)
	if err != nil {
	   log.Printf("Error: %v\n", err) 
	}

		for _, record := range dmarcRecords {
			 if strings.HasPrefix(record, "v=DMARC1") {
				hasDmarc = true
				dmrcRecord = record
				break
			 }
		}

		fmt.Printf("%v, %v, %v,  %v,  %v,  %v", domain, hasMX, hasSpf, sprRecord, hasDmarc, dmrcRecord)
}

