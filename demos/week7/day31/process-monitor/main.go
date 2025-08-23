package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

// ProcessInfo stores information about a monitored process
type ProcessInfo struct {
	PID       int
	Command   string
	StartTime time.Time
	CPU       float64
	Memory    int64
}

// ProcessMonitor manages process monitoring and system resources
type ProcessMonitor struct {
	processes map[int]*ProcessInfo
	mu        sync.RWMutex
	watcher   *fsnotify.Watcher
	done      chan struct{}
}

// NewProcessMonitor creates a new process monitor
func NewProcessMonitor() (*ProcessMonitor, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %v", err)
	}

	return &ProcessMonitor{
		processes: make(map[int]*ProcessInfo),
		watcher:   watcher,
		done:      make(chan struct{}),
	}, nil
}

// StartProcess starts a new process and monitors it
func (pm *ProcessMonitor) StartProcess(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	// Set up process attributes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start process: %v", err)
	}

	// Store process information
	pm.mu.Lock()
	pm.processes[cmd.Process.Pid] = &ProcessInfo{
		PID:       cmd.Process.Pid,
		Command:   command,
		StartTime: time.Now(),
	}
	pm.mu.Unlock()

	// Monitor process in background
	go func() {
		cmd.Wait()
		pm.mu.Lock()
		delete(pm.processes, cmd.Process.Pid)
		pm.mu.Unlock()
	}()

	return nil
}

// MonitorResourceUsage periodically updates resource usage stats
func (pm *ProcessMonitor) MonitorResourceUsage() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pm.mu.Lock()
			for pid, info := range pm.processes {
				// Get process resource usage
				usage := &syscall.Rusage{}
				err := syscall.Getrusage(pid, usage)
				if err != nil {
					continue
				}

				// Update process info
				info.CPU = float64(usage.Utime.Sec) + float64(usage.Utime.Usec)/1e6
				info.Memory = int64(usage.Maxrss)
			}
			pm.mu.Unlock()

		case <-pm.done:
			return
		}
	}
}

// WatchDirectory monitors a directory for changes
func (pm *ProcessMonitor) WatchDirectory(path string) error {
	if err := pm.watcher.Add(path); err != nil {
		return fmt.Errorf("failed to watch directory: %v", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-pm.watcher.Events:
				if !ok {
					return
				}
				log.Printf("File system event: %v", event)

			case err, ok := <-pm.watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Watcher error: %v", err)

			case <-pm.done:
				return
			}
		}
	}()

	return nil
}

// HandleSignals sets up signal handling
func (pm *ProcessMonitor) HandleSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			select {
			case sig := <-sigChan:
				log.Printf("Received signal: %v", sig)

				// Graceful shutdown
				pm.mu.Lock()
				for pid, info := range pm.processes {
					proc, err := os.FindProcess(pid)
					if err != nil {
						continue
					}

					log.Printf("Terminating process %d (%s)", pid, info.Command)
					proc.Signal(syscall.SIGTERM)
				}
				pm.mu.Unlock()

				if sig == syscall.SIGTERM || sig == syscall.SIGINT {
					close(pm.done)
					return
				}

			case <-pm.done:
				return
			}
		}
	}()
}

// PrintStatus prints current process status
func (pm *ProcessMonitor) PrintStatus() {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	fmt.Println("\nProcess Monitor Status:")
	fmt.Println("----------------------")
	for _, info := range pm.processes {
		fmt.Printf("PID: %d\n", info.PID)
		fmt.Printf("Command: %s\n", info.Command)
		fmt.Printf("Running for: %v\n", time.Since(info.StartTime))
		fmt.Printf("CPU Time: %.2f seconds\n", info.CPU)
		fmt.Printf("Memory: %d KB\n", info.Memory)
		fmt.Println("----------------------")
	}
}

func main() {
	monitor, err := NewProcessMonitor()
	if err != nil {
		log.Fatal(err)
	}

	// Start resource monitoring
	go monitor.MonitorResourceUsage()

	// Set up signal handling
	monitor.HandleSignals()

	// Watch current directory
	if err := monitor.WatchDirectory("."); err != nil {
		log.Printf("Failed to watch directory: %v", err)
	}

	// Start some test processes
	if err := monitor.StartProcess("sleep", "100"); err != nil {
		log.Printf("Failed to start sleep: %v", err)
	}

	if err := monitor.StartProcess("yes", ">", "/dev/null"); err != nil {
		log.Printf("Failed to start yes: %v", err)
	}

	// Print status periodically
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			monitor.PrintStatus()
		case <-monitor.done:
			return
		}
	}
}
