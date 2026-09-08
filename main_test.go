package main

import "testing"

// die Einzahlung - поповнення, внесення коштів
// der Betrag - сума
// das Guthaben - баланс, залишок на рахунку
func TestGeldEinzahlen(t *testing.T) {
	tests := []struct {
		name          string
		startGuthaben int
		einzahlBetrag int
		wantGuthaben  int
		wantErr       bool
	}{
		{
			name:          "Gültige Einzahlung",
			startGuthaben: 100,
			einzahlBetrag: 50,
			wantGuthaben:  150,
			wantErr:       false,
		},
		{
			name:          "Ungültiger Betrag (negativ)",
			startGuthaben: 100,
			einzahlBetrag: -50,
			wantGuthaben:  100,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			konto := &Konto{Guthaben: tt.startGuthaben}
			err := konto.GeldEinzahlen(tt.einzahlBetrag)

			if (err != nil) != tt.wantErr {
				t.Errorf("GeldEinzahlen() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}

			if konto.Guthaben != tt.wantGuthaben {
				t.Errorf("GeldEinzahlen() Guthaben = %d, erwartet Guthaben = %d", konto.Guthaben, tt.wantGuthaben)
			}
		})
	}
}

// die Abhebung - зняття грошей
// der Betrag - сума
// das Guthaben - баланс, залишок на рахунку
// nicht genug Geld - недостатньо грошей
func TestGeldAbheben(t *testing.T) {
	tests := []struct {
		name          string
		startGuthaben int
		abhebeBetrag  int
		wantGuthaben  int
		wantErr       bool
	}{
		{
			name:          "Gültige Abhebung",
			startGuthaben: 100,
			abhebeBetrag:  50,
			wantGuthaben:  50,
			wantErr:       false,
		},
		{
			name:          "Ungültiger Betrag (nicht genug Geld)",
			startGuthaben: 100,
			abhebeBetrag:  200,
			wantGuthaben:  100,
			wantErr:       true,
		},
		{
			name:          "Ungültiger Betrag (negativ oder null)",
			startGuthaben: 100,
			abhebeBetrag:  0,
			wantGuthaben:  100,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			konto := &Konto{Guthaben: tt.startGuthaben}
			err := konto.GeldAbheben(tt.abhebeBetrag)

			if (err != nil) != tt.wantErr {
				t.Errorf("GeldAbheben() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}

			if konto.Guthaben != tt.wantGuthaben {
				t.Errorf("GeldAbheben() Guthaben = %d, erwartet Guthaben = %d", konto.Guthaben, tt.wantGuthaben)
			}
		})
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
