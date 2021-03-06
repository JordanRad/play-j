// Code generated by goa v3.7.0, DO NOT EDIT.
//
// payment HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	paymentc "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/http/payment/client"
	subscriptionc "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/http/subscription/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `payment (get-account-payments|create-account-payment)
subscription get-account-subscription-status
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` payment get-account-payments --limit 6377591863020585586 --auth "Placeat voluptatum recusandae."` + "\n" +
		os.Args[0] + ` subscription get-account-subscription-status --auth "Et et."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		paymentFlags = flag.NewFlagSet("payment", flag.ContinueOnError)

		paymentGetAccountPaymentsFlags     = flag.NewFlagSet("get-account-payments", flag.ExitOnError)
		paymentGetAccountPaymentsLimitFlag = paymentGetAccountPaymentsFlags.String("limit", "REQUIRED", "")
		paymentGetAccountPaymentsAuthFlag  = paymentGetAccountPaymentsFlags.String("auth", "REQUIRED", "")

		paymentCreateAccountPaymentFlags    = flag.NewFlagSet("create-account-payment", flag.ExitOnError)
		paymentCreateAccountPaymentAuthFlag = paymentCreateAccountPaymentFlags.String("auth", "REQUIRED", "")

		subscriptionFlags = flag.NewFlagSet("subscription", flag.ContinueOnError)

		subscriptionGetAccountSubscriptionStatusFlags    = flag.NewFlagSet("get-account-subscription-status", flag.ExitOnError)
		subscriptionGetAccountSubscriptionStatusAuthFlag = subscriptionGetAccountSubscriptionStatusFlags.String("auth", "REQUIRED", "")
	)
	paymentFlags.Usage = paymentUsage
	paymentGetAccountPaymentsFlags.Usage = paymentGetAccountPaymentsUsage
	paymentCreateAccountPaymentFlags.Usage = paymentCreateAccountPaymentUsage

	subscriptionFlags.Usage = subscriptionUsage
	subscriptionGetAccountSubscriptionStatusFlags.Usage = subscriptionGetAccountSubscriptionStatusUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "payment":
			svcf = paymentFlags
		case "subscription":
			svcf = subscriptionFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "payment":
			switch epn {
			case "get-account-payments":
				epf = paymentGetAccountPaymentsFlags

			case "create-account-payment":
				epf = paymentCreateAccountPaymentFlags

			}

		case "subscription":
			switch epn {
			case "get-account-subscription-status":
				epf = subscriptionGetAccountSubscriptionStatusFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "payment":
			c := paymentc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-account-payments":
				endpoint = c.GetAccountPayments()
				data, err = paymentc.BuildGetAccountPaymentsPayload(*paymentGetAccountPaymentsLimitFlag, *paymentGetAccountPaymentsAuthFlag)
			case "create-account-payment":
				endpoint = c.CreateAccountPayment()
				data, err = paymentc.BuildCreateAccountPaymentPayload(*paymentCreateAccountPaymentAuthFlag)
			}
		case "subscription":
			c := subscriptionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-account-subscription-status":
				endpoint = c.GetAccountSubscriptionStatus()
				data, err = subscriptionc.BuildGetAccountSubscriptionStatusPayload(*subscriptionGetAccountSubscriptionStatusAuthFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// paymentUsage displays the usage of the payment command and its subcommands.
func paymentUsage() {
	fmt.Fprintf(os.Stderr, `Payment operations
Usage:
    %[1]s [globalflags] payment COMMAND [flags]

COMMAND:
    get-account-payments: GetAccountPayments implements getAccountPayments.
    create-account-payment: CreateAccountPayment implements createAccountPayment.

Additional help:
    %[1]s payment COMMAND --help
`, os.Args[0])
}
func paymentGetAccountPaymentsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] payment get-account-payments -limit INT -auth STRING

GetAccountPayments implements getAccountPayments.
    -limit INT: 
    -auth STRING: 

Example:
    %[1]s payment get-account-payments --limit 6377591863020585586 --auth "Placeat voluptatum recusandae."
`, os.Args[0])
}

func paymentCreateAccountPaymentUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] payment create-account-payment -auth STRING

CreateAccountPayment implements createAccountPayment.
    -auth STRING: 

Example:
    %[1]s payment create-account-payment --auth "Autem expedita laudantium earum."
`, os.Args[0])
}

// subscriptionUsage displays the usage of the subscription command and its
// subcommands.
func subscriptionUsage() {
	fmt.Fprintf(os.Stderr, `Service is the subscription service interface.
Usage:
    %[1]s [globalflags] subscription COMMAND [flags]

COMMAND:
    get-account-subscription-status: GetAccountSubscriptionStatus implements getAccountSubscriptionStatus.

Additional help:
    %[1]s subscription COMMAND --help
`, os.Args[0])
}
func subscriptionGetAccountSubscriptionStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] subscription get-account-subscription-status -auth STRING

GetAccountSubscriptionStatus implements getAccountSubscriptionStatus.
    -auth STRING: 

Example:
    %[1]s subscription get-account-subscription-status --auth "Et et."
`, os.Args[0])
}
