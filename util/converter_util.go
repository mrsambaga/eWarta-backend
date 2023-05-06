package util

import "stage01-project-backend/httperror"

func ConvertTypeToTypeId(postType string) (uint64, error) {
	switch postType {
	case "free":
		return 1, nil
	case "premium":
		return 2, nil
	case "vip":
		return 3, nil
	default:
		return 0, httperror.ErrInvalidType
	}
}

func ConvertCategoryToCategoryId(category string) (uint64, error) {
	switch category {
	case "Business":
		return 1, nil
	case "Technology":
		return 2, nil
	case "Politic":
		return 3, nil
	case "Sport":
		return 4, nil
	case "Science":
		return 5, nil
	case "Economy":
		return 6, nil
	default:
		return 0, httperror.ErrInvalidCategory
	}
}
