package utils

import (
	"fmt"
	"strconv"
	"strings"

	"insights-pulse/src/logger"
)

func ConvToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%.1f", v)
	default:
		return ""
	}
}

// Function to stringfy IDs into strings with a max of 20 IDs per string
//
//	[123, 987, 989] => "123-987-989"
func StringfyIds(ids []int, maxPerString int) []string {
	var result []string
	var currentChunk []string

	for i, id := range ids {
		currentChunk = append(currentChunk, fmt.Sprintf("%d", id))
		if (i+1)%maxPerString == 0 || i == len(ids)-1 {
			result = append(result, strings.Join(currentChunk, "-"))
			currentChunk = nil
		}
	}

	return result
}

func GetFloatFromPercentage(value string) (float64, error) {
	trimmedStr := strings.TrimSuffix(value, "%")
	number, err := strconv.ParseFloat(trimmedStr, 64)
	if err != nil {
		logger.GetLogger().Error("Cannot Parse String to float")
		return 0, err
	}

	// Round the number to one decimal place
	return float64(int(number*10+0.5)) / 10, nil
}
