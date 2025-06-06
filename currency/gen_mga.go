// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = MGA{}

// Malagasy Ariary currency type
type MGA struct{}

func (MGA) Code() string { return "MGA" }

func (MGA) Number() string { return "969" }

func (MGA) Name() string { return "Malagasy Ariary" }

func (MGA) MinorUnits() int { return 2 }

func (MGA) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 140 locales
	case locale.ES_PE, locale.KCG, locale.KOK_DEVA, locale.MOH, locale.GEZ_ER, locale.SAT, locale.KS_DEVA, locale.KPE, locale.LRC_IQ, locale.KAA_CYRL, locale.KK_ARAB, locale.NY, locale.KEN, locale.TRW, locale.MAI, locale.RAJ, locale.ANN, locale.CKB_IR, locale.FRR, locale.ARN, locale.SDH, locale.AZ_ARAB_IQ, locale.DV, locale.MNI_MTEI, locale.SHN, locale.EN_SHAW, locale.GN, locale.SW_KE, locale.CIC, locale.SMJ, locale.KPE_GN, locale.BAL_ARAB, locale.BSS, locale.MYV, locale.SCN, locale.MN, locale.MUS, locale.SAT_DEVA, locale.TIG, locale.QUC, locale.SKR, locale.HA_ARAB_SD, locale.QU_EC, locale.AZ_ARAB, locale.PIS, locale.SW, locale.CCH, locale.CAD, locale.CKB, locale.HNJ, locale.MDF, locale.MN_MONG, locale.UND, locale.AZ_ARAB_TR, locale.COP, locale.ZA, locale.GEZ, locale.BLT, locale.OSA, locale.II, locale.IU, locale.CU, locale.EN_DSRT, locale.WA, locale.BO_IN, locale.SW_UG, locale.AN, locale.RHG_ROHG, locale.AA_ER, locale.SMS, locale.TA_SG, locale.SMA_NO, locale.LAG, locale.SD_ARAB, locale.SYR_SY, locale.KAA, locale.PAP, locale.SMJ_NO, locale.LA, locale.BM_NKOO, locale.KAJ, locale.KOK, locale.VO, locale.LTG, locale.YI, locale.AA_DJ, locale.RHG, locale.APC, locale.TRV, locale.BO, locale.GAA, locale.SD_DEVA, locale.SYR, locale.BA, locale.AB, locale.BEW, locale.JBO, locale.MHN, locale.CSW, locale.LKT, locale.MIC, locale.QU, locale.SID, locale.SSY, locale.TO, locale.TYV, locale.WAL, locale.SMA, locale.ZH_LATN, locale.BYN, locale.WBP, locale.IO, locale.RHG_ROHG_BD, locale.SAT_OLCK, locale.MGO, locale.TA_MY, locale.MNI_BENG, locale.MI, locale.IU_LATN, locale.MNI, locale.SHN_TH, locale.CO, locale.HNJ_HMNP, locale.SD, locale.HA, locale.CHO, locale.MG, locale.BGC, locale.HA_GH, locale.NV, locale.BAL, locale.SDH_IQ, locale.HA_NE, locale.MZN, locale.RIF, locale.LRC, locale.PAP_AW, locale.AA, locale.HA_ARAB, locale.KAA_LATN:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 103 locales
	case locale.HAW, locale.ZH_HANS, locale.ML, locale.ES_MX, locale.TEO_KE, locale.CGG, locale.GV, locale.YUE_HANS, locale.TI, locale.ZH_HANS_MO, locale.DOI, locale.IG, locale.KO, locale.TI_ER, locale.VAI_VAII, locale.ES_US, locale.ES_SV, locale.YO_BJ, locale.SN, locale.VAI, locale.BM, locale.MN_MONG_MN, locale.CY, locale.ST_LS, locale.TEO, locale.ZH_HANS_SG, locale.SI, locale.MS, locale.SO_ET, locale.FIL, locale.SO_KE, locale.KN, locale.ZH_HANS_HK, locale.YUE, locale.PCM, locale.SO_DJ, locale.ZU, locale.GUZ, locale.EE_TG, locale.NUS, locale.ZH_HANT_MO, locale.ES_419, locale.KO_CN, locale.OR, locale.CEB, locale.OM, locale.EE, locale.ZH_HANT, locale.ROF, locale.KW, locale.MR, locale.ZH_HANT_HK, locale.ES_HN, locale.BEM, locale.KO_KP, locale.MT, locale.VUN, locale.ES_DO, locale.MS_SG, locale.DAV, locale.ES_NI, locale.KDE, locale.ZH_HANS_MY, locale.GA_GB, locale.MAS_TZ, locale.TH, locale.CHR, locale.NAQ, locale.YUE_HANT, locale.ES_BR, locale.SAQ, locale.VAI_LATN, locale.BHO, locale.JA, locale.KI, locale.JMC, locale.YO, locale.YUE_HANT_CN, locale.MER, locale.ND, locale.ES_BZ, locale.ZH_HANT_MY, locale.ES_PA, locale.AM, locale.ES_CU, locale.GD, locale.KAM, locale.MAS, locale.SO, locale.NYN, locale.ZH, locale.UG, locale.YUE_HANT_MO, locale.KLN, locale.MS_ARAB, locale.ES_PR, locale.GA, locale.OM_KE, locale.ST, locale.EN_MG, locale.EBU, locale.ES_GT, locale.AK:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 93 locales
	case locale.EN_BI, locale.EN_TV, locale.EN_VU, locale.EN_NU, locale.EN_RW, locale.EN_KE, locale.EN_ZM, locale.EN_UG, locale.EN_AG, locale.EN_BZ, locale.EN_GM, locale.EN_JM, locale.EN_PR, locale.EN_SL, locale.EN_BM, locale.EN_TK, locale.EN_CC, locale.EN_VC, locale.EN_CM, locale.EN_DM, locale.EN_TO, locale.EN_GY, locale.EN_DG, locale.EN_KY, locale.EN_SX, locale.EN_NR, locale.EN_GI, locale.EN_AS, locale.EN_FM, locale.EN_PG, locale.EN_PK, locale.EN_SS, locale.EN_UM, locale.EN_KN, locale.EN_KI, locale.EN_TC, locale.EN_SD, locale.EN_JE, locale.EN_LC, locale.EN_001, locale.EN_FK, locale.EN_GB, locale.EN_SH, locale.EN_IM, locale.EN_VI, locale.EN_MH, locale.EN_SZ, locale.EN_GG, locale.EN_GS, locale.EN_SC, locale.EN, locale.EN_VG, locale.EN_NZ, locale.EN_MO, locale.EN_CY, locale.EN_CK, locale.EN_IO, locale.EN_LR, locale.EN_GD, locale.EN_MS, locale.EN_AE, locale.EN_IE, locale.EN_PH, locale.EN_BB, locale.EN_NF, locale.EN_AU, locale.EN_FJ, locale.EN_WS, locale.EN_LS, locale.EN_MU, locale.EN_ER, locale.EN_PN, locale.EN_SG, locale.EN_NG, locale.EN_CX, locale.EN_PW, locale.EN_BW, locale.EN_AI, locale.EN_ZW, locale.EN_TT, locale.EN_GH, locale.EN_HK, locale.EN_IL, locale.EN_MT, locale.EN_CA, locale.EN_TZ, locale.EN_MP, locale.EN_GU, locale.EN_SB, locale.EN_BS, locale.EN_MY, locale.EN_NA, locale.EN_MW:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.RU_UA, locale.RU_BY, locale.RU_MD, locale.FF_LATN_GW, locale.HU, locale.CS, locale.PT_AO, locale.PT_MO, locale.FF_LATN_GM, locale.KY, locale.PT_ST, locale.SAH, locale.RU_KG, locale.FF_LATN_MR, locale.RU, locale.FF_LATN_NE, locale.FF_LATN_NG, locale.FF, locale.PT_CH, locale.KK_KZ, locale.EWO, locale.FF_LATN, locale.FF_LATN_GN, locale.HT, locale.KA, locale.BE, locale.UK, locale.PT_TL, locale.YAV, locale.LV, locale.BR, locale.PT_LU, locale.FF_LATN_BF, locale.PT_PT, locale.FF_LATN_SL, locale.KK_CYRL, locale.DUA, locale.FR_CA, locale.PT_MZ, locale.SMN, locale.PT_GW, locale.PL, locale.EO, locale.PT_CV, locale.SK, locale.CV, locale.PRG, locale.HY, locale.TZM, locale.UZ_CYRL, locale.BAS, locale.UZ, locale.PT_GQ, locale.NMG, locale.TK, locale.FF_LATN_LR, locale.KK, locale.DYO, locale.BE_TARASK, locale.FF_LATN_GH, locale.UZ_LATN, locale.TT, locale.SZL, locale.BG, locale.RU_KZ, locale.FF_LATN_CM, locale.KSF, locale.KEA, locale.TG:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 59 locales
	case locale.IT_SM, locale.SR_LATN_XK, locale.DE, locale.SR_LATN_BA, locale.CA_FR, locale.LN, locale.LN_CG, locale.CA_AD, locale.SR_LATN_ME, locale.BS_LATN, locale.SR_CYRL_ME, locale.DSB, locale.EL, locale.DE_IT, locale.DA, locale.NDS, locale.CA, locale.SC, locale.BS_CYRL, locale.GL, locale.CA_IT, locale.DA_GL, locale.SR_CYRL_BA, locale.AST, locale.ES, locale.BS, locale.CA_ES_VALENCIA, locale.ES_IC, locale.FR_MA, locale.AZ_LATN, locale.DE_LU, locale.VMW, locale.RO_MD, locale.MK, locale.IT_VA, locale.IS, locale.SR_CYRL_XK, locale.LB, locale.ES_PH, locale.LN_AO, locale.SR_LATN, locale.KU, locale.EL_CY, locale.IT, locale.AZ_CYRL, locale.VI, locale.RO, locale.SR, locale.ES_EA, locale.DE_BE, locale.LN_CF, locale.NDS_NL, locale.SR_CYRL, locale.EL_POLYTON, locale.HSB, locale.LIJ, locale.LLD, locale.AZ, locale.FR_LU:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_NC, locale.FR_BI, locale.FR_DJ, locale.FR_BJ, locale.FR_BE, locale.FR_RE, locale.FR_CH, locale.FR_BL, locale.FR_NE, locale.FR_ML, locale.FR_VU, locale.FR_CG, locale.FR_MG, locale.FR_PM, locale.FR_SN, locale.FR_MC, locale.FR_CM, locale.FR_CF, locale.FR_TD, locale.FR_YT, locale.FR, locale.FR_GF, locale.FR_TN, locale.FR_GQ, locale.FR_MF, locale.FR_CI, locale.FR_GP, locale.FR_RW, locale.FR_KM, locale.FR_SC, locale.FR_GN, locale.FR_DZ, locale.FR_TG, locale.FR_CD, locale.FR_SY, locale.FR_HT, locale.FR_MU, locale.FR_WF, locale.FR_MQ, locale.FR_MR, locale.FR_PF, locale.FR_BF, locale.FR_GA:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_YE, locale.AR_ER, locale.AR_BH, locale.AR_SS, locale.AR_TD, locale.AR_DJ, locale.AR_SA, locale.AR_EG, locale.AR_AE, locale.AR_KW, locale.AR_QA, locale.AR_SO, locale.AR_IQ, locale.AR, locale.AR_SD, locale.AR_OM, locale.AR_JO, locale.AR_SY, locale.AR_EH, locale.AR_IL, locale.AR_KM, locale.AR_PS:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.MGH, locale.IA, locale.YRL, locale.QU_BO, locale.KKJ, locale.YRL_VE, locale.MS_ARAB_BN, locale.RW, locale.MS_BN, locale.ES_CO, locale.YRL_CO, locale.JGO, locale.WO, locale.ES_AR, locale.PT, locale.FUR, locale.JV, locale.KGP, locale.SW_CD, locale.ES_UY, locale.NNH:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 17 locales
	case locale.KXV, locale.TE, locale.KXV_ORYA, locale.TA, locale.KXV_TELU, locale.KXV_DEVA, locale.BN_IN, locale.XNR, locale.DZ, locale.HI, locale.SA, locale.PA_GURU, locale.KXV_LATN, locale.PA, locale.GU, locale.TA_LK, locale.KOK_LATN:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_BF, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_NE, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_GW, locale.FF_ADLM_NG, locale.FF_ADLM_SL, locale.FF_ADLM_MR, locale.FF_ADLM_GM, locale.FF_ADLM_LR:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.EN_HU, locale.SQ_XK, locale.EN_PT, locale.EN_SK, locale.EN_CZ, locale.EN_SE, locale.SQ_MK, locale.EN_NO, locale.SQ, locale.EN_FI:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.LT, locale.SE_SE, locale.SV_FI, locale.ET, locale.SE, locale.SV_AX, locale.SV, locale.KSH, locale.SE_FI:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 9 locales
	case locale.TR_CY, locale.SU_LATN, locale.SU, locale.TR, locale.MS_ID, locale.ES_BO, locale.MUA, locale.ID, locale.ES_GQ:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.AGQ, locale.OC, locale.ZGH, locale.SHI_TFNG, locale.OC_ES, locale.KAB, locale.SHI_LATN, locale.SHI:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.EN_SI, locale.EN_DK, locale.EN_PL, locale.EN_RO, locale.EN_BE, locale.EN_ES, locale.EN_DE, locale.EN_IT:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_SX, locale.NL_AW, locale.NL, locale.NL_BE, locale.NL_BQ, locale.NL_SR, locale.NL_CW, locale.ES_PY:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.RWK, locale.LUO, locale.KM, locale.SBP, locale.BEZ, locale.LG, locale.KSB:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.VE, locale.SS, locale.SS_SZ, locale.ES_CR, locale.NR, locale.AF_NA, locale.AF:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_LB, locale.AR_TN, locale.AR_LY, locale.AR_MR, locale.AR_MA:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR, locale.SL, locale.EU, locale.FO, locale.HR_BA, locale.FO_DK:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.SG, locale.ES_CL, locale.ES_EC, locale.KL, locale.LO, locale.ES_VE:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN, locale.BGN_AF, locale.BGN_AE, locale.BGN_OM, locale.BGN_IR:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.MY, locale.XOG, locale.CE, locale.ASA, locale.TPI:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.TWQ, locale.KHQ, locale.SES:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.GSW_LI, locale.RM, locale.GSW:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NO, locale.NN, locale.NB_SJ, locale.NB:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.TS, locale.OS_RU, locale.DE_AT, locale.OS:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.AS, locale.NE:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.PS, locale.UZ_ARAB, locale.PS_PK:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 3 locales
	case locale.SEH, locale.LU, locale.RN:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_IN, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN, locale.TN_BW:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "MLGA",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "एम.जि.ए",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_NL:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߡߘߙ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "Ar",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "MGA",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
