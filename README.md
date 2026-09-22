# goblog-plugin-socialicons

A [goblog](https://github.com/goblogplatform/goblog) WebAssembly plugin that renders a row of icon links to your profiles in the footer of every page.

## Install

From your goblog's **Admin → Plugins → Browse**, search for *Social Icons* and click **Install**. Or download `plugin.wasm` from the [latest release](https://github.com/goblogplatform/goblog-plugin-socialicons/releases/latest) into `plugins/wasm/` as `socialicons.wasm` and restart goblog.

Upgrading from goblog ≤ 0.6.x, where this plugin was compiled in: install it from the directory and your existing profile URLs carry over (same plugin name, same setting keys).

## Settings

Under **Admin → Settings → Social Icons**: `enabled` (default `true`) and one URL per network — `github_url`, `linkedin_url`, `x_url`, `keybase_url`, `instagram_url`, `facebook_url`, `strava_url`, `spotify_url`, `xbox_url`, `steam_url`. Only networks with a URL are shown, in that order.

Icons are Font Awesome brand classes (`fab fa-github` …), which goblog's themes already load.

## For theme authors

The plugin also exposes the configured links as template data, so a theme can place them itself instead of using the footer row:

```
{{ range .plugins.socialicons.links }}<a href="{{ .URL }}" title="{{ .Name }}"><i class="{{ .Icon }}"></i></a>{{ end }}
```

## Build it yourself

```bash
GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -ldflags="-s -w" -o plugin.wasm .
go test ./...   # native tests for the rendering
```

Needs Go 1.25 or newer. Implements the `identity`, `settings`, `template_footer` and `template_data` exports and makes no network requests; see [goblog.live/docs/plugin-api](https://www.goblog.live/docs/plugin-api).

## License

Apache-2.0.
