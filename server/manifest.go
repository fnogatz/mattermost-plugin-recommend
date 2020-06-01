// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.jespino.recomend",
  "name": "Recommend",
  "description": "This plugin recommends you channels",
  "version": "0.0.1",
  "min_server_version": "5.16.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "RecommendOnJoinTeam",
        "display_name": "Recommend at team join",
        "type": "bool",
        "help_text": "",
        "placeholder": "",
        "default": null
      },
      {
        "key": "RecommendOnJoinChannel",
        "display_name": "Recommend at channel join",
        "type": "bool",
        "help_text": "",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}