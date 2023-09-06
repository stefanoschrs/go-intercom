package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/stefanoschrs/go-intercom/interfaces"
)

// ConversationRepository defines the interface for working with Conversations through the API.
type ConversationRepository interface {
	find(id string, displayType string) (Conversation, error)
	list(params ConversationListParams) (ConversationList, error)
	read(id string) (Conversation, error)
	reply(id string, reply *Reply) (Conversation, error)
	update(conversation *Conversation) (Conversation, error)
}

// ConversationAPI implements ConversationRepository
type ConversationAPI struct {
	httpClient interfaces.HTTPClient
}

type conversationReadRequest struct {
	Read bool `json:"read"`
}

func (api ConversationAPI) list(params ConversationListParams) (ConversationList, error) {
	convoList := ConversationList{}
	data, err := api.httpClient.Get("/conversations", params)
	if err != nil {
		return convoList, err
	}
	err = json.Unmarshal(data, &convoList)
	return convoList, err
}

func (api ConversationAPI) read(id string) (Conversation, error) {
	conversation := Conversation{}
	data, err := api.httpClient.Post(fmt.Sprintf("/conversations/%s", id), conversationReadRequest{Read: true})
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, err
}

func (api ConversationAPI) reply(id string, reply *Reply) (Conversation, error) {
	conversation := Conversation{}
	data, err := api.httpClient.Post(fmt.Sprintf("/conversations/%s/reply", id), reply)
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, nil
}

func (api ConversationAPI) find(id string, displayType string) (Conversation, error) {
	type findParams struct {
		DisplayAs string `url:"display_as"`
	}

	conversation := Conversation{}
	data, err := api.httpClient.Get(fmt.Sprintf("/conversations/%s", id), findParams{
		DisplayAs: displayType,
	})
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, err
}

func (api ConversationAPI) update(conversation *Conversation) (Conversation, error) {
	reqConv := api.buildRequestConversation(conversation)
	return unmarshalToConversation(api.httpClient.Put("/conversations/"+conversation.Id, &reqConv))
}

/**/
// Helpers
/**/

func unmarshalToConversation(data []byte, err error) (Conversation, error) {
	savedConversation := Conversation{}
	if err != nil {
		return savedConversation, err
	}
	err = json.Unmarshal(data, &savedConversation)
	return savedConversation, err
}

func (api ConversationAPI) buildRequestConversation(conversation *Conversation) (reqConv requestConversation) {
	conv, _ := json.Marshal(conversation)
	_ = json.Unmarshal(conv, &reqConv)
	return reqConv
}
