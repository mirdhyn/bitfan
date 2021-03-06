// Code generated by "bitfanDoc "; DO NOT EDIT
package beatsinput

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "beatsinput",
  Doc:      "",
  DocShort: "",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "Add_field",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Codec",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Congestion_threshold",
        Alias:        "",
        Doc:          "The number of seconds before we raise a timeout,\nthis option is useful to control how much time to wait if something is blocking\nthe pipeline",
        Required:     false,
        Type:         "int",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Host",
        Alias:        "",
        Doc:          "The IP address to listen on",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Port",
        Alias:        "",
        Doc:          "The port to listen on (default 5044)",
        Required:     false,
        Type:         "int",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl",
        Alias:        "",
        Doc:          "Events are by default send in plain text,\nyou can enable encryption by using ssl to true and\nconfiguring the ssl_certificate and ssl_key options",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl_certificate",
        Alias:        "",
        Doc:          "SSL certificate to use (path)",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl_certificate_authorities",
        Alias:        "",
        Doc:          "Validate client certificates against theses authorities\n You can defined multiples files or path, all the certificates will be read\n and added to the trust store.\n You need to configure the ssl_verify_mode to peer or force_peer to enable\n the verification.\nThis feature only support certificate directly signed by your root ca.\nIntermediate CA are currently not supported.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl_key",
        Alias:        "",
        Doc:          "SSL key to use (path)",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl_key_passphrase",
        Alias:        "",
        Doc:          "SSL key passphrase to use. (not yet implemented)",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ssl_verify_mode",
        Alias:        "",
        Doc:          "By default the server dont do any client verification,\npeer will make the server ask the client to provide a certificate,\n  if the client provide the certificate it will be validated.\nforce_peer will make the server ask the client for their certificate,\n  if the clients doesn’t provide it the connection will be closed.\nThis option need to be used with ssl_certificate_authorities and a defined list of CA.\nValue can be any of: none, peer, force_peer",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Tags",
        Alias:        "",
        Doc:          "Add any number of arbitrary tags to your event",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "",
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