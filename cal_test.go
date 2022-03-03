package main

import "testing"

func Test_Caltax_Thanes(t *testing.T) {
	name := "ธเนศ"

	giveIncome := 2120000.00
	giveLTF := 300000.00
	giveInsurance := 150000.00

	wantTax := 305000.00

	getTax := caltax(name, giveIncome, giveLTF, giveInsurance)
	if getTax != wantTax {
		t.Errorf("getTax %10.2f but wantTax %10.2f\n", getTax, wantTax)
	}
}

func Test_Caltax_Juraluk(t *testing.T) {
	name := "จุฬาลักษณ์"

	giveIncome := 845000.00
	giveLTF := 0.00
	giveInsurance := 30000.00

	wantTax := 66000.00

	getTax := caltax(name, giveIncome, giveLTF, giveInsurance)
	if getTax != wantTax {
		t.Errorf("getTax %10.2f but wantTax %10.2f\n", getTax, wantTax)
	}
}

func Test_Caltax_Sukkasem(t *testing.T) {
	name := "สุขเกษม"

	giveIncome := 1050000.00
	giveLTF := 25000.00
	giveInsurance := 80000.00

	wantTax := 92000.00

	getTax := caltax(name, giveIncome, giveLTF, giveInsurance)
	if getTax != wantTax {
		t.Errorf("getTax %10.2f but wantTax %10.2f\n", getTax, wantTax)
	}
}

func Test_Caltax_Suporn(t *testing.T) {
	name := "สุพร"

	giveIncome := 5550000.00
	giveLTF := 1200000.00
	giveInsurance := 400000.00

	wantTax := 1.3315e+06

	getTax := caltax(name, giveIncome, giveLTF, giveInsurance)
	if getTax != wantTax {
		t.Errorf("getTax %10.2f but wantTax %10.2f\n", getTax, wantTax)
	}
}
