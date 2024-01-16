package out

type Config struct {
	Source struct {
		Type            string `yaml:"type"`
		FirstContentRow int    `yaml:"firstContentRow"`
		Sheet           string `yaml:"sheet"`
	}
}

type Yml interface {
	ReadConfig() (Config, error)
}
