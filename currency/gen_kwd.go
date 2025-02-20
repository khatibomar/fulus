// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = KWD{}

// Kuwaiti Dinar currency type
type KWD struct{}

func (KWD) Code() string { return "KWD" }

func (KWD) Number() string { return "414" }

func (KWD) Name() string { return "Kuwaiti Dinar" }

func (KWD) MinorUnits() int { return 3 }

func (KWD) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 155 locales
	case locale.TI_ER, locale.EN_FJ, locale.JA, locale.EN_AI, locale.ML, locale.EN_VC, locale.YUE_HANT_CN, locale.EE_TG, locale.EN_CC, locale.EN_MG, locale.EN_NR, locale.ES_419, locale.EN_TT, locale.EN_SG, locale.ES_HN, locale.ZH_HANT_HK, locale.EN_GG, locale.EN_NA, locale.EN_SB, locale.KO, locale.CHR, locale.ES_US, locale.EN_GU, locale.EN_MO, locale.EN_NG, locale.EN_SC, locale.ZU, locale.ZH_HANT_MY, locale.EN_PW, locale.EN_TZ, locale.CY, locale.EN_DG, locale.EN_MT, locale.EN_PR, locale.EN_CA, locale.EN, locale.EN_KY, locale.EN_RW, locale.ES_SV, locale.OR, locale.TH, locale.EN_MU, locale.EN_SH, locale.EN_GD, locale.EN_TC, locale.EN_FK, locale.ZH_HANS_MO, locale.EN_BZ, locale.EN_VG, locale.ES_NI, locale.EN_SL, locale.EN_SD, locale.EN_UG, locale.EN_BS, locale.EN_IE, locale.EN_JE, locale.EN_MY, locale.EN_NU, locale.ES_MX, locale.ES_GT, locale.IG, locale.ZH_HANT, locale.EN_DM, locale.EN_VU, locale.KN, locale.EN_LR, locale.EN_LS, locale.ES_DO, locale.YO, locale.EN_GY, locale.EN_IO, locale.EN_PN, locale.EN_001, locale.EN_UM, locale.ES_PR, locale.PCM, locale.YO_BJ, locale.EN_TO, locale.AK, locale.EN_ER, locale.YUE, locale.EN_ZW, locale.EN_AU, locale.EN_CK, locale.ES_BZ, locale.EN_GM, locale.TI, locale.EN_JM, locale.EN_LC, locale.EN_GI, locale.EN_PK, locale.GD, locale.ES_CU, locale.ZH, locale.MS_SG, locale.SO_ET, locale.EN_CY, locale.EN_NZ, locale.SO_KE, locale.EN_IM, locale.EN_PG, locale.EN_AE, locale.GA, locale.AM, locale.SI, locale.ZH_HANS_HK, locale.CEB, locale.EN_MP, locale.EN_MW, locale.ES_PA, locale.EN_ZM, locale.FIL, locale.GA_GB, locale.ZH_HANS, locale.EN_CX, locale.EN_FM, locale.EN_VI, locale.SO_DJ, locale.EN_SZ, locale.EN_CM, locale.EN_WS, locale.EN_TV, locale.EN_MS, locale.EN_BW, locale.EN_KE, locale.EN_BB, locale.EN_PH, locale.MR, locale.EN_GH, locale.EN_AG, locale.EN_BM, locale.EN_IL, locale.EN_SS, locale.EN_AS, locale.EN_KI, locale.EN_KN, locale.KO_CN, locale.UG, locale.EN_TK, locale.EN_BI, locale.YUE_HANT, locale.EN_HK, locale.MS, locale.EN_GB, locale.KO_KP, locale.SO, locale.YUE_HANS, locale.ZH_HANS_MY, locale.ZH_HANS_SG, locale.EN_NF, locale.EN_MH, locale.EE, locale.ES_BR, locale.ZH_HANT_MO, locale.EN_SX:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.EL_POLYTON, locale.DA_GL, locale.EL_CY, locale.GL, locale.BS_CYRL, locale.SR_CYRL_BA, locale.EN_SI, locale.ES_PH, locale.IS, locale.CA_IT, locale.EN_DK, locale.BS, locale.SR_CYRL, locale.HSB, locale.ES, locale.IT_VA, locale.DE_BE, locale.AZ_LATN, locale.EL, locale.SR_LATN_BA, locale.SR_LATN, locale.LB, locale.RO_MD, locale.KU, locale.CA, locale.SR, locale.DE_LU, locale.AST, locale.FR_LU, locale.DE, locale.EN_DE, locale.LLD, locale.ES_EA, locale.DSB, locale.SR_CYRL_XK, locale.ES_IC, locale.EN_BE, locale.MK, locale.CA_FR, locale.DE_IT, locale.RO, locale.IT, locale.VI, locale.CA_AD, locale.BS_LATN, locale.SR_LATN_XK, locale.DA, locale.IT_SM, locale.SC, locale.SR_LATN_ME, locale.AZ, locale.FR_MA, locale.CA_ES_VALENCIA, locale.SR_CYRL_ME:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 45 locales
	case locale.KK_CYRL, locale.PT_AO, locale.HY, locale.KA, locale.KY, locale.RU_UA, locale.EN_SE, locale.PT_LU, locale.PT_GQ, locale.SK, locale.KK, locale.PT_CH, locale.EN_FI, locale.PT_CV, locale.TG, locale.KK_KZ, locale.PT_MO, locale.UZ, locale.SQ_MK, locale.TK, locale.BE_TARASK, locale.PT_PT, locale.PT_MZ, locale.CV, locale.SQ, locale.UK, locale.SQ_XK, locale.BR, locale.RU_BY, locale.BG, locale.CS, locale.PT_ST, locale.UZ_LATN, locale.BE, locale.LV, locale.TT, locale.FR_CA, locale.PT_TL, locale.PL, locale.HU, locale.RU_KG, locale.RU, locale.RU_MD, locale.RU_KZ, locale.PT_GW:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CH, locale.FR_CI, locale.FR_NC, locale.FR_VU, locale.FR_PF, locale.FR_SY, locale.FR_CG, locale.FR_HT, locale.FR_TN, locale.FR_CM, locale.FR_PM, locale.FR_ML, locale.FR_GF, locale.FR, locale.FR_MU, locale.FR_KM, locale.FR_BJ, locale.FR_NE, locale.FR_DJ, locale.FR_BE, locale.FR_SN, locale.FR_CD, locale.FR_MG, locale.FR_GQ, locale.FR_SC, locale.FR_BI, locale.FR_RW, locale.FR_GN, locale.FR_MQ, locale.FR_TG, locale.FR_GA, locale.FR_RE, locale.FR_YT, locale.FR_BL, locale.FR_MC, locale.FR_WF, locale.FR_MF, locale.FR_CF, locale.FR_BF, locale.FR_GP, locale.FR_MR, locale.FR_TD, locale.FR_DZ:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_YE, locale.AR_IQ, locale.AR_AE, locale.AR_KM, locale.AR_SA, locale.AR_QA, locale.AR_OM, locale.AR_SO, locale.AR_BH, locale.AR_SS, locale.AR_SD, locale.AR_TD, locale.AR_KW, locale.AR_SY, locale.AR_JO, locale.AR_IL, locale.AR_PS, locale.AR_EG, locale.AR_DJ, locale.AR_ER, locale.AR, locale.AR_EH:
		return CurrencyFormatInfo{
			Symbol:           "د.ك.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 20 locales
	case locale.SW_UG, locale.HA, locale.MN, locale.KOK_DEVA, locale.SW_KE, locale.HA_GH, locale.ES_PE, locale.HA_NE, locale.KOK, locale.TA_SG, locale.EN_MV, locale.MI, locale.TA_MY, locale.SW, locale.SD_ARAB, locale.BAL_LATN, locale.QU_EC, locale.SD, locale.MZN, locale.QU:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.YRL_CO, locale.QU_BO, locale.ES_CO, locale.ES_AR, locale.YRL_VE, locale.YRL, locale.WO, locale.JV, locale.ES_UY, locale.PT, locale.SW_CD, locale.IA, locale.MS_BN, locale.KGP, locale.EN_AT:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM, locale.FF_ADLM_GW, locale.FF_ADLM_GH, locale.FF_ADLM_SL, locale.FF_ADLM_LR, locale.FF_ADLM_SN, locale.FF_ADLM_CM, locale.FF_ADLM_NE, locale.FF_ADLM_NG, locale.FF_ADLM_GM, locale.FF_ADLM_BF, locale.FF_ADLM_MR:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.TE, locale.HI_LATN, locale.BN_IN, locale.PA, locale.DZ, locale.EN_IN, locale.HI, locale.XNR, locale.GU, locale.TA_LK, locale.PA_GURU, locale.TA:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_BE, locale.NL_SX, locale.EN_NL, locale.ES_PY, locale.NL_AW, locale.NL, locale.NL_BQ, locale.NL_SR, locale.NL_CW:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_ID, locale.ID, locale.TR_CY, locale.TR, locale.ES_BO, locale.MS_ID, locale.ES_GQ:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV_FI, locale.SV, locale.FI, locale.KSH, locale.LT, locale.SV_AX, locale.ET:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_MR, locale.AR_DZ, locale.AR_LY, locale.AR_TN, locale.AR_LB, locale.AR_MA:
		return CurrencyFormatInfo{
			Symbol:           "د.ك.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.SL, locale.FO, locale.HR_BA, locale.FO_DK, locale.EU, locale.HR:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.ES_CR, locale.AF, locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_VE, locale.ES_CL, locale.LO, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_FR, locale.RM, locale.GSW_LI:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB, locale.NB_SJ, locale.NN, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CE, locale.EN_150, locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "के.दब्ल्यु.दि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߞߥߘ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "KWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
