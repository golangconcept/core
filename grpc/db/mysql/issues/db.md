Ah, das sieht nach einem klassischen DNS-Auflösungsproblem aus. Die Fehlermeldung:

```
could not connect to MySQL: %w dial tcp: lookup mysql on 127.0.0.53:53: server misbehaving
```

bedeutet, dass dein Go-Programm versucht, den Hostnamen `mysql` aufzulösen, dieser aber über den lokalen DNS-Resolver `127.0.0.53` (oft systemd-resolved) nicht korrekt aufgelöst werden kann. Hier sind ein paar Lösungsansätze:

---

### ✅ **1. Ist der Hostname korrekt?**
Wenn du in deinem `cfg.Host` den Wert `"mysql"` drin hast, stellt sich die Frage:

- **Läuft der MySQL-Server wirklich unter diesem Namen?**
  - `"mysql"` funktioniert meistens nur in Docker-Netzwerken, wo ein Container `mysql` heißt.
  - In anderen Setups (z. B. lokal oder in der Cloud), musst du evtl. `localhost`, `127.0.0.1`, oder die tatsächliche IP angeben.

➡️ **Fix:** Stelle sicher, dass `cfg.Host` z. B. so gesetzt ist:
```go
cfg.Host = "localhost" // oder "127.0.0.1" oder "192.168.x.x"
```

---

### 🔁 **2. Läuft der DNS-Resolver korrekt?**
Die IP `127.0.0.53` gehört zu `systemd-resolved`. Manchmal kommt es dort zu Problemen.

**Checke auf deinem System:**

```bash
systemctl status systemd-resolved
```

Falls da Fehler sind, könntest du `/etc/resolv.conf` temporär so ändern:

```bash
sudo ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
```

Oder DNS direkt auf z. B. 8.8.8.8 setzen (nur temporär und testweise):

```bash
echo "nameserver 8.8.8.8" | sudo tee /etc/resolv.conf
```

---

### 🐳 **3. Docker im Spiel?**
Falls du Go in einem Container laufen lässt, dann:

- **Ist der MySQL-Container im selben Docker-Netzwerk?**
- **Heißt der MySQL-Container wirklich `mysql`?**

In einer `docker-compose.yml` kannst du dann tatsächlich einfach `mysql` als Hostnamen verwenden. Beispiel:

```yaml
services:
  app:
    build: .
    depends_on:
      - mysql
  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: root
    ports:
      - "3306:3306"
```

Dann funktioniert in Go:
```go
cfg.Host = "mysql"
```

---

### 🛠 Debug-Tipp:
Füge vor dem `sql.Open` mal folgenden Print ein:

```go
fmt.Println("DSN:", dsn)
```

Dann kannst du sehen, was genau als Verbindung versucht wird.

---

Wenn du magst, kannst du mir deinen `cfg`-Inhalt oder das Setup (lokal, Docker, etc.) zeigen, und ich helf dir noch genauer.