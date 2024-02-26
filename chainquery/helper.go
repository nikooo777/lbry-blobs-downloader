package chainquery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lbryio/lbry.go/v2/extras/errors"
)

// GetOriginalName attempts to get the original name of a stream from the chainquery public API
func GetOriginalName(sdHash string) (string, error) {
	url := fmt.Sprintf("https://chainquery.odysee.tv/api/sql?query=select%%20source_name%%20from%%20claim%%20where%%20sd_hash%%20=%%20%%27%s%%27%%20limit%%201", sdHash)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", errors.Err(err)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", errors.Err(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.Err("Failed to get original name")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Err(err)
	}

	type cQResponse struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error"`
		Data    []struct {
			SourceName string `json:"source_name"`
		} `json:"data"`
	}

	var cQResp cQResponse
	err = json.Unmarshal(body, &cQResp)
	if err != nil {
		return "", errors.Err(err)
	}

	if !cQResp.Success {
		return "", errors.Err("Failed to get original name")
	}

	if len(cQResp.Data) == 0 {
		return "", errors.Err("No data returned")
	}

	return cQResp.Data[0].SourceName, nil
}
