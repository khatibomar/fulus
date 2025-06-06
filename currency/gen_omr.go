// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = OMR{}

// Rial Omani currency type
type OMR struct{}

func (OMR) Code() string { return "OMR" }

func (OMR) Number() string { return "512" }

func (OMR) Name() string { return "Rial Omani" }

func (OMR) MinorUnits() int { return 3 }

func (OMR) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 157 locales
	case locale.EN_GY, locale.EN_NA, locale.EN_RW, locale.ML, locale.TI, locale.YUE_HANT, locale.EN_GG, locale.TH, locale.ZH_HANS_MO, locale.EN_BW, locale.EE, locale.EN_CX, locale.GA, locale.EN_AE, locale.FIL, locale.EN_NR, locale.ES_CU, locale.MR, locale.EN_HK, locale.EN_SD, locale.EN_GB, locale.EN_JE, locale.ES_MX, locale.ZU, locale.EN_FK, locale.EN_LR, locale.EN_PN, locale.EN_TC, locale.ES_BR, locale.EN_VI, locale.ZH_HANT, locale.EN_UG, locale.EN_CM, locale.ZH_HANT_MY, locale.EN_SH, locale.EN_IE, locale.EN_MU, locale.UG, locale.ES_SV, locale.EN_IL, locale.EN_JM, locale.EN_PK, locale.EN_ZW, locale.ES_PA, locale.EN_PR, locale.CEB, locale.SO_KE, locale.YUE_HANT_MO, locale.YUE_HANT_CN, locale.EN_FJ, locale.GD, locale.IG, locale.EN_AI, locale.EN_BM, locale.ES_DO, locale.OR, locale.AK, locale.ZH_HANS_HK, locale.EN_SX, locale.ZH_HANT_MO, locale.EN_MH, locale.EN_NZ, locale.KO_CN, locale.EN_PG, locale.EN_BB, locale.EN_GI, locale.EN_ZM, locale.EN_IO, locale.EN_MP, locale.EN_MT, locale.EN_SL, locale.JA, locale.ZH_HANS_SG, locale.CHR, locale.EN_CK, locale.EN_SB, locale.EN_KY, locale.ES_BZ, locale.EN_CY, locale.EN_DG, locale.EN_KI, locale.EN_GS, locale.YO_BJ, locale.EN_MW, locale.EN_CC, locale.EN_BS, locale.EN_MS, locale.EN_NG, locale.EN_DM, locale.EN_PH, locale.EN_GM, locale.ES_HN, locale.EN_FM, locale.EN_TO, locale.KO_KP, locale.TI_ER, locale.EN_MY, locale.EN_KE, locale.EN_AU, locale.EN_MG, locale.EN_SC, locale.YUE_HANS, locale.EN_NU, locale.EN_PW, locale.MS_SG, locale.ES_GT, locale.EN_TK, locale.ZH_HANS, locale.EN_BZ, locale.EN_VU, locale.EN_GU, locale.EN_IM, locale.EN_TZ, locale.YUE, locale.EN_KN, locale.YO, locale.EN_SS, locale.PCM, locale.EN_001, locale.EN_NF, locale.MS, locale.SO_DJ, locale.EN_LS, locale.EN_MO, locale.ES_419, locale.ZH_HANT_HK, locale.ZH_HANS_MY, locale.EN_UM, locale.EN_VC, locale.KN, locale.SI, locale.EN_AS, locale.EN_CA, locale.EN_GD, locale.GA_GB, locale.SO_ET, locale.EN_TT, locale.EN_ER, locale.EN_LC, locale.EN_GH, locale.ES_PR, locale.EN, locale.EN_TV, locale.ES_US, locale.CY, locale.KO, locale.EN_AG, locale.SO, locale.EN_BI, locale.AM, locale.EN_VG, locale.EN_WS, locale.EN_SG, locale.ES_NI, locale.ZH, locale.EE_TG, locale.EN_SZ:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 58 locales
	case locale.KU, locale.IT, locale.BS, locale.SR_LATN_ME, locale.FR_LU, locale.EL_POLYTON, locale.MK, locale.ES_PH, locale.EL_CY, locale.BS_CYRL, locale.DE_BE, locale.RO, locale.IT_SM, locale.EN_ES, locale.EN_BE, locale.EN_DE, locale.HSB, locale.VI, locale.SR_CYRL_XK, locale.SR_LATN, locale.ES_EA, locale.SR_LATN_XK, locale.CA, locale.DA, locale.EN_SI, locale.DE_LU, locale.EN_DK, locale.IS, locale.EN_PL, locale.GL, locale.CA_AD, locale.DA_GL, locale.AZ, locale.EN_RO, locale.RO_MD, locale.DE_IT, locale.BS_LATN, locale.CA_ES_VALENCIA, locale.AST, locale.DE, locale.SR_CYRL, locale.AZ_LATN, locale.DSB, locale.SR_CYRL_ME, locale.LB, locale.SR, locale.FR_MA, locale.ES_IC, locale.IT_VA, locale.SR_LATN_BA, locale.CA_IT, locale.EL, locale.EN_IT, locale.LLD, locale.ES, locale.CA_FR, locale.SR_CYRL_BA, locale.SC:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 51 locales
	case locale.EN_NO, locale.KK_CYRL, locale.PT_CH, locale.PT_TL, locale.PT_PT, locale.TG, locale.BR, locale.RU_BY, locale.KY, locale.PT_LU, locale.KA, locale.PT_MZ, locale.EN_PT, locale.RU_UA, locale.CS, locale.SQ, locale.HY, locale.EN_HU, locale.KK, locale.RU_MD, locale.KK_KZ, locale.BG, locale.EN_FI, locale.EN_SE, locale.CV, locale.PT_GQ, locale.HT, locale.RU_KG, locale.PT_GW, locale.PT_ST, locale.SK, locale.HU, locale.SQ_MK, locale.BE_TARASK, locale.TK, locale.PT_CV, locale.FR_CA, locale.LV, locale.EN_CZ, locale.UZ_LATN, locale.EN_SK, locale.PL, locale.TT, locale.PT_AO, locale.UK, locale.RU_KZ, locale.RU, locale.PT_MO, locale.SQ_XK, locale.BE, locale.UZ:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_CG, locale.FR_BI, locale.FR_VU, locale.FR_RW, locale.FR_BE, locale.FR_TD, locale.FR_PF, locale.FR_MQ, locale.FR_CD, locale.EN_FR, locale.FR_PM, locale.FR_MR, locale.FR_CI, locale.FR_DZ, locale.FR_RE, locale.FR_CH, locale.FR_SY, locale.FR_NC, locale.FR_BL, locale.FR_SN, locale.FR_YT, locale.FR_NE, locale.FR_TN, locale.FR_MU, locale.FR_GP, locale.FR_MG, locale.FR_GN, locale.FR_TG, locale.FR_DJ, locale.FR_WF, locale.FR_CF, locale.FR_BF, locale.FR_HT, locale.FR_SC, locale.FR_ML, locale.FR_MC, locale.FR_GA, locale.FR_MF, locale.FR_GF, locale.FR_KM, locale.FR_CM, locale.FR, locale.FR_GQ, locale.FR_BJ:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_EG, locale.AR_KW, locale.AR_PS, locale.AR_KM, locale.AR_YE, locale.AR_TD, locale.AR_EH, locale.AR_ER, locale.AR_OM, locale.AR_AE, locale.AR_IQ, locale.AR_SS, locale.AR_SA, locale.AR_QA, locale.AR_DJ, locale.AR_BH, locale.AR_SD, locale.AR_JO, locale.AR_IL, locale.AR, locale.AR_SO, locale.AR_SY:
		return FormatInfo{
			Symbol:           "ر.ع.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 20 locales
	case locale.HA, locale.SW_UG, locale.SD_ARAB, locale.EN_MV, locale.ES_PE, locale.QU, locale.TA_MY, locale.MZN, locale.SW, locale.QU_EC, locale.HA_NE, locale.SW_KE, locale.HA_GH, locale.MN, locale.BAL_LATN, locale.TA_SG, locale.MI, locale.SD, locale.KOK, locale.KOK_DEVA:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.KGP, locale.SW_CD, locale.JV, locale.ES_CO, locale.ES_UY, locale.YRL, locale.YRL_CO, locale.MS_BN, locale.QU_BO, locale.ES_AR, locale.EN_AT, locale.IA, locale.YRL_VE, locale.WO, locale.PT:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GM, locale.FF_ADLM_SL, locale.FF_ADLM_SN, locale.FF_ADLM_GW, locale.FF_ADLM_MR, locale.FF_ADLM_LR, locale.FF_ADLM_NG, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_NE, locale.FF_ADLM_BF, locale.FF_ADLM_GH:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.GU, locale.DZ, locale.PA, locale.HI_LATN, locale.HI, locale.TA_LK, locale.PA_GURU, locale.XNR, locale.BN_IN, locale.TA, locale.EN_IN, locale.TE:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.NL_SX, locale.NL_CW, locale.NL_SR, locale.ES_PY, locale.NL_BQ, locale.NL_AW, locale.EN_NL, locale.NL_BE:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.LT, locale.SV_AX, locale.FI, locale.ET, locale.KSH, locale.SV_FI, locale.SV:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.TR_CY, locale.TR, locale.ID, locale.MS_ID, locale.EN_ID, locale.ES_GQ, locale.ES_BO:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MA, locale.AR_MR, locale.AR_LB, locale.AR_LY, locale.AR_DZ, locale.AR_TN:
		return FormatInfo{
			Symbol:           "ر.ع.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.EU, locale.HR_BA, locale.HR, locale.SL, locale.FO_DK, locale.FO:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.ES_CR, locale.EN_ZA, locale.AF, locale.AF_NA:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_EC, locale.LO, locale.ES_CL, locale.ES_VE:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.RM, locale.GSW_FR, locale.GSW_LI:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB, locale.NB_SJ, locale.NN, locale.NO:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.BN, locale.CCP:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CE, locale.EN_150, locale.MY:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.EN_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "अ.एम.आर",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߏߡߙ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "OMR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
