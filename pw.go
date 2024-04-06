package main
import (
    "fmt"
    "math/rand"
    "time"
    "unicode"
)
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789^$*[]{}.?-!@#%&/,><;_~"
var seededRand *rand.Rand = rand.New(
    rand.NewSource(time.Now().UnixNano()))
func StringWithCharset(length int, charset string) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}
func main() {
    for {
        s := StringWithCharset(8, charset)
        if hasSpecialChar(s) && hasNumber(s) && hasLowerCase(s) && hasUpperCase(s) {
            fmt.Println(s)
            break
        }
    }
}
func hasSpecialChar(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
            return true
        }
    }
    return false
}
func hasNumber(s string) bool {
    for _, r := range s {
        if unicode.IsNumber(r) {
            return true
        }
    }
    return false
}
func hasLowerCase(s string) bool {
    for _, r := range s {
        if unicode.IsLower(r) {
            return true
        }
    }
    return false
}
func hasUpperCase(s string) bool {
    for _, r := range s {
        if unicode.IsUpper(r) {
            return true
        }
    }
    return false
}