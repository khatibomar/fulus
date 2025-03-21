// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = TWD{}

// New Taiwan Dollar currency type
type TWD struct{}

func (TWD) Code() string { return "TWD" }

func (TWD) Number() string { return "901" }

func (TWD) Name() string { return "New Taiwan Dollar" }

func (TWD) MinorUnits() int { return 2 }

func (TWD) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 178 locales
	case locale.ZH_HANT_MO, locale.OM, locale.YUE_HANT_CN, locale.EN_LR, locale.EN_TZ, locale.KN, locale.EN_IL, locale.EN_VU, locale.MN_MONG_MN, locale.NUS, locale.EN_NZ, locale.CEB, locale.YUE_HANT, locale.EN_PN, locale.EN_PK, locale.MS_SG, locale.EN_PH, locale.ROF, locale.EN_SG, locale.SO_ET, locale.TH, locale.EE, locale.YUE_HANS, locale.ZH_HANS_MY, locale.EN_MG, locale.MR, locale.EN_ER, locale.EN_BM, locale.EN_HK, locale.ZH, locale.EN_DM, locale.ZH_HANS_HK, locale.BEM, locale.EN_JE, locale.EN_PW, locale.EN_AG, locale.FIL, locale.EN_SB, locale.EN_MU, locale.BHO, locale.MS, locale.EN_JM, locale.EN_BW, locale.EN_IE, locale.EN_PG, locale.PCM, locale.EN_GU, locale.EN_MP, locale.TEO, locale.GV, locale.EN_BB, locale.EN_UM, locale.KLN, locale.MS_ARAB, locale.EN_NA, locale.EN_001, locale.EN_GY, locale.DAV, locale.EN_VC, locale.EN_CM, locale.EN_DG, locale.EN_KI, locale.KO, locale.EN_MO, locale.VAI, locale.CHR, locale.EN_AS, locale.EN_FJ, locale.EN_GD, locale.AM, locale.EN_BI, locale.EN_CC, locale.EN_GH, locale.YO_BJ, locale.EN_SZ, locale.EN_TK, locale.KDE, locale.KI, locale.EE_TG, locale.ND, locale.EN_CK, locale.EN_CX, locale.EN_SD, locale.SO, locale.MAS, locale.EN_CY, locale.EN_GI, locale.EN_FK, locale.KO_CN, locale.EN_SL, locale.EN_SX, locale.KAM, locale.MER, locale.YUE, locale.EN_BZ, locale.EN_TV, locale.EN_FM, locale.EN_CA, locale.EN_GG, locale.EN_IM, locale.EN_GM, locale.EN_NR, locale.CGG, locale.CY, locale.GA, locale.GD, locale.ZH_HANS, locale.ZH_HANS_SG, locale.EN_KE, locale.EN_LC, locale.EN_SS, locale.EN_MH, locale.VUN, locale.EN_NG, locale.IG, locale.EN_GB, locale.EN_ZM, locale.EN_TC, locale.NAQ, locale.GUZ, locale.JMC, locale.KO_KP, locale.VAI_LATN, locale.VAI_VAII, locale.EN_MT, locale.EN, locale.EN_TO, locale.EN_MW, locale.EN_SH, locale.SI, locale.ZH_HANT_HK, locale.EN_BS, locale.EN_NF, locale.ZH_HANS_MO, locale.UG, locale.EN_KN, locale.EN_MY, locale.KW, locale.EN_AI, locale.SAQ, locale.EN_MS, locale.EN_RW, locale.AK, locale.EN_IO, locale.EN_NU, locale.HAW, locale.DOI, locale.MT, locale.OM_KE, locale.OR, locale.EN_SC, locale.ST, locale.ST_LS, locale.TI, locale.ZU, locale.EN_VI, locale.SO_KE, locale.GA_GB, locale.ML, locale.EN_TT, locale.SO_DJ, locale.BM, locale.EN_KY, locale.EN_LS, locale.EN_UG, locale.EBU, locale.EN_ZW, locale.SN, locale.TEO_KE, locale.EN_PR, locale.EN_WS, locale.EN_VG, locale.YO, locale.JA, locale.NYN, locale.TI_ER, locale.EN_AE, locale.MAS_TZ:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 139 locales
	case locale.SSY, locale.LKT, locale.WBP, locale.MDF, locale.QU, locale.SDH_IQ, locale.EN_DSRT, locale.QUC, locale.NY, locale.SD_DEVA, locale.YI, locale.IO, locale.SAT_OLCK, locale.MAI, locale.BSS, locale.KPE, locale.HA_NE, locale.LRC, locale.KEN, locale.MNI, locale.AA_ER, locale.SD_ARAB, locale.SW_UG, locale.LA, locale.EN_SHAW, locale.JBO, locale.SYR, locale.ANN, locale.RIF, locale.ARN, locale.SHN, locale.APC, locale.BEW, locale.GEZ_ER, locale.KAA_LATN, locale.HNJ, locale.SMA, locale.BGC, locale.MUS, locale.KOK, locale.SMA_NO, locale.TYV, locale.AB, locale.SAT, locale.IU, locale.CIC, locale.SCN, locale.TA_SG, locale.KAA, locale.SAT_DEVA, locale.AZ_ARAB, locale.BYN, locale.MYV, locale.CO, locale.PAP_AW, locale.KPE_GN, locale.TIG, locale.GN, locale.CAD, locale.SD, locale.SKR, locale.TRW, locale.CHO, locale.MG, locale.BO_IN, locale.EN_MV, locale.KAA_CYRL, locale.KAJ, locale.QU_EC, locale.RHG_ROHG_BD, locale.SMJ_NO, locale.BAL_ARAB, locale.UND, locale.SW, locale.BAL, locale.DV, locale.OSA, locale.HA_GH, locale.AZ_ARAB_TR, locale.KK_ARAB, locale.TA_MY, locale.LRC_IQ, locale.TRV, locale.II, locale.RHG_ROHG, locale.WA, locale.PIS, locale.LTG, locale.MNI_BENG, locale.SYR_SY, locale.IU_LATN, locale.CKB, locale.MOH, locale.CU, locale.HNJ_HMNP, locale.WAL, locale.AA_DJ, locale.BA, locale.SHN_TH, locale.SMS, locale.SW_KE, locale.AN, locale.GEZ, locale.MN_MONG, locale.PAP, locale.TO, locale.CKB_IR, locale.HA_ARAB, locale.CSW, locale.MGO, locale.BM_NKOO, locale.LAG, locale.MZN, locale.CCH, locale.BLT, locale.BO, locale.MNI_MTEI, locale.RAJ, locale.AA, locale.GAA, locale.KS_DEVA, locale.SDH, locale.MN, locale.SMJ, locale.KOK_DEVA, locale.RHG, locale.VO, locale.FRR, locale.MIC, locale.ZH_LATN, locale.KCG, locale.MI, locale.HA_ARAB_SD, locale.HA, locale.MHN, locale.SID, locale.AZ_ARAB_IQ, locale.NV, locale.ZA:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 65 locales
	case locale.KSF, locale.FF, locale.KEA, locale.FF_LATN_NE, locale.UZ_CYRL, locale.BAS, locale.SQ_XK, locale.SQ_MK, locale.PT_ST, locale.PT_PT, locale.SMN, locale.SQ, locale.CV, locale.PT_MO, locale.UZ_LATN, locale.TG, locale.FF_LATN_GN, locale.SAH, locale.RU_BY, locale.KK, locale.DYO, locale.FF_LATN_LR, locale.PT_TL, locale.FF_LATN_SL, locale.CS, locale.PT_AO, locale.PT_CV, locale.YAV, locale.LV, locale.FF_LATN, locale.HY, locale.BE, locale.FF_LATN_BF, locale.TK, locale.PT_GQ, locale.EN_SE, locale.EO, locale.FF_LATN_CM, locale.RU_MD, locale.KK_KZ, locale.NMG, locale.UZ, locale.KA, locale.BE_TARASK, locale.SZL, locale.PT_MZ, locale.DUA, locale.PRG, locale.PT_LU, locale.RU, locale.FF_LATN_MR, locale.KK_CYRL, locale.EN_FI, locale.PT_GW, locale.EWO, locale.PT_CH, locale.RU_KZ, locale.FF_LATN_GH, locale.FF_LATN_GW, locale.TZM, locale.FF_LATN_NG, locale.RU_UA, locale.FF_LATN_GM, locale.RU_KG, locale.TT:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 50 locales
	case locale.EN_SI, locale.DSB, locale.SC, locale.CA_ES_VALENCIA, locale.LIJ, locale.DE_LU, locale.AZ_CYRL, locale.EL_CY, locale.AST, locale.CA_IT, locale.LB, locale.DE_IT, locale.SR_CYRL_BA, locale.SR_CYRL_ME, locale.SR_LATN_ME, locale.SR_CYRL_XK, locale.NDS_NL, locale.CA, locale.KU, locale.EN_DE, locale.AZ, locale.LLD, locale.VI, locale.LN, locale.EL_POLYTON, locale.EN_DK, locale.CA_AD, locale.SR_LATN_BA, locale.EL, locale.SR, locale.LN_AO, locale.SR_CYRL, locale.SR_LATN_XK, locale.SR_LATN, locale.BS_LATN, locale.CA_FR, locale.DA_GL, locale.DE, locale.LN_CF, locale.NDS, locale.BS_CYRL, locale.DE_BE, locale.EN_BE, locale.HSB, locale.VMW, locale.GL, locale.LN_CG, locale.DA, locale.AZ_LATN, locale.BS:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_CM, locale.FR_TD, locale.FR_BE, locale.FR_BJ, locale.FR_PF, locale.FR_CI, locale.FR_MR, locale.FR_YT, locale.FR_BI, locale.FR_GN, locale.FR_MQ, locale.FR_SC, locale.FR_CG, locale.FR_BF, locale.FR_MC, locale.FR_GP, locale.FR_BL, locale.FR_NC, locale.FR_MF, locale.FR_ML, locale.FR_DJ, locale.FR_RE, locale.FR_CD, locale.FR_TN, locale.FR_NE, locale.FR_VU, locale.FR_KM, locale.FR_HT, locale.FR_MG, locale.FR_MU, locale.FR_GF, locale.FR_PM, locale.FR_GA, locale.FR_RW, locale.FR_TG, locale.FR_GQ, locale.FR_WF, locale.FR_CF, locale.FR_DZ, locale.FR_CH, locale.FR, locale.FR_SY, locale.FR_SN:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_BH, locale.AR_JO, locale.AR_OM, locale.AR_KW, locale.AR_AE, locale.AR_YE, locale.AR_SA, locale.AR_TD, locale.AR_SD, locale.AR_IL, locale.AR_ER, locale.AR_SY, locale.AR_DJ, locale.AR_EH, locale.AR_QA, locale.AR_SO, locale.AR, locale.AR_KM, locale.AR_PS, locale.AR_IQ, locale.AR_SS, locale.AR_EG:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 19 locales
	case locale.KKJ, locale.RW, locale.KGP, locale.SW_CD, locale.MS_BN, locale.QU_BO, locale.YRL_VE, locale.FUR, locale.JGO, locale.MGH, locale.YRL_CO, locale.MS_ARAB_BN, locale.WO, locale.IA, locale.YRL, locale.NNH, locale.EN_AT, locale.JV, locale.PT:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.KXV, locale.KXV_DEVA, locale.TA, locale.GU, locale.PA_GURU, locale.EN_IN, locale.KOK_LATN, locale.KXV_LATN, locale.BN_IN, locale.PA, locale.SA, locale.KXV_TELU, locale.TE, locale.DZ, locale.TA_LK, locale.XNR, locale.KXV_ORYA, locale.HI_LATN, locale.HI:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 14 locales
	case locale.ES_DO, locale.ES_MX, locale.ES_CU, locale.ES_PR, locale.ES_PA, locale.ES_NI, locale.ES_GT, locale.ES_SV, locale.ES_419, locale.ES_BR, locale.EN_AU, locale.ES_HN, locale.ES_US, locale.ES_BZ:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 13 locales
	case locale.FR_LU, locale.ES_IC, locale.IT_VA, locale.IT_SM, locale.IT, locale.FR_MA, locale.RO_MD, locale.RO, locale.ES_PH, locale.ES_EA, locale.IS, locale.MK, locale.ES:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_GW, locale.FF_ADLM_CM, locale.FF_ADLM_GM, locale.FF_ADLM_NG, locale.FF_ADLM_SL, locale.FF_ADLM_LR, locale.FF_ADLM_NE, locale.FF_ADLM, locale.FF_ADLM_BF, locale.FF_ADLM_SN, locale.FF_ADLM_MR, locale.FF_ADLM_GH:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.BR, locale.FR_CA, locale.PL, locale.HU, locale.UK, locale.BG, locale.KY, locale.SK:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_CW, locale.EN_NL, locale.NL, locale.NL_BE, locale.NL_SX, locale.NL_SR, locale.NL_AW, locale.NL_BQ:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI_TFNG, locale.ZGH, locale.OC_ES, locale.SHI, locale.AGQ, locale.KAB, locale.OC, locale.SHI_LATN:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SU_LATN, locale.MUA, locale.SU, locale.MS_ID, locale.TR, locale.EN_ID, locale.TR_CY, locale.ID:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.AF_NA, locale.VE, locale.NR, locale.SS_SZ, locale.AF, locale.SS, locale.EN_ZA:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.LUO, locale.KM, locale.SBP, locale.RWK, locale.BEZ, locale.LG, locale.KSB:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_TN, locale.AR_LB, locale.AR_LY, locale.AR_MA, locale.AR_DZ, locale.AR_MR:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.EN_150, locale.ASA, locale.TPI, locale.XOG, locale.CE, locale.MY:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_AE, locale.BGN, locale.BGN_IR, locale.BGN_AF, locale.BGN_OM:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.LT, locale.SV_FI, locale.SV_AX, locale.FI, locale.SV:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 5 locales
	case locale.SE_FI, locale.ET, locale.SE, locale.SE_SE, locale.KSH:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_LI, locale.RM, locale.GSW_FR:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.KHQ, locale.SES, locale.TWQ, locale.DJE:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NB_SJ, locale.NB, locale.NO, locale.NN:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS, locale.DE_AT, locale.OS_RU, locale.TS:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.BN, locale.CCP, locale.CCP_IN:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_CO, locale.ES_UY, locale.ES_AR:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.ES_EC, locale.ES_VE, locale.ES_CL:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.FO, locale.FO_DK, locale.EU:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.LO, locale.KL, locale.SG:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE, locale.AS, locale.NE_IN:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.RN, locale.SEH, locale.LU:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SL, locale.HR, locale.HR_BA:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS_PK, locale.PS:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.EN_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ES_GQ, locale.ES_BO:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR_IN, locale.UR:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 2 locales
	case locale.WAE, locale.LMO:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.ZH_HANT, locale.ZH_HANT_MY:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "NTWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "एन.ति$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_CR:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.ES_PE:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.ES_PY:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.IT_CH:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.LUY:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return CurrencyFormatInfo{
			Symbol:           "ߕߥߘ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return CurrencyFormatInfo{
			Symbol:           "TWD",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return CurrencyFormatInfo{
			Symbol:           "NT$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
