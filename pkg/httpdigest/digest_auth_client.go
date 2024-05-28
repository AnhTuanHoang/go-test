package digest_auth_client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DigestRequest struct {
	Body       string
	Method     string
	Password   string
	URI        string
	Username   string
	Header     http.Header
	Auth       *authorization
	Wa         *wwwAuthenticate
	CertVal    bool
	HTTPClient *http.Client
}

type DigestTransport struct {
	Password   string
	Username   string
	HTTPClient *http.Client
}

// NewRequest creates a new DigestRequest object
func NewRequest(username, password, method, uri, body string) DigestRequest {
	dr := DigestRequest{}
	dr.UpdateRequest(username, password, method, uri, body)
	dr.CertVal = true
	return dr
}

func (dr *DigestRequest) getHTTPClient() *http.Client {
	if dr.HTTPClient != nil {
		return dr.HTTPClient
	}
	tlsConfig := tls.Config{}
	timeout := 30 * time.Second
	if !dr.CertVal {
		tlsConfig.InsecureSkipVerify = true
		return &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tlsConfig,
			},
		}
	}

	return &http.Client{
		Timeout: timeout,
	}
}

// UpdateRequest is called when you want to reuse an existing
//  DigestRequest connection with new request information
func (dr *DigestRequest) UpdateRequest(username, password, method, uri, body string) *DigestRequest {
	dr.Body = body
	dr.Method = method
	dr.Password = password
	dr.URI = uri
	dr.Username = username
	dr.Header = make(map[string][]string)
	return dr
}

func (dr *DigestRequest) ExecuteNewDigest(resp *http.Response) (resp2 *http.Response, err error) {
	var (
		auth     *authorization
		wa       *wwwAuthenticate
		waString string
	)

	// body not required for authentication, discarding and closing to reuse connection
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if waString = resp.Header.Get("WWW-Authenticate"); waString == "" {
		return nil, fmt.Errorf("failed to get WWW-Authenticate header, please check your server configuration")
	}
	wa = newWwwAuthenticate(waString)
	dr.Wa = wa

	if auth, err = newAuthorization(dr); err != nil {
		return nil, err
	}

	if resp2, err = dr.executeRequest(auth.toString()); err != nil {
		return nil, err
	}

	dr.Auth = auth
	return resp2, nil
}


func (dr *DigestRequest) executeRequest(authString string) (resp *http.Response, err error) {
	var req *http.Request

	if req, err = http.NewRequest(dr.Method, dr.URI, bytes.NewReader([]byte(dr.Body))); err != nil {
		return nil, err
	}
	req.Header = dr.Header
	req.Header.Add("Authorization", authString)

	client := dr.getHTTPClient()
	return client.Do(req)
}
