package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
)

const (
	// Protocol version
	Version1 = uint8(1)

	// Command types
	CmdPing    = uint8(1)
	CmdMessage = uint8(2)
	CmdFile    = uint8(3)

	// Maximum payload size
	MaxPayloadSize = 1 << 16
)

// Common errors
var (
	ErrInvalidChecksum = errors.New("invalid checksum")
	ErrInvalidVersion  = errors.New("invalid protocol version")
	ErrPayloadTooLarge = errors.New("payload too large")
	ErrInvalidCommand  = errors.New("invalid command")
)

// Frame represents a protocol frame
type Frame struct {
	Version    uint8
	Command    uint8
	PayloadLen uint16
	Payload    []byte
	Checksum   uint32
}

// NewFrame creates a new protocol frame
func NewFrame(command uint8, payload []byte) (*Frame, error) {
	if len(payload) > MaxPayloadSize {
		return nil, ErrPayloadTooLarge
	}

	if command != CmdPing && command != CmdMessage && command != CmdFile {
		return nil, ErrInvalidCommand
	}

	frame := &Frame{
		Version:    Version1,
		Command:    command,
		PayloadLen: uint16(len(payload)),
		Payload:    payload,
	}

	// Calculate checksum
	frame.Checksum = frame.calculateChecksum()
	return frame, nil
}

// Marshal converts the frame to bytes
func (f *Frame) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write header
	if err := binary.Write(buf, binary.BigEndian, f.Version); err != nil {
		return nil, fmt.Errorf("failed to write version: %v", err)
	}

	if err := binary.Write(buf, binary.BigEndian, f.Command); err != nil {
		return nil, fmt.Errorf("failed to write command: %v", err)
	}

	if err := binary.Write(buf, binary.BigEndian, f.PayloadLen); err != nil {
		return nil, fmt.Errorf("failed to write payload length: %v", err)
	}

	// Write payload
	if _, err := buf.Write(f.Payload); err != nil {
		return nil, fmt.Errorf("failed to write payload: %v", err)
	}

	// Write checksum
	if err := binary.Write(buf, binary.BigEndian, f.Checksum); err != nil {
		return nil, fmt.Errorf("failed to write checksum: %v", err)
	}

	return buf.Bytes(), nil
}

// Unmarshal reads a frame from bytes
func Unmarshal(r io.Reader) (*Frame, error) {
	frame := &Frame{}

	// Read header
	if err := binary.Read(r, binary.BigEndian, &frame.Version); err != nil {
		return nil, fmt.Errorf("failed to read version: %v", err)
	}

	if frame.Version != Version1 {
		return nil, ErrInvalidVersion
	}

	if err := binary.Read(r, binary.BigEndian, &frame.Command); err != nil {
		return nil, fmt.Errorf("failed to read command: %v", err)
	}

	if err := binary.Read(r, binary.BigEndian, &frame.PayloadLen); err != nil {
		return nil, fmt.Errorf("failed to read payload length: %v", err)
	}

	if frame.PayloadLen > MaxPayloadSize {
		return nil, ErrPayloadTooLarge
	}

	// Read payload
	frame.Payload = make([]byte, frame.PayloadLen)
	if _, err := io.ReadFull(r, frame.Payload); err != nil {
		return nil, fmt.Errorf("failed to read payload: %v", err)
	}

	// Read checksum
	if err := binary.Read(r, binary.BigEndian, &frame.Checksum); err != nil {
		return nil, fmt.Errorf("failed to read checksum: %v", err)
	}

	// Verify checksum
	if expected := frame.calculateChecksum(); expected != frame.Checksum {
		return nil, ErrInvalidChecksum
	}

	return frame, nil
}

// calculateChecksum calculates the CRC32 checksum of the frame
func (f *Frame) calculateChecksum() uint32 {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f.Version)
	binary.Write(buf, binary.BigEndian, f.Command)
	binary.Write(buf, binary.BigEndian, f.PayloadLen)
	buf.Write(f.Payload)
	return crc32.ChecksumIEEE(buf.Bytes())
}
