package alphavantage

// CurrencyCode is the enumeration of all supported physical and digital
// currencies
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

	// CurrencyCode1ST identifies the FirstBlood digital currency
	CurrencyCode1ST = "1ST"

	// CurrencyCode2GIVE identifies the GiveCoin digital currency
	CurrencyCode2GIVE = "2GIVE"

	// CurrencyCode808 identifies the 808Coin digital currency
	CurrencyCode808 = "808"

	// CurrencyCodeABT identifies the ArcBlock digital currency
	CurrencyCodeABT = "ABT"

	// CurrencyCodeABY identifies the ArtByte digital currency
	CurrencyCodeABY = "ABY"

	// CurrencyCodeAC identifies the AsiaCoin digital currency
	CurrencyCodeAC = "AC"

	// CurrencyCodeACT identifies the Achain digital currency
	CurrencyCodeACT = "ACT"

	// CurrencyCodeADA identifies the Cardano digital currency
	CurrencyCodeADA = "ADA"

	// CurrencyCodeADT identifies the adToken digital currency
	CurrencyCodeADT = "ADT"

	// CurrencyCodeADX identifies the AdEx digital currency
	CurrencyCodeADX = "ADX"

	// CurrencyCodeAE identifies the Aeternity digital currency
	CurrencyCodeAE = "AE"

	// CurrencyCodeAEON identifies the Aeon digital currency
	CurrencyCodeAEON = "AEON"

	// CurrencyCodeAGI identifies the SingularityNET digital currency
	CurrencyCodeAGI = "AGI"

	// CurrencyCodeAGRS identifies the IDNI-Agoras digital currency
	CurrencyCodeAGRS = "AGRS"

	// CurrencyCodeAI identifies the POLY-AI digital currency
	CurrencyCodeAI = "AI"

	// CurrencyCodeAID identifies the AidCoin digital currency
	CurrencyCodeAID = "AID"

	// CurrencyCodeAION identifies the Aion digital currency
	CurrencyCodeAION = "AION"

	// CurrencyCodeAIR identifies the AirToken digital currency
	CurrencyCodeAIR = "AIR"

	// CurrencyCodeAKY identifies the Akuya-Coin digital currency
	CurrencyCodeAKY = "AKY"

	// CurrencyCodeALIS identifies the ALIS digital currency
	CurrencyCodeALIS = "ALIS"

	// CurrencyCodeAMBER identifies the AmberCoin digital currency
	CurrencyCodeAMBER = "AMBER"

	// CurrencyCodeAMP identifies the Synereo digital currency
	CurrencyCodeAMP = "AMP"

	// CurrencyCodeAMPL identifies the Ampleforth digital currency
	CurrencyCodeAMPL = "AMPL"

	// CurrencyCodeANC identifies the Anoncoin digital currency
	CurrencyCodeANC = "ANC"

	// CurrencyCodeANT identifies the Aragon digital currency
	CurrencyCodeANT = "ANT"

	// CurrencyCodeAPPC identifies the AppCoins digital currency
	CurrencyCodeAPPC = "APPC"

	// CurrencyCodeAPX identifies the APX-Ventures digital currency
	CurrencyCodeAPX = "APX"

	// CurrencyCodeARDR identifies the Ardor digital currency
	CurrencyCodeARDR = "ARDR"

	// CurrencyCodeARK identifies the Ark digital currency
	CurrencyCodeARK = "ARK"

	// CurrencyCodeARN identifies the Aeron digital currency
	CurrencyCodeARN = "ARN"

	// CurrencyCodeAST identifies the AirSwap digital currency
	CurrencyCodeAST = "AST"

	// CurrencyCodeATB identifies the ATBCoin digital currency
	CurrencyCodeATB = "ATB"

	// CurrencyCodeATM identifies the ATMChain digital currency
	CurrencyCodeATM = "ATM"

	// CurrencyCodeATS identifies the Authorship digital currency
	CurrencyCodeATS = "ATS"

	// CurrencyCodeAUR identifies the Auroracoin digital currency
	CurrencyCodeAUR = "AUR"

	// CurrencyCodeAVT identifies the Aventus digital currency
	CurrencyCodeAVT = "AVT"

	// CurrencyCodeB3 identifies the B3Coin digital currency
	CurrencyCodeB3 = "B3"

	// CurrencyCodeBAT identifies the Basic-Attention-Token digital currency
	CurrencyCodeBAT = "BAT"

	// CurrencyCodeBAY identifies the BitBay digital currency
	CurrencyCodeBAY = "BAY"

	// CurrencyCodeBBR identifies the Boolberry digital currency
	CurrencyCodeBBR = "BBR"

	// CurrencyCodeBCAP identifies the BCAP digital currency
	CurrencyCodeBCAP = "BCAP"

	// CurrencyCodeBCC identifies the BitConnect digital currency
	CurrencyCodeBCC = "BCC"

	// CurrencyCodeBCD identifies the Bitcoin-Diamond digital currency
	CurrencyCodeBCD = "BCD"

	// CurrencyCodeBCH identifies the Bitcoin-Cash digital currency
	CurrencyCodeBCH = "BCH"

	// CurrencyCodeBCN identifies the Bytecoin digital currency
	CurrencyCodeBCN = "BCN"

	// CurrencyCodeBCPT identifies the BlockMason-Credit-Protocol-Token digital currency
	CurrencyCodeBCPT = "BCPT"

	// CurrencyCodeBCX identifies the BitcoinX digital currency
	CurrencyCodeBCX = "BCX"

	// CurrencyCodeBCY identifies the BitCrystals digital currency
	CurrencyCodeBCY = "BCY"

	// CurrencyCodeBDL identifies the Bitdeal digital currency
	CurrencyCodeBDL = "BDL"

	// CurrencyCodeBEE identifies the Bee-Token digital currency
	CurrencyCodeBEE = "BEE"

	// CurrencyCodeBELA identifies the BelaCoin digital currency
	CurrencyCodeBELA = "BELA"

	// CurrencyCodeBET identifies the DAO-Casino digital currency
	CurrencyCodeBET = "BET"

	// CurrencyCodeBFT identifies the BF-Token digital currency
	CurrencyCodeBFT = "BFT"

	// CurrencyCodeBIS identifies the Bismuth digital currency
	CurrencyCodeBIS = "BIS"

	// CurrencyCodeBITB identifies the BitBean digital currency
	CurrencyCodeBITB = "BITB"

	// CurrencyCodeBITBTC identifies the BitBTC digital currency
	CurrencyCodeBITBTC = "BITBTC"

	// CurrencyCodeBITCNY identifies the BitCNY digital currency
	CurrencyCodeBITCNY = "BITCNY"

	// CurrencyCodeBITEUR identifies the BitEUR digital currency
	CurrencyCodeBITEUR = "BITEUR"

	// CurrencyCodeBITGOLD identifies the BitGOLD digital currency
	CurrencyCodeBITGOLD = "BITGOLD"

	// CurrencyCodeBITSILVER identifies the BitSILVER digital currency
	CurrencyCodeBITSILVER = "BITSILVER"

	// CurrencyCodeBITUSD identifies the BitUSD digital currency
	CurrencyCodeBITUSD = "BITUSD"

	// CurrencyCodeBIX identifies the Bibox-Token digital currency
	CurrencyCodeBIX = "BIX"

	// CurrencyCodeBLITZ identifies the Blitzcash digital currency
	CurrencyCodeBLITZ = "BLITZ"

	// CurrencyCodeBLK identifies the Blackcoin digital currency
	CurrencyCodeBLK = "BLK"

	// CurrencyCodeBLN identifies the Bolenum digital currency
	CurrencyCodeBLN = "BLN"

	// CurrencyCodeBLOCK identifies the Blocknet digital currency
	CurrencyCodeBLOCK = "BLOCK"

	// CurrencyCodeBLZ identifies the Bluzelle digital currency
	CurrencyCodeBLZ = "BLZ"

	// CurrencyCodeBMC identifies the Blackmoon-Crypto digital currency
	CurrencyCodeBMC = "BMC"

	// CurrencyCodeBNB identifies the Binance-Coin digital currency
	CurrencyCodeBNB = "BNB"

	// CurrencyCodeBNT identifies the Bancor-Network-Token digital currency
	CurrencyCodeBNT = "BNT"

	// CurrencyCodeBNTY identifies the Bounty0x digital currency
	CurrencyCodeBNTY = "BNTY"

	// CurrencyCodeBOST identifies the BoostCoin digital currency
	CurrencyCodeBOST = "BOST"

	// CurrencyCodeBOT identifies the Bodhi digital currency
	CurrencyCodeBOT = "BOT"

	// CurrencyCodeBQ identifies the bitqy digital currency
	CurrencyCodeBQ = "BQ"

	// CurrencyCodeBRD identifies the Bread digital currency
	CurrencyCodeBRD = "BRD"

	// CurrencyCodeBRK identifies the Breakout-Coin digital currency
	CurrencyCodeBRK = "BRK"

	// CurrencyCodeBRX identifies the Breakout-Stake digital currency
	CurrencyCodeBRX = "BRX"

	// CurrencyCodeBTA identifies the Bata digital currency
	CurrencyCodeBTA = "BTA"

	// CurrencyCodeBTC identifies the Bitcoin digital currency
	CurrencyCodeBTC = "BTC"

	// CurrencyCodeBTCD identifies the BitcoinDark digital currency
	CurrencyCodeBTCD = "BTCD"

	// CurrencyCodeBTCP identifies the Bitcoin-Private digital currency
	CurrencyCodeBTCP = "BTCP"

	// CurrencyCodeBTG identifies the Bitcoin-Gold digital currency
	CurrencyCodeBTG = "BTG"

	// CurrencyCodeBTM identifies the Bitmark digital currency
	CurrencyCodeBTM = "BTM"

	// CurrencyCodeBTS identifies the BitShares digital currency
	CurrencyCodeBTS = "BTS"

	// CurrencyCodeBTSR identifies the BTSR digital currency
	CurrencyCodeBTSR = "BTSR"

	// CurrencyCodeBTX identifies the Bitcore digital currency
	CurrencyCodeBTX = "BTX"

	// CurrencyCodeBURST identifies the Burstcoin digital currency
	CurrencyCodeBURST = "BURST"

	// CurrencyCodeBUSD identifies the Binance-USD digital currency
	CurrencyCodeBUSD = "BUSD"

	// CurrencyCodeBUZZ identifies the BuzzCoin digital currency
	CurrencyCodeBUZZ = "BUZZ"

	// CurrencyCodeBYC identifies the Bytecent digital currency
	CurrencyCodeBYC = "BYC"

	// CurrencyCodeBYTOM identifies the Bytom digital currency
	CurrencyCodeBYTOM = "BYTOM"

	// CurrencyCodeC20 identifies the Crypto20 digital currency
	CurrencyCodeC20 = "C20"

	// CurrencyCodeCANN identifies the CannabisCoin digital currency
	CurrencyCodeCANN = "CANN"

	// CurrencyCodeCAT identifies the BlockCAT digital currency
	CurrencyCodeCAT = "CAT"

	// CurrencyCodeCCRB identifies the CryptoCarbon digital currency
	CurrencyCodeCCRB = "CCRB"

	// CurrencyCodeCDT identifies the Blox digital currency
	CurrencyCodeCDT = "CDT"

	// CurrencyCodeCFI identifies the Cofound-it digital currency
	CurrencyCodeCFI = "CFI"

	// CurrencyCodeCHAT identifies the ChatCoin digital currency
	CurrencyCodeCHAT = "CHAT"

	// CurrencyCodeCHIPS identifies the Chips digital currency
	CurrencyCodeCHIPS = "CHIPS"

	// CurrencyCodeCLAM identifies the Clams digital currency
	CurrencyCodeCLAM = "CLAM"

	// CurrencyCodeCLOAK identifies the CloakCoin digital currency
	CurrencyCodeCLOAK = "CLOAK"

	// CurrencyCodeCMP identifies the Compcoin digital currency
	CurrencyCodeCMP = "CMP"

	// CurrencyCodeCMT identifies the CyberMiles digital currency
	CurrencyCodeCMT = "CMT"

	// CurrencyCodeCND identifies the Cindicator digital currency
	CurrencyCodeCND = "CND"

	// CurrencyCodeCNX identifies the Cryptonex digital currency
	CurrencyCodeCNX = "CNX"

	// CurrencyCodeCOFI identifies the CoinFi digital currency
	CurrencyCodeCOFI = "COFI"

	// CurrencyCodeCOSS identifies the COSS digital currency
	CurrencyCodeCOSS = "COSS"

	// CurrencyCodeCOVAL identifies the Circuits-Of-Value digital currency
	CurrencyCodeCOVAL = "COVAL"

	// CurrencyCodeCRBIT identifies the CreditBIT digital currency
	CurrencyCodeCRBIT = "CRBIT"

	// CurrencyCodeCREA identifies the CreativeCoin digital currency
	CurrencyCodeCREA = "CREA"

	// CurrencyCodeCREDO identifies the Credo digital currency
	CurrencyCodeCREDO = "CREDO"

	// CurrencyCodeCRW identifies the Crown digital currency
	CurrencyCodeCRW = "CRW"

	// CurrencyCodeCSNO identifies the BitDice digital currency
	CurrencyCodeCSNO = "CSNO"

	// CurrencyCodeCTR identifies the Centra digital currency
	CurrencyCodeCTR = "CTR"

	// CurrencyCodeCTXC identifies the Cortex digital currency
	CurrencyCodeCTXC = "CTXC"

	// CurrencyCodeCURE identifies the CureCoin digital currency
	CurrencyCodeCURE = "CURE"

	// CurrencyCodeCVC identifies the Civic digital currency
	CurrencyCodeCVC = "CVC"

	// CurrencyCodeDAI identifies the Dai digital currency
	CurrencyCodeDAI = "DAI"

	// CurrencyCodeDAR identifies the Darcrus digital currency
	CurrencyCodeDAR = "DAR"

	// CurrencyCodeDASH identifies the Dash digital currency
	CurrencyCodeDASH = "DASH"

	// CurrencyCodeDATA identifies the DATAcoin digital currency
	CurrencyCodeDATA = "DATA"

	// CurrencyCodeDAY identifies the Chronologic digital currency
	CurrencyCodeDAY = "DAY"

	// CurrencyCodeDBC identifies the DeepBrain-Chain digital currency
	CurrencyCodeDBC = "DBC"

	// CurrencyCodeDBIX identifies the DubaiCoin digital currency
	CurrencyCodeDBIX = "DBIX"

	// CurrencyCodeDCN identifies the Dentacoin digital currency
	CurrencyCodeDCN = "DCN"

	// CurrencyCodeDCR identifies the Decred digital currency
	CurrencyCodeDCR = "DCR"

	// CurrencyCodeDCT identifies the DECENT digital currency
	CurrencyCodeDCT = "DCT"

	// CurrencyCodeDDF identifies the Digital-Developers-Fund digital currency
	CurrencyCodeDDF = "DDF"

	// CurrencyCodeDENT identifies the Dent digital currency
	CurrencyCodeDENT = "DENT"

	// CurrencyCodeDFS identifies the DFSCoin digital currency
	CurrencyCodeDFS = "DFS"

	// CurrencyCodeDGB identifies the DigiByte digital currency
	CurrencyCodeDGB = "DGB"

	// CurrencyCodeDGC identifies the Digitalcoin digital currency
	CurrencyCodeDGC = "DGC"

	// CurrencyCodeDGD identifies the DigixDAO digital currency
	CurrencyCodeDGD = "DGD"

	// CurrencyCodeDICE identifies the Etheroll digital currency
	CurrencyCodeDICE = "DICE"

	// CurrencyCodeDLT identifies the Agrello-Delta digital currency
	CurrencyCodeDLT = "DLT"

	// CurrencyCodeDMD identifies the Diamond digital currency
	CurrencyCodeDMD = "DMD"

	// CurrencyCodeDMT identifies the DMarket digital currency
	CurrencyCodeDMT = "DMT"

	// CurrencyCodeDNT identifies the district0x digital currency
	CurrencyCodeDNT = "DNT"

	// CurrencyCodeDOGE identifies the DogeCoin digital currency
	CurrencyCodeDOGE = "DOGE"

	// CurrencyCodeDOPE identifies the DopeCoin digital currency
	CurrencyCodeDOPE = "DOPE"

	// CurrencyCodeDRGN identifies the Dragonchain digital currency
	CurrencyCodeDRGN = "DRGN"

	// CurrencyCodeDTA identifies the Data digital currency
	CurrencyCodeDTA = "DTA"

	// CurrencyCodeDTB identifies the Databits digital currency
	CurrencyCodeDTB = "DTB"

	// CurrencyCodeDYN identifies the Dynamic digital currency
	CurrencyCodeDYN = "DYN"

	// CurrencyCodeEAC identifies the EarthCoin digital currency
	CurrencyCodeEAC = "EAC"

	// CurrencyCodeEBST identifies the eBoost digital currency
	CurrencyCodeEBST = "EBST"

	// CurrencyCodeEBTC identifies the eBTC digital currency
	CurrencyCodeEBTC = "EBTC"

	// CurrencyCodeECC identifies the ECC digital currency
	CurrencyCodeECC = "ECC"

	// CurrencyCodeECN identifies the E-coin digital currency
	CurrencyCodeECN = "ECN"

	// CurrencyCodeEDG identifies the Edgeless digital currency
	CurrencyCodeEDG = "EDG"

	// CurrencyCodeEDO identifies the Eidoo digital currency
	CurrencyCodeEDO = "EDO"

	// CurrencyCodeEFL identifies the Electronic-Gulden digital currency
	CurrencyCodeEFL = "EFL"

	// CurrencyCodeEGC identifies the EverGreenCoin digital currency
	CurrencyCodeEGC = "EGC"

	// CurrencyCodeEKT identifies the EDUCare digital currency
	CurrencyCodeEKT = "EKT"

	// CurrencyCodeELA identifies the Elastos digital currency
	CurrencyCodeELA = "ELA"

	// CurrencyCodeELEC identifies the Electrify.Asia digital currency
	CurrencyCodeELEC = "ELEC"

	// CurrencyCodeELF identifies the aelf digital currency
	CurrencyCodeELF = "ELF"

	// CurrencyCodeELIX identifies the Elixir digital currency
	CurrencyCodeELIX = "ELIX"

	// CurrencyCodeEMB identifies the Embercoin digital currency
	CurrencyCodeEMB = "EMB"

	// CurrencyCodeEMC identifies the Emercoin digital currency
	CurrencyCodeEMC = "EMC"

	// CurrencyCodeEMC2 identifies the Einsteinium digital currency
	CurrencyCodeEMC2 = "EMC2"

	// CurrencyCodeENG identifies the Enigma digital currency
	CurrencyCodeENG = "ENG"

	// CurrencyCodeENJ identifies the Enjin-Coin digital currency
	CurrencyCodeENJ = "ENJ"

	// CurrencyCodeENRG identifies the EnergyCoin digital currency
	CurrencyCodeENRG = "ENRG"

	// CurrencyCodeEOS identifies the EOS digital currency
	CurrencyCodeEOS = "EOS"

	// CurrencyCodeEOT identifies the EOT-Token digital currency
	CurrencyCodeEOT = "EOT"

	// CurrencyCodeEQT identifies the EquiTrader digital currency
	CurrencyCodeEQT = "EQT"

	// CurrencyCodeERC identifies the EuropeCoin digital currency
	CurrencyCodeERC = "ERC"

	// CurrencyCodeETC identifies the Ethereum-Classic digital currency
	CurrencyCodeETC = "ETC"

	// CurrencyCodeETH identifies the Ethereum digital currency
	CurrencyCodeETH = "ETH"

	// CurrencyCodeETHD identifies the Ethereum-Dark digital currency
	CurrencyCodeETHD = "ETHD"

	// CurrencyCodeETHOS identifies the Ethos digital currency
	CurrencyCodeETHOS = "ETHOS"

	// CurrencyCodeETN identifies the Electroneum digital currency
	CurrencyCodeETN = "ETN"

	// CurrencyCodeETP identifies the Metaverse-Entropy digital currency
	CurrencyCodeETP = "ETP"

	// CurrencyCodeETT identifies the EncryptoTel digital currency
	CurrencyCodeETT = "ETT"

	// CurrencyCodeEVE identifies the Devery digital currency
	CurrencyCodeEVE = "EVE"

	// CurrencyCodeEVX identifies the Everex digital currency
	CurrencyCodeEVX = "EVX"

	// CurrencyCodeEXCL identifies the ExclusiveCoin digital currency
	CurrencyCodeEXCL = "EXCL"

	// CurrencyCodeEXP identifies the Expanse digital currency
	CurrencyCodeEXP = "EXP"

	// CurrencyCodeFCT identifies the Factom digital currency
	CurrencyCodeFCT = "FCT"

	// CurrencyCodeFLDC identifies the FoldingCoin digital currency
	CurrencyCodeFLDC = "FLDC"

	// CurrencyCodeFLO identifies the FlorinCoin digital currency
	CurrencyCodeFLO = "FLO"

	// CurrencyCodeFLT identifies the FlutterCoin digital currency
	CurrencyCodeFLT = "FLT"

	// CurrencyCodeFRST identifies the FirstCoin digital currency
	CurrencyCodeFRST = "FRST"

	// CurrencyCodeFTC identifies the Feathercoin digital currency
	CurrencyCodeFTC = "FTC"

	// CurrencyCodeFUEL identifies the Etherparty digital currency
	CurrencyCodeFUEL = "FUEL"

	// CurrencyCodeFUN identifies the FunFair digital currency
	CurrencyCodeFUN = "FUN"

	// CurrencyCodeGAM identifies the Gambit digital currency
	CurrencyCodeGAM = "GAM"

	// CurrencyCodeGAME identifies the GameCredits digital currency
	CurrencyCodeGAME = "GAME"

	// CurrencyCodeGAS identifies the Gas digital currency
	CurrencyCodeGAS = "GAS"

	// CurrencyCodeGBG identifies the Golos Gold digital currency
	CurrencyCodeGBG = "GBG"

	// CurrencyCodeGBX identifies the GoByte digital currency
	CurrencyCodeGBX = "GBX"

	// CurrencyCodeGBYTE identifies the Byteball digital currency
	CurrencyCodeGBYTE = "GBYTE"

	// CurrencyCodeGCR identifies the GCRCoin digital currency
	CurrencyCodeGCR = "GCR"

	// CurrencyCodeGEO identifies the GeoCoin digital currency
	CurrencyCodeGEO = "GEO"

	// CurrencyCodeGLD identifies the GoldCoin digital currency
	CurrencyCodeGLD = "GLD"

	// CurrencyCodeGNO identifies the Gnosis-Token digital currency
	CurrencyCodeGNO = "GNO"

	// CurrencyCodeGNT identifies the Golem-Tokens digital currency
	CurrencyCodeGNT = "GNT"

	// CurrencyCodeGOLOS identifies the Golos digital currency
	CurrencyCodeGOLOS = "GOLOS"

	// CurrencyCodeGRC identifies the Gridcoin digital currency
	CurrencyCodeGRC = "GRC"

	// CurrencyCodeGRS identifies the Groestlcoin digital currency
	CurrencyCodeGRS = "GRS"

	// CurrencyCodeGRWI identifies the Growers-International digital currency
	CurrencyCodeGRWI = "GRWI"

	// CurrencyCodeGTC identifies the Game digital currency
	CurrencyCodeGTC = "GTC"

	// CurrencyCodeGTO identifies the Gifto digital currency
	CurrencyCodeGTO = "GTO"

	// CurrencyCodeGUP identifies the Guppy digital currency
	CurrencyCodeGUP = "GUP"

	// CurrencyCodeGVT identifies the Genesis-Vision digital currency
	CurrencyCodeGVT = "GVT"

	// CurrencyCodeGXS identifies the GXShares digital currency
	CurrencyCodeGXS = "GXS"

	// CurrencyCodeHBN identifies the HoboNickels digital currency
	CurrencyCodeHBN = "HBN"

	// CurrencyCodeHEAT identifies the HEAT digital currency
	CurrencyCodeHEAT = "HEAT"

	// CurrencyCodeHMQ identifies the Humaniq digital currency
	CurrencyCodeHMQ = "HMQ"

	// CurrencyCodeHPB identifies the High-Performance-Blockchain digital currency
	CurrencyCodeHPB = "HPB"

	// CurrencyCodeHSR identifies the Hshare digital currency
	CurrencyCodeHSR = "HSR"

	// CurrencyCodeHUSH identifies the Hush digital currency
	CurrencyCodeHUSH = "HUSH"

	// CurrencyCodeHVN identifies the Hive digital currency
	CurrencyCodeHVN = "HVN"

	// CurrencyCodeHXX identifies the HexxCoin digital currency
	CurrencyCodeHXX = "HXX"

	// CurrencyCodeICN identifies the ICONOMI digital currency
	CurrencyCodeICN = "ICN"

	// CurrencyCodeICX identifies the ICON digital currency
	CurrencyCodeICX = "ICX"

	// CurrencyCodeIFC identifies the Infinitecoin digital currency
	CurrencyCodeIFC = "IFC"

	// CurrencyCodeIFT identifies the investFeed digital currency
	CurrencyCodeIFT = "IFT"

	// CurrencyCodeIGNIS identifies the Ignis digital currency
	CurrencyCodeIGNIS = "IGNIS"

	// CurrencyCodeINCNT identifies the Incent digital currency
	CurrencyCodeINCNT = "INCNT"

	// CurrencyCodeIND identifies the Indorse-Token digital currency
	CurrencyCodeIND = "IND"

	// CurrencyCodeINF identifies the InfChain digital currency
	CurrencyCodeINF = "INF"

	// CurrencyCodeINK identifies the Ink digital currency
	CurrencyCodeINK = "INK"

	// CurrencyCodeINS identifies the INS-Ecosystem digital currency
	CurrencyCodeINS = "INS"

	// CurrencyCodeINSTAR identifies the Insights-Network digital currency
	CurrencyCodeINSTAR = "INSTAR"

	// CurrencyCodeINT identifies the Internet-Node-Token digital currency
	CurrencyCodeINT = "INT"

	// CurrencyCodeINXT identifies the Internxt digital currency
	CurrencyCodeINXT = "INXT"

	// CurrencyCodeIOC identifies the IOCoin digital currency
	CurrencyCodeIOC = "IOC"

	// CurrencyCodeION identifies the ION digital currency
	CurrencyCodeION = "ION"

	// CurrencyCodeIOP identifies the Internet-of-People digital currency
	CurrencyCodeIOP = "IOP"

	// CurrencyCodeIOST identifies the IOStoken digital currency
	CurrencyCodeIOST = "IOST"

	// CurrencyCodeIOTA identifies the IOTA digital currency
	CurrencyCodeIOTA = "IOTA"

	// CurrencyCodeIOTX identifies the IoTeX digital currency
	CurrencyCodeIOTX = "IOTX"

	// CurrencyCodeIQT identifies the Iquant-Chain digital currency
	CurrencyCodeIQT = "IQT"

	// CurrencyCodeITC identifies the IoT-Chain digital currency
	CurrencyCodeITC = "ITC"

	// CurrencyCodeIXC identifies the iXcoin digital currency
	CurrencyCodeIXC = "IXC"

	// CurrencyCodeIXT identifies the InsureX digital currency
	CurrencyCodeIXT = "IXT"

	// CurrencyCodeJ8T identifies the JET8 digital currency
	CurrencyCodeJ8T = "J8T"

	// CurrencyCodeJNT identifies the Jibrel-Network digital currency
	CurrencyCodeJNT = "JNT"

	// CurrencyCodeKCS identifies the KuCoin digital currency
	CurrencyCodeKCS = "KCS"

	// CurrencyCodeKICK identifies the KickCoin digital currency
	CurrencyCodeKICK = "KICK"

	// CurrencyCodeKIN identifies the KIN digital currency
	CurrencyCodeKIN = "KIN"

	// CurrencyCodeKMD identifies the Komodo digital currency
	CurrencyCodeKMD = "KMD"

	// CurrencyCodeKNC identifies the Kyber-Network digital currency
	CurrencyCodeKNC = "KNC"

	// CurrencyCodeKORE identifies the KoreCoin digital currency
	CurrencyCodeKORE = "KORE"

	// CurrencyCodeLBC identifies the LBRY-Credits digital currency
	CurrencyCodeLBC = "LBC"

	// CurrencyCodeLCC identifies the Litecoin-Cash digital currency
	CurrencyCodeLCC = "LCC"

	// CurrencyCodeLEND identifies the EthLend digital currency
	CurrencyCodeLEND = "LEND"

	// CurrencyCodeLEV identifies the Leverj digital currency
	CurrencyCodeLEV = "LEV"

	// CurrencyCodeLGD identifies the Legends-Room digital currency
	CurrencyCodeLGD = "LGD"

	// CurrencyCodeLINDA identifies the Linda digital currency
	CurrencyCodeLINDA = "LINDA"

	// CurrencyCodeLINK identifies the ChainLink digital currency
	CurrencyCodeLINK = "LINK"

	// CurrencyCodeLKK identifies the Lykke digital currency
	CurrencyCodeLKK = "LKK"

	// CurrencyCodeLMC identifies the LoMoCoin digital currency
	CurrencyCodeLMC = "LMC"

	// CurrencyCodeLOCI identifies the LOCIcoin digital currency
	CurrencyCodeLOCI = "LOCI"

	// CurrencyCodeLOOM identifies the Loom-Token digital currency
	CurrencyCodeLOOM = "LOOM"

	// CurrencyCodeLRC identifies the Loopring digital currency
	CurrencyCodeLRC = "LRC"

	// CurrencyCodeLSK identifies the Lisk digital currency
	CurrencyCodeLSK = "LSK"

	// CurrencyCodeLTC identifies the Litecoin digital currency
	CurrencyCodeLTC = "LTC"

	// CurrencyCodeLUN identifies the Lunyr digital currency
	CurrencyCodeLUN = "LUN"

	// CurrencyCodeMAID identifies the MaidSafeCoin digital currency
	CurrencyCodeMAID = "MAID"

	// CurrencyCodeMANA identifies the Decentraland digital currency
	CurrencyCodeMANA = "MANA"

	// CurrencyCodeMAX identifies the Maxcoin digital currency
	CurrencyCodeMAX = "MAX"

	// CurrencyCodeMBRS identifies the Embers digital currency
	CurrencyCodeMBRS = "MBRS"

	// CurrencyCodeMCAP identifies the MCAP digital currency
	CurrencyCodeMCAP = "MCAP"

	// CurrencyCodeMCO identifies the Monaco digital currency
	CurrencyCodeMCO = "MCO"

	// CurrencyCodeMDA identifies the Moeda-Loyalty-Points digital currency
	CurrencyCodeMDA = "MDA"

	// CurrencyCodeMEC identifies the Megacoin digital currency
	CurrencyCodeMEC = "MEC"

	// CurrencyCodeMED identifies the MediBlock digital currency
	CurrencyCodeMED = "MED"

	// CurrencyCodeMEME identifies the Memetic digital currency
	CurrencyCodeMEME = "MEME"

	// CurrencyCodeMER identifies the Mercury digital currency
	CurrencyCodeMER = "MER"

	// CurrencyCodeMGC identifies the MergeCoin digital currency
	CurrencyCodeMGC = "MGC"

	// CurrencyCodeMGO identifies the MobileGo digital currency
	CurrencyCodeMGO = "MGO"

	// CurrencyCodeMINEX identifies the Minex digital currency
	CurrencyCodeMINEX = "MINEX"

	// CurrencyCodeMINT identifies the Mintcoin digital currency
	CurrencyCodeMINT = "MINT"

	// CurrencyCodeMITH identifies the Mithril digital currency
	CurrencyCodeMITH = "MITH"

	// CurrencyCodeMKR identifies the Maker digital currency
	CurrencyCodeMKR = "MKR"

	// CurrencyCodeMLN identifies the Melon digital currency
	CurrencyCodeMLN = "MLN"

	// CurrencyCodeMNE identifies the Minereum digital currency
	CurrencyCodeMNE = "MNE"

	// CurrencyCodeMNX identifies the MinexCoin digital currency
	CurrencyCodeMNX = "MNX"

	// CurrencyCodeMOD identifies the Modum digital currency
	CurrencyCodeMOD = "MOD"

	// CurrencyCodeMONA identifies the MonaCoin digital currency
	CurrencyCodeMONA = "MONA"

	// CurrencyCodeMRT identifies the Miners-Reward-Token digital currency
	CurrencyCodeMRT = "MRT"

	// CurrencyCodeMSP identifies the Mothership digital currency
	CurrencyCodeMSP = "MSP"

	// CurrencyCodeMTH identifies the Monetha digital currency
	CurrencyCodeMTH = "MTH"

	// CurrencyCodeMTN identifies the MedToken digital currency
	CurrencyCodeMTN = "MTN"

	// CurrencyCodeMUE identifies the MonetaryUnit digital currency
	CurrencyCodeMUE = "MUE"

	// CurrencyCodeMUSIC identifies the Musicoin digital currency
	CurrencyCodeMUSIC = "MUSIC"

	// CurrencyCodeMYB identifies the MyBit-Token digital currency
	CurrencyCodeMYB = "MYB"

	// CurrencyCodeMYST identifies the Mysterium digital currency
	CurrencyCodeMYST = "MYST"

	// CurrencyCodeMZC identifies the Mazacoin digital currency
	CurrencyCodeMZC = "MZC"

	// CurrencyCodeNAMO identifies the Namocoin digital currency
	CurrencyCodeNAMO = "NAMO"

	// CurrencyCodeNANO identifies the Nano digital currency
	CurrencyCodeNANO = "NANO"

	// CurrencyCodeNAS identifies the Nebulas-Token digital currency
	CurrencyCodeNAS = "NAS"

	// CurrencyCodeNAV identifies the Nav-Coin digital currency
	CurrencyCodeNAV = "NAV"

	// CurrencyCodeNBT identifies the NuBits digital currency
	CurrencyCodeNBT = "NBT"

	// CurrencyCodeNCASH identifies the Nucleus-Vision digital currency
	CurrencyCodeNCASH = "NCASH"

	// CurrencyCodeNDC identifies the NeverDie-Coin digital currency
	CurrencyCodeNDC = "NDC"

	// CurrencyCodeNEBL identifies the Neblio digital currency
	CurrencyCodeNEBL = "NEBL"

	// CurrencyCodeNEO identifies the NEO digital currency
	CurrencyCodeNEO = "NEO"

	// CurrencyCodeNEOS identifies the NeosCoin digital currency
	CurrencyCodeNEOS = "NEOS"

	// CurrencyCodeNET identifies the Nimiq digital currency
	CurrencyCodeNET = "NET"

	// CurrencyCodeNLC2 identifies the NoLimitCoin digital currency
	CurrencyCodeNLC2 = "NLC2"

	// CurrencyCodeNLG identifies the Gulden digital currency
	CurrencyCodeNLG = "NLG"

	// CurrencyCodeNMC identifies the Namecoin digital currency
	CurrencyCodeNMC = "NMC"

	// CurrencyCodeNMR identifies the Numeraire digital currency
	CurrencyCodeNMR = "NMR"

	// CurrencyCodeNOBL identifies the NobleCoin digital currency
	CurrencyCodeNOBL = "NOBL"

	// CurrencyCodeNOTE identifies the DNotes digital currency
	CurrencyCodeNOTE = "NOTE"

	// CurrencyCodeNPXS identifies the Pundi-X-Token digital currency
	CurrencyCodeNPXS = "NPXS"

	// CurrencyCodeNSR identifies the NuShares digital currency
	CurrencyCodeNSR = "NSR"

	// CurrencyCodeNTO identifies the Fujinto digital currency
	CurrencyCodeNTO = "NTO"

	// CurrencyCodeNULS identifies the Nuls digital currency
	CurrencyCodeNULS = "NULS"

	// CurrencyCodeNVC identifies the Novacoin digital currency
	CurrencyCodeNVC = "NVC"

	// CurrencyCodeNXC identifies the Nexium digital currency
	CurrencyCodeNXC = "NXC"

	// CurrencyCodeNXS identifies the Nexus digital currency
	CurrencyCodeNXS = "NXS"

	// CurrencyCodeNXT identifies the Nxt digital currency
	CurrencyCodeNXT = "NXT"

	// CurrencyCodeOAX identifies the openANX digital currency
	CurrencyCodeOAX = "OAX"

	// CurrencyCodeOBITS identifies the Obits digital currency
	CurrencyCodeOBITS = "OBITS"

	// CurrencyCodeOCL identifies the Oceanlab digital currency
	CurrencyCodeOCL = "OCL"

	// CurrencyCodeOCN identifies the Odyssey digital currency
	CurrencyCodeOCN = "OCN"

	// CurrencyCodeODEM identifies the ODEM digital currency
	CurrencyCodeODEM = "ODEM"

	// CurrencyCodeODN identifies the Obsidian digital currency
	CurrencyCodeODN = "ODN"

	// CurrencyCodeOF identifies the OFCOIN digital currency
	CurrencyCodeOF = "OF"

	// CurrencyCodeOK identifies the OKCash digital currency
	CurrencyCodeOK = "OK"

	// CurrencyCodeOMG identifies the OmiseGo digital currency
	CurrencyCodeOMG = "OMG"

	// CurrencyCodeOMNI identifies the Omni digital currency
	CurrencyCodeOMNI = "OMNI"

	// CurrencyCodeONION identifies the DeepOnion digital currency
	CurrencyCodeONION = "ONION"

	// CurrencyCodeONT identifies the Ontology digital currency
	CurrencyCodeONT = "ONT"

	// CurrencyCodeOPT identifies the Opus digital currency
	CurrencyCodeOPT = "OPT"

	// CurrencyCodeOST identifies the Simple-Token digital currency
	CurrencyCodeOST = "OST"

	// CurrencyCodePART identifies the Particl digital currency
	CurrencyCodePART = "PART"

	// CurrencyCodePASC identifies the PascalCoin digital currency
	CurrencyCodePASC = "PASC"

	// CurrencyCodePAY identifies the TenX digital currency
	CurrencyCodePAY = "PAY"

	// CurrencyCodePBL identifies the Pebbles digital currency
	CurrencyCodePBL = "PBL"

	// CurrencyCodePBT identifies the Primalbase-Token digital currency
	CurrencyCodePBT = "PBT"

	// CurrencyCodePFR identifies the Payfair digital currency
	CurrencyCodePFR = "PFR"

	// CurrencyCodePING identifies the CryptoPing digital currency
	CurrencyCodePING = "PING"

	// CurrencyCodePINK identifies the Pinkcoin digital currency
	CurrencyCodePINK = "PINK"

	// CurrencyCodePIVX identifies the PIVX digital currency
	CurrencyCodePIVX = "PIVX"

	// CurrencyCodePIX identifies the Lampix digital currency
	CurrencyCodePIX = "PIX"

	// CurrencyCodePLBT identifies the Polybius digital currency
	CurrencyCodePLBT = "PLBT"

	// CurrencyCodePLR identifies the Pillar digital currency
	CurrencyCodePLR = "PLR"

	// CurrencyCodePLU identifies the Pluton digital currency
	CurrencyCodePLU = "PLU"

	// CurrencyCodePOA identifies the POA-Network digital currency
	CurrencyCodePOA = "POA"

	// CurrencyCodePOE identifies the Poet digital currency
	CurrencyCodePOE = "POE"

	// CurrencyCodePOLY identifies the Polymath digital currency
	CurrencyCodePOLY = "POLY"

	// CurrencyCodePOSW identifies the PoSW-Coin digital currency
	CurrencyCodePOSW = "POSW"

	// CurrencyCodePOT identifies the PotCoin digital currency
	CurrencyCodePOT = "POT"

	// CurrencyCodePOWR identifies the Power-Ledger digital currency
	CurrencyCodePOWR = "POWR"

	// CurrencyCodePPC identifies the Peercoin digital currency
	CurrencyCodePPC = "PPC"

	// CurrencyCodePPT identifies the Populous digital currency
	CurrencyCodePPT = "PPT"

	// CurrencyCodePPY identifies the Peerplays digital currency
	CurrencyCodePPY = "PPY"

	// CurrencyCodePRG identifies the Paragon-Coin digital currency
	CurrencyCodePRG = "PRG"

	// CurrencyCodePRL identifies the Oyster-Pearl digital currency
	CurrencyCodePRL = "PRL"

	// CurrencyCodePRO identifies the Propy digital currency
	CurrencyCodePRO = "PRO"

	// CurrencyCodePST identifies the Primas digital currency
	CurrencyCodePST = "PST"

	// CurrencyCodePTC identifies the Pesetacoin digital currency
	CurrencyCodePTC = "PTC"

	// CurrencyCodePTOY identifies the Patientory digital currency
	CurrencyCodePTOY = "PTOY"

	// CurrencyCodePURA identifies the Pura digital currency
	CurrencyCodePURA = "PURA"

	// CurrencyCodeQASH identifies the QASH digital currency
	CurrencyCodeQASH = "QASH"

	// CurrencyCodeQAU identifies the Quantum digital currency
	CurrencyCodeQAU = "QAU"

	// CurrencyCodeQLC identifies the Qlink digital currency
	CurrencyCodeQLC = "QLC"

	// CurrencyCodeQRK identifies the Quark digital currency
	CurrencyCodeQRK = "QRK"

	// CurrencyCodeQRL identifies the Quantum-Resistant-Ledger digital currency
	CurrencyCodeQRL = "QRL"

	// CurrencyCodeQSP identifies the Quantstamp digital currency
	CurrencyCodeQSP = "QSP"

	// CurrencyCodeQTL identifies the Quatloo digital currency
	CurrencyCodeQTL = "QTL"

	// CurrencyCodeQTUM identifies the Qtum digital currency
	CurrencyCodeQTUM = "QTUM"

	// CurrencyCodeQWARK identifies the Qwark digital currency
	CurrencyCodeQWARK = "QWARK"

	// CurrencyCodeR identifies the Revain digital currency
	CurrencyCodeR = "R"

	// CurrencyCodeRADS identifies the Radium digital currency
	CurrencyCodeRADS = "RADS"

	// CurrencyCodeRAIN identifies the Condensate digital currency
	CurrencyCodeRAIN = "RAIN"

	// CurrencyCodeRBIES identifies the Rubies digital currency
	CurrencyCodeRBIES = "RBIES"

	// CurrencyCodeRBX identifies the Ripto-Bux digital currency
	CurrencyCodeRBX = "RBX"

	// CurrencyCodeRBY identifies the RubyCoin digital currency
	CurrencyCodeRBY = "RBY"

	// CurrencyCodeRCN identifies the Ripio-Credit-Network digital currency
	CurrencyCodeRCN = "RCN"

	// CurrencyCodeRDD identifies the ReddCoin digital currency
	CurrencyCodeRDD = "RDD"

	// CurrencyCodeRDN identifies the Raiden-Network-Token digital currency
	CurrencyCodeRDN = "RDN"

	// CurrencyCodeREC identifies the Regalcoin digital currency
	CurrencyCodeREC = "REC"

	// CurrencyCodeRED identifies the Redcoin digital currency
	CurrencyCodeRED = "RED"

	// CurrencyCodeREP identifies the Augur digital currency
	CurrencyCodeREP = "REP"

	// CurrencyCodeREQ identifies the Request-Network digital currency
	CurrencyCodeREQ = "REQ"

	// CurrencyCodeRHOC identifies the RChain digital currency
	CurrencyCodeRHOC = "RHOC"

	// CurrencyCodeRIC identifies the Riecoin digital currency
	CurrencyCodeRIC = "RIC"

	// CurrencyCodeRISE identifies the Rise digital currency
	CurrencyCodeRISE = "RISE"

	// CurrencyCodeRLC identifies the RLC-Token digital currency
	CurrencyCodeRLC = "RLC"

	// CurrencyCodeRLT identifies the RouletteToken digital currency
	CurrencyCodeRLT = "RLT"

	// CurrencyCodeRPX identifies the Red-Pulse digital currency
	CurrencyCodeRPX = "RPX"

	// CurrencyCodeRRT identifies the Recovery-Right-Tokens digital currency
	CurrencyCodeRRT = "RRT"

	// CurrencyCodeRUFF identifies the Ruff digital currency
	CurrencyCodeRUFF = "RUFF"

	// CurrencyCodeRUP identifies the Rupee digital currency
	CurrencyCodeRUP = "RUP"

	// CurrencyCodeRVT identifies the Rivetz digital currency
	CurrencyCodeRVT = "RVT"

	// CurrencyCodeSAFEX identifies the SafeExchangeCoin digital currency
	CurrencyCodeSAFEX = "SAFEX"

	// CurrencyCodeSALT identifies the Salt digital currency
	CurrencyCodeSALT = "SALT"

	// CurrencyCodeSAN identifies the Santiment-Network-Token digital currency
	CurrencyCodeSAN = "SAN"

	// CurrencyCodeSBTC identifies the Super-Bitcoin digital currency
	CurrencyCodeSBTC = "SBTC"

	// CurrencyCodeSC identifies the Siacoin digital currency
	CurrencyCodeSC = "SC"

	// CurrencyCodeSEELE identifies the Seele digital currency
	CurrencyCodeSEELE = "SEELE"

	// CurrencyCodeSEQ identifies the Sequence digital currency
	CurrencyCodeSEQ = "SEQ"

	// CurrencyCodeSHIFT identifies the SHIFT digital currency
	CurrencyCodeSHIFT = "SHIFT"

	// CurrencyCodeSIB identifies the SIBCoin digital currency
	CurrencyCodeSIB = "SIB"

	// CurrencyCodeSIGMA identifies the SIGMAcoin digital currency
	CurrencyCodeSIGMA = "SIGMA"

	// CurrencyCodeSIGT identifies the Signatum digital currency
	CurrencyCodeSIGT = "SIGT"

	// CurrencyCodeSJCX identifies the Storjcoin-X digital currency
	CurrencyCodeSJCX = "SJCX"

	// CurrencyCodeSKIN identifies the SkinCoin digital currency
	CurrencyCodeSKIN = "SKIN"

	// CurrencyCodeSKY identifies the Skycoin digital currency
	CurrencyCodeSKY = "SKY"

	// CurrencyCodeSLR identifies the SolarCoin digital currency
	CurrencyCodeSLR = "SLR"

	// CurrencyCodeSLS identifies the SaluS digital currency
	CurrencyCodeSLS = "SLS"

	// CurrencyCodeSMART identifies the SmartCash digital currency
	CurrencyCodeSMART = "SMART"

	// CurrencyCodeSMT identifies the SmartMesh digital currency
	CurrencyCodeSMT = "SMT"

	// CurrencyCodeSNC identifies the SunContract digital currency
	CurrencyCodeSNC = "SNC"

	// CurrencyCodeSNGLS identifies the SingularDTV digital currency
	CurrencyCodeSNGLS = "SNGLS"

	// CurrencyCodeSNM identifies the SONM digital currency
	CurrencyCodeSNM = "SNM"

	// CurrencyCodeSNRG identifies the Synergy digital currency
	CurrencyCodeSNRG = "SNRG"

	// CurrencyCodeSNT identifies the Status-Network-Token digital currency
	CurrencyCodeSNT = "SNT"

	// CurrencyCodeSOC identifies the All-Sports digital currency
	CurrencyCodeSOC = "SOC"

	// CurrencyCodeSOUL identifies the Phantasma digital currency
	CurrencyCodeSOUL = "SOUL"

	// CurrencyCodeSPANK identifies the SpankChain digital currency
	CurrencyCodeSPANK = "SPANK"

	// CurrencyCodeSPC identifies the SpaceChain digital currency
	CurrencyCodeSPC = "SPC"

	// CurrencyCodeSPHR identifies the Sphere digital currency
	CurrencyCodeSPHR = "SPHR"

	// CurrencyCodeSPR identifies the SpreadCoin digital currency
	CurrencyCodeSPR = "SPR"

	// CurrencyCodeSNX identifies the Synthetix-Network-Token digital currency
	CurrencyCodeSNX = "SNX"

	// CurrencyCodeSRN identifies the Sirin-Labs-Token digital currency
	CurrencyCodeSRN = "SRN"

	// CurrencyCodeSTART identifies the Startcoin digital currency
	CurrencyCodeSTART = "START"

	// CurrencyCodeSTEEM identifies the Steem digital currency
	CurrencyCodeSTEEM = "STEEM"

	// CurrencyCodeSTK identifies the STK-Token digital currency
	CurrencyCodeSTK = "STK"

	// CurrencyCodeSTORJ identifies the Storj digital currency
	CurrencyCodeSTORJ = "STORJ"

	// CurrencyCodeSTORM identifies the Storm digital currency
	CurrencyCodeSTORM = "STORM"

	// CurrencyCodeSTQ identifies the Storiqa digital currency
	CurrencyCodeSTQ = "STQ"

	// CurrencyCodeSTRAT identifies the Stratis digital currency
	CurrencyCodeSTRAT = "STRAT"

	// CurrencyCodeSTX identifies the Stox digital currency
	CurrencyCodeSTX = "STX"

	// CurrencyCodeSUB identifies the Substratum digital currency
	CurrencyCodeSUB = "SUB"

	// CurrencyCodeSWFTC identifies the SwftCoin digital currency
	CurrencyCodeSWFTC = "SWFTC"

	// CurrencyCodeSWIFT identifies the Bitswift digital currency
	CurrencyCodeSWIFT = "SWIFT"

	// CurrencyCodeSWT identifies the Swarm-City digital currency
	CurrencyCodeSWT = "SWT"

	// CurrencyCodeSYNX identifies the Syndicate digital currency
	CurrencyCodeSYNX = "SYNX"

	// CurrencyCodeSYS identifies the SysCoin digital currency
	CurrencyCodeSYS = "SYS"

	// CurrencyCodeTAAS identifies the Taas digital currency
	CurrencyCodeTAAS = "TAAS"

	// CurrencyCodeTAU identifies the Lamden digital currency
	CurrencyCodeTAU = "TAU"

	// CurrencyCodeTCC identifies the The-ChampCoin digital currency
	CurrencyCodeTCC = "TCC"

	// CurrencyCodeTFL identifies the True-Flip digital currency
	CurrencyCodeTFL = "TFL"

	// CurrencyCodeTHC identifies the HempCoin digital currency
	CurrencyCodeTHC = "THC"

	// CurrencyCodeTHETA identifies the Theta-Token digital currency
	CurrencyCodeTHETA = "THETA"

	// CurrencyCodeTIME identifies the Time digital currency
	CurrencyCodeTIME = "TIME"

	// CurrencyCodeTIX identifies the Blocktix digital currency
	CurrencyCodeTIX = "TIX"

	// CurrencyCodeTKN identifies the TokenCard digital currency
	CurrencyCodeTKN = "TKN"

	// CurrencyCodeTKR identifies the Trackr digital currency
	CurrencyCodeTKR = "TKR"

	// CurrencyCodeTKS identifies the Tokes digital currency
	CurrencyCodeTKS = "TKS"

	// CurrencyCodeTNB identifies the Time-New-Bank digital currency
	CurrencyCodeTNB = "TNB"

	// CurrencyCodeTNT identifies the Tierion digital currency
	CurrencyCodeTNT = "TNT"

	// CurrencyCodeTOA identifies the ToaCoin digital currency
	CurrencyCodeTOA = "TOA"

	// CurrencyCodeTRAC identifies the OriginTrail digital currency
	CurrencyCodeTRAC = "TRAC"

	// CurrencyCodeTRC identifies the Terracoin digital currency
	CurrencyCodeTRC = "TRC"

	// CurrencyCodeTRCT identifies the Tracto digital currency
	CurrencyCodeTRCT = "TRCT"

	// CurrencyCodeTRIG identifies the Triggers digital currency
	CurrencyCodeTRIG = "TRIG"

	// CurrencyCodeTRST identifies the Trustcoin digital currency
	CurrencyCodeTRST = "TRST"

	// CurrencyCodeTRUE identifies the TrueChain digital currency
	CurrencyCodeTRUE = "TRUE"

	// CurrencyCodeTRUST identifies the TrustPlus digital currency
	CurrencyCodeTRUST = "TRUST"

	// CurrencyCodeTRX identifies the Tronix digital currency
	CurrencyCodeTRX = "TRX"

	// CurrencyCodeTUSD identifies the TrueUSD digital currency
	CurrencyCodeTUSD = "TUSD"

	// CurrencyCodeTX identifies the TransferCoin digital currency
	CurrencyCodeTX = "TX"

	// CurrencyCodeUBQ identifies the Ubiq digital currency
	CurrencyCodeUBQ = "UBQ"

	// CurrencyCodeUKG identifies the UnikoinGold digital currency
	CurrencyCodeUKG = "UKG"

	// CurrencyCodeULA identifies the Ulatech digital currency
	CurrencyCodeULA = "ULA"

	// CurrencyCodeUNB identifies the UnbreakableCoin digital currency
	CurrencyCodeUNB = "UNB"

	// CurrencyCodeUNITY identifies the SuperNET digital currency
	CurrencyCodeUNITY = "UNITY"

	// CurrencyCodeUNO identifies the Unobtanium digital currency
	CurrencyCodeUNO = "UNO"

	// CurrencyCodeUNY identifies the Unity-Ingot digital currency
	CurrencyCodeUNY = "UNY"

	// CurrencyCodeUP identifies the UpToken digital currency
	CurrencyCodeUP = "UP"

	// CurrencyCodeURO identifies the Uro digital currency
	CurrencyCodeURO = "URO"

	// CurrencyCodeUSDT identifies the Tether digital currency
	CurrencyCodeUSDT = "USDT"

	// CurrencyCodeUTK identifies the UTrust digital currency
	CurrencyCodeUTK = "UTK"

	// CurrencyCodeVEE identifies the BLOCKv digital currency
	CurrencyCodeVEE = "VEE"

	// CurrencyCodeVEN identifies the VeChain digital currency
	CurrencyCodeVEN = "VEN"

	// CurrencyCodeVERI identifies the Veritaseum digital currency
	CurrencyCodeVERI = "VERI"

	// CurrencyCodeVIA identifies the Viacoin digital currency
	CurrencyCodeVIA = "VIA"

	// CurrencyCodeVIB identifies the Viberate digital currency
	CurrencyCodeVIB = "VIB"

	// CurrencyCodeVIBE identifies the Vibe digital currency
	CurrencyCodeVIBE = "VIBE"

	// CurrencyCodeVIVO identifies the VIVO digital currency
	CurrencyCodeVIVO = "VIVO"

	// CurrencyCodeVOISE identifies the Voise digital currency
	CurrencyCodeVOISE = "VOISE"

	// CurrencyCodeVOX identifies the Voxels digital currency
	CurrencyCodeVOX = "VOX"

	// CurrencyCodeVPN identifies the VPNCoin digital currency
	CurrencyCodeVPN = "VPN"

	// CurrencyCodeVRC identifies the Vericoin digital currency
	CurrencyCodeVRC = "VRC"

	// CurrencyCodeVRM identifies the Verium digital currency
	CurrencyCodeVRM = "VRM"

	// CurrencyCodeVRS identifies the Veros digital currency
	CurrencyCodeVRS = "VRS"

	// CurrencyCodeVSL identifies the vSlice digital currency
	CurrencyCodeVSL = "VSL"

	// CurrencyCodeVTC identifies the Vertcoin digital currency
	CurrencyCodeVTC = "VTC"

	// CurrencyCodeVTR identifies the vTorrent digital currency
	CurrencyCodeVTR = "VTR"

	// CurrencyCodeWABI identifies the WaBi digital currency
	CurrencyCodeWABI = "WABI"

	// CurrencyCodeWAN identifies the Wanchain digital currency
	CurrencyCodeWAN = "WAN"

	// CurrencyCodeWAVES identifies the Waves digital currency
	CurrencyCodeWAVES = "WAVES"

	// CurrencyCodeWAX identifies the Wax-Token digital currency
	CurrencyCodeWAX = "WAX"

	// CurrencyCodeWCT identifies the Waves-Community digital currency
	CurrencyCodeWCT = "WCT"

	// CurrencyCodeWDC identifies the WorldCoin digital currency
	CurrencyCodeWDC = "WDC"

	// CurrencyCodeWGO identifies the WavesGo digital currency
	CurrencyCodeWGO = "WGO"

	// CurrencyCodeWGR identifies the Wagerr digital currency
	CurrencyCodeWGR = "WGR"

	// CurrencyCodeWINGS identifies the Wings digital currency
	CurrencyCodeWINGS = "WINGS"

	// CurrencyCodeWPR identifies the WePower digital currency
	CurrencyCodeWPR = "WPR"

	// CurrencyCodeWTC identifies the Walton digital currency
	CurrencyCodeWTC = "WTC"

	// CurrencyCodeWTT identifies the Giga-Watt-Token digital currency
	CurrencyCodeWTT = "WTT"

	// CurrencyCodeXAS identifies the Asch digital currency
	CurrencyCodeXAS = "XAS"

	// CurrencyCodeXAUR identifies the Xaurum digital currency
	CurrencyCodeXAUR = "XAUR"

	// CurrencyCodeXBC identifies the Bitcoin-Plus digital currency
	CurrencyCodeXBC = "XBC"

	// CurrencyCodeXBY identifies the XtraBYtes digital currency
	CurrencyCodeXBY = "XBY"

	// CurrencyCodeXCN identifies the Cryptonite digital currency
	CurrencyCodeXCN = "XCN"

	// CurrencyCodeXCP identifies the Counterparty digital currency
	CurrencyCodeXCP = "XCP"

	// CurrencyCodeXDN identifies the DigitalNote digital currency
	CurrencyCodeXDN = "XDN"

	// CurrencyCodeXEL identifies the Elastic digital currency
	CurrencyCodeXEL = "XEL"

	// CurrencyCodeXEM identifies the NEM digital currency
	CurrencyCodeXEM = "XEM"

	// CurrencyCodeXHV identifies the Haven-Protocol digital currency
	CurrencyCodeXHV = "XHV"

	// CurrencyCodeXID identifies the Sphere-Identity digital currency
	CurrencyCodeXID = "XID"

	// CurrencyCodeXLM identifies the Stellar digital currency
	CurrencyCodeXLM = "XLM"

	// CurrencyCodeXMG identifies the Magi digital currency
	CurrencyCodeXMG = "XMG"

	// CurrencyCodeXMR identifies the Monero digital currency
	CurrencyCodeXMR = "XMR"

	// CurrencyCodeXMT identifies the Metal digital currency
	CurrencyCodeXMT = "XMT"

	// CurrencyCodeXMY identifies the Myriadcoin digital currency
	CurrencyCodeXMY = "XMY"

	// CurrencyCodeXPM identifies the Primecoin digital currency
	CurrencyCodeXPM = "XPM"

	// CurrencyCodeXRL identifies the Rialto digital currency
	CurrencyCodeXRL = "XRL"

	// CurrencyCodeXRP identifies the Ripple digital currency
	CurrencyCodeXRP = "XRP"

	// CurrencyCodeXSPEC identifies the Spectrecoin digital currency
	CurrencyCodeXSPEC = "XSPEC"

	// CurrencyCodeXST identifies the Stealthcoin digital currency
	CurrencyCodeXST = "XST"

	// CurrencyCodeXTZ identifies the Tezos digital currency
	CurrencyCodeXTZ = "XTZ"

	// CurrencyCodeXUC identifies the Exchange-Union digital currency
	CurrencyCodeXUC = "XUC"

	// CurrencyCodeXVC identifies the Vcash digital currency
	CurrencyCodeXVC = "XVC"

	// CurrencyCodeXVG identifies the Verge digital currency
	CurrencyCodeXVG = "XVG"

	// CurrencyCodeXWC identifies the WhiteCoin digital currency
	CurrencyCodeXWC = "XWC"

	// CurrencyCodeXZC identifies the ZCoin digital currency
	CurrencyCodeXZC = "XZC"

	// CurrencyCodeXZR identifies the ZrCoin digital currency
	CurrencyCodeXZR = "XZR"

	// CurrencyCodeYEE identifies the Yee digital currency
	CurrencyCodeYEE = "YEE"

	// CurrencyCodeYOYOW identifies the YOYOW digital currency
	CurrencyCodeYOYOW = "YOYOW"

	// CurrencyCodeZCC identifies the ZcCoin digital currency
	CurrencyCodeZCC = "ZCC"

	// CurrencyCodeZCL identifies the Zclassic digital currency
	CurrencyCodeZCL = "ZCL"

	// CurrencyCodeZCO identifies the Zebi digital currency
	CurrencyCodeZCO = "ZCO"

	// CurrencyCodeZEC identifies the Zcash digital currency
	CurrencyCodeZEC = "ZEC"

	// CurrencyCodeZEN identifies the ZenCash digital currency
	CurrencyCodeZEN = "ZEN"

	// CurrencyCodeZET identifies the Zetacoin digital currency
	CurrencyCodeZET = "ZET"

	// CurrencyCodeZIL identifies the Zilliqa digital currency
	CurrencyCodeZIL = "ZIL"

	// CurrencyCodeZLA identifies the Zilla digital currency
	CurrencyCodeZLA = "ZLA"

	// CurrencyCodeZRX identifies the 0x digital currency
	CurrencyCodeZRX = "ZRX"
)

// IsPhysical returns true if the currency is a physical one, false otherwise
// (i.e. digital/crypto currencies)
func (c CurrencyCode) IsPhysical() bool {
	switch c {
	case CurrencyCodeAED:
		return true
	case CurrencyCodeAFN:
		return true
	case CurrencyCodeALL:
		return true
	case CurrencyCodeAMD:
		return true
	case CurrencyCodeANG:
		return true
	case CurrencyCodeAOA:
		return true
	case CurrencyCodeARS:
		return true
	case CurrencyCodeAUD:
		return true
	case CurrencyCodeAWG:
		return true
	case CurrencyCodeAZN:
		return true
	case CurrencyCodeBAM:
		return true
	case CurrencyCodeBBD:
		return true
	case CurrencyCodeBDT:
		return true
	case CurrencyCodeBGN:
		return true
	case CurrencyCodeBHD:
		return true
	case CurrencyCodeBIF:
		return true
	case CurrencyCodeBMD:
		return true
	case CurrencyCodeBND:
		return true
	case CurrencyCodeBOB:
		return true
	case CurrencyCodeBRL:
		return true
	case CurrencyCodeBSD:
		return true
	case CurrencyCodeBTN:
		return true
	case CurrencyCodeBWP:
		return true
	case CurrencyCodeBZD:
		return true
	case CurrencyCodeCAD:
		return true
	case CurrencyCodeCDF:
		return true
	case CurrencyCodeCHF:
		return true
	case CurrencyCodeCLF:
		return true
	case CurrencyCodeCLP:
		return true
	case CurrencyCodeCNH:
		return true
	case CurrencyCodeCNY:
		return true
	case CurrencyCodeCOP:
		return true
	case CurrencyCodeCUP:
		return true
	case CurrencyCodeCVE:
		return true
	case CurrencyCodeCZK:
		return true
	case CurrencyCodeDJF:
		return true
	case CurrencyCodeDKK:
		return true
	case CurrencyCodeDOP:
		return true
	case CurrencyCodeDZD:
		return true
	case CurrencyCodeEGP:
		return true
	case CurrencyCodeERN:
		return true
	case CurrencyCodeETB:
		return true
	case CurrencyCodeEUR:
		return true
	case CurrencyCodeFJD:
		return true
	case CurrencyCodeFKP:
		return true
	case CurrencyCodeGBP:
		return true
	case CurrencyCodeGEL:
		return true
	case CurrencyCodeGHS:
		return true
	case CurrencyCodeGIP:
		return true
	case CurrencyCodeGMD:
		return true
	case CurrencyCodeGNF:
		return true
	case CurrencyCodeGTQ:
		return true
	case CurrencyCodeGYD:
		return true
	case CurrencyCodeHKD:
		return true
	case CurrencyCodeHNL:
		return true
	case CurrencyCodeHRK:
		return true
	case CurrencyCodeHTG:
		return true
	case CurrencyCodeHUF:
		return true
	case CurrencyCodeIDR:
		return true
	case CurrencyCodeILS:
		return true
	case CurrencyCodeINR:
		return true
	case CurrencyCodeIQD:
		return true
	case CurrencyCodeIRR:
		return true
	case CurrencyCodeISK:
		return true
	case CurrencyCodeJEP:
		return true
	case CurrencyCodeJMD:
		return true
	case CurrencyCodeJOD:
		return true
	case CurrencyCodeJPY:
		return true
	case CurrencyCodeKES:
		return true
	case CurrencyCodeKGS:
		return true
	case CurrencyCodeKHR:
		return true
	case CurrencyCodeKMF:
		return true
	case CurrencyCodeKPW:
		return true
	case CurrencyCodeKRW:
		return true
	case CurrencyCodeKWD:
		return true
	case CurrencyCodeKYD:
		return true
	case CurrencyCodeKZT:
		return true
	case CurrencyCodeLAK:
		return true
	case CurrencyCodeLBP:
		return true
	case CurrencyCodeLKR:
		return true
	case CurrencyCodeLRD:
		return true
	case CurrencyCodeLSL:
		return true
	case CurrencyCodeLYD:
		return true
	case CurrencyCodeMAD:
		return true
	case CurrencyCodeMDL:
		return true
	case CurrencyCodeMGA:
		return true
	case CurrencyCodeMKD:
		return true
	case CurrencyCodeMMK:
		return true
	case CurrencyCodeMNT:
		return true
	case CurrencyCodeMOP:
		return true
	case CurrencyCodeMRO:
		return true
	case CurrencyCodeMRU:
		return true
	case CurrencyCodeMUR:
		return true
	case CurrencyCodeMVR:
		return true
	case CurrencyCodeMWK:
		return true
	case CurrencyCodeMXN:
		return true
	case CurrencyCodeMYR:
		return true
	case CurrencyCodeMZN:
		return true
	case CurrencyCodeNAD:
		return true
	case CurrencyCodeNGN:
		return true
	case CurrencyCodeNOK:
		return true
	case CurrencyCodeNPR:
		return true
	case CurrencyCodeNZD:
		return true
	case CurrencyCodeOMR:
		return true
	case CurrencyCodePAB:
		return true
	case CurrencyCodePEN:
		return true
	case CurrencyCodePGK:
		return true
	case CurrencyCodePHP:
		return true
	case CurrencyCodePKR:
		return true
	case CurrencyCodePLN:
		return true
	case CurrencyCodePYG:
		return true
	case CurrencyCodeQAR:
		return true
	case CurrencyCodeRON:
		return true
	case CurrencyCodeRSD:
		return true
	case CurrencyCodeRUB:
		return true
	case CurrencyCodeRUR:
		return true
	case CurrencyCodeRWF:
		return true
	case CurrencyCodeSAR:
		return true
	case CurrencyCodeSBD:
		return true
	case CurrencyCodeSCR:
		return true
	case CurrencyCodeSDG:
		return true
	case CurrencyCodeSEK:
		return true
	case CurrencyCodeSGD:
		return true
	case CurrencyCodeSHP:
		return true
	case CurrencyCodeSLL:
		return true
	case CurrencyCodeSOS:
		return true
	case CurrencyCodeSRD:
		return true
	case CurrencyCodeSYP:
		return true
	case CurrencyCodeSZL:
		return true
	case CurrencyCodeTHB:
		return true
	case CurrencyCodeTJS:
		return true
	case CurrencyCodeTMT:
		return true
	case CurrencyCodeTND:
		return true
	case CurrencyCodeTOP:
		return true
	case CurrencyCodeTRY:
		return true
	case CurrencyCodeTTD:
		return true
	case CurrencyCodeTWD:
		return true
	case CurrencyCodeTZS:
		return true
	case CurrencyCodeUAH:
		return true
	case CurrencyCodeUGX:
		return true
	case CurrencyCodeUSD:
		return true
	case CurrencyCodeUYU:
		return true
	case CurrencyCodeUZS:
		return true
	case CurrencyCodeVND:
		return true
	case CurrencyCodeVUV:
		return true
	case CurrencyCodeWST:
		return true
	case CurrencyCodeXAF:
		return true
	case CurrencyCodeXAG:
		return true
	case CurrencyCodeXAU:
		return true
	case CurrencyCodeXCD:
		return true
	case CurrencyCodeXDR:
		return true
	case CurrencyCodeXOF:
		return true
	case CurrencyCodeXPF:
		return true
	case CurrencyCodeYER:
		return true
	case CurrencyCodeZAR:
		return true
	case CurrencyCodeZMW:
		return true
	case CurrencyCodeZWL:
		return true
	default:
		return false
	}
}

// Code returns the 3 letter code representing the currency
func (c CurrencyCode) String() string {
	return string(c)
}
