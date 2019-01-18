package main

import (
	"strings"
	"testing"
)

func TestParseBuildLog(t *testing.T) {
	expected := []BuildOutput{
		{
			TestName: "TestAccAzureRMAppServiceCustomHostnameBinding",
			Passed:   true,
			Duration: 0.00,
			StdOut: `=== RUN   TestAccAzureRMAppServiceCustomHostnameBinding
=== RUN   TestAccAzureRMAppServiceCustomHostnameBinding/basic
=== RUN   TestAccAzureRMAppServiceCustomHostnameBinding/basic/basic
=== RUN   TestAccAzureRMAppServiceCustomHostnameBinding/basic/multiple
--- PASS: TestAccAzureRMAppServiceCustomHostnameBinding (0.00s)
    --- PASS: TestAccAzureRMAppServiceCustomHostnameBinding/basic (0.00s)
        --- SKIP: TestAccAzureRMAppServiceCustomHostnameBinding/basic/basic (0.00s)
            resource_arm_app_service_custom_hostname_binding_test.go:42: Skipping as "ARM_TEST_APP_SERVICE" is not specified
        --- SKIP: TestAccAzureRMAppServiceCustomHostnameBinding/basic/multiple (0.00s)
            resource_arm_app_service_custom_hostname_binding_test.go:80: Skipping as "ARM_TEST_APP_SERVICE" is not specified
PASS
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId",
			Passed:   false,
			Duration: 3163.27,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId
=== PAUSE TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId
=== CONT  TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId
2018/12/24 03:19:21 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:21 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:21 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

CREATE: azurerm_azuread_application.test
  application_id:    "" => "<computed>"
  homepage:          "" => "<computed>"
  identifier_uris.#: "" => "<computed>"
  name:              "" => "acctestspa7fada72e-96a9-ca6f-adb3-2e50f67975ac"
  reply_urls.#:      "" => "<computed>"
CREATE: azurerm_azuread_service_principal.test
  application_id: "" => "<computed>" (forces new resource)
  display_name:   "" => "<computed>"
CREATE: azurerm_azuread_service_principal_password.test
  end_date:             "" => "2020-01-01T01:02:03Z" (forces new resource)
  key_id:               "" => "1ff9cdbb-ad4f-e472-db6c-79efb769b546" (forces new resource)
  service_principal_id: "" => "<computed>" (forces new resource)
  start_date:           "" => "<computed>" (forces new resource)
  value:                "<sensitive>" => "<sensitive>" (forces new resource)

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryServicePrincipalPassword_customKeyId (3163.27s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryServicePrincipal_basic",
			Passed:   false,
			Duration: 7.46,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryServicePrincipal_basic
=== PAUSE TestAccAzureRMActiveDirectoryServicePrincipal_basic
=== CONT  TestAccAzureRMActiveDirectoryServicePrincipal_basic
2018/12/24 03:19:21 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:21 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:21 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

CREATE: azurerm_azuread_application.test
  application_id:    "" => "<computed>"
  homepage:          "" => "<computed>"
  identifier_uris.#: "" => "<computed>"
  name:              "" => "acctestspad4183b9b-acf8-4399-bc23-807a2743e6fd"
  reply_urls.#:      "" => "<computed>"
CREATE: azurerm_azuread_service_principal.test
  application_id: "" => "<computed>" (forces new resource)
  display_name:   "" => "<computed>"

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryServicePrincipal_basic (7.46s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryApplication_complete",
			Passed:   false,
			Duration: 6.95,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryApplication_complete
=== PAUSE TestAccAzureRMActiveDirectoryApplication_complete
=== CONT  TestAccAzureRMActiveDirectoryApplication_complete
2018/12/24 03:19:21 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:21 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:21 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

CREATE: azurerm_azuread_application.test
  application_id:             "" => "<computed>"
  homepage:                   "" => "https://homepage-302ea167-78da-4250-857e-fe496b6da7fc"
  identifier_uris.#:          "0" => "1"
  identifier_uris.0:          "" => "http://302ea167-78da-4250-857e-fe496b6da7fc.hashicorptest.com"
  name:                       "" => "acctest302ea167-78da-4250-857e-fe496b6da7fc"
  oauth2_allow_implicit_flow: "" => "true"
  reply_urls.#:               "0" => "1"
  reply_urls.0:               "" => "http://302ea167-78da-4250-857e-fe496b6da7fc.hashicorptest.com"

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryApplication_complete (6.95s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryApplication_availableToOtherTenants",
			Passed:   false,
			Duration: 7.00,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryApplication_availableToOtherTenants
=== PAUSE TestAccAzureRMActiveDirectoryApplication_availableToOtherTenants
=== CONT  TestAccAzureRMActiveDirectoryApplication_availableToOtherTenants
2018/12/24 03:19:21 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:21 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:21 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

 CREATE: azurerm_azuread_application.test
  application_id:             "" => "<computed>"
  available_to_other_tenants: "" => "true"
  homepage:                   "" => "<computed>"
  identifier_uris.#:          "0" => "1"
  identifier_uris.0:          "" => "https://196b15b4-ed17-4c2d-a05c-f5a6f48436b6.hashicorptest.com"
  name:                       "" => "acctest196b15b4-ed17-4c2d-a05c-f5a6f48436b6"
  reply_urls.#:               "" => "<computed>"

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryApplication_availableToOtherTenants (7.00s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryApplication_update",
			Passed:   false,
			Duration: 7.05,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryApplication_update
=== PAUSE TestAccAzureRMActiveDirectoryApplication_update
=== CONT  TestAccAzureRMActiveDirectoryApplication_update
2018/12/24 03:19:22 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:22 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:22 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

CREATE: azurerm_azuread_application.test
  application_id:    "" => "<computed>"
  homepage:          "" => "<computed>"
  identifier_uris.#: "" => "<computed>"
  name:              "" => "acctestf59cc989-2580-4cbd-a42a-3e285509346e"
  reply_urls.#:      "" => "<computed>"

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryApplication_update (7.05s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
		{
			TestName: "TestAccAzureRMActiveDirectoryApplication_basic",
			Passed:   false,
			Duration: 7.14,
			StdOut: `=== RUN   TestAccAzureRMActiveDirectoryApplication_basic
=== PAUSE TestAccAzureRMActiveDirectoryApplication_basic
=== CONT  TestAccAzureRMActiveDirectoryApplication_basic
2018/12/24 03:19:22 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:22 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:22 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:25 [WARN] Test: Step plan: DIFF:

CREATE: azurerm_azuread_application.test
  application_id:    "" => "<computed>"
  homepage:          "" => "<computed>"
  identifier_uris.#: "" => "<computed>"
  name:              "" => "acctestb11f4454-4966-430e-b9c8-a8ca700f7971"
  reply_urls.#:      "" => "<computed>"

STATE:

<no state>
2018/12/24 03:19:25 Testing if Service Principal / Client Certificate is applicable for Authentication..
2018/12/24 03:19:25 Testing if Service Principal / Client Secret is applicable for Authentication..
2018/12/24 03:19:25 Using Service Principal / Client Secret for Authentication
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalApplyPost, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [ERROR] root: eval: *terraform.EvalSequence, err: 1 error(s) occurred:

* azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
2018/12/24 03:19:28 [WARN] Test: Executing destroy step
2018/12/24 03:19:28 [WARN] Test: Step plan: DIFF:



STATE:

<no state>
--- FAIL: TestAccAzureRMActiveDirectoryApplication_basic (7.14s)
    testing.go:538: Step 0 error: Error applying: 1 error(s) occurred:

        * azurerm_azuread_application.test: 1 error(s) occurred:

        * azurerm_azuread_application.test: graphrbac.ApplicationsClient#Create: Failure responding to request: StatusCode=403 -- Original Error: autorest/azure: Service returned an error. Status=403 Code="Unknown" Message="Unknown service error" Details=[{"odata.error":{"code":"Authorization_RequestDenied","message":{"lang":"en","value":"Insufficient privileges to complete the operation."}}}]
FAIL
`,
		},
	}
	actual, err := ParseBuildLog("testdata/mini.log")
	if err != nil {
		t.Fatal(err)
	}

	if len(*actual) != len(expected) {
		t.Fatalf("Expected %d items but got %d", len(expected), len(*actual))
	}

	for i, _ := range expected {
		expectedValue := expected[i]
		actualValue := (*actual)[i]

		if expectedValue.TestName != actualValue.TestName {
			t.Fatalf("Expected Test Name to be %q but got %q", expectedValue.TestName, actualValue.TestName)
		}

		if expectedValue.Passed != actualValue.Passed {
			t.Fatalf("Expected Test Result to be %t but got %t for %q", expectedValue.Passed, actualValue.Passed, expectedValue.TestName)
		}

		if expectedValue.Duration != actualValue.Duration {
			t.Fatalf("Expected Test Name to be %.2f but got %.2f for %q", expectedValue.Duration, actualValue.Duration, expectedValue.TestName)
		}

		// compare the stdout's line by line by trimming the whitespace around then
		expectedLines := strings.Split(expectedValue.StdOut, "\n")
		actualLines := strings.Split(actualValue.StdOut, "\n")
		if len(expectedLines) != len(actualLines) {
			t.Fatalf("Expected Test %q to have %d lines but got %d lines", expectedValue.TestName, len(expectedLines), len(actualLines))
		}

		for lineNumber, _ := range expectedLines {
			expectedLine := expectedLines[lineNumber]
			actualLine := actualLines[lineNumber]
			if strings.TrimSpace(expectedLine) != strings.TrimSpace(actualLine) {
				t.Fatalf("Expected Test %q to have %q but got %q", expectedValue.TestName, expectedLine, actualLine)
			}
		}
	}
}
