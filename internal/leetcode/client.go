package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mutasim77/leetcode-api/internal/model"
)

const leetCodeGraphQLEndpoint = "https://leetcode.com/graphql"

type Client struct {
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{},
	}
}

func (c *Client) FetchUserData(username string) (*model.QueryLeetCodeResponse, error) {
	variables := map[string]string{
		"username": username,
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"query":     getUserInfoQuery,
		"variables": variables,
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", leetCodeGraphQLEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var leetCodeResponse model.QueryLeetCodeResponse
	err = json.Unmarshal(body, &leetCodeResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &leetCodeResponse, nil
}
