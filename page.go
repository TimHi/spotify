// Copyright 2014, 2015 Zac Bergquist
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spotify

import (
	"encoding/json"
	"errors"
)

// ErrNoMorePages is the error returned when you attempt to get the next
// (or previous) set of data but you've reached the end of the data set.
var ErrNoMorePages = errors.New("spotify: no more pages")

// This file contains the types that implement Spotify's paging object.
// See: https://developer.spotify.com/web-api/object-model/#paging-object

// basePage contains all of the fields in a Spotify paging object, except
// for the actual items.  This type is meant to be embedded in other types
// that add the Items field.
type basePage struct {
	// A link to the Web API Endpoint returning the full
	// result of this request.
	Endpoint string `json:"href"`
	// The maximum number of items in the response, as set
	// in the query (or default value if unset).
	Limit int `json:"limit"`
	// The offset of the items returned, as set in the query
	// (or default value if unset).
	Offset int `json:"offset"`
	// The total number of items available to return.
	Total int `json:"total"`
	// The URL to the next page of items (if available).
	Next string `json:"next"`
	// The URL to the previous page of items (if available).
	Previous string `json:"previous"`
}

// FullArtistPage contains FullArtists returned by the Web API.
type FullArtistPage struct {
	basePage
	Artists []FullArtist `json:"items"`
}

// SimpleAlbumPage contains SimpleAlbums returned by the Web API.
type SimpleAlbumPage struct {
	basePage
	Albums []SimpleAlbum `json:"items"`
}

// SavedAlbumPage contains SavedAlbums returned by the Web API.
type SavedAlbumPage struct {
	basePage
	Albums []SavedAlbum `json:"items"`
}

// SimplePlaylistPage contains SimplePlaylists returned by the Web API.
type SimplePlaylistPage struct {
	basePage
	Playlists []SimplePlaylist `json:"items"`
}

// SimpleTrackPage contains SimpleTracks returned by the Web API.
type SimpleTrackPage struct {
	basePage
	Tracks []SimpleTrack `json:"items"`
}

// FullTrackPage contains FullTracks returned by the Web API.
type FullTrackPage struct {
	basePage
	Tracks []FullTrack `json:"items"`
}

// SavedTrackPage contains SavedTracks return by the Web API.
type SavedTrackPage struct {
	basePage
	Tracks []SavedTrack `json:"items"`
}

// PlaylistTrackPage contains information about tracks in a playlist.
type PlaylistTrackPage struct {
	basePage
	Tracks []PlaylistTrack `json:"items"`
}

// CategoryPage contains Category objects returned by the Web API.
type CategoryPage struct {
	basePage
	Categories []Category `json:"items"`
}

// getPage GETs the data at the specified URL and unmarshals it into page.
func (c *Client) getPage(url string, page interface{}) error {
	resp, err := c.http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(page)
}
