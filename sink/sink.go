package sink

import (
	"net"

	"github.com/fatih/color"
)

func EnsureResolveOf(addresses ...string) error {
	for len(addresses) != 0 {
		for index, address := range addresses {
			// Ensure that we can resolve the DNS
			// Report any errors
			// Remove it from the slice
			if _, err := net.LookupHost(address); err != nil {
				switch err.(type) {
				case *net.DNSError:
					// Try again
				default:
					color.Red("Unable to resolve due to: %v", err)
					return err
				}
			} else {
				color.Green("Is able to resolve: %v", address)
				color.Green("Slice length:%v", len(addresses))
				if len(addresses) == 1 {
					addresses = append([]string{})
				} else {
					addresses = append(addresses[:index], addresses[index+1:]...)
				}
			}
		}
	}
	return nil
}
