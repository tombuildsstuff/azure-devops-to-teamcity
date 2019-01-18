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

func FormatForTeamCity(buildNumber string, input []BuildOutput) error {

	fmt.Println(fmt.Sprintf("##teamcity[buildNumber '%s']", buildNumber))

	fmt.Println("##teamcity[testSuiteStarted name='Tests']")

	for _, test := range input {
		// from https://confluence.jetbrains.com/display/TCD18/Build+Script+Interaction+with+TeamCity#BuildScriptInteractionwithTeamCity-Escapedvalues
		stdout := strings.Replace(test.StdOut, "|", "||", -1)
		stdout = strings.Replace(stdout, "'", "|'", -1)
		stdout = strings.Replace(stdout, "[", "|[", -1)
		stdout = strings.Replace(stdout, "]", "|]", -1)
		stdout = strings.Replace(stdout, "\n", "|n", -1)

		fmt.Println(fmt.Sprintf("##teamcity[testStarted name='%s' captureStandardOutput='true']", test.TestName))
		fmt.Println(test.StdOut)

		//fmt.Println(fmt.Sprintf("##teamcity[testStdOut name='%s' out='%s']", test.TestName, stdout))

		if !test.Passed {
			fmt.Println(fmt.Sprintf("##teamcity[testFailed name='%s' message='Test ended in failure']", test.TestName))
		}

		fmt.Println(fmt.Sprintf("##teamcity[testFinished name='%s' duration='%d']", test.TestName, int64(test.Duration.Seconds()*float64(1000))))
	}

	fmt.Println("##teamcity[testSuiteFinished name='Tests']")

	return nil
}
