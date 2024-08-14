// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import "context"

// Dispatcher is the full API for a dispatcher that facilitates the publishing
// of async events and the sending and receiving of async messages.
type Dispatcher interface {
	EventDispatcher
	MessageDispatcher
	// Start starts the dispatcher.
	Start(ctx context.Context) error
	// RegisterPublishers registers publishers to the dispatcher.
	RegisterPublishers(publishers ...Publisher) error
	// RegisterRoutes registers message routes to the dispatcher.
	RegisterRoutes(routes ...MessageRoute) error
	// Name returns the name of the dispatcher.
	Name() string
}

// EventDispatcher is the API for a dispatcher that facilitates the publishing
// of async events.
type EventDispatcher interface {
	// PublishEvent publishes an event to the dispatcher.
	PublishEvent(event BaseMessage) error
	// Subscribe subscribes the given channel to all events with the given event
	// ID.
	// Contract: the channel must be a Subscription[T], where T is the expected
	// type of the event data.
	Subscribe(eventID EventID, ch any) error
}

// MessageDispatcher is the API for a dispatcher that facilitates the sending
// and receiving of async messages.
type MessageDispatcher interface {
	// SendRequest sends a request to the dispatcher.
	SendRequest(req BaseMessage, future any) error
	// SendResponse sends a response to the dispatcher.
	SendResponse(resp BaseMessage) error
	// RegisterMsgReceiver registers the given channel as the message receiver
	// for the given message ID.
	RegisterMsgReceiver(messageID MessageID, ch any) error
}

// publisher is the interface that supports basic event publisher operations.
type Publisher interface {
	// Start starts the event publisher.
	Start(ctx context.Context)
	// Publish publishes the given event to the event publisher.
	Publish(event BaseMessage) error
	// Subscribe subscribes the given channel to the event publisher.
	Subscribe(ch any) error
	// Unsubscribe unsubscribes the given channel from the event publisher.
	Unsubscribe(ch any) error
	// EventID returns the event ID that the publisher is responsible for.
	EventID() EventID
}

// messageRoute is the interface that supports basic message route operations.
type MessageRoute interface {
	// RegisterRecipient sets the recipient for the route.
	RegisterReceiver(ch any) error
	// SendRequest sends a request to the recipient.
	SendRequest(msg BaseMessage, future any) error
	// SendResponse sends a response to the recipient.
	SendResponse(msg BaseMessage) error
	// MessageID returns the message ID that the route is responsible for.
	MessageID() MessageID
}
