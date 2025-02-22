// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = MOP{}

// Pataca currency type
type MOP struct{}

func (MOP) Code() string { return "MOP" }

func (MOP) Number() string { return "446" }

func (MOP) Name() string { return "Pataca" }

func (MOP) MinorUnits() int { return 2 }

func (MOP) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 152 locales
	case locale.EN_CK, locale.EN_RW, locale.EN_WS, locale.AK, locale.EN_001, locale.EN_NA, locale.EN_IO, locale.EN_MP, locale.EN_SH, locale.EN_FJ, locale.EN_FK, locale.EN_GU, locale.EN_TT, locale.YUE_HANT, locale.EN_CY, locale.JA, locale.ZH_HANT, locale.EN_AS, locale.EE, locale.EN_PN, locale.PCM, locale.EN_NU, locale.EN_HK, locale.EN_IL, locale.TI_ER, locale.ZH_HANS_MY, locale.EN_AG, locale.EN_AI, locale.EN_BZ, locale.SO, locale.EN_SX, locale.ES_BR, locale.EN_ZM, locale.EN_MW, locale.EN_PG, locale.KO, locale.OR, locale.EN_BM, locale.EN_PK, locale.ZH_HANS_HK, locale.EN_TZ, locale.ZH_HANT_MY, locale.EN_CC, locale.EN_ER, locale.EN_CM, locale.EN_VG, locale.EN_GM, locale.ZH, locale.YO, locale.EN_PH, locale.EN_LC, locale.EN_SZ, locale.ML, locale.EN_AU, locale.EN_NZ, locale.ES_419, locale.EN_ZW, locale.ES_MX, locale.EN_TK, locale.EN_SL, locale.GD, locale.EN_BW, locale.EN_UM, locale.EN_CX, locale.ES_HN, locale.YUE_HANT_CN, locale.EN_GG, locale.ES_PR, locale.EN_BI, locale.EN_IE, locale.ZH_HANS, locale.TH, locale.MS, locale.ZH_HANT_HK, locale.EN_JE, locale.EN_TO, locale.ZH_HANS_SG, locale.AM, locale.EN_PW, locale.EN_SD, locale.ES_CU, locale.EN_GY, locale.ES_PA, locale.SO_ET, locale.EN_KN, locale.UG, locale.SI, locale.SO_KE, locale.FIL, locale.KO_CN, locale.EN_FM, locale.EN_MT, locale.ES_SV, locale.EN_MU, locale.EN_NR, locale.EN_LR, locale.EN_DM, locale.EN_MG, locale.KO_KP, locale.EN_GB, locale.EN_GD, locale.EN_KI, locale.YUE_HANS, locale.EN_NF, locale.EN_KY, locale.ES_GT, locale.EN_BB, locale.EN_JM, locale.KN, locale.EN_SG, locale.EN_UG, locale.EN_SC, locale.EN_PR, locale.ES_BZ, locale.IG, locale.EN_SB, locale.ZU, locale.EN_DG, locale.EN_MY, locale.EN_KE, locale.YUE, locale.EN_SS, locale.ES_NI, locale.EN_BS, locale.EN_TC, locale.ES_US, locale.CY, locale.EN_MH, locale.EN_CA, locale.EN_LS, locale.SO_DJ, locale.YO_BJ, locale.EN_TV, locale.EN_VI, locale.TI, locale.EE_TG, locale.EN, locale.EN_GI, locale.EN_VC, locale.MR, locale.CHR, locale.EN_MS, locale.EN_NG, locale.CEB, locale.ES_DO, locale.GA, locale.GA_GB, locale.EN_AE, locale.EN_IM, locale.MS_SG, locale.EN_GH, locale.EN_VU:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.IT_SM, locale.EN_DE, locale.RO_MD, locale.IT, locale.SR, locale.SR_LATN_XK, locale.HSB, locale.IT_VA, locale.DE_BE, locale.CA_IT, locale.ES_PH, locale.GL, locale.EN_SI, locale.LB, locale.SC, locale.AZ_LATN, locale.AST, locale.EN_BE, locale.FR_LU, locale.EN_DK, locale.LLD, locale.BS_CYRL, locale.KU, locale.CA_ES_VALENCIA, locale.EL_CY, locale.BS, locale.SR_CYRL_BA, locale.CA_FR, locale.SR_LATN_BA, locale.MK, locale.DE, locale.SR_CYRL_XK, locale.DA_GL, locale.DSB, locale.SR_CYRL_ME, locale.ES, locale.DE_IT, locale.CA, locale.IS, locale.BS_LATN, locale.CA_AD, locale.SR_CYRL, locale.EL, locale.FR_MA, locale.AZ, locale.DE_LU, locale.SR_LATN, locale.SR_LATN_ME, locale.VI, locale.RO, locale.EL_POLYTON, locale.DA, locale.ES_EA, locale.ES_IC:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.SQ_XK, locale.HY, locale.PT_LU, locale.UZ_LATN, locale.TK, locale.HU, locale.PL, locale.BG, locale.RU_UA, locale.BE, locale.EN_FI, locale.KA, locale.RU_MD, locale.TG, locale.SK, locale.LV, locale.FR_CA, locale.BR, locale.KK, locale.PT_GQ, locale.SQ, locale.PT_GW, locale.RU, locale.KY, locale.BE_TARASK, locale.PT_PT, locale.KK_KZ, locale.RU_KG, locale.PT_CH, locale.SQ_MK, locale.PT_MZ, locale.EN_SE, locale.PT_TL, locale.KK_CYRL, locale.UZ, locale.CV, locale.PT_AO, locale.RU_BY, locale.PT_ST, locale.CS, locale.UK, locale.TT, locale.RU_KZ, locale.PT_CV:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CF, locale.FR_RE, locale.FR_ML, locale.FR_GN, locale.FR, locale.FR_SY, locale.FR_TN, locale.FR_GQ, locale.FR_CH, locale.FR_CM, locale.FR_MC, locale.FR_CI, locale.FR_DJ, locale.FR_HT, locale.FR_NC, locale.FR_GP, locale.FR_DZ, locale.FR_MF, locale.FR_SN, locale.FR_PF, locale.FR_MG, locale.FR_GA, locale.FR_BE, locale.FR_MQ, locale.FR_PM, locale.FR_YT, locale.FR_BL, locale.FR_BF, locale.FR_KM, locale.FR_CD, locale.FR_TD, locale.FR_RW, locale.FR_WF, locale.FR_BI, locale.FR_VU, locale.FR_SC, locale.FR_NE, locale.FR_TG, locale.FR_MU, locale.FR_BJ, locale.FR_CG, locale.FR_MR, locale.FR_GF:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SA, locale.AR_TD, locale.AR, locale.AR_EH, locale.AR_AE, locale.AR_YE, locale.AR_JO, locale.AR_ER, locale.AR_SD, locale.AR_KM, locale.AR_EG, locale.AR_IL, locale.AR_SO, locale.AR_PS, locale.AR_BH, locale.AR_SY, locale.AR_DJ, locale.AR_KW, locale.AR_SS, locale.AR_IQ, locale.AR_OM, locale.AR_QA:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.MN, locale.HA, locale.KOK, locale.EN_MV, locale.MI, locale.TA_SG, locale.SW, locale.SD, locale.SW_KE, locale.HA_GH, locale.HA_NE, locale.TA_MY, locale.MZN, locale.ES_PE, locale.KOK_DEVA, locale.SD_ARAB, locale.QU, locale.SW_UG, locale.QU_EC:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.WO, locale.MS_BN, locale.JV, locale.YRL_VE, locale.SW_CD, locale.PT, locale.QU_BO, locale.YRL, locale.EN_AT, locale.KGP, locale.IA, locale.YRL_CO, locale.ES_UY, locale.ES_CO, locale.ES_AR:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_GM, locale.FF_ADLM, locale.FF_ADLM_CM, locale.FF_ADLM_MR, locale.FF_ADLM_SL, locale.FF_ADLM_BF, locale.FF_ADLM_NE, locale.FF_ADLM_NG, locale.FF_ADLM_LR, locale.FF_ADLM_GH, locale.FF_ADLM_SN:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.XNR, locale.HI_LATN, locale.PA, locale.TE, locale.HI, locale.PA_GURU, locale.GU, locale.TA, locale.EN_IN, locale.BN_IN, locale.TA_LK:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.NL_SX, locale.ES_PY, locale.EN_NL, locale.NL_BE, locale.NL_AW, locale.NL_CW, locale.NL_BQ, locale.NL_SR:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_ID, locale.ES_GQ, locale.TR_CY, locale.MS_ID, locale.TR, locale.ID, locale.ES_BO:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSH, locale.LT, locale.SV_AX, locale.SV_FI, locale.ET, locale.SV, locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_MA, locale.AR_LB, locale.AR_LY, locale.AR_TN, locale.AR_MR:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.SL, locale.HR, locale.EU, locale.FO, locale.HR_BA:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.AF, locale.EN_ZA, locale.ES_CR:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_EC, locale.ES_CL, locale.LO, locale.ES_VE:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.GSW, locale.RM, locale.GSW_LI:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NN, locale.NB_SJ, locale.NO, locale.NB:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_150, locale.CE, locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ZH_HANT_MO, locale.ZH_HANS_MO, locale.EN_MO:
		return CurrencyFormatInfo{
			Symbol:           "MOP$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "MKNP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "एम.अ.पि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߡߏߔ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PT_MO:
		return CurrencyFormatInfo{
			Symbol:           "MOP$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "MOP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
