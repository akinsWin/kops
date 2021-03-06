// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package devicefarm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDeviceFarm_CreateDevicePool() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.CreateDevicePoolInput{
		Name:       aws.String("Name"),               // Required
		ProjectArn: aws.String("AmazonResourceName"), // Required
		Rules: []*devicefarm.Rule{ // Required
			{ // Required
				Attribute: aws.String("DeviceAttribute"),
				Operator:  aws.String("RuleOperator"),
				Value:     aws.String("String"),
			},
			// More values...
		},
		Description: aws.String("Message"),
	}
	resp, err := svc.CreateDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateProject() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.CreateProjectInput{
		Name: aws.String("Name"), // Required
	}
	resp, err := svc.CreateProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateRemoteAccessSession() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.CreateRemoteAccessSessionInput{
		DeviceArn:  aws.String("AmazonResourceName"), // Required
		ProjectArn: aws.String("AmazonResourceName"), // Required
		Configuration: &devicefarm.CreateRemoteAccessSessionConfiguration{
			BillingMethod: aws.String("BillingMethod"),
		},
		Name: aws.String("Name"),
	}
	resp, err := svc.CreateRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateUpload() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.CreateUploadInput{
		Name:        aws.String("Name"),               // Required
		ProjectArn:  aws.String("AmazonResourceName"), // Required
		Type:        aws.String("UploadType"),         // Required
		ContentType: aws.String("ContentType"),
	}
	resp, err := svc.CreateUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteDevicePool() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.DeleteDevicePoolInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteProject() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.DeleteProjectInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteRemoteAccessSession() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.DeleteRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteRun() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.DeleteRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteUpload() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.DeleteUploadInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetAccountSettings() {
	svc := devicefarm.New(session.New())

	var params *devicefarm.GetAccountSettingsInput
	resp, err := svc.GetAccountSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevice() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetDeviceInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevice(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevicePool() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetDevicePoolInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevicePoolCompatibility() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetDevicePoolCompatibilityInput{
		DevicePoolArn: aws.String("AmazonResourceName"), // Required
		AppArn:        aws.String("AmazonResourceName"),
		TestType:      aws.String("TestType"),
	}
	resp, err := svc.GetDevicePoolCompatibility(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetJob() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetJobInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetOfferingStatus() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetOfferingStatusInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.GetOfferingStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetProject() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetProjectInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetRemoteAccessSession() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetRun() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetSuite() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetSuiteInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetSuite(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetTest() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetTestInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetTest(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetUpload() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.GetUploadInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_InstallToRemoteAccessSession() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.InstallToRemoteAccessSessionInput{
		AppArn:                 aws.String("AmazonResourceName"), // Required
		RemoteAccessSessionArn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.InstallToRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListArtifacts() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListArtifactsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		Type:      aws.String("ArtifactCategory"),   // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListArtifacts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListDevicePools() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListDevicePoolsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
		Type:      aws.String("DevicePoolType"),
	}
	resp, err := svc.ListDevicePools(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListDevices() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListDevicesInput{
		Arn:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListDevices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListJobs() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListJobsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListOfferingTransactions() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListOfferingTransactionsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListOfferingTransactions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListOfferings() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListOfferingsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListOfferings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListProjects() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListProjectsInput{
		Arn:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListProjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListRemoteAccessSessions() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListRemoteAccessSessionsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListRemoteAccessSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListRuns() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListRunsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListSamples() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListSamplesInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSamples(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListSuites() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListSuitesInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSuites(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListTests() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListTestsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListTests(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListUniqueProblems() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListUniqueProblemsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUniqueProblems(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListUploads() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ListUploadsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUploads(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_PurchaseOffering() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.PurchaseOfferingInput{
		OfferingId: aws.String("OfferingIdentifier"),
		Quantity:   aws.Int64(1),
	}
	resp, err := svc.PurchaseOffering(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_RenewOffering() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.RenewOfferingInput{
		OfferingId: aws.String("OfferingIdentifier"),
		Quantity:   aws.Int64(1),
	}
	resp, err := svc.RenewOffering(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ScheduleRun() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.ScheduleRunInput{
		DevicePoolArn: aws.String("AmazonResourceName"), // Required
		ProjectArn:    aws.String("AmazonResourceName"), // Required
		Test: &devicefarm.ScheduleRunTest{ // Required
			Type:   aws.String("TestType"), // Required
			Filter: aws.String("Filter"),
			Parameters: map[string]*string{
				"Key": aws.String("String"), // Required
				// More values...
			},
			TestPackageArn: aws.String("AmazonResourceName"),
		},
		AppArn: aws.String("AmazonResourceName"),
		Configuration: &devicefarm.ScheduleRunConfiguration{
			AuxiliaryApps: []*string{
				aws.String("AmazonResourceName"), // Required
				// More values...
			},
			BillingMethod:       aws.String("BillingMethod"),
			ExtraDataPackageArn: aws.String("AmazonResourceName"),
			Locale:              aws.String("String"),
			Location: &devicefarm.Location{
				Latitude:  aws.Float64(1.0), // Required
				Longitude: aws.Float64(1.0), // Required
			},
			NetworkProfileArn: aws.String("AmazonResourceName"),
			Radios: &devicefarm.Radios{
				Bluetooth: aws.Bool(true),
				Gps:       aws.Bool(true),
				Nfc:       aws.Bool(true),
				Wifi:      aws.Bool(true),
			},
		},
		Name: aws.String("Name"),
	}
	resp, err := svc.ScheduleRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_StopRemoteAccessSession() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.StopRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.StopRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_StopRun() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.StopRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.StopRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_UpdateDevicePool() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.UpdateDevicePoolInput{
		Arn:         aws.String("AmazonResourceName"), // Required
		Description: aws.String("Message"),
		Name:        aws.String("Name"),
		Rules: []*devicefarm.Rule{
			{ // Required
				Attribute: aws.String("DeviceAttribute"),
				Operator:  aws.String("RuleOperator"),
				Value:     aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_UpdateProject() {
	svc := devicefarm.New(session.New())

	params := &devicefarm.UpdateProjectInput{
		Arn:  aws.String("AmazonResourceName"), // Required
		Name: aws.String("Name"),
	}
	resp, err := svc.UpdateProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
