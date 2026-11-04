package addon_test

import (
	"testing"

	"github.com/invopop/gobl.br.nfse/addon"
	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/rules"
	"github.com/invopop/gobl/tax"
	"github.com/stretchr/testify/assert"
)

func TestChargeValidation(t *testing.T) {
	tests := []struct {
		name   string
		charge *bill.Charge
		err    string
	}{
		{
			name: "tax estimate charge",
			charge: &bill.Charge{
				Key:         bill.ChargeKeyTax,
				Informative: true,
				Percent:     num.NewPercentage(65, 3),
				Ext:         tax.ExtensionsOf(cbc.CodeMap{addon.ExtKeyTaxEstimate: addon.TaxEstimateSimples}),
			},
		},
		{
			name:   "nil charge",
			charge: nil,
		},
		{
			name: "regular charge",
			charge: &bill.Charge{
				Amount: num.MakeAmount(100, 2),
			},
		},
		{
			name: "not informative",
			charge: &bill.Charge{
				Percent: num.NewPercentage(65, 3),
				Ext:     tax.ExtensionsOf(cbc.CodeMap{addon.ExtKeyTaxEstimate: addon.TaxEstimateSimples}),
			},
			err: "BR-NFSE-BILL-CHARGE-01",
		},
		{
			name: "invalid tax estimate code",
			charge: &bill.Charge{
				Informative: true,
				Percent:     num.NewPercentage(65, 3),
				Ext:         tax.ExtensionsOf(cbc.CodeMap{addon.ExtKeyTaxEstimate: "bogus"}),
			},
			err: "BR-NFSE-BILL-CHARGE-02",
		},
		{
			name: "tax estimate charge without percent",
			charge: &bill.Charge{
				Informative: true,
				Ext:         tax.ExtensionsOf(cbc.CodeMap{addon.ExtKeyTaxEstimate: addon.TaxEstimateSimples}),
			},
			err: "BR-NFSE-BILL-CHARGE-03",
		},
		{
			name: "informative without extension",
			charge: &bill.Charge{
				Informative: true,
				Percent:     num.NewPercentage(65, 3),
			},
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			err := rules.Validate(ts.charge, withAddonContext())
			if ts.err != "" {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), ts.err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
