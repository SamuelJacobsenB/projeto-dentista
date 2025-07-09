package utils

import "errors"

func ValidImageExtension(extension string) error {
	switch extension {
	case "jpg", "jpeg", "png", "webp":
		return nil
	default:
		return errors.New("extensão de imagem inválida")
	}
}
