# Vascular SDK for GO

The Vascular Go package allows using Vascular's APIs with apps written in Go.

## Installing

Use go get to retrieve the SDK to add it to your GOPATH workspace, or project's Go module dependencies.

```
$ go get github.com/vascular/vascular-go
```

To update the SDK use go get -u to retrieve the latest version of the SDK.

```
$ go get -u github.com/vascular/vascular-go
```

#### Example Usage

```
// Import package
import "github.com/vascular/vascular-go"

// initializeApp

apiKey := <API-KEY>
appKey := <APP-KEY>
vascular := vascular.New(*vascular.NewConfig().WithCredentials(&vascular.Credentials{
    ApiKey: &apiKey,
    AppKey: &appKey,
}).WithUserID(<USER-ID>))

// Create user

inboxID, err := vascular.CreateUser(<USER-ID>, "", "")

// Send message to a user

nowInSec := time.Now().Unix()

status, err := vascular.SendMessageToUser(&message.MessageData{
    Title: "Hello, world!",
    Body:  "This a text body.",
    Media: &message.MessageMedia{
        Thumbnail: "http://vascular.io",
        Image:     "http://vascular.io",
    },
    Actions: []*message.MessageAction{
        {
            Name:  "buy",
            Value: "10usd",
        },
    },
    Metadata: "{\n\"meta\": \"on\"\n\n}",
}, &timestamp.Timestamp{Seconds: nowInSec},
)

// Send message to a users

nowInSec := time.Now().Unix()
userIDs := make([]string, 1)
userIDs[0] = <USER-ID>
status, err := vascular.SendMessageToUsers(&message.MessageData{
    Title: "Hello, GO 1xx SDK!",
    Body:  "This a text body.",
    Media: &message.MessageMedia{
        Thumbnail: "http://vascular.io",
        Image:     "http://vascular.io",
    },
    Actions: []*message.MessageAction{
        {
            Name:  "buy",
            Value: "10usd",
        },
    },
    Metadata: "{\n\"meta\": \"on\"\n\n}",
}, &timestamp.Timestamp{Seconds: nowInSec},
		
