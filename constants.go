package notify

// Version of this client.
// This follows Semantic Versioning (http://semver.org/)
const Version = "1.0.0"

// BaseURLProduction is the The API endpoint for Notify production.
const BaseURLProduction = "https://api.notifications.service.gov.uk"

// PathNotificationList directs to the appropriate endpoint responsible for
// retrieving the list of notifications.
const PathNotificationList = "/v2/notifications"

// PathNotificationLookup directs to the appropriate endpoint responsible for
// lookup of any notifications.
const PathNotificationLookup = "/v2/notifications/%s"

// PathNotificationSendEmail directs to the appropriate endpoint responsible for
// sending an email message.
const PathNotificationSendEmail = "/v2/notifications/email"

// PathNotificationSendLetter directs to the appropriate endpoint responsible
// for sending a letter.
const PathNotificationSendLetter = "/v2/notifications/letter"

// PathNotificationSendSms directs to the appropriate endpoint responsible for
// sending a text message.
const PathNotificationSendSms = "/v2/notifications/sms"
