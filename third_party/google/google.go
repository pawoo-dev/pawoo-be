package third_party

import (
	third_party "github.com/pawoo-dev/pawoo-be/third_party/aws"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

const az = "ap-southeast-1"

type GoogleClient struct {
	Client *maps.Client
}

type google_api_key struct {
	GoogleApiKey string `json:"google_api_key"`
}

func GetGoogleApiKey() string {
	var apiKey google_api_key
	sm := third_party.SecretManager{
		Region:     az,
		SecretName: "evapp_api_key",
	}
	sm.GetSecrets(&apiKey)
	return apiKey.GoogleApiKey
}

func NewGoogleClient() *GoogleClient {
	key := GetGoogleApiKey()
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		logrus.WithField("err", err).Error("failied_to_init_google_client")
		panic(err)
	}
	return &GoogleClient{
		Client: c,
	}
}
