// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BYN{}

// Belarusian Ruble currency type
type BYN struct{}

func (BYN) Code() string { return "BYN" }

func (BYN) Number() string { return "933" }

func (BYN) Name() string { return "Belarusian Ruble" }

func (BYN) MinorUnits() int { return 2 }

func (BYN) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 109 locales
	case locale.EN_BI, locale.EN_GB, locale.EN_KE, locale.EN_ZW, locale.TI, locale.EN_IL, locale.EN_FM, locale.EN_GS, locale.EN_WS, locale.EN_SC, locale.EN_KN, locale.IG, locale.EN_VG, locale.SO_DJ, locale.TI_ER, locale.EN_DG, locale.EN_SH, locale.EN_TT, locale.YO, locale.EN_BW, locale.EN_CA, locale.EN_CX, locale.EN_001, locale.EN_UG, locale.AK, locale.EN_AS, locale.EN_MG, locale.EN_SG, locale.CEB, locale.EN_IE, locale.EN_MP, locale.EN_PH, locale.GA, locale.EN_NU, locale.EN_BS, locale.EN_KY, locale.EN_MO, locale.EN_UM, locale.EN_MT, locale.EN_NZ, locale.EN_MS, locale.EN_CM, locale.EN_JM, locale.EN_VC, locale.EN_GG, locale.EN_MU, locale.EN_TV, locale.EN_FJ, locale.EN_GH, locale.EN_AU, locale.EN_PW, locale.EN_SZ, locale.EN_MH, locale.SO, locale.EN_PG, locale.EN_VU, locale.MR, locale.EN_GD, locale.EN_GY, locale.EN_ZM, locale.EN_MY, locale.EN_NA, locale.EN_NG, locale.EN_SL, locale.EN_LC, locale.YO_BJ, locale.EN_AI, locale.EN_JE, locale.EN_ER, locale.EN_TC, locale.EN_TZ, locale.GA_GB, locale.EN_NR, locale.EN_SB, locale.EN_BZ, locale.EN_CY, locale.EN_PK, locale.EN_TO, locale.EN, locale.EN_AG, locale.EN_GI, locale.EN_IM, locale.EN_SX, locale.SO_ET, locale.EN_LR, locale.EN_CK, locale.EN_MW, locale.EN_LS, locale.EN_FK, locale.UG, locale.EN_KI, locale.EN_CC, locale.EN_BM, locale.EN_AE, locale.SO_KE, locale.EN_NF, locale.EN_RW, locale.EN_HK, locale.EN_VI, locale.EN_BB, locale.EN_PN, locale.EN_PR, locale.EN_GM, locale.EN_GU, locale.EN_TK, locale.EN_IO, locale.EN_SD, locale.EN_DM, locale.EN_SS:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 46 locales
	case locale.ES_DO, locale.MS, locale.ES_CU, locale.ES_NI, locale.EE_TG, locale.ES_GT, locale.ZH_HANS_MO, locale.ES_SV, locale.FIL, locale.ES_HN, locale.ZH, locale.KN, locale.ES_BR, locale.ZH_HANT_MO, locale.ZH_HANS_HK, locale.YUE_HANT_MO, locale.EE, locale.MT, locale.GD, locale.KO, locale.ZH_HANT_HK, locale.ZH_HANS_MY, locale.ZH_HANT_MY, locale.ES_BZ, locale.ES_PA, locale.YUE_HANT, locale.SI, locale.YUE, locale.TH, locale.CY, locale.KO_CN, locale.CHR, locale.JA, locale.ZH_HANS, locale.ES_PR, locale.ES_US, locale.ML, locale.YUE_HANT_CN, locale.AM, locale.OR, locale.KO_KP, locale.ZH_HANS_SG, locale.ES_419, locale.ZH_HANT, locale.MS_SG, locale.YUE_HANS:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_BL, locale.FR_HT, locale.FR_GA, locale.FR_PF, locale.FR_BI, locale.FR_CG, locale.FR_CM, locale.FR_CF, locale.FR_BJ, locale.FR_PM, locale.FR_GF, locale.FR_GQ, locale.FR_VU, locale.FR_TD, locale.FR_BE, locale.FR_YT, locale.FR_MU, locale.FR_SC, locale.FR_CI, locale.FR_MG, locale.FR_MQ, locale.FR_CH, locale.FR_DJ, locale.FR_MR, locale.FR_BF, locale.FR_MF, locale.FR_NC, locale.FR_WF, locale.FR, locale.FR_SN, locale.FR_DZ, locale.FR_RW, locale.FR_CD, locale.FR_NE, locale.FR_RE, locale.FR_GP, locale.FR_GN, locale.FR_TG, locale.FR_SY, locale.FR_MC, locale.FR_KM, locale.FR_TN, locale.FR_ML:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 32 locales
	case locale.TK, locale.KK, locale.UZ, locale.RU_MD, locale.HY, locale.LV, locale.RU_UA, locale.HU, locale.BR, locale.PT_GW, locale.KK_CYRL, locale.HT, locale.SK, locale.KA, locale.PT_GQ, locale.PT_CH, locale.RU_KG, locale.PT_AO, locale.PT_ST, locale.KY, locale.KK_KZ, locale.PT_TL, locale.RU, locale.PT_LU, locale.UZ_LATN, locale.PT_PT, locale.PT_MO, locale.UK, locale.PT_CV, locale.CS, locale.PT_MZ, locale.RU_KZ:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 31 locales
	case locale.ES_PH, locale.ES_EA, locale.FR_MA, locale.EL_CY, locale.CA_AD, locale.SR, locale.EL_POLYTON, locale.SR_CYRL_XK, locale.EL, locale.DE_BE, locale.DE_IT, locale.AZ, locale.CA_ES_VALENCIA, locale.DE_LU, locale.SR_CYRL_BA, locale.VI, locale.RO, locale.BS, locale.ES, locale.BS_LATN, locale.SR_CYRL_ME, locale.ES_IC, locale.CA_IT, locale.DE, locale.FR_LU, locale.CA, locale.RO_MD, locale.SR_CYRL, locale.AZ_LATN, locale.MK, locale.CA_FR:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_JO, locale.AR_YE, locale.AR_QA, locale.AR_EH, locale.AR_SO, locale.AR_AE, locale.AR_DJ, locale.AR_TD, locale.AR_SY, locale.AR_KM, locale.AR_SD, locale.AR_BH, locale.AR_EG, locale.AR_IQ, locale.AR_IL, locale.AR_OM, locale.AR_SA, locale.AR, locale.AR_ER, locale.AR_KW, locale.AR_PS, locale.AR_SS:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 17 locales
	case locale.EN_PL, locale.EN_SI, locale.EN_ES, locale.SC, locale.AST, locale.HSB, locale.IS, locale.EN_RO, locale.DSB, locale.KU, locale.EN_DE, locale.EN_BE, locale.EN_IT, locale.LB, locale.EN_DK, locale.BS_CYRL, locale.LLD:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.EN_SE, locale.EN_CZ, locale.EN_NO, locale.PL, locale.EN_PT, locale.TT, locale.EN_SK, locale.SQ, locale.SQ_XK, locale.TG, locale.EN_HU, locale.SQ_MK, locale.BG, locale.CV, locale.EN_FI:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_MR, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_NG, locale.FF_ADLM_NE, locale.FF_ADLM_LR, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_GM, locale.FF_ADLM_GW, locale.FF_ADLM_SL, locale.FF_ADLM_BF:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.SD, locale.TA_MY, locale.MN, locale.KOK_DEVA, locale.SW_UG, locale.ES_PE, locale.KOK, locale.TA_SG, locale.SW, locale.MZN, locale.SW_KE, locale.SD_ARAB:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.GU, locale.TE, locale.TA_LK, locale.HI, locale.BN_IN, locale.PA, locale.TA, locale.PA_GURU:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_AW, locale.NL, locale.NL_BQ, locale.NL_BE, locale.ES_PY, locale.NL_CW, locale.NL_SX, locale.NL_SR:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_MV, locale.HA_NE, locale.QU, locale.MI, locale.HA, locale.QU_EC, locale.HA_GH:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_TN, locale.AR_MA, locale.AR_LY, locale.AR_MR, locale.AR_LB:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ES_GQ, locale.ES_BO, locale.TR, locale.MS_ID, locale.TR_CY, locale.ID:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.ES_UY, locale.SW_CD, locale.ES_CO, locale.PT, locale.MS_BN, locale.ES_AR:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.JV, locale.IA, locale.FUR, locale.WO, locale.EN_AT, locale.QU_BO:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.SL, locale.FO, locale.EU, locale.HR, locale.HR_BA, locale.FO_DK:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.BE_TARASK, locale.BE, locale.FR_CA, locale.RU_BY:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_CL, locale.LO, locale.ES_VE, locale.ES_EC:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.RM, locale.GSW, locale.GSW_FR:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.IT_SM, locale.IT, locale.GL, locale.IT_VA:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.KGP, locale.YRL_CO, locale.YRL, locale.YRL_VE:
		return FormatInfo{
			Symbol:           "p.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NO, locale.NN, locale.NB:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.SR_LATN, locale.SR_LATN_BA, locale.SR_LATN_XK, locale.SR_LATN_ME:
		return FormatInfo{
			Symbol:           "r.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AF, locale.ES_CR, locale.AF_NA:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.BN, locale.CCP:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.FI, locale.KSH, locale.ET:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.NE, locale.NE_IN, locale.AS:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SV_FI, locale.SV_AX, locale.SV:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.XNR, locale.EN_IN, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.CE, locale.MY:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.DA_GL, locale.DA:
		return FormatInfo{
			Symbol:           "Br.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_MX, locale.PCM:
		return FormatInfo{
			Symbol:           "p.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "BLRR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "बि.वाई.एन",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_CH:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_NL:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "р",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.IT_CH:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.KM:
		return FormatInfo{
			Symbol:           "р.",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.LT:
		return FormatInfo{
			Symbol:           "Br",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߓߌߙ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ZU:
		return FormatInfo{
			Symbol:           "P.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "BYN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
