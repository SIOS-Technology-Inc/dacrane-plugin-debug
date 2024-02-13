package main

import (
	"encoding/json"

	dacranepdk "github.com/SIOS-Technology-Inc/dacrane-pdk"
)

func main() {
	dacranepdk.ExecPluginJob(dacranepdk.Plugin{
		Config: dacranepdk.NewDefaultPluginConfig(),
		Resources: dacranepdk.MapToFunc(map[string]dacranepdk.Resource{
			"print": PrintResource,
			"dummy": DummyResource,
		}),
	})
}

var PrintResource = dacranepdk.Resource{
	Create: func(parameter any, meta dacranepdk.PluginMeta) (any, error) {
		v, err := json.MarshalIndent(parameter, "", "  ")
		if err != nil {
			return nil, err
		}
		meta.Log(string(v))

		return parameter, nil
	},
	Update: func(current, previous any, meta dacranepdk.PluginMeta) (any, error) {
		v, err := json.MarshalIndent(current, "", "  ")
		if err != nil {
			return nil, err
		}
		meta.Log(string(v))

		return current, nil
	},
	Delete: func(parameter any, meta dacranepdk.PluginMeta) error {
		v, err := json.MarshalIndent(parameter, "", "  ")
		if err != nil {
			return err
		}

		meta.Log(string(v))

		return nil
	},
}

var DummyResource = dacranepdk.Resource{
	Create: func(parameter any, meta dacranepdk.PluginMeta) (any, error) {
		return parameter, nil
	},
	Update: func(current, previous any, meta dacranepdk.PluginMeta) (any, error) {
		return current, nil
	},
	Delete: func(parameter any, meta dacranepdk.PluginMeta) error {
		return nil
	},
}
