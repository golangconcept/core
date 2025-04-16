Alles klar! Du möchtest beim Ausführen deines Codes **automatisch mitverfolgen**, **welche Zeile gerade ausgeführt wird** – mit **Konsolenausgabe auf Deutsch** (gern mit englischer Übersetzung daneben, falls gewünscht).

---

## ✅ **Beispiel in Go – mit Zeilennummer und Nachricht (auf Deutsch)**

Nutze `runtime.Caller`, um Datei und Zeilennummer dynamisch zu ermitteln:

### 🔧 Hilfsfunktion: `logZeile()`

```go
import (
	"fmt"
	"runtime"
)

func logZeile(nachricht string) {
	_, datei, zeile, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("📍 %s (Zeile %d): %s\n", datei, zeile, nachricht)
	} else {
		fmt.Println("📍 [Unbekannte Position]:", nachricht)
	}
}
```

---

## 🧪 Beispielnutzung (z. B. im `main()` oder einem Test):

```go
func main() {
	logZeile("Programmstart") // Output z. B.: 📍 main.go (Zeile 15): Programmstart

	logZeile("Lade Konfiguration")
	config := loadConfig()

	logZeile("Stelle Verbindung zur Datenbank her")
	db, err := SetupMySQL(config)
	if err != nil {
		logZeile("❌ Fehler bei der Datenbankverbindung")
		return
	}
	logZeile("✅ Verbindung erfolgreich")
}
```

---

## 🔁 Mit Deutsch + Englisch kombinieren

Wenn du zweisprachig loggen willst:

```go
func logZeileZwei(deutsch, englisch string) {
	_, datei, zeile, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("📍 %s (Zeile %d): %s / %s\n", datei, zeile, deutsch, englisch)
	} else {
		fmt.Printf("📍 [Unbekannte Position]: %s / %s\n", deutsch, englisch)
	}
}
```

### Beispiel:

```go
logZeileZwei("Lade Benutzerkonfiguration", "Loading user configuration")
```

---

### ✅ Beispielausgabe:
```
📍 main.go (Zeile 21): Lade Benutzerkonfiguration / Loading user configuration
```

---

Möchtest du daraus ein Logging-Paket bauen, das man überall im Projekt verwenden kann? Oder lieber etwas Leichtes für kleine Tools?