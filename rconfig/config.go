// Create by Yale 2018/3/9 9:55
package rconfig

type RConfig struct {
	Replace []struct {
		Paths      []string `json:"paths"`
		RegReplace []struct {
			Regex       string `json:"regex"`
			Replacement string `json:"replacement"`
		} `json:"regReplace"`
		Recursive bool `json:"recursive"`
		Silent    bool `json:"silent"`
	} `json:"replace"`
}