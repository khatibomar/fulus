// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = UZS{}

// Uzbekistan Sum currency type
type UZS struct{}

func (UZS) Code() string { return "UZS" }

func (UZS) Number() string { return "860" }

func (UZS) Name() string { return "Uzbekistan Sum" }

func (UZS) MinorUnits() int { return 2 }

func (UZS) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 157 locales
	case locale.EN_PW, locale.EN_SS, locale.SO, locale.EN_CC, locale.EN_SB, locale.EN_TV, locale.YUE_HANT, locale.EN_MU, locale.EN_ZM, locale.KN, locale.EN_NA, locale.ES_DO, locale.GA_GB, locale.ZH_HANT_MY, locale.ES_BR, locale.ES_CU, locale.EN_LS, locale.ES_BZ, locale.ZH_HANT_HK, locale.EN_SL, locale.EN_NR, locale.EN_SG, locale.ZU, locale.KO_KP, locale.ZH_HANS_MY, locale.ML, locale.EN_CA, locale.ZH_HANS_SG, locale.EN_HK, locale.EN_MG, locale.EN_CY, locale.EN_WS, locale.ES_PA, locale.EN_VI, locale.ES_PR, locale.JA, locale.EN_GI, locale.TI, locale.EN_TT, locale.PCM, locale.ZH, locale.ES_US, locale.GA, locale.MR, locale.EN_KN, locale.EN_UM, locale.TH, locale.EN_GG, locale.IG, locale.ZH_HANS_HK, locale.AK, locale.EN_AU, locale.EN_MT, locale.EN_SC, locale.EN_AS, locale.EN_GB, locale.EN_VU, locale.EN_AE, locale.SI, locale.EN_FM, locale.EN_IL, locale.EN_LR, locale.EN_PK, locale.YO, locale.EN_FK, locale.EN_TZ, locale.EN_SH, locale.EN_BS, locale.EN_KY, locale.CY, locale.EN_SZ, locale.EN_MY, locale.EN_AI, locale.KO_CN, locale.ZH_HANS, locale.EN_MW, locale.EN_GU, locale.ZH_HANT_MO, locale.EN_NU, locale.EN_PR, locale.EN_SD, locale.ES_SV, locale.EN_MO, locale.EN_BI, locale.EN_MP, locale.YO_BJ, locale.YUE_HANT_MO, locale.EN_DM, locale.EN_IO, locale.EN_PN, locale.TI_ER, locale.EN_LC, locale.AM, locale.EN_KI, locale.ES_GT, locale.ES_HN, locale.ES_MX, locale.GD, locale.ZH_HANS_MO, locale.EN_DG, locale.YUE_HANS, locale.EN_GD, locale.EE_TG, locale.EN_AG, locale.EN_GM, locale.EE, locale.EN_GY, locale.EN_MS, locale.EN_PG, locale.EN_SX, locale.EN_TC, locale.EN_TO, locale.CEB, locale.EN_NZ, locale.EN_NF, locale.ES_419, locale.EN_NG, locale.EN_FJ, locale.EN_IE, locale.FIL, locale.SO_KE, locale.EN_001, locale.SO_DJ, locale.EN_PH, locale.EN_TK, locale.OR, locale.EN_BB, locale.EN_BM, locale.YUE, locale.SO_ET, locale.EN_ZW, locale.UG, locale.EN_IM, locale.EN_VG, locale.KO, locale.CHR, locale.EN_ER, locale.EN_KE, locale.MS_SG, locale.EN_VC, locale.EN_RW, locale.ES_NI, locale.EN_CX, locale.EN_GH, locale.EN_MH, locale.EN_BZ, locale.EN_CM, locale.EN_BW, locale.ZH_HANT, locale.EN_UG, locale.YUE_HANT_CN, locale.EN_GS, locale.EN_CK, locale.EN, locale.EN_JE, locale.EN_JM, locale.MS:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 58 locales
	case locale.KU, locale.HSB, locale.IT, locale.SC, locale.SR, locale.CA, locale.RO, locale.EN_DK, locale.SR_CYRL, locale.DA, locale.CA_FR, locale.EN_BE, locale.FR_MA, locale.RO_MD, locale.CA_IT, locale.VI, locale.AZ_LATN, locale.BS_CYRL, locale.EL_CY, locale.ES_EA, locale.DE_BE, locale.DE_LU, locale.GL, locale.DSB, locale.ES_PH, locale.EL_POLYTON, locale.LB, locale.SR_CYRL_XK, locale.EN_ES, locale.FR_LU, locale.LLD, locale.BS, locale.EN_PL, locale.EN_SI, locale.EN_DE, locale.DE, locale.SR_LATN, locale.SR_LATN_ME, locale.MK, locale.AST, locale.ES, locale.EL, locale.DE_IT, locale.IT_VA, locale.EN_IT, locale.IT_SM, locale.CA_AD, locale.SR_LATN_BA, locale.DA_GL, locale.ES_IC, locale.BS_LATN, locale.EN_RO, locale.SR_LATN_XK, locale.AZ, locale.SR_CYRL_BA, locale.CA_ES_VALENCIA, locale.IS, locale.SR_CYRL_ME:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 49 locales
	case locale.PT_MZ, locale.HU, locale.KK, locale.SQ_XK, locale.EN_FI, locale.PT_ST, locale.EN_HU, locale.HY, locale.SK, locale.EN_NO, locale.RU_MD, locale.BE, locale.PT_TL, locale.EN_PT, locale.SQ_MK, locale.TT, locale.LV, locale.CV, locale.RU_KZ, locale.KK_KZ, locale.SQ, locale.EN_SE, locale.EN_SK, locale.HT, locale.RU_BY, locale.PT_CH, locale.TG, locale.KA, locale.PT_LU, locale.RU_UA, locale.UK, locale.PT_GQ, locale.RU_KG, locale.PT_PT, locale.KY, locale.FR_CA, locale.CS, locale.RU, locale.PL, locale.EN_CZ, locale.PT_MO, locale.BG, locale.PT_AO, locale.PT_GW, locale.BR, locale.KK_CYRL, locale.PT_CV, locale.TK, locale.BE_TARASK:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_CD, locale.FR_MQ, locale.FR_DJ, locale.FR_SN, locale.FR_BF, locale.FR_MC, locale.FR_NC, locale.FR_CM, locale.FR_MG, locale.EN_FR, locale.FR_VU, locale.FR_TG, locale.FR_PM, locale.FR_GN, locale.FR_HT, locale.FR, locale.FR_CH, locale.FR_BI, locale.FR_BJ, locale.FR_ML, locale.FR_GQ, locale.FR_SC, locale.FR_MU, locale.FR_RW, locale.FR_RE, locale.FR_TN, locale.FR_WF, locale.FR_PF, locale.FR_TD, locale.FR_SY, locale.FR_GA, locale.FR_YT, locale.FR_MR, locale.FR_GP, locale.FR_KM, locale.FR_CG, locale.FR_BE, locale.FR_DZ, locale.FR_BL, locale.FR_GF, locale.FR_CF, locale.FR_NE, locale.FR_CI, locale.FR_MF:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SA, locale.AR_DJ, locale.AR_EG, locale.AR_TD, locale.AR_YE, locale.AR_JO, locale.AR_EH, locale.AR_KW, locale.AR_BH, locale.AR_SD, locale.AR_IQ, locale.AR_PS, locale.AR_OM, locale.AR_KM, locale.AR_SY, locale.AR_SO, locale.AR_QA, locale.AR_ER, locale.AR_AE, locale.AR, locale.AR_SS, locale.AR_IL:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.KOK, locale.SW_UG, locale.SW_KE, locale.TA_MY, locale.TA_SG, locale.HA_NE, locale.QU_EC, locale.ES_PE, locale.SW, locale.HA_GH, locale.QU, locale.MN, locale.MZN, locale.SD_ARAB, locale.MI, locale.SD, locale.EN_MV, locale.KOK_DEVA, locale.HA:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.KGP, locale.WO, locale.IA, locale.MS_BN, locale.EN_AT, locale.ES_UY, locale.PT, locale.ES_CO, locale.ES_AR, locale.YRL_CO, locale.QU_BO, locale.SW_CD, locale.YRL, locale.JV, locale.YRL_VE:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_SN, locale.FF_ADLM_BF, locale.FF_ADLM_CM, locale.FF_ADLM_GH, locale.FF_ADLM, locale.FF_ADLM_NE, locale.FF_ADLM_NG, locale.FF_ADLM_GM, locale.FF_ADLM_LR, locale.FF_ADLM_SL, locale.FF_ADLM_MR:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.TE, locale.TA, locale.XNR, locale.GU, locale.DZ, locale.HI_LATN, locale.PA_GURU, locale.BN_IN, locale.TA_LK, locale.EN_IN, locale.PA, locale.HI:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_CW, locale.NL, locale.NL_SX, locale.ES_PY, locale.NL_BQ, locale.EN_NL, locale.NL_BE, locale.NL_AW, locale.NL_SR:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_ID, locale.ES_GQ, locale.TR, locale.ES_BO, locale.MS_ID, locale.ID, locale.TR_CY:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV_FI, locale.SV, locale.SV_AX, locale.ET, locale.FI, locale.KSH, locale.LT:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_LY, locale.AR_MA, locale.AR_LB, locale.AR_DZ, locale.AR_TN, locale.AR_MR:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.FO, locale.HR_BA, locale.EU, locale.SL, locale.HR:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.ES_CR, locale.AF, locale.EN_ZA:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_CL, locale.LO, locale.ES_VE, locale.ES_EC:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_LI, locale.GSW_FR, locale.RM:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NB, locale.NO:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_150, locale.MY, locale.CE:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_CH, locale.IT_CH, locale.DE_CH:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.UZ, locale.UZ_LATN:
		return FormatInfo{
			Symbol:           "soʻm",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "OZBS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "इउ.जेत.एस",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߎߗ߭ߛ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.UZ_CYRL:
		return FormatInfo{
			Symbol:           "сўм",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "UZS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
