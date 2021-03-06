// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualNetworkGatewayInvalidSkuRule checks the pattern is valid
type AzurermVirtualNetworkGatewayInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualNetworkGatewayInvalidSkuRule returns new rule with default attributes
func NewAzurermVirtualNetworkGatewayInvalidSkuRule() *AzurermVirtualNetworkGatewayInvalidSkuRule {
	return &AzurermVirtualNetworkGatewayInvalidSkuRule{
		resourceType:  "azurerm_virtual_network_gateway",
		attributeName: "sku",
		enum: []string{
			"Basic",
			"HighPerformance",
			"Standard",
			"UltraPerformance",
			"VpnGw1",
			"VpnGw2",
			"VpnGw3",
			"VpnGw4",
			"VpnGw5",
			"VpnGw1AZ",
			"VpnGw2AZ",
			"VpnGw3AZ",
			"VpnGw4AZ",
			"VpnGw5AZ",
			"ErGw1AZ",
			"ErGw2AZ",
			"ErGw3AZ",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualNetworkGatewayInvalidSkuRule) Name() string {
	return "azurerm_virtual_network_gateway_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualNetworkGatewayInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualNetworkGatewayInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualNetworkGatewayInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualNetworkGatewayInvalidSkuRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
