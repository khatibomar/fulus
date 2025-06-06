// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BHD{}

// Bahraini Dinar currency type
type BHD struct{}

func (BHD) Code() string { return "BHD" }

func (BHD) Number() string { return "048" }

func (BHD) Name() string { return "Bahraini Dinar" }

func (BHD) MinorUnits() int { return 3 }

func (BHD) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 182 locales
	case locale.EN_BZ, locale.BM, locale.KO_CN, locale.EN_MS, locale.EN_TV, locale.TH, locale.TI_ER, locale.DAV, locale.EN_SZ, locale.EE_TG, locale.EN_GY, locale.VAI_LATN, locale.EN_SS, locale.YUE_HANT_MO, locale.UG, locale.TEO, locale.EN_AS, locale.SN, locale.ZH_HANT_HK, locale.AK, locale.EN_KN, locale.SO_ET, locale.EN_SB, locale.ML, locale.KDE, locale.EN_CY, locale.EN_GM, locale.KO_KP, locale.ZH_HANT, locale.AM, locale.EN_MG, locale.EN_NG, locale.EN_IE, locale.EN_UM, locale.GA_GB, locale.VUN, locale.MAS, locale.EN_TK, locale.KN, locale.ZH, locale.EN_GU, locale.VAI_VAII, locale.ZH_HANS_MY, locale.EN_AU, locale.EN_NA, locale.EN_GH, locale.EN_MH, locale.MR, locale.EN_PN, locale.EN_ZM, locale.EN_TT, locale.NAQ, locale.EN_AE, locale.FIL, locale.ZH_HANT_MY, locale.EE, locale.EN_MP, locale.TEO_KE, locale.EN_NR, locale.ES_PA, locale.ES_SV, locale.JA, locale.EN_NF, locale.PCM, locale.EN_PK, locale.KI, locale.EN_TO, locale.EN_KY, locale.ES_419, locale.EN_BW, locale.YUE_HANT, locale.EN_ER, locale.IG, locale.EN_PH, locale.MS, locale.EN_BI, locale.EN_MW, locale.ZH_HANS_SG, locale.EN_BS, locale.EN_JE, locale.EN_LC, locale.YUE_HANS, locale.TI, locale.EN_FM, locale.EN_ZW, locale.ES_NI, locale.EN, locale.EN_BB, locale.EN_IO, locale.ROF, locale.SO, locale.VAI, locale.EN_GD, locale.EN_IM, locale.EN_UG, locale.SO_KE, locale.EN_GB, locale.YUE_HANT_CN, locale.YO, locale.EN_NZ, locale.EN_VI, locale.ES_US, locale.EN_TC, locale.ES_GT, locale.EN_MT, locale.ZH_HANS_MO, locale.EN_001, locale.EN_BM, locale.EN_PG, locale.OR, locale.EN_MY, locale.EN_VG, locale.ND, locale.EN_CC, locale.NYN, locale.EN_SG, locale.ZU, locale.EN_CX, locale.EN_JM, locale.EN_PR, locale.ES_PR, locale.EN_CA, locale.ES_BR, locale.EN_AI, locale.EN_WS, locale.EN_CK, locale.EN_TZ, locale.EN_SC, locale.ES_BZ, locale.ES_DO, locale.EN_SD, locale.EN_SL, locale.CY, locale.EN_DG, locale.EN_VU, locale.CGG, locale.EN_RW, locale.EN_IL, locale.KO, locale.ZH_HANS_HK, locale.EN_FK, locale.EN_FJ, locale.EN_SX, locale.EN_NU, locale.EN_MO, locale.KLN, locale.SI, locale.YUE, locale.CHR, locale.EN_SH, locale.JMC, locale.EN_CM, locale.EN_DM, locale.MAS_TZ, locale.ZH_HANT_MO, locale.EN_HK, locale.EN_GS, locale.ES_CU, locale.SO_DJ, locale.GUZ, locale.ZH_HANS, locale.GA, locale.EN_GG, locale.ES_MX, locale.EN_GI, locale.EN_VC, locale.SAQ, locale.EN_LS, locale.CEB, locale.EN_PW, locale.EN_KE, locale.EN_MU, locale.GD, locale.EN_AG, locale.YO_BJ, locale.EBU, locale.EN_KI, locale.EN_LR, locale.ES_HN, locale.MER, locale.MS_SG, locale.KAM:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 72 locales
	case locale.EN_PT, locale.TK, locale.SQ_XK, locale.FF_LATN_LR, locale.CV, locale.KEA, locale.UZ_LATN, locale.BE, locale.EN_HU, locale.FF_LATN_SL, locale.PT_GQ, locale.PT_TL, locale.EN_CZ, locale.KK, locale.FF_LATN_NG, locale.FF_LATN_MR, locale.PT_GW, locale.PT_CV, locale.HU, locale.FR_CA, locale.TT, locale.BR, locale.UK, locale.RU_BY, locale.FF_LATN, locale.PT_PT, locale.CS, locale.YAV, locale.SK, locale.RU_KZ, locale.TG, locale.SQ, locale.KK_CYRL, locale.NMG, locale.EWO, locale.FF, locale.FF_LATN_BF, locale.PL, locale.KY, locale.RU_MD, locale.UZ, locale.PT_MO, locale.PT_AO, locale.FF_LATN_GN, locale.FF_LATN_GM, locale.FF_LATN_NE, locale.BE_TARASK, locale.KA, locale.LV, locale.BAS, locale.PT_LU, locale.EN_SK, locale.FF_LATN_GW, locale.KSF, locale.EN_FI, locale.PT_ST, locale.EN_NO, locale.SQ_MK, locale.TZM, locale.FF_LATN_CM, locale.HT, locale.RU_UA, locale.PT_CH, locale.FF_LATN_GH, locale.BG, locale.PT_MZ, locale.RU, locale.RU_KG, locale.HY, locale.DYO, locale.EN_SE, locale.KK_KZ:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 62 locales
	case locale.EN_IT, locale.EN_SI, locale.SR_LATN, locale.RO_MD, locale.SR_CYRL_BA, locale.BS, locale.EL_POLYTON, locale.GL, locale.LN_CF, locale.ES, locale.LN_AO, locale.FR_LU, locale.FR_MA, locale.DA_GL, locale.CA_ES_VALENCIA, locale.EL_CY, locale.EN_PL, locale.IS, locale.SR_CYRL_XK, locale.IT_VA, locale.SR, locale.DE_LU, locale.EN_RO, locale.ES_PH, locale.DA, locale.HSB, locale.EL, locale.RO, locale.EN_ES, locale.CA_FR, locale.ES_IC, locale.AZ, locale.MK, locale.CA_AD, locale.CA_IT, locale.SR_CYRL_ME, locale.CA, locale.DE_BE, locale.LLD, locale.SR_CYRL, locale.SR_LATN_XK, locale.ES_EA, locale.SR_LATN_ME, locale.AST, locale.LB, locale.DSB, locale.SC, locale.EN_DK, locale.VI, locale.BS_CYRL, locale.DE_IT, locale.AZ_LATN, locale.DE, locale.KU, locale.LN, locale.IT, locale.IT_SM, locale.EN_BE, locale.BS_LATN, locale.LN_CG, locale.EN_DE, locale.SR_LATN_BA:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_BL, locale.FR_SN, locale.FR_BI, locale.FR_BE, locale.FR_YT, locale.FR_GA, locale.FR_BF, locale.FR_TN, locale.FR_ML, locale.FR_RW, locale.FR_KM, locale.FR_TG, locale.FR_BJ, locale.FR_MC, locale.FR_WF, locale.FR_CI, locale.FR_CF, locale.FR_CH, locale.FR_SY, locale.FR_DJ, locale.FR_NE, locale.EN_FR, locale.FR, locale.FR_TD, locale.FR_MF, locale.FR_CD, locale.FR_HT, locale.FR_GN, locale.FR_NC, locale.FR_PF, locale.FR_MG, locale.FR_CG, locale.FR_RE, locale.FR_PM, locale.FR_GF, locale.FR_CM, locale.FR_MU, locale.FR_VU, locale.FR_GQ, locale.FR_MR, locale.FR_SC, locale.FR_GP, locale.FR_DZ, locale.FR_MQ:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SA, locale.AR_AE, locale.AR_JO, locale.AR_SO, locale.AR_SS, locale.AR_EG, locale.AR_IL, locale.AR_EH, locale.AR_OM, locale.AR_ER, locale.AR_SY, locale.AR_SD, locale.AR_YE, locale.AR_QA, locale.AR_BH, locale.AR_PS, locale.AR, locale.AR_IQ, locale.AR_KM, locale.AR_TD, locale.AR_KW, locale.AR_DJ:
		return FormatInfo{
			Symbol:           "د.ب.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.EN_MV, locale.MZN, locale.ES_PE, locale.MI, locale.KOK_DEVA, locale.TA_MY, locale.BAL_LATN, locale.HA, locale.SD, locale.SD_ARAB, locale.MN, locale.SW, locale.SW_UG, locale.HA_GH, locale.TA_SG, locale.KOK, locale.LAG, locale.MG, locale.SW_KE, locale.QU, locale.QU_EC, locale.HA_NE:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.KGP, locale.SW_CD, locale.IA, locale.PT, locale.EN_AT, locale.ES_UY, locale.YRL, locale.ES_AR, locale.ES_CO, locale.QU_BO, locale.YRL_CO, locale.WO, locale.YRL_VE, locale.MS_BN, locale.JV:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_SN, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_GH, locale.FF_ADLM_BF, locale.FF_ADLM_SL, locale.FF_ADLM_GM, locale.FF_ADLM_LR, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_NG:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.PA_GURU, locale.GU, locale.TA, locale.TA_LK, locale.HI_LATN, locale.BN_IN, locale.TE, locale.HI, locale.PA, locale.EN_IN, locale.XNR:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_CW, locale.NL, locale.ES_PY, locale.NL_SX, locale.EN_NL, locale.NL_BE, locale.NL_AW, locale.NL_SR, locale.NL_BQ:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.EN_ID, locale.TR, locale.TR_CY, locale.MS_ID, locale.ES_GQ, locale.ES_BO, locale.MUA, locale.ID:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.ET, locale.FI, locale.SV_AX, locale.LT, locale.SV_FI, locale.KSH, locale.SV:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.KSB, locale.SBP, locale.LUO, locale.KM, locale.BEZ, locale.LG, locale.RWK:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_TN, locale.AR_LB, locale.AR_DZ, locale.AR_MA, locale.AR_MR, locale.AR_LY:
		return FormatInfo{
			Symbol:           "د.ب.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO_DK, locale.EU, locale.FO, locale.SL, locale.HR, locale.HR_BA:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SHI_LATN, locale.KAB, locale.ZGH, locale.SHI, locale.SHI_TFNG, locale.AGQ:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.CE, locale.MY, locale.ASA, locale.EN_150, locale.XOG:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.ES_VE, locale.LO, locale.ES_CL, locale.ES_EC, locale.SG:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.ES_CR, locale.EN_ZA, locale.AF:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.GSW, locale.RM, locale.GSW_FR:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.KHQ, locale.DJE, locale.SES, locale.TWQ:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NO, locale.NN, locale.NB:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.IT_CH, locale.EN_CH:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.LU, locale.RN:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS_PK, locale.PS:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BLO:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "बि.ऐत्स.दि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MFE:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߓߤߘ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "BHD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
