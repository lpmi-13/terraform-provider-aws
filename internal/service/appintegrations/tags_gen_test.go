// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package appintegrations_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	tfappintegrations "github.com/hashicorp/terraform-provider-aws/internal/service/appintegrations"
)

func expectFullResourceTags(ctx context.Context, resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfappintegrations.ServicePackage(ctx), resourceAddress, knownValue)
}

func expectFullDataSourceTags(ctx context.Context, resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfappintegrations.ServicePackage(ctx), resourceAddress, knownValue)
}
