package main

import (
	"fmt"
	"github.com/Lucas-Mol/go-studies/mini-exercises/price-calculator/filemanager"
	"github.com/Lucas-Mol/go-studies/mini-exercises/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prasdasdices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process price included")
			fmt.Println(err)
		}
	}
}
