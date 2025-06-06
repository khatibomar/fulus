// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = DKK{}

// Danish Krone currency type
type DKK struct{}

func (DKK) Code() string { return "DKK" }

func (DKK) Number() string { return "208" }

func (DKK) Name() string { return "Danish Krone" }

func (DKK) MinorUnits() int { return 2 }

func (DKK) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 140 locales
	case locale.LAG, locale.ANN, locale.IO, locale.LKT, locale.LA, locale.PAP_AW, locale.RHG_ROHG, locale.BM_NKOO, locale.QUC, locale.WBP, locale.AZ_ARAB, locale.SSY, locale.JBO, locale.MDF, locale.BA, locale.TRW, locale.KOK_DEVA, locale.WAL, locale.KAA_LATN, locale.SKR, locale.GN, locale.SDH, locale.GEZ_ER, locale.SCN, locale.MIC, locale.BLT, locale.LRC_IQ, locale.HA_ARAB, locale.HA_ARAB_SD, locale.KOK, locale.AA_DJ, locale.SD, locale.BEW, locale.SAT_OLCK, locale.KPE, locale.KK_ARAB, locale.SD_DEVA, locale.MUS, locale.II, locale.CAD, locale.KAJ, locale.LTG, locale.EN_DSRT, locale.FRR, locale.DV, locale.SMA_NO, locale.BSS, locale.HA_NE, locale.QU_EC, locale.SDH_IQ, locale.SHN_TH, locale.YI, locale.HNJ, locale.UND, locale.CKB_IR, locale.SAT_DEVA, locale.MHN, locale.SAT, locale.CSW, locale.MNI, locale.NV, locale.PAP, locale.GEZ, locale.MYV, locale.BGC, locale.COP, locale.KCG, locale.CCH, locale.RAJ, locale.WA, locale.AZ_ARAB_IQ, locale.KAA, locale.ZH_LATN, locale.SMA, locale.SW_UG, locale.BO_IN, locale.AA_ER, locale.AN, locale.SHN, locale.VO, locale.MN_MONG, locale.SW, locale.KPE_GN, locale.MGO, locale.AB, locale.GAA, locale.MNI_MTEI, locale.BYN, locale.NY, locale.CHO, locale.APC, locale.LRC, locale.MN, locale.MZN, locale.CIC, locale.RIF, locale.SMS, locale.SMJ_NO, locale.TO, locale.TIG, locale.TRV, locale.SD_ARAB, locale.MNI_BENG, locale.RHG_ROHG_BD, locale.CU, locale.CKB, locale.IU, locale.MG, locale.ARN, locale.BAL, locale.MAI, locale.CO, locale.SW_KE, locale.HA, locale.MOH, locale.ES_PE, locale.SYR_SY, locale.IU_LATN, locale.KEN, locale.EN_SHAW, locale.MI, locale.ZA, locale.TA_SG, locale.TYV, locale.HA_GH, locale.KS_DEVA, locale.PIS, locale.AA, locale.KAA_CYRL, locale.QU, locale.AZ_ARAB_TR, locale.BAL_ARAB, locale.TA_MY, locale.SMJ, locale.BO, locale.SYR, locale.HNJ_HMNP, locale.OSA, locale.RHG, locale.SID:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 101 locales
	case locale.ES_419, locale.ES_PA, locale.GA_GB, locale.ST, locale.YUE_HANT, locale.ROF, locale.DOI, locale.GD, locale.MT, locale.MS, locale.KDE, locale.ES_DO, locale.ZH_HANS_SG, locale.MN_MONG_MN, locale.HAW, locale.OM_KE, locale.SO, locale.VAI_LATN, locale.AK, locale.KN, locale.KW, locale.UG, locale.YUE, locale.YUE_HANT_MO, locale.CHR, locale.TEO, locale.KO_CN, locale.DAV, locale.YO_BJ, locale.ZH_HANS_HK, locale.GUZ, locale.SO_KE, locale.KO, locale.BM, locale.VUN, locale.FIL, locale.ZH_HANT_MO, locale.JA, locale.GV, locale.ZH_HANT_MY, locale.MAS, locale.SN, locale.MR, locale.ES_NI, locale.OR, locale.SAQ, locale.YUE_HANT_CN, locale.EE, locale.ES_BZ, locale.ES_US, locale.KAM, locale.YUE_HANS, locale.ND, locale.ES_CU, locale.YO, locale.MAS_TZ, locale.AM, locale.JMC, locale.MER, locale.ZH_HANS_MO, locale.EE_TG, locale.ES_PR, locale.ES_GT, locale.TH, locale.ES_BR, locale.PCM, locale.SO_DJ, locale.ES_SV, locale.ZH_HANT, locale.ZH, locale.SI, locale.VAI_VAII, locale.OM, locale.BEM, locale.ML, locale.ES_MX, locale.CGG, locale.IG, locale.ZH_HANS, locale.ES_HN, locale.KO_KP, locale.NAQ, locale.NUS, locale.MS_SG, locale.TI_ER, locale.BHO, locale.KLN, locale.SO_ET, locale.VAI, locale.ZH_HANS_MY, locale.EBU, locale.NYN, locale.CY, locale.ZH_HANT_HK, locale.TEO_KE, locale.TI, locale.CEB, locale.ST_LS, locale.GA, locale.KI, locale.MS_ARAB:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 94 locales
	case locale.EN_IO, locale.EN_CA, locale.EN_BZ, locale.EN_GD, locale.EN_KY, locale.EN_GI, locale.EN_VC, locale.EN_TO, locale.EN_001, locale.EN_DM, locale.EN_PH, locale.EN_CM, locale.EN_SX, locale.EN_AG, locale.EN_GH, locale.EN_IM, locale.EN_IE, locale.EN_TC, locale.EN_MP, locale.EN_IL, locale.EN_MH, locale.EN_GB, locale.EN_BM, locale.EN_LC, locale.EN_ZW, locale.EN_CY, locale.EN_SC, locale.EN_UM, locale.EN_GU, locale.EN_NF, locale.EN_VG, locale.EN_SL, locale.EN_AI, locale.EN_FM, locale.EN_MT, locale.EN_MU, locale.EN_TV, locale.EN_WS, locale.EN_SS, locale.EN_PW, locale.EN_NA, locale.EN_SH, locale.EN_CX, locale.EN_BW, locale.EN_MO, locale.EN_NR, locale.EN_UG, locale.EN_TK, locale.EN_MG, locale.EN_GY, locale.EN_HK, locale.EN_KE, locale.EN_BB, locale.EN_ER, locale.EN_SG, locale.EN_PG, locale.EN_GG, locale.EN_CC, locale.EN_LS, locale.EN_PN, locale.EN_SD, locale.EN_AS, locale.EN_KI, locale.EN_PR, locale.EN_VU, locale.EN_MS, locale.EN_FJ, locale.EN_SZ, locale.EN_JM, locale.EN_NG, locale.EN_TZ, locale.EN_GS, locale.EN_NU, locale.EN_RW, locale.EN_LR, locale.EN_PK, locale.EN_VI, locale.EN_TT, locale.EN_AU, locale.EN_CK, locale.EN_GM, locale.EN_BS, locale.EN_NZ, locale.EN_FK, locale.EN_MW, locale.EN_SB, locale.EN_BI, locale.EN_DG, locale.EN_ZM, locale.EN, locale.EN_AE, locale.EN_MY, locale.EN_KN, locale.EN_JE:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.PT_CV, locale.FF_LATN_NE, locale.BG, locale.RU, locale.FF_LATN_BF, locale.KEA, locale.TZM, locale.DYO, locale.PT_PT, locale.CV, locale.BAS, locale.KK_CYRL, locale.FF_LATN_GH, locale.FF, locale.PT_GW, locale.PRG, locale.FF_LATN_SL, locale.PT_GQ, locale.HY, locale.PT_AO, locale.KSF, locale.FF_LATN_GN, locale.UK, locale.FR_CA, locale.HU, locale.CS, locale.BE, locale.SMN, locale.UZ, locale.PT_MZ, locale.FF_LATN_NG, locale.RU_BY, locale.TT, locale.DUA, locale.BE_TARASK, locale.PT_MO, locale.BR, locale.EWO, locale.TK, locale.SK, locale.FF_LATN_GW, locale.YAV, locale.KK_KZ, locale.PL, locale.PT_CH, locale.PT_LU, locale.RU_UA, locale.UZ_CYRL, locale.HT, locale.UZ_LATN, locale.KY, locale.LV, locale.FF_LATN_MR, locale.FF_LATN_LR, locale.NMG, locale.RU_KG, locale.RU_MD, locale.KK, locale.RU_KZ, locale.SZL, locale.PT_TL, locale.FF_LATN, locale.FF_LATN_GM, locale.FF_LATN_CM, locale.KA, locale.PT_ST, locale.EO, locale.TG, locale.SAH:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 55 locales
	case locale.KU, locale.CA_FR, locale.LIJ, locale.DE_BE, locale.BS, locale.IT, locale.SR, locale.DE_IT, locale.CA_IT, locale.ES_EA, locale.AZ, locale.EL_CY, locale.ES, locale.NDS, locale.ES_PH, locale.BS_LATN, locale.CA_ES_VALENCIA, locale.CA, locale.LB, locale.CA_AD, locale.MK, locale.SC, locale.DSB, locale.VI, locale.FR_MA, locale.LN_AO, locale.SR_CYRL_XK, locale.BS_CYRL, locale.LN_CG, locale.EL, locale.LLD, locale.VMW, locale.DE_LU, locale.SR_CYRL_BA, locale.FR_LU, locale.SR_LATN_ME, locale.ES_IC, locale.GL, locale.LN_CF, locale.RO_MD, locale.EL_POLYTON, locale.SR_LATN, locale.NDS_NL, locale.SR_CYRL, locale.SR_CYRL_ME, locale.AZ_CYRL, locale.LN, locale.SR_LATN_BA, locale.SR_LATN_XK, locale.IT_VA, locale.RO, locale.HSB, locale.IT_SM, locale.AZ_LATN, locale.DE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR, locale.FR_VU, locale.FR_RE, locale.FR_TG, locale.FR_GN, locale.FR_TD, locale.FR_MG, locale.FR_PF, locale.FR_GF, locale.FR_CM, locale.FR_MU, locale.FR_MQ, locale.FR_SC, locale.FR_BF, locale.FR_GA, locale.FR_HT, locale.FR_BI, locale.FR_MR, locale.FR_NE, locale.FR_CI, locale.FR_ML, locale.FR_BE, locale.FR_KM, locale.FR_MF, locale.FR_CG, locale.FR_GQ, locale.FR_CD, locale.FR_MC, locale.FR_SN, locale.FR_DJ, locale.FR_BL, locale.FR_CF, locale.FR_PM, locale.FR_YT, locale.FR_WF, locale.FR_DZ, locale.FR_GP, locale.FR_SY, locale.FR_BJ, locale.FR_RW, locale.FR_TN, locale.FR_CH, locale.FR_NC:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_EH, locale.AR_EG, locale.AR_DJ, locale.AR_JO, locale.AR_SO, locale.AR, locale.AR_YE, locale.AR_KM, locale.AR_SD, locale.AR_QA, locale.AR_KW, locale.AR_SA, locale.AR_IL, locale.AR_OM, locale.AR_PS, locale.AR_TD, locale.AR_SS, locale.AR_ER, locale.AR_BH, locale.AR_SY, locale.AR_IQ, locale.AR_AE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.RW, locale.YRL_VE, locale.ES_AR, locale.PT, locale.ES_CO, locale.QU_BO, locale.IA, locale.MS_BN, locale.MGH, locale.WO, locale.KGP, locale.YRL_CO, locale.SW_CD, locale.JV, locale.JGO, locale.NNH, locale.MS_ARAB_BN, locale.KKJ, locale.YRL, locale.FUR, locale.ES_UY:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 17 locales
	case locale.BN_IN, locale.KOK_LATN, locale.HI, locale.DZ, locale.PA_GURU, locale.KXV_LATN, locale.KXV, locale.PA, locale.KXV_DEVA, locale.TA_LK, locale.KXV_ORYA, locale.GU, locale.SA, locale.TA, locale.TE, locale.XNR, locale.KXV_TELU:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_MR, locale.FF_ADLM_LR, locale.FF_ADLM_NG, locale.FF_ADLM_BF, locale.FF_ADLM_GW, locale.FF_ADLM_SL, locale.FF_ADLM_GM, locale.FF_ADLM, locale.FF_ADLM_NE, locale.FF_ADLM_CM:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.EN_NO, locale.EN_SE, locale.SQ, locale.EN_SK, locale.SQ_MK, locale.EN_FI, locale.EN_CZ, locale.EN_HU, locale.SQ_XK, locale.EN_PT:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ES_BO, locale.TR_CY, locale.SU_LATN, locale.MS_ID, locale.ID, locale.TR, locale.ES_GQ, locale.MUA, locale.SU:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.AST, locale.EN_DE, locale.EN_BE, locale.EN_SI, locale.EN_IT, locale.EN_PL, locale.EN_RO, locale.EN_ES:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.KAB, locale.OC, locale.SHI_TFNG, locale.OC_ES, locale.AGQ, locale.SHI_LATN, locale.ZGH, locale.SHI:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_BE, locale.NL_SX, locale.NL_BQ, locale.ES_PY, locale.NL_AW, locale.NL_SR, locale.NL, locale.NL_CW:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.BEZ, locale.LUO, locale.RWK, locale.SBP, locale.LG, locale.KSB, locale.KM:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.VE, locale.AF, locale.NR, locale.AF_NA, locale.SS, locale.SS_SZ, locale.ES_CR:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MA, locale.AR_LB, locale.AR_LY, locale.AR_MR, locale.AR_TN, locale.AR_DZ:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.SV_FI, locale.SE_SE, locale.SE, locale.SE_FI, locale.SV_AX, locale.SV:
		return FormatInfo{
			Symbol:           "Dkr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.BGN_IR, locale.BGN_OM, locale.BGN_AE, locale.BGN_AF, locale.BGN:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.FO, locale.HR, locale.SL, locale.HR_BA, locale.EU:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.SG, locale.ES_EC, locale.LO, locale.ES_VE, locale.ES_CL:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.XOG, locale.CE, locale.MY, locale.TPI, locale.ASA:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DE_AT, locale.OS_RU, locale.TS, locale.OS:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.RM, locale.GSW_FR, locale.GSW:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.IS, locale.DA_GL, locale.DA, locale.EN_DK:
		return FormatInfo{
			Symbol:           "kr.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NO, locale.NB:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.TWQ, locale.KHQ, locale.SES, locale.DJE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.BN, locale.CCP_IN:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LT, locale.KSH, locale.ET:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.LU, locale.SEH, locale.RN:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS_PK, locale.PS:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.HI_LATN, locale.EN_IN:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "DNMK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "दि.के.के",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_NL:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FO_DK:
		return FormatInfo{
			Symbol:           "kr.",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.KL:
		return FormatInfo{
			Symbol:           "kr.",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߘߞߞ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ZU:
		return FormatInfo{
			Symbol:           "Kr",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "DKK",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
