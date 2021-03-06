package media

// Code generated by cdproto-gen. DO NOT EDIT.

// EventPlayerPropertiesChanged this can be called multiple times, and can be
// used to set / override / remove player properties. A null propValue indicates
// removal.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#event-playerPropertiesChanged
type EventPlayerPropertiesChanged struct {
	PlayerID   PlayerID          `json:"playerId"`
	Properties []*PlayerProperty `json:"properties"`
}

// EventPlayerEventsAdded send events as a list, allowing them to be batched
// on the browser for less congestion. If batched, events must ALWAYS be in
// chronological order.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#event-playerEventsAdded
type EventPlayerEventsAdded struct {
	PlayerID PlayerID       `json:"playerId"`
	Events   []*PlayerEvent `json:"events"`
}

// EventPlayerMessagesLogged send a list of any messages that need to be
// delivered.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#event-playerMessagesLogged
type EventPlayerMessagesLogged struct {
	PlayerID PlayerID         `json:"playerId"`
	Messages []*PlayerMessage `json:"messages"`
}

// EventPlayerErrorsRaised send a list of any errors that need to be
// delivered.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#event-playerErrorsRaised
type EventPlayerErrorsRaised struct {
	PlayerID PlayerID       `json:"playerId"`
	Errors   []*PlayerError `json:"errors"`
}

// EventPlayersCreated called whenever a player is created, or when a new
// agent joins and receives a list of active players. If an agent is restored,
// it will receive the full list of player ids and all events again.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#event-playersCreated
type EventPlayersCreated struct {
	Players []PlayerID `json:"players"`
}
