// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = NPR{}

// Nepalese Rupee currency type
type NPR struct{}

func (NPR) Code() string { return "NPR" }

func (NPR) Number() string { return "524" }

func (NPR) Name() string { return "Nepalese Rupee" }

func (NPR) MinorUnits() int { return 2 }

func (NPR) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 140 locales
	case locale.HNJ_HMNP, locale.LA, locale.BYN, locale.BO, locale.NY, locale.SW_UG, locale.CCH, locale.QU, locale.MZN, locale.DV, locale.KAJ, locale.HA_ARAB_SD, locale.HNJ, locale.MNI_BENG, locale.MOH, locale.SD_DEVA, locale.TRW, locale.BA, locale.SYR_SY, locale.HA_NE, locale.VO, locale.AA, locale.AA_ER, locale.KPE, locale.LAG, locale.MAI, locale.GEZ, locale.KAA, locale.MI, locale.COP, locale.CIC, locale.ANN, locale.PAP, locale.AZ_ARAB_IQ, locale.GAA, locale.HA_ARAB, locale.SID, locale.CSW, locale.TIG, locale.JBO, locale.MN, locale.SCN, locale.CO, locale.BM_NKOO, locale.MG, locale.IU, locale.EN_SHAW, locale.SYR, locale.BO_IN, locale.LTG, locale.SW_KE, locale.MIC, locale.KPE_GN, locale.AZ_ARAB, locale.PIS, locale.QU_EC, locale.SAT_OLCK, locale.BAL_ARAB, locale.RHG_ROHG, locale.UND, locale.MN_MONG, locale.MHN, locale.RIF, locale.SSY, locale.KS_DEVA, locale.BEW, locale.AZ_ARAB_TR, locale.MNI, locale.TRV, locale.LKT, locale.SKR, locale.BGC, locale.QUC, locale.SW, locale.KAA_LATN, locale.PAP_AW, locale.EN_DSRT, locale.IU_LATN, locale.KCG, locale.TA_SG, locale.TO, locale.SMJ, locale.WAL, locale.IO, locale.CAD, locale.II, locale.SMA_NO, locale.SMS, locale.ZH_LATN, locale.MDF, locale.SHN_TH, locale.SAT, locale.KAA_CYRL, locale.MUS, locale.APC, locale.OSA, locale.KEN, locale.MNI_MTEI, locale.BAL, locale.HA_GH, locale.SMJ_NO, locale.SD, locale.SD_ARAB, locale.WBP, locale.BLT, locale.KK_ARAB, locale.MYV, locale.RHG_ROHG_BD, locale.MGO, locale.KOK_DEVA, locale.AA_DJ, locale.KOK, locale.ZA, locale.ARN, locale.YI, locale.AB, locale.CHO, locale.CKB_IR, locale.FRR, locale.NV, locale.SAT_DEVA, locale.LRC_IQ, locale.RAJ, locale.TYV, locale.BSS, locale.HA, locale.CU, locale.CKB, locale.SDH, locale.GEZ_ER, locale.LRC, locale.SHN, locale.WA, locale.GN, locale.AN, locale.RHG, locale.TA_MY, locale.ES_PE, locale.SDH_IQ, locale.SMA:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 102 locales
	case locale.ES_CU, locale.IG, locale.OM, locale.OR, locale.ES_SV, locale.KAM, locale.MR, locale.TEO_KE, locale.YUE_HANT_MO, locale.ES_NI, locale.ES_HN, locale.SO, locale.PCM, locale.BEM, locale.ES_DO, locale.KI, locale.ND, locale.ZH_HANS_MY, locale.KO, locale.ES_MX, locale.YUE, locale.AK, locale.ES_BR, locale.ZH_HANT, locale.ZH_HANT_MO, locale.EE_TG, locale.YUE_HANT_CN, locale.GA_GB, locale.FIL, locale.CGG, locale.OM_KE, locale.SN, locale.ZH, locale.TI, locale.DAV, locale.HAW, locale.GA, locale.MS_SG, locale.VAI_VAII, locale.YO_BJ, locale.ML, locale.GUZ, locale.SO_DJ, locale.MAS, locale.MT, locale.EBU, locale.AM, locale.ES_BZ, locale.NUS, locale.SAQ, locale.NAQ, locale.ES_419, locale.ZH_HANS_SG, locale.KDE, locale.JA, locale.KO_CN, locale.ST_LS, locale.UG, locale.CY, locale.MS_ARAB, locale.ZU, locale.TEO, locale.JMC, locale.GV, locale.MER, locale.TI_ER, locale.ROF, locale.BHO, locale.KW, locale.VAI_LATN, locale.GD, locale.YO, locale.YUE_HANS, locale.TH, locale.VUN, locale.YUE_HANT, locale.ES_US, locale.ZH_HANS, locale.MS, locale.BM, locale.MN_MONG_MN, locale.ES_PA, locale.SO_ET, locale.ZH_HANS_HK, locale.ST, locale.ZH_HANT_MY, locale.KLN, locale.DOI, locale.ES_PR, locale.SO_KE, locale.SI, locale.KN, locale.ZH_HANT_HK, locale.KO_KP, locale.CHR, locale.MAS_TZ, locale.ZH_HANS_MO, locale.ES_GT, locale.CEB, locale.EE, locale.NYN, locale.VAI:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 94 locales
	case locale.EN_MS, locale.EN_001, locale.EN_PW, locale.EN_SB, locale.EN_NZ, locale.EN_SZ, locale.EN_ZM, locale.EN_SH, locale.EN_IE, locale.EN_PN, locale.EN_UM, locale.EN_SG, locale.EN_GH, locale.EN_GM, locale.EN_MG, locale.EN_CY, locale.EN_NA, locale.EN_TV, locale.EN_VU, locale.EN_MU, locale.EN_SL, locale.EN_CX, locale.EN_RW, locale.EN_BB, locale.EN_MO, locale.EN_CA, locale.EN_FJ, locale.EN_SD, locale.EN_SX, locale.EN_GY, locale.EN_BI, locale.EN_MT, locale.EN_MY, locale.EN_ZW, locale.EN_BS, locale.EN_IO, locale.EN_FK, locale.EN_PK, locale.EN_JE, locale.EN_NR, locale.EN_BW, locale.EN_JM, locale.EN_AG, locale.EN_LS, locale.EN_TZ, locale.EN_HK, locale.EN_BZ, locale.EN_GU, locale.EN_VC, locale.EN_GG, locale.EN_LC, locale.EN_SC, locale.EN_DG, locale.EN_IL, locale.EN_NU, locale.EN_SS, locale.EN_UG, locale.EN_NG, locale.EN_AS, locale.EN_CC, locale.EN, locale.EN_ER, locale.EN_KY, locale.EN_BM, locale.EN_PR, locale.EN_CK, locale.EN_CM, locale.EN_AE, locale.EN_VG, locale.EN_GD, locale.EN_GI, locale.EN_GB, locale.EN_KE, locale.EN_KI, locale.EN_AU, locale.EN_FM, locale.EN_AI, locale.EN_TC, locale.EN_GS, locale.EN_TK, locale.EN_PG, locale.EN_MW, locale.EN_KN, locale.EN_LR, locale.EN_TO, locale.EN_NF, locale.EN_TT, locale.EN_MH, locale.EN_PH, locale.EN_WS, locale.EN_MP, locale.EN_VI, locale.EN_DM, locale.EN_IM:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.PT_CV, locale.BG, locale.BAS, locale.PL, locale.YAV, locale.CV, locale.DYO, locale.FF_LATN_GM, locale.FF, locale.FF_LATN_NE, locale.RU_KG, locale.KK_CYRL, locale.FR_CA, locale.RU, locale.EO, locale.NMG, locale.FF_LATN_GW, locale.KA, locale.TT, locale.PT_AO, locale.KK_KZ, locale.PT_MO, locale.HU, locale.PT_LU, locale.RU_BY, locale.UZ, locale.PT_MZ, locale.RU_MD, locale.FF_LATN_MR, locale.FF_LATN_LR, locale.FF_LATN_NG, locale.HY, locale.PT_TL, locale.TZM, locale.HT, locale.SK, locale.BE_TARASK, locale.BR, locale.FF_LATN_CM, locale.PT_GQ, locale.SAH, locale.FF_LATN_BF, locale.KK, locale.PT_GW, locale.PT_ST, locale.UZ_LATN, locale.SMN, locale.KY, locale.FF_LATN, locale.PRG, locale.CS, locale.FF_LATN_SL, locale.EWO, locale.KEA, locale.UZ_CYRL, locale.LV, locale.FF_LATN_GN, locale.TK, locale.PT_PT, locale.UK, locale.DUA, locale.TG, locale.FF_LATN_GH, locale.BE, locale.RU_KZ, locale.PT_CH, locale.RU_UA, locale.SZL, locale.KSF:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 59 locales
	case locale.VI, locale.FR_MA, locale.ES_IC, locale.EL, locale.DE_LU, locale.EL_CY, locale.CA, locale.SR, locale.IT_SM, locale.DSB, locale.LLD, locale.IS, locale.GL, locale.BS_CYRL, locale.CA_ES_VALENCIA, locale.DA_GL, locale.DE, locale.HSB, locale.ES_PH, locale.KU, locale.IT_VA, locale.AZ, locale.AZ_CYRL, locale.ES, locale.CA_IT, locale.NDS_NL, locale.DE_BE, locale.VMW, locale.DE_IT, locale.NDS, locale.SR_LATN_BA, locale.MK, locale.SR_CYRL_XK, locale.LN_CG, locale.SR_CYRL_ME, locale.SR_CYRL_BA, locale.BS, locale.CA_FR, locale.IT, locale.LB, locale.LN_CF, locale.SR_LATN_ME, locale.BS_LATN, locale.AST, locale.DA, locale.LN, locale.LIJ, locale.SC, locale.SR_LATN_XK, locale.AZ_LATN, locale.CA_AD, locale.ES_EA, locale.RO_MD, locale.RO, locale.SR_CYRL, locale.LN_AO, locale.SR_LATN, locale.EL_POLYTON, locale.FR_LU:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_SC, locale.FR_CI, locale.FR_MQ, locale.FR_BI, locale.FR_GF, locale.FR_SN, locale.FR_RW, locale.FR_NE, locale.FR_DJ, locale.FR_GP, locale.FR_MU, locale.FR_PM, locale.FR_ML, locale.FR, locale.FR_BL, locale.FR_WF, locale.FR_TN, locale.FR_MR, locale.FR_VU, locale.FR_MF, locale.FR_TG, locale.FR_CH, locale.FR_NC, locale.FR_TD, locale.FR_SY, locale.FR_CF, locale.FR_CD, locale.FR_MC, locale.FR_GN, locale.FR_GQ, locale.FR_BJ, locale.FR_CG, locale.FR_BE, locale.FR_YT, locale.FR_MG, locale.FR_DZ, locale.FR_CM, locale.FR_KM, locale.FR_RE, locale.FR_BF, locale.FR_PF, locale.FR_HT, locale.FR_GA:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_DJ, locale.AR_SS, locale.AR_OM, locale.AR_KM, locale.AR_EG, locale.AR_KW, locale.AR_IL, locale.AR_PS, locale.AR_AE, locale.AR, locale.AR_SY, locale.AR_BH, locale.AR_SD, locale.AR_SO, locale.AR_TD, locale.AR_EH, locale.AR_JO, locale.AR_IQ, locale.AR_YE, locale.AR_SA, locale.AR_QA, locale.AR_ER:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.MS_BN, locale.ES_CO, locale.WO, locale.FUR, locale.MS_ARAB_BN, locale.YRL_VE, locale.JGO, locale.JV, locale.ES_AR, locale.IA, locale.KGP, locale.ES_UY, locale.YRL, locale.NNH, locale.RW, locale.SW_CD, locale.MGH, locale.YRL_CO, locale.QU_BO, locale.KKJ, locale.PT:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 17 locales
	case locale.KXV_LATN, locale.TE, locale.KOK_LATN, locale.KXV_DEVA, locale.PA, locale.KXV_ORYA, locale.GU, locale.HI, locale.BN_IN, locale.XNR, locale.DZ, locale.SA, locale.TA, locale.KXV, locale.TA_LK, locale.KXV_TELU, locale.PA_GURU:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_SL, locale.FF_ADLM_GM, locale.FF_ADLM_NG, locale.FF_ADLM_CM, locale.FF_ADLM_MR, locale.FF_ADLM_BF, locale.FF_ADLM_GH, locale.FF_ADLM_LR, locale.FF_ADLM_NE, locale.FF_ADLM_SN, locale.FF_ADLM:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.EN_SK, locale.EN_FI, locale.SQ_MK, locale.EN_NO, locale.EN_CZ, locale.EN_SE, locale.SQ_XK, locale.EN_HU, locale.SQ, locale.EN_PT:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ES_GQ, locale.SU, locale.ID, locale.TR_CY, locale.ES_BO, locale.MUA, locale.TR, locale.SU_LATN, locale.MS_ID:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.SE_FI, locale.KSH, locale.SV_AX, locale.ET, locale.SE, locale.SV_FI, locale.SE_SE, locale.LT, locale.SV:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 8 locales
	case locale.EN_BE, locale.EN_ES, locale.EN_SI, locale.EN_IT, locale.EN_RO, locale.EN_DE, locale.EN_DK, locale.EN_PL:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_SR, locale.NL_AW, locale.NL_BQ, locale.ES_PY, locale.NL, locale.NL_SX, locale.NL_BE, locale.NL_CW:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI, locale.AGQ, locale.OC, locale.SHI_TFNG, locale.KAB, locale.OC_ES, locale.SHI_LATN, locale.ZGH:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.AF, locale.ES_CR, locale.AF_NA, locale.SS_SZ, locale.SS, locale.NR, locale.VE:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.RWK, locale.KSB, locale.KM, locale.BEZ, locale.LG, locale.LUO, locale.SBP:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_LB, locale.AR_MA, locale.AR_MR, locale.AR_TN, locale.AR_LY, locale.AR_DZ:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.FO, locale.HR_BA, locale.HR, locale.FO_DK, locale.SL, locale.EU:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SG, locale.KL, locale.LO, locale.ES_CL, locale.ES_EC, locale.ES_VE:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_OM, locale.BGN_IR, locale.BGN_AF, locale.BGN, locale.BGN_AE:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.TPI, locale.MY, locale.XOG, locale.ASA, locale.CE:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DE_AT, locale.TS, locale.OS_RU, locale.OS:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.GSW, locale.GSW_FR, locale.RM:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NN, locale.NB_SJ, locale.NO, locale.NB:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.TWQ, locale.KHQ, locale.DJE, locale.SES:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.BN, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.RN, locale.LU:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.AS, locale.BRX:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_IN, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NE_IN, locale.NE:
		return FormatInfo{
			Symbol:           "नेरू",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "NPLR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_NL:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߣߔߙ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "Rs",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "NPR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
