package chainquery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

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

type Bool bool

func (bit *Bool) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.Err("Boolean unmarshal error: invalid input %s", asString)
	}
	return nil
}

type ClaimMetadata struct {
	ID                    uint64    `json:"id"`
	TransactionHashID     *string   `json:"transaction_hash_id,omitempty"`
	Vout                  uint      `json:"vout"`
	Name                  string    `json:"name"`
	ClaimID               string    `json:"claim_id"`
	ClaimType             int8      `json:"claim_type"`
	PublisherID           *string   `json:"publisher_id,omitempty"`
	PublisherSig          *string   `json:"publisher_sig,omitempty"`
	Certificate           *string   `json:"certificate,omitempty"`
	SDHash                *string   `json:"sd_hash,omitempty"`
	TransactionTime       *uint64   `json:"transaction_time,omitempty"`
	Version               *string   `json:"version,omitempty"`
	ValidAtHeight         uint      `json:"valid_at_height"`
	Height                uint      `json:"height"`
	EffectiveAmount       uint64    `json:"effective_amount"`
	Author                *string   `json:"author,omitempty"`
	Description           *string   `json:"description,omitempty"`
	ContentType           *string   `json:"content_type,omitempty"`
	IsNSFW                Bool      `json:"is_nsfw"`
	Language              *string   `json:"language,omitempty"`
	ThumbnailURL          *string   `json:"thumbnail_url,omitempty"`
	Title                 *string   `json:"title,omitempty"`
	Fee                   string    `json:"fee"`
	FeeCurrency           *string   `json:"fee_currency,omitempty"`
	FeeAddress            *string   `json:"fee_address,omitempty"`
	IsFiltered            Bool      `json:"is_filtered"`
	BidState              string    `json:"bid_state"`
	CreatedAt             time.Time `json:"created_at"`
	ModifiedAt            time.Time `json:"modified_at"`
	ClaimAddress          string    `json:"claim_address"`
	IsCertValid           Bool      `json:"is_cert_valid"`
	IsCertProcessed       Bool      `json:"is_cert_processed"`
	License               *string   `json:"license,omitempty"`
	Type                  *string   `json:"type,omitempty"`
	ReleaseTime           *uint64   `json:"release_time,omitempty"`
	SourceHash            *string   `json:"source_hash,omitempty"`
	SourceName            *string   `json:"source_name,omitempty"`
	SourceSize            *uint64   `json:"source_size,omitempty"`
	SourceMediaType       *string   `json:"source_media_type,omitempty"`
	SourceURL             *string   `json:"source_url,omitempty"`
	FrameWidth            *uint64   `json:"frame_width,omitempty"`
	FrameHeight           *uint64   `json:"frame_height,omitempty"`
	Duration              *uint64   `json:"duration,omitempty"`
	AudioDuration         *uint64   `json:"audio_duration,omitempty"`
	Email                 *string   `json:"email,omitempty"`
	ClaimReference        *string   `json:"claim_reference,omitempty"`
	TransactionHashUpdate *string   `json:"transaction_hash_update,omitempty"`
	VoutUpdate            *uint     `json:"vout_update,omitempty"`
	ClaimCount            int64     `json:"claim_count"`
}

func GetClaimMetadata(claimId string) (*ClaimMetadata, error) {
	query := url.QueryEscape(fmt.Sprintf("select * from claim where claim_id='%s'", claimId))
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
		return nil, errors.Err("Failed to get thumbnail")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Err(err)
	}

	type cQResponse struct {
		Success bool            `json:"success"`
		Error   interface{}     `json:"error"`
		Data    []ClaimMetadata `json:"data"`
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

	return &cQResp.Data[0], nil
}
