package cmd

import (
	"Translate_CLI/api"
	"fmt"
)

type ServiceContainer struct {
	Translator api.Translator
}

// NewServiceContainer creates a new ServiceContainer instance
func NewServiceContainer() (*ServiceContainer, error) {
	translator, err := api.NewAzureTranslator() // 假设 NewAzureTranslator 可能返回 error
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Translator: %w", err)
	}
	return &ServiceContainer{Translator: translator}, nil
}
