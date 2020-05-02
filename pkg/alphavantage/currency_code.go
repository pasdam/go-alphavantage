package alphavantage

// CurrencyCode is the enumeration of all supported currencies
type CurrencyCode string

const (
	// CurrencyCodeAED identifies the United Arab Emirates Dirham
	CurrencyCodeAED = "AED"

	// CurrencyCodeAFN identifies the Afghan Afghani
	CurrencyCodeAFN = "AFN"

	// CurrencyCodeALL identifies the Albanian Lek
	CurrencyCodeALL = "ALL"

	// CurrencyCodeAMD identifies the Armenian Dram
	CurrencyCodeAMD = "AMD"

	// CurrencyCodeANG identifies the Netherlands Antillean Guilder
	CurrencyCodeANG = "ANG"

	// CurrencyCodeAOA identifies the Angolan Kwanza
	CurrencyCodeAOA = "AOA"

	// CurrencyCodeARS identifies the Argentine Peso
	CurrencyCodeARS = "ARS"

	// CurrencyCodeAUD identifies the Australian Dollar
	CurrencyCodeAUD = "AUD"

	// CurrencyCodeAWG identifies the Aruban Florin
	CurrencyCodeAWG = "AWG"

	// CurrencyCodeAZN identifies the Azerbaijani Manat
	CurrencyCodeAZN = "AZN"

	// CurrencyCodeBAM identifies the Bosnia-Herzegovina Convertible Mark
	CurrencyCodeBAM = "BAM"

	// CurrencyCodeBBD identifies the Barbadian Dollar
	CurrencyCodeBBD = "BBD"

	// CurrencyCodeBDT identifies the Bangladeshi Taka
	CurrencyCodeBDT = "BDT"

	// CurrencyCodeBGN identifies the Bulgarian Lev
	CurrencyCodeBGN = "BGN"

	// CurrencyCodeBHD identifies the Bahraini Dinar
	CurrencyCodeBHD = "BHD"

	// CurrencyCodeBIF identifies the Burundian Franc
	CurrencyCodeBIF = "BIF"

	// CurrencyCodeBMD identifies the Bermudan Dollar
	CurrencyCodeBMD = "BMD"

	// CurrencyCodeBND identifies the Brunei Dollar
	CurrencyCodeBND = "BND"

	// CurrencyCodeBOB identifies the Bolivian Boliviano
	CurrencyCodeBOB = "BOB"

	// CurrencyCodeBRL identifies the Brazilian Real
	CurrencyCodeBRL = "BRL"

	// CurrencyCodeBSD identifies the Bahamian Dollar
	CurrencyCodeBSD = "BSD"

	// CurrencyCodeBTN identifies the Bhutanese Ngultrum
	CurrencyCodeBTN = "BTN"

	// CurrencyCodeBWP identifies the Botswanan Pula
	CurrencyCodeBWP = "BWP"

	// CurrencyCodeBZD identifies the Belize Dollar
	CurrencyCodeBZD = "BZD"

	// CurrencyCodeCAD identifies the Canadian Dollar
	CurrencyCodeCAD = "CAD"

	// CurrencyCodeCDF identifies the Congolese Franc
	CurrencyCodeCDF = "CDF"

	// CurrencyCodeCHF identifies the Swiss Franc
	CurrencyCodeCHF = "CHF"

	// CurrencyCodeCLF identifies the Chilean Unit of Account UF
	CurrencyCodeCLF = "CLF"

	// CurrencyCodeCLP identifies the Chilean Peso
	CurrencyCodeCLP = "CLP"

	// CurrencyCodeCNH identifies the Chinese Yuan Offshore
	CurrencyCodeCNH = "CNH"

	// CurrencyCodeCNY identifies the Chinese Yuan
	CurrencyCodeCNY = "CNY"

	// CurrencyCodeCOP identifies the Colombian Peso
	CurrencyCodeCOP = "COP"

	// CurrencyCodeCUP identifies the Cuban Peso
	CurrencyCodeCUP = "CUP"

	// CurrencyCodeCVE identifies the Cape Verdean Escudo
	CurrencyCodeCVE = "CVE"

	// CurrencyCodeCZK identifies the Czech Republic Koruna
	CurrencyCodeCZK = "CZK"

	// CurrencyCodeDJF identifies the Djiboutian Franc
	CurrencyCodeDJF = "DJF"

	// CurrencyCodeDKK identifies the Danish Krone
	CurrencyCodeDKK = "DKK"

	// CurrencyCodeDOP identifies the Dominican Peso
	CurrencyCodeDOP = "DOP"

	// CurrencyCodeDZD identifies the Algerian Dinar
	CurrencyCodeDZD = "DZD"

	// CurrencyCodeEGP identifies the Egyptian Pound
	CurrencyCodeEGP = "EGP"

	// CurrencyCodeERN identifies the Eritrean Nakfa
	CurrencyCodeERN = "ERN"

	// CurrencyCodeETB identifies the Ethiopian Birr
	CurrencyCodeETB = "ETB"

	// CurrencyCodeEUR identifies the Euro
	CurrencyCodeEUR = "EUR"

	// CurrencyCodeFJD identifies the Fijian Dollar
	CurrencyCodeFJD = "FJD"

	// CurrencyCodeFKP identifies the Falkland Islands Pound
	CurrencyCodeFKP = "FKP"

	// CurrencyCodeGBP identifies the British Pound Sterling
	CurrencyCodeGBP = "GBP"

	// CurrencyCodeGEL identifies the Georgian Lari
	CurrencyCodeGEL = "GEL"

	// CurrencyCodeGHS identifies the Ghanaian Cedi
	CurrencyCodeGHS = "GHS"

	// CurrencyCodeGIP identifies the Gibraltar Pound
	CurrencyCodeGIP = "GIP"

	// CurrencyCodeGMD identifies the Gambian Dalasi
	CurrencyCodeGMD = "GMD"

	// CurrencyCodeGNF identifies the Guinean Franc
	CurrencyCodeGNF = "GNF"

	// CurrencyCodeGTQ identifies the Guatemalan Quetzal
	CurrencyCodeGTQ = "GTQ"

	// CurrencyCodeGYD identifies the Guyanaese Dollar
	CurrencyCodeGYD = "GYD"

	// CurrencyCodeHKD identifies the Hong Kong Dollar
	CurrencyCodeHKD = "HKD"

	// CurrencyCodeHNL identifies the Honduran Lempira
	CurrencyCodeHNL = "HNL"

	// CurrencyCodeHRK identifies the Croatian Kuna
	CurrencyCodeHRK = "HRK"

	// CurrencyCodeHTG identifies the Haitian Gourde
	CurrencyCodeHTG = "HTG"

	// CurrencyCodeHUF identifies the Hungarian Forint
	CurrencyCodeHUF = "HUF"

	// CurrencyCodeIDR identifies the Indonesian Rupiah
	CurrencyCodeIDR = "IDR"

	// CurrencyCodeILS identifies the Israeli New Sheqel
	CurrencyCodeILS = "ILS"

	// CurrencyCodeINR identifies the Indian Rupee
	CurrencyCodeINR = "INR"

	// CurrencyCodeIQD identifies the Iraqi Dinar
	CurrencyCodeIQD = "IQD"

	// CurrencyCodeIRR identifies the Iranian Rial
	CurrencyCodeIRR = "IRR"

	// CurrencyCodeISK identifies the Icelandic Krona
	CurrencyCodeISK = "ISK"

	// CurrencyCodeJEP identifies the Jersey Pound
	CurrencyCodeJEP = "JEP"

	// CurrencyCodeJMD identifies the Jamaican Dollar
	CurrencyCodeJMD = "JMD"

	// CurrencyCodeJOD identifies the Jordanian Dinar
	CurrencyCodeJOD = "JOD"

	// CurrencyCodeJPY identifies the Japanese Yen
	CurrencyCodeJPY = "JPY"

	// CurrencyCodeKES identifies the Kenyan Shilling
	CurrencyCodeKES = "KES"

	// CurrencyCodeKGS identifies the Kyrgystani Som
	CurrencyCodeKGS = "KGS"

	// CurrencyCodeKHR identifies the Cambodian Riel
	CurrencyCodeKHR = "KHR"

	// CurrencyCodeKMF identifies the Comorian Franc
	CurrencyCodeKMF = "KMF"

	// CurrencyCodeKPW identifies the North Korean Won
	CurrencyCodeKPW = "KPW"

	// CurrencyCodeKRW identifies the South Korean Won
	CurrencyCodeKRW = "KRW"

	// CurrencyCodeKWD identifies the Kuwaiti Dinar
	CurrencyCodeKWD = "KWD"

	// CurrencyCodeKYD identifies the Cayman Islands Dollar
	CurrencyCodeKYD = "KYD"

	// CurrencyCodeKZT identifies the Kazakhstani Tenge
	CurrencyCodeKZT = "KZT"

	// CurrencyCodeLAK identifies the Laotian Kip
	CurrencyCodeLAK = "LAK"

	// CurrencyCodeLBP identifies the Lebanese Pound
	CurrencyCodeLBP = "LBP"

	// CurrencyCodeLKR identifies the Sri Lankan Rupee
	CurrencyCodeLKR = "LKR"

	// CurrencyCodeLRD identifies the Liberian Dollar
	CurrencyCodeLRD = "LRD"

	// CurrencyCodeLSL identifies the Lesotho Loti
	CurrencyCodeLSL = "LSL"

	// CurrencyCodeLYD identifies the Libyan Dinar
	CurrencyCodeLYD = "LYD"

	// CurrencyCodeMAD identifies the Moroccan Dirham
	CurrencyCodeMAD = "MAD"

	// CurrencyCodeMDL identifies the Moldovan Leu
	CurrencyCodeMDL = "MDL"

	// CurrencyCodeMGA identifies the Malagasy Ariary
	CurrencyCodeMGA = "MGA"

	// CurrencyCodeMKD identifies the Macedonian Denar
	CurrencyCodeMKD = "MKD"

	// CurrencyCodeMMK identifies the Myanma Kyat
	CurrencyCodeMMK = "MMK"

	// CurrencyCodeMNT identifies the Mongolian Tugrik
	CurrencyCodeMNT = "MNT"

	// CurrencyCodeMOP identifies the Macanese Pataca
	CurrencyCodeMOP = "MOP"

	// CurrencyCodeMRO identifies the Mauritanian Ouguiya (pre-2018)
	CurrencyCodeMRO = "MRO"

	// CurrencyCodeMRU identifies the Mauritanian Ouguiya
	CurrencyCodeMRU = "MRU"

	// CurrencyCodeMUR identifies the Mauritian Rupee
	CurrencyCodeMUR = "MUR"

	// CurrencyCodeMVR identifies the Maldivian Rufiyaa
	CurrencyCodeMVR = "MVR"

	// CurrencyCodeMWK identifies the Malawian Kwacha
	CurrencyCodeMWK = "MWK"

	// CurrencyCodeMXN identifies the Mexican Peso
	CurrencyCodeMXN = "MXN"

	// CurrencyCodeMYR identifies the Malaysian Ringgit
	CurrencyCodeMYR = "MYR"

	// CurrencyCodeMZN identifies the Mozambican Metical
	CurrencyCodeMZN = "MZN"

	// CurrencyCodeNAD identifies the Namibian Dollar
	CurrencyCodeNAD = "NAD"

	// CurrencyCodeNGN identifies the Nigerian Naira
	CurrencyCodeNGN = "NGN"

	// CurrencyCodeNOK identifies the Norwegian Krone
	CurrencyCodeNOK = "NOK"

	// CurrencyCodeNPR identifies the Nepalese Rupee
	CurrencyCodeNPR = "NPR"

	// CurrencyCodeNZD identifies the New Zealand Dollar
	CurrencyCodeNZD = "NZD"

	// CurrencyCodeOMR identifies the Omani Rial
	CurrencyCodeOMR = "OMR"

	// CurrencyCodePAB identifies the Panamanian Balboa
	CurrencyCodePAB = "PAB"

	// CurrencyCodePEN identifies the Peruvian Nuevo Sol
	CurrencyCodePEN = "PEN"

	// CurrencyCodePGK identifies the Papua New Guinean Kina
	CurrencyCodePGK = "PGK"

	// CurrencyCodePHP identifies the Philippine Peso
	CurrencyCodePHP = "PHP"

	// CurrencyCodePKR identifies the Pakistani Rupee
	CurrencyCodePKR = "PKR"

	// CurrencyCodePLN identifies the Polish Zloty
	CurrencyCodePLN = "PLN"

	// CurrencyCodePYG identifies the Paraguayan Guarani
	CurrencyCodePYG = "PYG"

	// CurrencyCodeQAR identifies the Qatari Rial
	CurrencyCodeQAR = "QAR"

	// CurrencyCodeRON identifies the Romanian Leu
	CurrencyCodeRON = "RON"

	// CurrencyCodeRSD identifies the Serbian Dinar
	CurrencyCodeRSD = "RSD"

	// CurrencyCodeRUB identifies the Russian Ruble
	CurrencyCodeRUB = "RUB"

	// CurrencyCodeRUR identifies the Old Russian Ruble
	CurrencyCodeRUR = "RUR"

	// CurrencyCodeRWF identifies the Rwandan Franc
	CurrencyCodeRWF = "RWF"

	// CurrencyCodeSAR identifies the Saudi Riyal
	CurrencyCodeSAR = "SAR"

	// CurrencyCodeSBD identifies the ,Solomon Islands Dollar
	CurrencyCodeSBD = "SBD"

	// CurrencyCodeSCR identifies the Seychellois Rupee
	CurrencyCodeSCR = "SCR"

	// CurrencyCodeSDG identifies the Sudanese Pound
	CurrencyCodeSDG = "SDG"

	// CurrencyCodeSEK identifies the Swedish Krona
	CurrencyCodeSEK = "SEK"

	// CurrencyCodeSGD identifies the Singapore Dollar
	CurrencyCodeSGD = "SGD"

	// CurrencyCodeSHP identifies the Saint Helena Pound
	CurrencyCodeSHP = "SHP"

	// CurrencyCodeSLL identifies the Sierra Leonean Leone
	CurrencyCodeSLL = "SLL"

	// CurrencyCodeSOS identifies the Somali Shilling
	CurrencyCodeSOS = "SOS"

	// CurrencyCodeSRD identifies the Surinamese Dollar
	CurrencyCodeSRD = "SRD"

	// CurrencyCodeSYP identifies the Syrian Pound
	CurrencyCodeSYP = "SYP"

	// CurrencyCodeSZL identifies the Swazi Lilangeni
	CurrencyCodeSZL = "SZL"

	// CurrencyCodeTHB identifies the Thai Baht
	CurrencyCodeTHB = "THB"

	// CurrencyCodeTJS identifies the Tajikistani Somoni
	CurrencyCodeTJS = "TJS"

	// CurrencyCodeTMT identifies the Turkmenistani Manat
	CurrencyCodeTMT = "TMT"

	// CurrencyCodeTND identifies the Tunisian Dinar
	CurrencyCodeTND = "TND"

	// CurrencyCodeTOP identifies the Tongan Pa'anga
	CurrencyCodeTOP = "TOP"

	// CurrencyCodeTRY identifies the Turkish Lira
	CurrencyCodeTRY = "TRY"

	// CurrencyCodeTTD identifies the Trinidad and Tobago Dollar
	CurrencyCodeTTD = "TTD"

	// CurrencyCodeTWD identifies the New Taiwan Dollar
	CurrencyCodeTWD = "TWD"

	// CurrencyCodeTZS identifies the Tanzanian Shilling
	CurrencyCodeTZS = "TZS"

	// CurrencyCodeUAH identifies the Ukrainian Hryvnia
	CurrencyCodeUAH = "UAH"

	// CurrencyCodeUGX identifies the Ugandan Shilling
	CurrencyCodeUGX = "UGX"

	// CurrencyCodeUSD identifies the United States Dollar
	CurrencyCodeUSD = "USD"

	// CurrencyCodeUYU identifies the Uruguayan Peso
	CurrencyCodeUYU = "UYU"

	// CurrencyCodeUZS identifies the Uzbekistan Som
	CurrencyCodeUZS = "UZS"

	// CurrencyCodeVND identifies the Vietnamese Dong
	CurrencyCodeVND = "VND"

	// CurrencyCodeVUV identifies the Vanuatu Vatu
	CurrencyCodeVUV = "VUV"

	// CurrencyCodeWST identifies the Samoan Tala
	CurrencyCodeWST = "WST"

	// CurrencyCodeXAF identifies the CFA Franc BEAC
	CurrencyCodeXAF = "XAF"

	// CurrencyCodeXAG identifies the Silver Ounce
	CurrencyCodeXAG = "XAG"

	// CurrencyCodeXAU identifies the Gold Ounce
	CurrencyCodeXAU = "XAU"

	// CurrencyCodeXCD identifies the East Caribbean Dollar
	CurrencyCodeXCD = "XCD"

	// CurrencyCodeXDR identifies the Special Drawing Rights
	CurrencyCodeXDR = "XDR"

	// CurrencyCodeXOF identifies the CFA Franc BCEAO
	CurrencyCodeXOF = "XOF"

	// CurrencyCodeXPF identifies the CFP Franc
	CurrencyCodeXPF = "XPF"

	// CurrencyCodeYER identifies the Yemeni Rial
	CurrencyCodeYER = "YER"

	// CurrencyCodeZAR identifies the South African Rand
	CurrencyCodeZAR = "ZAR"

	// CurrencyCodeZMW identifies the Zambian Kwacha
	CurrencyCodeZMW = "ZMW"

	// CurrencyCodeZWL identifies the Zimbabwean Dollar
	CurrencyCodeZWL = "ZWL"
)

// Code returns the 3 letter code representing the currency
func (c CurrencyCode) String() string {
	return string(c)
}
