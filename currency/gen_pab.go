// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = PAB{}

// Balboa currency type
type PAB struct{}

func (PAB) Code() string { return "PAB" }

func (PAB) Number() string { return "590" }

func (PAB) Name() string { return "Balboa" }

func (PAB) MinorUnits() int { return 2 }

func (PAB) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 154 locales
	case locale.EN_GM, locale.EN_KE, locale.EN_UG, locale.TI_ER, locale.EN_GH, locale.EN_KN, locale.TH, locale.ES_GT, locale.EN_MY, locale.EN_SB, locale.EN_TT, locale.EN_VG, locale.TI, locale.EN_CX, locale.EN_IM, locale.EN_TK, locale.EN_BS, locale.FIL, locale.ZH_HANS, locale.EN_MU, locale.EN_BW, locale.EN_GU, locale.EN_NR, locale.EN_CK, locale.EN_MT, locale.YO_BJ, locale.EN_GD, locale.EN_NG, locale.EN_RW, locale.EN_001, locale.EN_BB, locale.AK, locale.EN_SD, locale.EN_NU, locale.MR, locale.EN_IO, locale.EN_WS, locale.ES_CU, locale.IG, locale.SO_DJ, locale.EN_TV, locale.ZH_HANT_HK, locale.AM, locale.YUE_HANS, locale.EN_MP, locale.EN_TO, locale.ES_NI, locale.ZH_HANT_MO, locale.EN_FK, locale.EN_AI, locale.EN_AU, locale.EN_IL, locale.EN_VU, locale.CHR, locale.EN_JM, locale.ES_MX, locale.EN_PH, locale.SO_KE, locale.EN_DM, locale.EN_PG, locale.EN_MW, locale.ZH_HANS_MY, locale.EN_GB, locale.EN_HK, locale.EN_VC, locale.EN_GY, locale.YUE_HANT, locale.EN_KY, locale.SO, locale.EN_CM, locale.EN_UM, locale.ZH_HANT, locale.EN_DG, locale.JA, locale.OR, locale.YUE, locale.EN_SX, locale.EN_SL, locale.SO_ET, locale.EN_SG, locale.ES_SV, locale.EN_CY, locale.PCM, locale.EE_TG, locale.EN_PW, locale.EN_SZ, locale.ES_419, locale.GA, locale.ES_US, locale.EN, locale.EN_FM, locale.EN_JE, locale.ZH, locale.EN_TC, locale.EN_TZ, locale.EN_CC, locale.ES_DO, locale.MS_SG, locale.EN_MH, locale.EN_PK, locale.EN_BM, locale.ZH_HANS_MO, locale.EE, locale.ZH_HANS_HK, locale.ZH_HANT_MY, locale.EN_ER, locale.ZH_HANS_SG, locale.EN_PN, locale.EN_VI, locale.EN_SH, locale.MS, locale.EN_CA, locale.EN_SC, locale.ES_HN, locale.UG, locale.KO_CN, locale.CY, locale.ZU, locale.KN, locale.EN_GI, locale.EN_NZ, locale.EN_ZW, locale.EN_LC, locale.EN_AG, locale.ES_PR, locale.EN_GG, locale.YO, locale.EN_LS, locale.ES_BZ, locale.EN_SS, locale.GA_GB, locale.EN_LR, locale.EN_NA, locale.EN_MO, locale.ES_BR, locale.EN_BZ, locale.ML, locale.EN_ZM, locale.KO, locale.EN_AE, locale.EN_AS, locale.EN_KI, locale.CEB, locale.EN_BI, locale.EN_IE, locale.EN_NF, locale.SI, locale.KO_KP, locale.YUE_HANT_CN, locale.EN_MG, locale.EN_PR, locale.GD, locale.EN_FJ, locale.EN_MS:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.SR, locale.SC, locale.IT, locale.DE, locale.RO_MD, locale.KU, locale.BS_LATN, locale.SR_CYRL, locale.DE_BE, locale.EN_DK, locale.GL, locale.ES_IC, locale.LLD, locale.EL_CY, locale.BS, locale.DA, locale.EL, locale.DSB, locale.DE_LU, locale.EN_SI, locale.SR_CYRL_ME, locale.SR_LATN_XK, locale.CA_FR, locale.ES_PH, locale.CA_IT, locale.ES, locale.SR_CYRL_XK, locale.DA_GL, locale.SR_CYRL_BA, locale.VI, locale.CA_AD, locale.EN_BE, locale.FR_LU, locale.CA, locale.CA_ES_VALENCIA, locale.AZ, locale.AZ_LATN, locale.SR_LATN_ME, locale.RO, locale.FR_MA, locale.SR_LATN, locale.SR_LATN_BA, locale.MK, locale.AST, locale.IT_SM, locale.ES_EA, locale.IT_VA, locale.IS, locale.EN_DE, locale.BS_CYRL, locale.HSB, locale.LB, locale.EL_POLYTON, locale.DE_IT:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 46 locales
	case locale.SK, locale.UZ_LATN, locale.PT_ST, locale.PL, locale.RU_UA, locale.PT_TL, locale.HU, locale.SQ_XK, locale.BR, locale.CV, locale.UZ_CYRL, locale.SQ, locale.PT_CH, locale.SQ_MK, locale.UZ, locale.UK, locale.BE_TARASK, locale.KK_CYRL, locale.PT_GW, locale.PT_MO, locale.HY, locale.KY, locale.TG, locale.KA, locale.PT_LU, locale.RU_KZ, locale.LV, locale.TT, locale.EN_FI, locale.PT_AO, locale.FR_CA, locale.RU_MD, locale.CS, locale.KK, locale.PT_CV, locale.RU, locale.RU_KG, locale.BE, locale.PT_PT, locale.PT_GQ, locale.TK, locale.EN_SE, locale.KK_KZ, locale.BG, locale.PT_MZ, locale.RU_BY:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_SC, locale.FR_WF, locale.FR_CD, locale.FR_VU, locale.FR_CM, locale.FR_RW, locale.FR_MU, locale.FR_NE, locale.FR_BE, locale.FR_MC, locale.FR_RE, locale.FR_GP, locale.FR_GF, locale.FR_BI, locale.FR_HT, locale.FR_GA, locale.FR_CH, locale.FR_MG, locale.FR_PM, locale.FR_BL, locale.FR_TG, locale.FR_TN, locale.FR_ML, locale.FR_MQ, locale.FR_GQ, locale.FR, locale.FR_KM, locale.FR_CI, locale.FR_TD, locale.FR_CG, locale.FR_NC, locale.FR_GN, locale.FR_MR, locale.FR_MF, locale.FR_BF, locale.FR_BJ, locale.FR_CF, locale.FR_SY, locale.FR_YT, locale.FR_SN, locale.FR_PF, locale.FR_DJ, locale.FR_DZ:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_JO, locale.AR_EH, locale.AR_IL, locale.AR_BH, locale.AR_AE, locale.AR_SS, locale.AR_QA, locale.AR_SY, locale.AR_KM, locale.AR_KW, locale.AR, locale.AR_SO, locale.AR_OM, locale.AR_ER, locale.AR_TD, locale.AR_DJ, locale.AR_YE, locale.AR_SD, locale.AR_PS, locale.AR_EG, locale.AR_IQ, locale.AR_SA:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.MN, locale.KOK, locale.EN_MV, locale.SW_UG, locale.MZN, locale.SW_KE, locale.HA_NE, locale.SD, locale.TA_SG, locale.TA_MY, locale.HA, locale.MI, locale.QU_EC, locale.HA_GH, locale.KOK_DEVA, locale.ES_PE, locale.QU, locale.SD_ARAB, locale.SW:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.EN_AT, locale.KGP, locale.ES_AR, locale.WO, locale.MS_BN, locale.SW_CD, locale.ES_CO, locale.YRL_CO, locale.PT, locale.YRL, locale.QU_BO, locale.YRL_VE, locale.ES_UY, locale.IA, locale.JV:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_GW, locale.FF_ADLM, locale.FF_ADLM_NG, locale.FF_ADLM_MR, locale.FF_ADLM_BF, locale.FF_ADLM_GH, locale.FF_ADLM_SL, locale.FF_ADLM_NE, locale.FF_ADLM_GM, locale.FF_ADLM_CM, locale.FF_ADLM_SN:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.GU, locale.PA_GURU, locale.TE, locale.BN_IN, locale.XNR, locale.DZ, locale.HI_LATN, locale.TA_LK, locale.EN_IN, locale.HI, locale.TA, locale.PA:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_SR, locale.NL_AW, locale.NL_SX, locale.ES_PY, locale.NL, locale.NL_BQ, locale.NL_CW, locale.NL_BE, locale.EN_NL:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.MS_ID, locale.EN_ID, locale.ES_BO, locale.TR, locale.ES_GQ, locale.ID, locale.TR_CY:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV_FI, locale.SV_AX, locale.ET, locale.SV, locale.KSH, locale.FI, locale.LT:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_LB, locale.AR_LY, locale.AR_MR, locale.AR_MA, locale.AR_TN:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.SL, locale.FO_DK, locale.FO, locale.HR_BA, locale.HR, locale.EU:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.EN_ZA, locale.AF_NA, locale.ES_CR, locale.AF:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_VE, locale.ES_CL, locale.ES_EC, locale.LO:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NO, locale.NB:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW_FR, locale.GSW_LI, locale.GSW:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.CCP, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.NE, locale.AS:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.CE, locale.EN_150:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "PNMB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "पि.ए.बि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_PA:
		return CurrencyFormatInfo{
			Symbol:           "B/.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "B/.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߔߊߓ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "PAB",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
