package config

// Environment variable keys for Claude Code configuration
const (
	EnvAuthToken            = "ANTHROPIC_AUTH_TOKEN"
	EnvBaseURL             = "ANTHROPIC_BASE_URL"
	EnvDefaultHaikuModel   = "ANTHROPIC_DEFAULT_HAIKU_MODEL"
	EnvDefaultOpusModel    = "ANTHROPIC_DEFAULT_OPUS_MODEL"
	EnvDefaultSonnetModel  = "ANTHROPIC_DEFAULT_SONNET_MODEL"
	EnvModel               = "ANTHROPIC_MODEL"
)

// EnvKeys is the list of all environment variable keys for profiles
var EnvKeys = []string{
	EnvAuthToken,
	EnvBaseURL,
	EnvDefaultHaikuModel,
	EnvDefaultOpusModel,
	EnvDefaultSonnetModel,
	EnvModel,
}

// EnvLabels provides human-readable labels for environment keys
var EnvLabels = map[string]string{
	EnvAuthToken:           "API Token",
	EnvBaseURL:             "Base URL",
	EnvDefaultHaikuModel:   "Haiku Model",
	EnvDefaultOpusModel:    "Opus Model",
	EnvDefaultSonnetModel:  "Sonnet Model",
	EnvModel:               "Default Model",
}

// SetAuthToken sets the ANTHROPIC_AUTH_TOKEN
func (p *Profile) SetAuthToken(token string) {
	p.SetEnv(EnvAuthToken, token)
}

// GetAuthToken gets the ANTHROPIC_AUTH_TOKEN
func (p *Profile) GetAuthToken() (string, bool) {
	return p.GetEnv(EnvAuthToken)
}

// SetBaseURL sets the ANTHROPIC_BASE_URL
func (p *Profile) SetBaseURL(url string) {
	p.SetEnv(EnvBaseURL, url)
}

// GetBaseURL gets the ANTHROPIC_BASE_URL
func (p *Profile) GetBaseURL() (string, bool) {
	return p.GetEnv(EnvBaseURL)
}

// SetHaikuModel sets the ANTHROPIC_DEFAULT_HAIKU_MODEL
func (p *Profile) SetHaikuModel(model string) {
	p.SetEnv(EnvDefaultHaikuModel, model)
}

// GetHaikuModel gets the ANTHROPIC_DEFAULT_HAIKU_MODEL
func (p *Profile) GetHaikuModel() (string, bool) {
	return p.GetEnv(EnvDefaultHaikuModel)
}

// SetOpusModel sets the ANTHROPIC_DEFAULT_OPUS_MODEL
func (p *Profile) SetOpusModel(model string) {
	p.SetEnv(EnvDefaultOpusModel, model)
}

// GetOpusModel gets the ANTHROPIC_DEFAULT_OPUS_MODEL
func (p *Profile) GetOpusModel() (string, bool) {
	return p.GetEnv(EnvDefaultOpusModel)
}

// SetSonnetModel sets the ANTHROPIC_DEFAULT_SONNET_MODEL
func (p *Profile) SetSonnetModel(model string) {
	p.SetEnv(EnvDefaultSonnetModel, model)
}

// GetSonnetModel gets the ANTHROPIC_DEFAULT_SONNET_MODEL
func (p *Profile) GetSonnetModel() (string, bool) {
	return p.GetEnv(EnvDefaultSonnetModel)
}

// SetModel sets the ANTHROPIC_MODEL
func (p *Profile) SetModel(model string) {
	p.SetEnv(EnvModel, model)
}

// GetModel gets the ANTHROPIC_MODEL
func (p *Profile) GetModel() (string, bool) {
	return p.GetEnv(EnvModel)
}
