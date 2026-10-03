package main

import (
	"strings"
	"testing"
)

func TestFooterHTMLRendersOnlyConfiguredLinks(t *testing.T) {
	out := footerHTML(map[string]string{"enabled": "true", "github_url": "https://github.com/me", "x_url": "https://x.com/me"})
	if !strings.Contains(out, `href="https://github.com/me"`) || !strings.Contains(out, `fa-github`) {
		t.Fatalf("github link missing: %q", out)
	}
	if !strings.Contains(out, `href="https://x.com/me"`) || !strings.Contains(out, `fa-x`) {
		t.Fatalf("x link missing: %q", out)
	}
	if strings.Contains(out, "linkedin") {
		t.Fatalf("unconfigured network rendered: %q", out)
	}
	if strings.Count(out, `rel="me noopener noreferrer"`) != 2 {
		t.Fatalf("each link needs rel=me and noopener: %q", out)
	}
}

func TestFooterHTMLEscapesURLs(t *testing.T) {
	out := footerHTML(map[string]string{"enabled": "true", "github_url": `https://g.h/"><script>`})
	if strings.Contains(out, `<script>`) {
		t.Fatalf("url not escaped: %q", out)
	}
	if !strings.Contains(out, `&lt;script&gt;`) {
		t.Fatalf("expected escaped url: %q", out)
	}
}

func TestFooterHTMLEmptyWhenDisabled(t *testing.T) {
	if out := footerHTML(map[string]string{"enabled": "false", "github_url": "https://github.com/me"}); out != "" {
		t.Fatalf("want empty, got %q", out)
	}
}

func TestLinksKeepsDeclarationOrder(t *testing.T) {
	links := links(map[string]string{"enabled": "true", "steam_url": "s", "github_url": "g", "linkedin_url": "l"})
	var names []string
	for _, l := range links {
		names = append(names, l.Name)
	}
	if got := strings.Join(names, ","); got != "GitHub,LinkedIn,Steam" {
		t.Fatalf("order = %s", got)
	}
	if links[0].URL != "g" || links[0].Icon != "fab fa-github" {
		t.Fatalf("first link = %+v", links[0])
	}
}

func TestLinksNilWhenDisabled(t *testing.T) {
	if l := links(map[string]string{"enabled": "false", "github_url": "g"}); l != nil {
		t.Fatalf("want nil, got %v", l)
	}
}

func TestSettingsKeysMatchTheCompiledInPlugin(t *testing.T) {
	// The same keys keep an upgraded site's plugin_settings rows working.
	var keys []string
	for _, s := range settingDefs() {
		keys = append(keys, s.Key)
	}
	want := "enabled,github_url,linkedin_url,x_url,keybase_url,instagram_url,facebook_url,strava_url,spotify_url,xbox_url,steam_url"
	if got := strings.Join(keys, ","); got != want {
		t.Fatalf("keys = %s", got)
	}
	if settingDefs()[0].Default != "true" {
		t.Fatalf("enabled must default to true, as before")
	}
}
