// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.jespino.recommend",
  "name": "Recommend",
  "description": "This plugin recommends you channels",
  "homepage_url": "https://github.com/jespino/mattermost-plugin-recommend",
  "support_url": "https://github.com/jespino/mattermost-plugin-recommend/issues",
  "icon_path": "assets/icon.svg",
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
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "RecommendOnJoinTeam",
        "display_name": "Recommend at team join",
        "type": "bool",
        "help_text": "When user joins to a team, recommend bot is going to recommend interesting channels in that team.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "RecommendOnJoinChannel",
        "display_name": "Recommend at channel join",
        "type": "bool",
        "help_text": "When user joins to a channel, recommend bot is going to recommend other channels in the team based on the people in that channel.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "GracePeriod",
        "display_name": "Grace period",
        "type": "number",
        "help_text": "Give a period of time since the user was created before start sending automatic messages on join.",
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
