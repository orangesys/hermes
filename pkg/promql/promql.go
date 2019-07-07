package promql

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	Server *url.URL
}

func NewClient(addr string) (*Client, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		Server: u,
	}, nil
}

type QueryRangeResponse struct {
	Status string                  `json:"status"`
	Data   *QueryRangeResponseData `json:"data"`
}

type QueryRangeResponseData struct {
	Result []*QueryRangeResponseResult `json:"result"`
}

type QueryRangeResponseResult struct {
	Metric map[string]string          `json:"metric"`
	Values []*QueryRangeResponseValue `json:"values"`
}

type QueryRangeResponseValue []interface{}

func (v *QueryRangeResponseValue) Time() time.Time {
	t := (*v)[0].(float64)
	return time.Unix(int64(t), 0)
}

func (v *QueryRangeResponseValue) Value() (float64, error) {
	s := (*v)[1].(string)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// QueryRange is count nodes every day, default step 1h with prome unit stripe
// TODO env addr
func (c *Client) QueryRange(query string, start, end int64) (*QueryRangeResponse, error) {
	u, err := url.Parse(fmt.Sprintf("./api/v1/query_range?query=%s&start=%s&end=%s&step=%s",
		url.QueryEscape(query),
		url.QueryEscape(fmt.Sprintf("%d", start)),
		url.QueryEscape(fmt.Sprintf("%d", end)),
		url.QueryEscape(fmt.Sprintf("%dh", int(1))),
	))
	if err != nil {
		return nil, err
	}

	u = c.Server.ResolveReference(u)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if err == io.EOF {
			return &QueryRangeResponse{}, nil
		}
		return nil, err
	}

	if 400 <= res.StatusCode {
		return nil, fmt.Errorf("error response: %s\n", string(body))
	}
	resp := &QueryRangeResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
