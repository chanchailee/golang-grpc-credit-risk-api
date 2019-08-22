package mockrecords

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var riskFactorScoreHeader = []string{
	"customer_id",
	"payment_history_score",
	"amount_owned_score",
	"length_of_credit_history_score",
	"new_credit_score",
	"type_of_credit_score",
}

// randScore : random a score as a float64
// with min,max,size as parameters and returns as an array of string
func randScore(min, max float64, size int) []string {
	score := make([]string, size+1)
	for i := 1; i < len(score); i++ {
		score[i] = fmt.Sprintf("%.2f", math.Round(min+rand.Float64()*(max-min)*100)/100)
	}

	return score
}

// createRecords : this function will create records with a particular size
// and it will return records as map[int][]string
func createRecords(size int) map[int][]string {
	records := make(map[int][]string, size)
	for i := 0; i < size; i++ {
		records[i] = randScore(0, 1, 5)
		records[i][0] = strconv.FormatInt(int64(i), 10)
	}

	return records
}

// Generate function will generate sample of FICO factors as a csv file by using "size"
// and "filename" as parameters.
// 	Parameters:
//		size : is number of the records (integer type).
//		filename : is a record file-name.
func Generate(size int, filename string) error {
	records := createRecords(size)
	csvfile, err := os.Create(filename)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Write(riskFactorScoreHeader)
	for _, record := range records {
		csvwriter.Write(record)
	}

	csvwriter.Flush()
	csvfile.Close()

	return nil
}
