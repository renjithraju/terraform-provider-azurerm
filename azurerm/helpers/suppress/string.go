package suppress

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func CaseDifference(_, old, new string, _ *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

func IgnoreIfNotSet(_, old, new string, _ *schema.ResourceData) bool {
	if new == "" {
		return true
	} else {
		return strings.EqualFold(old, new)
	}
}
