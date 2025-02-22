// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = KZT{}

// Tenge currency type
type KZT struct{}

func (KZT) Code() string { return "KZT" }

func (KZT) Number() string { return "398" }

func (KZT) Name() string { return "Tenge" }

func (KZT) MinorUnits() int { return 2 }

func (KZT) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 194 locales
	case locale.EN_CY, locale.EN_NR, locale.KAM, locale.MS, locale.EN_TO, locale.HAW, locale.SAQ, locale.DOI, locale.EN_SH, locale.EN, locale.BEM, locale.EN_001, locale.EE, locale.TEO_KE, locale.EN_BB, locale.KLN, locale.KO_CN, locale.EN_UG, locale.YO, locale.EN_HK, locale.EN_LC, locale.EN_SC, locale.SO_DJ, locale.EN_IM, locale.IG, locale.EN_BI, locale.EN_WS, locale.ROF, locale.GUZ, locale.EN_BM, locale.EN_LR, locale.EN_AU, locale.EN_VI, locale.ES_419, locale.GA_GB, locale.YUE, locale.EN_SX, locale.FIL, locale.EN_BZ, locale.YUE_HANT, locale.EN_NU, locale.TH, locale.EN_ZM, locale.KN, locale.EN_PG, locale.MT, locale.EN_SD, locale.ST_LS, locale.VAI_LATN, locale.EN_VC, locale.NUS, locale.ZH, locale.EN_SZ, locale.ES_BR, locale.KDE, locale.EN_CA, locale.ML, locale.PCM, locale.EN_SL, locale.EN_TK, locale.EN_VU, locale.EN_ZW, locale.EN_CX, locale.ES_PR, locale.MS_ARAB, locale.SI, locale.CEB, locale.EN_RW, locale.ES_SV, locale.OR, locale.ZH_HANT_MO, locale.MER, locale.CHR, locale.EN_GG, locale.ZH_HANS, locale.ZH_HANT_HK, locale.EN_AS, locale.EN_FJ, locale.ES_MX, locale.VUN, locale.EN_GI, locale.EN_KI, locale.EN_VG, locale.VAI, locale.EN_JE, locale.EN_PN, locale.CGG, locale.EN_SB, locale.CY, locale.EN_ER, locale.EN_TZ, locale.EN_AI, locale.EN_FM, locale.ES_PA, locale.BM, locale.EN_KE, locale.ZH_HANS_SG, locale.GV, locale.EN_SG, locale.YUE_HANS, locale.DAV, locale.EN_MH, locale.EN_MU, locale.GD, locale.EN_GH, locale.ZH_HANS_HK, locale.ZH_HANT_MY, locale.EN_CC, locale.EN_MY, locale.KO, locale.NYN, locale.ES_GT, locale.ZH_HANS_MO, locale.EE_TG, locale.EN_PW, locale.TI, locale.YO_BJ, locale.EN_LS, locale.ES_US, locale.MN_MONG_MN, locale.EN_MT, locale.SO_KE, locale.EN_KN, locale.ZU, locale.ES_BZ, locale.EN_CK, locale.ES_DO, locale.MAS, locale.EN_NF, locale.SO, locale.KI, locale.EN_MS, locale.ZH_HANS_MY, locale.EN_GD, locale.EN_NG, locale.ES_NI, locale.YUE_HANT_CN, locale.ZH_HANT, locale.EN_GM, locale.EN_GU, locale.ND, locale.SO_ET, locale.EN_CM, locale.EN_MW, locale.EN_SS, locale.KW, locale.EN_NZ, locale.EN_PH, locale.UG, locale.OM, locale.OM_KE, locale.EN_GB, locale.EN_JM, locale.MR, locale.MS_SG, locale.EN_DM, locale.ES_HN, locale.GA, locale.EN_BS, locale.EN_IO, locale.EN_TV, locale.AM, locale.EN_TT, locale.EN_FK, locale.EN_NA, locale.EN_PR, locale.MAS_TZ, locale.EN_GY, locale.EN_BW, locale.EN_PK, locale.TI_ER, locale.AK, locale.EN_KY, locale.EN_AG, locale.NAQ, locale.JMC, locale.EN_MP, locale.KO_KP, locale.JA, locale.VAI_VAII, locale.ES_CU, locale.SN, locale.TEO, locale.EN_MG, locale.EN_MO, locale.EBU, locale.EN_DG, locale.EN_IL, locale.ST, locale.EN_AE, locale.EN_IE, locale.EN_TC, locale.EN_UM, locale.BHO:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.BA, locale.TA_MY, locale.UND, locale.IO, locale.MYV, locale.KAA_LATN, locale.KAJ, locale.MN, locale.QU_EC, locale.JBO, locale.RHG_ROHG_BD, locale.CU, locale.AZ_ARAB_IQ, locale.AA_ER, locale.HNJ, locale.BM_NKOO, locale.HNJ_HMNP, locale.SMA_NO, locale.HA_NE, locale.MG, locale.IU_LATN, locale.KAA, locale.RIF, locale.ZH_LATN, locale.LKT, locale.ES_PE, locale.MNI_BENG, locale.MOH, locale.SD, locale.KPE, locale.SD_DEVA, locale.BAL_ARAB, locale.SSY, locale.CO, locale.SHN_TH, locale.BLT, locale.NY, locale.QU, locale.WA, locale.BYN, locale.GEZ, locale.MAI, locale.MN_MONG, locale.SID, locale.MUS, locale.KEN, locale.QUC, locale.AA, locale.BSS, locale.MI, locale.KS_DEVA, locale.FRR, locale.NV, locale.CSW, locale.HA, locale.PAP, locale.SW_UG, locale.SHN, locale.TYV, locale.ZA, locale.GAA, locale.ANN, locale.BO, locale.PAP_AW, locale.SAT_DEVA, locale.CHO, locale.MNI, locale.LRC_IQ, locale.MIC, locale.SCN, locale.SD_ARAB, locale.HA_GH, locale.MGO, locale.RHG_ROHG, locale.TRV, locale.KK_ARAB, locale.SMJ, locale.SW, locale.AZ_ARAB, locale.VO, locale.CKB, locale.PIS, locale.SDH_IQ, locale.EN_SHAW, locale.GEZ_ER, locale.AZ_ARAB_TR, locale.SW_KE, locale.TA_SG, locale.BAL, locale.BO_IN, locale.BEW, locale.HA_ARAB, locale.SYR_SY, locale.EN_MV, locale.IU, locale.SYR, locale.CAD, locale.WBP, locale.LA, locale.APC, locale.AA_DJ, locale.AN, locale.KPE_GN, locale.KOK_DEVA, locale.CIC, locale.SAT_OLCK, locale.RHG, locale.AB, locale.YI, locale.MZN, locale.EN_DSRT, locale.II, locale.LRC, locale.OSA, locale.SMJ_NO, locale.SAT, locale.LAG, locale.MNI_MTEI, locale.CCH, locale.KOK, locale.HA_ARAB_SD, locale.DV, locale.ARN, locale.SDH, locale.SMA, locale.WAL, locale.TO, locale.CKB_IR, locale.SKR, locale.TIG, locale.RAJ, locale.GN, locale.KCG, locale.BGC, locale.MDF, locale.SMS, locale.KAA_CYRL, locale.TRW, locale.LTG, locale.MHN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.PT_ST, locale.UK, locale.SK, locale.RU_MD, locale.KK, locale.EO, locale.RU_UA, locale.FF_LATN_GW, locale.PT_CV, locale.KK_CYRL, locale.KK_KZ, locale.PT_LU, locale.EWO, locale.BAS, locale.FF, locale.UZ_CYRL, locale.KY, locale.FF_LATN_BF, locale.KSF, locale.DYO, locale.EN_SE, locale.PT_TL, locale.FF_LATN_CM, locale.PL, locale.FF_LATN_GN, locale.FF_LATN_LR, locale.HY, locale.RU_BY, locale.BE_TARASK, locale.EN_FI, locale.PT_MZ, locale.TZM, locale.UZ, locale.UZ_LATN, locale.PT_CH, locale.BE, locale.RU, locale.FF_LATN_GM, locale.SZL, locale.KA, locale.FF_LATN, locale.FF_LATN_NE, locale.KEA, locale.SMN, locale.FF_LATN_MR, locale.FF_LATN_NG, locale.TK, locale.PT_AO, locale.RU_KG, locale.CV, locale.PT_MO, locale.TG, locale.NMG, locale.FF_LATN_GH, locale.BR, locale.SAH, locale.YAV, locale.FF_LATN_SL, locale.PT_GQ, locale.HU, locale.PT_PT, locale.CS, locale.RU_KZ, locale.FR_CA, locale.PT_GW, locale.TT, locale.DUA, locale.LV, locale.PRG:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 63 locales
	case locale.DE_IT, locale.VI, locale.HSB, locale.LN_AO, locale.LN_CG, locale.RO, locale.AZ_LATN, locale.IS, locale.SC, locale.SR_LATN, locale.DE, locale.DE_LU, locale.CA, locale.NDS_NL, locale.BS_CYRL, locale.DSB, locale.IT_VA, locale.LN, locale.MK, locale.SR_CYRL_XK, locale.EN_DE, locale.CA_FR, locale.FR_MA, locale.LB, locale.DA, locale.SR_LATN_BA, locale.SR, locale.DE_BE, locale.LN_CF, locale.IT, locale.SR_CYRL_BA, locale.ES_EA, locale.AZ, locale.EL, locale.DA_GL, locale.EN_BE, locale.CA_ES_VALENCIA, locale.SR_CYRL, locale.EN_SI, locale.GL, locale.CA_AD, locale.KU, locale.EL_CY, locale.SR_CYRL_ME, locale.EL_POLYTON, locale.LIJ, locale.VMW, locale.BS_LATN, locale.EN_DK, locale.AZ_CYRL, locale.NDS, locale.IT_SM, locale.BS, locale.ES_PH, locale.RO_MD, locale.ES_IC, locale.AST, locale.FR_LU, locale.CA_IT, locale.LLD, locale.SR_LATN_ME, locale.ES, locale.SR_LATN_XK:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_TN, locale.FR_BI, locale.FR_SY, locale.FR_BL, locale.FR_SC, locale.FR_MR, locale.FR_NC, locale.FR_PF, locale.FR_VU, locale.FR_CF, locale.FR, locale.FR_BE, locale.FR_DZ, locale.FR_CG, locale.FR_TG, locale.FR_HT, locale.FR_TD, locale.FR_MU, locale.FR_WF, locale.FR_NE, locale.FR_MC, locale.FR_CM, locale.FR_GA, locale.FR_YT, locale.FR_KM, locale.FR_GN, locale.FR_MF, locale.FR_GQ, locale.FR_MQ, locale.FR_MG, locale.FR_DJ, locale.FR_CH, locale.FR_ML, locale.FR_GP, locale.FR_RW, locale.FR_SN, locale.FR_BJ, locale.FR_CI, locale.FR_PM, locale.FR_GF, locale.FR_RE, locale.FR_CD, locale.FR_BF:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_KM, locale.AR_SA, locale.AR_QA, locale.AR_DJ, locale.AR_EH, locale.AR_SY, locale.AR_SO, locale.AR_SD, locale.AR_BH, locale.AR_EG, locale.AR_IL, locale.AR_AE, locale.AR_KW, locale.AR_YE, locale.AR_IQ, locale.AR_OM, locale.AR_JO, locale.AR_SS, locale.AR, locale.AR_ER, locale.AR_TD, locale.AR_PS:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.SW_CD, locale.KGP, locale.RW, locale.JGO, locale.ES_CO, locale.MGH, locale.NNH, locale.MS_BN, locale.ES_UY, locale.PT, locale.KKJ, locale.MS_ARAB_BN, locale.YRL_CO, locale.EN_AT, locale.WO, locale.QU_BO, locale.YRL, locale.YRL_VE, locale.JV, locale.IA, locale.ES_AR, locale.FUR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.KXV_ORYA, locale.KXV_LATN, locale.PA_GURU, locale.TE, locale.DZ, locale.SA, locale.KXV_DEVA, locale.TA_LK, locale.XNR, locale.KXV_TELU, locale.TA, locale.KXV, locale.PA, locale.GU, locale.BN_IN, locale.HI_LATN, locale.HI, locale.KOK_LATN, locale.EN_IN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_NG, locale.FF_ADLM_GW, locale.FF_ADLM_GM, locale.FF_ADLM_SL, locale.FF_ADLM_NE, locale.FF_ADLM_BF, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM, locale.FF_ADLM_CM, locale.FF_ADLM_MR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.TR_CY, locale.ES_BO, locale.EN_ID, locale.ES_GQ, locale.MS_ID, locale.SU_LATN, locale.ID, locale.SU, locale.MUA, locale.TR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.EN_NL, locale.NL_BQ, locale.NL_CW, locale.NL, locale.NL_AW, locale.ES_PY, locale.NL_SX, locale.NL_BE, locale.NL_SR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.SV, locale.ET, locale.SE_SE, locale.SV_FI, locale.SE_FI, locale.LT, locale.SV_AX, locale.KSH, locale.SE:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 8 locales
	case locale.SS, locale.AF, locale.NR, locale.EN_ZA, locale.SS_SZ, locale.AF_NA, locale.VE, locale.ES_CR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.ZGH, locale.SHI_TFNG, locale.OC_ES, locale.SHI_LATN, locale.OC, locale.KAB, locale.SHI, locale.AGQ:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KM, locale.LG, locale.SBP, locale.LUO, locale.KSB, locale.BEZ, locale.RWK:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_TN, locale.AR_MR, locale.AR_LB, locale.AR_MA, locale.AR_LY:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ES_CL, locale.SG, locale.LO, locale.ES_VE, locale.KL, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.FO, locale.HR, locale.HR_BA, locale.SL, locale.FO_DK, locale.EU:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.TPI, locale.EN_150, locale.CE, locale.MY, locale.XOG, locale.ASA:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_AF, locale.BGN_IR, locale.BGN_OM, locale.BGN_AE, locale.BGN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DE_AT, locale.OS, locale.TS, locale.OS_RU:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.RM, locale.GSW_LI, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB, locale.NO, locale.NB_SJ, locale.NN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.SES, locale.TWQ, locale.KHQ, locale.DJE:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.SQ_XK, locale.SQ, locale.BG, locale.SQ_MK:
		return CurrencyFormatInfo{
			Symbol:           "KZT",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LU, locale.SEH, locale.RN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.PS_PK, locale.UZ_ARAB, locale.PS:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "KZKT",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "के.जेत.ति",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "KZT",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߞߗ߭ߕ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "₸",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
