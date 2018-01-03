package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {

	icao, err := parseInput(req)
	if err != nil {
		return err.Error()
	}
	ep := fmt.Sprintf("http://avwx.rest/api/metar/%s?options=info,speech", icao)
	resp, err := http.Get(ep)
	if err != nil {
		return fmt.Sprintf("Error when calling endpoint %s. Error: %s", ep, err)
	}

	// check our status code is 200
	if resp.StatusCode != http.StatusOK {
		// something bad happened, try to parse the response
		// body and echo that back to the user along with the
		// status code. Ignore the error if we can't get the body
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Sprintf("Unexpected status code, %v, response body: %s, endpoint: %s", resp.StatusCode, b, ep)
	}

	var r MetarResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return fmt.Sprintf("Error marshalling response body, %s.", err)
	}

	if r.Error != "" {
		return fmt.Sprintf("Error received when calling endpoint: %s", r.Error)
	}
	return fmt.Sprintf("Latest METAR for %s. %s", r.Info.Name, r.Speech)
}

func parseInput(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("Oops, no weather station was provided, ensure you provide a station using the Phonetic Alphabet, ie Echo Golf Alpha Charlie")
	}
	sIn := strings.Split(string(input), " ")
	var icao []byte
	for _, v := range sIn {
		icao = append(icao, v[0])
	}
	return icao, nil
}

// MetarResponse is the root level object returned
// by the wx api. There are many more properties
// available but we don't them for this function
type MetarResponse struct {
	Speech string
	Info   Info
	Error  string
}

// Info contains the name of the current weather station
// we can echo this back to the user so they have confidence
// they are receiving the correct station
type Info struct {
	Name string
}
