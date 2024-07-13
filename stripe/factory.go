package stripe

// ClientFactory is the factory that will create instances of StripeClient
type ClientFactory struct{}

// NewClientFactory returns a new instance of ClientFactory
func NewClientFactory() *ClientFactory {
	return &ClientFactory{}
}

// CreateClient creates a new StripeClient
func (f *ClientFactory) CreateClient() StripeClient {
	return NewStripeClient()
}
