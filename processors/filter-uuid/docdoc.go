// Code generated by "bitfanDoc "; DO NOT EDIT
package uuid

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "uuid",
  Doc:      "The uuid filter allows you to generate a UUID and add it as a field to each processed event.\nThis is useful if you need to generate a string that’s unique for every event, even if the same input is processed multiple times. If you want to generate strings that are identical each time a event with a given content is processed (i.e. a hash) you should use the fingerprint filter instead.\nThe generated UUIDs follow the version 4 definition in RFC 4122).",
  DocShort: "Adds a UUID to events",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "Add_field",
        Alias:        "",
        Doc:          "If this filter is successful, add any arbitrary fields to this event.\nField names can be dynamic and include parts of the event using the %{field}.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Add_tag",
        Alias:        "",
        Doc:          "If this filter is successful, add arbitrary tags to the event.\nTags can be dynamic and include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_field",
        Alias:        "",
        Doc:          "If this filter is successful, remove arbitrary fields from this event.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_Tag",
        Alias:        "",
        Doc:          "If this filter is successful, remove arbitrary tags from the event.\nTags can be dynamic and include parts of the event using the %{field} syntax",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Overwrite",
        Alias:        "",
        Doc:          "If the value in the field currently (if any) should be overridden by the generated UUID.\nDefaults to false (i.e. if the field is present, with ANY value, it won’t be overridden)",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Target",
        Alias:        "",
        Doc:          "Add a UUID to a field",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}