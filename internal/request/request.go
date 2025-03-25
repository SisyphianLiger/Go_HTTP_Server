package request

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine 
}

type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method	string
}

type chunkReader struct {
	data            string
	numBytesPerRead int
	pos             int
}

const crlf = "\r\n"

func verifyHttpVersionIsUpper(httpVersion string) (string,error) {
	
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

func parseRequestLine(request []byte) (*RequestLine, error) {
	
	idx := bytes.Index(request, []byte(crlf))
	if idx == -1 {
		return &RequestLine{}, fmt.Errorf("Could not find crlf in request-line")
	}

	parsedData := string(request[:idx])
	requestLines := strings.Split(parsedData, " ")	

	if len(requestLines) != 3 {
		return &RequestLine{}, fmt.Errorf("Invalid Format: Require METHOD TARGET HTTPVERSION, Got: %s\n", requestLines)
	}

	requestTarget := requestLines[1]

	method, methodErr := verifyHttpVersionIsUpper(requestLines[0]);
	if  methodErr != nil {
		return &RequestLine{}, methodErr
	}

	version, versionErr := verifySemanticHTTP(requestLines[2]);
	if  versionErr != nil {
		return &RequestLine{}, versionErr	
	}

	return &RequestLine{
		HttpVersion: version,
		RequestTarget: requestTarget,
		Method: method,
	}, nil 
}


func RequestFromReader(reader io.Reader) (*Request, error) {
	rawBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
		

	reqData, err := parseRequestLine(rawBytes)

	if err != nil {
		return nil, err
	}

	return &Request{ RequestLine: *reqData }, nil
	 
}



// Read reads up to len(p) or numBytesPerRead bytes from the string per call
// its useful for simulating reading a variable number of bytes per chunk from a network connection
func (cr *chunkReader) Read(p []byte) (n int, err error) {
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}
	endIndex := cr.pos + cr.numBytesPerRead
	if endIndex > len(cr.data) {
		endIndex = len(cr.data)
	}
	n = copy(p, cr.data[cr.pos:endIndex])
	cr.pos += n
	if n > cr.numBytesPerRead {
		n = cr.numBytesPerRead
		cr.pos -= n - cr.numBytesPerRead
	}
	return n, nil
}
