// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule checks the pattern is valid
type AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule struct {
	resourceType  string
	attributeName string
	min           int
}

// NewAzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule returns new rule with default attributes
func NewAzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule() *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule {
	return &AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule{
		resourceType:  "azurerm_data_factory_integration_runtime_managed",
		attributeName: "number_of_nodes",
		min:           1,
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule) Name() string {
	return "azurerm_data_factory_integration_runtime_managed_invalid_number_of_nodes"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidNumberOfNodesRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val < r.min {
				runner.EmitIssueOnExpr(
					r,
					"number_of_nodes must be 1 or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
