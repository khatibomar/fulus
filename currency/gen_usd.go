// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = USD{}

// US Dollar currency type
type USD struct{}

func (USD) Code() string { return "USD" }

func (USD) Number() string { return "840" }

func (USD) Name() string { return "US Dollar" }

func (USD) MinorUnits() int { return 2 }

func (USD) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 152 locales
	case locale.EN_BB, locale.JMC, locale.EN_SD, locale.ZH_HANS_MY, locale.EN_MG, locale.EN_SX, locale.EN_JE, locale.EN_IO, locale.YUE_HANS, locale.EN_BZ, locale.EN_CC, locale.KO_CN, locale.EN_JM, locale.YUE, locale.KLN, locale.EN_SL, locale.VAI_LATN, locale.EN_TK, locale.EN_FJ, locale.KDE, locale.SO, locale.EN_LC, locale.ZH_HANT_MO, locale.EN_IL, locale.EE, locale.KI, locale.EN_IM, locale.EN_CX, locale.EN_SB, locale.EN_GD, locale.SI, locale.EN_CM, locale.EN_LR, locale.ND, locale.EN_GM, locale.EN_PN, locale.EN_AG, locale.EN_PK, locale.MT, locale.OM_KE, locale.BEM, locale.ZH_HANS_MO, locale.EN_MO, locale.EN_SZ, locale.EN_VC, locale.KAM, locale.NYN, locale.EN_TC, locale.VUN, locale.MER, locale.KO, locale.ZH_HANS_SG, locale.EN_BM, locale.BM, locale.EN_SH, locale.VAI_VAII, locale.EN_FM, locale.NAQ, locale.EE_TG, locale.EN_GS, locale.EN_ZW, locale.EN_NF, locale.KW, locale.YUE_HANT, locale.YUE_HANT_MO, locale.EN_GB, locale.EN_KN, locale.EN_LS, locale.TI_ER, locale.ZH_HANT_HK, locale.EN_SG, locale.OM, locale.PCM, locale.SO_KE, locale.EN_VU, locale.EN_NU, locale.ES_CU, locale.ROF, locale.ZH_HANT, locale.EN_GG, locale.EN_TT, locale.CY, locale.AM, locale.EN_IE, locale.EN_BS, locale.EN_001, locale.KO_KP, locale.AK, locale.VAI, locale.EN_RW, locale.TEO_KE, locale.GV, locale.EBU, locale.EN_CA, locale.EN_TO, locale.SO_DJ, locale.ST, locale.SO_ET, locale.TEO, locale.MAS_TZ, locale.ZH_HANS_HK, locale.EN_AI, locale.EN_BW, locale.EN_FK, locale.EN_TV, locale.ES_DO, locale.EN_WS, locale.EN_CK, locale.EN_ER, locale.EN_GY, locale.MAS, locale.EN_CY, locale.EN_NG, locale.NUS, locale.TI, locale.CGG, locale.SN, locale.EN_KE, locale.EN_NR, locale.EN_SC, locale.MS_ARAB, locale.DAV, locale.EN_GH, locale.EN_UG, locale.TH, locale.EN_DM, locale.EN_HK, locale.EN_KI, locale.EN_TZ, locale.ZH_HANS, locale.EN_MT, locale.EN_MY, locale.EN_ZM, locale.EN_MU, locale.EN_VG, locale.EN_KY, locale.EN_PG, locale.EN_SS, locale.EN_MS, locale.EN_GI, locale.ZH, locale.ST_LS, locale.GUZ, locale.SAQ, locale.BHO, locale.EN_DG, locale.EN_NZ, locale.ZH_HANT_MY, locale.EN_PW, locale.YUE_HANT_CN, locale.EN_MW, locale.EN_NA:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 117 locales
	case locale.RIF, locale.TA_SG, locale.KEN, locale.KPE, locale.RHG, locale.UND, locale.MGO, locale.LTG, locale.SHN, locale.KAA, locale.KOK, locale.SID, locale.BLT, locale.LAG, locale.KOK_DEVA, locale.IU_LATN, locale.JBO, locale.EN_MV, locale.COP, locale.VO, locale.QUC, locale.ZH_LATN, locale.GEZ, locale.PAP, locale.CHO, locale.AA_ER, locale.KK_ARAB, locale.BGC, locale.SYR, locale.TYV, locale.ZA, locale.KAJ, locale.CKB_IR, locale.TO, locale.ARN, locale.KAA_CYRL, locale.RAJ, locale.YI, locale.LA, locale.TIG, locale.KPE_GN, locale.MYV, locale.SMA_NO, locale.BSS, locale.MG, locale.TRW, locale.LRC_IQ, locale.BO_IN, locale.BM_NKOO, locale.II, locale.SHN_TH, locale.MIC, locale.SSY, locale.BA, locale.MDF, locale.SW_UG, locale.SKR, locale.AZ_ARAB_IQ, locale.DV, locale.ANN, locale.SMS, locale.AA_DJ, locale.BAL, locale.BYN, locale.SAT_OLCK, locale.SAT, locale.SYR_SY, locale.WAL, locale.BAL_ARAB, locale.EN_SHAW, locale.HA_ARAB_SD, locale.SCN, locale.IO, locale.WA, locale.MHN, locale.AA, locale.KAA_LATN, locale.WBP, locale.AN, locale.CCH, locale.MNI_MTEI, locale.RHG_ROHG_BD, locale.SD_ARAB, locale.AB, locale.SD, locale.NV, locale.CKB, locale.PIS, locale.SDH, locale.BEW, locale.CU, locale.HA_ARAB, locale.IU, locale.MOH, locale.SDH_IQ, locale.GAA, locale.SW, locale.FRR, locale.KCG, locale.RHG_ROHG, locale.AZ_ARAB, locale.PAP_AW, locale.TRV, locale.CO, locale.GEZ_ER, locale.GN, locale.SMA, locale.MI, locale.CSW, locale.BO, locale.AZ_ARAB_TR, locale.NY, locale.SAT_DEVA, locale.SMJ_NO, locale.APC, locale.LRC, locale.SMJ:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.BAS, locale.FF_LATN_GN, locale.PT_PT, locale.SQ_MK, locale.FF_LATN_NE, locale.EN_FI, locale.KA, locale.FF_LATN_GW, locale.PT_MZ, locale.DUA, locale.FF_LATN_GM, locale.UZ_LATN, locale.FF_LATN, locale.EN_NO, locale.EWO, locale.FF_LATN_BF, locale.TZM, locale.PT_MO, locale.SZL, locale.NMG, locale.FF, locale.FF_LATN_CM, locale.PT_TL, locale.PRG, locale.PT_CH, locale.UZ, locale.FF_LATN_NG, locale.PT_ST, locale.SQ, locale.UZ_CYRL, locale.FF_LATN_MR, locale.TK, locale.EN_PT, locale.SAH, locale.SMN, locale.PT_GQ, locale.PT_LU, locale.PT_CV, locale.FF_LATN_SL, locale.YAV, locale.EN_HU, locale.EO, locale.EN_SE, locale.KSF, locale.SQ_XK, locale.EN_SK, locale.PT_AO, locale.KEA, locale.CS, locale.PT_GW, locale.EN_CZ, locale.FF_LATN_GH, locale.DYO, locale.FF_LATN_LR:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_BE, locale.FR_HT, locale.FR_DZ, locale.FR_BL, locale.FR_MF, locale.FR_CD, locale.FR_GN, locale.FR_RW, locale.FR_TG, locale.FR_SC, locale.FR_CH, locale.FR_MU, locale.FR_MC, locale.FR_BF, locale.FR_DJ, locale.FR_CF, locale.FR_PF, locale.FR_MQ, locale.FR_NE, locale.FR_KM, locale.FR_GQ, locale.FR_CI, locale.FR_ML, locale.FR_TN, locale.FR_BI, locale.FR_MR, locale.FR_RE, locale.FR_SN, locale.FR_MG, locale.FR_CG, locale.FR_NC, locale.FR_GA, locale.FR_PM, locale.FR_VU, locale.FR_GF, locale.FR_BJ, locale.FR_TD, locale.FR_YT, locale.FR_SY, locale.FR_WF, locale.FR, locale.FR_CM, locale.FR_GP:
		return FormatInfo{
			Symbol:           "$US",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 39 locales
	case locale.SR_CYRL_XK, locale.LN, locale.SR_CYRL, locale.LN_CG, locale.EN_SI, locale.SR, locale.SR_LATN_BA, locale.SR_CYRL_ME, locale.EN_IT, locale.AZ_CYRL, locale.EN_DK, locale.ES, locale.LIJ, locale.VI, locale.NDS, locale.SC, locale.VMW, locale.LLD, locale.AZ_LATN, locale.EN_DE, locale.SR_CYRL_BA, locale.EN_PL, locale.AZ, locale.ES_EA, locale.MK, locale.EN_BE, locale.LN_AO, locale.DA_GL, locale.ES_PH, locale.SR_LATN_ME, locale.DA, locale.NDS_NL, locale.BS_CYRL, locale.ES_IC, locale.SR_LATN, locale.EN_ES, locale.EN_RO, locale.SR_LATN_XK, locale.LN_CF:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 32 locales
	case locale.GA_GB, locale.YO, locale.HAW, locale.JA, locale.GD, locale.ES_PR, locale.UG, locale.ZU, locale.EN_GU, locale.DOI, locale.EN, locale.YO_BJ, locale.EN_BI, locale.EN_MP, locale.IG, locale.EN_MH, locale.EN_AS, locale.ES_SV, locale.EN_PR, locale.ES_US, locale.GA, locale.KN, locale.MN_MONG_MN, locale.OR, locale.EN_PH, locale.EN_AE, locale.EN_UM, locale.FIL, locale.EN_VI, locale.CHR, locale.ML, locale.MR:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_EH, locale.AR_DJ, locale.AR_IL, locale.AR_JO, locale.AR_YE, locale.AR_SO, locale.AR_SS, locale.AR_BH, locale.AR_TD, locale.AR_KW, locale.AR_QA, locale.AR_EG, locale.AR_SD, locale.AR_IQ, locale.AR_AE, locale.AR_ER, locale.AR_SA, locale.AR_SY, locale.AR_PS, locale.AR_KM, locale.AR, locale.AR_OM:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.HA, locale.OSA, locale.EN_DSRT, locale.CAD, locale.BAL_LATN, locale.MUS, locale.MN, locale.MNI_BENG, locale.HA_NE, locale.CIC, locale.QU_EC, locale.SD_DEVA, locale.MN_MONG, locale.LKT, locale.SW_KE, locale.KS_DEVA, locale.MZN, locale.MNI, locale.TA_MY, locale.MAI, locale.HA_GH:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.EN_AT, locale.ES_UY, locale.JV, locale.MGH, locale.KKJ, locale.RW, locale.YRL_CO, locale.IA, locale.KGP, locale.ES_CO, locale.ES_AR, locale.NNH, locale.YRL_VE, locale.JGO, locale.MS_ARAB_BN, locale.YRL, locale.FUR, locale.PT, locale.SW_CD:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 16 locales
	case locale.KK_CYRL, locale.RU_BY, locale.RU_UA, locale.HY, locale.RU_KG, locale.KK_KZ, locale.KK, locale.TT, locale.RU_MD, locale.BE_TARASK, locale.CV, locale.RU_KZ, locale.BE, locale.TG, locale.LV, locale.RU:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 13 locales
	case locale.DE_LU, locale.HSB, locale.DE_IT, locale.DE_BE, locale.AST, locale.KU, locale.DE, locale.EL_POLYTON, locale.GL, locale.EL_CY, locale.EL, locale.DSB, locale.LB:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 13 locales
	case locale.IS, locale.IT_VA, locale.CA_AD, locale.CA_ES_VALENCIA, locale.IT_SM, locale.CA_IT, locale.RO, locale.IT, locale.CA, locale.RO_MD, locale.CA_FR, locale.BS, locale.BS_LATN:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_SN, locale.FF_ADLM_GM, locale.FF_ADLM, locale.FF_ADLM_MR, locale.FF_ADLM_BF, locale.FF_ADLM_GH, locale.FF_ADLM_NG, locale.FF_ADLM_NE, locale.FF_ADLM_SL, locale.FF_ADLM_CM, locale.FF_ADLM_LR:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.XNR, locale.SA, locale.PA_GURU, locale.GU, locale.KXV, locale.KXV_ORYA, locale.KXV_LATN, locale.KOK_LATN, locale.KXV_DEVA, locale.PA, locale.KXV_TELU, locale.DZ:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 11 locales
	case locale.ES_BR, locale.ES_PA, locale.ES_HN, locale.ES_BZ, locale.EN_AU, locale.ES_MX, locale.ES_419, locale.MS, locale.ES_GT, locale.ES_NI, locale.MS_SG:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.OC_ES, locale.SHI_LATN, locale.OC, locale.SHI, locale.KAB, locale.ZGH, locale.SHI_TFNG, locale.AGQ:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSH, locale.SV_FI, locale.SE_FI, locale.SE, locale.SE_SE, locale.SV, locale.SV_AX:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 7 locales
	case locale.NL_AW, locale.NL_SX, locale.NL_CW, locale.NL_SR, locale.NL, locale.EN_NL, locale.NL_BE:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.TA, locale.TE, locale.HI_LATN, locale.EN_IN, locale.BN_IN, locale.HI, locale.TA_LK:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_TN, locale.AR_MR, locale.AR_LY, locale.AR_DZ, locale.AR_LB, locale.AR_MA:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.CE, locale.MY, locale.EN_150, locale.TPI, locale.XOG, locale.ASA:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.LG, locale.SBP, locale.KSB, locale.RWK, locale.BEZ, locale.LUO:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN, locale.BGN_IR, locale.BGN_AF, locale.BGN_OM, locale.BGN_AE:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.NR, locale.EN_ZA, locale.SS, locale.SS_SZ, locale.VE:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.UK, locale.SK, locale.KY, locale.PL, locale.HU:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.SES, locale.TWQ, locale.KHQ:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_GQ, locale.EN_ID, locale.MUA, locale.ID:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.LO, locale.ES_CL, locale.KL, locale.SG:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB, locale.NN, locale.NB_SJ, locale.NO:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW_FR, locale.GSW_LI, locale.GSW:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.TR_CY, locale.SU, locale.TR, locale.SU_LATN:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AF_NA, locale.ES_CR, locale.AF:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_AT, locale.OS_RU, locale.OS:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.FO, locale.FO_DK, locale.EU:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.LU, locale.SEH, locale.RN:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.AS, locale.NE:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_BO, locale.MS_ID:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.FI, locale.ET:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 2 locales
	case locale.FR_MA, locale.FR_LU:
		return FormatInfo{
			Symbol:           "$US",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.HNJ_HMNP, locale.HNJ:
		return FormatInfo{
			Symbol:           "𞅎",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.HR_BA, locale.HR:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BG:
		return FormatInfo{
			Symbol:           "щ.д.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BR:
		return FormatInfo{
			Symbol:           "$ SU",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.CEB:
		return FormatInfo{
			Symbol:           "US $",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_CH:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_EC:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_PE:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_PY:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_VE:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "$",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FR_CA:
		return FormatInfo{
			Symbol:           "$ US",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "$",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.HT:
		return FormatInfo{
			Symbol:           "$US",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.IT_CH:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.KM:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.LMO:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.LT:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.MS_BN:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.NL_BQ:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߊߞߘ$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.QU:
		return FormatInfo{
			Symbol:           "$US",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.QU_BO:
		return FormatInfo{
			Symbol:           "$US",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.SL:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.TS:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.UZ_ARAB:
		return FormatInfo{
			Symbol:           "US$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "USD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.WAE:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.WO:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
