// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = GTQ{}

// Quetzal currency type
type GTQ struct{}

func (GTQ) Code() string { return "GTQ" }

func (GTQ) Number() string { return "320" }

func (GTQ) Name() string { return "Quetzal" }

func (GTQ) MinorUnits() int { return 2 }

func (GTQ) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 194 locales
	case locale.AM, locale.EN_PH, locale.KI, locale.EN_CA, locale.EN_IO, locale.KW, locale.EN_MY, locale.ES_NI, locale.EN_AG, locale.EN_CX, locale.EN_MT, locale.EN_LS, locale.EN_MU, locale.ES_419, locale.YUE, locale.EE, locale.EN_SG, locale.EN_TC, locale.EN_VC, locale.EN_WS, locale.EN_GH, locale.EN_GY, locale.YO_BJ, locale.EN_BM, locale.EN_BI, locale.GV, locale.EN_SC, locale.ES_US, locale.DOI, locale.EN_SL, locale.EN_JM, locale.ES_BR, locale.ES_CU, locale.MN_MONG_MN, locale.ST_LS, locale.ZH_HANS_HK, locale.EN_AU, locale.EN_VU, locale.ES_PA, locale.EN_MH, locale.EN_MO, locale.JA, locale.EN_MG, locale.TI, locale.EN_CM, locale.ROF, locale.ES_BZ, locale.ES_MX, locale.MS, locale.EN_KE, locale.EN_NF, locale.ST, locale.MS_ARAB, locale.EN_LC, locale.EN_PK, locale.EN_JE, locale.ZH_HANS, locale.HAW, locale.EN_BB, locale.EN_BZ, locale.EN_ZM, locale.ZH_HANT_HK, locale.ES_HN, locale.IG, locale.VAI, locale.EN_KN, locale.EN_SB, locale.ES_PR, locale.GD, locale.UG, locale.ZH_HANS_MY, locale.EN_RW, locale.MS_SG, locale.EN_GM, locale.EN_IE, locale.EN_MW, locale.EN_SD, locale.ES_DO, locale.TH, locale.YO, locale.EN_CY, locale.EN_TV, locale.KAM, locale.ZH, locale.EN_NZ, locale.MR, locale.EN_HK, locale.EN_PN, locale.EN_SX, locale.GA_GB, locale.EN_GI, locale.EE_TG, locale.VUN, locale.EN_BS, locale.EN_NR, locale.KO_KP, locale.ZH_HANS_SG, locale.EN_GG, locale.EN_SZ, locale.GA, locale.SAQ, locale.ZH_HANT_MO, locale.EN_KI, locale.SI, locale.EN_VI, locale.SO_ET, locale.EN_DM, locale.EN_MS, locale.EN_PW, locale.EN_SS, locale.KO, locale.ZH_HANT_MY, locale.NUS, locale.NYN, locale.TEO, locale.EN_KY, locale.EN_UG, locale.ES_SV, locale.EN_AE, locale.EN_NG, locale.CHR, locale.EN_PG, locale.ML, locale.OR, locale.OM, locale.EN_DG, locale.EN_ER, locale.EN_SH, locale.MT, locale.EBU, locale.KLN, locale.JMC, locale.ES_GT, locale.TI_ER, locale.YUE_HANT_CN, locale.OM_KE, locale.MAS, locale.SO_KE, locale.YUE_HANT, locale.KDE, locale.EN_AS, locale.GUZ, locale.CEB, locale.EN_FJ, locale.EN_LR, locale.FIL, locale.MER, locale.EN_CK, locale.EN_GB, locale.DAV, locale.EN_GD, locale.EN_AI, locale.EN_UM, locale.EN_CC, locale.EN_GU, locale.EN_ZW, locale.ZH_HANT, locale.EN_BW, locale.EN_FK, locale.EN_TT, locale.MAS_TZ, locale.CGG, locale.EN_TK, locale.ZH_HANS_MO, locale.CY, locale.EN_PR, locale.TEO_KE, locale.EN_MP, locale.NAQ, locale.AK, locale.EN, locale.ND, locale.PCM, locale.VAI_VAII, locale.EN_001, locale.YUE_HANS, locale.EN_NA, locale.EN_TO, locale.BM, locale.EN_IM, locale.SO, locale.EN_NU, locale.EN_FM, locale.KN, locale.ZU, locale.BEM, locale.BHO, locale.VAI_LATN, locale.SO_DJ, locale.EN_TZ, locale.EN_IL, locale.SN, locale.EN_VG, locale.KO_CN:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.FRR, locale.NY, locale.KAA_CYRL, locale.QUC, locale.DV, locale.MAI, locale.SMS, locale.AA, locale.GEZ_ER, locale.ZH_LATN, locale.TRW, locale.LRC_IQ, locale.MNI, locale.HA_GH, locale.PAP, locale.LRC, locale.TA_MY, locale.BO_IN, locale.SD_ARAB, locale.HA_ARAB_SD, locale.SW_UG, locale.MHN, locale.HNJ, locale.MI, locale.TO, locale.AA_ER, locale.LKT, locale.CKB, locale.SD, locale.QU, locale.RHG_ROHG, locale.AN, locale.SW_KE, locale.MYV, locale.IO, locale.BGC, locale.AZ_ARAB, locale.KPE, locale.KS_DEVA, locale.MOH, locale.IU, locale.RHG_ROHG_BD, locale.CU, locale.CSW, locale.RAJ, locale.APC, locale.CCH, locale.SHN, locale.SW, locale.AB, locale.WAL, locale.PAP_AW, locale.KK_ARAB, locale.SHN_TH, locale.SKR, locale.BAL_ARAB, locale.YI, locale.KOK, locale.KEN, locale.ARN, locale.MIC, locale.BEW, locale.SID, locale.SYR_SY, locale.GEZ, locale.LTG, locale.SAT, locale.HA, locale.PIS, locale.BAL, locale.GAA, locale.CHO, locale.WA, locale.JBO, locale.SMJ_NO, locale.RHG, locale.OSA, locale.VO, locale.ZA, locale.MN_MONG, locale.KCG, locale.MG, locale.SAT_DEVA, locale.GN, locale.SSY, locale.KAA, locale.AZ_ARAB_TR, locale.WBP, locale.KOK_DEVA, locale.AZ_ARAB_IQ, locale.SD_DEVA, locale.UND, locale.ES_PE, locale.EN_SHAW, locale.MUS, locale.CAD, locale.EN_DSRT, locale.TRV, locale.SAT_OLCK, locale.AA_DJ, locale.BA, locale.SCN, locale.CO, locale.MNI_BENG, locale.QU_EC, locale.HA_ARAB, locale.LA, locale.SYR, locale.SDH, locale.MNI_MTEI, locale.BYN, locale.ANN, locale.BSS, locale.IU_LATN, locale.KAJ, locale.RIF, locale.MGO, locale.CIC, locale.LAG, locale.EN_MV, locale.II, locale.BO, locale.BLT, locale.BM_NKOO, locale.CKB_IR, locale.SMA_NO, locale.MZN, locale.SMJ, locale.TIG, locale.NV, locale.HA_NE, locale.HNJ_HMNP, locale.KAA_LATN, locale.MDF, locale.SDH_IQ, locale.KPE_GN, locale.TA_SG, locale.SMA, locale.MN, locale.TYV:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 70 locales
	case locale.TT, locale.CS, locale.PT_MO, locale.PT_TL, locale.KK_CYRL, locale.CV, locale.UZ, locale.UZ_LATN, locale.FF_LATN_LR, locale.TG, locale.PT_CH, locale.FF_LATN_CM, locale.DUA, locale.SK, locale.EO, locale.TK, locale.RU, locale.FF_LATN_GH, locale.KEA, locale.FF, locale.BG, locale.PT_GW, locale.KK_KZ, locale.PT_CV, locale.PT_AO, locale.FF_LATN_NG, locale.KA, locale.BAS, locale.SAH, locale.HY, locale.RU_KG, locale.HU, locale.EWO, locale.KY, locale.FF_LATN_NE, locale.RU_KZ, locale.FF_LATN_GM, locale.RU_MD, locale.FF_LATN_SL, locale.BR, locale.FR_CA, locale.LV, locale.PT_LU, locale.FF_LATN_MR, locale.EN_SE, locale.FF_LATN_GW, locale.FF_LATN, locale.TZM, locale.KK, locale.UZ_CYRL, locale.BE, locale.PT_ST, locale.PL, locale.PT_MZ, locale.FF_LATN_BF, locale.PRG, locale.KSF, locale.SZL, locale.FF_LATN_GN, locale.YAV, locale.UK, locale.EN_FI, locale.BE_TARASK, locale.PT_GQ, locale.PT_PT, locale.RU_UA, locale.NMG, locale.DYO, locale.SMN, locale.RU_BY:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 63 locales
	case locale.HSB, locale.NDS, locale.LB, locale.CA_ES_VALENCIA, locale.AST, locale.LN, locale.ES, locale.FR_LU, locale.CA_IT, locale.SR_LATN, locale.SC, locale.SR, locale.EL, locale.IS, locale.BS, locale.DE, locale.EL_CY, locale.IT_VA, locale.SR_LATN_XK, locale.LN_AO, locale.LLD, locale.EN_BE, locale.CA_FR, locale.VI, locale.AZ_CYRL, locale.SR_CYRL_ME, locale.DSB, locale.AZ_LATN, locale.DE_BE, locale.ES_PH, locale.NDS_NL, locale.VMW, locale.BS_CYRL, locale.DA, locale.LIJ, locale.ES_EA, locale.CA_AD, locale.SR_CYRL_XK, locale.GL, locale.MK, locale.IT, locale.DE_LU, locale.SR_CYRL_BA, locale.RO_MD, locale.DE_IT, locale.AZ, locale.EN_DE, locale.LN_CF, locale.LN_CG, locale.EL_POLYTON, locale.EN_SI, locale.CA, locale.SR_LATN_ME, locale.ES_IC, locale.RO, locale.EN_DK, locale.BS_LATN, locale.DA_GL, locale.IT_SM, locale.KU, locale.SR_LATN_BA, locale.FR_MA, locale.SR_CYRL:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_NC, locale.FR_GF, locale.FR_GP, locale.FR_MR, locale.FR_CI, locale.FR_NE, locale.FR_BE, locale.FR_DJ, locale.FR_MF, locale.FR_VU, locale.FR_CF, locale.FR_TG, locale.FR_TN, locale.FR_SY, locale.FR_TD, locale.FR_CM, locale.FR_HT, locale.FR_MC, locale.FR_BJ, locale.FR_RE, locale.FR_CD, locale.FR_DZ, locale.FR_PM, locale.FR_MQ, locale.FR_BF, locale.FR_KM, locale.FR_CG, locale.FR_CH, locale.FR_SC, locale.FR_SN, locale.FR_GQ, locale.FR_GN, locale.FR_MU, locale.FR, locale.FR_WF, locale.FR_PF, locale.FR_RW, locale.FR_ML, locale.FR_BI, locale.FR_BL, locale.FR_MG, locale.FR_GA, locale.FR_YT:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SD, locale.AR_EG, locale.AR_TD, locale.AR_KW, locale.AR_DJ, locale.AR_IQ, locale.AR_KM, locale.AR_PS, locale.AR_YE, locale.AR_AE, locale.AR_EH, locale.AR_SO, locale.AR_SA, locale.AR_ER, locale.AR_JO, locale.AR, locale.AR_QA, locale.AR_IL, locale.AR_SY, locale.AR_OM, locale.AR_SS, locale.AR_BH:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.YRL_CO, locale.ES_UY, locale.ES_CO, locale.MGH, locale.YRL, locale.FUR, locale.RW, locale.MS_BN, locale.SW_CD, locale.JV, locale.KKJ, locale.KGP, locale.YRL_VE, locale.JGO, locale.ES_AR, locale.QU_BO, locale.MS_ARAB_BN, locale.IA, locale.PT, locale.NNH, locale.EN_AT:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.TE, locale.HI_LATN, locale.PA_GURU, locale.DZ, locale.HI, locale.KXV_LATN, locale.EN_IN, locale.TA, locale.KXV_ORYA, locale.KOK_LATN, locale.GU, locale.SA, locale.PA, locale.KXV_TELU, locale.BN_IN, locale.XNR, locale.KXV_DEVA, locale.TA_LK, locale.KXV:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GH, locale.FF_ADLM_NE, locale.FF_ADLM_LR, locale.FF_ADLM_NG, locale.FF_ADLM_GM, locale.FF_ADLM_CM, locale.FF_ADLM_BF, locale.FF_ADLM_GW, locale.FF_ADLM, locale.FF_ADLM_SL, locale.FF_ADLM_MR, locale.FF_ADLM_SN:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.ID, locale.MS_ID, locale.SU_LATN, locale.SU, locale.TR_CY, locale.EN_ID, locale.TR, locale.MUA, locale.ES_GQ, locale.ES_BO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ES_PY, locale.NL_BQ, locale.EN_NL, locale.NL_AW, locale.NL_SX, locale.NL_BE, locale.NL_SR, locale.NL_CW, locale.NL:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ET, locale.SE_FI, locale.SV_FI, locale.LT, locale.KSH, locale.SE, locale.SE_SE, locale.SV, locale.SV_AX:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 8 locales
	case locale.SHI, locale.KAB, locale.ZGH, locale.OC, locale.OC_ES, locale.AGQ, locale.SHI_LATN, locale.SHI_TFNG:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SS_SZ, locale.VE, locale.AF_NA, locale.AF, locale.NR, locale.EN_ZA, locale.ES_CR, locale.SS:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.LG, locale.LUO, locale.KM, locale.RWK, locale.BEZ, locale.SBP, locale.KSB:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_DZ, locale.AR_MR, locale.AR_TN, locale.AR_MA, locale.AR_LY, locale.AR_LB:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ASA, locale.XOG, locale.MY, locale.EN_150, locale.TPI, locale.CE:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.ES_VE, locale.ES_EC, locale.LO, locale.SG, locale.KL, locale.ES_CL:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.SL, locale.HR_BA, locale.HR, locale.EU, locale.FO_DK, locale.FO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.BGN_OM, locale.BGN, locale.BGN_IR, locale.BGN_AE, locale.BGN_AF:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.KHQ, locale.SES, locale.DJE, locale.TWQ:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NN, locale.NB_SJ, locale.NB, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS_RU, locale.OS, locale.TS, locale.DE_AT:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW, locale.GSW_LI, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.EN_CH, locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE, locale.NE_IN, locale.AS:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.RN, locale.LU:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SQ, locale.SQ_MK, locale.SQ_XK:
		return CurrencyFormatInfo{
			Symbol:           "GTQ",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "GTMK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "जि.ति.किउ",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "GTQ",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߜ߭ߕߞ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.WO:
		return CurrencyFormatInfo{
			Symbol:           "GT Q",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "Q",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
