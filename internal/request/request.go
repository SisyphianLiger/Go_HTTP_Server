package request

import (
	"bytes"
	"fmt"
	"io"
	"errors"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine 
	State RequestState
}

type RequestState int

const (
	Initialized = iota
	Done 
)


type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method	string
}


const crlf = "\r\n"
const bufferSize = 8

func verifyMethodIsUpper(httpVersion string) (string,error) {
	
	if httpVersion == "" {
		return "",fmt.Errorf("The String is Empty\n")
	}

	for _,c := range httpVersion {
		if unicode.IsLower(c){
			return "",fmt.Errorf("The HTTP Version is not UPPERCASE: %s, Incorrect HTTP Version Specification\n", httpVersion)
		}
	}
	return httpVersion, nil
}

func verifySemanticHTTP(http string) (string,error) {
	if http != "HTTP/1.1" {
		return "", fmt.Errorf("Not the correct HTTP Version: Found %s", http)
	}

	HTTPVersion := strings.Split(http, "/")

	return HTTPVersion[1], nil
}


func parseRequestLine(data []byte) (*RequestLine, int, error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return nil, 0, nil
	}
	requestLineText := string(data[:idx])
	requestLine, err := requestLineFromString(requestLineText)
	if err != nil {
		return nil, 0, err
	}
	return requestLine, idx + 2, nil
}

func requestLineFromString(str string) (*RequestLine, error) {

	requestLines := strings.Split(str, " ")


	if len(requestLines) != 3 {
		return &RequestLine{}, fmt.Errorf("Invalid Format: Require METHOD TARGET HTTPVERSION, Got: %s with Length: %d\n", requestLines, len(requestLines))
	}


	requestTarget := requestLines[1]

	method, methodErr := verifyMethodIsUpper(requestLines[0]);
	if  methodErr != nil {
		return &RequestLine{}, methodErr
	}

	version, versionErr := verifySemanticHTTP(requestLines[2]);
	if  versionErr != nil {
		return  &RequestLine{}, versionErr	
	}
	

	return &RequestLine{ HttpVersion: version, Method: method, RequestTarget: requestTarget }, nil 
}


func RequestFromReader(reader io.Reader) (*Request, error) {
	
	buf := make([]byte, bufferSize, bufferSize)
	readToIndex := 0

	req := &Request{ State: Initialized}
	

	for req.State != Done {
		if readToIndex >= len(buf) {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf)
			buf = newBuf
		}

		numBytesRead, err := reader.Read(buf[readToIndex:])
		if err != nil {
			if errors.Is(io.EOF, err) {
				req.State = Done
				break
			}
			return nil, err
		}
		readToIndex += numBytesRead

		numBytesParsed, err := req.parse(buf[:readToIndex])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[numBytesParsed:])
		readToIndex -= numBytesParsed
	}
	return req, nil
	 
}

// Accepts Next slice of bytes for request
// updates state to parser in RequestLine
// returns number of bytes consumed (if successful) and error 
func (r *Request) parse(data []byte) (int, error) {
	switch r.State {
		case Initialized:
			requestLine, bytes, err := parseRequestLine(data)
			if err != nil {	
				return 0, err
			}
			if bytes == 0 {
				return 0, nil
			}
			r.RequestLine = *requestLine
			r.State = Done
			return bytes,nil
		case Done:
			return 0, fmt.Errorf("Error: Trying to read data in a done state")
		default:
			return 0, fmt.Errorf("Error: Unknown State")
	}
	
}

