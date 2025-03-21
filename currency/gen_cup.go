// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 46.1.0
// ISO 4217 Data Last Updated: Sun, 16 Feb 2025 04:00:34 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = CUP{}

// Cuban Peso currency type
type CUP struct{}

func (CUP) Code() string { return "CUP" }

func (CUP) Number() string { return "192" }

func (CUP) Name() string { return "Cuban Peso" }

func (CUP) MinorUnits() int { return 2 }

func (CUP) FormatInfo(localeType locale.Locale) CurrencyFormatInfo {
	switch localeType {
	// Group of 193 locales
	case locale.TH, locale.EN_GM, locale.SN, locale.VAI_VAII, locale.EN_CC, locale.MN_MONG_MN, locale.EN_001, locale.KDE, locale.ST, locale.EN_MU, locale.ES_HN, locale.YUE_HANS, locale.CGG, locale.EN_CA, locale.EN_LC, locale.EN_NA, locale.EN_TZ, locale.MT, locale.EN_DM, locale.EN_GY, locale.EN_MS, locale.EN_SB, locale.ROF, locale.VAI_LATN, locale.EN_FM, locale.EN_GD, locale.EN_KI, locale.MS_ARAB, locale.GV, locale.EN_JM, locale.AK, locale.EN_VU, locale.JA, locale.KW, locale.ND, locale.EN_MP, locale.GUZ, locale.EN_GB, locale.ES_SV, locale.GA_GB, locale.YO, locale.EN_AG, locale.ES_MX, locale.UG, locale.EN_GH, locale.EN_IM, locale.EN_VG, locale.ZH_HANS_MO, locale.EN_KY, locale.EN_TT, locale.EN_UM, locale.ML, locale.TEO_KE, locale.ZH_HANT_MO, locale.OM_KE, locale.ES_GT, locale.KO_KP, locale.VUN, locale.BHO, locale.EN_SH, locale.EN_VI, locale.ES_DO, locale.VAI, locale.ZH_HANT_MY, locale.EN_NZ, locale.MS, locale.SO_KE, locale.EN_MW, locale.ZU, locale.KI, locale.MS_SG, locale.SI, locale.EN_GG, locale.YUE_HANT, locale.ES_BZ, locale.CHR, locale.MR, locale.ZH_HANS_MY, locale.NAQ, locale.OR, locale.EN_SZ, locale.SO_ET, locale.EN_BI, locale.OM, locale.ZH_HANT, locale.SO, locale.YO_BJ, locale.CEB, locale.EN_NR, locale.TEO, locale.GA, locale.ZH_HANT_HK, locale.ES_PR, locale.KN, locale.EN_HK, locale.EN_PK, locale.EN_FK, locale.EN_PG, locale.ZH_HANS_SG, locale.EN_FJ, locale.EN_TK, locale.EN_BS, locale.EN_PW, locale.EN_BB, locale.EN_BW, locale.EN_DG, locale.EN_VC, locale.DOI, locale.EN_MT, locale.EN_TO, locale.KAM, locale.EN_KE, locale.GD, locale.EN_MY, locale.EN_LR, locale.ST_LS, locale.EN_CK, locale.EN_UG, locale.ES_US, locale.BEM, locale.EN_SG, locale.ZH, locale.DAV, locale.AM, locale.EN_AI, locale.BM, locale.JMC, locale.EE_TG, locale.EN, locale.EN_CM, locale.EN_ER, locale.EN_LS, locale.EN_PH, locale.EN_TC, locale.HAW, locale.EN_PR, locale.ES_NI, locale.EN_ZM, locale.ES_PA, locale.TI, locale.TI_ER, locale.EN_NF, locale.SO_DJ, locale.EN_IE, locale.ES_419, locale.CY, locale.EN_MO, locale.EN_SD, locale.EN_MH, locale.ZH_HANS, locale.EN_GI, locale.EN_NU, locale.YUE_HANT_CN, locale.EE, locale.EN_JE, locale.NUS, locale.EN_IL, locale.SAQ, locale.ZH_HANS_HK, locale.EN_ZW, locale.EN_WS, locale.FIL, locale.YUE, locale.EN_SS, locale.ES_BR, locale.PCM, locale.EN_NG, locale.EN_BZ, locale.EN_CX, locale.IG, locale.EN_KN, locale.EN_GU, locale.EN_MG, locale.EN_SX, locale.MER, locale.EN_AE, locale.EN_PN, locale.EN_SC, locale.EN_TV, locale.EN_AS, locale.KLN, locale.MAS_TZ, locale.EBU, locale.EN_RW, locale.KO, locale.EN_BM, locale.EN_CY, locale.EN_SL, locale.KO_CN, locale.NYN, locale.MAS, locale.EN_IO, locale.ES_CU:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 140 locales
	case locale.MN_MONG, locale.WAL, locale.ZH_LATN, locale.CSW, locale.KOK, locale.LKT, locale.MUS, locale.GAA, locale.CHO, locale.AB, locale.AN, locale.APC, locale.CIC, locale.EN_SHAW, locale.RAJ, locale.QUC, locale.ANN, locale.BO, locale.SMA_NO, locale.EN_MV, locale.MG, locale.BLT, locale.SID, locale.UND, locale.WBP, locale.HA, locale.KS_DEVA, locale.SSY, locale.KPE, locale.MI, locale.LTG, locale.QU, locale.HNJ, locale.WA, locale.SD_ARAB, locale.ARN, locale.TA_MY, locale.CCH, locale.LA, locale.MHN, locale.SMA, locale.BAL, locale.TO, locale.MGO, locale.CU, locale.MNI_BENG, locale.EN_DSRT, locale.MZN, locale.NV, locale.HNJ_HMNP, locale.PAP_AW, locale.CKB, locale.SAT, locale.SHN, locale.SMJ, locale.KAA_LATN, locale.KAJ, locale.SW, locale.BYN, locale.FRR, locale.CAD, locale.CO, locale.AA_DJ, locale.SD_DEVA, locale.SYR, locale.TA_SG, locale.HA_NE, locale.MIC, locale.KOK_DEVA, locale.GEZ_ER, locale.BO_IN, locale.TIG, locale.JBO, locale.MAI, locale.SAT_DEVA, locale.BSS, locale.TRV, locale.MDF, locale.CKB_IR, locale.PIS, locale.BGC, locale.DV, locale.KAA_CYRL, locale.YI, locale.ES_PE, locale.SW_UG, locale.AZ_ARAB, locale.LRC, locale.MOH, locale.RIF, locale.ZA, locale.HA_GH, locale.SDH, locale.BAL_ARAB, locale.HA_ARAB_SD, locale.TRW, locale.BM_NKOO, locale.KEN, locale.KK_ARAB, locale.MNI, locale.OSA, locale.SAT_OLCK, locale.II, locale.SD, locale.SYR_SY, locale.MN, locale.RHG_ROHG_BD, locale.IO, locale.GN, locale.RHG, locale.KPE_GN, locale.RHG_ROHG, locale.BA, locale.PAP, locale.VO, locale.AA_ER, locale.HA_ARAB, locale.SDH_IQ, locale.TYV, locale.MYV, locale.IU_LATN, locale.AZ_ARAB_TR, locale.SKR, locale.SHN_TH, locale.LAG, locale.KCG, locale.NY, locale.AA, locale.QU_EC, locale.SCN, locale.AZ_ARAB_IQ, locale.MNI_MTEI, locale.LRC_IQ, locale.SMJ_NO, locale.SMS, locale.SW_KE, locale.BEW, locale.GEZ, locale.KAA, locale.IU:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 67 locales
	case locale.RU_UA, locale.UZ_LATN, locale.SK, locale.CV, locale.RU_BY, locale.TZM, locale.FF_LATN_CM, locale.RU_KZ, locale.TK, locale.BR, locale.PT_MO, locale.TG, locale.KY, locale.PT_PT, locale.PT_ST, locale.PT_GQ, locale.UK, locale.CS, locale.FF_LATN_SL, locale.PT_GW, locale.PT_LU, locale.EWO, locale.RU_KG, locale.SAH, locale.PL, locale.KEA, locale.FF_LATN_GH, locale.UZ, locale.PT_AO, locale.KSF, locale.EN_FI, locale.FF_LATN_GN, locale.KK_CYRL, locale.FF_LATN_BF, locale.SZL, locale.PT_CH, locale.BAS, locale.KK_KZ, locale.NMG, locale.HU, locale.TT, locale.KK, locale.PT_MZ, locale.PT_CV, locale.KA, locale.YAV, locale.FF_LATN_GW, locale.RU_MD, locale.FF_LATN_NE, locale.FR_CA, locale.EN_SE, locale.UZ_CYRL, locale.FF_LATN_NG, locale.LV, locale.RU, locale.PT_TL, locale.EO, locale.DUA, locale.FF_LATN_GM, locale.HY, locale.FF_LATN_MR, locale.FF_LATN, locale.DYO, locale.SMN, locale.FF, locale.FF_LATN_LR, locale.PRG:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 63 locales
	case locale.MK, locale.NDS, locale.DE, locale.AZ_CYRL, locale.IT_SM, locale.GL, locale.LN_CG, locale.SR_CYRL_BA, locale.RO, locale.LN, locale.SR_CYRL, locale.AZ_LATN, locale.HSB, locale.LN_AO, locale.EL, locale.DA, locale.SR_CYRL_ME, locale.CA_FR, locale.SC, locale.AST, locale.RO_MD, locale.NDS_NL, locale.DE_BE, locale.EN_SI, locale.SR_LATN_XK, locale.LIJ, locale.DSB, locale.ES_IC, locale.DE_LU, locale.SR_LATN_ME, locale.FR_LU, locale.LN_CF, locale.AZ, locale.EL_POLYTON, locale.ES_PH, locale.IT_VA, locale.CA, locale.EL_CY, locale.EN_DE, locale.EN_BE, locale.EN_DK, locale.ES_EA, locale.IS, locale.BS_LATN, locale.CA_AD, locale.LLD, locale.SR_LATN, locale.CA_IT, locale.FR_MA, locale.ES, locale.CA_ES_VALENCIA, locale.SR, locale.DA_GL, locale.IT, locale.BS, locale.VI, locale.SR_CYRL_XK, locale.VMW, locale.BS_CYRL, locale.SR_LATN_BA, locale.KU, locale.LB, locale.DE_IT:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_BI, locale.FR, locale.FR_CD, locale.FR_BL, locale.FR_VU, locale.FR_BJ, locale.FR_MC, locale.FR_GP, locale.FR_CM, locale.FR_BE, locale.FR_ML, locale.FR_GQ, locale.FR_NC, locale.FR_MF, locale.FR_SN, locale.FR_SY, locale.FR_MQ, locale.FR_GF, locale.FR_RW, locale.FR_PM, locale.FR_CG, locale.FR_CI, locale.FR_KM, locale.FR_TD, locale.FR_SC, locale.FR_MG, locale.FR_PF, locale.FR_GN, locale.FR_TG, locale.FR_CH, locale.FR_TN, locale.FR_WF, locale.FR_DJ, locale.FR_MR, locale.FR_BF, locale.FR_DZ, locale.FR_NE, locale.FR_CF, locale.FR_MU, locale.FR_GA, locale.FR_RE, locale.FR_HT, locale.FR_YT:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_PS, locale.AR_SY, locale.AR_KM, locale.AR_SS, locale.AR_OM, locale.AR_EH, locale.AR_EG, locale.AR_SO, locale.AR_SA, locale.AR_YE, locale.AR_TD, locale.AR_IL, locale.AR_SD, locale.AR, locale.AR_IQ, locale.AR_QA, locale.AR_AE, locale.AR_DJ, locale.AR_KW, locale.AR_BH, locale.AR_JO, locale.AR_ER:
		return CurrencyFormatInfo{
			Symbol:           "CU$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 22 locales
	case locale.JGO, locale.ES_AR, locale.ES_CO, locale.IA, locale.KKJ, locale.PT, locale.MS_BN, locale.RW, locale.FUR, locale.QU_BO, locale.KGP, locale.MGH, locale.WO, locale.MS_ARAB_BN, locale.YRL_VE, locale.YRL, locale.JV, locale.ES_UY, locale.NNH, locale.EN_AT, locale.YRL_CO, locale.SW_CD:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 19 locales
	case locale.KXV_DEVA, locale.HI, locale.PA, locale.KXV_LATN, locale.TA_LK, locale.DZ, locale.GU, locale.TE, locale.KXV_TELU, locale.HI_LATN, locale.KXV_ORYA, locale.EN_IN, locale.KOK_LATN, locale.TA, locale.BN_IN, locale.KXV, locale.SA, locale.PA_GURU, locale.XNR:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_BF, locale.FF_ADLM_LR, locale.FF_ADLM_NE, locale.FF_ADLM_SN, locale.FF_ADLM_NG, locale.FF_ADLM_GM, locale.FF_ADLM_CM, locale.FF_ADLM_GH, locale.FF_ADLM_SL, locale.FF_ADLM_MR, locale.FF_ADLM_GW, locale.FF_ADLM:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.ES_BO, locale.SU_LATN, locale.ES_GQ, locale.EN_ID, locale.TR, locale.SU, locale.MUA, locale.ID, locale.MS_ID, locale.TR_CY:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 9 locales
	case locale.ET, locale.SE, locale.SE_SE, locale.LT, locale.SE_FI, locale.SV_FI, locale.SV_AX, locale.KSH, locale.SV:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 9 locales
	case locale.NL_SX, locale.NL_SR, locale.ES_PY, locale.NL_AW, locale.EN_NL, locale.NL, locale.NL_BQ, locale.NL_BE, locale.NL_CW:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.EN_ZA, locale.ES_CR, locale.SS, locale.AF_NA, locale.AF, locale.VE, locale.NR, locale.SS_SZ:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI_TFNG, locale.SHI, locale.OC_ES, locale.AGQ, locale.ZGH, locale.OC, locale.KAB, locale.SHI_LATN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.BEZ, locale.RWK, locale.SBP, locale.KM, locale.LUO, locale.KSB, locale.LG:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_TN, locale.AR_LY, locale.AR_DZ, locale.AR_MA, locale.AR_MR, locale.AR_LB:
		return CurrencyFormatInfo{
			Symbol:           "CU$",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.ES_CL, locale.ES_VE, locale.LO, locale.ES_EC, locale.KL, locale.SG:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.EU, locale.FO_DK, locale.HR, locale.HR_BA, locale.SL, locale.FO:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.XOG, locale.EN_150, locale.ASA, locale.MY, locale.CE, locale.TPI:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_OM, locale.BGN, locale.BGN_AF, locale.BGN_IR, locale.BGN_AE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.NN, locale.NB, locale.NO, locale.NB_SJ:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.OS, locale.DE_AT, locale.TS, locale.OS_RU:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.RM, locale.GSW_LI, locale.GSW_FR, locale.GSW:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.SES, locale.DJE, locale.KHQ, locale.TWQ:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.SQ_XK, locale.SQ, locale.SQ_MK, locale.BG:
		return CurrencyFormatInfo{
			Symbol:           "CUP",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP_IN, locale.BN, locale.CCP:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.EN_CH, locale.IT_CH, locale.DE_CH:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.NE_IN, locale.AS, locale.NE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.LU, locale.RN:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.BE, locale.BE_TARASK:
		return CurrencyFormatInfo{
			Symbol:           "$MN",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.IE, locale.BLO:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return CurrencyFormatInfo{
			Symbol:           "$",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
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
	case locale.BAL_LATN:
		return CurrencyFormatInfo{
			Symbol:           "CBAP",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return CurrencyFormatInfo{
			Symbol:           "सि.इउ.पि",
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
	case locale.EN_AU:
		return CurrencyFormatInfo{
			Symbol:           "₱",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
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
			Symbol:           "CUP",
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
			Symbol:           "ߞߎ߳ߔ",
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
