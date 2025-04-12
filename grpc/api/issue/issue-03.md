Gute Frage! 😄  
Warum also:

```bash
go run main.go rechner.go
```

obwohl du auch einfach `go run .` schreiben könntest?

---

## 📌 Erklärung: `go run` braucht zu wissen, **was es ausführen soll**

### ✳️ `go run main.go rechner.go`

Du gibst hier **explizit die Go-Dateien an**, die ausgeführt werden sollen.

Das ist nötig, wenn:

- Du **kein Modul (`go.mod`)** hast
- Oder dein Code ist **nicht im aktuellen Ordner** organisiert wie ein Modul
- Du willst **einzeln bestimmte Dateien** ausführen (nicht den ganzen Ordner)

📦 Wenn du also in einem simplen Projekt bist ohne `go.mod`, sagst du:  
> „Hey Go, hier sind die Dateien, aus denen das Programm bestehen soll.“

---

### ✅ `go run .` (empfohlen mit `go.mod`)

```bash
go run .
```

Das bedeutet:
> „Führe das Go-Programm im aktuellen Ordner aus.“

Das funktioniert **nur**, wenn du ein **Go-Modul** hast (`go.mod`), weil Go dann weiß:
- Welche Datei `package main` enthält
- Welche Dateien dazugehören (alle im gleichen Ordner und Package)

---

## 🛠 Beispiele

### 🟧 Ohne `go.mod`:

Du musst alle Dateien angeben:

```bash
go run main.go rechner.go
```

### 🟩 Mit `go.mod`:

Dann geht’s auch bequem:

```bash
go mod init meinprojekt
go run .
```

---

## ✅ TL;DR

| Befehl | Wann verwenden |
|--------|----------------|
| `go run main.go rechner.go` | Wenn du **kein Modul** hast oder nur bestimmte Dateien ausführen willst |
| `go run .` | Wenn du ein **Go-Modul** hast (`go.mod`) im Ordner |

---

Wenn du magst, kann ich dir zeigen, wie du in 5 Sekunden ein Modul einrichtest und nie wieder Dateinamen tippen musst 😄