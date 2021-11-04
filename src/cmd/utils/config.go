package utils

type Configuration struct {
	Server ServerSettings
	Queue  QueueSettings
}

type ServerSettings struct {
	Port string
}

type QueueSettings struct {
	Url string
}
