// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BGN{}

// Bulgarian Lev currency type
type BGN struct{}

func (BGN) Code() string { return "BGN" }

func (BGN) Number() string { return "975" }

func (BGN) Name() string { return "Bulgarian Lev" }

func (BGN) MinorUnits() int { return 2 }

func (BGN) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 157 locales
	case locale.EN_SD, locale.EN_NF, locale.SO, locale.ES_DO, locale.ZH, locale.EN_KY, locale.KO, locale.EN_GY, locale.EN_GH, locale.SO_ET, locale.EN_HK, locale.KO_KP, locale.SO_KE, locale.EN_LR, locale.ES_HN, locale.TI_ER, locale.EN, locale.MS_SG, locale.ZH_HANS, locale.EN_DG, locale.ZH_HANT_HK, locale.ZU, locale.TI, locale.EN_AG, locale.EN_MU, locale.EN_NG, locale.ES_PR, locale.EN_RW, locale.EN_SB, locale.ES_MX, locale.EN_BM, locale.EN_GB, locale.EN_NZ, locale.EN_DM, locale.GD, locale.EN_IM, locale.SI, locale.GA_GB, locale.EN_GD, locale.YO, locale.EN_TT, locale.FIL, locale.EN_TZ, locale.ES_NI, locale.EN_FJ, locale.EN_BZ, locale.EN_CY, locale.EN_MO, locale.CHR, locale.EN_UG, locale.ZH_HANT, locale.EN_PH, locale.YUE_HANT_MO, locale.EN_JM, locale.EN_SL, locale.MR, locale.EN_KE, locale.KN, locale.EN_SC, locale.EN_BI, locale.EN_GG, locale.EN_MG, locale.ES_PA, locale.ML, locale.EN_001, locale.EN_AE, locale.EN_AI, locale.EN_CX, locale.EN_GU, locale.EN_MP, locale.YUE_HANT_CN, locale.ZH_HANS_MO, locale.CEB, locale.EN_TK, locale.JA, locale.EN_PR, locale.EN_UM, locale.EN_GS, locale.EN_PN, locale.ZH_HANT_MO, locale.EN_JE, locale.EN_WS, locale.ZH_HANS_MY, locale.IG, locale.EN_MH, locale.EN_SH, locale.EN_BB, locale.EN_AS, locale.EN_KI, locale.EN_LC, locale.UG, locale.ZH_HANS_HK, locale.EE_TG, locale.EN_SZ, locale.EN_CC, locale.EN_ER, locale.YUE_HANS, locale.EN_NA, locale.EN_PK, locale.ZH_HANS_SG, locale.EN_LS, locale.EN_MY, locale.EN_VC, locale.ES_US, locale.EN_CK, locale.EN_ZM, locale.EN_TV, locale.ES_SV, locale.EN_PW, locale.OR, locale.EN_IL, locale.EN_SS, locale.ES_BR, locale.YUE_HANT, locale.EN_PG, locale.TH, locale.EN_CA, locale.EN_GM, locale.ZH_HANT_MY, locale.AK, locale.EN_IO, locale.EN_TO, locale.EN_AU, locale.EN_BW, locale.CY, locale.ES_CU, locale.EN_MS, locale.EN_VU, locale.AM, locale.MS, locale.EN_FK, locale.EN_ZW, locale.EN_VG, locale.EN_TC, locale.ES_BZ, locale.PCM, locale.EN_MT, locale.EN_VI, locale.SO_DJ, locale.YO_BJ, locale.EN_FM, locale.EN_NU, locale.EN_BS, locale.EN_SG, locale.EE, locale.EN_GI, locale.ES_GT, locale.EN_KN, locale.GA, locale.EN_MW, locale.EN_SX, locale.YUE, locale.EN_CM, locale.EN_IE, locale.ES_419, locale.KO_CN, locale.EN_NR:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 58 locales
	case locale.IT, locale.KU, locale.EN_IT, locale.SR_LATN_XK, locale.DE_BE, locale.IS, locale.EN_DE, locale.BS_CYRL, locale.LB, locale.CA_IT, locale.SR_CYRL_XK, locale.VI, locale.SR_CYRL, locale.DSB, locale.EN_SI, locale.SR_LATN_BA, locale.SR_LATN_ME, locale.EL_POLYTON, locale.HSB, locale.EN_RO, locale.FR_MA, locale.RO_MD, locale.IT_SM, locale.DE, locale.GL, locale.AST, locale.EN_DK, locale.AZ_LATN, locale.SR_CYRL_BA, locale.FR_LU, locale.SR_LATN, locale.IT_VA, locale.CA_AD, locale.EN_BE, locale.EL, locale.ES_PH, locale.CA_ES_VALENCIA, locale.AZ, locale.RO, locale.EL_CY, locale.DE_LU, locale.DA_GL, locale.LLD, locale.SR_CYRL_ME, locale.CA_FR, locale.BS_LATN, locale.DE_IT, locale.EN_PL, locale.ES_IC, locale.SC, locale.EN_ES, locale.MK, locale.CA, locale.SR, locale.DA, locale.BS, locale.ES, locale.ES_EA:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 50 locales
	case locale.LV, locale.BR, locale.EN_FI, locale.UZ, locale.FR_CA, locale.RU_MD, locale.RU_BY, locale.EN_CZ, locale.PT_MZ, locale.CS, locale.PT_LU, locale.EN_HU, locale.RU_KG, locale.RU_KZ, locale.SQ, locale.KK_KZ, locale.TT, locale.PT_ST, locale.PT_CH, locale.TG, locale.EN_SE, locale.HY, locale.PT_GW, locale.HU, locale.SQ_XK, locale.KY, locale.UK, locale.CV, locale.BE_TARASK, locale.RU_UA, locale.PT_AO, locale.EN_PT, locale.KK_CYRL, locale.PT_GQ, locale.PT_TL, locale.KK, locale.PT_MO, locale.PT_PT, locale.RU, locale.SK, locale.EN_SK, locale.BE, locale.KA, locale.UZ_LATN, locale.EN_NO, locale.HT, locale.PL, locale.TK, locale.SQ_MK, locale.PT_CV:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_BL, locale.FR_MQ, locale.FR_MG, locale.FR_CG, locale.FR_CI, locale.FR_BE, locale.FR_CM, locale.FR_SY, locale.FR_GF, locale.FR_TG, locale.FR_RW, locale.FR_PF, locale.FR_WF, locale.FR_CD, locale.FR, locale.FR_DJ, locale.FR_MF, locale.FR_CF, locale.FR_RE, locale.FR_SC, locale.FR_MU, locale.FR_GP, locale.FR_SN, locale.EN_FR, locale.FR_BF, locale.FR_BI, locale.FR_MC, locale.FR_HT, locale.FR_TD, locale.FR_MR, locale.FR_BJ, locale.FR_TN, locale.FR_NC, locale.FR_NE, locale.FR_VU, locale.FR_PM, locale.FR_DZ, locale.FR_GQ, locale.FR_KM, locale.FR_ML, locale.FR_GA, locale.FR_CH, locale.FR_GN, locale.FR_YT:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_TD, locale.AR_IL, locale.AR_KW, locale.AR_DJ, locale.AR_AE, locale.AR_EH, locale.AR_SY, locale.AR, locale.AR_SS, locale.AR_SD, locale.AR_JO, locale.AR_ER, locale.AR_PS, locale.AR_KM, locale.AR_EG, locale.AR_IQ, locale.AR_BH, locale.AR_SA, locale.AR_SO, locale.AR_YE, locale.AR_OM, locale.AR_QA:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.ES_PE, locale.MZN, locale.MI, locale.SD, locale.SW_KE, locale.KOK, locale.QU, locale.SD_ARAB, locale.HA, locale.HA_NE, locale.SW_UG, locale.HA_GH, locale.EN_MV, locale.KOK_DEVA, locale.QU_EC, locale.TA_SG, locale.SW, locale.MN, locale.TA_MY:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.QU_BO, locale.MS_BN, locale.YRL_VE, locale.ES_UY, locale.EN_AT, locale.SW_CD, locale.JV, locale.WO, locale.YRL_CO, locale.KGP, locale.ES_AR, locale.PT, locale.YRL, locale.IA, locale.ES_CO:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_BF, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_GM, locale.FF_ADLM_LR, locale.FF_ADLM_SN, locale.FF_ADLM_SL, locale.FF_ADLM_GH, locale.FF_ADLM, locale.FF_ADLM_GW, locale.FF_ADLM_NG, locale.FF_ADLM_CM:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.HI_LATN, locale.BN_IN, locale.PA, locale.XNR, locale.HI, locale.TA_LK, locale.PA_GURU, locale.EN_IN, locale.GU, locale.TE, locale.TA:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ES_PY, locale.NL_SX, locale.NL_BQ, locale.NL_CW, locale.NL_AW, locale.EN_NL, locale.NL, locale.NL_BE, locale.NL_SR:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.ES_BO, locale.MS_ID, locale.TR, locale.EN_ID, locale.TR_CY, locale.ES_GQ, locale.ID:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV_AX, locale.LT, locale.FI, locale.KSH, locale.ET, locale.SV, locale.SV_FI:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_LB, locale.AR_MA, locale.AR_TN, locale.AR_DZ, locale.AR_MR, locale.AR_LY:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.SL, locale.EU, locale.FO, locale.HR_BA, locale.HR:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.AF, locale.EN_ZA, locale.ES_CR:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.LO, locale.ES_VE, locale.ES_EC, locale.ES_CL:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB, locale.NN, locale.NO, locale.NB_SJ:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW_FR, locale.GSW_LI, locale.GSW:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_CH, locale.DE_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.MY, locale.EN_150, locale.CE:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE, locale.AS, locale.NE_IN:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "BLGL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BG:
		return FormatInfo{
			Symbol:           "лв.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "बि.जि.एन",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߓߜ߭ߟ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
