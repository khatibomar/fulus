// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = EUR{}

// Euro currency type
type EUR struct{}

func (EUR) Code() string { return "EUR" }

func (EUR) Number() string { return "978" }

func (EUR) Name() string { return "Euro" }

func (EUR) MinorUnits() int { return 2 }

func (EUR) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 182 locales
	case locale.ZH_HANS_SG, locale.EN_SS, locale.UG, locale.EN_CY, locale.CY, locale.DAV, locale.EN_IL, locale.EN_SD, locale.ZH_HANT_MY, locale.YUE_HANT, locale.EN_PH, locale.EN_PW, locale.EN_VI, locale.ZH, locale.EN_MU, locale.NAQ, locale.ZH_HANS_MO, locale.EN_SX, locale.EN_AE, locale.MN_MONG_MN, locale.EN_MO, locale.EN_SG, locale.EN_ZW, locale.YUE_HANS, locale.EN_MW, locale.YO, locale.TEO_KE, locale.GA_GB, locale.BM, locale.EN_SB, locale.MAS_TZ, locale.EN_GB, locale.EN_NG, locale.EN_RW, locale.KN, locale.EN_DG, locale.EN_HK, locale.EN_BM, locale.EN_FK, locale.EN_MP, locale.EN_TV, locale.EN_CC, locale.EN_MS, locale.GUZ, locale.SO_KE, locale.ZH_HANS_HK, locale.EN_PG, locale.EN_BI, locale.ZH_HANT_MO, locale.EN_MT, locale.KDE, locale.EN_NR, locale.CHR, locale.VAI_VAII, locale.EN_BW, locale.EN_UM, locale.MAS, locale.EN_PK, locale.EN_DM, locale.MER, locale.EN_CM, locale.EN_LC, locale.EBU, locale.ML, locale.EN_WS, locale.KO_CN, locale.EN_KE, locale.MR, locale.ST_LS, locale.EN_MG, locale.EN_ZM, locale.EE_TG, locale.MS, locale.SO, locale.NUS, locale.EN_CX, locale.EN_SZ, locale.KO, locale.OM, locale.EN_GY, locale.EN_NF, locale.EN_ER, locale.YO_BJ, locale.OR, locale.SN, locale.SO_DJ, locale.EN_MY, locale.BHO, locale.TEO, locale.EN_AG, locale.EN_PR, locale.EN_UG, locale.TI, locale.EN_GG, locale.EN_GS, locale.EN_IO, locale.EN_NZ, locale.MS_SG, locale.EN_GD, locale.EN_SL, locale.BEM, locale.DOI, locale.CGG, locale.EN_GU, locale.HAW, locale.MT, locale.ST, locale.EN_001, locale.EN_TZ, locale.KO_KP, locale.EN_TO, locale.VAI, locale.EN_KN, locale.OM_KE, locale.EN_JE, locale.EN_AI, locale.ZH_HANT_HK, locale.EN_KY, locale.JA, locale.KAM, locale.MS_ARAB, locale.CEB, locale.EN_VG, locale.SO_ET, locale.YUE_HANT_MO, locale.KI, locale.ZH_HANS, locale.EN_JM, locale.EN_TC, locale.ZH_HANS_MY, locale.EN_IM, locale.EN_VU, locale.EN_GH, locale.EN_AS, locale.EN_TK, locale.ZH_HANT, locale.GD, locale.PCM, locale.NYN, locale.EE, locale.EN, locale.FIL, locale.SAQ, locale.AK, locale.TH, locale.YUE, locale.EN_BB, locale.EN_LR, locale.EN_SH, locale.EN_KI, locale.EN_CA, locale.EN_TT, locale.KLN, locale.VAI_LATN, locale.EN_LS, locale.KW, locale.AM, locale.EN_GI, locale.EN_GM, locale.EN_BZ, locale.EN_FJ, locale.ZU, locale.EN_BS, locale.EN_CK, locale.TI_ER, locale.EN_FM, locale.EN_IE, locale.YUE_HANT_CN, locale.ND, locale.EN_NA, locale.JMC, locale.EN_SC, locale.EN_MH, locale.EN_PN, locale.VUN, locale.IG, locale.GA, locale.GV, locale.ROF, locale.EN_NU, locale.SI, locale.EN_VC:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 141 locales
	case locale.TIG, locale.TO, locale.DV, locale.KAJ, locale.SD_DEVA, locale.COP, locale.AA, locale.AB, locale.SID, locale.SW_UG, locale.MYV, locale.VO, locale.CCH, locale.LAG, locale.SYR_SY, locale.KAA_CYRL, locale.SD, locale.QU_EC, locale.TA_SG, locale.KOK_DEVA, locale.LTG, locale.MI, locale.CKB, locale.HA_ARAB_SD, locale.PIS, locale.BYN, locale.CO, locale.BAL_ARAB, locale.TRW, locale.IO, locale.LA, locale.OSA, locale.SDH, locale.MN, locale.BA, locale.SAT_DEVA, locale.SW, locale.AA_DJ, locale.BAL, locale.CAD, locale.EN_DSRT, locale.MHN, locale.SAT_OLCK, locale.CHO, locale.ANN, locale.SMS, locale.KEN, locale.RAJ, locale.HNJ_HMNP, locale.BAL_LATN, locale.MUS, locale.AZ_ARAB_IQ, locale.KAA, locale.II, locale.PAP_AW, locale.MNI_BENG, locale.SW_KE, locale.AN, locale.FRR, locale.AZ_ARAB_TR, locale.GEZ, locale.KS_DEVA, locale.SHN_TH, locale.BM_NKOO, locale.MIC, locale.RHG_ROHG_BD, locale.BEW, locale.KAA_LATN, locale.BLT, locale.GN, locale.WBP, locale.GEZ_ER, locale.LRC, locale.MN_MONG, locale.MGO, locale.HA_ARAB, locale.EN_SHAW, locale.QUC, locale.BO_IN, locale.GAA, locale.HNJ, locale.KK_ARAB, locale.KPE, locale.RHG_ROHG, locale.YI, locale.AZ_ARAB, locale.MAI, locale.SCN, locale.SDH_IQ, locale.CSW, locale.SMA_NO, locale.JBO, locale.MZN, locale.SYR, locale.IU_LATN, locale.MNI, locale.HA, locale.KOK, locale.BSS, locale.APC, locale.KCG, locale.SMJ_NO, locale.WA, locale.WAL, locale.SAT, locale.ZA, locale.SSY, locale.MDF, locale.RHG, locale.LRC_IQ, locale.NY, locale.CIC, locale.MG, locale.CKB_IR, locale.MNI_MTEI, locale.TA_MY, locale.AA_ER, locale.TRV, locale.BGC, locale.MOH, locale.SD_ARAB, locale.ZH_LATN, locale.CU, locale.LKT, locale.SKR, locale.EN_MV, locale.RIF, locale.SMJ, locale.ARN, locale.SHN, locale.TYV, locale.NV, locale.KPE_GN, locale.PAP, locale.QU, locale.IU, locale.UND, locale.HA_NE, locale.SMA, locale.BO, locale.HA_GH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 76 locales
	case locale.FF_LATN_MR, locale.KK, locale.CV, locale.FF_LATN_CM, locale.FF_LATN_SL, locale.RU_KG, locale.KK_CYRL, locale.PT_GQ, locale.FF_LATN_GH, locale.CS, locale.KSF, locale.PT_ST, locale.EN_HU, locale.RU_BY, locale.SZL, locale.SAH, locale.SK, locale.FF_LATN_GW, locale.FF_LATN, locale.PT_PT, locale.SMN, locale.FF_LATN_NG, locale.PT_MZ, locale.UZ_LATN, locale.HT, locale.EN_NO, locale.NMG, locale.KY, locale.TZM, locale.FF, locale.PT_GW, locale.FF_LATN_GN, locale.KA, locale.PRG, locale.BG, locale.FF_LATN_GM, locale.SQ_MK, locale.KEA, locale.EN_SE, locale.PT_AO, locale.FF_LATN_NE, locale.SQ, locale.RU_MD, locale.EWO, locale.FF_LATN_LR, locale.TT, locale.TG, locale.YAV, locale.BR, locale.FR_CA, locale.EN_CZ, locale.LV, locale.DYO, locale.RU_UA, locale.RU_KZ, locale.EN_SK, locale.UZ, locale.BAS, locale.PT_CV, locale.PT_MO, locale.FF_LATN_BF, locale.RU, locale.EO, locale.SQ_XK, locale.HY, locale.EN_PT, locale.BE_TARASK, locale.PL, locale.BE, locale.UZ_CYRL, locale.EN_FI, locale.PT_CH, locale.PT_TL, locale.PT_LU, locale.DUA, locale.KK_KZ:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 64 locales
	case locale.VMW, locale.EN_IT, locale.EL_CY, locale.SR_LATN, locale.AZ_LATN, locale.ES_PH, locale.LB, locale.SR, locale.FR_MA, locale.AST, locale.EN_DK, locale.IT, locale.VI, locale.KU, locale.DE_IT, locale.ES_EA, locale.DE, locale.LIJ, locale.MK, locale.SR_LATN_ME, locale.CA, locale.IT_SM, locale.LN, locale.AZ, locale.FR_LU, locale.HSB, locale.BS_CYRL, locale.EN_RO, locale.LN_CG, locale.SR_CYRL_XK, locale.LLD, locale.EN_ES, locale.NDS, locale.ES, locale.EN_PL, locale.AZ_CYRL, locale.EL, locale.CA_AD, locale.EN_SI, locale.BS, locale.CA_IT, locale.DA_GL, locale.EN_DE, locale.EL_POLYTON, locale.DE_LU, locale.DE_BE, locale.BS_LATN, locale.DSB, locale.EN_BE, locale.SR_CYRL, locale.IT_VA, locale.SR_CYRL_BA, locale.LN_AO, locale.DA, locale.CA_ES_VALENCIA, locale.CA_FR, locale.ES_IC, locale.NDS_NL, locale.SR_LATN_BA, locale.SR_CYRL_ME, locale.SR_LATN_XK, locale.LN_CF, locale.GL, locale.SC:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 44 locales
	case locale.FR_SY, locale.FR_DZ, locale.FR_VU, locale.FR_BJ, locale.FR_MQ, locale.FR_GP, locale.FR_RE, locale.FR_BF, locale.FR_CF, locale.FR_TG, locale.FR_BL, locale.FR_WF, locale.FR_DJ, locale.FR_MG, locale.FR_MU, locale.FR_PF, locale.FR_HT, locale.FR_GA, locale.FR, locale.FR_CG, locale.FR_CH, locale.FR_ML, locale.FR_CD, locale.FR_MR, locale.FR_MC, locale.FR_CM, locale.FR_TD, locale.FR_YT, locale.FR_CI, locale.FR_BE, locale.FR_NE, locale.EN_FR, locale.FR_GQ, locale.FR_GN, locale.FR_GF, locale.FR_RW, locale.FR_SC, locale.FR_TN, locale.FR_MF, locale.FR_KM, locale.FR_BI, locale.FR_NC, locale.FR_SN, locale.FR_PM:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_EG, locale.AR_SA, locale.AR_SS, locale.AR_SD, locale.AR_AE, locale.AR_KW, locale.AR_SY, locale.AR_PS, locale.AR_SO, locale.AR_TD, locale.AR_JO, locale.AR_OM, locale.AR, locale.AR_ER, locale.AR_YE, locale.AR_IL, locale.AR_KM, locale.AR_IQ, locale.AR_DJ, locale.AR_EH, locale.AR_BH, locale.AR_QA:
		return FormatInfo{
			Symbol:           "€",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.EN_AT, locale.SW_CD, locale.YRL, locale.YRL_VE, locale.IA, locale.WO, locale.JV, locale.FUR, locale.MS_BN, locale.KKJ, locale.MGH, locale.PT, locale.JGO, locale.RW, locale.MS_ARAB_BN, locale.YRL_CO, locale.KGP, locale.QU_BO, locale.NNH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.HI_LATN, locale.HI, locale.KXV_TELU, locale.GU, locale.XNR, locale.BN_IN, locale.KXV_LATN, locale.KXV, locale.DZ, locale.EN_IN, locale.SA, locale.TE, locale.KXV_DEVA, locale.TA_LK, locale.PA_GURU, locale.TA, locale.KXV_ORYA, locale.PA, locale.KOK_LATN:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.EN_AU, locale.ES_GT, locale.ES_NI, locale.ES_DO, locale.ES_BZ, locale.ES_CU, locale.ES_419, locale.ES_SV, locale.ES_PR, locale.ES_PA, locale.ES_US, locale.ES_HN, locale.ES_BR, locale.ES_MX:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_GW, locale.FF_ADLM_GM, locale.FF_ADLM_BF, locale.FF_ADLM_NG, locale.FF_ADLM_SN, locale.FF_ADLM_SL, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_GH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.SV_AX, locale.FI, locale.SE, locale.SV_FI, locale.SE_SE, locale.ET, locale.LT, locale.SV, locale.SE_FI, locale.KSH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 9 locales
	case locale.TR_CY, locale.MUA, locale.EN_ID, locale.ID, locale.MS_ID, locale.ES_GQ, locale.SU_LATN, locale.TR, locale.SU:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.KAB, locale.OC_ES, locale.SHI, locale.SHI_LATN, locale.SHI_TFNG, locale.OC, locale.AGQ, locale.ZGH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_CW, locale.NL_AW, locale.NL_SX, locale.NL_BQ, locale.NL_SR, locale.EN_NL, locale.NL_BE, locale.NL:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSB, locale.KM, locale.LG, locale.LUO, locale.SBP, locale.BEZ, locale.RWK:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.NR, locale.EN_ZA, locale.AF, locale.VE, locale.SS_SZ, locale.SS, locale.AF_NA:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_TN, locale.AR_MA, locale.AR_LY, locale.AR_MR, locale.AR_LB:
		return FormatInfo{
			Symbol:           "€",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.SL, locale.HR, locale.FO, locale.FO_DK, locale.HR_BA, locale.EU:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.XOG, locale.MY, locale.TPI, locale.ASA, locale.EN_150, locale.CE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_IR, locale.BGN_OM, locale.BGN, locale.BGN_AF, locale.BGN_AE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.RM, locale.GSW_FR, locale.GSW:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NN, locale.NO, locale.NB:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NE_IN, locale.NE, locale.AS, locale.BRX:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.OS_RU, locale.DE_AT, locale.OS, locale.TS:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.TWQ, locale.DJE, locale.SES, locale.KHQ:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.CCP, locale.BN:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_AR, locale.ES_CO, locale.ES_UY:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_EC, locale.ES_CL, locale.ES_VE:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.HU, locale.UK, locale.TK:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LO, locale.KL, locale.SG:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.RN, locale.SEH, locale.LU:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.RO, locale.IS, locale.RO_MD:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS_PK, locale.PS:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN, locale.TN_BW:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.DE_CH:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_BO:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_CR:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_PE:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_PY:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "€",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "€",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "EUR",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "€",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "€",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
