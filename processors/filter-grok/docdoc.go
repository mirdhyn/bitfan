// Code generated by "bitfanDoc "; DO NOT EDIT
package grok

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "grok",
  Doc:      "",
  DocShort: "",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "AddField",
        Alias:        "add_field",
        Doc:          "If this filter is successful, add any arbitrary fields to this event. Field names can\nbe dynamic and include parts of the event using the %{field}.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "AddTag",
        Alias:        "add_tag",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "BreakOnMatch",
        Alias:        "break_on_match",
        Doc:          "Break on first match. The first successful match by grok will result in the filter being\nfinished. If you want grok to try all patterns (maybe you are parsing different things),\nthen set this to false",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "KeepEmptyCaptures",
        Alias:        "keep_empty_captures",
        Doc:          "If true, keep empty captures as event fields",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Match",
        Alias:        "match",
        Doc:          "A hash of matches of field ⇒ value\n@nodefault\n\nFor example:\n```\n    filter {\n      grok { match => { \"message\" => \"Duration: %{NUMBER:duration}\" } }\n    }\n```\nIf you need to match multiple patterns against a single field, the value can be an array of patterns\n```\n    filter {\n      grok { match => { \"message\" => [ \"Duration: %{NUMBER:duration}\", \"Speed: %{NUMBER:speed}\" ] } }\n    }\n```",
        Required:     true,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "NamedCapturesOnly",
        Alias:        "named_capture_only",
        Doc:          "If true, only store named captures from grok.",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "PatternsDir",
        Alias:        "patterns_dir",
        Doc:          "BitFan ships by default with a bunch of patterns, so you don’t necessarily need to\ndefine this yourself unless you are adding additional patterns. You can point to\nmultiple pattern directories using this setting Note that Grok will read all files\nin the directory and assume its a pattern file (including any tilde backup files)",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RemoveField",
        Alias:        "remove_field",
        Doc:          "If this filter is successful, remove arbitrary fields from this event",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RemoveTag",
        Alias:        "remove_tag",
        Doc:          "If this filter is successful, remove arbitrary tags from the event.\nTags can be dynamic and include parts of the event using the %{field} syntax",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "TagOnFailure",
        Alias:        "tag_on_failure",
        Doc:          "Append values to the tags field when there has been no successful match",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{
    &doc.ProcessorPort{
      Default: true,
      Name:    "PORT_SUCCESS",
      Number:  0,
      Doc:     "",
    },
  },
}
}