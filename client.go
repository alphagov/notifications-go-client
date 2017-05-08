package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client for accessing GOV.UK Notify.
//
// Before using this client you must have:
//  - created an account with GOV.UK Notify
//  - found your Service ID and generated an API Key.
//  - created at least one template and know its ID.
type Client struct {
	Configuration Configuration
}

/**
 * Internal not exported
 **/

func (c *Client) buildHeaders(r *http.Request) error {
	token, err := c.Configuration.Authenticate(c.Configuration.APIKey)
	if err != nil {
		return err
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-type", "application/json")
	r.Header.Set("User-agent", fmt.Sprintf("NOTIFY-API-GO-CLIENT/%s", Version))

	return nil
}

func (c *Client) handleInvalidResponse(res *http.Response) error {
	if res.StatusCode >= http.StatusBadRequest {
		e := APIError{
			Message:    "api: encountered following errors",
			StatusCode: res.StatusCode,
		}

		err := jsonResponse(res.Body, &e.Errors)
		if err != nil {
			return err
		}

		return &e
	}

	return nil
}

func (c *Client) httpCall(method, url string, payload *[]byte) (*http.Response, error) {
	var body []byte
	if payload != nil {
		body = *payload
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	err = c.buildHeaders(req)
	if err != nil {
		return nil, err
	}

	return c.Configuration.HTTPClient.Do(req)
}

func (c *Client) httpGet(path string, query *Filters) (*http.Response, error) {
	newURL := fmt.Sprintf("%s%s", c.Configuration.BaseURL.String(), path)
	u, err := url.Parse(newURL)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := query.ToURLValues()
		if len(q) > 0 {
			u.RawQuery = q.Encode()
		}
	}

	return c.httpCall("GET", u.String(), nil)
}

func (c *Client) httpPost(path string, payload *Payload) (*http.Response, error) {
	newURL := fmt.Sprintf("%s%s", c.Configuration.BaseURL.String(), path)
	u, err := url.Parse(newURL)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return c.httpCall("POST", u.String(), &body)
}

/**
 * Internal exported
 **/

// GetNotification will fire a request that returns details about the passed
// notification ID.
func (c *Client) GetNotification(id string) (*Notification, error) {
	path := fmt.Sprintf(PathNotificationLookup, id)
	notification := Notification{}

	res, err := c.httpGet(path, nil)
	if err != nil {
		return nil, err
	}

	err = c.handleInvalidResponse(res)
	if err != nil {
		return nil, err
	}

	err = jsonResponse(res.Body, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

// ListNotifications will fire a request that returns a list of all
// notifications for the current Service ID.
func (c *Client) ListNotifications(filters Filters) (*NotificationList, error) {
	path := PathNotificationList
	notificationList := NotificationList{Client: c}

	res, err := c.httpGet(path, &filters)
	if err != nil {
		return nil, err
	}

	err = c.handleInvalidResponse(res)
	if err != nil {
		return nil, err
	}

	err = jsonResponse(res.Body, &notificationList)
	if err != nil {
		return nil, err
	}

	return &notificationList, nil
}

// SendEmail will fire a request to Send an Email message.
func (c *Client) SendEmail(emailAddress, templateID string, personalisation templateData, reference string) (*NotificationEntry, error) {
	payload := NewPayload(
		"email",
		emailAddress,
		templateID,
		personalisation,
		reference,
	)
	apiResponse := NotificationEntry{}

	res, err := c.httpPost(PathNotificationSendEmail, payload)
	if err != nil {
		return nil, err
	}

	err = c.handleInvalidResponse(res)
	if err != nil {
		return nil, err
	}

	err = jsonResponse(res.Body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// SendLetter will fire a request to Send a Letter.
// TODO Establish what this actually is and prepare for usage.
func (c *Client) SendLetter(letter, templateID string, personalisation templateData, reference string) (*NotificationEntry, error) {
	payload := NewPayload(
		"letter",
		letter,
		templateID,
		personalisation,
		reference,
	)
	apiResponse := NotificationEntry{}

	res, err := c.httpPost(PathNotificationSendLetter, payload)
	if err != nil {
		return nil, err
	}

	err = c.handleInvalidResponse(res)
	if err != nil {
		return nil, err
	}

	err = jsonResponse(res.Body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// SendSms will fire a request to Send a SMS message.
func (c *Client) SendSms(phoneNumber, templateID string, personalisation templateData, reference string) (*NotificationEntry, error) {
	payload := NewPayload(
		"sms",
		phoneNumber,
		templateID,
		personalisation,
		reference,
	)
	apiResponse := NotificationEntry{}

	res, err := c.httpPost(PathNotificationSendSms, payload)
	if err != nil {
		return nil, err
	}

	err = c.handleInvalidResponse(res)
	if err != nil {
		return nil, err
	}

	err = jsonResponse(res.Body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

/**
 * External exported
 **/

// New instance of a client will be generated for the use
func New(configuration Configuration) (*Client, error) {
	// Make sure the HTTP client is always set up.
	if configuration.HTTPClient == nil {
		configuration.HTTPClient = &http.Client{}
	}

	// Make sure the BaseURL is always set up.
	if configuration.BaseURL == nil {
		url, err := url.Parse(BaseURLProduction)
		if err != nil {
			return nil, err
		}

		configuration.BaseURL = url
	}

	c := Client{
		Configuration: configuration,
	}

	return &c, nil
}
