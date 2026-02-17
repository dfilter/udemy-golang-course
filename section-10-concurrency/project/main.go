package main

import (
	"fmt"
	"price_calculator/filemanager"
	"price_calculator/prices"
)

func main() {
	taxRates := [4]float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result-%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[index], errorChannels[index])
	}

	for index := range taxRates {
		select {
		case done := <-doneChannels[index]:
			if done {
				fmt.Println("Done!")
			}
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
