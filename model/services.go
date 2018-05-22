package model

import (
	"github.com/loukmho/bcaccount_api/config"
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func CalcTaxDoc(taxtype int, taxrate float64, totalamount float64) (beforetaxamount float64, taxamount float64) {
	fmt.Println("taxtype,taxrate,total", taxtype, taxrate, totalamount)

	switch taxtype {
	case 0:
		taxamount = ToFixed(((totalamount*(100+float64(taxrate)))/(100))-totalamount, 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 1:
		taxamount = ToFixed(totalamount-((totalamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 2:
		beforetaxamount = ToFixed(totalamount, 2)
		taxamount = 0
	}

	fmt.Println("Before,Tax", beforetaxamount, taxamount)

	return beforetaxamount, taxamount
}

func CalcTaxItem(taxtype int, taxrate float64, afterdiscountamount float64) (beforetaxamount float64, taxamount float64, totalamount float64) {
	switch taxtype {
	case 0:
		beforetaxamount = ToFixed(afterdiscountamount, config.MoneyPrecision)
		taxamount = ToFixed(((afterdiscountamount*(100+float64(taxrate)))/(100))-afterdiscountamount, config.MoneyPrecision)
		totalamount = ToFixed(beforetaxamount+taxamount, config.MoneyPrecision)
	case 1:
		taxamount = ToFixed(afterdiscountamount-((afterdiscountamount*100)/(100+float64(taxrate))), config.MoneyPrecision)
		beforetaxamount = ToFixed(afterdiscountamount-taxamount, config.MoneyPrecision)
		totalamount = ToFixed(afterdiscountamount, config.MoneyPrecision)
	case 2:
		beforetaxamount = ToFixed(afterdiscountamount, config.MoneyPrecision)
		taxamount = 0
		totalamount = ToFixed(afterdiscountamount, config.MoneyPrecision)
	}

	fmt.Println("taxtype,taxrate,beforetaxamount,taxamount,totalamount", taxtype, taxrate, beforetaxamount, taxamount, totalamount)

	return beforetaxamount, taxamount, totalamount
}


func CalcTaxCredit(taxtype int, taxrate float64, totalamount float64) (beforetaxamount float64, taxamount float64) {
	fmt.Println("taxtype,taxrate,total", taxtype, taxrate, totalamount)

	switch taxtype {
	case 0:
		beforetaxamount = ToFixed(((totalamount * (100)) / (100 + float64(taxrate))), 2)
		taxamount = ToFixed(totalamount-beforetaxamount, 2)
	case 1:
		taxamount = ToFixed(totalamount-((totalamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 2:
		beforetaxamount = ToFixed(totalamount, 2)
		taxamount = 0
	}

	fmt.Println("Before,Tax", beforetaxamount, taxamount)

	return beforetaxamount, taxamount
}