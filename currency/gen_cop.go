// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = COP{}

// Colombian Peso currency type
type COP struct{}

func (COP) Code() string { return "COP" }

func (COP) Number() string { return "170" }

func (COP) Name() string { return "Colombian Peso" }

func (COP) MinorUnits() int { return 2 }

func (COP) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 194 locales
	case locale.YUE_HANT, locale.MAS, locale.EN_RW, locale.ES_CU, locale.KDE, locale.EN_KN, locale.MS, locale.EN_ER, locale.EN_PK, locale.ROF, locale.KO, locale.EN_JM, locale.EN_GY, locale.TI, locale.ZU, locale.EN_PG, locale.ES_MX, locale.ES_419, locale.EN_GD, locale.EN_MW, locale.YO_BJ, locale.YUE_HANS, locale.KO_CN, locale.EN_UG, locale.EN_ZM, locale.TH, locale.EN_FJ, locale.GUZ, locale.EN_FK, locale.EN_KY, locale.EN_IO, locale.EN_LC, locale.EN_SH, locale.EN_SZ, locale.EN_BM, locale.EN_IM, locale.MAS_TZ, locale.ZH_HANS_MO, locale.EN_DM, locale.EN_001, locale.EN_SS, locale.OM_KE, locale.JMC, locale.PCM, locale.ZH_HANS_SG, locale.ES_SV, locale.ES_PA, locale.ES_US, locale.VAI_LATN, locale.EN_GI, locale.TEO, locale.MER, locale.KLN, locale.KN, locale.EN, locale.FIL, locale.EN_TC, locale.EN_SC, locale.EN_AG, locale.EN_AS, locale.EN_UM, locale.ES_GT, locale.KW, locale.EN_CK, locale.EN_GM, locale.EN_MS, locale.CHR, locale.CY, locale.EN_NR, locale.EN_ZW, locale.EN_CC, locale.BHO, locale.EN_MY, locale.ZH, locale.ES_HN, locale.EN_BB, locale.EN_DG, locale.EN_LS, locale.SAQ, locale.MS_SG, locale.EN_SG, locale.EN_SL, locale.TI_ER, locale.EN_AI, locale.EN_MU, locale.DAV, locale.EN_MH, locale.EN_TK, locale.OM, locale.EN_PN, locale.SO_DJ, locale.DOI, locale.MR, locale.NYN, locale.JA, locale.BEM, locale.EN_VI, locale.EN_GH, locale.EN_PR, locale.EN_TZ, locale.VUN, locale.KI, locale.EN_JE, locale.GA, locale.NUS, locale.EN_IL, locale.SN, locale.EN_PH, locale.EN_BI, locale.EN_VU, locale.YUE_HANT_CN, locale.ZH_HANT_MY, locale.AK, locale.EN_NA, locale.ES_BZ, locale.UG, locale.GV, locale.MT, locale.TEO_KE, locale.ZH_HANT_HK, locale.KO_KP, locale.MS_ARAB, locale.ZH_HANS, locale.EN_TV, locale.GD, locale.SO_ET, locale.CGG, locale.EN_MP, locale.ES_BR, locale.VAI, locale.ZH_HANT_MO, locale.EN_MG, locale.NAQ, locale.EN_AE, locale.EN_SX, locale.ZH_HANT, locale.EN_MO, locale.EN_NF, locale.IG, locale.ZH_HANS_MY, locale.HAW, locale.EBU, locale.EN_CX, locale.VAI_VAII, locale.ST_LS, locale.EN_SD, locale.ZH_HANS_HK, locale.EN_TO, locale.ML, locale.EN_BW, locale.EN_SB, locale.EN_GG, locale.ST, locale.YO, locale.SO, locale.EN_KE, locale.EN_TT, locale.EE, locale.EN_NZ, locale.EN_BS, locale.EN_HK, locale.EN_MT, locale.ES_DO, locale.BM, locale.EN_CA, locale.EN_FM, locale.EN_LR, locale.ES_NI, locale.EN_GB, locale.SO_KE, locale.GA_GB, locale.EN_GU, locale.EN_KI, locale.EN_VG, locale.MN_MONG_MN, locale.EN_WS, locale.OR, locale.EE_TG, locale.EN_CY, locale.EN_IE, locale.EN_NG, locale.ND, locale.EN_PW, locale.ES_PR, locale.CEB, locale.EN_NU, locale.YUE, locale.AM, locale.EN_CM, locale.EN_VC, locale.SI, locale.EN_BZ, locale.KAM, locale.EN_AU:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.TRW, locale.CO, locale.MNI_MTEI, locale.PAP_AW, locale.SID, locale.HNJ, locale.LTG, locale.MNI_BENG, locale.AZ_ARAB_TR, locale.SYR, locale.II, locale.LA, locale.FRR, locale.VO, locale.AN, locale.RHG_ROHG, locale.MNI, locale.CIC, locale.KOK, locale.KS_DEVA, locale.SSY, locale.BAL_ARAB, locale.CU, locale.JBO, locale.GAA, locale.UND, locale.SCN, locale.TIG, locale.TYV, locale.SD, locale.RIF, locale.GN, locale.SDH, locale.SMS, locale.KPE_GN, locale.CAD, locale.HA_ARAB, locale.IU_LATN, locale.LAG, locale.AB, locale.EN_DSRT, locale.SKR, locale.EN_SHAW, locale.SW_KE, locale.WBP, locale.KEN, locale.OSA, locale.AA_DJ, locale.HA_NE, locale.MHN, locale.SW, locale.IU, locale.MGO, locale.MI, locale.MN_MONG, locale.BO, locale.LRC, locale.AA_ER, locale.BEW, locale.QUC, locale.KOK_DEVA, locale.RAJ, locale.SAT_OLCK, locale.TA_SG, locale.BAL, locale.KAJ, locale.MOH, locale.TRV, locale.MUS, locale.HA, locale.HA_GH, locale.MAI, locale.BLT, locale.SHN_TH, locale.DV, locale.AZ_ARAB_IQ, locale.BO_IN, locale.QU_EC, locale.TO, locale.ZA, locale.GEZ_ER, locale.SDH_IQ, locale.MDF, locale.MYV, locale.PAP, locale.MZN, locale.RHG_ROHG_BD, locale.HA_ARAB_SD, locale.IO, locale.CSW, locale.BSS, locale.SD_DEVA, locale.LKT, locale.CHO, locale.PIS, locale.SAT_DEVA, locale.SMJ, locale.BGC, locale.KAA_CYRL, locale.BM_NKOO, locale.CCH, locale.CKB, locale.LRC_IQ, locale.NV, locale.SMA_NO, locale.ES_PE, locale.SW_UG, locale.WAL, locale.SHN, locale.MN, locale.SD_ARAB, locale.AA, locale.KK_ARAB, locale.NY, locale.SYR_SY, locale.SMJ_NO, locale.EN_MV, locale.KAA, locale.KPE, locale.SAT, locale.HNJ_HMNP, locale.ZH_LATN, locale.APC, locale.MIC, locale.KCG, locale.YI, locale.KAA_LATN, locale.RHG, locale.ARN, locale.GEZ, locale.SMA, locale.BYN, locale.WA, locale.ANN, locale.QU, locale.MG, locale.TA_MY, locale.CKB_IR, locale.AZ_ARAB, locale.BA:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 68 locales
	case locale.LV, locale.RU_KG, locale.EN_FI, locale.NMG, locale.TG, locale.FF, locale.FF_LATN_MR, locale.KK_KZ, locale.UZ_LATN, locale.CS, locale.HU, locale.EWO, locale.KEA, locale.RU_KZ, locale.UZ_CYRL, locale.YAV, locale.SK, locale.PT_CV, locale.FF_LATN, locale.SZL, locale.PT_MO, locale.FF_LATN_SL, locale.PT_PT, locale.KK_CYRL, locale.TK, locale.FF_LATN_NE, locale.PT_CH, locale.UK, locale.PT_GW, locale.PT_LU, locale.PRG, locale.FF_LATN_LR, locale.DUA, locale.FF_LATN_GH, locale.FF_LATN_GW, locale.CV, locale.PT_ST, locale.SMN, locale.SAH, locale.RU_BY, locale.PT_GQ, locale.EN_SE, locale.BAS, locale.HY, locale.RU_MD, locale.UZ, locale.KY, locale.BR, locale.FF_LATN_BF, locale.KSF, locale.FF_LATN_GN, locale.TZM, locale.BE, locale.PL, locale.FF_LATN_NG, locale.PT_MZ, locale.PT_AO, locale.FF_LATN_GM, locale.RU, locale.RU_UA, locale.EO, locale.BE_TARASK, locale.FF_LATN_CM, locale.PT_TL, locale.KA, locale.TT, locale.DYO, locale.KK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 61 locales
	case locale.GL, locale.HSB, locale.VMW, locale.AZ_CYRL, locale.DE_BE, locale.IT, locale.IT_VA, locale.EN_BE, locale.LB, locale.LN_AO, locale.NDS_NL, locale.RO_MD, locale.AZ, locale.DSB, locale.IS, locale.LLD, locale.ES, locale.SR_CYRL, locale.MK, locale.SC, locale.LN_CF, locale.DA, locale.NDS, locale.SR_CYRL_XK, locale.EL, locale.DE_IT, locale.CA_FR, locale.KU, locale.LIJ, locale.AZ_LATN, locale.CA_AD, locale.SR_CYRL_BA, locale.ES_PH, locale.SR_CYRL_ME, locale.AST, locale.BS, locale.DE_LU, locale.EN_DK, locale.CA_ES_VALENCIA, locale.EN_SI, locale.BS_CYRL, locale.IT_SM, locale.EN_DE, locale.LN, locale.LN_CG, locale.SR_LATN_ME, locale.ES_IC, locale.DA_GL, locale.SR_LATN_XK, locale.SR_LATN_BA, locale.VI, locale.BS_LATN, locale.RO, locale.EL_CY, locale.EL_POLYTON, locale.SR, locale.ES_EA, locale.CA_IT, locale.DE, locale.CA, locale.SR_LATN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CD, locale.FR_MF, locale.FR_MC, locale.FR_GP, locale.FR_MR, locale.FR_CH, locale.FR, locale.FR_GQ, locale.FR_GN, locale.FR_BL, locale.FR_GF, locale.FR_PF, locale.FR_YT, locale.FR_CF, locale.FR_TN, locale.FR_SN, locale.FR_BJ, locale.FR_KM, locale.FR_BF, locale.FR_SC, locale.FR_NC, locale.FR_TG, locale.FR_MG, locale.FR_DZ, locale.FR_CI, locale.FR_NE, locale.FR_TD, locale.FR_GA, locale.FR_BE, locale.FR_PM, locale.FR_VU, locale.FR_DJ, locale.FR_RE, locale.FR_MQ, locale.FR_RW, locale.FR_CM, locale.FR_ML, locale.FR_WF, locale.FR_HT, locale.FR_CG, locale.FR_BI, locale.FR_MU, locale.FR_SY:
		return CurrencyFormatInfo{
			Symbol:           "$CO",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SO, locale.AR_EH, locale.AR_YE, locale.AR_ER, locale.AR_QA, locale.AR_OM, locale.AR, locale.AR_SS, locale.AR_IQ, locale.AR_KM, locale.AR_DJ, locale.AR_EG, locale.AR_BH, locale.AR_TD, locale.AR_AE, locale.AR_PS, locale.AR_JO, locale.AR_SD, locale.AR_IL, locale.AR_KW, locale.AR_SY, locale.AR_SA:
		return CurrencyFormatInfo{
			Symbol:           "CO$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.IA, locale.JGO, locale.SW_CD, locale.KKJ, locale.JV, locale.ES_CO, locale.WO, locale.MGH, locale.FUR, locale.ES_UY, locale.PT, locale.ES_AR, locale.NNH, locale.QU_BO, locale.YRL, locale.YRL_CO, locale.EN_AT, locale.KGP, locale.MS_BN, locale.YRL_VE, locale.MS_ARAB_BN, locale.RW:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.XNR, locale.KXV_LATN, locale.KXV_DEVA, locale.EN_IN, locale.PA_GURU, locale.SA, locale.HI_LATN, locale.KXV_TELU, locale.BN_IN, locale.KXV, locale.DZ, locale.HI, locale.PA, locale.GU, locale.KOK_LATN, locale.KXV_ORYA, locale.TA, locale.TA_LK, locale.TE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM, locale.FF_ADLM_MR, locale.FF_ADLM_GH, locale.FF_ADLM_LR, locale.FF_ADLM_BF, locale.FF_ADLM_GW, locale.FF_ADLM_CM, locale.FF_ADLM_SL, locale.FF_ADLM_NG, locale.FF_ADLM_NE, locale.FF_ADLM_GM, locale.FF_ADLM_SN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.ES_BO, locale.MS_ID, locale.SU_LATN, locale.SU, locale.TR, locale.TR_CY, locale.EN_ID, locale.ES_GQ, locale.ID, locale.MUA:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_BQ, locale.NL_CW, locale.NL_AW, locale.ES_PY, locale.EN_NL, locale.NL_BE, locale.NL_SX, locale.NL_SR, locale.NL:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.SV, locale.SE_SE, locale.ET, locale.SV_FI, locale.SV_AX, locale.SE, locale.LT, locale.KSH, locale.SE_FI:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 8 locales
	case locale.EN_ZA, locale.VE, locale.NR, locale.SS, locale.SS_SZ, locale.ES_CR, locale.AF, locale.AF_NA:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI, locale.SHI_TFNG, locale.AGQ, locale.OC, locale.SHI_LATN, locale.ZGH, locale.OC_ES, locale.KAB:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.BEZ, locale.SBP, locale.LUO, locale.KSB, locale.LG, locale.KM, locale.RWK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_LB, locale.AR_MR, locale.AR_MA, locale.AR_DZ, locale.AR_TN, locale.AR_LY:
		return CurrencyFormatInfo{
			Symbol:           "CO$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ASA, locale.MY, locale.EN_150, locale.CE, locale.XOG, locale.TPI:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.ES_EC, locale.KL, locale.LO, locale.ES_VE, locale.ES_CL, locale.SG:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.HR, locale.SL, locale.FO, locale.EU, locale.HR_BA, locale.FO_DK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.BGN_AF, locale.BGN_AE, locale.BGN, locale.BGN_IR, locale.BGN_OM:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.SQ_MK, locale.SQ, locale.FR_CA, locale.BG, locale.SQ_XK:
		return CurrencyFormatInfo{
			Symbol:           "COP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_LI, locale.RM, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.KHQ, locale.DJE, locale.TWQ, locale.SES:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NN, locale.NB, locale.NB_SJ, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS, locale.DE_AT, locale.OS_RU, locale.TS:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.EN_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.NE, locale.AS:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.PS, locale.UZ_ARAB, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 3 locales
	case locale.SEH, locale.RN, locale.LU:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.FR_MA, locale.FR_LU:
		return CurrencyFormatInfo{
			Symbol:           "$CO",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN, locale.TN_BW:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "KLBP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "सि.अ.पि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "COP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߞߐߔ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
