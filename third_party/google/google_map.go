package third_party

import (
	"context"
	"strings"

	"googlemaps.github.io/maps"
)

func (g *GoogleClient) GetRecommendation(keyword string) ([]maps.AutocompletePrediction, error) {

	res, err := g.Client.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
		Input:        keyword,
		StrictBounds: true,
		Location: &maps.LatLng{
			Lat: 1.3937605,
			Lng: 103.8066033,
		},
		Radius: 50000,
	})
	var newRR []maps.AutocompletePrediction

	if err != nil {
		return nil, err
	}

	for _, p := range res.Predictions {
		if strings.Contains(p.Description, "Singapore") {
			newRR = append(newRR, p)
		}
	}
	return newRR, nil
}

func (g *GoogleClient) GetPlaceDetails(placeId string) (maps.PlaceDetailsResult, error) {
	res, err := g.Client.PlaceDetails(context.Background(), &maps.PlaceDetailsRequest{
		PlaceID: placeId,
	})
	return res, err
}
