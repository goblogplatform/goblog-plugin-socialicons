package main

import "html"

type social struct {
	key   string // setting key
	label string // display label
	icon  string // Font Awesome icon class
}

// socials is rendered in this order.
var socials = []social{
	{"github_url", "GitHub", "fab fa-github"},
	{"linkedin_url", "LinkedIn", "fab fa-linkedin"},
	{"x_url", "X", "fab fa-x"},
	{"keybase_url", "Keybase", "fab fa-keybase"},
	{"instagram_url", "Instagram", "fab fa-instagram"},
	{"facebook_url", "Facebook", "fab fa-facebook"},
	{"strava_url", "Strava", "fab fa-strava"},
	{"spotify_url", "Spotify", "fab fa-spotify"},
	{"xbox_url", "Xbox", "fab fa-xbox"},
	{"steam_url", "Steam", "fab fa-steam"},
}

// setting mirrors goblog's settings JSON shape.
type setting struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// settingDefs keeps the keys of the compiled-in plugin this one replaces,
// so an upgraded site's saved settings carry over.
func settingDefs() []setting {
	defs := []setting{
		{Key: "enabled", Type: "text", Default: "true", Label: "Enabled", Description: "Set to 'true' to show social icons"},
	}
	for _, s := range socials {
		defs = append(defs, setting{Key: s.key, Type: "text", Default: "", Label: s.label + " URL", Description: "Full URL to your " + s.label + " profile"})
	}
	return defs
}

// link is one configured profile, exposed to templates as .links.
type link struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
	Icon string `json:"Icon"`
}

// links returns the configured profiles in display order, or nil when the
// plugin is disabled.
func links(settings map[string]string) []link {
	if settings["enabled"] != "true" {
		return nil
	}
	var out []link
	for _, s := range socials {
		if url := settings[s.key]; url != "" {
			out = append(out, link{Name: s.label, URL: url, Icon: s.icon})
		}
	}
	return out
}

// footerHTML renders the icon row. Setting values are admin-typed, but
// the browser trusts whatever comes back, so everything is escaped.
func footerHTML(settings map[string]string) string {
	ls := links(settings)
	if ls == nil {
		return ""
	}
	out := `<div class="text-center" style="padding: 10px 0;">`
	for _, l := range ls {
		out += `<a href="` + html.EscapeString(l.URL) + `" target="_blank" rel="noopener noreferrer" title="` + html.EscapeString(l.Name) + `" style="margin: 0 6px; color: inherit;"><i class="` + l.Icon + ` fa-1x"></i></a>`
	}
	return out + `</div>`
}
