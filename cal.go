package main

import "fmt"

const expense float64 = 60000.00
const maxltf float64 = 200000.00
const maxinsurance float64 = 100000.00

func main() {
	fmt.Println(caltax("ธเนศ", 2120000, 300000, 150000))
	fmt.Println(caltax("จุฬาลักษณ์", 845000, 0, 30000))
	fmt.Println(caltax("สุขเกษม", 1050000, 25000, 80000))
	fmt.Println(caltax("สุพร", 5550000, 1200000, 400000))
}

// func caltax(name string, income float64, ltf float64, insurance float64) string {
func caltax(name string, income float64, ltf float64, insurance float64) float64 {
	var tax float64

	if maxltf < ltf {
		ltf = maxltf
	}
	if maxinsurance < insurance {
		insurance = maxinsurance
	}

	net := income - expense - ltf - insurance

	if net <= 150000 {
		tax = net * 0.05
	} else if net <= 300000 {
		tax = (net-150000)*0.05 + 7500
	} else if net <= 500000 {
		tax = (net-300000)*0.10 + 7500
	} else if net <= 750000 {
		tax = (net-500000)*0.15 + 27500
	} else if net <= 1000000 {
		tax = (net-750000)*0.20 + 65000
	} else if net <= 2000000 {
		tax = (net-1000000)*0.25 + 115000
	} else if net <= 5000000 {
		tax = (net-2000000)*0.30 + 365000
	} else {
		tax = (net-5000000)*0.35 + 1265000
	}
	// return fmt.Sprintf("%s เงินภาษีที่ต้องจ่ายคือ %10.2f \n", name, tax)

	return tax

}
