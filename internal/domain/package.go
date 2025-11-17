package domain

type Package struct {
	Name       string `yaml:"-"`
	Repository string `yaml:"repository"`
}
