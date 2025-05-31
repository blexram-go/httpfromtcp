package request

import (
	"errors"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	r, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	req, err := parseRequestLine(string(r))
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *req,
	}, nil
}

func parseRequestLine(request string) (*RequestLine, error) {
	parts := strings.Split(request, "\r\n")

	requestLineParts := strings.Split(parts[0], " ")
	if len(requestLineParts) != 3 {
		return nil, errors.New("Invalid request line")
	}

	requestLineMethod := requestLineParts[0]
	if !isMethodUpper(requestLineMethod) {
		return nil, errors.New("Invalid method")
	}

	requestLineTarget := requestLineParts[1]
	requestLineVersion := requestLineParts[2]
	httpVersion, err := getHTTPVersion(requestLineVersion)
	if err != nil {
		return nil, err
	}

	return &RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestLineTarget,
		Method:        requestLineMethod,
	}, nil
}

func isMethodUpper(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func getHTTPVersion(s string) (string, error) {
	if s == "" {
		return "", errors.New("No HTTP version provided")
	}

	parts := strings.Split(s, "/")
	if parts[0] == "HTTP" && parts[1] == "1.1" {
		return parts[1], nil
	}
	return "", errors.New("Unsupported HTTP version")
}
