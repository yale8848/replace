// Create by Yale 2018/3/9 9:45
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"replace/rconfig"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getFiles(files []string, file string, recursive bool) []string {

	f, err := os.Open(file)
	checkError(err)
	if s, err := f.Stat(); err == nil {
		if s.IsDir() {
			fi, err := ioutil.ReadDir(file)
			checkError(err)
			for _, f := range fi {
				if f.IsDir() {
					if recursive {
						return getFiles(files, filepath.Join(file, f.Name()), recursive)
					}
				} else {
					files = append(files, filepath.Join(file, f.Name()))
				}
			}

		} else {
			files = append(files, file)
		}
	}
	f.Close()

	return files

}

func main() {

	config := flag.String("c", "config.json", "config.json")

	flag.Parse()

	dat, err := ioutil.ReadFile(*config)
	checkError(err)

	rc := &rconfig.RConfig{}
	err = json.Unmarshal(dat, &rc)
	checkError(err)

	for _, c := range rc.Replace {

		for _, p := range c.Paths {

			f, err := os.Open(p)
			checkError(err)
			files := make([]string, 0)

			files = getFiles(files, p, c.Recursive)

			for _, f := range files {
				dat, err := ioutil.ReadFile(f)
				checkError(err)
				data := string(dat)

				rp := data
				for _, rpls := range c.RegReplace {
					m, _ := regexp.MatchString(rpls.Regex, data)
					if !m {
						continue
					}
					r, _ := regexp.Compile(rpls.Regex)
					rp = r.ReplaceAllString(rp, rpls.Replacement)
					if !c.Silent {
						fmt.Printf("[replacing]:%v: %v-->%v\r\n", f, rpls.Regex, rpls.Replacement)
					}
				}
				if rp != data {
					fmt.Println("[replaced]:" + f)
					err = ioutil.WriteFile(f, []byte(rp), 0666)
					checkError(err)
				}
			}
			f.Close()

		}
	}

}
