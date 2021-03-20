package mapping

import "github.com/mitchellh/mapstructure"

type Option func(config *mapstructure.DecoderConfig)

func Mapping(input interface{}, output interface{}, opts ...Option) error {
	config := &mapstructure.DecoderConfig{
		Result:   output,
		Metadata: nil,
	}
	for _, opt := range opts {
		opt(config)
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func Decode(input interface{}, opts ...Option) map[string]interface{} {
	result := make(map[string]interface{})
	_ = Mapping(input, &result, opts...)
	return result
}
