package notification

type Country struct {
	Name     string
	DialCode string
	Code     string
}

// DialCode example: +351
type DialCode string

var by_alpha2 map[string]Country

func init() {
	by_alpha2 = map[string]Country{
		"AF": Country{Name: "Afghanistan", DialCode: "+93"},
		"AL": Country{Name: "Albania", DialCode: "+355"},
		"DZ": Country{Name: "Algeria", DialCode: "+213"},
		"AS": Country{Name: "AmericanSamoa", DialCode: "+1 684"},
		"AD": Country{Name: "Andorra", DialCode: "+376"},
		"AO": Country{Name: "Angola", DialCode: "+244"},
		"AI": Country{Name: "Anguilla", DialCode: "+1 264"},
		"AG": Country{Name: "Antigua and Barbuda", DialCode: "+1268"},
		"AR": Country{Name: "Argentina", DialCode: "+54"},
		"AM": Country{Name: "Armenia", DialCode: "+374"},
		"AW": Country{Name: "Aruba", DialCode: "+297"},
		"AU": Country{Name: "Australia", DialCode: "+61"},
		"AT": Country{Name: "Austria", DialCode: "+43"},
		"AZ": Country{Name: "Azerbaijan", DialCode: "+994"},
		"BS": Country{Name: "Bahamas", DialCode: "+1 242"},
		"BH": Country{Name: "Bahrain", DialCode: "+973"},
		"BD": Country{Name: "Bangladesh", DialCode: "+880"},
		"BB": Country{Name: "Barbados", DialCode: "+1 246"},
		"BY": Country{Name: "Belarus", DialCode: "+375"},
		"BE": Country{Name: "Belgium", DialCode: "+32"},
		"BZ": Country{Name: "Belize", DialCode: "+501"},
		"BJ": Country{Name: "Benin", DialCode: "+229"},
		"BM": Country{Name: "Bermuda", DialCode: "+1 441"},
		"BT": Country{Name: "Bhutan", DialCode: "+975"},
		"BA": Country{Name: "Bosnia and Herzegovina", DialCode: "+387"},
		"BW": Country{Name: "Botswana", DialCode: "+267"},
		"BR": Country{Name: "Brazil", DialCode: "+55"},
		"IO": Country{Name: "British Indian Ocean Territory", DialCode: "+246"},
		"BG": Country{Name: "Bulgaria", DialCode: "+359"},
		"BF": Country{Name: "Burkina Faso", DialCode: "+226"},
		"BI": Country{Name: "Burundi", DialCode: "+257"},
		"KH": Country{Name: "Cambodia", DialCode: "+855"},
		"CM": Country{Name: "Cameroon", DialCode: "+237"},
		"CA": Country{Name: "Canada", DialCode: "+1"},
		"CV": Country{Name: "Cape Verde", DialCode: "+238"},
		"KY": Country{Name: "Cayman Islands", DialCode: "+ 345"},
		"CF": Country{Name: "Central African Republic", DialCode: "+236"},
		"TD": Country{Name: "Chad", DialCode: "+235"},
		"CL": Country{Name: "Chile", DialCode: "+56"},
		"CN": Country{Name: "China", DialCode: "+86"},
		"CX": Country{Name: "Christmas Island", DialCode: "+61"},
		"CO": Country{Name: "Colombia", DialCode: "+57"},
		"KM": Country{Name: "Comoros", DialCode: "+269"},
		"CG": Country{Name: "Congo", DialCode: "+242"},
		"CK": Country{Name: "Cook Islands", DialCode: "+682"},
		"CR": Country{Name: "Costa Rica", DialCode: "+506"},
		"HR": Country{Name: "Croatia", DialCode: "+385"},
		"CU": Country{Name: "Cuba", DialCode: "+53"},
		"CY": Country{Name: "Cyprus", DialCode: "+537"},
		"CZ": Country{Name: "Czech Republic", DialCode: "+420"},
		"DK": Country{Name: "Denmark", DialCode: "+45"},
		"DJ": Country{Name: "Djibouti", DialCode: "+253"},
		"DM": Country{Name: "Dominica", DialCode: "+1 767"},
		"DO": Country{Name: "Dominican Republic", DialCode: "+1 849"},
		"EC": Country{Name: "Ecuador", DialCode: "+593"},
		"EG": Country{Name: "Egypt", DialCode: "+20"},
		"SV": Country{Name: "El Salvador", DialCode: "+503"},
		"GQ": Country{Name: "Equatorial Guinea", DialCode: "+240"},
		"ER": Country{Name: "Eritrea", DialCode: "+291"},
		"EE": Country{Name: "Estonia", DialCode: "+372"},
		"ET": Country{Name: "Ethiopia", DialCode: "+251"},
		"FO": Country{Name: "Faroe Islands", DialCode: "+298"},
		"FJ": Country{Name: "Fiji", DialCode: "+679"},
		"FI": Country{Name: "Finland", DialCode: "+358"},
		"FR": Country{Name: "France", DialCode: "+33"},
		"GF": Country{Name: "French Guiana", DialCode: "+594"},
		"PF": Country{Name: "French Polynesia", DialCode: "+689"},
		"GA": Country{Name: "Gabon", DialCode: "+241"},
		"GM": Country{Name: "Gambia", DialCode: "+220"},
		"GE": Country{Name: "Georgia", DialCode: "+995"},
		"DE": Country{Name: "Germany", DialCode: "+49"},
		"GH": Country{Name: "Ghana", DialCode: "+233"},
		"GI": Country{Name: "Gibraltar", DialCode: "+350"},
		"GR": Country{Name: "Greece", DialCode: "+30"},
		"GL": Country{Name: "Greenland", DialCode: "+299"},
		"GD": Country{Name: "Grenada", DialCode: "+1 473"},
		"GP": Country{Name: "Guadeloupe", DialCode: "+590"},
		"GU": Country{Name: "Guam", DialCode: "+1 671"},
		"GT": Country{Name: "Guatemala", DialCode: "+502"},
		"GN": Country{Name: "Guinea", DialCode: "+224"},
		"GW": Country{Name: "Guinea-Bissau", DialCode: "+245"},
		"GY": Country{Name: "Guyana", DialCode: "+595"},
		"HT": Country{Name: "Haiti", DialCode: "+509"},
		"HN": Country{Name: "Honduras", DialCode: "+504"},
		"HU": Country{Name: "Hungary", DialCode: "+36"},
		"IS": Country{Name: "Iceland", DialCode: "+354"},
		"IN": Country{Name: "India", DialCode: "+91"},
		"ID": Country{Name: "Indonesia", DialCode: "+62"},
		"IQ": Country{Name: "Iraq", DialCode: "+964"},
		"IE": Country{Name: "Ireland", DialCode: "+353"},
		"IL": Country{Name: "Israel", DialCode: "+972"},
		"IT": Country{Name: "Italy", DialCode: "+39"},
		"JM": Country{Name: "Jamaica", DialCode: "+1 876"},
		"JP": Country{Name: "Japan", DialCode: "+81"},
		"JO": Country{Name: "Jordan", DialCode: "+962"},
		"KZ": Country{Name: "Kazakhstan", DialCode: "+7 7"},
		"KE": Country{Name: "Kenya", DialCode: "+254"},
		"KI": Country{Name: "Kiribati", DialCode: "+686"},
		"KW": Country{Name: "Kuwait", DialCode: "+965"},
		"KG": Country{Name: "Kyrgyzstan", DialCode: "+996"},
		"LV": Country{Name: "Latvia", DialCode: "+371"},
		"LB": Country{Name: "Lebanon", DialCode: "+961"},
		"LS": Country{Name: "Lesotho", DialCode: "+266"},
		"LR": Country{Name: "Liberia", DialCode: "+231"},
		"LI": Country{Name: "Liechtenstein", DialCode: "+423"},
		"LT": Country{Name: "Lithuania", DialCode: "+370"},
		"LU": Country{Name: "Luxembourg", DialCode: "+352"},
		"MG": Country{Name: "Madagascar", DialCode: "+261"},
		"MW": Country{Name: "Malawi", DialCode: "+265"},
		"MY": Country{Name: "Malaysia", DialCode: "+60"},
		"MV": Country{Name: "Maldives", DialCode: "+960"},
		"ML": Country{Name: "Mali", DialCode: "+223"},
		"MT": Country{Name: "Malta", DialCode: "+356"},
		"MH": Country{Name: "Marshall Islands", DialCode: "+692"},
		"MQ": Country{Name: "Martinique", DialCode: "+596"},
		"MR": Country{Name: "Mauritania", DialCode: "+222"},
		"MU": Country{Name: "Mauritius", DialCode: "+230"},
		"YT": Country{Name: "Mayotte", DialCode: "+262"},
		"MX": Country{Name: "Mexico", DialCode: "+52"},
		"MC": Country{Name: "Monaco", DialCode: "+377"},
		"MN": Country{Name: "Mongolia", DialCode: "+976"},
		"ME": Country{Name: "Montenegro", DialCode: "+382"},
		"MS": Country{Name: "Montserrat", DialCode: "+1664"},
		"MA": Country{Name: "Morocco", DialCode: "+212"},
		"MM": Country{Name: "Myanmar", DialCode: "+95"},
		"NA": Country{Name: "Namibia", DialCode: "+264"},
		"NR": Country{Name: "Nauru", DialCode: "+674"},
		"NP": Country{Name: "Nepal", DialCode: "+977"},
		"NL": Country{Name: "Netherlands", DialCode: "+31"},
		"AN": Country{Name: "Netherlands Antilles", DialCode: "+599"},
		"NC": Country{Name: "New Caledonia", DialCode: "+687"},
		"NZ": Country{Name: "New Zealand", DialCode: "+64"},
		"NI": Country{Name: "Nicaragua", DialCode: "+505"},
		"NE": Country{Name: "Niger", DialCode: "+227"},
		"NG": Country{Name: "Nigeria", DialCode: "+234"},
		"NU": Country{Name: "Niue", DialCode: "+683"},
		"NF": Country{Name: "Norfolk Island", DialCode: "+672"},
		"MP": Country{Name: "Northern Mariana Islands", DialCode: "+1 670"},
		"NO": Country{Name: "Norway", DialCode: "+47"},
		"OM": Country{Name: "Oman", DialCode: "+968"},
		"PK": Country{Name: "Pakistan", DialCode: "+92"},
		"PW": Country{Name: "Palau", DialCode: "+680"},
		"PA": Country{Name: "Panama", DialCode: "+507"},
		"PG": Country{Name: "Papua New Guinea", DialCode: "+675"},
		"PY": Country{Name: "Paraguay", DialCode: "+595"},
		"PE": Country{Name: "Peru", DialCode: "+51"},
		"PH": Country{Name: "Philippines", DialCode: "+63"},
		"PL": Country{Name: "Poland", DialCode: "+48"},
		"PT": Country{Name: "Portugal", DialCode: "+351"},
		"PR": Country{Name: "Puerto Rico", DialCode: "+1 939"},
		"QA": Country{Name: "Qatar", DialCode: "+974"},
		"RO": Country{Name: "Romania", DialCode: "+40"},
		"RW": Country{Name: "Rwanda", DialCode: "+250"},
		"WS": Country{Name: "Samoa", DialCode: "+685"},
		"SM": Country{Name: "San Marino", DialCode: "+378"},
		"SA": Country{Name: "Saudi Arabia", DialCode: "+966"},
		"SN": Country{Name: "Senegal", DialCode: "+221"},
		"RS": Country{Name: "Serbia", DialCode: "+381"},
		"SC": Country{Name: "Seychelles", DialCode: "+248"},
		"SL": Country{Name: "Sierra Leone", DialCode: "+232"},
		"SG": Country{Name: "Singapore", DialCode: "+65"},
		"SK": Country{Name: "Slovakia", DialCode: "+421"},
		"SI": Country{Name: "Slovenia", DialCode: "+386"},
		"SB": Country{Name: "Solomon Islands", DialCode: "+677"},
		"ZA": Country{Name: "South Africa", DialCode: "+27"},
		"GS": Country{Name: "South Georgia and the South Sandwich Islands", DialCode: "+500"},
		"ES": Country{Name: "Spain", DialCode: "+34"},
		"LK": Country{Name: "Sri Lanka", DialCode: "+94"},
		"SD": Country{Name: "Sudan", DialCode: "+249"},
		"SR": Country{Name: "Suriname", DialCode: "+597"},
		"SZ": Country{Name: "Swaziland", DialCode: "+268"},
		"SE": Country{Name: "Sweden", DialCode: "+46"},
		"CH": Country{Name: "Switzerland", DialCode: "+41"},
		"TJ": Country{Name: "Tajikistan", DialCode: "+992"},
		"TH": Country{Name: "Thailand", DialCode: "+66"},
		"TG": Country{Name: "Togo", DialCode: "+228"},
		"TK": Country{Name: "Tokelau", DialCode: "+690"},
		"TO": Country{Name: "Tonga", DialCode: "+676"},
		"TT": Country{Name: "Trinidad and Tobago", DialCode: "+1 868"},
		"TN": Country{Name: "Tunisia", DialCode: "+216"},
		"TR": Country{Name: "Turkey", DialCode: "+90"},
		"TM": Country{Name: "Turkmenistan", DialCode: "+993"},
		"TC": Country{Name: "Turks and Caicos Islands", DialCode: "+1 649"},
		"TV": Country{Name: "Tuvalu", DialCode: "+688"},
		"UG": Country{Name: "Uganda", DialCode: "+256"},
		"UA": Country{Name: "Ukraine", DialCode: "+380"},
		"AE": Country{Name: "United Arab Emirates", DialCode: "+971"},
		"GB": Country{Name: "United Kingdom", DialCode: "+44"},
		"US": Country{Name: "United States", DialCode: "+1"},
		"UY": Country{Name: "Uruguay", DialCode: "+598"},
		"UZ": Country{Name: "Uzbekistan", DialCode: "+998"},
		"VU": Country{Name: "Vanuatu", DialCode: "+678"},
		"WF": Country{Name: "Wallis and Futuna", DialCode: "+681"},
		"YE": Country{Name: "Yemen", DialCode: "+967"},
		"ZM": Country{Name: "Zambia", DialCode: "+260"},
		"ZW": Country{Name: "Zimbabwe", DialCode: "+263"},
		"AX": Country{Name: "land Islands", DialCode: ""},
		"AQ": Country{Name: "Antarctica", DialCode: "null"},
		"BO": Country{Name: "Bolivia, Plurinational State of", DialCode: "+591"},
		"BN": Country{Name: "Brunei Darussalam", DialCode: "+673"},
		"CC": Country{Name: "Cocos (Keeling) Islands", DialCode: "+61"},
		"CD": Country{Name: "Congo, The Democratic Republic of the", DialCode: "+243"},
		"CI": Country{Name: "Cote d'Ivoire", DialCode: "+225"},
		"FK": Country{Name: "Falkland Islands (Malvinas)", DialCode: "+500"},
		"GG": Country{Name: "Guernsey", DialCode: "+44"},
		"VA": Country{Name: "Holy See (Vatican City State)", DialCode: "+379"},
		"HK": Country{Name: "Hong Kong", DialCode: "+852"},
		"IR": Country{Name: "Iran, Islamic Republic of", DialCode: "+98"},
		"IM": Country{Name: "Isle of Man", DialCode: "+44"},
		"JE": Country{Name: "Jersey", DialCode: "+44"},
		"KP": Country{Name: "Korea, Democratic People's Republic of", DialCode: "+850"},
		"KR": Country{Name: "Korea, Republic of", DialCode: "+82"},
		"LA": Country{Name: "Lao People's Democratic Republic", DialCode: "+856"},
		"LY": Country{Name: "Libyan Arab Jamahiriya", DialCode: "+218"},
		"MO": Country{Name: "Macao", DialCode: "+853"},
		"MK": Country{Name: "Macedonia, The Former Yugoslav Republic of", DialCode: "+389"},
		"FM": Country{Name: "Micronesia, Federated States of", DialCode: "+691"},
		"MD": Country{Name: "Moldova, Republic of", DialCode: "+373"},
		"MZ": Country{Name: "Mozambique", DialCode: "+258"},
		"PS": Country{Name: "Palestinian Territory, Occupied", DialCode: "+970"},
		"PN": Country{Name: "Pitcairn", DialCode: "+872"},
		"RE": Country{Name: "Réunion", DialCode: "+262"},
		"RU": Country{Name: "Russia", DialCode: "+7"},
		"BL": Country{Name: "Saint Barthélemy", DialCode: "+590"},
		"SH": Country{Name: "Saint Helena, Ascension and Tristan Da Cunha", DialCode: "+290"},
		"KN": Country{Name: "Saint Kitts and Nevis", DialCode: "+1 869"},
		"LC": Country{Name: "Saint Lucia", DialCode: "+1 758"},
		"MF": Country{Name: "Saint Martin", DialCode: "+590"},
		"PM": Country{Name: "Saint Pierre and Miquelon", DialCode: "+508"},
		"VC": Country{Name: "Saint Vincent and the Grenadines", DialCode: "+1 784"},
		"ST": Country{Name: "Sao Tome and Principe", DialCode: "+239"},
		"SO": Country{Name: "Somalia", DialCode: "+252"},
		"SJ": Country{Name: "Svalbard and Jan Mayen", DialCode: "+47"},
		"SY": Country{Name: "Syrian Arab Republic", DialCode: "+963"},
		"TW": Country{Name: "Taiwan, Province of China", DialCode: "+886"},
		"TZ": Country{Name: "Tanzania, United Republic of", DialCode: "+255"},
		"TL": Country{Name: "Timor-Leste", DialCode: "+670"},
		"VE": Country{Name: "Venezuela, Bolivarian Republic of", DialCode: "+58"},
		"VN": Country{Name: "Viet Nam", DialCode: "+84"},
		"VG": Country{Name: "Virgin Islands, British", DialCode: "+1 284"},
		"VI": Country{Name: "Virgin Islands, U.S.", DialCode: "+1 340"},
	}
}

// Valid validates if the dial code exists in the list
func (d DialCode) Valid() bool {
	for _, element := range by_alpha2 {
		c := "+" + string(d)
		if element.DialCode == c {
			return true
		}
	}

	return false
}