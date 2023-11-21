package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Asset struct {
	Segment        interface{} `json:"segment"`
	Code           string      `json:"cod"`
	Asset          string      `json:"asset"`
	Type           string      `json:"type"`
	Part           string      `json:"part"`
	PartAccum      interface{} `json:"partAcum"`
	TheoreticalQty string      `json:"theoricalQty"`
}
type Response struct {
	Page    map[string]interface{} `json:"page"`
	Header  map[string]interface{} `json:"header"`
	Results []Asset                `json:"results"`
}

// this function get the data from b3 website and return a list of assets
func GetB3Ifixdata() []Asset {
	url := "https://sistemaswebb3-listados.b3.com.br/indexProxy/indexCall/GetPortfolioDay/eyJsYW5ndWFnZSI6InB0LWJyIiwicGFnZU51bWJlciI6MSwicGFnZVNpemUiOjEyMCwiaW5kZXgiOiJJRklYIiwic2VnbWVudCI6IjEifQ=="

	// Create an HTTP client
	client := &http.Client{}
	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// Send the request and get the response
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	var parsedResponse Response
	err = json.Unmarshal(body, &parsedResponse)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	sort.Slice(parsedResponse.Results, func(i, j int) bool {
		part1 := parsePart(parsedResponse.Results[i].Part)
		part2 := parsePart(parsedResponse.Results[j].Part)
		return part1 > part2
	})
	// Accessing only the "results" array
	return parsedResponse.Results

}
func parsePart(part string) float64 {
	// Remove any non-numeric characters except the dot
	part = strings.ReplaceAll(part, ".", "")
	part = strings.ReplaceAll(part, ",", ".")

	// Convert to float64
	parsedPart, _ := strconv.ParseFloat(part, 64)
	return parsedPart
}
