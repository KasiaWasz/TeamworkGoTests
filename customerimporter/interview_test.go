package customerimporter

import (
	"os"
	"testing"
)

func TestCountEmailDomains(t *testing.T) {

	testFilePath := "test_customers.csv"
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(testFilePath)

	csvData := `first_name,last_name,email,gender,ip_address
	Anastazja,Kruk,a.kruk@gmail.com,Female,38.194.51.128 
	Anna,Nowak,anna.nowak@yahoo.com,Female,190.106.124.105
	Robert,Johnson,robertjohnson@wp.pl,Male,170.254.122.46`

	err := os.WriteFile(testFilePath, []byte(csvData), 0644)

	if err != nil {

		t.Fatal(err)
	}

	if err := CountEmailDomains(testFilePath); err != nil {

		t.Errorf("Expected nil error, got: %v", err)
	}
}
