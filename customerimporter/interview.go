// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func CountEmailDomains(filePath string) error {

	file, err := os.Open(filePath)

	if err != nil {

		return err
	}

	defer file.Close()

	domainCounts := make(map[string]int)

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1 // Ustawienie na -1 ignoruje błędy związane z różną liczbą pól na rekord

	for {

		record, err := csvReader.Read()

		if err != nil {

			break
		}

		if len(record) >= 3 {

			email := record[2]
			parts := strings.Split(email, "@")

			if len(parts) == 2 {

				domain := parts[1]
				domainCounts[domain]++
			}
		}
	}

	if len(domainCounts) == 0 {

		return errors.New("no email domains found")
	}

	var domains []string

	for domain := range domainCounts {

		domains = append(domains, domain)
	}

	sort.Strings(domains)

	for _, domain := range domains {

		fmt.Printf("%s: %d\n", domain, domainCounts[domain])
	}
	return nil
}