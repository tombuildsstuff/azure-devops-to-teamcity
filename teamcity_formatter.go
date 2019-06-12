package main

import (
	"fmt"
	"strings"
)

/*
	https://confluence.jetbrains.com/display/TCD18/Build+Script+Interaction+with+TeamCity#BuildScriptInteractionwithTeamCity-ReportingTests
	// There must be only one testStdOut and one testStdErr message per test.

	##teamcity[testFailed name='package_or_namespace.ClassName.TestName' message='The number should be 20000' details='junit.framework.AssertionFailedError: expected:<20000> but was:<10000>|n|r    at junit.framework.Assert.fail(Assert.java:47)|n|r    at junit.framework.Assert.failNotEquals(Assert.java:280)|n|r...']

	##teamcity[testStarted name='className.testName']
	##teamcity[testStdOut name='className.testName' out='text']
	##teamcity[testStdErr name='className.testName' out='error text']
	##teamcity[testFinished name='className.testName' duration='50']
*/

func escapeString(input string) string {
	// from https://confluence.jetbrains.com/display/TCD18/Build+Script+Interaction+with+TeamCity#BuildScriptInteractionwithTeamCity-Escapedvalues
	out := strings.Replace(input, "|", "||", -1)
	out = strings.Replace(out, "'", "|'", -1)
	out = strings.Replace(out, "[", "|[", -1)
	out = strings.Replace(out, "]", "|]", -1)
	out = strings.Replace(out, "\n", "|n", -1)
	return out
}

func OutputFailureForTeamCity(message string) error {
	out := escapeString(message)
	fmt.Println(fmt.Sprintf("##teamcity[message text='Build Failed' errorDetails='%q' status='ERROR']", out))
	return nil
}

func FormatForTeamCity(buildNumber string, input []BuildOutput) error {
	fmt.Println(fmt.Sprintf("##teamcity[buildNumber '%s']", buildNumber))

	fmt.Println("##teamcity[testSuiteStarted name='Tests']")

	for _, test := range input {
		stdout := escapeString(test.StdOut)

		fmt.Println(fmt.Sprintf("##teamcity[testStarted name='%s' captureStandardOutput='true']", test.TestName))
		fmt.Println(stdout)

		//fmt.Println(fmt.Sprintf("##teamcity[testStdOut name='%s' out='%s']", test.TestName, stdout))

		if test.Skipped {
			fmt.Println(fmt.Sprintf("##teamcity[testIgnored name='%s' message='Test was Skipped']", test.TestName))
			continue
		}

		if !test.Passed {
			fmt.Println(fmt.Sprintf("##teamcity[testFailed name='%s' message='Test ended in failure']", test.TestName))
		} else {
			fmt.Println(fmt.Sprintf("##teamcity[testFinished name='%s' duration='%d']", test.TestName, int64(test.Duration.Seconds()*float64(1000))))
		}
	}

	fmt.Println("##teamcity[testSuiteFinished name='Tests']")

	return nil
}
