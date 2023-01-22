package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigShouldSucceed(t *testing.T) {
	//given
	correctConfigPath := "../../assets/test/correct_config.json"
	expectedConfig := &Config{
		Token: "exampleBotToken",
	}

	//when
	config, err := LoadConfig(correctConfigPath)

	//then
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestLoadConfigShouldFailWhenWrongPathWasProvided(t *testing.T) {
	//given
	nonexistentConfigPath := "../../assets/test/nonexistent_config.json"
	var expectedConfig *Config

	//when
	config, err := LoadConfig(nonexistentConfigPath)

	//then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, expectedConfig, config)
}

func TestLoadConfigShouldFailWhenIncorrectFileWasProvided(t *testing.T) {
	//given
	incorrectConfigPath := "../../assets/test/incorrect_config.json"
	var expectedConfig *Config

	//when
	config, err := LoadConfig(incorrectConfigPath)

	//then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
	assert.Equal(t, expectedConfig, config)
}
