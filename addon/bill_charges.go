package addon

import (
	"fmt"

	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/rules"
	"github.com/invopop/gobl/rules/is"
	"github.com/invopop/gobl/tax"
)

func billChargeRules() *rules.Set {
	return rules.For(new(bill.Charge),
		rules.When(is.Func("has tax estimate ext", chargeHasTaxEstimateExt),
			rules.Assert("01", "tax estimate charge must be informative",
				is.Expr("Informative"),
			),
			rules.Field("ext",
				rules.Assert("02", fmt.Sprintf("charge extension '%s' must be a valid code", ExtKeyTaxEstimate),
					tax.ExtensionHasValidCode(ExtKeyTaxEstimate),
				),
			),
			rules.Field("percent",
				rules.Assert("03", "tax estimate charge requires a percent", is.Present),
			),
		),
	)
}

func chargeHasTaxEstimateExt(val any) bool {
	c, ok := val.(*bill.Charge)
	return ok && c != nil && c.Ext.Has(ExtKeyTaxEstimate)
}
