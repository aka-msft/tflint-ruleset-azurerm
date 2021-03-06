// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCdnEndpointInvalidOptimizationTypeRule checks the pattern is valid
type AzurermCdnEndpointInvalidOptimizationTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermCdnEndpointInvalidOptimizationTypeRule returns new rule with default attributes
func NewAzurermCdnEndpointInvalidOptimizationTypeRule() *AzurermCdnEndpointInvalidOptimizationTypeRule {
	return &AzurermCdnEndpointInvalidOptimizationTypeRule{
		resourceType:  "azurerm_cdn_endpoint",
		attributeName: "optimization_type",
		enum: []string{
			"GeneralWebDelivery",
			"GeneralMediaStreaming",
			"VideoOnDemandMediaStreaming",
			"LargeFileDownload",
			"DynamicSiteAcceleration",
		},
	}
}

// Name returns the rule name
func (r *AzurermCdnEndpointInvalidOptimizationTypeRule) Name() string {
	return "azurerm_cdn_endpoint_invalid_optimization_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCdnEndpointInvalidOptimizationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCdnEndpointInvalidOptimizationTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCdnEndpointInvalidOptimizationTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCdnEndpointInvalidOptimizationTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as optimization_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
