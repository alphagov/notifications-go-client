# GOV.UK Notify Go client

## Installing

The Notify Go Client can be installed with
[`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies).

```sh
go get github.com/alphagov/notifications-go-client
```

## Getting started

Create an instance of the Client using:

```go
package main

import "github.com/alphagov/notifications-go-client"

func main() {
	// Configure the client.
	config := notify.Configuration{
		APIKey: "{your api key}",
		ServiceID: "{your service id}",
	}

	// Initialise the client.
	client, err := notify.Client.New(config)
	if err != nil {
		panic(err)
	}
}
```

Generate an API key by logging in to
[GOV.UK Notify](https://www.notifications.service.gov.uk) and going to the
**API integration** page.

## Send messages

### Text message

The method signature is:
```go
SendSms(phoneNumber, templateID string, personalisation templateData, reference string) (*NotificationEntry, error)
```

An example request would look like:

```go
data := map[string]string{
	"name": "Betty Smith",
	"dob": "12 July 1968",
}

response, err := client.SendSms("+447777111222", "df10a23e-2c6d-4ea5-87fb-82e520cbf93a", data, "")
```

<details>
<summary>
Response
</summary>

If the request is successful, `response` will be a `*notify.NotificationEntry`:

```go
type NotificationEntry struct {
	Content   map[string]string
	ID        string
	Reference string
	Template  type Template struct {
		Version int64
		ID      int64
		URI     string
	}
	URI       string
}
```

Otherwise the `notify.APIError` will be returned:
<table>
<thead>
<tr>
<th>`error["status_code"]`</th>
<th>`error["message"]`</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<pre>429</pre>
</td>
<td>
<pre>
[{
	"error": "RateLimitError",
	"message": "Exceeded rate limit for key type TEAM of 10 requests per 10 seconds"
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>429</pre>
</td>
<td>
<pre>
[{
	"error": "TooManyRequestsError",
	"message": "Exceeded send limits (50) for today"
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "BadRequestError",
	"message": "Can"t send to this recipient using a team-only API key"
]}
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "BadRequestError",
	"message": "Can"t send to this recipient when service is in trial mode
	              - see https://www.notifications.service.gov.uk/trial-mode"
}]
</pre>
</td>
</tr>
</tbody>
</table>
</details>

### Email

The method signature is:
```go
SendEmail(emailAddress, templateID string, personalisation templateData, reference string) (*NotificationEntry, error)
```

An example request would look like:

```go
data := map[string]string{
	"name": "Betty Smith",
	"dob": "12 July 1968",
}

response, err := SendEmail("betty@exmple.com", "df10a23e-2c0d-4ea5-87fb-82e520cbf93c", data, "")
```

<details>
<summary>
Response
</summary>

If the request is successful, `response` will be an `*notify.NotificationEntry`:

```go
type NotificationEntry struct {
	Content   map[string]string
	ID        string
	Reference string
	Template  type Template struct {
		Version int64
		ID      int64
		URI     string
	}
	URI       string
}
```

Otherwise the client will raise a ``Alphagov\Notifications\Exception\NotifyException``:
<table>
<thead>
<tr>
<th>`error["status_code"]`</th>
<th>`error["message"]`</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<pre>429</pre>
</td>
<td>
<pre>
[{
	"error": "RateLimitError",
	"message": "Exceeded rate limit for key type TEAM of 10 requests per 10 seconds"
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>429</pre>
</td>
<td>
<pre>
[{
	"error": "TooManyRequestsError",
	"message": "Exceeded send limits (50) for today"
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "BadRequestError",
	"message": "Can"t send to this recipient using a team-only API key"
]}
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "BadRequestError",
	"message": "Can"t send to this recipient when service is in trial mode
	              - see https://www.notifications.service.gov.uk/trial-mode"
}]
</pre>
</td>
</tr>
</tbody>
</table>
</details>


### Arguments


#### `templateId`

Find by clicking **API info** for the template you want to send.

#### `personalisation`

If a template has placeholders you need to provide their values. For example:

```go
personalisation := map[string]string{
	"name": "Betty Smith",
	"dob": "12 July 1968",
}
```

#### `reference`

An optional identifier you generate if you don’t want to use Notify’s `id`. It can be used to identify a single notification or a batch of notifications.

## Get the status of one message

The method signature is:
```go
GetNotification(id string) (*Notification, error)
```

An example request would look like:

```go
notification, err := client.GetNotification("c32e9c89-a423-42d2-85b7-a21cd4486a2a")
```

<details>
<summary>
Response
</summary>

If the request is successful, `notification` will be an `*notify.Notification`:

```go
type Notification struct {
	ID        string
	Body      string
	Subject   string
	Reference string
	Email     string
	Phone     string
	Line1     string
	Line2     string
	Line3     string
	Line4     string
	Line5     string
	Line6     string
	Postcode  string
	Type      string
	Status    string
	Template  type Template struct {
		ID      int64
		URI     string
		Version int64
	}
	CreatedAt time.Time
	SentAt    time.Time
}
```

Otherwise the client will raise a `notify.APIError`:
<table>
<thead>
<tr>
<th>`error["status_code"]`</th>
<th>`error["message"]`</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<pre>404</pre>
</td>
<td>
<pre>
[{
	"error": "NoResultFound",
	"message": "No result found"
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "ValidationError",
	"message": "id is not a valid UUID"
}]
</pre>
</td>
</tr>
</tbody>
</table>
</details>

## Get the status of all messages
The method signature is:
```go
ListNotifications(filters notify.Filters) (*NotificationList, error)
```

An example request would look like:

```go
filters := notify.Filters{
	OlderThan: "c32e9c89-a423-42d2-85b7-a21cd4486a2a",
	Reference: "weekly-reminders",
	Status: "delivered",
	TemplateType: "sms",
}

list, err := client.ListNotifications(filters)
```

<details>
<summary>
Response
</summary>

If the request is successful, `list` will be an `*notify.NotificationList`:

```go
type NotificationList struct {
	Client *Client

	Notifications []Notification
	Links         Pagination
}
```

Otherwise the client will raise a `notify.APIError`:
<table>
<thead>
<tr>
<th>`error["status_code"]`</th>
<th>`error["message"]`</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	'error': 'ValidationError',
	'message': 'bad status is not one of [created, sending, delivered, pending, failed, technical-failure, temporary-failure, permanent-failure]'
}]
</pre>
</td>
</tr>
<tr>
<td>
<pre>400</pre>
</td>
<td>
<pre>
[{
	"error": "ValidationError",
	"message": "Apple is not one of [sms, email, letter]"
}]
</pre>
</td>
</tr>
</tbody>
</table>
</details>

### Notification list pagination

The method signatures are:

```go
Next() error
Previous() error
```

An example request would look like:

```go
var err error
a := list.Notifications[0].ID // Recorded for demonstration only.

err = list.Next()
if err != nil {
	fmt.Printf("Could be that there are no more pages that way. %#v", err)
}

fmt.Println(list.Notifications[0].ID == a) // false

err = list.Previous()
if err != nil {
	fmt.Printf("Could be that there are no more pages that way. %#v", err)
}

fmt.Println(list.Notifications[0].ID == a) // true
```

### Arguments

#### `older_than`

If omitted all messages are returned. Otherwise you can filter to retrieve all notifications older than the given notification `id`.

#### `template_type`

If omitted all messages are returned. Otherwise you can filter by:

* `email`
* `sms`
* `letter`


#### `status`

If omitted all messages are returned. Otherwise you can filter by:

* `sending` - the message is queued to be sent by the provider.
* `delivered` - the message was successfully delivered.
* `failed` - this will return all failure statuses `permanent-failure`, `temporary-failure` and `technical-failure`.
* `permanent-failure` - the provider was unable to deliver message, email or phone number does not exist; remove this recipient from your list.
* `temporary-failure` - the provider was unable to deliver message, email box was full or the phone was turned off; you can try to send the message again.
* `technical-failure` - Notify had a technical failure; you can try to send the message again.

#### `reference`


This is the `reference` you gave at the time of sending the notification. This can be omitted to ignore the filter.


## Development

#### Tests

There are unit and integration tests that can be run to test functionality of
the client.

To run the tests:

```sh
ginkgo
```

Or in the traditional way:

```sh
go test
```

## License

The Notify Go Client is released under the MIT license, a copy of which can be found in [LICENSE](LICENSE.txt).
