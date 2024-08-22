package chainquery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

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

func GetSdHash(claimId string) (string, error) {
	url := fmt.Sprintf("https://chainquery.odysee.tv/api/sql?query=select%%20sd_hash%%20from%%20claim%%20where%%20claim_id%%20=%%20%%27%s%%27", claimId)
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
		return "", errors.Err("Failed to get sd hash")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Err(err)
	}

	type cQResponse struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error"`
		Data    []struct {
			SdHash string `json:"sd_hash"`
		} `json:"data"`
	}

	var cQResp cQResponse
	err = json.Unmarshal(body, &cQResp)
	if err != nil {
		return "", errors.Err(err)
	}

	if !cQResp.Success {
		return "", errors.Err("Failed to get sd hash")
	}

	if len(cQResp.Data) == 0 {
		return "", errors.Err("No data returned")
	}

	return cQResp.Data[0].SdHash, nil
}

// channelClaimId returns all unspent claim_ids and sd_hashes of streams belonging to a channel
type stream struct {
	ClaimId string `json:"claim_id"`
	SdHash  string `json:"sd_hash"`
}

func GetChannelStreams(channelClaimId string) ([]stream, error) {
	query := url.QueryEscape(fmt.Sprintf("select claim_id,sd_hash from claim where publisher_id='%s' and bid_state != 'spent' and type = 'stream'", channelClaimId))
	url := fmt.Sprintf("https://chainquery.odysee.tv/api/sql?query=%s", query)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Err(err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Err(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.Err("Failed to get streams")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Err(err)
	}

	type cQResponse struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error"`
		Data    []stream    `json:"data"`
	}
	var cQResp cQResponse
	err = json.Unmarshal(body, &cQResp)
	if err != nil {
		return nil, errors.Err(err)
	}

	if !cQResp.Success {
		return nil, errors.Err("Failed to get sd hash")
	}

	if len(cQResp.Data) == 0 {
		return nil, errors.Err("No data returned")
	}

	return cQResp.Data, nil
}

type thumbnailResponse struct {
	ThumbnailUrl string `json:"thumbnail_url"`
}

func GetClaimThumbnail(claimId string) (string, error) {
	query := url.QueryEscape(fmt.Sprintf("select thumbnail_url from claim where claim_id='%s'", claimId))
	url := fmt.Sprintf("https://chainquery.odysee.tv/api/sql?query=%s", query)
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
		return "", errors.Err("Failed to get thumbnail")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Err(err)
	}

	type cQResponse struct {
		Success bool                `json:"success"`
		Error   interface{}         `json:"error"`
		Data    []thumbnailResponse `json:"data"`
	}
	var cQResp cQResponse
	err = json.Unmarshal(body, &cQResp)
	if err != nil {
		return "", errors.Err(err)
	}

	if !cQResp.Success {
		return "", errors.Err("Failed to get sd hash")
	}

	if len(cQResp.Data) == 0 {
		return "", errors.Err("No data returned")
	}

	return cQResp.Data[0].ThumbnailUrl, nil
}
