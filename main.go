
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "github.com/karlrink/gurl/http"
)

func main() {

    url := os.Args[1]

    stdin := bufio.NewReader(os.Stdin)

    requestBody, err := ioutil.ReadAll(stdin)
    if err != nil {
        fmt.Printf("Failed to read request body: %v \n", err)
        os.Exit(1)
    }


    header := map[string]string{
        "Content-Type": "application/json",
	}

    response, err := http.Request("POST", url, header, requestBody)
    if err != nil {
        fmt.Printf("Failed HTTP: %v \n", err)
        os.Exit(1)
    }

    for name, values := range response.Header {
    // Loop over all values for the name.
        for _, value := range values {
            fmt.Println(name, value)
        }
    }


    content, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    fmt.Printf("Read body: %v", err)
    }

    responseString := string(content)

    fmt.Printf("%s", responseString)

    fmt.Println(response.StatusCode)


}
