package utils

import (
	"fmt"
	"strings"
)

func FormatPrice(price float32) string {
    parts := strings.Split(fmt.Sprintf("%.2f", price), ".")
    intPart := parts[0]
    decPart := parts[1]

    // Insert commas in the integer part
    n := len(intPart)
    if n <= 3 {
        return fmt.Sprintf("$%s.%s", intPart, decPart)
    }

    result := ""
    for i := 0; i < n; i++ {
        if i > 0 && (n-i)%3 == 0 {
            result += ","
        }
        result += string(intPart[i])
    }

    return fmt.Sprintf("$%s.%s", result, decPart)
}