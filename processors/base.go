package processors

import (
	"github.com/mitchellh/mapstructure"
	"github.com/vjeantet/bitfan/processors/doc"
	"gopkg.in/go-playground/validator.v8"
)

const TimeFormat = "2006-01-02T15:04:05.999Z07:00"

type Base struct {
	Send                  PacketSender
	NewPacket             PacketBuilder
	Logger                Logger
	Memory                Memory
	ConfigWorkingLocation string
	DataLocation          string
}

func (b *Base) Doc() *doc.Processor {
	return &doc.Processor{}
}

func (b *Base) Configure(ctx ProcessorContext, conf map[string]interface{}) error { return nil }

func (b *Base) Receive(e IPacket) error { return nil }

func (b *Base) Tick(e IPacket) error { return nil }

func (b *Base) Start(e IPacket) error { return nil }

func (b *Base) Stop(e IPacket) error { return nil }

func (b *Base) MaxConcurent() int {
	return 0
}

func (b *Base) ConfigureAndValidate(ctx ProcessorContext, conf map[string]interface{}, rawVal interface{}) error {

	// Logger
	if ctx.Log != nil {
		b.Logger = ctx.Log()
	} else {
		// TODO set a dummy logger
	}

	b.ConfigWorkingLocation = ctx.ConfigWorkingLocation()

	// Packet Sender func
	if ctx.PacketSender != nil {
		b.Send = ctx.PacketSender()
	} else {
		// TODO set a dummy packetSender
	}

	// Packet Builder func
	if ctx.PacketBuilder != nil {
		b.NewPacket = ctx.PacketBuilder()
	} else {
		// TODO set a dummy PacketBuilder
	}

	// StoreSpace
	if ctx.Memory() != nil {
		b.Memory = ctx.Memory()
	} else {
		// TODO set a dummy PacketBuilder
	}

	b.DataLocation = ctx.DataLocation()

	// Set processor's user options
	if err := mapstructure.Decode(conf, rawVal); err != nil {
		return err
	}

	// validates processor's user options
	if err := validator.New(&validator.Config{TagName: "validate"}).Struct(rawVal); err != nil {
		return err
	}

	return nil
}
