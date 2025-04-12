Perfekt, wenn **alle Dateien im selben Package** sind, ist es sogar noch einfacher – dann brauchst du **keinen Import**, und kannst Funktionen einfach direkt aufrufen. 🎯

---

## 🧱 Beispiel: Alle Dateien im selben Package `main`

### 📂 Struktur:

```
meinprojekt/
├── main.go
└── rechner.go
```

---

### 📄 `main.go`

```go
package main

import "fmt"

func main() {
    ergebnis := Addiere(5, 7)
    fmt.Println("Das Ergebnis ist:", ergebnis)

    Begruessung()
}
```

---

### 📄 `rechner.go`

```go
package main

import "fmt"

func Addiere(a int, b int) int {
    return a + b
}

func Begruessung() {
    fmt.Println("Willkommen in meiner App!")
}
```

> ✅ Wichtig: Beide Dateien haben `package main`

---

### 🛠 Kompilieren / Ausführen

```bash
go run .
```

Oder explizit:

```bash
go run main.go rechner.go
```

---

## 🧠 Warum das funktioniert

- In Go zählt das **Package**, nicht der Dateiname.
- Wenn zwei Dateien im gleichen Ordner sind **und das gleiche Package** (`main`, z. B.), dann **gehören sie zusammen**.
- Der Go Compiler behandelt sie wie **eine große Datei**.

---

## ✅ TL;DR

| Ziel | Was du tun musst |
|------|------------------|
| Funktion aus anderer Datei aufrufen | Einfach direkt aufrufen |
| Voraussetzung | Beide Dateien im gleichen `package` (z. B. `main`) |
| Kein `import` nötig | Solange sie im gleichen Ordner und Package sind |

---

Wenn du willst, kann ich dir gleich ein Beispiel mit mehreren Funktionen oder Structs zeigen – oder sogar mit Tests! Sag einfach, was du brauchst.