
package cmd

import (
    "fmt"
    "net/http"
)

func Response(method string, url string, header map[string]string, data map[string]string) (*http.Response, error) {

    //client := &http.Client{}

    // don't follow redirect
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    //req, _ := http.NewRequest("GET", url, nil)
    req, _ := http.NewRequest(method, url, nil)

    //req.Header.Set("Header_Key", "Header_Value")
    for key, value := range header {
        //fmt.Println(key, ":", value)
        req.Header.Set(key, value)
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Client Error making HTTP request: %v", err)
    }

    return res, nil
}

