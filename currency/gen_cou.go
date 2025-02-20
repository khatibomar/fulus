// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = COU{}

// Unidad de Valor Real currency type
type COU struct{}

func (COU) Code() string { return "COU" }

func (COU) Number() string { return "970" }

func (COU) Name() string { return "Unidad de Valor Real" }

func (COU) MinorUnits() int { return 2 }

func (COU) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 131 locales
	case locale.EN_NA, locale.EN_TV, locale.EN_BB, locale.EN_BI, locale.EN_NU, locale.EN_SX, locale.EN_IO, locale.EN_NF, locale.EN_ZM, locale.EN_GM, locale.ZH_HANS, locale.ZH_HANS_MY, locale.EE_TG, locale.EN_MO, locale.YUE_HANT_CN, locale.EN_KE, locale.EN_FK, locale.EN_LS, locale.ES_MX, locale.ES_PA, locale.EN_WS, locale.EN_RW, locale.EN_BZ, locale.EN_KI, locale.GD, locale.EN_IM, locale.KO, locale.YUE_HANS, locale.EN_KN, locale.EN_PR, locale.EN_SG, locale.EN_PW, locale.EN_ER, locale.EN_LC, locale.ZH_HANS_MO, locale.ZH_HANT_MO, locale.EN_SS, locale.ES_HN, locale.TH, locale.EN_MT, locale.ZH_HANS_HK, locale.EN_AG, locale.EN_GB, locale.ES_GT, locale.UG, locale.EN_JM, locale.EN_SH, locale.EN_TO, locale.ZH_HANT_MY, locale.EN_CY, locale.ES_US, locale.EN_MG, locale.EN_CA, locale.EN_HK, locale.ES_BR, locale.EN_PG, locale.EN_IE, locale.ZH_HANT, locale.EN_001, locale.EN_GH, locale.EN_AE, locale.EN_AU, locale.EN_DM, locale.EN_UM, locale.EN_BM, locale.EN_IL, locale.EN_KY, locale.JA, locale.EN_AI, locale.ES_DO, locale.EN_SZ, locale.EN_TT, locale.EN_TC, locale.ES_BZ, locale.ML, locale.EN_GI, locale.EN_JE, locale.EN_MS, locale.ZH_HANS_SG, locale.EN_PK, locale.EN_SL, locale.EN_TZ, locale.EN_BW, locale.EN_PN, locale.EN_VG, locale.CY, locale.EN_PH, locale.EN_BS, locale.EN_AS, locale.EN_CX, locale.EN_FM, locale.EN_SB, locale.ZH_HANT_HK, locale.EN_CM, locale.EN_VU, locale.EE, locale.ES_CU, locale.YUE, locale.EN_MY, locale.KO_KP, locale.EN_NG, locale.EN_VI, locale.KO_CN, locale.YUE_HANT, locale.EN_GD, locale.EN_FJ, locale.EN_GY, locale.EN_ZW, locale.ZH, locale.EN_NR, locale.EN_SC, locale.EN_MH, locale.EN_MU, locale.ES_NI, locale.EN_GG, locale.ES_419, locale.EN_CC, locale.EN_VC, locale.EN_CK, locale.EN, locale.EN_DG, locale.EN_LR, locale.EN_MP, locale.EN_MW, locale.EN_UG, locale.ES_PR, locale.EN_TK, locale.EN_GU, locale.EN_NZ, locale.EN_SD, locale.ES_SV:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_NE, locale.FR_TD, locale.FR_CF, locale.FR_MC, locale.FR_MG, locale.FR_BF, locale.FR_MF, locale.FR_PF, locale.FR_MQ, locale.FR_RE, locale.FR_BE, locale.FR_NC, locale.FR_ML, locale.FR_GN, locale.FR, locale.FR_HT, locale.FR_CM, locale.FR_CH, locale.FR_RW, locale.FR_BJ, locale.FR_SC, locale.FR_GQ, locale.FR_MU, locale.FR_GF, locale.FR_KM, locale.FR_VU, locale.FR_BL, locale.FR_CD, locale.FR_SY, locale.FR_CG, locale.FR_TN, locale.FR_GP, locale.FR_WF, locale.FR_MR, locale.FR_TG, locale.FR_CI, locale.FR_PM, locale.FR_SN, locale.FR_DZ, locale.FR_YT, locale.FR_DJ, locale.FR_BI, locale.FR_GA:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 34 locales
	case locale.SR_LATN, locale.EN_SI, locale.ES, locale.SR_CYRL_ME, locale.FR_LU, locale.SR, locale.BS_LATN, locale.ES_PH, locale.SR_CYRL, locale.SR_CYRL_BA, locale.CA_IT, locale.CA_FR, locale.DE_LU, locale.CA_ES_VALENCIA, locale.DE_IT, locale.SR_CYRL_XK, locale.EN_BE, locale.EN_DE, locale.AST, locale.SR_LATN_ME, locale.BS, locale.DE_BE, locale.SR_LATN_BA, locale.VI, locale.BS_CYRL, locale.ES_EA, locale.FR_MA, locale.CA_AD, locale.SR_LATN_XK, locale.ES_IC, locale.EN_DK, locale.SC, locale.CA, locale.DE:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 26 locales
	case locale.UK, locale.PT_LU, locale.RU_MD, locale.PT_PT, locale.RU_UA, locale.PT_GQ, locale.CS, locale.PT_MO, locale.FR_CA, locale.PT_CH, locale.SK, locale.RU_KZ, locale.EN_SE, locale.PT_MZ, locale.BG, locale.PT_AO, locale.EN_FI, locale.LV, locale.HU, locale.PT_CV, locale.RU_KG, locale.RU_BY, locale.PT_ST, locale.PT_TL, locale.PT_GW, locale.RU:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_LR, locale.FF_ADLM_GM, locale.FF_ADLM_SN, locale.FF_ADLM_SL, locale.FF_ADLM_NE, locale.FF_ADLM_MR, locale.FF_ADLM_NG, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_GH, locale.FF_ADLM_BF:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_SR, locale.NL, locale.NL_SX, locale.NL_AW, locale.NL_BE, locale.ES_PY, locale.NL_BQ, locale.NL_CW, locale.EN_NL:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.PT, locale.EN_AT, locale.ES_UY, locale.ES_CO, locale.KGP, locale.YRL_CO, locale.YRL_VE, locale.ES_AR, locale.YRL:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.ES_BO, locale.ID, locale.TR_CY, locale.EN_ID, locale.ES_GQ, locale.TR:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.LT, locale.SV_AX, locale.SV, locale.FI, locale.SV_FI:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.ES_VE, locale.LO, locale.ES_EC, locale.ES_CL:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.HR, locale.SL, locale.EU, locale.HR_BA:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NB, locale.NO, locale.NN:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW_FR, locale.GSW, locale.GSW_LI:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.HI_LATN, locale.BN_IN, locale.EN_IN:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_MV, locale.ES_PE:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_CR, locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "COU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
