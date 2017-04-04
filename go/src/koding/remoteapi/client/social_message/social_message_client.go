package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new social message API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for social message API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SocialMessageByID social message by Id API
*/
func (a *Client) SocialMessageByID(params *SocialMessageByIDParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.byId",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.byId",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageByIDOK), nil

}

/*
SocialMessageBySlug social message by slug API
*/
func (a *Client) SocialMessageBySlug(params *SocialMessageBySlugParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageBySlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageBySlugParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.bySlug",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.bySlug",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageBySlugReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageBySlugOK), nil

}

/*
SocialMessageDelete social message delete API
*/
func (a *Client) SocialMessageDelete(params *SocialMessageDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.delete",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.delete",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageDeleteOK), nil

}

/*
SocialMessageEdit social message edit API
*/
func (a *Client) SocialMessageEdit(params *SocialMessageEditParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageEditParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.edit",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.edit",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageEditReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageEditOK), nil

}

/*
SocialMessageFetch social message fetch API
*/
func (a *Client) SocialMessageFetch(params *SocialMessageFetchParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageFetchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageFetchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.fetch",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.fetch",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageFetchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageFetchOK), nil

}

/*
SocialMessageFetchDataFromEmbedly Method SocialMessage.fetchDataFromEmbedly
*/
func (a *Client) SocialMessageFetchDataFromEmbedly(params *SocialMessageFetchDataFromEmbedlyParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageFetchDataFromEmbedlyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageFetchDataFromEmbedlyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.fetchDataFromEmbedly",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.fetchDataFromEmbedly",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageFetchDataFromEmbedlyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageFetchDataFromEmbedlyOK), nil

}

/*
SocialMessageFetchPrivateMessageCount social message fetch private message count API
*/
func (a *Client) SocialMessageFetchPrivateMessageCount(params *SocialMessageFetchPrivateMessageCountParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageFetchPrivateMessageCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageFetchPrivateMessageCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.fetchPrivateMessageCount",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.fetchPrivateMessageCount",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageFetchPrivateMessageCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageFetchPrivateMessageCountOK), nil

}

/*
SocialMessageFetchPrivateMessages social message fetch private messages API
*/
func (a *Client) SocialMessageFetchPrivateMessages(params *SocialMessageFetchPrivateMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageFetchPrivateMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageFetchPrivateMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.fetchPrivateMessages",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.fetchPrivateMessages",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageFetchPrivateMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageFetchPrivateMessagesOK), nil

}

/*
SocialMessageInitPrivateMessage social message init private message API
*/
func (a *Client) SocialMessageInitPrivateMessage(params *SocialMessageInitPrivateMessageParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageInitPrivateMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageInitPrivateMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.initPrivateMessage",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.initPrivateMessage",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageInitPrivateMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageInitPrivateMessageOK), nil

}

/*
SocialMessageLike social message like API
*/
func (a *Client) SocialMessageLike(params *SocialMessageLikeParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageLikeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageLikeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.like",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.like",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageLikeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageLikeOK), nil

}

/*
SocialMessageListLikers social message list likers API
*/
func (a *Client) SocialMessageListLikers(params *SocialMessageListLikersParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageListLikersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageListLikersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.listLikers",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.listLikers",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageListLikersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageListLikersOK), nil

}

/*
SocialMessageListReplies social message list replies API
*/
func (a *Client) SocialMessageListReplies(params *SocialMessageListRepliesParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageListRepliesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageListRepliesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.listReplies",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.listReplies",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageListRepliesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageListRepliesOK), nil

}

/*
SocialMessagePaymentSubscribe social message payment subscribe API
*/
func (a *Client) SocialMessagePaymentSubscribe(params *SocialMessagePaymentSubscribeParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessagePaymentSubscribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessagePaymentSubscribeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.paymentSubscribe",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.paymentSubscribe",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessagePaymentSubscribeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessagePaymentSubscribeOK), nil

}

/*
SocialMessagePost social message post API
*/
func (a *Client) SocialMessagePost(params *SocialMessagePostParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessagePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessagePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.post",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.post",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessagePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessagePostOK), nil

}

/*
SocialMessageReply social message reply API
*/
func (a *Client) SocialMessageReply(params *SocialMessageReplyParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageReplyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageReplyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.reply",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.reply",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageReplyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageReplyOK), nil

}

/*
SocialMessageSearch social message search API
*/
func (a *Client) SocialMessageSearch(params *SocialMessageSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.search",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.search",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageSearchOK), nil

}

/*
SocialMessageSendPrivateMessage social message send private message API
*/
func (a *Client) SocialMessageSendPrivateMessage(params *SocialMessageSendPrivateMessageParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageSendPrivateMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageSendPrivateMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.sendPrivateMessage",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.sendPrivateMessage",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageSendPrivateMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageSendPrivateMessageOK), nil

}

/*
SocialMessageUnlike social message unlike API
*/
func (a *Client) SocialMessageUnlike(params *SocialMessageUnlikeParams, authInfo runtime.ClientAuthInfoWriter) (*SocialMessageUnlikeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialMessageUnlikeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialMessage.unlike",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialMessage.unlike",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialMessageUnlikeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialMessageUnlikeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
