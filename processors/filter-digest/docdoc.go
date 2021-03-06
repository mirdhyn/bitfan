// Code generated by "bitfanDoc "; DO NOT EDIT
package digest

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "digest",
  Doc:      "",
  DocShort: "Digest events every x",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "AddField",
        Alias:        "add_field",
        Doc:          "If this filter is successful, add any arbitrary fields to this event.",
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
        Name:         "RemoveField",
        Alias:        "remove_field",
        Doc:          "If this filter is successful, remove arbitrary fields from this event. Example:\n` kv {\n`   remove_field => [ \"foo_%{somefield}\" ]\n` }",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RemoveTag",
        Alias:        "remove_tag",
        Doc:          "If this filter is successful, remove arbitrary tags from the event. Tags can be dynamic and include parts of the event using the %{field} syntax.\nExample:\n` kv {\n`   remove_tag => [ \"foo_%{somefield}\" ]\n` }\nIf the event has field \"somefield\" == \"hello\" this filter, on success, would remove the tag foo_hello if it is present. The second example would remove a sad, unwanted tag as well.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "KeyMap",
        Alias:        "key_map",
        Doc:          "Add received event fields to the digest field named with the key map_key\nWhen this setting is empty, digest will merge fields from coming events",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "key_map => \"type\"",
      },
      &doc.ProcessorOption{
        Name:         "Interval",
        Alias:        "interval",
        Doc:          "When should Digest send a digested event ?\nUse CRON or BITFAN notation",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "interval => \"every_10s\"",
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