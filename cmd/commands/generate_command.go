package commands

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/wesleysnt/finance-api/app/config"
	"gopkg.in/yaml.v3"
)

var generateJwtKeyCmd = &cobra.Command{
	Use:   "generate:jwt_key",
	Short: "Generate Secret Key For JWT",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		confFile := fmt.Sprintf("%s/.yaml", pwd)

		data, err := os.ReadFile(confFile)

		if err != nil {
			color.Redln("Error reading environment file")
			return
		}

		config := config.Env{}
		yaml.Unmarshal(data, &config)

		if config.Jwt.Secret != "" {
			color.Redln("JWT secret already exists")
			return
		}

		secretBytes := make([]byte, 32)

		rand.Read([]byte(secretBytes))

		jwtSecret := base64.URLEncoding.EncodeToString(secretBytes)

		config.Jwt.Secret = jwtSecret

		data, err = yaml.Marshal(&config)
		if err != nil {
			color.Redln("failed to marshal YAML")
			return
		}

		err = os.WriteFile(confFile, data, 0644)
		if err != nil {
			color.Redln("Error saving file: %v", err)
			return
		}

		color.Greenln("JWT Secret Generated!")
		return
	},
}

func init() {
	gobaseCommand.AddCommand(generateJwtKeyCmd)
}
