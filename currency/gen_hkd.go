// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = HKD{}

// Hong Kong Dollar currency type
type HKD struct{}

func (HKD) Code() string { return "HKD" }

func (HKD) Number() string { return "344" }

func (HKD) Name() string { return "Hong Kong Dollar" }

func (HKD) MinorUnits() int { return 2 }

func (HKD) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 180 locales
	case locale.SN, locale.BM, locale.EN_CM, locale.EN_GI, locale.FIL, locale.EN_GD, locale.EN_TK, locale.KI, locale.NYN, locale.EN_BW, locale.TH, locale.EN_BI, locale.EN_KI, locale.EN_AS, locale.MT, locale.EN_NZ, locale.ST_LS, locale.SAQ, locale.EN_HK, locale.EN_PW, locale.KAM, locale.EN_UG, locale.OR, locale.EN_JE, locale.EN_VC, locale.MER, locale.EE, locale.EN_VG, locale.EN_KN, locale.EN_MT, locale.EN_VI, locale.EN_IM, locale.VAI, locale.KLN, locale.KW, locale.EN_CA, locale.EN_GH, locale.EN_AI, locale.EN_BZ, locale.EN_NG, locale.EN_TO, locale.KO_CN, locale.TI, locale.EN_LR, locale.HAW, locale.ST, locale.EN_SZ, locale.MS_SG, locale.KO_KP, locale.EBU, locale.EN_MO, locale.NUS, locale.EN_CK, locale.EN_TT, locale.EN_MH, locale.EN_NR, locale.EN_SD, locale.EN_PH, locale.VAI_VAII, locale.EN_NF, locale.EN_PN, locale.GA_GB, locale.EN_DG, locale.EN_MG, locale.EN_SG, locale.EN_MW, locale.EN_TC, locale.IG, locale.YUE_HANT, locale.EN_AE, locale.EN_UM, locale.MN_MONG_MN, locale.SO_ET, locale.EN_MP, locale.ZH_HANS_MY, locale.ND, locale.UG, locale.EN_WS, locale.SO_KE, locale.EN_001, locale.YO, locale.EN, locale.EN_NA, locale.EN_SB, locale.CY, locale.EE_TG, locale.MAS_TZ, locale.YUE_HANS, locale.EN_PR, locale.MS_ARAB, locale.EN_IE, locale.ZH_HANS, locale.BHO, locale.EN_FJ, locale.EN_KY, locale.MR, locale.CGG, locale.EN_DM, locale.EN_ER, locale.MS, locale.EN_GG, locale.YUE_HANT_CN, locale.BEM, locale.TI_ER, locale.ZH_HANS_SG, locale.ZH_HANT_MO, locale.ZH, locale.EN_FK, locale.GV, locale.JA, locale.EN_CC, locale.YUE, locale.ZH_HANS_HK, locale.EN_LS, locale.ZH_HANT_HK, locale.AK, locale.EN_CY, locale.EN_PK, locale.OM, locale.EN_BM, locale.EN_VU, locale.EN_SL, locale.EN_SS, locale.EN_ZM, locale.GD, locale.KO, locale.ZH_HANT_MY, locale.EN_KE, locale.NAQ, locale.EN_JM, locale.SI, locale.AM, locale.SO, locale.DAV, locale.EN_SX, locale.EN_PG, locale.EN_TV, locale.TEO_KE, locale.VUN, locale.EN_BS, locale.EN_TZ, locale.EN_IO, locale.YO_BJ, locale.EN_SH, locale.KN, locale.CHR, locale.VAI_LATN, locale.EN_LC, locale.DOI, locale.EN_FM, locale.EN_MS, locale.TEO, locale.ZH_HANS_MO, locale.KDE, locale.EN_GY, locale.ZU, locale.PCM, locale.EN_CX, locale.OM_KE, locale.EN_SC, locale.EN_MY, locale.ZH_HANT, locale.ROF, locale.SO_DJ, locale.EN_ZW, locale.GA, locale.EN_NU, locale.GUZ, locale.MAS, locale.EN_GB, locale.JMC, locale.EN_IL, locale.EN_RW, locale.CEB, locale.EN_AG, locale.EN_BB, locale.EN_MU, locale.EN_GM, locale.EN_GU, locale.ML:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 139 locales
	case locale.MNI_MTEI, locale.MNI_BENG, locale.DV, locale.RHG_ROHG_BD, locale.SHN, locale.CSW, locale.KEN, locale.BM_NKOO, locale.CU, locale.EN_MV, locale.KAA, locale.KOK, locale.HA_NE, locale.AB, locale.EN_SHAW, locale.MGO, locale.MN_MONG, locale.CCH, locale.LTG, locale.FRR, locale.MNI, locale.MYV, locale.VO, locale.IO, locale.RAJ, locale.AN, locale.LA, locale.TRV, locale.TRW, locale.MG, locale.SCN, locale.LKT, locale.KS_DEVA, locale.SMS, locale.HA_ARAB_SD, locale.PIS, locale.TYV, locale.CO, locale.SW_UG, locale.CHO, locale.KAJ, locale.RHG, locale.SDH, locale.AA, locale.SYR, locale.ANN, locale.PAP_AW, locale.SAT, locale.SHN_TH, locale.HA, locale.MAI, locale.WBP, locale.BAL_ARAB, locale.CAD, locale.LAG, locale.SMA, locale.SW, locale.BAL, locale.QU_EC, locale.KPE_GN, locale.RIF, locale.SMJ, locale.ZH_LATN, locale.OSA, locale.SYR_SY, locale.PAP, locale.SID, locale.EN_DSRT, locale.MHN, locale.KK_ARAB, locale.MOH, locale.YI, locale.TO, locale.AA_ER, locale.JBO, locale.MN, locale.KAA_LATN, locale.CKB_IR, locale.GEZ, locale.MI, locale.SW_KE, locale.BSS, locale.IU, locale.GAA, locale.HA_ARAB, locale.HNJ, locale.WA, locale.SAT_OLCK, locale.CKB, locale.MIC, locale.SDH_IQ, locale.BO, locale.BYN, locale.QUC, locale.HNJ_HMNP, locale.KOK_DEVA, locale.AZ_ARAB, locale.GN, locale.IU_LATN, locale.BO_IN, locale.KPE, locale.LRC_IQ, locale.SMJ_NO, locale.TIG, locale.BGC, locale.AZ_ARAB_TR, locale.SD_ARAB, locale.ARN, locale.SD, locale.MZN, locale.AZ_ARAB_IQ, locale.MDF, locale.II, locale.SMA_NO, locale.ZA, locale.BA, locale.CIC, locale.HA_GH, locale.NV, locale.RHG_ROHG, locale.SSY, locale.TA_MY, locale.LRC, locale.SKR, locale.BLT, locale.SAT_DEVA, locale.APC, locale.KAA_CYRL, locale.AA_DJ, locale.UND, locale.KCG, locale.MUS, locale.TA_SG, locale.WAL, locale.BEW, locale.GEZ_ER, locale.SD_DEVA, locale.QU, locale.NY:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 64 locales
	case locale.EN_FI, locale.PT_PT, locale.KK, locale.SAH, locale.RU_KZ, locale.DUA, locale.PT_CH, locale.FF_LATN_CM, locale.KK_CYRL, locale.FF_LATN_BF, locale.SMN, locale.FF_LATN_GN, locale.TK, locale.CS, locale.FF_LATN_NG, locale.PRG, locale.SQ_XK, locale.RU_UA, locale.FF_LATN_GH, locale.FF_LATN_GM, locale.SQ, locale.SZL, locale.RU_BY, locale.EO, locale.DYO, locale.KSF, locale.PT_GQ, locale.PT_AO, locale.UZ, locale.PT_TL, locale.UZ_LATN, locale.FF, locale.FF_LATN_GW, locale.PT_MO, locale.SQ_MK, locale.PT_LU, locale.KK_KZ, locale.PT_MZ, locale.TT, locale.NMG, locale.HY, locale.FF_LATN_SL, locale.KEA, locale.LV, locale.PT_CV, locale.BE_TARASK, locale.FF_LATN_NE, locale.PT_GW, locale.RU_MD, locale.RU_KG, locale.TG, locale.YAV, locale.BE, locale.FF_LATN_MR, locale.TZM, locale.UZ_CYRL, locale.RU, locale.BAS, locale.EN_SE, locale.PT_ST, locale.CV, locale.FF_LATN, locale.FF_LATN_LR, locale.EWO:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 49 locales
	case locale.DE, locale.DE_BE, locale.LIJ, locale.LN, locale.GL, locale.LN_CF, locale.SR_CYRL_XK, locale.DA, locale.BS_CYRL, locale.KU, locale.EL_CY, locale.AZ, locale.HSB, locale.SR_CYRL_ME, locale.NDS_NL, locale.EN_DK, locale.AST, locale.EN_DE, locale.CA_FR, locale.LB, locale.SR_LATN_XK, locale.DE_IT, locale.CA_AD, locale.NDS, locale.SR_CYRL_BA, locale.SR_LATN_BA, locale.EN_BE, locale.DSB, locale.CA, locale.SR, locale.SR_LATN_ME, locale.VI, locale.SC, locale.EL, locale.DE_LU, locale.IS, locale.AZ_LATN, locale.LN_CG, locale.SR_LATN, locale.VMW, locale.LN_AO, locale.EN_SI, locale.LLD, locale.DA_GL, locale.AZ_CYRL, locale.CA_ES_VALENCIA, locale.SR_CYRL, locale.EL_POLYTON, locale.CA_IT:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CI, locale.FR_NC, locale.FR_GA, locale.FR_GN, locale.FR_RE, locale.FR_TN, locale.FR_SN, locale.FR_SC, locale.FR_PF, locale.FR_GP, locale.FR, locale.FR_MF, locale.FR_BJ, locale.FR_TG, locale.FR_MG, locale.FR_NE, locale.FR_PM, locale.FR_CD, locale.FR_CM, locale.FR_ML, locale.FR_VU, locale.FR_GF, locale.FR_KM, locale.FR_WF, locale.FR_TD, locale.FR_MQ, locale.FR_GQ, locale.FR_RW, locale.FR_MU, locale.FR_YT, locale.FR_BL, locale.FR_DZ, locale.FR_BI, locale.FR_SY, locale.FR_MC, locale.FR_CF, locale.FR_HT, locale.FR_MR, locale.FR_BE, locale.FR_BF, locale.FR_CH, locale.FR_CG, locale.FR_DJ:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SD, locale.AR, locale.AR_JO, locale.AR_TD, locale.AR_ER, locale.AR_IQ, locale.AR_OM, locale.AR_SS, locale.AR_DJ, locale.AR_YE, locale.AR_SY, locale.AR_KM, locale.AR_EG, locale.AR_PS, locale.AR_IL, locale.AR_QA, locale.AR_AE, locale.AR_BH, locale.AR_SO, locale.AR_EH, locale.AR_SA, locale.AR_KW:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.KXV_LATN, locale.PA, locale.KXV_ORYA, locale.KXV, locale.SA, locale.KXV_DEVA, locale.XNR, locale.TA_LK, locale.KXV_TELU, locale.HI, locale.TE, locale.PA_GURU, locale.TA, locale.KOK_LATN, locale.HI_LATN, locale.DZ, locale.GU, locale.EN_IN, locale.BN_IN:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.NNH, locale.KKJ, locale.WO, locale.YRL_VE, locale.FUR, locale.MS_BN, locale.SW_CD, locale.MGH, locale.IA, locale.MS_ARAB_BN, locale.PT, locale.KGP, locale.QU_BO, locale.RW, locale.YRL_CO, locale.YRL, locale.JV, locale.JGO, locale.EN_AT:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.ES_BZ, locale.ES_CU, locale.ES_NI, locale.ES_US, locale.EN_AU, locale.ES_DO, locale.ES_HN, locale.ES_419, locale.ES_PA, locale.ES_MX, locale.ES_BR, locale.ES_GT, locale.ES_PR, locale.ES_SV:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.MK, locale.IT_VA, locale.ES_PH, locale.FR_LU, locale.FR_MA, locale.RO_MD, locale.IT, locale.IT_SM, locale.BS_LATN, locale.ES, locale.BS, locale.ES_EA, locale.RO, locale.ES_IC:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_NE, locale.FF_ADLM_SL, locale.FF_ADLM, locale.FF_ADLM_MR, locale.FF_ADLM_NG, locale.FF_ADLM_BF, locale.FF_ADLM_GM, locale.FF_ADLM_GW, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_CM:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_BQ, locale.NL_BE, locale.EN_NL, locale.NL_SX, locale.NL_AW, locale.NL_CW, locale.NL, locale.NL_SR:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI_LATN, locale.KAB, locale.AGQ, locale.SHI_TFNG, locale.OC, locale.OC_ES, locale.ZGH, locale.SHI:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.TR_CY, locale.EN_ID, locale.MUA, locale.TR, locale.ID, locale.SU, locale.SU_LATN, locale.MS_ID:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.AF, locale.VE, locale.EN_ZA, locale.SS, locale.SS_SZ, locale.NR, locale.AF_NA:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.BEZ, locale.SBP, locale.RWK, locale.LUO, locale.KM, locale.LG, locale.KSB:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KY, locale.KA, locale.UK, locale.SK, locale.HU, locale.PL, locale.BG:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_TN, locale.AR_LB, locale.AR_LY, locale.AR_MA, locale.AR_MR:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.CE, locale.TPI, locale.MY, locale.EN_150, locale.XOG, locale.ASA:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_AE, locale.BGN, locale.BGN_AF, locale.BGN_OM, locale.BGN_IR:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.LT, locale.SV, locale.SV_FI, locale.SV_AX, locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.SE_SE, locale.SE, locale.ET, locale.SE_FI, locale.KSH:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.DE_AT, locale.OS_RU, locale.OS, locale.TS:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.KHQ, locale.TWQ, locale.SES:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.EU, locale.FO_DK, locale.SL, locale.FO:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.RM, locale.GSW, locale.GSW_LI:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NN, locale.NB, locale.NB_SJ, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.BN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_UY, locale.ES_CO, locale.ES_AR:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_VE, locale.ES_CL, locale.ES_EC:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.KL, locale.SG, locale.LO:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.AS, locale.NE:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.LU, locale.RN:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.EN_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_BO, locale.ES_GQ:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.HR_BA, locale.HR:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "HGKD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BR:
		return CurrencyFormatInfo{
			Symbol:           "$ HK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "ऐत्स.के$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_CR:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_PE:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_PY:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "$HK",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "$HK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FR_CA:
		return CurrencyFormatInfo{
			Symbol:           "$ HK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߤߞߘ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "HKD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "HK$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
