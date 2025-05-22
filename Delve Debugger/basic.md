Setting up the **Delve debugger** in a Go project is essential for step-by-step debugging. Delve is the **official debugger for Go**, and it integrates well with **VS Code**, **Goland**, or CLI-based workflows.

---

## ✅ 1. **Install Delve**

### 🖥️ On Linux/macOS:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

---

## 🚀 2. **Use Delve from the Command Line**

### Run your app in debug mode:

```bash
dlv debug
```

This:

* Builds your app
* Starts it with debugging
* Drops into an interactive debug shell

### Common Delve commands:

| Command               | Description                   |
| --------------------- | ----------------------------- |
| `break main.main`     | Set a breakpoint              |
| `break filename:line` | Break at specific file + line |
| `continue` or `c`     | Continue execution            |
| `next` or `n`         | Step to next line             |
| `step` or `s`         | Step into function            |
| `print varName`       | Print variable                |
| `locals`              | Show local variables          |
| `goroutines`          | Show goroutines               |
| `exit`                | Exit debugger                 |

---

## 💻 3. **Debugging in VS Code**

### 🔧 Install:

* Install the **Go extension** by Google.
* Ensure Delve is in your PATH.

### 🛠 `launch.json`:

Create or edit `.vscode/launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug Go App",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}",
      "showLog": true
    }
  ]
}
```

Now you can:

* Set breakpoints in the editor
* Hit **F5** to start debugging
* Inspect variables, call stack, etc.

---

## 🐳 4. **Debug Go inside Docker with Delve**

### Add this to your Dockerfile:

```dockerfile
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient"]
```

### docker-compose:

Expose port `2345`:

```yaml
ports:
  - "2345:2345"
```

### VS Code Remote Debug:

Update `launch.json`:

```json
{
  "name": "Remote Docker Debug",
  "type": "go",
  "request": "launch",
  "mode": "remote",
  "remotePath": "/app",            // Path inside container
  "port": 2345,
  "host": "127.0.0.1",
  "program": "${workspaceFolder}",
  "showLog": true
}
```

---

## 🔐 Security Notes

* **Never expose Delve in production** without securing the `--headless` server with:

  * firewalls
  * `--only-same-user`
  * Docker networking restrictions

---

## ✅ Summary

* `dlv debug` for CLI debugging
* Integrates easily with VS Code
* Supports Docker remote debugging
* Essential for stepping into bugs in live dev

---

Would you like a **working Docker + Delve + VS Code** setup for your Gin + PostgreSQL project? I can generate a full config for you.
