// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BGN{}

// Bulgarian Lev currency type
type BGN struct{}

func (BGN) Code() string { return "BGN" }

func (BGN) Number() string { return "975" }

func (BGN) Name() string { return "Bulgarian Lev" }

func (BGN) MinorUnits() int { return 2 }

func (BGN) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 155 locales
	case locale.EN_NG, locale.EN_PN, locale.ZH_HANS_MO, locale.EN_TK, locale.EN_WS, locale.EN_KN, locale.EN_KY, locale.EN_DM, locale.EN_SB, locale.YO, locale.YUE_HANT, locale.GA, locale.GD, locale.EN_SC, locale.TI, locale.SO_KE, locale.EN_FK, locale.EN_MS, locale.EN_SD, locale.EN_JM, locale.EN_MP, locale.EN_GU, locale.EN_UG, locale.EN_AG, locale.EE_TG, locale.EN_MY, locale.IG, locale.EN_PW, locale.EN_VG, locale.JA, locale.KO_CN, locale.EN_SS, locale.MS, locale.EN_PK, locale.EN_GM, locale.ZH_HANT_MY, locale.EN_TO, locale.EN_AI, locale.EN_BS, locale.EN_KE, locale.EN_JE, locale.EN_BW, locale.ES_PA, locale.EN_KI, locale.EN_MO, locale.ES_GT, locale.EN_FJ, locale.EN_MH, locale.ES_HN, locale.TI_ER, locale.PCM, locale.EN_BM, locale.ML, locale.EN_NA, locale.EN, locale.KN, locale.YUE_HANS, locale.CHR, locale.KO_KP, locale.ES_US, locale.MR, locale.EN_MT, locale.EN_NF, locale.ES_NI, locale.EN_001, locale.EN_MW, locale.ZH, locale.ZH_HANS, locale.SO_DJ, locale.ZH_HANT_HK, locale.AK, locale.MS_SG, locale.GA_GB, locale.EN_GD, locale.ES_PR, locale.EN_RW, locale.ES_DO, locale.EN_ER, locale.EN_LC, locale.EN_UM, locale.AM, locale.ZH_HANT_MO, locale.EN_TV, locale.TH, locale.EN_SZ, locale.EN_VI, locale.EN_AS, locale.EN_BZ, locale.EN_MG, locale.ZH_HANS_SG, locale.EN_IE, locale.EN_PH, locale.ES_BZ, locale.YUE, locale.EN_BI, locale.EN_CA, locale.EN_CM, locale.EN_GB, locale.EN_CK, locale.EN_LR, locale.EN_SH, locale.EN_GH, locale.EE, locale.EN_IO, locale.OR, locale.SI, locale.EN_IL, locale.EN_GG, locale.EN_ZW, locale.EN_AU, locale.EN_NU, locale.EN_PR, locale.SO, locale.EN_BB, locale.EN_MU, locale.EN_PG, locale.YUE_HANT_CN, locale.EN_CY, locale.EN_VC, locale.ZH_HANS_MY, locale.EN_LS, locale.ZH_HANS_HK, locale.EN_SL, locale.FIL, locale.EN_IM, locale.EN_SX, locale.EN_TC, locale.EN_DG, locale.SO_ET, locale.ZU, locale.EN_SG, locale.EN_AE, locale.ES_BR, locale.EN_NR, locale.ES_CU, locale.EN_CC, locale.EN_VU, locale.KO, locale.ZH_HANT, locale.CY, locale.ES_MX, locale.UG, locale.EN_TZ, locale.CEB, locale.EN_NZ, locale.EN_HK, locale.EN_FM, locale.EN_GY, locale.ES_SV, locale.ES_419, locale.EN_TT, locale.EN_CX, locale.EN_GI, locale.YO_BJ, locale.EN_ZM:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.SR_CYRL_XK, locale.AZ, locale.ES, locale.EN_SI, locale.BS, locale.HSB, locale.SR_CYRL_ME, locale.ES_EA, locale.SR, locale.SR_LATN, locale.BS_LATN, locale.EN_BE, locale.CA_AD, locale.DE_BE, locale.EL_POLYTON, locale.SR_CYRL_BA, locale.SR_LATN_ME, locale.SC, locale.IS, locale.LLD, locale.AZ_LATN, locale.MK, locale.DA, locale.CA_ES_VALENCIA, locale.DSB, locale.FR_MA, locale.CA_FR, locale.EL_CY, locale.GL, locale.DE_IT, locale.EN_DE, locale.IT_VA, locale.KU, locale.SR_CYRL, locale.SR_LATN_XK, locale.AST, locale.IT_SM, locale.VI, locale.IT, locale.ES_PH, locale.LB, locale.BS_CYRL, locale.ES_IC, locale.EL, locale.CA, locale.DE_LU, locale.FR_LU, locale.RO_MD, locale.SR_LATN_BA, locale.RO, locale.CA_IT, locale.DA_GL, locale.DE, locale.EN_DK:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.UZ_LATN, locale.EN_FI, locale.SK, locale.SQ_XK, locale.PT_ST, locale.CV, locale.HU, locale.PT_GQ, locale.RU_BY, locale.KK_KZ, locale.RU_UA, locale.KK_CYRL, locale.PT_GW, locale.BE_TARASK, locale.KY, locale.TG, locale.FR_CA, locale.PT_CV, locale.PT_AO, locale.HY, locale.CS, locale.TT, locale.UK, locale.PT_MZ, locale.RU, locale.TK, locale.PT_TL, locale.PT_MO, locale.PL, locale.LV, locale.BE, locale.RU_KG, locale.PT_PT, locale.UZ, locale.SQ_MK, locale.RU_MD, locale.BR, locale.KA, locale.SQ, locale.PT_LU, locale.KK, locale.PT_CH, locale.RU_KZ, locale.EN_SE:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_SY, locale.FR_RE, locale.FR_TG, locale.FR_CF, locale.FR_PM, locale.FR_BJ, locale.FR_MC, locale.FR_CD, locale.FR_NC, locale.FR_MG, locale.FR_BL, locale.FR_GP, locale.FR_CG, locale.FR_KM, locale.FR_CM, locale.FR_GF, locale.FR_TN, locale.FR_DZ, locale.FR_CI, locale.FR_MU, locale.FR_MF, locale.FR_SC, locale.FR_DJ, locale.FR_GA, locale.FR_BE, locale.FR_PF, locale.FR_GQ, locale.FR_SN, locale.FR_MQ, locale.FR_GN, locale.FR_TD, locale.FR, locale.FR_CH, locale.FR_NE, locale.FR_ML, locale.FR_VU, locale.FR_BF, locale.FR_BI, locale.FR_HT, locale.FR_YT, locale.FR_MR, locale.FR_RW, locale.FR_WF:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_OM, locale.AR_PS, locale.AR_KW, locale.AR_SD, locale.AR, locale.AR_SS, locale.AR_SA, locale.AR_YE, locale.AR_JO, locale.AR_KM, locale.AR_QA, locale.AR_EG, locale.AR_EH, locale.AR_SY, locale.AR_IL, locale.AR_SO, locale.AR_ER, locale.AR_TD, locale.AR_DJ, locale.AR_BH, locale.AR_AE, locale.AR_IQ:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.MI, locale.QU_EC, locale.SW, locale.SD, locale.MZN, locale.EN_MV, locale.HA, locale.ES_PE, locale.TA_SG, locale.HA_GH, locale.SD_ARAB, locale.QU, locale.SW_KE, locale.SW_UG, locale.TA_MY, locale.MN, locale.KOK, locale.HA_NE, locale.KOK_DEVA:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.ES_CO, locale.ES_AR, locale.PT, locale.MS_BN, locale.EN_AT, locale.KGP, locale.ES_UY, locale.JV, locale.SW_CD, locale.YRL_CO, locale.IA, locale.YRL_VE, locale.WO, locale.QU_BO, locale.YRL:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_CM, locale.FF_ADLM_GM, locale.FF_ADLM_MR, locale.FF_ADLM_GW, locale.FF_ADLM_SL, locale.FF_ADLM_NE, locale.FF_ADLM_SN, locale.FF_ADLM_BF, locale.FF_ADLM, locale.FF_ADLM_GH, locale.FF_ADLM_NG:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.GU, locale.HI, locale.PA_GURU, locale.TA, locale.BN_IN, locale.PA, locale.TA_LK, locale.XNR, locale.EN_IN, locale.HI_LATN, locale.TE:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_AW, locale.NL_BQ, locale.NL_CW, locale.NL_BE, locale.NL_SX, locale.NL, locale.EN_NL, locale.ES_PY, locale.NL_SR:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.ET, locale.KSH, locale.FI, locale.SV_AX, locale.SV_FI, locale.SV, locale.LT:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.TR, locale.ID, locale.ES_GQ, locale.ES_BO, locale.TR_CY, locale.MS_ID, locale.EN_ID:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_LY, locale.AR_MR, locale.AR_TN, locale.AR_DZ, locale.AR_MA, locale.AR_LB:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR_BA, locale.FO, locale.HR, locale.FO_DK, locale.EU, locale.SL:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.EN_ZA, locale.ES_CR, locale.AF, locale.AF_NA:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_EC, locale.LO, locale.ES_VE, locale.ES_CL:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_LI, locale.RM, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NB, locale.NN, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.BN, locale.CCP_IN:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.EN_CH, locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_150, locale.MY, locale.CE:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "BLGL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BG:
		return CurrencyFormatInfo{
			Symbol:           "лв.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "बि.जि.एन",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߓߜ߭ߟ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "BGN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
