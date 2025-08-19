# Development Setup Guide

This guide helps you set up your development environment for the Go training course.

## Prerequisites

1. **Go Installation**
   ```bash
   # Linux
   wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   ```

2. **VS Code Extensions**
   - Go extension
   - Code Runner
   - Go Test Explorer
   - Marp for VS Code

3. **Additional Tools**
   ```bash
   # Install essential Go tools
   go install golang.org/x/tools/gopls@latest
   go install github.com/fatih/gomodifytags@latest
   go install github.com/cweill/gotests/gotests@latest
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

## Project Setup

1. **Clone Repository**
   ```bash
   git clone <repository-url>
   cd visual-path-training
   ```

2. **VS Code Settings**
   ```json
   {
     "go.useLanguageServer": true,
     "go.testOnSave": true,
     "go.formatTool": "gofmt",
     "[go]": {
       "editor.formatOnSave": true,
       "editor.codeActionsOnSave": {
         "source.organizeImports": true
       }
     }
   }
   ```

3. **Verify Setup**
   ```bash
   go version
   go test ./...
   ```

## Common Issues

1. **GOPATH not set**
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

2. **Missing tools**
   ```bash
   go mod download
   go mod tidy
   ```

3. **Presentation rendering**
   - Install Marp CLI: `npm install -g @marp-team/marp-cli`
   - Test: `marp presentations/week1/day1/Introduction_to_Go.marp.md`

## Testing

Run tests with coverage:
```bash
go test -v -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
