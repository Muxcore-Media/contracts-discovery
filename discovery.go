package discovery

import "context"

// Note: imports contracts-media for MediaType.
// In your go.mod: require github.com/Muxcore-Media/contracts-media v1.0.0
// type MediaType = media.MediaType  (or just use string)

type DiscoveryItem struct {
	Title       string
	Year        int
	MediaType   string
	ExternalID  string
	Source      string
	Synopsis    string
	PosterURL   string
	BackdropURL string
	Rating      float64
	Genres      []string
	Metadata    map[string]string
}

type DiscoveryFilter struct {
	MediaType string
	Genre     string
	YearMin   int
	YearMax   int
	RatingMin float64
	Language  string
}

type MediaDiscovery interface {
	Trending(ctx context.Context, mediaType string, limit int) ([]DiscoveryItem, error)
	Popular(ctx context.Context, mediaType string, limit int) ([]DiscoveryItem, error)
	Upcoming(ctx context.Context, mediaType string, limit int) ([]DiscoveryItem, error)
	SearchExternal(ctx context.Context, query string, filter DiscoveryFilter) ([]DiscoveryItem, error)
}

const EventDiscoveryUpdated = "discovery.updated"
