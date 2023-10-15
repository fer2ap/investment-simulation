package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

func main() {
	var taxes = flag.String("type", "Regressive", "What kind of investment you want: Regressive or tax-free")
	var duration = flag.Int("duration", 1, "For how many months are you going to be invested")
	var interest = flag.Float64("interest", 0.085, "Interest per year of your investment")
	var amount = flag.Float64("amount", 1000.0, "Initial invested amount")
	var multiplier = flag.Float64("mult", 1.0, "Multiplier that modifies the interest, i.e.: you want to compare an invest that pays 120% selic and another that pays 115%, so you set interest as the base interest and use the multiplier to 1.2 and 1.15 to simulate both.")
	flag.Parse()

	effectiveIntereset := *interest * *multiplier
	interestPerMonth := math.Pow((float64(1)+effectiveIntereset), float64(0.0833)) - float64(1)

	var history []float64
	completeHistory := calculateInterest(append(history, *amount), interestPerMonth, 1, *duration)

	totalAmount := completeHistory[len(completeHistory)-1]

	fmt.Println("Interest: ", *interest*100, "%, Multiplier: ", *multiplier*100, "%, Taxes: "+*taxes)
	fmt.Println(payTaxes(totalAmount, *amount, *taxes, *duration))
}

func payTaxes(totalAmount float64, initialAmount float64, taxType string, duration int) float64 {
	if strings.ToLower(taxType) == "regressive" {
		if duration <= 6 {
			return initialAmount + (totalAmount-initialAmount)*(1-0.2250)
		}
		if duration <= 12 {
			return initialAmount + (totalAmount-initialAmount)*(1-0.200)
		}
		if duration <= 24 {
			return initialAmount + (totalAmount-initialAmount)*(1-0.1750)
		}
		return initialAmount + (totalAmount-initialAmount)*(1-0.1500)
	}
	return totalAmount
}

func calculateInterest(history []float64, interestPerMonth float64, actualMonth int, totalMonths int) []float64 {
	if actualMonth > totalMonths {
		return history
	}

	actualAmount := history[len(history)-1]
	return calculateInterest(append(history, actualAmount+actualAmount*interestPerMonth), interestPerMonth, (actualMonth + 1), totalMonths)
}
