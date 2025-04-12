Got it! Let's build a **custom eBPF firewall in Go** using the excellent [`cilium/ebpf`](https://github.com/cilium/ebpf) library. Go is a solid choice — the eBPF ecosystem for Go is growing fast and it's a clean way to manage programs, maps, and syscalls.

---

## ✅ Overview

We'll write:

1. 🧠 A small **eBPF program** (in C) that drops packets from a hardcoded IP.
2. 🐹 A **Go loader** that:
   - Loads the eBPF bytecode
   - Attaches it to an interface using **XDP**
   - Reads logs from `trace_pipe`

---

## 🧱 Prerequisites

Install:

```bash
sudo apt install clang llvm gcc make iproute2 bpftool
go install github.com/cilium/ebpf/cmd/bpf2go@latest
```

> Note: `bpf2go` will generate Go bindings for the eBPF program.

Make sure `$GOPATH/bin` is in your `$PATH`.

---

## 📁 Project Structure

```
go-ebpf-firewall/
├── main.go
├── xdp_firewall.c
├── Makefile
```

---

## 🔥 Step 1: Write eBPF Program (C - `xdp_firewall.c`)

```c
// +build ignore
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <bpf/bpf_helpers.h>

SEC("xdp")
int xdp_firewall_prog(struct xdp_md *ctx) {
    void *data_end = (void *)(long)ctx->data_end;
    void *data     = (void *)(long)ctx->data;

    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end) return XDP_PASS;

    if (eth->h_proto != __bpf_htons(ETH_P_IP)) return XDP_PASS;

    struct iphdr *ip = data + sizeof(*eth);
    if ((void *)(ip + 1) > data_end) return XDP_PASS;

    // Drop 192.168.1.100
    if (ip->saddr == __bpf_htonl(0xc0a80164)) {
        bpf_printk("Dropping packet from 192.168.1.100\\n");
        return XDP_DROP;
    }

    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
```

---

## ⚙️ Step 2: Generate Go Bindings

```bash
bpf2go -cc clang -cflags "-O2 -g -Wall" XDPFirewall xdp_firewall.c -- -I/usr/include
```

This creates:

- `xdp_firewall.bpf.o` – your compiled eBPF program
- `xdp_firewall_bpfel.go` – Go bindings to your program

---

## 🐹 Step 3: Write Go Loader (`main.go`)

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"golang.org/x/sys/unix"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-O2 -g -Wall" XDPFirewall xdp_firewall.c -- -I/usr/include

func main() {
	iface := "eth0"
	if len(os.Args) > 1 {
		iface = os.Args[1]
	}
	ifIndex, err := netInterfaceIndex(iface)
	if err != nil {
		log.Fatalf("Could not get interface %s: %v", iface, err)
	}

	// Load compiled eBPF program
	objs := XDPFirewallObjects{}
	if err := LoadXDPFirewallObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Attach the program using XDP
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.XdpFirewallProg,
		Interface: ifIndex,
		Flags:     link.XDPGenericMode, // use XDPNativeMode if supported
	})
	if err != nil {
		log.Fatalf("Could not attach program: %v", err)
	}
	defer link.Close()

	fmt.Printf("Firewall running on %s. Press Ctrl+C to exit.\n", iface)

	// Print logs from trace_pipe
	go tracePipeReader()

	// Wait for Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

func netInterfaceIndex(name string) (int, error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return 0, err
	}
	return iface.Index, nil
}

func tracePipeReader() {
	f, err := os.Open("/sys/kernel/debug/tracing/trace_pipe")
	if err != nil {
		log.Printf("Failed to open trace_pipe: %v", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			log.Printf("trace_pipe read error: %v", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Print(string(buf[:n]))
	}
}
```

---

## 🛠 Makefile

```makefile
all:
	bpf2go -cc clang -cflags "-O2 -g -Wall" XDPFirewall xdp_firewall.c -- -I/usr/include
	go build -o firewall
```

---

## ▶️ Run It!

```bash
make
sudo ./firewall eth0
```

You should see:

```
Firewall running on eth0. Press Ctrl+C to exit.
Dropping packet from 192.168.1.100
```

---

## 🚀 What You Just Built

- An **eBPF XDP firewall in C**, compiled and controlled by **Go**
- Drop packets at the NIC level for **192.168.1.100**
- View logs via `trace_pipe`

---

## 🔧 Want to Extend?

- Replace hardcoded IP with **map-based IP list**
- Add userspace **gRPC/REST API** to add/remove IPs
- Export packet drop stats using **eBPF maps** + Go access
- Build a CLI tool to dynamically attach/detach

---

## Need Help With:

- Creating a **map to store blocklisted IPs**?
- Logging dropped packets to Go?
- Making it a full **eBPF-based security tool**?

Let me know how deep you wanna go, and I’ll walk you through the next steps!