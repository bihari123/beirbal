package structs

import (
	"github.com/spf13/cast"
)

type Insurance struct {
	Age     int     `json:"pii-age"`
	Sex     string  `json:"pii-sex"`
	BMI     float64 `json:"phi-bmi"`
	Region  string  `json:"pii-geo"`
	Charges float64 `json:"pfi-money"`
}

func CreateInsuranceData(csvData [][]string) (insuranceData []Insurance) {
	for row, col := range csvData {
		if row == 0 {
			continue
		}
		insuranceData = append(insuranceData, Insurance{
			Age:     cast.ToInt(col[0]),
			Sex:     col[1],
			BMI:     cast.ToFloat64(col[2]),
			Region:  col[5],
			Charges: cast.ToFloat64(col[6]),
		})
	}
	return
}
