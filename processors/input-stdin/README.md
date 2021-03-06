# STDIN
Read events from standard input.
By default, each event is assumed to be one line. If you want to join lines, you’ll want to use the multiline filter.

## Synopsys


|  SETTING  |  TYPE  | REQUIRED | DEFAULT VALUE |
|-----------|--------|----------|---------------|
| Add_field | hash   | false    | {}            |
| Tags      | array  | false    | []            |
| Type      | string | false    | ""            |
| Codec     | string | false    | ""            |


## Details

### Add_field
* Value type is hash
* Default value is `{}`

Add a field to an event

### Tags
* Value type is array
* Default value is `[]`

Add any number of arbitrary tags to your event.
This can help with processing later.

### Type
* Value type is string
* Default value is `""`

Add a type field to all events handled by this input

### Codec
* Value type is string
* Default value is `""`

The codec used for input data. Input codecs are a convenient method for decoding
your data before it enters the input, without needing a separate filter in your bitfan pipeline



## Configuration blueprint

```
stdin{
	add_field => {}
	tags => []
	type => ""
	codec => ""
}
```
