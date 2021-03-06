// Code generated by protoc-gen-validate
// source: cilium/api/nphds.proto
// DO NOT EDIT!!!

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on NetworkPolicyHosts with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NetworkPolicyHosts) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Policy

	for idx, item := range m.GetHostAddresses() {
		_, _ = idx, item

		if len(item) < 1 {
			return NetworkPolicyHostsValidationError{
				Field:  fmt.Sprintf("HostAddresses[%v]", idx),
				Reason: "value length must be at least 1 bytes",
			}
		}

	}

	return nil
}

// NetworkPolicyHostsValidationError is the validation error returned by
// NetworkPolicyHosts.Validate if the designated constraints aren't met.
type NetworkPolicyHostsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e NetworkPolicyHostsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkPolicyHosts.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = NetworkPolicyHostsValidationError{}
