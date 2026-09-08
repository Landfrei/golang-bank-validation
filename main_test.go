package main

import "testing"

func TestGeldEinzahlen(t *testing.T) {
	konto := Konto{Guthaben: 100}

	// Перевірка коректного поповнення
	err := konto.GeldEinzahlen(100)
	if err != nil {
		t.Errorf("Очікувалася відсутність помилки, але отримали: %v", err)
	}

	// Перевірка від'ємної суми
	err = konto.GeldEinzahlen(-50)
	if err == nil {
		t.Errorf("Очікувалася помилка для від'ємної суми, але її немає")
	}

	if konto.Guthaben != 200 {
		t.Errorf("Очікувався баланс 200, але отримали: %d", konto.Guthaben)
	}
}

func TestGeldAbheben(t *testing.T) {
	konto := Konto{Guthaben: 100}

	// Перевірка зняття суми в межах балансу
	err := konto.GeldAbheben(50)
	if err != nil {
		t.Errorf("Очікувалася відсутність помилки, але отримали: %v", err)
	}

	// Перевірка спроби зняти більше, ніж є на рахунку
	err = konto.GeldAbheben(10000)
	if err == nil {
		t.Errorf("Очікувалася помилка перевищення балансу, але її немає")
	}
}

func TestIstIBANGueltig(t *testing.T) {
	// Перевірка коректної IBAN
	konto := &Konto{}
	err := konto.IstIBANGueltig("DE89370400440532013000")
	if err != nil {
		t.Errorf("Очікувалася відсутність помилки для коректної IBAN, але отримали: %v", err)
	}

	// Перевірка некоректної IBAN (неправильна довжина)
	err = konto.IstIBANGueltig("DE8937040044053201300")
	if err == nil {
		t.Errorf("Очікувалася помилка для некоректної IBAN (неправильна довжина), але її немає")
	}

	// Перевірка некоректної IBAN (неправильний код країни)
	err = konto.IstIBANGueltig("FR89370400440532013000")
	if err == nil {
		t.Errorf("Очікувалася помилка для некоректної IBAN (неправильний код країни), але її немає")
	}
}
