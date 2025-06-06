// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = ERN{}

// Nakfa currency type
type ERN struct{}

func (ERN) Code() string { return "ERN" }

func (ERN) Number() string { return "232" }

func (ERN) Name() string { return "Nakfa" }

func (ERN) MinorUnits() int { return 2 }

func (ERN) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 180 locales
	case locale.EN_JM, locale.EN_TO, locale.EN_VC, locale.ML, locale.YUE_HANS, locale.EN_BW, locale.EN_MT, locale.EN_BZ, locale.ZH_HANT, locale.AK, locale.EN_FK, locale.EN_JE, locale.EN_KI, locale.ZH, locale.KDE, locale.CGG, locale.EN_AS, locale.EN_LR, locale.VUN, locale.AM, locale.CHR, locale.EN_GS, locale.OR, locale.UG, locale.YUE_HANT, locale.SI, locale.EN_GD, locale.EN_BB, locale.EN_PN, locale.ZH_HANT_MY, locale.ES_HN, locale.ZH_HANS_HK, locale.EN_UG, locale.ZH_HANS, locale.EN_VI, locale.KN, locale.EN_GU, locale.MS_SG, locale.EN_BM, locale.EN_NZ, locale.EN_TK, locale.EN_IO, locale.YO_BJ, locale.ES_US, locale.EN_GH, locale.VAI, locale.ZH_HANT_HK, locale.ES_SV, locale.EN_MS, locale.ND, locale.SO_KE, locale.EBU, locale.KO_KP, locale.EN_KN, locale.EN_SZ, locale.JA, locale.NYN, locale.ROF, locale.EE, locale.EN_BS, locale.JMC, locale.EN_SL, locale.ES_GT, locale.EN_TT, locale.KAM, locale.EN_MU, locale.EN_MP, locale.SAQ, locale.BM, locale.ES_PA, locale.PCM, locale.EN_MO, locale.EN_NU, locale.DAV, locale.EN_MW, locale.EN_MY, locale.EN_NG, locale.EN_CX, locale.EN_NF, locale.MR, locale.EN_CA, locale.TH, locale.EN_FM, locale.NAQ, locale.EN_LC, locale.GA, locale.EN_SX, locale.SO, locale.EN_IL, locale.EN_ZW, locale.YUE_HANT_CN, locale.ZU, locale.ES_MX, locale.EN_HK, locale.YO, locale.IG, locale.EN_GM, locale.EN_VU, locale.YUE, locale.ZH_HANS_SG, locale.EN, locale.EN_FJ, locale.MER, locale.YUE_HANT_MO, locale.KO_CN, locale.EN_IE, locale.EN_SH, locale.EN_VG, locale.FIL, locale.KO, locale.MAS_TZ, locale.EN_TC, locale.SN, locale.EN_SD, locale.EN_CK, locale.EN_001, locale.SO_DJ, locale.EN_NR, locale.KLN, locale.KI, locale.SO_ET, locale.EN_GI, locale.TEO_KE, locale.EN_TV, locale.VAI_LATN, locale.EN_CM, locale.EN_WS, locale.ES_PR, locale.EN_AU, locale.ZH_HANS_MY, locale.CY, locale.ES_BZ, locale.ES_CU, locale.ZH_HANS_MO, locale.EN_SS, locale.EN_GB, locale.EN_SG, locale.EN_TZ, locale.ES_DO, locale.EN_AI, locale.EN_BI, locale.EN_RW, locale.MAS, locale.EN_AG, locale.ES_BR, locale.TI, locale.EN_KY, locale.MS, locale.EN_AE, locale.EN_PW, locale.TEO, locale.EN_PG, locale.EN_SC, locale.VAI_VAII, locale.ES_NI, locale.EN_ZM, locale.EN_IM, locale.EN_SB, locale.EN_MH, locale.EN_PH, locale.ZH_HANT_MO, locale.EN_CC, locale.EN_NA, locale.GD, locale.CEB, locale.EN_GY, locale.EN_DG, locale.EE_TG, locale.EN_CY, locale.EN_DM, locale.GA_GB, locale.EN_PK, locale.ES_419, locale.EN_KE, locale.EN_GG, locale.EN_UM, locale.EN_LS, locale.EN_PR, locale.GUZ, locale.EN_MG:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 72 locales
	case locale.FF_LATN_SL, locale.LV, locale.SQ, locale.EWO, locale.UZ, locale.KEA, locale.TG, locale.EN_SK, locale.FR_CA, locale.UZ_LATN, locale.CS, locale.FF_LATN_GW, locale.FF_LATN_NG, locale.BAS, locale.SK, locale.HU, locale.PT_ST, locale.KK, locale.PT_GQ, locale.RU_UA, locale.EN_NO, locale.FF_LATN_BF, locale.RU_BY, locale.PL, locale.EN_CZ, locale.FF_LATN_GM, locale.HT, locale.PT_CV, locale.BR, locale.KA, locale.FF_LATN_LR, locale.PT_GW, locale.BE_TARASK, locale.NMG, locale.FF_LATN_MR, locale.PT_MO, locale.TZM, locale.TT, locale.CV, locale.UK, locale.RU, locale.BE, locale.HY, locale.SQ_XK, locale.FF_LATN_GH, locale.RU_MD, locale.FF, locale.SQ_MK, locale.TK, locale.PT_LU, locale.PT_TL, locale.PT_PT, locale.RU_KG, locale.EN_HU, locale.PT_MZ, locale.EN_SE, locale.EN_FI, locale.FF_LATN, locale.RU_KZ, locale.BG, locale.KK_KZ, locale.KK_CYRL, locale.KY, locale.EN_PT, locale.DYO, locale.FF_LATN_CM, locale.PT_CH, locale.KSF, locale.PT_AO, locale.YAV, locale.FF_LATN_GN, locale.FF_LATN_NE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 62 locales
	case locale.CA, locale.KU, locale.SR, locale.CA_ES_VALENCIA, locale.EL_CY, locale.SR_LATN, locale.FR_MA, locale.IT, locale.MK, locale.EL_POLYTON, locale.SC, locale.EN_BE, locale.LLD, locale.EN_SI, locale.IT_VA, locale.AZ_LATN, locale.DE_LU, locale.ES_PH, locale.DE_IT, locale.VI, locale.AZ, locale.IT_SM, locale.ES, locale.CA_IT, locale.BS, locale.RO_MD, locale.IS, locale.SR_LATN_ME, locale.DE_BE, locale.EL, locale.EN_ES, locale.SR_CYRL_XK, locale.CA_AD, locale.FR_LU, locale.SR_CYRL, locale.SR_LATN_XK, locale.EN_DE, locale.GL, locale.DE, locale.SR_CYRL_ME, locale.SR_LATN_BA, locale.ES_EA, locale.ES_IC, locale.BS_LATN, locale.BS_CYRL, locale.AST, locale.EN_DK, locale.LB, locale.DSB, locale.CA_FR, locale.DA_GL, locale.EN_PL, locale.LN_AO, locale.LN_CG, locale.LN_CF, locale.RO, locale.SR_CYRL_BA, locale.DA, locale.LN, locale.EN_IT, locale.HSB, locale.EN_RO:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_VU, locale.FR_CH, locale.FR_GQ, locale.FR_DJ, locale.FR_MR, locale.FR_MG, locale.FR_GP, locale.FR_NC, locale.FR_NE, locale.FR_ML, locale.FR_RW, locale.FR_GF, locale.FR_CD, locale.FR_CM, locale.FR, locale.FR_MU, locale.FR_PM, locale.FR_BI, locale.FR_MC, locale.FR_KM, locale.FR_SC, locale.FR_SN, locale.FR_TN, locale.FR_GN, locale.FR_RE, locale.FR_PF, locale.FR_BJ, locale.FR_SY, locale.FR_CF, locale.FR_CI, locale.FR_CG, locale.FR_DZ, locale.FR_WF, locale.FR_YT, locale.FR_BL, locale.FR_GA, locale.FR_BF, locale.FR_TG, locale.FR_TD, locale.EN_FR, locale.FR_BE, locale.FR_HT, locale.FR_MF, locale.FR_MQ:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 21 locales
	case locale.AR_EH, locale.AR_OM, locale.AR_PS, locale.AR_EG, locale.AR_IQ, locale.AR_IL, locale.AR_SY, locale.AR_DJ, locale.AR_KM, locale.AR_SS, locale.AR_KW, locale.AR_SO, locale.AR_QA, locale.AR_JO, locale.AR_YE, locale.AR_AE, locale.AR_SD, locale.AR_TD, locale.AR_BH, locale.AR_SA, locale.AR:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.KOK, locale.HA_GH, locale.TA_SG, locale.QU_EC, locale.EN_MV, locale.MN, locale.SD_ARAB, locale.QU, locale.SW, locale.SW_UG, locale.MZN, locale.ES_PE, locale.HA, locale.HA_NE, locale.TA_MY, locale.MI, locale.MG, locale.SD, locale.SW_KE, locale.KOK_DEVA, locale.LAG:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.QU_BO, locale.MS_BN, locale.YRL_CO, locale.ES_UY, locale.JV, locale.KGP, locale.YRL, locale.EN_AT, locale.SW_CD, locale.ES_AR, locale.IA, locale.ES_CO, locale.YRL_VE, locale.PT, locale.WO:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_MR, locale.FF_ADLM_BF, locale.FF_ADLM_NE, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM, locale.FF_ADLM_CM, locale.FF_ADLM_SL, locale.FF_ADLM_GW, locale.FF_ADLM_GM, locale.FF_ADLM_NG, locale.FF_ADLM_LR:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.TA, locale.XNR, locale.BN_IN, locale.TA_LK, locale.HI, locale.GU, locale.PA, locale.PA_GURU, locale.EN_IN, locale.HI_LATN, locale.TE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.NL_BE, locale.NL_BQ, locale.NL_CW, locale.NL_SR, locale.EN_NL, locale.NL_AW, locale.ES_PY, locale.NL_SX:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.TR, locale.ES_GQ, locale.MUA, locale.ID, locale.MS_ID, locale.TR_CY, locale.EN_ID, locale.ES_BO:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.FI, locale.SV_FI, locale.KSH, locale.ET, locale.LT, locale.SV_AX, locale.SV:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.RWK, locale.SBP, locale.LUO, locale.KM, locale.BEZ, locale.KSB, locale.LG:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_LB, locale.AR_TN, locale.AR_LY, locale.AR_MA, locale.AR_MR:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.FO, locale.SL, locale.HR_BA, locale.EU, locale.HR:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.KAB, locale.ZGH, locale.SHI_LATN, locale.AGQ, locale.SHI, locale.SHI_TFNG:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BYN, locale.GEZ_ER, locale.SSY, locale.TIG, locale.AA_ER:
		return FormatInfo{
			Symbol:           "Nfk",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.LO, locale.ES_CL, locale.SG, locale.ES_EC, locale.ES_VE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.XOG, locale.CE, locale.EN_150, locale.ASA, locale.MY:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.EN_ZA, locale.AF, locale.AF_NA, locale.ES_CR:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.GSW_FR, locale.RM, locale.GSW:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NB, locale.NO, locale.NN:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.SES, locale.KHQ, locale.TWQ, locale.DJE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.EN_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.RN, locale.LU, locale.SEH:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_ER, locale.TI_ER:
		return FormatInfo{
			Symbol:           "Nfk",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.AR_ER:
		return FormatInfo{
			Symbol:           "Nfk",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "ERTN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "इ.आर.एन",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MFE:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߋߙߝ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "ERN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
