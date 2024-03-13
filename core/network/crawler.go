package network

import "fmt"

func Crawl(request interface{}) interface{} {
	req, ok := request.(*Request)
	if !ok {
		fmt.Println("Invalid request type")
		return nil
	}

	tmpResponse, err := req.Send()

	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	return tmpResponse
}
