// Package player : supports playing videos via different players: Web (HTML), MPV and later
// other video players
package mplayer

import (
	"fmt"
	"strings"
)

// Player : interface functions for every Player
type Player interface {
	Play()
	SetTitle(string)
	SetURL(string)
}

// Props : Props to be passed to Player
type Props struct {
	URL   string
	Title string
}

func (p *Props) SetTitle(title string) {
	p.Title = title
}

func (p *Props) SetURL(URL string) {
	p.URL = URL
}

// GetPlayers : A map of all available players
func GetPlayers() map[string]Player {
	engines := make(map[string]Player)
	engines["browser"] = &BrowserPlayer{}
	engines["mpv"] = &MPVPlayer{}
	engines["vlc"] = &VLCPlayer{}
	//  engines["mpv"] = NewFzEngine()
	return engines
}

// GetPlayer : get a player for streaming
func GetPlayer(player string) (Player, error) {
	e := GetPlayers()[strings.ToLower(player)]
	if e == nil {
		return nil, fmt.Errorf("Player %s Does not exist", player)
	}
	return e, nil
}
