package mapping

import "github.com/mitchellh/mapstructure"

func WithTag(tag string) Option {
	return func(config *mapstructure.DecoderConfig) {
		config.TagName = tag
	}
}

func WithMeta(meta *mapstructure.Metadata) Option {
	return func(config *mapstructure.DecoderConfig) {
		config.Metadata = meta
	}
}

func WithSquash() Option {
	return func(config *mapstructure.DecoderConfig) {
		config.Squash = true
	}
}

func WithDecodeHook(hook mapstructure.DecodeHookFunc) Option {
	return func(config *mapstructure.DecoderConfig) {
		config.DecodeHook = hook
	}
}

func WithErrorUnused() Option {
	return func(config *mapstructure.DecoderConfig) {
		config.ErrorUnused = true
	}
}

func WithZeroFields() Option {
	return func(config *mapstructure.DecoderConfig) {
		config.ZeroFields = true
	}
}

func WithWeaklyTypedInput() Option {
	return func(config *mapstructure.DecoderConfig) {
		config.WeaklyTypedInput = true
	}
}
