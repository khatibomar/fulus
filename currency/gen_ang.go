// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = ANG{}

// Netherlands Antillean Guilder currency type
type ANG struct{}

func (ANG) Code() string { return "ANG" }

func (ANG) Number() string { return "532" }

func (ANG) Name() string { return "Netherlands Antillean Guilder" }

func (ANG) MinorUnits() int { return 2 }

func (ANG) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 154 locales
	case locale.EN_MS, locale.EN_SL, locale.EN_MT, locale.ES_419, locale.EN_GY, locale.EN_SD, locale.JA, locale.EN_SS, locale.YUE_HANT_CN, locale.EN_AU, locale.EN_FK, locale.EN_AS, locale.EN_TO, locale.ES_GT, locale.MR, locale.EN_IL, locale.EN_PW, locale.EN_001, locale.EN_IM, locale.EN_MU, locale.EN_JM, locale.EN_KY, locale.EN_LR, locale.EN_VU, locale.ZH_HANT_MY, locale.ES_NI, locale.ES_SV, locale.UG, locale.EN_GI, locale.EE, locale.EN_VC, locale.EN_KI, locale.YUE_HANS, locale.EN_NG, locale.ES_PA, locale.EN_GM, locale.EN_HK, locale.KN, locale.EN_MY, locale.ES_DO, locale.EN_SH, locale.EN_LC, locale.CEB, locale.EN_VI, locale.EN_KE, locale.YO, locale.ZH_HANS_MY, locale.ZH_HANT_MO, locale.EN_TV, locale.EN_VG, locale.EN_SC, locale.EN_GB, locale.CY, locale.EN_SB, locale.EN_CK, locale.IG, locale.YUE, locale.EN_ER, locale.ES_BR, locale.EN_AI, locale.EN_AE, locale.ZH, locale.YO_BJ, locale.EN_PG, locale.ZH_HANS_HK, locale.EN_FM, locale.EN_LS, locale.EN_MO, locale.SI, locale.EN_BZ, locale.YUE_HANT, locale.EN_NZ, locale.EN_SG, locale.ES_CU, locale.ZH_HANS_MO, locale.EN_GH, locale.EN_PK, locale.EN_UG, locale.EN_TT, locale.ZH_HANS, locale.EN_GU, locale.GD, locale.SO_ET, locale.EN_KN, locale.EN_SZ, locale.EN_PN, locale.EN_TC, locale.EN_MP, locale.EN_NA, locale.ES_US, locale.EN_NR, locale.TI, locale.KO_CN, locale.ES_HN, locale.GA_GB, locale.EN_BS, locale.EN_PH, locale.ES_BZ, locale.ES_MX, locale.FIL, locale.OR, locale.PCM, locale.SO_KE, locale.EN_FJ, locale.EN_MW, locale.EN_DG, locale.CHR, locale.MS_SG, locale.AK, locale.EN_MG, locale.TH, locale.EN_BB, locale.EN_IO, locale.EN_TZ, locale.EN_MH, locale.ZH_HANT, locale.SO_DJ, locale.EE_TG, locale.EN_ZW, locale.EN_AG, locale.EN_CY, locale.ML, locale.EN_CX, locale.MS, locale.ZH_HANS_SG, locale.EN_TK, locale.EN_WS, locale.SO, locale.AM, locale.GA, locale.ZU, locale.EN_BI, locale.ZH_HANT_HK, locale.EN_BM, locale.KO, locale.KO_KP, locale.EN_GD, locale.EN_GG, locale.EN, locale.EN_IE, locale.EN_PR, locale.TI_ER, locale.EN_CA, locale.EN_DM, locale.ES_PR, locale.EN_JE, locale.EN_ZM, locale.EN_UM, locale.EN_CC, locale.EN_NF, locale.EN_RW, locale.EN_BW, locale.EN_CM, locale.EN_NU:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.KU, locale.ES, locale.IT, locale.DE_LU, locale.IT_SM, locale.CA_FR, locale.MK, locale.SR_LATN_BA, locale.SR_LATN_XK, locale.CA_ES_VALENCIA, locale.DA_GL, locale.BS_LATN, locale.ES_IC, locale.DE_BE, locale.DSB, locale.CA_IT, locale.EL, locale.SR_LATN_ME, locale.RO_MD, locale.AZ, locale.VI, locale.DA, locale.SR_LATN, locale.LLD, locale.IT_VA, locale.EL_CY, locale.SR_CYRL, locale.EN_SI, locale.FR_LU, locale.IS, locale.ES_PH, locale.BS, locale.HSB, locale.DE_IT, locale.FR_MA, locale.EN_DK, locale.DE, locale.EL_POLYTON, locale.SR_CYRL_XK, locale.EN_DE, locale.GL, locale.AZ_LATN, locale.BS_CYRL, locale.SC, locale.SR_CYRL_BA, locale.EN_BE, locale.ES_EA, locale.SR_CYRL_ME, locale.RO, locale.SR, locale.AST, locale.LB, locale.CA_AD, locale.CA:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 46 locales
	case locale.UZ, locale.PT_MZ, locale.PT_TL, locale.RU, locale.UZ_LATN, locale.KY, locale.PT_AO, locale.TT, locale.KK, locale.TG, locale.TK, locale.BG, locale.HU, locale.PT_CH, locale.PT_GQ, locale.UZ_CYRL, locale.SQ, locale.SQ_MK, locale.BE, locale.PL, locale.PT_MO, locale.CV, locale.KK_KZ, locale.EN_FI, locale.BR, locale.PT_CV, locale.EN_SE, locale.RU_UA, locale.RU_KZ, locale.CS, locale.RU_MD, locale.UK, locale.BE_TARASK, locale.SQ_XK, locale.RU_BY, locale.HY, locale.PT_ST, locale.SK, locale.KK_CYRL, locale.FR_CA, locale.PT_LU, locale.LV, locale.PT_GW, locale.RU_KG, locale.PT_PT, locale.KA:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_YT, locale.FR_BJ, locale.FR_PM, locale.FR_GA, locale.FR_ML, locale.FR_SN, locale.FR_CD, locale.FR_CF, locale.FR_DJ, locale.FR_WF, locale.FR_BE, locale.FR_RE, locale.FR_TD, locale.FR_GQ, locale.FR_MR, locale.FR_NC, locale.FR_CG, locale.FR_GF, locale.FR_PF, locale.FR_CM, locale.FR_TG, locale.FR_TN, locale.FR_KM, locale.FR_MF, locale.FR_MU, locale.FR_MQ, locale.FR_CI, locale.FR_GN, locale.FR_BF, locale.FR_CH, locale.FR_MG, locale.FR_MC, locale.FR_RW, locale.FR_VU, locale.FR, locale.FR_SC, locale.FR_BI, locale.FR_BL, locale.FR_NE, locale.FR_HT, locale.FR_GP, locale.FR_SY, locale.FR_DZ:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_EG, locale.AR_SY, locale.AR_IL, locale.AR_KW, locale.AR_SS, locale.AR_QA, locale.AR_KM, locale.AR_TD, locale.AR_BH, locale.AR_YE, locale.AR_PS, locale.AR_IQ, locale.AR_EH, locale.AR_OM, locale.AR_SA, locale.AR_JO, locale.AR, locale.AR_SO, locale.AR_DJ, locale.AR_AE, locale.AR_ER, locale.AR_SD:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.SD_ARAB, locale.HA_NE, locale.HA_GH, locale.TA_MY, locale.SW_UG, locale.SW, locale.SW_KE, locale.SD, locale.MN, locale.HA, locale.MI, locale.KOK, locale.ES_PE, locale.QU, locale.TA_SG, locale.KOK_DEVA, locale.QU_EC, locale.MZN, locale.EN_MV:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 15 locales
	case locale.ES_AR, locale.IA, locale.SW_CD, locale.MS_BN, locale.WO, locale.ES_CO, locale.KGP, locale.YRL_VE, locale.JV, locale.PT, locale.YRL, locale.EN_AT, locale.ES_UY, locale.YRL_CO, locale.QU_BO:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_NG, locale.FF_ADLM, locale.FF_ADLM_SN, locale.FF_ADLM_BF, locale.FF_ADLM_CM, locale.FF_ADLM_LR, locale.FF_ADLM_GM, locale.FF_ADLM_SL, locale.FF_ADLM_GH:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.HI_LATN, locale.BN_IN, locale.GU, locale.EN_IN, locale.TA_LK, locale.TE, locale.TA, locale.PA, locale.PA_GURU, locale.HI, locale.XNR:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_NL, locale.NL_BQ, locale.NL_SR, locale.NL_AW, locale.ES_PY, locale.NL_BE, locale.NL:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SV, locale.SV_FI, locale.ET, locale.SV_AX, locale.LT, locale.FI, locale.KSH:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.TR, locale.TR_CY, locale.ES_GQ, locale.MS_ID, locale.EN_ID, locale.ES_BO, locale.ID:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MR, locale.AR_MA, locale.AR_TN, locale.AR_LY, locale.AR_DZ, locale.AR_LB:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR_BA, locale.EU, locale.SL, locale.FO, locale.HR, locale.FO_DK:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.ES_CR, locale.AF_NA, locale.AF, locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.GSW, locale.GSW_LI, locale.RM:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.LO, locale.ES_VE, locale.ES_CL, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NB, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.IT_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_150, locale.CE:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NL_CW, locale.NL_SX:
		return CurrencyFormatInfo{
			Symbol:           "NAf.",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "NLAG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "ए.एन.जि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_SX:
		return CurrencyFormatInfo{
			Symbol:           "NAf.",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KM:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "NAf",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߊ߲ߕߝ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "ANG",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
