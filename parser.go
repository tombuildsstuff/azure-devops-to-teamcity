package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type BuildOutput struct {
	TestName string
	Passed   bool
	Duration time.Duration
	StdOut   string
}

func ParseBuildLog(fileName string) (*[]BuildOutput, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outputs := make([]BuildOutput, 0)
	scanner := bufio.NewScanner(file)

	currentTestName := ""
	currentStdOut := ""
	var currentTestDuration *time.Duration

	reachedTests := false
	for scanner.Scan() {
		line := scanner.Text()
		lineRegex := "^((\\d){4}-(\\d){2}-(\\d){2}T(\\d){2}:(\\d){2}:(\\d){2}.(\\d){2,}Z)"
		matches, err := regexp.MatchString(lineRegex, line)
		if err != nil {
			panic(err)
		}

		if !matches {
			continue
		}

		if len(line) < 30 {
			if !reachedTests {
				continue
			} else {
				// append the empty line and move on
				currentStdOut += "\n"
				continue
			}
		}

		lineWithoutDate := line[29:]

		// if we havne't found the start of the tests then continue until we find one
		// === RUN   TestAccAzureRMAppServiceCustomHostnameBinding
		if !reachedTests && !strings.HasPrefix(lineWithoutDate, "=== RUN") {
			continue
		}

		// handles nested tests
		if currentTestName == "" && strings.HasPrefix(lineWithoutDate, "=== RUN") {
			reachedTests = true
			// parse the test name out of the line
			currentTestName = strings.Replace(lineWithoutDate, "=== RUN   ", "", 1)
		}

		// check if the current test has crashed
		if currentTestName != "" {
			match, err := regexp.MatchString("exit status (\\d+){1,}", lineWithoutDate)
			if err != nil {
				return nil, err
			}

			if match {
				currentTestName = ""
				continue
			}
		}

		if currentTestName != "" {
			currentStdOut += fmt.Sprintf("%s\n", lineWithoutDate)
		}

		// if it's a test failure line e.g.
		// --- FAIL: TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId (7.46s)
		// then we can capture the stderr

		// if this line contains the run time for the test then parse that out
		// --- [PASS|FAIL]{4}: [A-Za-z0-9_]{1,} \((\d){1,}.(\d){1,}s\)
		durationRegex := fmt.Sprintf("--- [PASS|FAIL]{4}: [A-Za-z0-9_]{1,} \\((\\d+){1,}")
		match, err := regexp.MatchString(durationRegex, lineWithoutDate)
		if match {

			//  TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId (3163.27s)

			// parse the duration out
			name := lineWithoutDate
			name = strings.Replace(name, "--- FAIL:", "", 1)
			name = strings.Replace(name, "--- PASS:", "", 1)
			name = strings.Replace(name, currentTestName, "", 1)
			name = strings.Replace(name, "(", "", 1)
			name = strings.Replace(name, ")", "", 1)
			name = strings.TrimSpace(name)

			log.Printf("[DEBUG] Duration is %q", name)

			duration, err := time.ParseDuration(name)
			if err != nil {
				return nil, fmt.Errorf("Error parsing duration for %s: %s", currentTestName, err)
			}

			currentTestDuration = &duration
		}

		// the end of a Test either ends with a PASS or a FAIL
		if lineWithoutDate == "FAIL" || lineWithoutDate == "PASS" {
			currentTestPassed := lineWithoutDate == "PASS"
			output := BuildOutput{
				TestName: currentTestName,
				StdOut:   currentStdOut,
				Passed:   currentTestPassed,
				Duration: *currentTestDuration,
			}
			outputs = append(outputs, output)

			// then reset the values
			currentTestName = ""
			currentStdOut = ""
			currentTestDuration = nil
		}
	}

	return &outputs, nil
}
