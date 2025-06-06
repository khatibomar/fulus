// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = XSU{}

// Sucre currency type
type XSU struct{}

func (XSU) Code() string { return "XSU" }

func (XSU) Number() string { return "994" }

func (XSU) Name() string { return "Sucre" }

func (XSU) MinorUnits() int { return 2 }

func (XSU) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 114 locales
	case locale.YUE_HANT, locale.ZH_HANT, locale.EN_CA, locale.EN_NZ, locale.EN_PR, locale.EN_SS, locale.EN_UM, locale.YUE_HANT_MO, locale.EN_BB, locale.EN_BS, locale.EN_IL, locale.EN_KY, locale.EN_SB, locale.EN_JE, locale.EN_MG, locale.EN_MT, locale.EN_NG, locale.EN_PK, locale.EN_DG, locale.EN_MU, locale.EN_TC, locale.EN_LS, locale.EN_PG, locale.EN_BW, locale.EN_GG, locale.EN_GM, locale.EN_IM, locale.EN_LR, locale.ZH_HANT_HK, locale.EN_SL, locale.EN_FK, locale.EN_GB, locale.EN_AS, locale.EN_TZ, locale.TH, locale.ZH_HANS_SG, locale.GD, locale.JA, locale.EN_001, locale.EN_ER, locale.EN_KN, locale.EN_RW, locale.EN_VI, locale.EN_MW, locale.EN_SG, locale.EN_SH, locale.EN_TK, locale.ZH_HANT_MY, locale.EN_GY, locale.EN_HK, locale.EN_NR, locale.EN_SC, locale.ZH_HANS_MY, locale.EN_GU, locale.EN_KI, locale.ZH_HANS_HK, locale.ZH_HANT_MO, locale.EN_AI, locale.EN_BM, locale.EN_BZ, locale.EN_GI, locale.EN_PH, locale.EN_WS, locale.EN_IE, locale.EN_MO, locale.ZH_HANS, locale.EN_CX, locale.EN_DM, locale.EN_FM, locale.EN_KE, locale.EN_LC, locale.EN_SZ, locale.EN_AG, locale.EN_MH, locale.YUE_HANT_CN, locale.ZH, locale.EN_SD, locale.EN, locale.EN_CK, locale.EN_VG, locale.EN_GD, locale.EN_MP, locale.EN_MS, locale.EN_VC, locale.YUE_HANS, locale.EN_AE, locale.EN_PN, locale.ZH_HANS_MO, locale.EN_JM, locale.EN_MY, locale.EN_VU, locale.EN_CY, locale.EN_FJ, locale.EN_GS, locale.EN_ZW, locale.YUE, locale.EN_NU, locale.EN_SX, locale.EN_TT, locale.EN_ZM, locale.UG, locale.EN_UG, locale.CY, locale.EN_CM, locale.EN_TO, locale.EN_AU, locale.EN_BI, locale.EN_GH, locale.EN_NF, locale.EN_IO, locale.EN_NA, locale.EN_PW, locale.EN_TV, locale.EN_CC:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.DE_IT, locale.DE_LU, locale.EN_DK, locale.EN_PL, locale.EN_DE, locale.DE, locale.EN_ES, locale.AST, locale.EN_SI, locale.SC, locale.DE_BE, locale.EN_RO, locale.EN_BE, locale.EN_IT:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.EN_FI, locale.CS, locale.EN_SE, locale.SK, locale.EN_HU, locale.EN_SK, locale.EN_PT, locale.EN_NO, locale.EN_CZ:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_BQ, locale.NL_AW, locale.NL_BE, locale.NL, locale.NL_SR, locale.NL_CW, locale.EN_NL, locale.NL_SX:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.SV, locale.SV_FI, locale.FI, locale.SV_AX, locale.LT:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NO, locale.NN, locale.NB, locale.NB_SJ:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.EN_ID, locale.TR_CY, locale.TR:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.HR, locale.EU, locale.HR_BA:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.EN_CH:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_IN, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "XSU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
