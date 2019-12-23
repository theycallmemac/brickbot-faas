package main
import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)
func main() {
	payload :=  []byte(strings.Join(os.Args[1:], ""))
	hashed := fmt.Sprintf("%X", sha1.Sum(payload))
	prefix := hashed[:5]

	c := http.Client{}

	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefix), nil)

	res, _ := c.Do(req)

	var bytesOut []byte
	if res.Body != nil {
		defer res.Body.Close()
		bytesOut, _ = ioutil.ReadAll(res.Body)
	}

	passwords := string(bytesOut)
	foundTimes, _ := findPassword(passwords, prefix, hashed)
	result := result{Found: foundTimes}
	output, _ := json.Marshal(result)
	fmt.Println(string(output))
}

type result struct {
	Found int `json:"found"`
}

func findPassword(passwords string, prefix string, hashed string) (int, error) {
	foundTimes := 0

	for _, passwordLine := range strings.Split(passwords, "\r\n") {
		if passwordLine != "" {
			parts := strings.Split(passwordLine, ":")

			if fmt.Sprintf("%s%s", prefix, parts[0]) == hashed {
				var convErr error
				foundTimes, convErr = strconv.Atoi(parts[1])
				if convErr != nil {
					return 0, fmt.Errorf(`Cannot convert value: "%s", error: "%s"\n`, parts[1], convErr)
				}
				break
			}
		}
	}
	return foundTimes, nil
}
