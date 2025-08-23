---
marp: true
theme: default
paginate: true
---

# Advanced System Programming with Go
## Day 31: System-Level Programming and OS Integration

---

# Overview

1. System Calls in Go
2. Process Management
3. IPC (Inter-Process Communication)
4. Signal Handling
5. File System Operations
6. Low-Level System Access

---

# System Calls in Go

```go
// Direct system call example
import "syscall"

func main() {
    // File descriptor operations
    fd, err := syscall.Open("file.txt", 
        syscall.O_RDWR|syscall.O_CREAT, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer syscall.Close(fd)
}
```

---

# Process Management

1. Creating Processes
```go
cmd := exec.Command("ls", "-l")
output, err := cmd.Output()
```

2. Process Control
```go
process, _ := os.FindProcess(pid)
process.Signal(syscall.SIGTERM)
```

3. Process States
- Creation
- Execution
- Termination

---

# Inter-Process Communication (IPC)

1. Pipes
```go
reader, writer := io.Pipe()
```

2. Shared Memory
```go
// Using mmap for shared memory
data, err := syscall.Mmap(
    fd, 0, size,
    syscall.PROT_READ|syscall.PROT_WRITE,
    syscall.MAP_SHARED)
```

3. Unix Domain Sockets

---

# Signal Handling

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan,
    syscall.SIGINT,
    syscall.SIGTERM,
    syscall.SIGHUP)

go func() {
    sig := <-sigChan
    log.Printf("Received signal: %v", sig)
    // Graceful shutdown
}()
```

---

# File System Operations

1. Raw File Operations
```go
fd, err := syscall.Open(
    "file.txt",
    syscall.O_RDWR|syscall.O_CREAT,
    0644)
```

2. File System Events
```go
watcher, _ := fsnotify.NewWatcher()
watcher.Add("/path/to/watch")
```

---

# Low-Level System Access

1. Memory Management
```go
// Memory page size
pageSize := syscall.Getpagesize()

// Memory allocation
addr, err := syscall.Mmap(
    -1, 0, pageSize,
    syscall.PROT_READ|syscall.PROT_WRITE,
    syscall.MAP_PRIVATE|syscall.MAP_ANON)
```

2. System Information
```go
var uname syscall.Utsname
syscall.Uname(&uname)
```

---

# System Resources Management

1. Resource Limits
```go
var rLimit syscall.Rlimit
syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
```

2. Process Priority
```go
syscall.Setpriority(
    syscall.PRIO_PROCESS,
    os.Getpid(),
    -10)
```

---

# Error Handling in System Programming

1. System Error Types
```go
if err == syscall.EAGAIN {
    // Resource temporarily unavailable
}
```

2. Error Recovery Strategies
- Retry mechanisms
- Fallback options
- Graceful degradation

---

# Best Practices

1. Always check error returns
2. Use proper permissions
3. Clean up resources
4. Handle signals appropriately
5. Consider platform differences
6. Test on multiple OS versions

---

# Security Considerations

1. File Permissions
2. Process Privileges
3. Memory Protection
4. System Call Filtering
5. Resource Limitations

---

# Hands-on Exercise

Build a process monitor that:
1. Tracks system resources
2. Handles signals
3. Manages child processes
4. Implements IPC
5. Monitors file system changes

---

# Key Takeaways

1. Go provides extensive system-level access
2. System programming requires careful error handling
3. Resource management is critical
4. Security must be considered
5. Platform differences matter

---

# Next Steps

1. Advanced IPC Patterns
2. Custom System Calls
3. Resource Monitoring
4. System Hardening
5. Performance Optimization
