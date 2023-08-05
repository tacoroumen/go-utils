package utils

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)
const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hypixelXPToLevel(lvl float64) int {
	var xp int = int(lvl)
	level := 1 + int(math.Floor(math.Sqrt(float64(2*xp+30625))/50)-2.5)
	return level
}

// functionName is a description of what the function does. 
//
// Charinput selects the values you want (Default(no input) = 0-9 a-z A-Z) -
// (1 = 0-9 a-z A-Z) -
// (2 = 0-9 a-z) -
// (3 = 0-9 A-Z) -
// (4 = a-z A-Z) -
// (5 = 0-9) -
// (6 = a-z) -
// (7 = A-Z) -
func generateRandomString(charinput int, length int) string {
	b := make([]byte, length)
	charset := ""
	switch charinput {
	case 1:
		const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 2:
		const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	case 3:
		const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 4:
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 5:
		const charset = "0123456789"
	case 6:
		const charset = "abcdefghijklmnopqrstuvwxyz"
	case 7:
		const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	default:
		const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// functionName is a description of what the function does. 
//
// str = the string that you want to insert the character in
//
// char = the character/string you want to insert in str in format of string ("")
//
// n = int of the interval to insert char
func insertCharacterIntoString(str string, char string, n int) string {
	var result string
	for i := 0; i < len(str); i += n {
		end := i + n
		if end > len(str) {
			end = len(str)
		}
		result += str[i:end]
		if end != len(str) {
			result += char
		}
	}
	return result
}

// GenerateRandomNumber generates a random number between min and max (inclusive).
func GenerateRandomNumber(min, max int) int {
    if max <= min {
        return min
    }
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min+1) + min
}

// IsPrime checks if a number is a prime number.
func IsPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

// ReverseString reverses a given string.
func ReverseString(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// SliceContains checks if a slice contains a specific element.
func SliceContains(slice []interface{}, item interface{}) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// GetMax returns the maximum value from a slice of integers.
func GetMax(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    max := nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    return max
}

// GetMin returns the minimum value from a slice of integers.
func GetMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    min := nums[0]
    for _, num := range nums {
        if num < min {
            min = num
        }
    }
    return min
}

// ShuffleSlice shuffles the elements of a slice randomly.
func ShuffleSlice(slice []interface{}) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(slice), func(i, j int) {
        slice[i], slice[j] = slice[j], slice[i]
    })
}

// FormatNumber adds comma as a thousand separator to a number.
func FormatNumber(num int) string {
    return fmt.Sprintf("%d", num)
}

// FormatFloat formats a float number to a string with the specified precision.
func FormatFloat(f float64, precision int) string {
    return fmt.Sprintf("%."+fmt.Sprintf("%d", precision)+"f", f)
}