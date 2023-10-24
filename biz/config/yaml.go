package config

type YamlPart struct{}

func (y *YamlPart) Read() error {
	return nil
}

func (y *YamlPart) Validate() error {
	return nil
}
