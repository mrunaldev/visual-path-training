# Process Monitor Demo

This demo implements a system-level process monitor in Go that demonstrates various aspects of system programming.

## Features

1. Process Management
   - Start and monitor processes
   - Track process resources (CPU, Memory)
   - Handle process lifecycle

2. File System Monitoring
   - Watch directory for changes
   - Real-time event notifications
   - Handle file system events

3. Signal Handling
   - Handle system signals (SIGINT, SIGTERM, SIGHUP)
   - Graceful shutdown
   - Process termination

4. Resource Monitoring
   - CPU usage tracking
   - Memory usage tracking
   - Process statistics

## Implementation Details

### Process Management
- Uses `os/exec` for process creation
- Implements process tracking with sync.RWMutex
- Maintains process metadata

### File System Monitoring
- Uses fsnotify for file system events
- Real-time directory watching
- Event logging and handling

### Signal Handling
- Graceful shutdown on SIGTERM/SIGINT
- Process cleanup
- Resource release

### Resource Usage
- Uses syscall.Rusage for resource stats
- Periodic monitoring
- Statistics reporting

## Usage

1. Build the program:
   ```bash
   go build
   ```

2. Run the monitor:
   ```bash
   ./process-monitor
   ```

3. The monitor will:
   - Start sample processes
   - Watch current directory
   - Print process stats every 10 seconds
   - Handle signals for graceful shutdown

## Example Output

```
Process Monitor Status:
----------------------
PID: 1234
Command: sleep
Running for: 5s
CPU Time: 0.02 seconds
Memory: 1024 KB
----------------------
```

## Key Concepts Demonstrated

1. System Calls
   - Process management
   - Resource monitoring
   - File operations

2. Concurrent Programming
   - Goroutines for monitoring
   - Mutex for synchronization
   - Channel-based communication

3. Resource Management
   - Memory tracking
   - CPU usage monitoring
   - File descriptor handling

4. Error Handling
   - Graceful error recovery
   - Resource cleanup
   - Proper shutdown

## Best Practices

1. Always clean up resources
2. Use proper synchronization
3. Handle errors appropriately
4. Implement graceful shutdown
5. Monitor system resources

## Requirements

- Go 1.21 or later
- github.com/fsnotify/fsnotify package
