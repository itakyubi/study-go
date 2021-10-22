package database

type Config struct {
	Database struct {
		Type            string `yaml:"type" json:"type" validate:"nonzero"`
		URL             string `yaml:"url" json:"url" validate:"nonzero"`
		MaxOpenConns    int    `yaml:"maxOpenConns" json:"maxOpenConns" default:20`
		MaxIdleConns    int    `yaml:"maxIdleConns" json:"maxIdleConns" default:20`
		ConnMaxLifetime int    `yaml:"connMaxLifetime" json:"connMaxLifetime" default:150`
	} `yaml:"database" json:"database" default:"{}"`
}
