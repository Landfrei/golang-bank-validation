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

// gültig - дійсний, коректний, чинний
// ungültig - недійсний, некоректний
// die Länge - довжина
// das Landeskürzel - код країни
// der Fehler - помилка
// erwarten - очікувати (erwartet - очікується)
func TestIstIBANGueltig(t *testing.T) {
	konto := &Konto{}

	tests := []struct {
		name    string
		iban    string
		wantErr bool
	}{
		{
			name:    "Gültige deutsche IBAN",
			iban:    "DE89370400440532013000",
			wantErr: false,
		},
		{
			name:    "Ungültige IBAN-Länge",
			iban:    "DE8937040044053201300",
			wantErr: true,
		},
		{
			name:    "Ungültiges Landeskürzel",
			iban:    "FR89370400440532013000",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := konto.IstIBANGueltig(tt.iban)
			if (err != nil) != tt.wantErr {
				t.Errorf("IstIBANGueltig() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}
		})
	}
}
