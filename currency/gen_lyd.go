// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = LYD{}

// Libyan Dinar currency type
type LYD struct{}

func (LYD) Code() string { return "LYD" }

func (LYD) Number() string { return "434" }

func (LYD) Name() string { return "Libyan Dinar" }

func (LYD) MinorUnits() int { return 3 }

func (LYD) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 182 locales
	case locale.EN_NR, locale.TH, locale.EBU, locale.EN_SB, locale.KI, locale.EN_HK, locale.EN_MH, locale.EN_BS, locale.ML, locale.DAV, locale.EE_TG, locale.EN_LS, locale.EN_MG, locale.EN_GI, locale.GUZ, locale.EN_CK, locale.NAQ, locale.TI, locale.KAM, locale.EN_TT, locale.VAI, locale.ES_CU, locale.EN_VI, locale.SO_ET, locale.UG, locale.EN_VG, locale.SAQ, locale.EN_FM, locale.EN_NU, locale.CY, locale.YUE_HANT_CN, locale.EE, locale.EN_SZ, locale.YO, locale.EN_AI, locale.EN_IO, locale.EN_UG, locale.ZH_HANS_MO, locale.EN_KN, locale.EN_WS, locale.ES_BR, locale.VAI_VAII, locale.ES_MX, locale.YO_BJ, locale.EN_TC, locale.AM, locale.EN_SL, locale.ES_US, locale.EN_BB, locale.EN_LC, locale.EN_PR, locale.EN_VC, locale.ES_BZ, locale.TEO, locale.EN_AG, locale.KN, locale.ND, locale.SO, locale.ZH_HANT_MO, locale.EN_LR, locale.JMC, locale.EN_RW, locale.TI_ER, locale.EN_CM, locale.EN_NF, locale.ZH_HANS_SG, locale.EN_GY, locale.KO, locale.GA, locale.KO_CN, locale.EN_MT, locale.EN_TZ, locale.ES_HN, locale.YUE_HANT_MO, locale.CEB, locale.EN_JE, locale.EN_MO, locale.GA_GB, locale.EN_CA, locale.EN_SH, locale.ES_DO, locale.EN_DM, locale.EN_UM, locale.EN_VU, locale.ES_PR, locale.JA, locale.EN_001, locale.EN_MS, locale.ES_PA, locale.SI, locale.EN_CY, locale.ES_GT, locale.EN_ZW, locale.ZH_HANT, locale.EN_BW, locale.EN_IE, locale.TEO_KE, locale.VAI_LATN, locale.EN_SD, locale.YUE_HANS, locale.KDE, locale.EN_TV, locale.EN_KE, locale.EN_NA, locale.EN_GM, locale.PCM, locale.EN_IM, locale.EN_MP, locale.SO_KE, locale.ES_NI, locale.EN_KI, locale.EN_PH, locale.EN_CC, locale.VUN, locale.EN_AU, locale.EN_NG, locale.EN_ER, locale.EN_GB, locale.MAS, locale.EN_FJ, locale.EN_MW, locale.EN_BI, locale.EN_PG, locale.EN, locale.EN_GG, locale.KLN, locale.EN_MY, locale.ROF, locale.YUE_HANT, locale.EN_SX, locale.KO_KP, locale.FIL, locale.EN_GD, locale.EN_NZ, locale.MS, locale.EN_JM, locale.EN_BZ, locale.ZU, locale.EN_SC, locale.GD, locale.CHR, locale.EN_TK, locale.IG, locale.ZH_HANT_HK, locale.EN_GH, locale.EN_MU, locale.MR, locale.MAS_TZ, locale.ZH_HANS_HK, locale.EN_PN, locale.EN_BM, locale.EN_AE, locale.EN_DG, locale.EN_SS, locale.NYN, locale.EN_IL, locale.EN_SG, locale.ES_419, locale.OR, locale.ZH, locale.EN_PK, locale.EN_PW, locale.MER, locale.SO_DJ, locale.ZH_HANT_MY, locale.EN_TO, locale.SN, locale.ZH_HANS, locale.CGG, locale.AK, locale.EN_CX, locale.BM, locale.EN_AS, locale.EN_GS, locale.EN_KY, locale.EN_ZM, locale.ZH_HANS_MY, locale.MS_SG, locale.YUE, locale.ES_SV, locale.EN_GU, locale.EN_FK:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 72 locales
	case locale.SQ_MK, locale.EN_CZ, locale.FF_LATN_CM, locale.FF_LATN_LR, locale.UZ_CYRL, locale.BR, locale.PT_AO, locale.SQ, locale.CV, locale.TZM, locale.RU_KG, locale.KK_CYRL, locale.PT_MZ, locale.RU_BY, locale.FF_LATN_GN, locale.RU_UA, locale.KK, locale.EN_HU, locale.RU_MD, locale.BG, locale.KSF, locale.UK, locale.PT_PT, locale.PT_TL, locale.FF, locale.FF_LATN_SL, locale.NMG, locale.FF_LATN_GW, locale.KY, locale.PT_CV, locale.TT, locale.FF_LATN_BF, locale.PL, locale.DYO, locale.FF_LATN_NE, locale.EN_SE, locale.EN_SK, locale.PT_MO, locale.FF_LATN_GM, locale.BE_TARASK, locale.UZ_LATN, locale.HU, locale.BE, locale.TK, locale.KA, locale.FR_CA, locale.KK_KZ, locale.HY, locale.FF_LATN, locale.PT_GQ, locale.CS, locale.KEA, locale.FF_LATN_GH, locale.PT_ST, locale.BAS, locale.EN_FI, locale.HT, locale.RU_KZ, locale.UZ, locale.PT_CH, locale.EWO, locale.TG, locale.SK, locale.EN_PT, locale.FF_LATN_NG, locale.SQ_XK, locale.PT_GW, locale.PT_LU, locale.RU, locale.FF_LATN_MR, locale.LV, locale.EN_NO:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 62 locales
	case locale.DA_GL, locale.IT_VA, locale.SR_CYRL, locale.KU, locale.IT_SM, locale.EN_SI, locale.DE, locale.CA_ES_VALENCIA, locale.EL_CY, locale.SR_LATN_XK, locale.EN_RO, locale.SR_CYRL_ME, locale.LN_CF, locale.FR_LU, locale.SR, locale.DE_BE, locale.BS_LATN, locale.LN, locale.LB, locale.ES_IC, locale.HSB, locale.EN_DK, locale.AST, locale.DE_LU, locale.SR_CYRL_XK, locale.ES, locale.IS, locale.GL, locale.IT, locale.EN_BE, locale.BS, locale.EL_POLYTON, locale.LLD, locale.SR_LATN, locale.ES_EA, locale.EN_PL, locale.SC, locale.AZ, locale.SR_LATN_ME, locale.AZ_LATN, locale.EN_ES, locale.MK, locale.EN_DE, locale.ES_PH, locale.BS_CYRL, locale.RO, locale.EN_IT, locale.RO_MD, locale.EL, locale.SR_LATN_BA, locale.LN_AO, locale.LN_CG, locale.CA_AD, locale.FR_MA, locale.CA, locale.VI, locale.DE_IT, locale.SR_CYRL_BA, locale.CA_IT, locale.CA_FR, locale.DA, locale.DSB:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR, locale.FR_GP, locale.FR_BL, locale.FR_NE, locale.FR_GN, locale.FR_NC, locale.FR_PF, locale.FR_BF, locale.FR_VU, locale.FR_KM, locale.FR_GQ, locale.FR_MU, locale.FR_ML, locale.EN_FR, locale.FR_TD, locale.FR_GF, locale.FR_DJ, locale.FR_PM, locale.FR_MG, locale.FR_MC, locale.FR_MR, locale.FR_CM, locale.FR_SY, locale.FR_CF, locale.FR_RE, locale.FR_BJ, locale.FR_CD, locale.FR_BI, locale.FR_DZ, locale.FR_MF, locale.FR_RW, locale.FR_TG, locale.FR_CI, locale.FR_CG, locale.FR_YT, locale.FR_BE, locale.FR_GA, locale.FR_HT, locale.FR_CH, locale.FR_TN, locale.FR_WF, locale.FR_SC, locale.FR_SN, locale.FR_MQ:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_IL, locale.AR_SY, locale.AR_AE, locale.AR_KM, locale.AR_PS, locale.AR_KW, locale.AR_EH, locale.AR_SA, locale.AR_JO, locale.AR_OM, locale.AR_QA, locale.AR_SS, locale.AR_EG, locale.AR_IQ, locale.AR_YE, locale.AR_ER, locale.AR_BH, locale.AR_SO, locale.AR_DJ, locale.AR, locale.AR_TD, locale.AR_SD:
		return FormatInfo{
			Symbol:           "د.ل.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.SW, locale.SW_KE, locale.TA_MY, locale.ES_PE, locale.QU, locale.SW_UG, locale.MZN, locale.TA_SG, locale.KOK_DEVA, locale.QU_EC, locale.HA_GH, locale.MG, locale.SD, locale.HA_NE, locale.HA, locale.MN, locale.KOK, locale.EN_MV, locale.LAG, locale.SD_ARAB, locale.MI:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.KGP, locale.IA, locale.EN_AT, locale.SW_CD, locale.QU_BO, locale.ES_UY, locale.YRL_VE, locale.PT, locale.JV, locale.YRL_CO, locale.WO, locale.ES_CO, locale.MS_BN, locale.YRL, locale.ES_AR:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.DZ, locale.PA_GURU, locale.GU, locale.HI, locale.BN_IN, locale.TA, locale.EN_IN, locale.PA, locale.TA_LK, locale.XNR, locale.TE, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM, locale.FF_ADLM_GH, locale.FF_ADLM_SN, locale.FF_ADLM_NG, locale.FF_ADLM_LR, locale.FF_ADLM_CM, locale.FF_ADLM_GM, locale.FF_ADLM_SL, locale.FF_ADLM_NE, locale.FF_ADLM_GW, locale.FF_ADLM_MR, locale.FF_ADLM_BF:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_CW, locale.NL_AW, locale.NL_BE, locale.ES_PY, locale.NL, locale.NL_BQ, locale.NL_SR, locale.EN_NL, locale.NL_SX:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.EN_ID, locale.ES_BO, locale.ID, locale.TR, locale.TR_CY, locale.MS_ID, locale.ES_GQ, locale.MUA:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSH, locale.LT, locale.SV_AX, locale.SV, locale.ET, locale.FI, locale.SV_FI:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.LUO, locale.BEZ, locale.LG, locale.SBP, locale.RWK, locale.KSB, locale.KM:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_LB, locale.AR_DZ, locale.AR_MA, locale.AR_LY, locale.AR_MR, locale.AR_TN:
		return FormatInfo{
			Symbol:           "د.ل.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR, locale.HR_BA, locale.SL, locale.FO_DK, locale.EU, locale.FO:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SHI, locale.SHI_TFNG, locale.ZGH, locale.SHI_LATN, locale.KAB, locale.AGQ:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.MY, locale.XOG, locale.CE, locale.EN_150, locale.ASA:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.SG, locale.ES_VE, locale.ES_EC, locale.ES_CL, locale.LO:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.AF, locale.EN_ZA, locale.ES_CR:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.KHQ, locale.TWQ, locale.SES:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.RM, locale.GSW_LI, locale.GSW_FR:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NB, locale.NO:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.CCP, locale.BN:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.EN_CH, locale.DE_CH:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE, locale.AS, locale.NE_IN:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.LU, locale.RN:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "LBYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "एल.वाई.दि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MFE:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߟߓߘ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "LYD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
