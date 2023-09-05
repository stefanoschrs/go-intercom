# go-intercom

![Intercom API Version](https://img.shields.io/badge/Intercom%20API%20Version-2.9-blue)

> The [official](https://github.com/intercom/intercom-go) client for the [Intercom API](https://developers.intercom.com/docs/references/introduction/) is not maintained anymore and it's a few years behind the latest API.
> 
> This fork is by far not a complete binding to the API, I'm adding methods as I need them. All the examples listed below have been tested with the API version (2.9) and they work as expected.
> 
> Feel free to contribute. 

## Install

`go get gopkg.in/intercom/intercom-go.v2`

## Usage

### Getting a Client

```go
import "github.com/stefanoschrs/go-intercom"

intercomClient := intercom.NewClient("appId", "apiToken")
```

#### Client Options

```go
intercomClient.Option(intercom.BaseURI("https://api.intercom.io"))  // change the base uri used
intercomClient.Option(intercom.ApiVersion("2.9")) // change the api version used
```

### Events

#### Save

```go
event := intercom.Event{
    UserID: "27",
    EventName: "bought_item",
    CreatedAt: int64(time.Now().Unix()),
    Metadata: map[string]interface{}{"item_name": "PocketWatch"},
}
err := ic.Events.Save(&event)
```

- One of `UserID`, `ID`, or `Email` is required (With leads you need to use ID).
- `EventName` is required.
- `CreatedAt` is required, must be an integer representing seconds since Unix Epoch. Will be set to _now_ unless given.
- `Metadata` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.

### Contacts

#### Search

```go
searchParams := intercom.ContactSearchParams{
    Query: map[string]string{
        "field":    "external_id",
        "operator": "=",
        "value":    "xxx",
    },
}
contacts, err := intercomClient.Contacts.Search(searchParams)
```

#### Update

```go
contact := intercom.Contact{
    UserID: "abc-13d-3",
    Name: "SomeContact",
    CustomAttributes: map[string]interface{}{"is_cool": true},
}
savedContact, err := ic.Contacts.Update(&contact)
```

- ID or UserID is required.
- Will not create new contacts.
