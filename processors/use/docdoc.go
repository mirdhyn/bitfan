// Code generated by "bitfanDoc "; DO NOT EDIT
package use

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "use",
  Doc:      "",
  DocShort: "Include a config file",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "Add_field",
        Alias:        "",
        Doc:          "If this processor is successful, add any arbitrary fields to this event.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Add_tag",
        Alias:        "",
        Doc:          "If this processor is successful, add arbitrary tags to the event.\nTags can be dynamic and include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_field",
        Alias:        "",
        Doc:          "If this processor is successful, remove arbitrary fields from this event.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_tag",
        Alias:        "",
        Doc:          "If this processor is successful, remove arbitrary tags from the event.\nTags can be dynamic and include parts of the event using the %{field} syntax",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "",
        Doc:          "Add a type field to all events handled by this processor",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Path",
        Alias:        "path",
        Doc:          "Path to configuration to import in this pipeline, it could be a local file or an url\ncan be relative path to the current configuration.\nSPLIT and JOIN : in filter Section, set multiples path to make a split and join into your pipeline",
        Required:     true,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "path=> [\"meteo-input.conf\"]",
      },
      &doc.ProcessorOption{
        Name:         "Var",
        Alias:        "var",
        Doc:          "You can set variable references in the used configuration by using ${var}.\neach reference will be replaced by the value of the variable found in this option\nThe replacement is case-sensitive.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "var => {\"hostname\"=>\"myhost\",\"varname\"=>\"varvalue\"}",
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