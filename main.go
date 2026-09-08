package main

import (
	"fmt"
)

// das Guthaneb = баланс
// das Konto = обліковий запис
type Konto struct {
	Guthaben int
}

// die Einzahlung = поповнення грошами
func (k *Konto) GeldEinzahlen(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Einzahlung: %d", betrag)
	}
	k.Guthaben += betrag
	return nil
}

// ungültig = некоректний
// der Betrag = сума
// die Abhebung = зняття грошей
func (k *Konto) GeldAbheben(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Abhebung: %d", betrag)
	}
	if betrag > k.Guthaben {
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %d Geld abzuheben", betrag)
	}
	k.Guthaben -= betrag
	return nil
}

// gültig - дійсний, коректний, чинний
// das Landeskürzel - код країни
func (k *Konto) IstIBANGueltig(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("ungültige IBAN-Länge: %d (muss 22 sein)", len(iban))
	}
	if iban[:2] != "DE" {
		return fmt.Errorf("ungültiges Landeskürzel: %s (muss DE sein)", iban[:2])
	}
	return nil
}

func main() {
	meinBank := Konto{
		Guthaben: 100,
	}
	err1 := meinBank.GeldEinzahlen(-50)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Mein Guthaben ist ", meinBank.Guthaben, "Geld")
	err2 := meinBank.GeldAbheben(200)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("Mein Guthaben ist %d Geld", meinBank.Guthaben)
}
