// Code generated by "bitfanDoc "; DO NOT EDIT
package sysloginput

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "sysloginput",
  Doc:      "",
  DocShort: "",
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
        Name:         "Format",
        Alias:        "format",
        Doc:          "Which format to use to decode syslog messages. Default value is \"automatic\"\nValue can be \"automatic\",\"rfc3164\", \"rfc5424\" or \"rfc6587\"\n\nNote on \"automatic\" format: if you don't know which format to select,\nor have multiple incoming formats, this is the one to go for.\nThere is a theoretical performance penalty (it has to look at a few bytes\nat the start of the frame), and a risk that you may parse things you don't want to parse\n(rogue syslog clients using other formats), so if you can be absolutely sure of your syslog\nformat, it would be best to select it explicitly.",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Port",
        Alias:        "port",
        Doc:          "Port number to listen on",
        Required:     false,
        Type:         "int",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Protocol",
        Alias:        "protocol",
        Doc:          "Protocol to use to listen to syslog messages\nValue can be either \"tcp\" or \"udp\"",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Tags",
        Alias:        "tags",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "type",
        Doc:          "Add a type field to all events handled by this input",
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