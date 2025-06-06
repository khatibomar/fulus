// Code generated by generator.go; DO NOT EDIT.

// CLDR Version: 47.0.0
// ISO 4217 Data Last Updated: Sat, 31 May 2025 03:00:43 GMT

package currency

import "github.com/khatibomar/fulus/locale"

var _ Currency = KGS{}

// Som currency type
type KGS struct{}

func (KGS) Code() string { return "KGS" }

func (KGS) Number() string { return "417" }

func (KGS) Name() string { return "Som" }

func (KGS) MinorUnits() int { return 2 }

func (KGS) FormatInfo(localeType locale.Locale) FormatInfo {
	switch localeType {
	// Group of 140 locales
	case locale.CHO, locale.HA_ARAB_SD, locale.QU_EC, locale.QU, locale.LA, locale.MNI_MTEI, locale.SAT_OLCK, locale.COP, locale.QUC, locale.AA_DJ, locale.SDH_IQ, locale.SMJ_NO, locale.MHN, locale.NV, locale.CIC, locale.HA, locale.KS_DEVA, locale.LTG, locale.KOK, locale.PIS, locale.BLT, locale.SD_ARAB, locale.SMS, locale.KCG, locale.BM_NKOO, locale.MNI, locale.HNJ_HMNP, locale.MYV, locale.KAJ, locale.LRC_IQ, locale.MDF, locale.MNI_BENG, locale.KPE, locale.ES_PE, locale.SMJ, locale.BEW, locale.GAA, locale.LAG, locale.WA, locale.AZ_ARAB, locale.TIG, locale.MN, locale.TA_MY, locale.IU, locale.NY, locale.TA_SG, locale.SW_KE, locale.ANN, locale.MUS, locale.AB, locale.MIC, locale.BYN, locale.RAJ, locale.SMA_NO, locale.SSY, locale.SKR, locale.TYV, locale.AA_ER, locale.KK_ARAB, locale.YI, locale.MZN, locale.SYR_SY, locale.CAD, locale.CKB, locale.BAL_ARAB, locale.HA_NE, locale.FRR, locale.KAA_LATN, locale.ARN, locale.CCH, locale.MG, locale.MAI, locale.TRV, locale.SDH, locale.VO, locale.KPE_GN, locale.KEN, locale.GN, locale.MGO, locale.AA, locale.SW, locale.CSW, locale.MN_MONG, locale.SAT, locale.SD_DEVA, locale.APC, locale.HA_GH, locale.GEZ_ER, locale.SHN_TH, locale.SID, locale.EN_DSRT, locale.ZA, locale.HNJ, locale.KAA_CYRL, locale.SD, locale.AN, locale.ZH_LATN, locale.OSA, locale.BA, locale.PAP, locale.MOH, locale.II, locale.PAP_AW, locale.IU_LATN, locale.RHG_ROHG_BD, locale.LRC, locale.SMA, locale.IO, locale.SCN, locale.SHN, locale.BSS, locale.HA_ARAB, locale.MI, locale.EN_SHAW, locale.SW_UG, locale.CU, locale.BO_IN, locale.TRW, locale.BO, locale.CO, locale.GEZ, locale.CKB_IR, locale.RHG_ROHG, locale.JBO, locale.SAT_DEVA, locale.RIF, locale.KOK_DEVA, locale.UND, locale.KAA, locale.DV, locale.SYR, locale.AZ_ARAB_TR, locale.BGC, locale.WAL, locale.LKT, locale.RHG, locale.WBP, locale.AZ_ARAB_IQ, locale.BAL, locale.TO:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 102 locales
	case locale.FIL, locale.BM, locale.JMC, locale.NUS, locale.OM, locale.AM, locale.NAQ, locale.GA_GB, locale.YO_BJ, locale.ES_PA, locale.SO_ET, locale.ZH_HANS_SG, locale.EE, locale.MT, locale.ZH_HANT, locale.KLN, locale.YUE_HANT, locale.BEM, locale.TEO, locale.CHR, locale.SAQ, locale.ZH_HANS_HK, locale.CY, locale.MAS, locale.TEO_KE, locale.IG, locale.YUE_HANT_MO, locale.ES_CU, locale.MS, locale.ES_SV, locale.OM_KE, locale.SN, locale.YUE, locale.GD, locale.MS_SG, locale.MER, locale.ES_US, locale.MS_ARAB, locale.PCM, locale.ROF, locale.ES_HN, locale.ZH_HANS, locale.TI_ER, locale.MAS_TZ, locale.GA, locale.KDE, locale.BHO, locale.KO, locale.MN_MONG_MN, locale.GV, locale.ML, locale.ZU, locale.ES_PR, locale.NYN, locale.JA, locale.YUE_HANT_CN, locale.KO_CN, locale.EE_TG, locale.UG, locale.VAI, locale.VAI_VAII, locale.HAW, locale.ST, locale.TH, locale.CGG, locale.KO_KP, locale.ES_BR, locale.ZH_HANT_HK, locale.ZH_HANS_MY, locale.ND, locale.YUE_HANS, locale.SI, locale.AK, locale.CEB, locale.ST_LS, locale.SO, locale.SO_DJ, locale.TI, locale.YO, locale.MR, locale.GUZ, locale.ZH_HANT_MY, locale.VAI_LATN, locale.DAV, locale.ES_BZ, locale.OR, locale.VUN, locale.ZH_HANS_MO, locale.KW, locale.SO_KE, locale.ES_GT, locale.ES_MX, locale.KAM, locale.ES_419, locale.ES_DO, locale.KI, locale.ZH, locale.KN, locale.ZH_HANT_MO, locale.DOI, locale.ES_NI, locale.EBU:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 94 locales
	case locale.EN_AS, locale.EN_GI, locale.EN_VC, locale.EN_AI, locale.EN_CC, locale.EN_MY, locale.EN_AE, locale.EN_BB, locale.EN_PR, locale.EN_LS, locale.EN_MW, locale.EN_MG, locale.EN_SG, locale.EN_JM, locale.EN_SD, locale.EN_LC, locale.EN_DG, locale.EN_WS, locale.EN_ZW, locale.EN_TV, locale.EN_BZ, locale.EN_PK, locale.EN_SH, locale.EN_KI, locale.EN_LR, locale.EN_JE, locale.EN_001, locale.EN_IO, locale.EN_TT, locale.EN_IL, locale.EN_NR, locale.EN_BS, locale.EN_CX, locale.EN_GU, locale.EN_SX, locale.EN_SZ, locale.EN_MS, locale.EN_GS, locale.EN_VI, locale.EN_MO, locale.EN_AG, locale.EN_CM, locale.EN_TC, locale.EN_BI, locale.EN_AU, locale.EN_RW, locale.EN_ZM, locale.EN_GG, locale.EN_MU, locale.EN_PN, locale.EN_VU, locale.EN_SL, locale.EN_GD, locale.EN_CA, locale.EN_IM, locale.EN_PG, locale.EN_HK, locale.EN_SC, locale.EN_GY, locale.EN_GB, locale.EN, locale.EN_KN, locale.EN_PH, locale.EN_BW, locale.EN_FM, locale.EN_VG, locale.EN_NG, locale.EN_TO, locale.EN_GH, locale.EN_ER, locale.EN_CY, locale.EN_UM, locale.EN_CK, locale.EN_NF, locale.EN_NZ, locale.EN_UG, locale.EN_FJ, locale.EN_MH, locale.EN_MP, locale.EN_PW, locale.EN_DM, locale.EN_GM, locale.EN_TZ, locale.EN_KY, locale.EN_BM, locale.EN_FK, locale.EN_NU, locale.EN_SS, locale.EN_SB, locale.EN_IE, locale.EN_MT, locale.EN_KE, locale.EN_TK, locale.EN_NA:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 70 locales
	case locale.PT_AO, locale.RU, locale.YAV, locale.PT_TL, locale.TZM, locale.UK, locale.DUA, locale.FF_LATN_BF, locale.RU_MD, locale.TK, locale.RU_BY, locale.FF_LATN_GN, locale.PT_CH, locale.EO, locale.BR, locale.PL, locale.FF_LATN_GM, locale.KSF, locale.HU, locale.PT_LU, locale.FF, locale.FF_LATN_NE, locale.EWO, locale.FF_LATN_CM, locale.FF_LATN_LR, locale.FF_LATN_SL, locale.NMG, locale.PT_GW, locale.KEA, locale.FF_LATN_GW, locale.UZ, locale.HT, locale.LV, locale.PT_MZ, locale.KK_CYRL, locale.SQ_XK, locale.UZ_CYRL, locale.CV, locale.TG, locale.BE_TARASK, locale.SMN, locale.FF_LATN, locale.FF_LATN_MR, locale.KK, locale.SQ, locale.KK_KZ, locale.SQ_MK, locale.PT_GQ, locale.CS, locale.FR_CA, locale.TT, locale.HY, locale.PT_MO, locale.RU_KZ, locale.DYO, locale.FF_LATN_NG, locale.PT_PT, locale.KA, locale.PT_CV, locale.SK, locale.BE, locale.BAS, locale.RU_UA, locale.PT_ST, locale.FF_LATN_GH, locale.SZL, locale.UZ_LATN, locale.BG, locale.PRG, locale.SAH:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 59 locales
	case locale.ES_PH, locale.CA_ES_VALENCIA, locale.NDS, locale.RO, locale.CA, locale.SR_LATN_XK, locale.FR_LU, locale.LN_CF, locale.LB, locale.DSB, locale.EL, locale.SR_CYRL, locale.CA_FR, locale.SR_CYRL_XK, locale.FR_MA, locale.ES_IC, locale.SR_LATN_BA, locale.LN_CG, locale.VI, locale.SC, locale.DE_IT, locale.NDS_NL, locale.CA_IT, locale.DA_GL, locale.SR_LATN_ME, locale.SR_CYRL_ME, locale.IT_SM, locale.CA_AD, locale.HSB, locale.BS_CYRL, locale.SR_LATN, locale.AZ_CYRL, locale.AST, locale.EL_POLYTON, locale.MK, locale.ES, locale.DE_BE, locale.EL_CY, locale.AZ, locale.DA, locale.GL, locale.VMW, locale.BS_LATN, locale.IT, locale.DE, locale.IS, locale.LIJ, locale.SR, locale.LN_AO, locale.KU, locale.LN, locale.BS, locale.IT_VA, locale.DE_LU, locale.AZ_LATN, locale.LLD, locale.RO_MD, locale.SR_CYRL_BA, locale.ES_EA:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 43 locales
	case locale.FR_BI, locale.FR_MG, locale.FR_MQ, locale.FR, locale.FR_PF, locale.FR_PM, locale.FR_DJ, locale.FR_ML, locale.FR_TG, locale.FR_BL, locale.FR_CG, locale.FR_BF, locale.FR_CI, locale.FR_SC, locale.FR_CM, locale.FR_HT, locale.FR_GF, locale.FR_CF, locale.FR_SY, locale.FR_VU, locale.FR_GA, locale.FR_SN, locale.FR_NC, locale.FR_NE, locale.FR_MR, locale.FR_YT, locale.FR_KM, locale.FR_TD, locale.FR_MF, locale.FR_BE, locale.FR_CD, locale.FR_GQ, locale.FR_GP, locale.FR_WF, locale.FR_MC, locale.FR_RE, locale.FR_BJ, locale.FR_CH, locale.FR_RW, locale.FR_GN, locale.FR_MU, locale.FR_DZ, locale.FR_TN:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 22 locales
	case locale.AR_SY, locale.AR_TD, locale.AR_DJ, locale.AR_KM, locale.AR_EG, locale.AR_ER, locale.AR_SS, locale.AR_SA, locale.AR_EH, locale.AR_IQ, locale.AR_OM, locale.AR_AE, locale.AR_BH, locale.AR, locale.AR_QA, locale.AR_KW, locale.AR_YE, locale.AR_IL, locale.AR_JO, locale.AR_SO, locale.AR_PS, locale.AR_SD:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	// Group of 21 locales
	case locale.MS_ARAB_BN, locale.FUR, locale.YRL_VE, locale.PT, locale.YRL_CO, locale.KKJ, locale.MS_BN, locale.ES_AR, locale.QU_BO, locale.RW, locale.MGH, locale.IA, locale.JV, locale.ES_UY, locale.YRL, locale.NNH, locale.JGO, locale.SW_CD, locale.ES_CO, locale.KGP, locale.WO:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 17 locales
	case locale.TA, locale.HI, locale.KXV, locale.KXV_LATN, locale.SA, locale.TE, locale.DZ, locale.GU, locale.PA_GURU, locale.KXV_DEVA, locale.XNR, locale.KOK_LATN, locale.KXV_ORYA, locale.KXV_TELU, locale.BN_IN, locale.TA_LK, locale.PA:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 12 locales
	case locale.FF_ADLM_BF, locale.FF_ADLM_NG, locale.FF_ADLM_MR, locale.FF_ADLM_SN, locale.FF_ADLM_GH, locale.FF_ADLM_GW, locale.FF_ADLM_GM, locale.FF_ADLM_SL, locale.FF_ADLM, locale.FF_ADLM_NE, locale.FF_ADLM_LR, locale.FF_ADLM_CM:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "⹁",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 10 locales
	case locale.SE_FI, locale.SE_SE, locale.SE, locale.FI, locale.SV, locale.LT, locale.SV_FI, locale.ET, locale.SV_AX, locale.KSH:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 9 locales
	case locale.SU, locale.SU_LATN, locale.TR_CY, locale.ES_BO, locale.TR, locale.ID, locale.ES_GQ, locale.MS_ID, locale.MUA:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.EN_DE, locale.EN_IT, locale.EN_DK, locale.EN_PL, locale.EN_SI, locale.EN_ES, locale.EN_RO, locale.EN_BE:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.NL_SX, locale.NL, locale.NL_CW, locale.NL_SR, locale.ES_PY, locale.NL_BQ, locale.NL_AW, locale.NL_BE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 8 locales
	case locale.SHI, locale.AGQ, locale.ZGH, locale.SHI_LATN, locale.OC_ES, locale.KAB, locale.SHI_TFNG, locale.OC:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.EN_SE, locale.EN_SK, locale.EN_CZ, locale.EN_HU, locale.EN_NO, locale.EN_FI, locale.EN_PT:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.KM, locale.KSB, locale.BEZ, locale.SBP, locale.LG, locale.LUO, locale.RWK:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 7 locales
	case locale.SS_SZ, locale.AF, locale.SS, locale.AF_NA, locale.NR, locale.ES_CR, locale.VE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 6 locales
	case locale.AR_TN, locale.AR_LY, locale.AR_DZ, locale.AR_MA, locale.AR_MR, locale.AR_LB:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "\u200f#,##0.00 ¤;\u200f-#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e-",
		}
	// Group of 6 locales
	case locale.EU, locale.HR, locale.HR_BA, locale.SL, locale.FO_DK, locale.FO:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 6 locales
	case locale.LO, locale.ES_VE, locale.ES_CL, locale.KL, locale.SG, locale.ES_EC:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00;¤-#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.BGN_OM, locale.BGN, locale.BGN_IR, locale.BGN_AE, locale.BGN_AF:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: "٫",
			MinusSign:        "-",
		}
	// Group of 5 locales
	case locale.CE, locale.MY, locale.ASA, locale.XOG, locale.TPI:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DE_AT, locale.OS_RU, locale.OS, locale.TS:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.DJE, locale.TWQ, locale.KHQ, locale.SES:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00¤",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 4 locales
	case locale.GSW, locale.GSW_FR, locale.GSW_LI, locale.RM:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "−",
		}
	// Group of 4 locales
	case locale.NN, locale.NB, locale.NB_SJ, locale.NO:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤;-#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "−",
		}
	// Group of 3 locales
	case locale.AS, locale.NE_IN, locale.NE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.CCP, locale.CCP_IN, locale.BN:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##,##0.00¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.SEH, locale.RN, locale.LU:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00¤",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 3 locales
	case locale.UZ_ARAB, locale.PS, locale.PS_PK:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "\u200e−",
		}
	// Group of 2 locales
	case locale.BLO, locale.IE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.DE_CH, locale.IT_CH:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.EN_IN, locale.HI_LATN:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤#,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KS_ARAB, locale.KS:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.KY, locale.RU_KG:
		return FormatInfo{
			Symbol:           "сом",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.LMO, locale.WAE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.NSO, locale.MFE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.TN_BW, locale.TN:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	// Group of 2 locales
	case locale.UR, locale.UR_IN:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.BAL_LATN:
		return FormatInfo{
			Symbol:           "KGSS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.BRX:
		return FormatInfo{
			Symbol:           "के.जि.एस",
			Format:           "¤ #,##,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.DE_LI:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_150:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_AT:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_CH:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤ #,##0.00;¤-#,##0.00",
			GroupSeparator:   "’",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_FR:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ID:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_MV:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.EN_NL:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤ #,##0.00;¤ -#,##0.00",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.EN_ZA:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.FA:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "\u200e¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FA_AF:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e−",
		}
	case locale.FY:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00;¤ #,##0.00-",
			GroupSeparator:   ".",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.HE:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "\u200f#,##0.00 \u200f¤;\u200f-#,##0.00 \u200f¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.LUY:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00;¤- #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.NQO:
		return FormatInfo{
			Symbol:           "ߞߜ߭ߛ",
			Format:           "¤ #,##0.00",
			GroupSeparator:   "،",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	case locale.PA_ARAB:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤ #,##0.00",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "\u200e-",
		}
	case locale.TOK:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,#0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.VEC:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   " ",
			DecimalSeparator: ",",
			MinusSign:        "-",
		}
	case locale.XH:
		return FormatInfo{
			Symbol:           "⃀",
			Format:           "¤#,##0.00",
			GroupSeparator:   " ",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	default:
		return FormatInfo{
			Symbol:           "KGS",
			Format:           "#,##0.00 ¤",
			GroupSeparator:   ",",
			DecimalSeparator: ".",
			MinusSign:        "-",
		}
	}
}
