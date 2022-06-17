package hl7

import "io"

type Parser interface {
	Parse(r io.Reader) ([]byte, error)
}

func NewParser() Parser {
	return &hl7Parser{}
}

type hl7Parser struct {
}

// Parse implements Parser
func (p hl7Parser) Parse(r io.Reader) ([]byte, error) {
	return nil, nil
}
