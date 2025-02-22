// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = BDT{}

// Taka currency type
type BDT struct{}

func (BDT) Code() string { return "BDT" }

func (BDT) Number() string { return "050" }

func (BDT) Name() string { return "Taka" }

func (BDT) MinorUnits() int { return 2 }

func (BDT) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 192 locales
	case locale.EBU, locale.EN, locale.TEO, locale.EN_SB, locale.SAQ, locale.EN_BB, locale.EN_UG, locale.EN_GU, locale.MT, locale.EN_GB, locale.EN_JM, locale.EN_BW, locale.EN_BI, locale.EN_VG, locale.ES_GT, locale.EN_TZ, locale.TH, locale.EN_KY, locale.EN_NG, locale.EN_SG, locale.NUS, locale.AK, locale.ML, locale.SO_DJ, locale.ZH_HANS_MY, locale.EN_DM, locale.EN_MG, locale.EN_LS, locale.GD, locale.OM, locale.EN_CA, locale.NYN, locale.OM_KE, locale.EN_AG, locale.OR, locale.EN_IO, locale.EN_NU, locale.EN_BZ, locale.EN_GM, locale.ES_DO, locale.EN_IL, locale.KAM, locale.ST, locale.DOI, locale.EN_MO, locale.EN_PG, locale.ES_BZ, locale.ES_CU, locale.VAI, locale.SO, locale.EN_FJ, locale.EN_CX, locale.EN_MT, locale.GA, locale.GUZ, locale.EN_ER, locale.ND, locale.EN_PK, locale.EN_NA, locale.EN_SS, locale.ES_BR, locale.ES_MX, locale.BM, locale.EN_LC, locale.EN_TK, locale.EN_MU, locale.EN_SL, locale.EN_VC, locale.HAW, locale.JA, locale.EN_SD, locale.GV, locale.JMC, locale.SO_ET, locale.VAI_VAII, locale.TI, locale.KW, locale.MR, locale.EN_RW, locale.ES_PR, locale.IG, locale.PCM, locale.BEM, locale.EN_SZ, locale.ES_SV, locale.EE_TG, locale.EN_ZW, locale.EN_FM, locale.UG, locale.EE, locale.EN_SC, locale.EN_GI, locale.EN_IE, locale.FIL, locale.EN_LR, locale.SN, locale.EN_KI, locale.EN_PW, locale.KDE, locale.SO_KE, locale.EN_NF, locale.YUE, locale.EN_BM, locale.ZH_HANS_SG, locale.SI, locale.VAI_LATN, locale.ZH_HANT_HK, locale.EN_MS, locale.MAS, locale.NAQ, locale.YUE_HANS, locale.EN_UM, locale.ES_US, locale.MS_SG, locale.TI_ER, locale.AM, locale.EN_AE, locale.EN_SX, locale.CHR, locale.MS, locale.CGG, locale.EN_CK, locale.BHO, locale.ZU, locale.ES_PA, locale.EN_JE, locale.EN_TV, locale.KI, locale.ZH_HANT_MO, locale.EN_IM, locale.ES_HN, locale.VUN, locale.EN_PN, locale.EN_MP, locale.KO, locale.KO_CN, locale.EN_SH, locale.MAS_TZ, locale.EN_TO, locale.EN_VI, locale.EN_WS, locale.KO_KP, locale.CEB, locale.ROF, locale.EN_PH, locale.ZH_HANT, locale.ST_LS, locale.TEO_KE, locale.YO, locale.EN_KN, locale.EN_NR, locale.EN_AS, locale.EN_BS, locale.EN_GD, locale.ZH_HANS, locale.EN_001, locale.KLN, locale.EN_ZM, locale.EN_CY, locale.EN_NZ, locale.EN_TT, locale.DAV, locale.EN_CC, locale.EN_GY, locale.EN_HK, locale.EN_DG, locale.YUE_HANT, locale.ES_419, locale.EN_GH, locale.EN_KE, locale.ES_NI, locale.ZH, locale.EN_PR, locale.YO_BJ, locale.EN_CM, locale.EN_MW, locale.KN, locale.YUE_HANT_CN, locale.ZH_HANS_MO, locale.EN_GG, locale.EN_AI, locale.MER, locale.EN_MY, locale.EN_FK, locale.ZH_HANT_MY, locale.EN_VU, locale.ZH_HANS_HK, locale.MN_MONG_MN, locale.MS_ARAB, locale.EN_TC, locale.EN_MH, locale.GA_GB:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.KAA_CYRL, locale.SYR_SY, locale.SSY, locale.WBP, locale.CCH, locale.SHN, locale.IU, locale.MYV, locale.BO, locale.SAT_OLCK, locale.DV, locale.MUS, locale.SID, locale.ZH_LATN, locale.ANN, locale.PAP_AW, locale.AZ_ARAB_IQ, locale.MZN, locale.HA_ARAB_SD, locale.SW, locale.YI, locale.BSS, locale.IO, locale.TRV, locale.GN, locale.GEZ_ER, locale.SYR, locale.VO, locale.FRR, locale.ES_PE, locale.EN_SHAW, locale.MNI_BENG, locale.SW_KE, locale.BA, locale.BGC, locale.MIC, locale.RHG, locale.KEN, locale.PIS, locale.LA, locale.MI, locale.SAT, locale.SDH, locale.BAL_ARAB, locale.II, locale.KS_DEVA, locale.SAT_DEVA, locale.CHO, locale.SDH_IQ, locale.ZA, locale.PAP, locale.GEZ, locale.TIG, locale.KOK, locale.BM_NKOO, locale.MN_MONG, locale.SCN, locale.AZ_ARAB_TR, locale.CU, locale.HA_ARAB, locale.CO, locale.KPE_GN, locale.BLT, locale.AA_ER, locale.SMJ, locale.MGO, locale.RHG_ROHG, locale.SMJ_NO, locale.WA, locale.HA, locale.HNJ, locale.KK_ARAB, locale.BEW, locale.CAD, locale.MDF, locale.RHG_ROHG_BD, locale.KCG, locale.NV, locale.BYN, locale.CKB, locale.TRW, locale.AB, locale.SD_DEVA, locale.LKT, locale.UND, locale.ARN, locale.RAJ, locale.SD_ARAB, locale.WAL, locale.JBO, locale.LRC_IQ, locale.SMA, locale.TA_MY, locale.AN, locale.GAA, locale.CKB_IR, locale.MN, locale.QU_EC, locale.KPE, locale.SMS, locale.HA_NE, locale.AZ_ARAB, locale.TA_SG, locale.IU_LATN, locale.NY, locale.SD, locale.AA, locale.TO, locale.AA_DJ, locale.MNI, locale.BO_IN, locale.HA_GH, locale.MOH, locale.LRC, locale.MAI, locale.CIC, locale.CSW, locale.TYV, locale.LAG, locale.MNI_MTEI, locale.QUC, locale.RIF, locale.KOK_DEVA, locale.BAL, locale.KAA, locale.KAA_LATN, locale.QU, locale.LTG, locale.KAJ, locale.MG, locale.OSA, locale.EN_DSRT, locale.APC, locale.HNJ_HMNP, locale.MHN, locale.SMA_NO, locale.EN_MV, locale.SHN_TH, locale.SKR, locale.SW_UG:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.HY, locale.PT_GQ, locale.FF_LATN_SL, locale.CV, locale.FF_LATN_MR, locale.KK_KZ, locale.YAV, locale.FF_LATN_BF, locale.PT_AO, locale.SZL, locale.KSF, locale.KY, locale.TK, locale.EWO, locale.FF_LATN_GW, locale.SAH, locale.PT_MZ, locale.FF_LATN_GN, locale.RU, locale.FF_LATN_LR, locale.PT_PT, locale.FF_LATN, locale.DYO, locale.LV, locale.FF_LATN_GM, locale.KEA, locale.BAS, locale.NMG, locale.PT_CH, locale.EN_SE, locale.PT_GW, locale.RU_KZ, locale.UK, locale.TG, locale.FF_LATN_GH, locale.PT_ST, locale.SK, locale.FF, locale.FF_LATN_NE, locale.UZ_LATN, locale.TT, locale.HU, locale.PT_TL, locale.KK_CYRL, locale.UZ, locale.CS, locale.KK, locale.PT_CV, locale.EO, locale.RU_BY, locale.FF_LATN_NG, locale.DUA, locale.FR_CA, locale.PT_LU, locale.RU_KG, locale.BE_TARASK, locale.TZM, locale.RU_MD, locale.UZ_CYRL, locale.PL, locale.KA, locale.BE, locale.PRG, locale.PT_MO, locale.BR, locale.FF_LATN_CM, locale.EN_FI, locale.SMN, locale.RU_UA:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 63 locales
	case locale.CA_FR, locale.CA_IT, locale.DA, locale.IT, locale.EL_CY, locale.AZ, locale.BS_CYRL, locale.MK, locale.DE_LU, locale.KU, locale.LN, locale.SR, locale.EN_BE, locale.BS, locale.EL_POLYTON, locale.VMW, locale.CA_ES_VALENCIA, locale.EN_SI, locale.GL, locale.LB, locale.EL, locale.IT_VA, locale.CA_AD, locale.ES, locale.FR_LU, locale.SR_CYRL_XK, locale.FR_MA, locale.LLD, locale.RO_MD, locale.SR_LATN_XK, locale.EN_DK, locale.NDS, locale.RO, locale.EN_DE, locale.SR_CYRL, locale.DE, locale.DA_GL, locale.ES_IC, locale.HSB, locale.SR_CYRL_BA, locale.AST, locale.SR_LATN_ME, locale.DE_BE, locale.BS_LATN, locale.IT_SM, locale.NDS_NL, locale.DE_IT, locale.SC, locale.SR_CYRL_ME, locale.ES_EA, locale.SR_LATN_BA, locale.LIJ, locale.CA, locale.AZ_CYRL, locale.LN_CF, locale.LN_CG, locale.AZ_LATN, locale.ES_PH, locale.SR_LATN, locale.IS, locale.VI, locale.LN_AO, locale.DSB:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_TD, locale.FR_WF, locale.FR_TG, locale.FR_CH, locale.FR_GA, locale.FR_GP, locale.FR_TN, locale.FR_MR, locale.FR_GQ, locale.FR_CG, locale.FR_HT, locale.FR_PM, locale.FR_MC, locale.FR_MG, locale.FR_CD, locale.FR_RW, locale.FR_MF, locale.FR_CF, locale.FR_ML, locale.FR_SY, locale.FR, locale.FR_CM, locale.FR_KM, locale.FR_YT, locale.FR_BE, locale.FR_BL, locale.FR_PF, locale.FR_VU, locale.FR_DJ, locale.FR_BJ, locale.FR_SC, locale.FR_MQ, locale.FR_NC, locale.FR_MU, locale.FR_GF, locale.FR_SN, locale.FR_BF, locale.FR_DZ, locale.FR_CI, locale.FR_RE, locale.FR_NE, locale.FR_BI, locale.FR_GN:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_IL, locale.AR_KM, locale.AR_SY, locale.AR_IQ, locale.AR, locale.AR_QA, locale.AR_SS, locale.AR_OM, locale.AR_AE, locale.AR_EG, locale.AR_PS, locale.AR_EH, locale.AR_YE, locale.AR_DJ, locale.AR_SA, locale.AR_SD, locale.AR_ER, locale.AR_TD, locale.AR_SO, locale.AR_BH, locale.AR_JO, locale.AR_KW:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.YRL_VE, locale.KGP, locale.NNH, locale.JV, locale.SW_CD, locale.FUR, locale.QU_BO, locale.ES_UY, locale.IA, locale.PT, locale.MS_BN, locale.ES_CO, locale.EN_AT, locale.WO, locale.ES_AR, locale.MGH, locale.YRL, locale.YRL_CO, locale.KKJ, locale.MS_ARAB_BN, locale.JGO, locale.RW:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.BN_IN, locale.HI_LATN, locale.TA_LK, locale.KXV_ORYA, locale.DZ, locale.KOK_LATN, locale.GU, locale.KXV_LATN, locale.KXV_DEVA, locale.PA_GURU, locale.XNR, locale.PA, locale.HI, locale.KXV, locale.TA, locale.KXV_TELU, locale.SA, locale.EN_IN, locale.TE:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_GM, locale.FF_ADLM_NG, locale.FF_ADLM, locale.FF_ADLM_SN, locale.FF_ADLM_BF, locale.FF_ADLM_CM, locale.FF_ADLM_GW, locale.FF_ADLM_GH, locale.FF_ADLM_LR, locale.FF_ADLM_SL:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.SU_LATN, locale.SU, locale.ID, locale.MS_ID, locale.EN_ID, locale.MUA, locale.ES_BO, locale.TR_CY, locale.ES_GQ, locale.TR:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL_AW, locale.NL_SX, locale.EN_NL, locale.NL_BQ, locale.NL, locale.NL_SR, locale.NL_CW, locale.ES_PY, locale.NL_BE:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.ET, locale.SE_SE, locale.SE_FI, locale.KSH, locale.SE, locale.SV, locale.SV_FI, locale.SV_AX:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 8 locales
	case locale.SHI_TFNG, locale.SHI_LATN, locale.ZGH, locale.OC_ES, locale.OC, locale.SHI, locale.KAB, locale.AGQ:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SS_SZ, locale.AF_NA, locale.NR, locale.SS, locale.VE, locale.EN_ZA, locale.ES_CR, locale.AF:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KSB, locale.RWK, locale.LUO, locale.BEZ, locale.SBP, locale.KM, locale.LG:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MR, locale.AR_DZ, locale.AR_LB, locale.AR_TN, locale.AR_MA, locale.AR_LY:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.HR, locale.SL, locale.FO_DK, locale.FO, locale.EU, locale.HR_BA:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.LO, locale.ES_EC, locale.SG, locale.ES_VE, locale.KL, locale.ES_CL:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.MY, locale.TPI, locale.XOG, locale.CE, locale.ASA, locale.EN_150:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_AE, locale.BGN, locale.BGN_IR, locale.BGN_AF, locale.BGN_OM:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.BRX, locale.NE, locale.NE_IN, locale.AS:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_FR, locale.GSW_LI, locale.GSW, locale.RM:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB, locale.NN, locale.NO, locale.NB_SJ:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS_RU, locale.OS, locale.DE_AT, locale.TS:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.SES, locale.DJE, locale.TWQ, locale.KHQ:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.SQ, locale.BG, locale.SQ_MK, locale.SQ_XK:
		return CurrencyFormatInfo{
			Symbol:           "BDT",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.DE_CH, locale.IT_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.RN, locale.LU, locale.SEH:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.FI, locale.LT:
		return CurrencyFormatInfo{
			Symbol:           "BDT",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "BGDT",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.CY:
		return CurrencyFormatInfo{
			Symbol:           "TK",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AU:
		return CurrencyFormatInfo{
			Symbol:           "Tk",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߓߘߕ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "৳",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
