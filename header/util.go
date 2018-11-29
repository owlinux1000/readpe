package header

import (
    "fmt"
    "strings"
    "encoding/hex"
)


func decodeStringFromHex(i interface{}) string {
    ret, err := hex.DecodeString(fmt.Sprintf("%x", i))
    if err != nil {
        panic(err)
    }
    return string(ret)
}

func hexString(i interface{}) string {
    result := []string{}
    for _, v := range i.([]uint16) {
        result = append(result, fmt.Sprintf("%02x", v))
    }
    return strings.Join(result, " ")
}


func reverse(s string) (r string) {
    for i := len(s)-1; i >= 0; i-- {
        r += string(s[i])
    }
    return r
}
