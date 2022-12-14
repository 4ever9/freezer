package freezer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type Freezer struct {
	PinataToken string
}

type PinataResponse struct {
	IpfsHash string
}

func New(PinataToken string) *Freezer {
	return &Freezer{PinataToken: PinataToken}
}

func (f *Freezer) PinFile(file io.Reader) (string, error) {
	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part1, err := writer.CreateFormFile("file", filepath.Base("file"))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part1, file)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", f.PinataToken))
	r.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ret := &PinataResponse{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ipfs://%s", ret.IpfsHash), nil
}

func (f *Freezer) PinJson(content map[string]interface{}) (string, error) {
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"

	data, err := json.Marshal(content)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", f.PinataToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ret := &PinataResponse{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ipfs://%s", ret.IpfsHash), nil
}

func (f *Freezer) PinERC1155(content map[string]interface{}, file io.Reader) (string, string, error) {
	assetCid, err := f.PinFile(file)
	if err != nil {
		return "", "", err
	}

	content["image"] = assetCid

	cid, err := f.PinJson(content)
	if err != nil {
		return "", "", err
	}

	return cid, assetCid, err
}
