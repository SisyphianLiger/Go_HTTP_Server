package request

import (
	"fmt"
	"io"
	"os"
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

func parseIntoRequestLine(request string) (RequestLine, error) {

	parsedData := strings.Split(request, "\r\n")
	requestLines := strings.Split(parsedData[0], " ")	

	if len(requestLines) != 3 {
		return RequestLine{}, fmt.Errorf("Invalid Format: Require METHOD TARGET HTTPVERSION, Got: %s\n", parsedData[0])
	}

	requestTarget := requestLines[1]

	method, methodErr := verifyHttpVersionIsUpper(requestLines[0]);
	if  methodErr != nil {
		return RequestLine{}, methodErr
	}

	version, versionErr := verifySemanticHTTP(requestLines[2]);
	if  versionErr != nil {
		return RequestLine{}, versionErr	
	}

	return RequestLine{
		HttpVersion: version,
		RequestTarget: requestTarget,
		Method: method,
	}, nil 
}


func requestFromReader(reader io.Reader) (*Request, error) {
	req, reqErr := io.ReadAll(reader)	
		
	if reqErr != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Data: %v\n", reqErr)
		os.Exit(1)
	}

	// discard everything after request line
	reqData, err := parseIntoRequestLine(string(req))
	if err != nil {
		return nil, err
	}

	return &Request{ reqData }, nil
	 
}
