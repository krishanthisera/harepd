package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Delete lines
func DeleteLines(file, HapIp string) error {

	content, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		return readErr
	}

	m := regexp.MustCompile(fmt.Sprintf(".*\\s%s/\\d.*", HapIp))
	alter := m.ReplaceAllString(string(content), "")
	writeErr := ioutil.WriteFile(file, []byte(alter), 0600)

	if writeErr != nil {
		return writeErr
	}

	return nil
}

// Insert Line
func InjectLines(file, line string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(strings.TrimSuffix(line, "\n")); err != nil {
		return err
	}
	return nil
}

func TideUp(file string) error {
	content, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		return readErr
	}

	m := regexp.MustCompile("\n\n\n")
	alter := m.ReplaceAllString(string(content), "\n")
	writeErr := ioutil.WriteFile(file, []byte(alter), 0600)

	if writeErr != nil {
		return writeErr
	}
	return nil
}

//IpScraper will be filtered and return all the IPs in a stirng
func IpScrapper(str string) []string {
	m := regexp.MustCompile(`([0-9]{0,3}\.){3}[0-9]{0,3}`)
	return m.FindAllString(str, -1)
}
