// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = CDF{}

// Congolese Franc currency type
type CDF struct{}

func (CDF) Code() string { return "CDF" }

func (CDF) Number() string { return "976" }

func (CDF) Name() string { return "Congolese Franc" }

func (CDF) MinorUnits() int { return 2 }

func (CDF) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 180 locales
	case locale.EN_001, locale.ES_PA, locale.KO_CN, locale.MER, locale.EN_LC, locale.GA, locale.YUE_HANT_CN, locale.ZH_HANT, locale.EN_VC, locale.EN_VI, locale.EN_SG, locale.EN_VU, locale.EN_TK, locale.ZH, locale.EN_TT, locale.EN_NZ, locale.EN_SH, locale.EN_ZW, locale.ES_GT, locale.GUZ, locale.EN_RW, locale.EN_CK, locale.EN_SC, locale.ZH_HANT_MO, locale.YUE, locale.EN_PH, locale.KAM, locale.EN_DG, locale.EN_SD, locale.ZU, locale.EN_GU, locale.EN_WS, locale.BM, locale.ES_BR, locale.SO, locale.ZH_HANT_MY, locale.EN_FJ, locale.EN_GH, locale.ES_US, locale.NAQ, locale.EN_BZ, locale.MS, locale.ES_HN, locale.VUN, locale.MAS, locale.ML, locale.EN_GM, locale.EN_JM, locale.EN_VG, locale.EN_BW, locale.JA, locale.TI_ER, locale.EN_FK, locale.EN_LR, locale.PCM, locale.EN_MO, locale.EN_PW, locale.EN_TV, locale.ES_BZ, locale.SAQ, locale.ZH_HANS, locale.EN_TZ, locale.CGG, locale.EN_KY, locale.TI, locale.ZH_HANT_HK, locale.EN_CC, locale.EN_MT, locale.ES_DO, locale.KI, locale.KLN, locale.KO, locale.EN_AU, locale.EN_BS, locale.EN_GD, locale.ES_NI, locale.EN_IE, locale.EN_JE, locale.EN_MH, locale.EN_AS, locale.EN_PG, locale.KO_KP, locale.YO_BJ, locale.EN_KE, locale.EN_NF, locale.GA_GB, locale.ZH_HANS_MY, locale.ES_MX, locale.TEO, locale.EN_UG, locale.NYN, locale.AM, locale.VAI_VAII, locale.EN_GI, locale.YUE_HANS, locale.SN, locale.EN_MU, locale.EN_SL, locale.ES_SV, locale.EN_NR, locale.EN_AI, locale.EN_MS, locale.GD, locale.EE_TG, locale.EN_KN, locale.EN_MY, locale.SI, locale.ZH_HANS_MO, locale.EN_LS, locale.EN_DM, locale.YO, locale.EN_IM, locale.EN_NA, locale.EN, locale.SO_KE, locale.ZH_HANS_HK, locale.ZH_HANS_SG, locale.ROF, locale.EN_SZ, locale.EN_PK, locale.EN_SB, locale.EN_SS, locale.VAI, locale.EN_IL, locale.EN_PR, locale.EN_ER, locale.EN_GY, locale.EN_TC, locale.JMC, locale.MS_SG, locale.ND, locale.SO_ET, locale.EN_HK, locale.DAV, locale.EN_BM, locale.EN_CM, locale.EN_AE, locale.ES_419, locale.TEO_KE, locale.AK, locale.EN_MW, locale.EN_NU, locale.EN_ZM, locale.ES_PR, locale.SO_DJ, locale.EN_GG, locale.UG, locale.CHR, locale.EN_MG, locale.EN_BI, locale.EN_IO, locale.OR, locale.MAS_TZ, locale.KDE, locale.EBU, locale.EN_FM, locale.EN_GB, locale.EN_UM, locale.FIL, locale.CY, locale.EN_AG, locale.CEB, locale.EN_CY, locale.EN_SX, locale.EN_MP, locale.TH, locale.EE, locale.EN_NG, locale.VAI_LATN, locale.EN_BB, locale.EN_CX, locale.YUE_HANT, locale.EN_CA, locale.EN_TO, locale.KN, locale.MR, locale.EN_PN, locale.ES_CU, locale.IG, locale.EN_KI:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 66 locales
	case locale.CV, locale.FF_LATN, locale.BE, locale.FF_LATN_GW, locale.FF_LATN_NE, locale.SQ_XK, locale.TG, locale.RU_MD, locale.PT_CH, locale.KK_CYRL, locale.BR, locale.FF_LATN_GN, locale.FF_LATN_BF, locale.FF_LATN_LR, locale.UZ, locale.FF_LATN_SL, locale.PT_GW, locale.EN_SE, locale.HY, locale.TZM, locale.PT_ST, locale.FF_LATN_GM, locale.KK_KZ, locale.FR_CA, locale.KA, locale.HU, locale.RU, locale.KSF, locale.SQ_MK, locale.PT_MZ, locale.PT_TL, locale.FF_LATN_CM, locale.UK, locale.SK, locale.FF, locale.UZ_LATN, locale.FF_LATN_GH, locale.FF_LATN_NG, locale.EN_FI, locale.RU_KZ, locale.PT_AO, locale.BAS, locale.KEA, locale.LV, locale.FF_LATN_MR, locale.PT_MO, locale.RU_KG, locale.RU_UA, locale.CS, locale.TK, locale.BG, locale.PT_CV, locale.YAV, locale.KK, locale.KY, locale.NMG, locale.PT_LU, locale.PT_PT, locale.TT, locale.PL, locale.PT_GQ, locale.EWO, locale.BE_TARASK, locale.DYO, locale.RU_BY, locale.SQ:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.BS_CYRL, locale.SR_LATN_BA, locale.IS, locale.GL, locale.SR_CYRL_ME, locale.SR_LATN_ME, locale.BS, locale.EN_DK, locale.DA, locale.EL, locale.ES, locale.IT_SM, locale.CA_FR, locale.DA_GL, locale.IT_VA, locale.EL_POLYTON, locale.SR_CYRL, locale.IT, locale.DE, locale.SC, locale.CA_AD, locale.DSB, locale.SR_LATN_XK, locale.RO_MD, locale.HSB, locale.SR, locale.FR_MA, locale.CA, locale.CA_IT, locale.MK, locale.ES_PH, locale.SR_CYRL_XK, locale.VI, locale.RO, locale.KU, locale.EN_BE, locale.DE_BE, locale.LLD, locale.LB, locale.AZ, locale.CA_ES_VALENCIA, locale.SR_CYRL_BA, locale.FR_LU, locale.EN_DE, locale.ES_EA, locale.AST, locale.EN_SI, locale.BS_LATN, locale.DE_IT, locale.SR_LATN, locale.ES_IC, locale.AZ_LATN, locale.DE_LU, locale.EL_CY:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 42 locales
	case locale.FR_DJ, locale.FR_SN, locale.FR_MU, locale.FR_MR, locale.FR, locale.FR_CG, locale.FR_TG, locale.FR_CH, locale.FR_MG, locale.FR_MC, locale.FR_BF, locale.FR_GP, locale.FR_PF, locale.FR_RE, locale.FR_BJ, locale.FR_NC, locale.FR_HT, locale.FR_ML, locale.FR_YT, locale.FR_GF, locale.FR_WF, locale.FR_SY, locale.FR_KM, locale.FR_BI, locale.FR_MQ, locale.FR_MF, locale.FR_DZ, locale.FR_CI, locale.FR_GN, locale.FR_BL, locale.FR_VU, locale.FR_BE, locale.FR_NE, locale.FR_TD, locale.FR_CM, locale.FR_GQ, locale.FR_GA, locale.FR_TN, locale.FR_RW, locale.FR_SC, locale.FR_PM, locale.FR_CF:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SD, locale.AR_SO, locale.AR_PS, locale.AR_EG, locale.AR_DJ, locale.AR_SS, locale.AR_OM, locale.AR_IQ, locale.AR_TD, locale.AR_EH, locale.AR_IL, locale.AR_SY, locale.AR_KM, locale.AR_JO, locale.AR, locale.AR_YE, locale.AR_BH, locale.AR_ER, locale.AR_QA, locale.AR_AE, locale.AR_SA, locale.AR_KW:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.LAG, locale.HA, locale.MN, locale.MZN, locale.MG, locale.QU_EC, locale.KOK_DEVA, locale.MI, locale.QU, locale.SW_KE, locale.TA_SG, locale.TA_MY, locale.EN_MV, locale.SW, locale.SW_UG, locale.HA_GH, locale.ES_PE, locale.SD, locale.HA_NE, locale.KOK, locale.SD_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.KGP, locale.YRL_VE, locale.MS_BN, locale.IA, locale.WO, locale.ES_AR, locale.JV, locale.QU_BO, locale.YRL, locale.ES_UY, locale.YRL_CO, locale.PT, locale.EN_AT, locale.ES_CO:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM, locale.FF_ADLM_GM, locale.FF_ADLM_NE, locale.FF_ADLM_LR, locale.FF_ADLM_BF, locale.FF_ADLM_SL, locale.FF_ADLM_NG, locale.FF_ADLM_SN, locale.FF_ADLM_CM, locale.FF_ADLM_GH, locale.FF_ADLM_GW, locale.FF_ADLM_MR:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.PA_GURU, locale.TA, locale.BN_IN, locale.TE, locale.TA_LK, locale.EN_IN, locale.PA, locale.HI_LATN, locale.GU, locale.HI, locale.XNR:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.NL_BE, locale.NL_SX, locale.EN_NL, locale.NL_SR, locale.NL_CW, locale.NL_AW, locale.ES_PY, locale.NL_BQ:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.ES_BO, locale.ES_GQ, locale.TR_CY, locale.ID, locale.MS_ID, locale.MUA, locale.TR, locale.EN_ID:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.RWK, locale.KM, locale.KSB, locale.LUO, locale.SBP, locale.BEZ, locale.LG:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV_FI, locale.SV_AX, locale.ET, locale.SV, locale.LT, locale.FI, locale.KSH:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.AR_LY, locale.AR_MA, locale.AR_LB, locale.AR_DZ, locale.AR_MR, locale.AR_TN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR, locale.SL, locale.FO, locale.FO_DK, locale.EU, locale.HR_BA:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SHI, locale.AGQ, locale.KAB, locale.SHI_TFNG, locale.ZGH, locale.SHI_LATN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.ASA, locale.MY, locale.EN_150, locale.CE, locale.XOG:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.ES_VE, locale.LO, locale.ES_CL, locale.SG, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.AF_NA, locale.AF, locale.ES_CR, locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.RM, locale.GSW_LI, locale.GSW:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.LN_AO, locale.LN_CF, locale.LN_CG, locale.LN:
		return CurrencyFormatInfo{
			Symbol:           "FC",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB, locale.NB_SJ, locale.NO, locale.NN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.TWQ, locale.DJE, locale.SES, locale.KHQ:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.CCP, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_CH, locale.IT_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS_PK, locale.PS:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.RN, locale.SEH:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "KNGF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "सि.दि.एफ",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FR_CD:
		return CurrencyFormatInfo{
			Symbol:           "FC",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LU:
		return CurrencyFormatInfo{
			Symbol:           "FC",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MFE:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߞߝ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.SW_CD:
		return CurrencyFormatInfo{
			Symbol:           "FC",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "CDF",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
