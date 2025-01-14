package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type HTTPClient interface {
	Get(string, interface{}) ([]byte, error)
	Post(string, interface{}) ([]byte, error)
	Put(string, interface{}) ([]byte, error)
	Patch(string, interface{}) ([]byte, error)
	Delete(string, interface{}) ([]byte, error)
}

type IntercomHTTPClient struct {
	*http.Client

	AppID  string
	APIKey string

	BaseURI       *string
	APIVersion    *string
	ClientVersion *string
	Debug         *bool
}

func NewIntercomHTTPClient(appID, apiKey string, baseURI, apiVersion, clientVersion *string, debug *bool) IntercomHTTPClient {
	return IntercomHTTPClient{
		Client: &http.Client{},

		AppID:  appID,
		APIKey: apiKey,

		BaseURI:       baseURI,
		ClientVersion: clientVersion,
		APIVersion:    apiVersion,
		Debug:         debug,
	}
}

func (c IntercomHTTPClient) UserAgentHeader() string {
	return fmt.Sprintf("go-intercom/%s", *c.ClientVersion)
}

func (c IntercomHTTPClient) Get(url string, queryParams interface{}) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest(http.MethodGet, *c.BaseURI+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgentHeader())
	req.Header.Add("Intercom-Version", *c.APIVersion)
	addQueryParams(req, queryParams)
	if *c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func addQueryParams(req *http.Request, params interface{}) {
	v, _ := query.Values(params)
	req.URL.RawQuery = v.Encode()
}

func (c IntercomHTTPClient) Put(url string, body interface{}) ([]byte, error) {
	return c.postOrPatchOrPut(http.MethodPut, url, body)
}

func (c IntercomHTTPClient) Patch(url string, body interface{}) ([]byte, error) {
	return c.postOrPatchOrPut(http.MethodPatch, url, body)
}

func (c IntercomHTTPClient) Post(url string, body interface{}) ([]byte, error) {
	return c.postOrPatchOrPut(http.MethodPost, url, body)
}

func (c IntercomHTTPClient) postOrPatchOrPut(method, url string, body interface{}) ([]byte, error) {
	// Marshal our body
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return nil, err
	}

	// Setup request
	req, err := http.NewRequest(method, *c.BaseURI+url, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgentHeader())
	req.Header.Add("Intercom-Version", *c.APIVersion)
	if *c.Debug {
		fmt.Printf("%s %s %s\n", req.Method, req.URL, buffer)
	}

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func (c IntercomHTTPClient) Delete(url string, queryParams interface{}) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest("DELETE", *c.BaseURI+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgentHeader())
	req.Header.Add("Intercom-Version", *c.APIVersion)
	addQueryParams(req, queryParams)
	if *c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

type IntercomError interface {
	Error() string
	GetStatusCode() int
	GetCode() string
	GetMessage() string
}

func (c IntercomHTTPClient) parseResponseError(data []byte, statusCode int) IntercomError {
	errorList := HTTPErrorList{}
	err := json.Unmarshal(data, &errorList)
	if err != nil {
		return NewUnknownHTTPError(statusCode)
	}
	if len(errorList.Errors) == 0 {
		return NewUnknownHTTPError(statusCode)
	}
	httpError := errorList.Errors[0]
	httpError.StatusCode = statusCode
	return httpError // only care about the first
}

func (c IntercomHTTPClient) readAll(body io.Reader) ([]byte, error) {
	b, err := io.ReadAll(body)
	if *c.Debug {
		fmt.Println(string(b))
		fmt.Println("")
	}
	return b, err
}
