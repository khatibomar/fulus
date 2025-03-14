// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = VED{}

// Bolívar Soberano currency type
type VED struct{}

func (VED) Code() string { return "VED" }

func (VED) Number() string { return "926" }

func (VED) Name() string { return "Bolívar Soberano" }

func (VED) MinorUnits() int { return 2 }

func (VED) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 100 locales
	case locale.EN_WS, locale.EN_GU, locale.EN_MU, locale.EN_CK, locale.EN_JM, locale.EN_CX, locale.EN_001, locale.EN_LR, locale.EN_NZ, locale.EN_MG, locale.EN_NR, locale.ZH_HANS_MO, locale.EN_AE, locale.EN_MO, locale.EN_UG, locale.EN_KN, locale.EN_GB, locale.EN_GY, locale.EN_MT, locale.EN_SB, locale.GD, locale.ZH_HANS_HK, locale.EN_AU, locale.EN_BM, locale.EN_FJ, locale.EN_IE, locale.EN_NF, locale.EN_PH, locale.EN_SL, locale.EN, locale.EN_PK, locale.EN_VC, locale.EN_VG, locale.EN_DG, locale.EN_MY, locale.ZH_HANS_MY, locale.EN_DM, locale.EN_MH, locale.EN_TO, locale.EN_TT, locale.ZH, locale.EN_BW, locale.EN_MP, locale.EN_SD, locale.EN_SS, locale.EN_CA, locale.EN_ZM, locale.EN_FM, locale.EN_GM, locale.EN_TC, locale.EN_UM, locale.EN_VI, locale.EN_SG, locale.EN_SH, locale.EN_GG, locale.EN_IL, locale.EN_MS, locale.EN_SZ, locale.EN_TV, locale.EN_NA, locale.EN_NU, locale.EN_TZ, locale.EN_GD, locale.EN_KI, locale.EN_LC, locale.EN_ER, locale.EN_CC, locale.EN_IM, locale.EN_IO, locale.ZH_HANS, locale.ZH_HANS_SG, locale.EN_CM, locale.EN_HK, locale.EN_TK, locale.EN_SX, locale.EN_ZW, locale.EN_KE, locale.EN_BZ, locale.EN_FK, locale.EN_MW, locale.EN_VU, locale.EN_BS, locale.EN_AG, locale.EN_AS, locale.EN_GH, locale.EN_KY, locale.EN_LS, locale.EN_PN, locale.EN_RW, locale.EN_BI, locale.EN_GI, locale.EN_BB, locale.EN_NG, locale.EN_PR, locale.EN_AI, locale.EN_CY, locale.EN_JE, locale.EN_PG, locale.EN_PW, locale.EN_SC:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GH, locale.FF_ADLM_NE, locale.FF_ADLM_SN, locale.FF_ADLM_GM, locale.FF_ADLM_CM, locale.FF_ADLM_SL, locale.FF_ADLM_GW, locale.FF_ADLM_BF, locale.FF_ADLM, locale.FF_ADLM_NG, locale.FF_ADLM_MR, locale.FF_ADLM_LR:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_CW, locale.NL_SR, locale.NL_BQ, locale.NL_SX, locale.NL, locale.EN_NL, locale.NL_AW, locale.NL_BE:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.EN_DK, locale.EN_DE, locale.EN_SI, locale.SC, locale.EN_BE:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_FI, locale.EN_SE:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.HI_LATN, locale.EN_IN:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EU:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "VED",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
