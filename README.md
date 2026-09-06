# German Bank Account Validation Service in Go

Ein einfaches Go-Programm zur Validierung von Bankkonto-Transaktionen (Einzahlungen und Abhebungen) mit deutscher Fehlermeldungslogik.

---

## 📌 Domain Vocabulary (Fachbegriffe)

| Begriff (Term) | Übersetzung (Translation) |
| :--- | :--- |
| **das Konto** | Bank Account |
| **das Guthaben** | Balance |
| **der Betrag** | Amount |
| **die Einzahlung** | Deposit |
| **die Abhebung** | Withdrawal |
| **ungültig** | Invalid |

---

## 🚀 Features & Validation Rules

- **`GeldEinzahlen(betrag int)`**: Validates that deposit amounts are greater than zero (`betrag > 0`).
- **`GeldAbheben(betrag int)`**: Validates that withdrawal amounts are greater than zero and do not exceed the current balance (`betrag <= Guthaben`).
- **Idiomatic Error Handling**: Returns descriptive German error messages using Go's standard `fmt.Errorf`.

---

## 🛠️ Usage

```bash
go run main.go