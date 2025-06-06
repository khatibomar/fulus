// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BOV{}

// Mvdol currency type
type BOV struct{}

func (BOV) Code() string { return "BOV" }

func (BOV) Number() string { return "984" }

func (BOV) Name() string { return "Mvdol" }

func (BOV) MinorUnits() int { return 2 }

func (BOV) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 135 locales
	case locale.EN_KI, locale.EN_MW, locale.ES_BR, locale.EN_AI, locale.EN_BB, locale.EN_NR, locale.JA, locale.EN_LS, locale.EN_PG, locale.ES_GT, locale.EN_CM, locale.EN_KY, locale.EN_UG, locale.EN_UM, locale.EN_AG, locale.EN_BW, locale.ES_CU, locale.ES_MX, locale.ES_PR, locale.KO_CN, locale.YUE_HANT, locale.EN_KE, locale.UG, locale.EN_MY, locale.EN_NF, locale.EN_SC, locale.KO, locale.EN_IO, locale.EN_SG, locale.ES_US, locale.ES_PA, locale.GA, locale.EN_TO, locale.EN_DM, locale.EN_SD, locale.EN_ZM, locale.EN_AS, locale.EN_IM, locale.TH, locale.EN_PH, locale.EN_TT, locale.ES_DO, locale.EE_TG, locale.EN_TC, locale.EN_SX, locale.EN_WS, locale.ZH_HANS, locale.EN_CC, locale.EN_GS, locale.EN_MO, locale.EN_NZ, locale.EN_BS, locale.EN_NA, locale.EN_PN, locale.ML, locale.YUE_HANS, locale.YUE_HANT_CN, locale.EN_MP, locale.EN_DG, locale.EN_GB, locale.EN_MU, locale.EN_PR, locale.ZH_HANT, locale.EN_SL, locale.ES_BZ, locale.ZH_HANS_MY, locale.EN_BM, locale.EN_VI, locale.ZH_HANT_HK, locale.EN_MH, locale.EN_ZW, locale.EN_FK, locale.EN_RW, locale.EN_SZ, locale.GA_GB, locale.EN_BI, locale.EN_BZ, locale.KO_KP, locale.EN_MG, locale.EN_FJ, locale.EN_TZ, locale.EN, locale.EN_ER, locale.EN_MS, locale.EN_GH, locale.EN_CA, locale.EN_NU, locale.EN_TK, locale.ES_419, locale.EN_PK, locale.EN_SB, locale.EN_GM, locale.EN_MT, locale.EN_CY, locale.EN_GU, locale.EN_JM, locale.EN_TV, locale.GD, locale.EN_FM, locale.EN_VU, locale.EN_LC, locale.ZH_HANS_SG, locale.EN_LR, locale.ES_HN, locale.ZH, locale.ZH_HANS_HK, locale.EN_KN, locale.CY, locale.ZH_HANT_MO, locale.YUE_HANT_MO, locale.EN_AE, locale.ES_SV, locale.EN_NG, locale.EN_VC, locale.EN_JE, locale.EN_SH, locale.YUE, locale.EN_IE, locale.EN_SS, locale.EE, locale.EN_HK, locale.EN_PW, locale.EN_VG, locale.ZH_HANT_MY, locale.EN_001, locale.EN_AU, locale.EN_GI, locale.ES_NI, locale.EN_GG, locale.ZH_HANS_MO, locale.EN_GD, locale.EN_GY, locale.EN_CK, locale.EN_CX, locale.EN_IL:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 55 locales
	case locale.DE_BE, locale.EN_ES, locale.DSB, locale.IT_SM, locale.ES, locale.AZ, locale.SR_CYRL, locale.SR_CYRL_ME, locale.DE, locale.SR_CYRL_BA, locale.BS, locale.VI, locale.ES_IC, locale.IT, locale.SR_LATN_XK, locale.SC, locale.CA_AD, locale.EN_RO, locale.IT_VA, locale.LB, locale.RO_MD, locale.RO, locale.EN_BE, locale.FR_MA, locale.CA_ES_VALENCIA, locale.SR_CYRL_XK, locale.ES_EA, locale.EL_CY, locale.ES_PH, locale.DA, locale.DE_LU, locale.EN_IT, locale.BS_LATN, locale.AST, locale.EN_DE, locale.CA_IT, locale.EL, locale.FR_LU, locale.SR_LATN_BA, locale.DE_IT, locale.EN_DK, locale.AZ_LATN, locale.BS_CYRL, locale.DA_GL, locale.EN_SI, locale.SR_LATN, locale.CA, locale.GL, locale.SR_LATN_ME, locale.HSB, locale.EL_POLYTON, locale.SR, locale.IS, locale.CA_FR, locale.EN_PL:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_SY, locale.FR_MR, locale.FR_MU, locale.FR_PM, locale.FR_TD, locale.FR_CF, locale.FR_DJ, locale.FR_MG, locale.FR_BE, locale.FR_CD, locale.FR_NC, locale.FR_BF, locale.FR_CM, locale.FR_VU, locale.FR_PF, locale.EN_FR, locale.FR_SN, locale.FR_RE, locale.FR_SC, locale.FR_DZ, locale.FR_BL, locale.FR_GN, locale.FR_GQ, locale.FR_MC, locale.FR_GF, locale.FR_MQ, locale.FR_GA, locale.FR_BJ, locale.FR_RW, locale.FR_KM, locale.FR_CH, locale.FR_NE, locale.FR_GP, locale.FR_CG, locale.FR_MF, locale.FR_WF, locale.FR_TN, locale.FR_TG, locale.FR_YT, locale.FR_CI, locale.FR_BI, locale.FR_HT, locale.FR_ML, locale.FR:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 32 locales
	case locale.PT_MO, locale.PT_MZ, locale.FR_CA, locale.RU_KG, locale.PT_CV, locale.RU, locale.CS, locale.BG, locale.PT_AO, locale.PT_TL, locale.RU_KZ, locale.EN_HU, locale.PT_LU, locale.PT_ST, locale.EN_SE, locale.PT_PT, locale.EN_SK, locale.EN_NO, locale.SK, locale.PT_CH, locale.RU_UA, locale.EN_FI, locale.HT, locale.PT_GW, locale.HU, locale.PT_GQ, locale.EN_CZ, locale.UK, locale.EN_PT, locale.RU_MD, locale.PL, locale.RU_BY:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_DJ, locale.AR_KW, locale.AR_YE, locale.AR_JO, locale.AR_SD, locale.AR_TD, locale.AR_BH, locale.AR_QA, locale.AR_IL, locale.AR_KM, locale.AR_PS, locale.AR_IQ, locale.AR_ER, locale.AR_EH, locale.AR_OM, locale.AR_SY, locale.AR, locale.AR_EG, locale.AR_SO, locale.AR_SS, locale.AR_AE, locale.AR_SA:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM, locale.FF_ADLM_GH, locale.FF_ADLM_GM, locale.FF_ADLM_NE, locale.FF_ADLM_GW, locale.FF_ADLM_SN, locale.FF_ADLM_NG, locale.FF_ADLM_CM, locale.FF_ADLM_BF, locale.FF_ADLM_MR, locale.FF_ADLM_SL:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.ES_PY, locale.NL_BE, locale.NL_SR, locale.NL_CW, locale.EN_NL, locale.NL_SX, locale.NL_BQ, locale.NL_AW:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.YRL_VE, locale.ES_CO, locale.EN_AT, locale.YRL, locale.PT, locale.KGP, locale.ES_UY, locale.YRL_CO, locale.ES_AR:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MR, locale.AR_LY, locale.AR_TN, locale.AR_MA, locale.AR_DZ, locale.AR_LB:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ES_BO, locale.ES_GQ, locale.ID, locale.EN_ID, locale.TR_CY, locale.TR:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.PA, locale.HI_LATN, locale.EN_IN, locale.PA_GURU, locale.BN_IN:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.SV, locale.FI, locale.SV_AX, locale.LT, locale.SV_FI:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.RM, locale.GSW_FR, locale.GSW:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.HR, locale.SL, locale.HR_BA, locale.EU:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.LO, locale.ES_CL, locale.ES_EC, locale.ES_VE:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NO, locale.NB_SJ, locale.NB, locale.NN:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.DE_CH, locale.EN_CH:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_CR, locale.EN_ZA:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_PE, locale.EN_MV:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "BOV",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
