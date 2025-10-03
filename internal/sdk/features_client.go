package sdk

import "net/http"

type FeaturesClient struct {
	inner    HTTPClient
	features string
}

func NewFeaturesClient(inner HTTPClient, features string) *FeaturesClient {
	return &FeaturesClient{inner: inner, features: features}
}

func (this *FeaturesClient) Do(request *http.Request) (*http.Response, error) {
	if len(this.features) > 0 {
		values := request.URL.Query()
		values.Set("features", this.features)
		request.URL.RawQuery = values.Encode()
	}
	return this.inner.Do(request)
}
