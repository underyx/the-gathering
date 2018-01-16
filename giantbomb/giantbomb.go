package giantbomb

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Client struct {
	key string
}

func NewClient(key string) *Client {
	// TODO: check valid key
	return &Client{
		key: key,
	}
}

func (c *Client) Search(name string) (*GameType, error) {
	searchUrl, _ := url.Parse("http://www.giantbomb.com/api/search/")
	opt := SearchRequest{
		Key:       c.key,
		Fields:    "name,original_release_date,platforms,deck",
		Format:    "json",
		Query:     name,
		Resources: "game",
	}
	queryValues, _ := query.Values(opt)
	searchUrl.RawQuery = queryValues.Encode()

	resp, err := http.Get(searchUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload SearchResponse
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Results[0], nil
}

type SearchRequest struct {
	Key       string `url:"api_key"`
	Fields    string `url:"field_list"`
	Format    string `url:"format"`
	Query     string `url:"query"`
	Resources string `url:"resources"`
}

type SearchResponse struct {
	Results []*GameType `json:"results"`
}

type GameType struct {
	Name                string     `json:"name"`
	Deck                string     `json:"deck"`
	OriginalReleaseDate string     `json:"original_release_date"`
	Platforms           []Platform `json:"platforms"`
}

type Platform struct {
	Abbreviation string `json:"abbreviation"`
}
