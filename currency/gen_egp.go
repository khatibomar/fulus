// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = EGP{}

// Egyptian Pound currency type
type EGP struct{}

func (EGP) Code() string { return "EGP" }

func (EGP) Number() string { return "818" }

func (EGP) Name() string { return "Egyptian Pound" }

func (EGP) MinorUnits() int { return 2 }

func (EGP) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 193 locales
	case locale.EN_PG, locale.EN_SX, locale.ZH_HANS_HK, locale.EN_BI, locale.EN_PN, locale.EN_IM, locale.TEO_KE, locale.TI_ER, locale.EN_KN, locale.SO_ET, locale.ES_SV, locale.KAM, locale.EN_DG, locale.EN_IL, locale.EN_PH, locale.EN_GH, locale.NAQ, locale.EN_SL, locale.EN_001, locale.ES_US, locale.EN_ER, locale.EN_GU, locale.EN_ZM, locale.BM, locale.EN_UM, locale.EN_VC, locale.OM, locale.EN_GY, locale.EN_AI, locale.EN_AS, locale.ES_419, locale.ES_BR, locale.EN_GM, locale.EN_JM, locale.EN_SH, locale.EN_LC, locale.EN_TK, locale.ST_LS, locale.ES_PA, locale.EN_GB, locale.OR, locale.EN_KY, locale.NUS, locale.ZH_HANT_MY, locale.EN_BS, locale.EN, locale.KW, locale.EN_NF, locale.EN_WS, locale.VAI_VAII, locale.EN_IO, locale.UG, locale.EN_SG, locale.KI, locale.EN_CY, locale.EN_LR, locale.EN_SB, locale.EN_AG, locale.EN_MY, locale.SAQ, locale.EN_VG, locale.EN_PW, locale.MAS, locale.MR, locale.EN_NZ, locale.ES_DO, locale.YUE_HANT, locale.EN_BW, locale.YUE_HANT_CN, locale.ZH, locale.EE_TG, locale.EN_NR, locale.GD, locale.KO, locale.OM_KE, locale.ES_BZ, locale.IG, locale.YO_BJ, locale.ZH_HANS, locale.EN_IE, locale.EN_DM, locale.AM, locale.EN_BM, locale.EN_SC, locale.ES_MX, locale.GV, locale.EN_CM, locale.EN_TC, locale.BHO, locale.EN_MT, locale.EN_PR, locale.SO, locale.EN_LS, locale.EN_MG, locale.ES_NI, locale.EN_MH, locale.GA_GB, locale.MS, locale.EN_VU, locale.GUZ, locale.NYN, locale.TH, locale.YO, locale.EN_MS, locale.EN_MW, locale.KO_CN, locale.KO_KP, locale.VAI, locale.ES_CU, locale.ZH_HANS_SG, locale.MER, locale.ZU, locale.EN_AE, locale.EN_MU, locale.PCM, locale.EN_KE, locale.AK, locale.KN, locale.DAV, locale.EN_RW, locale.EN_GD, locale.EBU, locale.EN_TO, locale.EN_TV, locale.EN_NG, locale.SN, locale.EN_SS, locale.MN_MONG_MN, locale.JMC, locale.ROF, locale.EN_HK, locale.EN_GI, locale.HAW, locale.ZH_HANT_HK, locale.MT, locale.ZH_HANT_MO, locale.EN_BZ, locale.EN_NA, locale.EN_UG, locale.ND, locale.FIL, locale.SO_DJ, locale.SI, locale.YUE_HANS, locale.EN_SZ, locale.EN_TT, locale.ML, locale.ZH_HANS_MO, locale.EN_TZ, locale.GA, locale.KDE, locale.EN_BB, locale.EN_GG, locale.EN_VI, locale.VUN, locale.CHR, locale.DOI, locale.EE, locale.CEB, locale.ZH_HANS_MY, locale.ZH_HANT, locale.EN_MP, locale.EN_NU, locale.CY, locale.EN_CK, locale.ST, locale.ES_PR, locale.CGG, locale.EN_ZW, locale.EN_KI, locale.EN_SD, locale.YUE, locale.EN_JE, locale.EN_MO, locale.VAI_LATN, locale.BEM, locale.KLN, locale.EN_PK, locale.EN_FJ, locale.EN_FM, locale.ES_HN, locale.MS_ARAB, locale.EN_CX, locale.ES_GT, locale.JA, locale.SO_KE, locale.EN_FK, locale.TI, locale.TEO, locale.EN_CA, locale.EN_CC, locale.MAS_TZ, locale.MS_SG:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.CIC, locale.KS_DEVA, locale.IU_LATN, locale.MYV, locale.SDH, locale.RHG, locale.CSW, locale.HA_ARAB_SD, locale.NV, locale.QUC, locale.GAA, locale.OSA, locale.II, locale.MUS, locale.QU_EC, locale.IU, locale.BYN, locale.KPE_GN, locale.TO, locale.BM_NKOO, locale.CKB, locale.MHN, locale.SHN_TH, locale.BA, locale.AA_ER, locale.KOK_DEVA, locale.NY, locale.SW_UG, locale.KEN, locale.TYV, locale.FRR, locale.SMJ_NO, locale.MAI, locale.QU, locale.SAT_OLCK, locale.SMS, locale.LRC, locale.TRV, locale.WBP, locale.PAP, locale.BAL_ARAB, locale.PIS, locale.ES_PE, locale.MNI, locale.KCG, locale.KAA_LATN, locale.RHG_ROHG, locale.APC, locale.LRC_IQ, locale.MOH, locale.SHN, locale.SMA, locale.YI, locale.LA, locale.UND, locale.SMA_NO, locale.AZ_ARAB, locale.AB, locale.PAP_AW, locale.TA_MY, locale.SYR, locale.BO, locale.AA, locale.MIC, locale.RIF, locale.SAT, locale.EN_MV, locale.JBO, locale.ANN, locale.MZN, locale.TIG, locale.BGC, locale.SW_KE, locale.SYR_SY, locale.DV, locale.CO, locale.KAA_CYRL, locale.HA_GH, locale.CU, locale.GN, locale.WA, locale.KOK, locale.AA_DJ, locale.EN_DSRT, locale.HNJ_HMNP, locale.ARN, locale.CHO, locale.TRW, locale.KPE, locale.SID, locale.SKR, locale.BAL, locale.ZA, locale.GEZ_ER, locale.BSS, locale.SDH_IQ, locale.BLT, locale.MG, locale.EN_SHAW, locale.AZ_ARAB_TR, locale.AN, locale.LAG, locale.SCN, locale.SW, locale.WAL, locale.MN, locale.MNI_BENG, locale.CCH, locale.HNJ, locale.KAJ, locale.BEW, locale.VO, locale.CAD, locale.HA_ARAB, locale.MI, locale.SD_ARAB, locale.GEZ, locale.LTG, locale.LKT, locale.MGO, locale.MN_MONG, locale.RHG_ROHG_BD, locale.SMJ, locale.KK_ARAB, locale.MDF, locale.RAJ, locale.AZ_ARAB_IQ, locale.BO_IN, locale.SAT_DEVA, locale.ZH_LATN, locale.KAA, locale.CKB_IR, locale.HA_NE, locale.SD, locale.MNI_MTEI, locale.TA_SG, locale.SD_DEVA, locale.SSY, locale.HA, locale.IO:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 67 locales
	case locale.FF_LATN_CM, locale.RU_MD, locale.DYO, locale.HY, locale.RU_UA, locale.HU, locale.RU, locale.UK, locale.UZ_LATN, locale.UZ, locale.KK_CYRL, locale.KK, locale.PT_AO, locale.KA, locale.RU_BY, locale.LV, locale.DUA, locale.EWO, locale.PT_GQ, locale.CS, locale.FF_LATN_BF, locale.UZ_CYRL, locale.CV, locale.EN_SE, locale.FF_LATN, locale.PT_ST, locale.SZL, locale.PRG, locale.RU_KZ, locale.SAH, locale.FF_LATN_GH, locale.FF_LATN_GN, locale.FF_LATN_NE, locale.PT_MO, locale.BE_TARASK, locale.PT_GW, locale.FF_LATN_MR, locale.RU_KG, locale.TZM, locale.SK, locale.TT, locale.KSF, locale.YAV, locale.PT_MZ, locale.FF_LATN_LR, locale.TK, locale.PT_PT, locale.EN_FI, locale.PT_TL, locale.PL, locale.BE, locale.BG, locale.FF_LATN_GM, locale.KK_KZ, locale.KEA, locale.PT_LU, locale.SMN, locale.FF_LATN_GW, locale.PT_CV, locale.PT_CH, locale.FF_LATN_NG, locale.NMG, locale.EO, locale.TG, locale.BAS, locale.FF_LATN_SL, locale.FF:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 54 locales
	case locale.SR_LATN_XK, locale.LN, locale.LN_AO, locale.DE, locale.EL_POLYTON, locale.CA_IT, locale.DE_BE, locale.LB, locale.EN_DK, locale.LIJ, locale.VI, locale.LN_CG, locale.RO_MD, locale.EN_DE, locale.SR, locale.CA, locale.VMW, locale.AST, locale.EN_SI, locale.IS, locale.AZ_CYRL, locale.DA_GL, locale.CA_ES_VALENCIA, locale.HSB, locale.BS_LATN, locale.AZ, locale.CA_FR, locale.DE_LU, locale.EL_CY, locale.LLD, locale.BS_CYRL, locale.NDS_NL, locale.SR_CYRL_BA, locale.SR_CYRL, locale.GL, locale.KU, locale.SR_LATN_ME, locale.SR_CYRL_ME, locale.CA_AD, locale.EN_BE, locale.SR_LATN, locale.MK, locale.EL, locale.DSB, locale.DA, locale.DE_IT, locale.LN_CF, locale.SR_LATN_BA, locale.SR_CYRL_XK, locale.SC, locale.RO, locale.NDS, locale.AZ_LATN, locale.BS:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_BE, locale.FR_HT, locale.FR_GA, locale.FR_CH, locale.FR_CG, locale.FR_VU, locale.FR_RW, locale.FR_GN, locale.FR_NC, locale.FR_SY, locale.FR_BI, locale.FR_CF, locale.FR_TG, locale.FR_MF, locale.FR_PM, locale.FR_MR, locale.FR_BL, locale.FR_ML, locale.FR_TN, locale.FR_WF, locale.FR_CD, locale.FR_GQ, locale.FR_CI, locale.FR_BF, locale.FR_CM, locale.FR_KM, locale.FR_YT, locale.FR_GF, locale.FR_SN, locale.FR_PF, locale.FR_MG, locale.FR_MC, locale.FR_BJ, locale.FR_MQ, locale.FR, locale.FR_TD, locale.FR_NE, locale.FR_RE, locale.FR_MU, locale.FR_GP, locale.FR_DZ, locale.FR_DJ, locale.FR_SC:
		return CurrencyFormatInfo{
			Symbol:           "£E",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_KM, locale.AR_JO, locale.AR_SO, locale.AR_SS, locale.AR_KW, locale.AR_SA, locale.AR_EG, locale.AR_EH, locale.AR_IL, locale.AR_DJ, locale.AR_PS, locale.AR_OM, locale.AR_ER, locale.AR_IQ, locale.AR_AE, locale.AR_SD, locale.AR_YE, locale.AR_TD, locale.AR_SY, locale.AR_QA, locale.AR, locale.AR_BH:
		return CurrencyFormatInfo{
			Symbol:           "ج.م.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.PT, locale.ES_AR, locale.MGH, locale.MS_ARAB_BN, locale.YRL, locale.MS_BN, locale.QU_BO, locale.YRL_CO, locale.IA, locale.ES_CO, locale.JV, locale.ES_UY, locale.RW, locale.EN_AT, locale.KKJ, locale.NNH, locale.KGP, locale.JGO, locale.YRL_VE, locale.FUR, locale.SW_CD:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.DZ, locale.HI, locale.PA_GURU, locale.XNR, locale.KXV_LATN, locale.EN_IN, locale.KXV, locale.TE, locale.KXV_TELU, locale.HI_LATN, locale.TA, locale.TA_LK, locale.KXV_ORYA, locale.KXV_DEVA, locale.SA, locale.PA, locale.BN_IN, locale.GU, locale.KOK_LATN:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_LR, locale.FF_ADLM_CM, locale.FF_ADLM, locale.FF_ADLM_SL, locale.FF_ADLM_GM, locale.FF_ADLM_GH, locale.FF_ADLM_NE, locale.FF_ADLM_NG, locale.FF_ADLM_BF, locale.FF_ADLM_MR, locale.FF_ADLM_SN, locale.FF_ADLM_GW:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.NL, locale.EN_NL, locale.NL_BQ, locale.NL_BE, locale.ES_PY, locale.NL_SR, locale.NL_AW, locale.NL_SX, locale.NL_CW:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.TR, locale.MUA, locale.ID, locale.SU_LATN, locale.TR_CY, locale.ES_BO, locale.EN_ID, locale.SU, locale.MS_ID:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.AGQ, locale.ZGH, locale.SHI, locale.KAB, locale.SHI_TFNG, locale.SHI_LATN, locale.OC, locale.OC_ES:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.ES_CR, locale.SS_SZ, locale.AF, locale.VE, locale.EN_ZA, locale.AF_NA, locale.NR, locale.SS:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KM, locale.RWK, locale.LUO, locale.LG, locale.SBP, locale.KSB, locale.BEZ:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_LY, locale.AR_DZ, locale.AR_TN, locale.AR_MR, locale.AR_LB, locale.AR_MA:
		return CurrencyFormatInfo{
			Symbol:           "ج.م.\u200f",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.EN_150, locale.ASA, locale.XOG, locale.TPI, locale.CE, locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.EU, locale.HR, locale.SL, locale.FO_DK, locale.HR_BA, locale.FO:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.LO, locale.KL, locale.ES_EC, locale.ES_VE, locale.ES_CL, locale.SG:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.LT, locale.KSH, locale.SE, locale.ET, locale.SE_SE, locale.SE_FI:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.BGN_IR, locale.BGN_AF, locale.BGN, locale.BGN_OM, locale.BGN_AE:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.IT_VA, locale.FR_MA, locale.IT, locale.FR_LU, locale.IT_SM:
		return CurrencyFormatInfo{
			Symbol:           "£E",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.ES_IC, locale.ES_EA, locale.ES_PH, locale.ES:
		return CurrencyFormatInfo{
			Symbol:           "EGP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW_LI, locale.RM, locale.GSW, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.KHQ, locale.DJE, locale.TWQ, locale.SES:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NN, locale.NB_SJ, locale.NB, locale.NO:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS_RU, locale.DE_AT, locale.TS, locale.OS:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.AS, locale.NE, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.LU, locale.SEH, locale.RN:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SQ, locale.SQ_MK, locale.SQ_XK:
		return CurrencyFormatInfo{
			Symbol:           "EGP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SV_FI, locale.SV, locale.SV_AX:
		return CurrencyFormatInfo{
			Symbol:           "EG£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.EN_CH:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.MFE, locale.NSO:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN, locale.TN_BW:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "MSRP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BR:
		return CurrencyFormatInfo{
			Symbol:           "£ E",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "ई.जि.पि",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AU:
		return CurrencyFormatInfo{
			Symbol:           "£",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_GQ:
		return CurrencyFormatInfo{
			Symbol:           "EGP",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FI:
		return CurrencyFormatInfo{
			Symbol:           "EGP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	case locale.FR_CA:
		return CurrencyFormatInfo{
			Symbol:           "£E",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "£E",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.KY:
		return CurrencyFormatInfo{
			Symbol:           "LE",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߡߛߔ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.WO:
		return CurrencyFormatInfo{
			Symbol:           "EGPP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "E£",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
