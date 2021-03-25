package api

// Client interface with 3-rd party joke API
type Client interface {
	// GetJoke returns one Joke
	GetJoke() (*JokeResponse, error)
}
