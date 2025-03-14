// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = OMR{}

// Rial Omani currency type
type OMR struct{}

func (OMR) Code() string { return "OMR" }

func (OMR) Number() string { return "512" }

func (OMR) Name() string { return "Rial Omani" }

func (OMR) MinorUnits() int { return 3 }

func (OMR) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 155 locales
	case locale.EN_LR, locale.ES_BZ, locale.EN_GY, locale.EN_UG, locale.ZH, locale.EN_KN, locale.EN_PW, locale.JA, locale.GA_GB, locale.MS_SG, locale.EN_FM, locale.EN_MU, locale.ES_DO, locale.MS, locale.EN_TO, locale.EE_TG, locale.EN_MW, locale.EN_MO, locale.EN_PR, locale.ZU, locale.EN_MY, locale.ES_CU, locale.YUE, locale.EN_NA, locale.EN_NR, locale.EN_TK, locale.CHR, locale.EN_LS, locale.EN_GM, locale.ES_419, locale.ES_PR, locale.CEB, locale.PCM, locale.EN_NU, locale.EN_AU, locale.ES_PA, locale.ZH_HANS, locale.AK, locale.EN_GU, locale.YUE_HANS, locale.EN_VG, locale.EN_AS, locale.EN_GD, locale.EN_PG, locale.ZH_HANS_MY, locale.KO, locale.EN_IE, locale.EN_JM, locale.TI_ER, locale.EN_BB, locale.EN_WS, locale.SO_DJ, locale.EN_MH, locale.EN_MS, locale.EN_MT, locale.EN_RW, locale.EN_CX, locale.EN_CY, locale.EN_IL, locale.ES_US, locale.ZH_HANS_MO, locale.ES_GT, locale.YO_BJ, locale.ZH_HANS_HK, locale.EN_FJ, locale.EN_PH, locale.EN_DG, locale.EN_SX, locale.EN_TC, locale.OR, locale.SO_ET, locale.EN_PK, locale.EE, locale.EN_KY, locale.EN_SL, locale.EN_UM, locale.YUE_HANT, locale.EN_BM, locale.EN_BZ, locale.EN_SD, locale.KO_CN, locale.EN_BS, locale.SI, locale.EN_GB, locale.EN_TZ, locale.YUE_HANT_CN, locale.EN_VI, locale.EN_DM, locale.EN_GG, locale.SO_KE, locale.EN_PN, locale.KN, locale.EN_SC, locale.ZH_HANT_HK, locale.ZH_HANT_MY, locale.EN_AE, locale.EN_HK, locale.KO_KP, locale.TH, locale.EN_FK, locale.EN_SH, locale.EN_SS, locale.ES_MX, locale.EN_IM, locale.EN_001, locale.ZH_HANT, locale.CY, locale.EN_MP, locale.EN_NF, locale.EN_ZW, locale.ES_HN, locale.ML, locale.ZH_HANS_SG, locale.EN_BW, locale.EN_TV, locale.FIL, locale.EN_NZ, locale.MR, locale.YO, locale.EN_BI, locale.EN_SZ, locale.EN_CA, locale.EN_ER, locale.EN_AG, locale.EN_MG, locale.ES_NI, locale.EN_CM, locale.GD, locale.IG, locale.ES_BR, locale.EN_CC, locale.EN_IO, locale.EN_SB, locale.EN_VC, locale.EN_VU, locale.EN_SG, locale.EN_GI, locale.EN_NG, locale.UG, locale.AM, locale.EN_AI, locale.EN_KI, locale.SO, locale.TI, locale.EN_KE, locale.EN_LC, locale.EN_ZM, locale.GA, locale.EN_TT, locale.EN_CK, locale.EN_GH, locale.ES_SV, locale.EN_JE, locale.EN, locale.ZH_HANT_MO:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.SC, locale.EN_DE, locale.FR_LU, locale.SR_CYRL_BA, locale.DE_LU, locale.CA_AD, locale.ES, locale.BS_CYRL, locale.SR_LATN_BA, locale.VI, locale.SR_CYRL, locale.IT_SM, locale.LLD, locale.SR_LATN_ME, locale.AST, locale.IS, locale.SR, locale.EL, locale.CA_FR, locale.MK, locale.DSB, locale.DA_GL, locale.FR_MA, locale.SR_LATN, locale.EL_POLYTON, locale.RO_MD, locale.GL, locale.EN_SI, locale.DA, locale.HSB, locale.DE, locale.SR_LATN_XK, locale.DE_BE, locale.ES_PH, locale.ES_IC, locale.EL_CY, locale.SR_CYRL_XK, locale.LB, locale.BS, locale.ES_EA, locale.CA_IT, locale.BS_LATN, locale.IT, locale.CA_ES_VALENCIA, locale.SR_CYRL_ME, locale.CA, locale.RO, locale.IT_VA, locale.KU, locale.AZ, locale.DE_IT, locale.EN_BE, locale.AZ_LATN, locale.EN_DK:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 45 locales
	case locale.PT_LU, locale.TG, locale.KK, locale.PT_CV, locale.PT_CH, locale.SQ_MK, locale.PT_PT, locale.HY, locale.PT_GQ, locale.KY, locale.RU_BY, locale.HU, locale.BG, locale.PT_AO, locale.EN_FI, locale.BE, locale.EN_SE, locale.RU_MD, locale.CS, locale.UZ, locale.RU_KG, locale.KK_KZ, locale.BR, locale.SQ, locale.TT, locale.PT_GW, locale.PL, locale.KA, locale.RU_UA, locale.FR_CA, locale.UZ_LATN, locale.LV, locale.PT_MO, locale.BE_TARASK, locale.KK_CYRL, locale.PT_TL, locale.PT_ST, locale.SK, locale.UK, locale.CV, locale.SQ_XK, locale.PT_MZ, locale.RU, locale.TK, locale.RU_KZ:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_NE, locale.FR_MF, locale.FR_SC, locale.FR_GF, locale.FR_GP, locale.FR_MU, locale.FR_KM, locale.FR_HT, locale.FR_SN, locale.FR_CF, locale.FR_CM, locale.FR_DJ, locale.FR_SY, locale.FR_ML, locale.FR_CG, locale.FR_YT, locale.FR_MQ, locale.FR_PF, locale.FR_DZ, locale.FR_RE, locale.FR_TG, locale.FR_CD, locale.FR_GN, locale.FR, locale.FR_WF, locale.FR_MG, locale.FR_VU, locale.FR_TD, locale.FR_BL, locale.FR_MC, locale.FR_GQ, locale.FR_BF, locale.FR_TN, locale.FR_CH, locale.FR_MR, locale.FR_NC, locale.FR_PM, locale.FR_BI, locale.FR_BJ, locale.FR_RW, locale.FR_GA, locale.FR_BE, locale.FR_CI:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_PS, locale.AR_DJ, locale.AR, locale.AR_TD, locale.AR_OM, locale.AR_ER, locale.AR_SA, locale.AR_JO, locale.AR_SD, locale.AR_QA, locale.AR_IL, locale.AR_KM, locale.AR_YE, locale.AR_SO, locale.AR_SY, locale.AR_EH, locale.AR_EG, locale.AR_KW, locale.AR_AE, locale.AR_BH, locale.AR_IQ, locale.AR_SS:
		return CurrencyFormatInfo{
			Symbol:           "ر.ع.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 20 locales
	case locale.BAL_LATN, locale.MZN, locale.MI, locale.HA_GH, locale.QU_EC, locale.SW_KE, locale.KOK, locale.TA_SG, locale.KOK_DEVA, locale.ES_PE, locale.HA_NE, locale.EN_MV, locale.SW, locale.TA_MY, locale.SD_ARAB, locale.MN, locale.QU, locale.HA, locale.SD, locale.SW_UG:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.IA, locale.PT, locale.JV, locale.KGP, locale.ES_AR, locale.SW_CD, locale.YRL, locale.ES_UY, locale.QU_BO, locale.ES_CO, locale.YRL_VE, locale.YRL_CO, locale.EN_AT, locale.MS_BN, locale.WO:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_BF, locale.FF_ADLM_SL, locale.FF_ADLM_MR, locale.FF_ADLM, locale.FF_ADLM_NG, locale.FF_ADLM_GW, locale.FF_ADLM_LR, locale.FF_ADLM_GH, locale.FF_ADLM_SN, locale.FF_ADLM_NE, locale.FF_ADLM_CM, locale.FF_ADLM_GM:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.PA, locale.TA_LK, locale.HI_LATN, locale.TA, locale.TE, locale.EN_IN, locale.BN_IN, locale.HI, locale.PA_GURU, locale.GU, locale.DZ, locale.XNR:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.EN_NL, locale.NL_BQ, locale.NL_CW, locale.ES_PY, locale.NL_AW, locale.NL_BE, locale.NL_SR, locale.NL_SX:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.ES_BO, locale.ES_GQ, locale.TR_CY, locale.EN_ID, locale.TR, locale.ID, locale.MS_ID:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSH, locale.LT, locale.SV_FI, locale.SV_AX, locale.FI, locale.SV, locale.ET:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_MR, locale.AR_TN, locale.AR_MA, locale.AR_LB, locale.AR_DZ, locale.AR_LY:
		return CurrencyFormatInfo{
			Symbol:           "ر.ع.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.EU, locale.HR_BA, locale.SL, locale.HR, locale.FO:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.EN_ZA, locale.ES_CR, locale.AF:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_VE, locale.ES_CL, locale.LO, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NB, locale.NO, locale.NN:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW, locale.GSW_LI, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.IT_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_150, locale.CE, locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "अ.एम.आर",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߏߡߙ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
