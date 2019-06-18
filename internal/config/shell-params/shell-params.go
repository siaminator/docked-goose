package shell_params

type dbConfig struct {
	source string
}

func NewShellParamConfig(source string) *dbConfig {
	dbConfig := dbConfig{}
	if source == "" {
		source = "user=local password=local dbname=local sslmode=disable host=postgres port=5432"
	}
	dbConfig.source = source
	return &dbConfig
}

func (dbc *dbConfig) GetDriver() (string, error) {
	return "postgres", nil
}

func (dbc *dbConfig) GetDbString() (string, error) {
	return dbc.source, nil
}
