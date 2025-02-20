// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = SZL{}

// Lilangeni currency type
type SZL struct{}

func (SZL) Code() string { return "SZL" }

func (SZL) Number() string { return "748" }

func (SZL) Name() string { return "Lilangeni" }

func (SZL) MinorUnits() int { return 2 }

func (SZL) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 178 locales
	case locale.KN, locale.EN_GH, locale.EN_IM, locale.EN_KN, locale.ND, locale.EN_TC, locale.EE_TG, locale.KLN, locale.EN_BZ, locale.EN_IO, locale.KO_CN, locale.EN_VG, locale.EN_WS, locale.EN_AU, locale.EN_SC, locale.EN_PK, locale.EN_AG, locale.EN_MO, locale.IG, locale.ZU, locale.NAQ, locale.EN_VI, locale.ZH_HANT_MO, locale.EN_SL, locale.EN_LR, locale.EBU, locale.EN_GD, locale.PCM, locale.EN_AI, locale.EN_UM, locale.EN_ER, locale.EN_KY, locale.VUN, locale.EN_TZ, locale.EN_GI, locale.MAS, locale.ML, locale.ZH, locale.EN_NF, locale.VAI_VAII, locale.EN_DG, locale.KAM, locale.ROF, locale.ZH_HANS_MY, locale.AK, locale.BM, locale.EN_SD, locale.EN, locale.EN_GG, locale.EN_MP, locale.EN_PG, locale.EN_TK, locale.GUZ, locale.AM, locale.TEO, locale.YUE_HANS, locale.ES_SV, locale.EN_FK, locale.EN_PH, locale.EN_JE, locale.EN_001, locale.SO_ET, locale.EN_PN, locale.GA_GB, locale.SN, locale.VAI, locale.EN_ZW, locale.EN_GB, locale.EN_MY, locale.EN_MU, locale.EN_SB, locale.EN_RW, locale.EN_BI, locale.ES_US, locale.ES_NI, locale.EN_VC, locale.GA, locale.SAQ, locale.EN_AE, locale.EN_NR, locale.YUE_HANT_CN, locale.MAS_TZ, locale.KDE, locale.NYN, locale.EN_NG, locale.EN_MT, locale.YO, locale.EN_BS, locale.DAV, locale.EN_BB, locale.JA, locale.EN_JM, locale.ZH_HANT, locale.EN_NA, locale.ES_419, locale.TI_ER, locale.KI, locale.EN_SX, locale.EN_DM, locale.EN_MW, locale.SO_KE, locale.EN_CM, locale.EN_CY, locale.ES_PR, locale.FIL, locale.MER, locale.EN_BW, locale.EN_GY, locale.EN_UG, locale.ES_PA, locale.EN_AS, locale.EN_PR, locale.SO, locale.YUE, locale.ZH_HANS_SG, locale.OR, locale.EN_LS, locale.EN_TO, locale.JMC, locale.EN_MS, locale.EN_NU, locale.MS, locale.UG, locale.EN_FM, locale.MR, locale.CY, locale.TI, locale.EN_MH, locale.EN_VU, locale.ES_BR, locale.EN_CK, locale.EN_FJ, locale.SI, locale.EN_NZ, locale.EN_KI, locale.TH, locale.ES_MX, locale.ES_BZ, locale.EN_LC, locale.ZH_HANT_MY, locale.EN_CX, locale.VAI_LATN, locale.KO, locale.TEO_KE, locale.ZH_HANT_HK, locale.EN_CC, locale.KO_KP, locale.EN_HK, locale.EN_SH, locale.EN_BM, locale.ZH_HANS, locale.EN_IE, locale.EN_ZM, locale.SO_DJ, locale.ZH_HANS_MO, locale.EN_CA, locale.EN_KE, locale.CHR, locale.ZH_HANS_HK, locale.EN_GM, locale.YUE_HANT, locale.EN_TT, locale.ES_HN, locale.ES_DO, locale.EN_GU, locale.ES_GT, locale.EE, locale.EN_IL, locale.EN_SG, locale.EN_SS, locale.GD, locale.YO_BJ, locale.CEB, locale.EN_PW, locale.EN_MG, locale.ES_CU, locale.EN_TV, locale.MS_SG:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 64 locales
	case locale.BG, locale.PT_ST, locale.KK, locale.CV, locale.EN_SE, locale.FF_LATN_GM, locale.HY, locale.PT_CV, locale.KY, locale.KK_KZ, locale.NMG, locale.SQ_XK, locale.KEA, locale.PT_TL, locale.BE, locale.UZ_LATN, locale.FF, locale.FF_LATN_NE, locale.SK, locale.FF_LATN_MR, locale.PT_LU, locale.RU_KG, locale.FF_LATN_GW, locale.CS, locale.FF_LATN_SL, locale.RU_KZ, locale.FF_LATN_GH, locale.PT_MO, locale.RU_UA, locale.HU, locale.KA, locale.RU_MD, locale.FF_LATN_LR, locale.PT_GW, locale.TT, locale.PT_MZ, locale.PT_PT, locale.BE_TARASK, locale.EN_FI, locale.BAS, locale.BR, locale.FF_LATN_CM, locale.TK, locale.LV, locale.RU_BY, locale.FF_LATN, locale.PL, locale.EWO, locale.RU, locale.SQ_MK, locale.FF_LATN_BF, locale.UK, locale.SQ, locale.PT_AO, locale.UZ, locale.PT_CH, locale.TG, locale.FF_LATN_NG, locale.KK_CYRL, locale.FR_CA, locale.TZM, locale.KSF, locale.PT_GQ, locale.FF_LATN_GN:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 58 locales
	case locale.BS, locale.CA_IT, locale.CA, locale.EL_POLYTON, locale.IT, locale.BS_CYRL, locale.LB, locale.EN_SI, locale.CA_AD, locale.AZ_LATN, locale.EL, locale.ES_EA, locale.FR_LU, locale.CA_ES_VALENCIA, locale.LN_CG, locale.DE_BE, locale.IT_SM, locale.DE_IT, locale.HSB, locale.LN, locale.ES_PH, locale.MK, locale.RO, locale.AZ, locale.ES, locale.EN_DK, locale.EN_BE, locale.DA, locale.IS, locale.SR_CYRL_BA, locale.ES_IC, locale.SR_CYRL, locale.SR_LATN, locale.SR_CYRL_ME, locale.SR_LATN_BA, locale.LN_AO, locale.SR_CYRL_XK, locale.CA_FR, locale.DSB, locale.VI, locale.EN_DE, locale.SR_LATN_ME, locale.RO_MD, locale.AST, locale.DA_GL, locale.KU, locale.LN_CF, locale.EL_CY, locale.SR, locale.GL, locale.IT_VA, locale.DE_LU, locale.SR_LATN_XK, locale.SC, locale.LLD, locale.DE, locale.FR_MA, locale.BS_LATN:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CH, locale.FR_GF, locale.FR_RE, locale.FR_DJ, locale.FR_TD, locale.FR_DZ, locale.FR_BE, locale.FR_PF, locale.FR_BJ, locale.FR_GQ, locale.FR_NC, locale.FR_CG, locale.FR_SY, locale.FR_WF, locale.FR_KM, locale.FR_MC, locale.FR_ML, locale.FR_TN, locale.FR_CF, locale.FR_RW, locale.FR_SN, locale.FR_MQ, locale.FR_SC, locale.FR_MG, locale.FR_GP, locale.FR_MU, locale.FR_CI, locale.FR_YT, locale.FR_PM, locale.FR_NE, locale.FR_GA, locale.FR_MF, locale.FR_HT, locale.FR_CM, locale.FR_VU, locale.FR_BI, locale.FR_CD, locale.FR_BL, locale.FR_BF, locale.FR_GN, locale.FR_TG, locale.FR, locale.FR_MR:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_QA, locale.AR_SS, locale.AR_DJ, locale.AR_SY, locale.AR_KW, locale.AR_IL, locale.AR_SD, locale.AR_SO, locale.AR_AE, locale.AR_SA, locale.AR_IQ, locale.AR_BH, locale.AR_EG, locale.AR_JO, locale.AR_PS, locale.AR_TD, locale.AR_KM, locale.AR_OM, locale.AR_YE, locale.AR, locale.AR_EH, locale.AR_ER:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.KOK, locale.MI, locale.MN, locale.MZN, locale.SW_KE, locale.MG, locale.SD, locale.LAG, locale.EN_MV, locale.SW_UG, locale.KOK_DEVA, locale.QU, locale.QU_EC, locale.ES_PE, locale.TA_SG, locale.SD_ARAB, locale.HA_GH, locale.SW, locale.HA, locale.TA_MY, locale.HA_NE:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.IA, locale.WO, locale.YRL_VE, locale.ES_AR, locale.YRL, locale.KGP, locale.JV, locale.ES_UY, locale.EN_AT, locale.QU_BO, locale.PT, locale.SW_CD, locale.MS_BN, locale.YRL_CO, locale.ES_CO:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_NG, locale.FF_ADLM_BF, locale.FF_ADLM, locale.FF_ADLM_SN, locale.FF_ADLM_GM, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_SL, locale.FF_ADLM_LR, locale.FF_ADLM_GH, locale.FF_ADLM_GW, locale.FF_ADLM_CM:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.PA, locale.BN_IN, locale.HI, locale.TE, locale.TA, locale.EN_IN, locale.XNR, locale.GU, locale.HI_LATN, locale.PA_GURU, locale.TA_LK:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_SR, locale.NL_BE, locale.NL_AW, locale.ES_PY, locale.NL, locale.EN_NL, locale.NL_SX, locale.NL_BQ, locale.NL_CW:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.TR_CY, locale.TR, locale.MUA, locale.ES_GQ, locale.ID, locale.ES_BO, locale.MS_ID, locale.EN_ID:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.ET, locale.SV_AX, locale.KSH, locale.FI, locale.LT, locale.SV, locale.SV_FI:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.SBP, locale.LG, locale.KSB, locale.LUO, locale.BEZ, locale.KM, locale.RWK:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_LY, locale.AR_LB, locale.AR_MA, locale.AR_MR, locale.AR_TN:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR_BA, locale.EU, locale.FO, locale.FO_DK, locale.SL, locale.HR:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SHI_LATN, locale.SHI, locale.KAB, locale.SHI_TFNG, locale.ZGH, locale.AGQ:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.SG, locale.ES_EC, locale.ES_VE, locale.ES_CL, locale.LO:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.XOG, locale.MY, locale.EN_150, locale.ASA, locale.CE:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_CR, locale.AF, locale.EN_ZA, locale.AF_NA:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.RM, locale.GSW_FR, locale.GSW_LI:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.KHQ, locale.TWQ, locale.DJE, locale.SES:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NO, locale.NN, locale.NB, locale.NB_SJ:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.BN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_CH, locale.DE_CH, locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LU, locale.RN, locale.SEH:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.SS, locale.SS_SZ:
		return CurrencyFormatInfo{
			Symbol:           "E",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "SWZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "एस.जेत.एल",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_SZ:
		return CurrencyFormatInfo{
			Symbol:           "E",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MFE:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߛߖ߭ߟ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "SZL",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
