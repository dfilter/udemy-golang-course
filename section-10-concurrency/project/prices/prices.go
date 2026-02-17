package prices

import (
	"fmt"
	"price_calculator/converstion"
	"price_calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	priceStrings, err := job.IOManager.ReadData()
	if err != nil {
		return err
	}

	job.InputPrices, err = converstion.StringsToFloats(priceStrings)
	if err != nil {
		return err
	}

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errorChannel chan error) {
	err := job.LoadData()
	if err != nil {
		errorChannel <- err
		doneChannel <- true
		return
	}

	pricesMap := make(map[string]string, len(job.InputPrices))
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		pricesMap[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = pricesMap

	errorChannel <- job.IOManager.WriteData(job)
	doneChannel <- true
}

func NewTaxIncludedPriceJob(iomngr iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iomngr,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
