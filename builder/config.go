package builder

type JenkinsConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	JenkinsURL string `yaml:"url"`
}
