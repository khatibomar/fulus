// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = CRC{}

// Costa Rican Colon currency type
type CRC struct{}

func (CRC) Code() string { return "CRC" }

func (CRC) Number() string { return "188" }

func (CRC) Name() string { return "Costa Rican Colon" }

func (CRC) MinorUnits() int { return 2 }

func (CRC) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 194 locales
	case locale.EN_PG, locale.SO, locale.EN_PH, locale.ZH, locale.EN_SC, locale.ES_MX, locale.OM, locale.EN_NU, locale.AK, locale.TEO, locale.EN_ER, locale.KDE, locale.EN_GG, locale.EN_MY, locale.EN_TO, locale.EN_VU, locale.EN_IM, locale.BEM, locale.EN_BB, locale.EN_GI, locale.KW, locale.EN_MU, locale.ZH_HANT_MY, locale.EBU, locale.EN_CC, locale.MS_SG, locale.PCM, locale.EN_GD, locale.AM, locale.CGG, locale.CY, locale.DOI, locale.EN_BS, locale.EN_NA, locale.KAM, locale.ZH_HANS, locale.EN_SS, locale.ES_DO, locale.KLN, locale.EN_BZ, locale.EN_AI, locale.EN_SD, locale.EN, locale.EN_JM, locale.EN_KY, locale.EN_PK, locale.MS, locale.DAV, locale.EN_BM, locale.EN_IO, locale.EN_ZW, locale.NUS, locale.ZH_HANS_SG, locale.BM, locale.ES_BR, locale.YO, locale.ZH_HANT, locale.EN_MT, locale.EN_NZ, locale.EN_BI, locale.EN_CK, locale.MER, locale.EN_LR, locale.ZH_HANS_HK, locale.EN_SZ, locale.EN_WS, locale.YO_BJ, locale.EN_AU, locale.ES_GT, locale.ROF, locale.SAQ, locale.EN_MG, locale.EN_SX, locale.MAS_TZ, locale.ES_HN, locale.EN_TZ, locale.EE, locale.ZH_HANS_MO, locale.GD, locale.EN_KN, locale.EN_SG, locale.ML, locale.EN_LC, locale.SN, locale.VAI_LATN, locale.EN_NR, locale.OM_KE, locale.GV, locale.TI, locale.ZH_HANT_HK, locale.VAI_VAII, locale.EN_NF, locale.EN_RW, locale.EN_TT, locale.EN_PR, locale.YUE_HANS, locale.JA, locale.YUE, locale.EN_FJ, locale.ES_PR, locale.IG, locale.SO_KE, locale.KI, locale.UG, locale.EN_MH, locale.EN_MO, locale.TH, locale.EN_TC, locale.SO_DJ, locale.SO_ET, locale.EN_GY, locale.EN_VG, locale.EN_BW, locale.EN_CX, locale.OR, locale.GA_GB, locale.ES_US, locale.EN_MW, locale.EN_ZM, locale.ES_SV, locale.MT, locale.EN_AS, locale.MR, locale.GA, locale.EN_FM, locale.EN_KI, locale.FIL, locale.EN_HK, locale.EN_VC, locale.YUE_HANT, locale.EN_IE, locale.ES_BZ, locale.EN_PN, locale.EN_GM, locale.EN_VI, locale.EN_GU, locale.EN_NG, locale.MAS, locale.EN_FK, locale.EN_PW, locale.EN_UG, locale.EN_JE, locale.EN_LS, locale.KO_KP, locale.GUZ, locale.TI_ER, locale.EN_TK, locale.ES_CU, locale.EN_GH, locale.EN_AE, locale.EN_SH, locale.KO_CN, locale.EN_001, locale.ES_PA, locale.ST, locale.VUN, locale.BHO, locale.EN_SB, locale.NAQ, locale.EE_TG, locale.MS_ARAB, locale.VAI, locale.ZU, locale.HAW, locale.EN_DG, locale.EN_AG, locale.ST_LS, locale.EN_MS, locale.EN_UM, locale.EN_GB, locale.EN_TV, locale.ZH_HANS_MY, locale.EN_MP, locale.EN_SL, locale.CEB, locale.EN_CA, locale.KO, locale.EN_DM, locale.MN_MONG_MN, locale.YUE_HANT_CN, locale.ES_419, locale.ES_NI, locale.TEO_KE, locale.SI, locale.CHR, locale.EN_CM, locale.JMC, locale.EN_IL, locale.NYN, locale.EN_CY, locale.KN, locale.ND, locale.ZH_HANT_MO, locale.EN_KE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.HNJ_HMNP, locale.MZN, locale.SSY, locale.IO, locale.ANN, locale.AZ_ARAB_TR, locale.NY, locale.BAL, locale.CO, locale.MN_MONG, locale.SCN, locale.MN, locale.HA_NE, locale.AN, locale.MNI_BENG, locale.BAL_ARAB, locale.AZ_ARAB_IQ, locale.CU, locale.SDH, locale.MUS, locale.SYR, locale.HA, locale.EN_DSRT, locale.LA, locale.MNI_MTEI, locale.PAP_AW, locale.SDH_IQ, locale.BM_NKOO, locale.KAA_CYRL, locale.KCG, locale.SMJ, locale.SMJ_NO, locale.CSW, locale.MGO, locale.HA_ARAB, locale.IU_LATN, locale.AA_ER, locale.ARN, locale.KAJ, locale.SW, locale.KOK_DEVA, locale.CHO, locale.KK_ARAB, locale.CAD, locale.GAA, locale.RAJ, locale.KAA_LATN, locale.KOK, locale.BO, locale.CKB_IR, locale.EN_SHAW, locale.MNI, locale.RHG_ROHG, locale.IU, locale.SW_KE, locale.MOH, locale.TA_MY, locale.CCH, locale.LKT, locale.KEN, locale.BO_IN, locale.SD_DEVA, locale.LRC_IQ, locale.TYV, locale.APC, locale.PAP, locale.GEZ_ER, locale.LRC, locale.AB, locale.SKR, locale.PIS, locale.QU_EC, locale.KS_DEVA, locale.ZH_LATN, locale.MI, locale.MDF, locale.TRW, locale.WAL, locale.BA, locale.BEW, locale.OSA, locale.SAT, locale.MG, locale.AA_DJ, locale.QU, locale.SMA_NO, locale.HA_GH, locale.JBO, locale.SAT_DEVA, locale.SD_ARAB, locale.WA, locale.BLT, locale.HA_ARAB_SD, locale.TA_SG, locale.SMA, locale.SAT_OLCK, locale.AA, locale.TO, locale.BSS, locale.HNJ, locale.CKB, locale.SMS, locale.LAG, locale.BGC, locale.QUC, locale.RHG_ROHG_BD, locale.SID, locale.YI, locale.GEZ, locale.MAI, locale.CIC, locale.SHN, locale.ES_PE, locale.KPE, locale.ZA, locale.RHG, locale.TRV, locale.II, locale.MIC, locale.UND, locale.LTG, locale.TIG, locale.SD, locale.KPE_GN, locale.SYR_SY, locale.BYN, locale.DV, locale.NV, locale.SHN_TH, locale.EN_MV, locale.VO, locale.MYV, locale.RIF, locale.GN, locale.KAA, locale.AZ_ARAB, locale.MHN, locale.FRR, locale.SW_UG, locale.WBP:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 69 locales
	case locale.FF_LATN, locale.FF_LATN_MR, locale.FF_LATN_SL, locale.PT_ST, locale.RU_UA, locale.LV, locale.PT_LU, locale.UZ_LATN, locale.TG, locale.FF_LATN_GM, locale.SAH, locale.CS, locale.RU_KZ, locale.HY, locale.RU_KG, locale.EN_FI, locale.TZM, locale.UZ_CYRL, locale.RU, locale.UZ, locale.FF_LATN_NE, locale.YAV, locale.EO, locale.KSF, locale.KK_CYRL, locale.KK, locale.PT_PT, locale.PT_MZ, locale.KEA, locale.PT_GW, locale.BR, locale.UK, locale.DYO, locale.BAS, locale.KY, locale.HU, locale.PRG, locale.FR_CA, locale.PT_CV, locale.BE_TARASK, locale.CV, locale.FF_LATN_LR, locale.TT, locale.EWO, locale.FF_LATN_CM, locale.PT_CH, locale.BE, locale.PT_TL, locale.TK, locale.EN_SE, locale.PT_AO, locale.FF_LATN_GW, locale.SZL, locale.PL, locale.KA, locale.PT_MO, locale.SK, locale.FF_LATN_GN, locale.RU_MD, locale.FF, locale.SMN, locale.NMG, locale.RU_BY, locale.FF_LATN_NG, locale.FF_LATN_GH, locale.KK_KZ, locale.FF_LATN_BF, locale.PT_GQ, locale.DUA:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 63 locales
	case locale.DE_BE, locale.SR_LATN_ME, locale.IS, locale.LLD, locale.SR_CYRL_XK, locale.SR_CYRL_ME, locale.EN_DK, locale.MK, locale.CA_ES_VALENCIA, locale.EL, locale.BS_LATN, locale.SR_LATN_BA, locale.ES_IC, locale.CA_AD, locale.ES_PH, locale.DA_GL, locale.AST, locale.VI, locale.FR_MA, locale.AZ_LATN, locale.SR_CYRL, locale.IT_VA, locale.RO, locale.LIJ, locale.SC, locale.CA_IT, locale.KU, locale.DSB, locale.BS_CYRL, locale.HSB, locale.AZ_CYRL, locale.ES_EA, locale.RO_MD, locale.EN_BE, locale.EN_SI, locale.ES, locale.DE_IT, locale.EN_DE, locale.BS, locale.LN, locale.DA, locale.SR, locale.SR_LATN_XK, locale.AZ, locale.LN_AO, locale.LN_CF, locale.EL_CY, locale.CA, locale.GL, locale.DE_LU, locale.SR_LATN, locale.IT_SM, locale.LN_CG, locale.LB, locale.NDS_NL, locale.SR_CYRL_BA, locale.DE, locale.FR_LU, locale.VMW, locale.CA_FR, locale.EL_POLYTON, locale.NDS, locale.IT:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_MF, locale.FR_MC, locale.FR_HT, locale.FR_DJ, locale.FR_PF, locale.FR_ML, locale.FR_BJ, locale.FR_CF, locale.FR_TN, locale.FR_KM, locale.FR_NE, locale.FR_SC, locale.FR_BI, locale.FR_SN, locale.FR_YT, locale.FR_MG, locale.FR_NC, locale.FR_MQ, locale.FR_BF, locale.FR_CM, locale.FR_MU, locale.FR_BE, locale.FR_GA, locale.FR_BL, locale.FR_SY, locale.FR, locale.FR_MR, locale.FR_VU, locale.FR_RW, locale.FR_PM, locale.FR_GP, locale.FR_DZ, locale.FR_GF, locale.FR_CG, locale.FR_WF, locale.FR_TG, locale.FR_CD, locale.FR_CI, locale.FR_RE, locale.FR_GN, locale.FR_CH, locale.FR_GQ, locale.FR_TD:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_KM, locale.AR_SA, locale.AR, locale.AR_EH, locale.AR_IQ, locale.AR_YE, locale.AR_SO, locale.AR_JO, locale.AR_KW, locale.AR_EG, locale.AR_SD, locale.AR_DJ, locale.AR_IL, locale.AR_BH, locale.AR_QA, locale.AR_PS, locale.AR_SY, locale.AR_AE, locale.AR_TD, locale.AR_ER, locale.AR_OM, locale.AR_SS:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.FUR, locale.MS_ARAB_BN, locale.SW_CD, locale.ES_AR, locale.ES_CO, locale.KKJ, locale.QU_BO, locale.EN_AT, locale.PT, locale.ES_UY, locale.JGO, locale.KGP, locale.YRL_VE, locale.NNH, locale.YRL_CO, locale.MGH, locale.YRL, locale.MS_BN, locale.IA, locale.WO, locale.JV, locale.RW:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.KOK_LATN, locale.KXV_DEVA, locale.KXV_TELU, locale.KXV_ORYA, locale.DZ, locale.PA_GURU, locale.TA_LK, locale.TE, locale.KXV_LATN, locale.TA, locale.PA, locale.XNR, locale.GU, locale.HI, locale.HI_LATN, locale.KXV, locale.EN_IN, locale.BN_IN, locale.SA:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_CM, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_GM, locale.FF_ADLM_LR, locale.FF_ADLM_MR, locale.FF_ADLM_NE, locale.FF_ADLM_SL, locale.FF_ADLM_NG, locale.FF_ADLM, locale.FF_ADLM_BF:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.SU_LATN, locale.ID, locale.MUA, locale.TR_CY, locale.SU, locale.TR, locale.MS_ID, locale.EN_ID, locale.ES_BO, locale.ES_GQ:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.LT, locale.SV, locale.SV_AX, locale.SE_FI, locale.ET, locale.SV_FI, locale.SE, locale.KSH, locale.SE_SE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 9 locales
	case locale.NL_CW, locale.ES_PY, locale.NL, locale.NL_SX, locale.NL_SR, locale.EN_NL, locale.NL_AW, locale.NL_BQ, locale.NL_BE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.OC, locale.ZGH, locale.SHI_TFNG, locale.KAB, locale.SHI, locale.AGQ, locale.SHI_LATN, locale.OC_ES:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.VE, locale.SS, locale.AF, locale.ES_CR, locale.SS_SZ, locale.EN_ZA, locale.AF_NA, locale.NR:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.BEZ, locale.RWK, locale.KSB, locale.LUO, locale.KM, locale.SBP, locale.LG:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_MA, locale.AR_TN, locale.AR_LY, locale.AR_MR, locale.AR_DZ, locale.AR_LB:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ASA, locale.XOG, locale.EN_150, locale.TPI, locale.MY, locale.CE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.FO, locale.HR, locale.HR_BA, locale.SL, locale.FO_DK, locale.EU:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.KL, locale.ES_CL, locale.ES_EC, locale.LO, locale.ES_VE, locale.SG:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN, locale.BGN_AE, locale.BGN_OM, locale.BGN_IR, locale.BGN_AF:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.BG, locale.SQ, locale.SQ_XK, locale.SQ_MK:
		return CurrencyFormatInfo{
			Symbol:           "CRC",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.TWQ, locale.SES, locale.KHQ:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_FR, locale.GSW_LI, locale.RM:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NB, locale.NN, locale.NB_SJ, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS, locale.DE_AT, locale.OS_RU, locale.TS:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP_IN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.IT_CH, locale.EN_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LU, locale.SEH, locale.RN:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE, locale.NE_IN, locale.AS:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.PS_PK, locale.UZ_ARAB, locale.PS:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS, locale.KS_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN, locale.TN_BW:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "KRK",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "सि.आर.सि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "CRC",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߞߙߞ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "₡",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
