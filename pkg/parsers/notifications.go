package parsers

type general struct {
	Token   string `yaml:"token"`
	Color   string `yaml:"color"`
	Channel string `yaml:"channel"`
}

type telegram struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}
