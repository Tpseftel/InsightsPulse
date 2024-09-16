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

// ConvertToFloat64Ptr converts various types of input to a pointer to a float64.
func ConvertToFloat64Ptr(input interface{}) *float64 {
	var result float64

	switch v := input.(type) {
	case float64:
		result = v
	case int:
		result = float64(v)
	case int8:
		result = float64(v)
	case int16:
		result = float64(v)
	case int32:
		result = float64(v)
	case int64:
		result = float64(v)
	case uint:
		result = float64(v)
	case uint8:
		result = float64(v)
	case uint16:
		result = float64(v)
	case uint32:
		result = float64(v)
	case uint64:
		result = float64(v)
	case float32:
		result = float64(v)

	case string:
		parsed, err := GetFloatFromPercentage(v)
		if err != nil {
			return nil // Return nil if the string cannot be converted to a float
		}
		result = parsed
	case nil:
		return nil // Handle null values
	default:
		return nil // Return nil for unsupported types
	}

	return &result
}

func StructToString(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var sb strings.Builder
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()
		sb.WriteString(fmt.Sprintf("%s: %v, ", field.Name, value))
	}

	result := sb.String()
	if len(result) > 0 {
		result = result[:len(result)-2] // Remove the trailing comma and space
	}
	return result
}
