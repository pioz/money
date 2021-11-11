package money

// Currency represents a currency.
type Currency struct {
	// The name of the currency.
	Name string
	// The 3 char identifier of the currency (ISO 4217).
	IsoCode string
	// The symbol of the currency.
	Symbol string
	// Is true if the symbol will be displayed before the amount.
	SymbolFirst bool
	// Number of subunits that compose the unit. For example USD is made of 100
	// cents, so SubunitToUnit is 100.
	SubunitToUnit int
	// Thousands separator.
	ThousandsSeparator rune
	// Decimal mark.
	DecimalMark rune
}

var (
	// From Ruby Money:
	// Money::Currency.each { |c| puts "#{c.iso_code} = Currency{Name: \"#{c.name}\", IsoCode: \"#{c.iso_code}\", Symbol: \"#{c.symbol}\", SymbolFirst: #{c.symbol_first}, SubunitToUnit: #{c.subunit_to_unit}, ThousandsSeparator: '#{c.thousands_separator}', DecimalMark: '#{c.decimal_mark}'}" }; nil

	AED = Currency{Name: "United Arab Emirates Dirham", IsoCode: "AED", Symbol: "د.إ", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	AFN = Currency{Name: "Afghan Afghani", IsoCode: "AFN", Symbol: "؋", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ALL = Currency{Name: "Albanian Lek", IsoCode: "ALL", Symbol: "L", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	AMD = Currency{Name: "Armenian Dram", IsoCode: "AMD", Symbol: "դր.", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ANG = Currency{Name: "Netherlands Antillean Gulden", IsoCode: "ANG", Symbol: "ƒ", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	AOA = Currency{Name: "Angolan Kwanza", IsoCode: "AOA", Symbol: "Kz", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ARS = Currency{Name: "Argentine Peso", IsoCode: "ARS", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	AUD = Currency{Name: "Australian Dollar", IsoCode: "AUD", Symbol: "A$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	AWG = Currency{Name: "Aruban Florin", IsoCode: "AWG", Symbol: "ƒ", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	AZN = Currency{Name: "Azerbaijani Manat", IsoCode: "AZN", Symbol: "₼", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BAM = Currency{Name: "Bosnia and Herzegovina Convertible Mark", IsoCode: "BAM", Symbol: "КМ", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BBD = Currency{Name: "Barbadian Dollar", IsoCode: "BBD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BCH = Currency{Name: "Bitcoin Cash", IsoCode: "BCH", Symbol: "₿", SymbolFirst: false, SubunitToUnit: 100000000, ThousandsSeparator: ',', DecimalMark: '.'}
	BDT = Currency{Name: "Bangladeshi Taka", IsoCode: "BDT", Symbol: "৳", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BGN = Currency{Name: "Bulgarian Lev", IsoCode: "BGN", Symbol: "лв.", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BHD = Currency{Name: "Bahraini Dinar", IsoCode: "BHD", Symbol: "ب.د", SymbolFirst: true, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	BIF = Currency{Name: "Burundian Franc", IsoCode: "BIF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	BMD = Currency{Name: "Bermudian Dollar", IsoCode: "BMD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BND = Currency{Name: "Brunei Dollar", IsoCode: "BND", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BOB = Currency{Name: "Bolivian Boliviano", IsoCode: "BOB", Symbol: "Bs.", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BRL = Currency{Name: "Brazilian Real", IsoCode: "BRL", Symbol: "R$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	BSD = Currency{Name: "Bahamian Dollar", IsoCode: "BSD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BTC = Currency{Name: "Bitcoin", IsoCode: "BTC", Symbol: "₿", SymbolFirst: true, SubunitToUnit: 100000000, ThousandsSeparator: ',', DecimalMark: '.'}
	BTN = Currency{Name: "Bhutanese Ngultrum", IsoCode: "BTN", Symbol: "Nu.", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BWP = Currency{Name: "Botswana Pula", IsoCode: "BWP", Symbol: "P", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	BYN = Currency{Name: "Belarusian Ruble", IsoCode: "BYN", Symbol: "Br", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ' ', DecimalMark: ','}
	BZD = Currency{Name: "Belize Dollar", IsoCode: "BZD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CAD = Currency{Name: "Canadian Dollar", IsoCode: "CAD", Symbol: "C$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CDF = Currency{Name: "Congolese Franc", IsoCode: "CDF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CHF = Currency{Name: "Swiss Franc", IsoCode: "CHF", Symbol: "CHF", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CLF = Currency{Name: "Unidad de Fomento", IsoCode: "CLF", Symbol: "UF", SymbolFirst: true, SubunitToUnit: 10000, ThousandsSeparator: '.', DecimalMark: ','}
	CLP = Currency{Name: "Chilean Peso", IsoCode: "CLP", Symbol: "$", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: '.', DecimalMark: ','}
	CNH = Currency{Name: "Chinese Renminbi Yuan Offshore", IsoCode: "CNH", Symbol: "¥", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CNY = Currency{Name: "Chinese Renminbi Yuan", IsoCode: "CNY", Symbol: "¥", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	COP = Currency{Name: "Colombian Peso", IsoCode: "COP", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	CRC = Currency{Name: "Costa Rican Colón", IsoCode: "CRC", Symbol: "₡", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	CUC = Currency{Name: "Cuban Convertible Peso", IsoCode: "CUC", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CUP = Currency{Name: "Cuban Peso", IsoCode: "CUP", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CVE = Currency{Name: "Cape Verdean Escudo", IsoCode: "CVE", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	CZK = Currency{Name: "Czech Koruna", IsoCode: "CZK", Symbol: "Kč", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ' ', DecimalMark: ','}
	DJF = Currency{Name: "Djiboutian Franc", IsoCode: "DJF", Symbol: "Fdj", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	DKK = Currency{Name: "Danish Krone", IsoCode: "DKK", Symbol: "kr.", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	DOP = Currency{Name: "Dominican Peso", IsoCode: "DOP", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	DZD = Currency{Name: "Algerian Dinar", IsoCode: "DZD", Symbol: "د.ج", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	EEK = Currency{Name: "Estonian Kroon", IsoCode: "EEK", Symbol: "KR", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	EGP = Currency{Name: "Egyptian Pound", IsoCode: "EGP", Symbol: "ج.م", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ERN = Currency{Name: "Eritrean Nakfa", IsoCode: "ERN", Symbol: "Nfk", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ETB = Currency{Name: "Ethiopian Birr", IsoCode: "ETB", Symbol: "Br", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	EUR = Currency{Name: "Euro", IsoCode: "EUR", Symbol: "€", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	FJD = Currency{Name: "Fijian Dollar", IsoCode: "FJD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	FKP = Currency{Name: "Falkland Pound", IsoCode: "FKP", Symbol: "£", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GBP = Currency{Name: "British Pound", IsoCode: "GBP", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GBX = Currency{Name: "British Penny", IsoCode: "GBX", Symbol: "", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	GEL = Currency{Name: "Georgian Lari", IsoCode: "GEL", Symbol: "ლ", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GGP = Currency{Name: "Guernsey Pound", IsoCode: "GGP", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GHS = Currency{Name: "Ghanaian Cedi", IsoCode: "GHS", Symbol: "₵", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GIP = Currency{Name: "Gibraltar Pound", IsoCode: "GIP", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GMD = Currency{Name: "Gambian Dalasi", IsoCode: "GMD", Symbol: "D", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GNF = Currency{Name: "Guinean Franc", IsoCode: "GNF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	GTQ = Currency{Name: "Guatemalan Quetzal", IsoCode: "GTQ", Symbol: "Q", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	GYD = Currency{Name: "Guyanese Dollar", IsoCode: "GYD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	HKD = Currency{Name: "Hong Kong Dollar", IsoCode: "HKD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	HNL = Currency{Name: "Honduran Lempira", IsoCode: "HNL", Symbol: "L", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	HRK = Currency{Name: "Croatian Kuna", IsoCode: "HRK", Symbol: "kn", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	HTG = Currency{Name: "Haitian Gourde", IsoCode: "HTG", Symbol: "G", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	HUF = Currency{Name: "Hungarian Forint", IsoCode: "HUF", Symbol: "Ft", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ' ', DecimalMark: ','}
	IDR = Currency{Name: "Indonesian Rupiah", IsoCode: "IDR", Symbol: "Rp", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	ILS = Currency{Name: "Israeli New Sheqel", IsoCode: "ILS", Symbol: "₪", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	IMP = Currency{Name: "Isle of Man Pound", IsoCode: "IMP", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	INR = Currency{Name: "Indian Rupee", IsoCode: "INR", Symbol: "₹", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	IQD = Currency{Name: "Iraqi Dinar", IsoCode: "IQD", Symbol: "ع.د", SymbolFirst: false, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	IRR = Currency{Name: "Iranian Rial", IsoCode: "IRR", Symbol: "﷼", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ISK = Currency{Name: "Icelandic Króna", IsoCode: "ISK", Symbol: "kr", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: '.', DecimalMark: ','}
	JEP = Currency{Name: "Jersey Pound", IsoCode: "JEP", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	JMD = Currency{Name: "Jamaican Dollar", IsoCode: "JMD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	JOD = Currency{Name: "Jordanian Dinar", IsoCode: "JOD", Symbol: "د.ا", SymbolFirst: true, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	JPY = Currency{Name: "Japanese Yen", IsoCode: "JPY", Symbol: "¥", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	KES = Currency{Name: "Kenyan Shilling", IsoCode: "KES", Symbol: "KSh", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	KGS = Currency{Name: "Kyrgyzstani Som", IsoCode: "KGS", Symbol: "som", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	KHR = Currency{Name: "Cambodian Riel", IsoCode: "KHR", Symbol: "៛", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	KMF = Currency{Name: "Comorian Franc", IsoCode: "KMF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	KPW = Currency{Name: "North Korean Won", IsoCode: "KPW", Symbol: "₩", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	KRW = Currency{Name: "South Korean Won", IsoCode: "KRW", Symbol: "₩", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	KWD = Currency{Name: "Kuwaiti Dinar", IsoCode: "KWD", Symbol: "د.ك", SymbolFirst: true, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	KYD = Currency{Name: "Cayman Islands Dollar", IsoCode: "KYD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	KZT = Currency{Name: "Kazakhstani Tenge", IsoCode: "KZT", Symbol: "₸", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LAK = Currency{Name: "Lao Kip", IsoCode: "LAK", Symbol: "₭", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LBP = Currency{Name: "Lebanese Pound", IsoCode: "LBP", Symbol: "ل.ل", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LKR = Currency{Name: "Sri Lankan Rupee", IsoCode: "LKR", Symbol: "₨", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LRD = Currency{Name: "Liberian Dollar", IsoCode: "LRD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LSL = Currency{Name: "Lesotho Loti", IsoCode: "LSL", Symbol: "L", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LTL = Currency{Name: "Lithuanian Litas", IsoCode: "LTL", Symbol: "Lt", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LVL = Currency{Name: "Latvian Lats", IsoCode: "LVL", Symbol: "Ls", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	LYD = Currency{Name: "Libyan Dinar", IsoCode: "LYD", Symbol: "ل.د", SymbolFirst: false, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	MAD = Currency{Name: "Moroccan Dirham", IsoCode: "MAD", Symbol: "د.م.", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MDL = Currency{Name: "Moldovan Leu", IsoCode: "MDL", Symbol: "L", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MGA = Currency{Name: "Malagasy Ariary", IsoCode: "MGA", Symbol: "Ar", SymbolFirst: true, SubunitToUnit: 5, ThousandsSeparator: ',', DecimalMark: '.'}
	MKD = Currency{Name: "Macedonian Denar", IsoCode: "MKD", Symbol: "ден", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MMK = Currency{Name: "Myanmar Kyat", IsoCode: "MMK", Symbol: "K", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MNT = Currency{Name: "Mongolian Tögrög", IsoCode: "MNT", Symbol: "₮", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MOP = Currency{Name: "Macanese Pataca", IsoCode: "MOP", Symbol: "P", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MRO = Currency{Name: "Mauritanian Ouguiya", IsoCode: "MRO", Symbol: "UM", SymbolFirst: false, SubunitToUnit: 5, ThousandsSeparator: ',', DecimalMark: '.'}
	MTL = Currency{Name: "Maltese Lira", IsoCode: "MTL", Symbol: "₤", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MUR = Currency{Name: "Mauritian Rupee", IsoCode: "MUR", Symbol: "₨", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MVR = Currency{Name: "Maldivian Rufiyaa", IsoCode: "MVR", Symbol: "MVR", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MWK = Currency{Name: "Malawian Kwacha", IsoCode: "MWK", Symbol: "MK", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MXN = Currency{Name: "Mexican Peso", IsoCode: "MXN", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MYR = Currency{Name: "Malaysian Ringgit", IsoCode: "MYR", Symbol: "RM", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	MZN = Currency{Name: "Mozambican Metical", IsoCode: "MZN", Symbol: "MTn", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	NAD = Currency{Name: "Namibian Dollar", IsoCode: "NAD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	NGN = Currency{Name: "Nigerian Naira", IsoCode: "NGN", Symbol: "₦", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	NIO = Currency{Name: "Nicaraguan Córdoba", IsoCode: "NIO", Symbol: "C$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	NOK = Currency{Name: "Norwegian Krone", IsoCode: "NOK", Symbol: "kr", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	NPR = Currency{Name: "Nepalese Rupee", IsoCode: "NPR", Symbol: "₨", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	NZD = Currency{Name: "New Zealand Dollar", IsoCode: "NZD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	OMR = Currency{Name: "Omani Rial", IsoCode: "OMR", Symbol: "ر.ع.", SymbolFirst: true, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	PAB = Currency{Name: "Panamanian Balboa", IsoCode: "PAB", Symbol: "B/.", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	PEN = Currency{Name: "Peruvian Sol", IsoCode: "PEN", Symbol: "S/", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	PGK = Currency{Name: "Papua New Guinean Kina", IsoCode: "PGK", Symbol: "K", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	PHP = Currency{Name: "Philippine Peso", IsoCode: "PHP", Symbol: "₱", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	PKR = Currency{Name: "Pakistani Rupee", IsoCode: "PKR", Symbol: "₨", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	PLN = Currency{Name: "Polish Złoty", IsoCode: "PLN", Symbol: "zł", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ' ', DecimalMark: ','}
	PYG = Currency{Name: "Paraguayan Guaraní", IsoCode: "PYG", Symbol: "₲", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	QAR = Currency{Name: "Qatari Riyal", IsoCode: "QAR", Symbol: "ر.ق", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	RON = Currency{Name: "Romanian Leu", IsoCode: "RON", Symbol: "Lei", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	RSD = Currency{Name: "Serbian Dinar", IsoCode: "RSD", Symbol: "РСД", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	RUB = Currency{Name: "Russian Ruble", IsoCode: "RUB", Symbol: "₽", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	RWF = Currency{Name: "Rwandan Franc", IsoCode: "RWF", Symbol: "FRw", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	SAR = Currency{Name: "Saudi Riyal", IsoCode: "SAR", Symbol: "ر.س", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SBD = Currency{Name: "Solomon Islands Dollar", IsoCode: "SBD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SCR = Currency{Name: "Seychellois Rupee", IsoCode: "SCR", Symbol: "₨", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SDG = Currency{Name: "Sudanese Pound", IsoCode: "SDG", Symbol: "£", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SEK = Currency{Name: "Swedish Krona", IsoCode: "SEK", Symbol: "kr", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ' ', DecimalMark: ','}
	SGD = Currency{Name: "Singapore Dollar", IsoCode: "SGD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SHP = Currency{Name: "Saint Helenian Pound", IsoCode: "SHP", Symbol: "£", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SKK = Currency{Name: "Slovak Koruna", IsoCode: "SKK", Symbol: "Sk", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SLL = Currency{Name: "Sierra Leonean Leone", IsoCode: "SLL", Symbol: "Le", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SOS = Currency{Name: "Somali Shilling", IsoCode: "SOS", Symbol: "Sh", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SRD = Currency{Name: "Surinamese Dollar", IsoCode: "SRD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SSP = Currency{Name: "South Sudanese Pound", IsoCode: "SSP", Symbol: "£", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	STD = Currency{Name: "São Tomé and Príncipe Dobra", IsoCode: "STD", Symbol: "Db", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SVC = Currency{Name: "Salvadoran Colón", IsoCode: "SVC", Symbol: "₡", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SYP = Currency{Name: "Syrian Pound", IsoCode: "SYP", Symbol: "£S", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	SZL = Currency{Name: "Swazi Lilangeni", IsoCode: "SZL", Symbol: "E", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	THB = Currency{Name: "Thai Baht", IsoCode: "THB", Symbol: "฿", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TJS = Currency{Name: "Tajikistani Somoni", IsoCode: "TJS", Symbol: "ЅМ", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TMT = Currency{Name: "Turkmenistani Manat", IsoCode: "TMT", Symbol: "T", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TND = Currency{Name: "Tunisian Dinar", IsoCode: "TND", Symbol: "د.ت", SymbolFirst: false, SubunitToUnit: 1000, ThousandsSeparator: ',', DecimalMark: '.'}
	TOP = Currency{Name: "Tongan Paʻanga", IsoCode: "TOP", Symbol: "T$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TRY = Currency{Name: "Turkish Lira", IsoCode: "TRY", Symbol: "₺", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	TTD = Currency{Name: "Trinidad and Tobago Dollar", IsoCode: "TTD", Symbol: "$", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TWD = Currency{Name: "New Taiwan Dollar", IsoCode: "TWD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	TZS = Currency{Name: "Tanzanian Shilling", IsoCode: "TZS", Symbol: "Sh", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	UAH = Currency{Name: "Ukrainian Hryvnia", IsoCode: "UAH", Symbol: "₴", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	UGX = Currency{Name: "Ugandan Shilling", IsoCode: "UGX", Symbol: "USh", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	USD = Currency{Name: "United States Dollar", IsoCode: "USD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	UYU = Currency{Name: "Uruguayan Peso", IsoCode: "UYU", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	UZS = Currency{Name: "Uzbekistan Som", IsoCode: "UZS", Symbol: "so'm", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	VEF = Currency{Name: "Venezuelan Bolívar", IsoCode: "VEF", Symbol: "Bs.F", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	VES = Currency{Name: "Venezuelan Bolívar Soberano", IsoCode: "VES", Symbol: "Bs", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: '.', DecimalMark: ','}
	VND = Currency{Name: "Vietnamese Đồng", IsoCode: "VND", Symbol: "₫", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: '.', DecimalMark: ','}
	VUV = Currency{Name: "Vanuatu Vatu", IsoCode: "VUV", Symbol: "Vt", SymbolFirst: true, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	WST = Currency{Name: "Samoan Tala", IsoCode: "WST", Symbol: "T", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	XAF = Currency{Name: "Central African Cfa Franc", IsoCode: "XAF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XAG = Currency{Name: "Silver (Troy Ounce)", IsoCode: "XAG", Symbol: "oz t", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XAU = Currency{Name: "Gold (Troy Ounce)", IsoCode: "XAU", Symbol: "oz t", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XBA = Currency{Name: "European Composite Unit", IsoCode: "XBA", Symbol: "", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XBB = Currency{Name: "European Monetary Unit", IsoCode: "XBB", Symbol: "", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XBC = Currency{Name: "European Unit of Account 9", IsoCode: "XBC", Symbol: "", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XBD = Currency{Name: "European Unit of Account 17", IsoCode: "XBD", Symbol: "", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XCD = Currency{Name: "East Caribbean Dollar", IsoCode: "XCD", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	XDR = Currency{Name: "Special Drawing Rights", IsoCode: "XDR", Symbol: "SDR", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XFU = Currency{Name: "UIC Franc", IsoCode: "XFU", Symbol: "", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	XOF = Currency{Name: "West African Cfa Franc", IsoCode: "XOF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XPD = Currency{Name: "Palladium", IsoCode: "XPD", Symbol: "oz t", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XPF = Currency{Name: "Cfp Franc", IsoCode: "XPF", Symbol: "Fr", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XPT = Currency{Name: "Platinum", IsoCode: "XPT", Symbol: "oz t", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	XTS = Currency{Name: "Codes specifically reserved for testing purposes", IsoCode: "XTS", Symbol: "", SymbolFirst: false, SubunitToUnit: 1, ThousandsSeparator: ',', DecimalMark: '.'}
	YER = Currency{Name: "Yemeni Rial", IsoCode: "YER", Symbol: "﷼", SymbolFirst: false, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ZAR = Currency{Name: "South African Rand", IsoCode: "ZAR", Symbol: "R", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ZMW = Currency{Name: "Zambian Kwacha", IsoCode: "ZMW", Symbol: "K", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
	ZWL = Currency{Name: "Zimbabwean Dollar", IsoCode: "ZWL", Symbol: "$", SymbolFirst: true, SubunitToUnit: 100, ThousandsSeparator: ',', DecimalMark: '.'}
)

// A slice with all pre-defined currencies.
var AllCurrencies []Currency = []Currency{
	AED,
	AFN,
	ALL,
	AMD,
	ANG,
	AOA,
	ARS,
	AUD,
	AWG,
	AZN,
	BAM,
	BBD,
	BCH,
	BDT,
	BGN,
	BHD,
	BIF,
	BMD,
	BND,
	BOB,
	BRL,
	BSD,
	BTC,
	BTN,
	BWP,
	BYN,
	BZD,
	CAD,
	CDF,
	CHF,
	CLF,
	CLP,
	CNH,
	CNY,
	COP,
	CRC,
	CUC,
	CUP,
	CVE,
	CZK,
	DJF,
	DKK,
	DOP,
	DZD,
	EEK,
	EGP,
	ERN,
	ETB,
	EUR,
	FJD,
	FKP,
	GBP,
	GBX,
	GEL,
	GGP,
	GHS,
	GIP,
	GMD,
	GNF,
	GTQ,
	GYD,
	HKD,
	HNL,
	HRK,
	HTG,
	HUF,
	IDR,
	ILS,
	IMP,
	INR,
	IQD,
	IRR,
	ISK,
	JEP,
	JMD,
	JOD,
	JPY,
	KES,
	KGS,
	KHR,
	KMF,
	KPW,
	KRW,
	KWD,
	KYD,
	KZT,
	LAK,
	LBP,
	LKR,
	LRD,
	LSL,
	LTL,
	LVL,
	LYD,
	MAD,
	MDL,
	MGA,
	MKD,
	MMK,
	MNT,
	MOP,
	MRO,
	MTL,
	MUR,
	MVR,
	MWK,
	MXN,
	MYR,
	MZN,
	NAD,
	NGN,
	NIO,
	NOK,
	NPR,
	NZD,
	OMR,
	PAB,
	PEN,
	PGK,
	PHP,
	PKR,
	PLN,
	PYG,
	QAR,
	RON,
	RSD,
	RUB,
	RWF,
	SAR,
	SBD,
	SCR,
	SDG,
	SEK,
	SGD,
	SHP,
	SKK,
	SLL,
	SOS,
	SRD,
	SSP,
	STD,
	SVC,
	SYP,
	SZL,
	THB,
	TJS,
	TMT,
	TND,
	TOP,
	TRY,
	TTD,
	TWD,
	TZS,
	UAH,
	UGX,
	USD,
	UYU,
	UZS,
	VEF,
	VES,
	VND,
	VUV,
	WST,
	XAF,
	XAG,
	XAU,
	XBA,
	XBB,
	XBC,
	XBD,
	XCD,
	XDR,
	XFU,
	XOF,
	XPD,
	XPF,
	XPT,
	XTS,
	YER,
	ZAR,
	ZMW,
	ZWL,
}
