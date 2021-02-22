// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pe

var localeMap = map[uint16]string{
	1025:  "Arabic - Saudi Arabia",
	1026:  "Bulgarian",
	1027:  "Catalan",
	1028:  "Chinese - Taiwan",
	1029:  "Czech",
	1030:  "Danish",
	1031:  "German - Germany",
	1032:  "Greek",
	1033:  "English - United States",
	1034:  "Spanish - Spain (Traditional Sort)",
	1035:  "Finnish",
	1036:  "French - France",
	1037:  "Hebrew",
	1038:  "Hungarian",
	1039:  "Icelandic",
	1040:  "Italian - Italy",
	1041:  "Japanese",
	1042:  "Korean",
	1043:  "Dutch - Netherlands",
	1044:  "Norwegian (Bokmål)",
	1045:  "Polish",
	1046:  "Portuguese - Brazil",
	1047:  "Rhaeto-Romanic",
	1048:  "Romanian",
	1049:  "Russian",
	1050:  "Croatian",
	1051:  "Slovak",
	1052:  "Albanian - Albania",
	1053:  "Swedish",
	1054:  "Thai",
	1055:  "Turkish",
	1056:  "Urdu - Pakistan",
	1057:  "Indonesian",
	1058:  "Ukrainian",
	1059:  "Belarusian",
	1060:  "Slovenian",
	1061:  "Estonian",
	1062:  "Latvian",
	1063:  "Lithuanian",
	1064:  "Tajik",
	1065:  "Persian",
	1066:  "Vietnamese",
	1067:  "Armenian - Armenia",
	1068:  "Azeri (Latin)",
	1069:  "Basque",
	1070:  "Sorbian",
	1071:  "F.Y.R.O. Macedonian",
	1072:  "Sutu",
	1073:  "Tsonga",
	1074:  "Tswana",
	1075:  "Venda",
	1076:  "Xhosa",
	1077:  "Zulu",
	1078:  "Afrikaans - South Africa",
	1079:  "Georgian",
	1080:  "Faroese",
	1081:  "Hindi",
	1082:  "Maltese",
	1083:  "Sami",
	1084:  "Gaelic (Scotland)",
	1085:  "Yiddish",
	1086:  "Malay - Malaysia",
	1087:  "Kazakh",
	1088:  "Kyrgyz (Cyrillic)",
	1089:  "Swahili",
	1090:  "Turkmen",
	1091:  "Uzbek (Latin)",
	1092:  "Tatar",
	1093:  "Bengali (India)",
	1094:  "Punjabi",
	1095:  "Gujarati",
	1096:  "Oriya",
	1097:  "Tamil",
	1098:  "Telugu",
	1099:  "Kannada",
	1100:  "Malayalam",
	1101:  "Assamese",
	1102:  "Marathi",
	1103:  "Sanskrit",
	1104:  "Mongolian (Cyrillic)",
	1105:  "Tibetan - People's Republic of China",
	1106:  "Welsh",
	1107:  "Khmer",
	1108:  "Lao",
	1109:  "Burmese",
	1110:  "Galician",
	1111:  "Konkani",
	1112:  "Manipuri",
	1113:  "Sindhi - India",
	1114:  "Syriac",
	1115:  "Sinhalese - Sri Lanka",
	1116:  "Cherokee - United States",
	1117:  "Inuktitut",
	1118:  "Amharic - Ethiopia",
	1119:  "Tamazight (Arabic)",
	1120:  "Kashmiri (Arabic)",
	1121:  "Nepali",
	1122:  "Frisian - Netherlands",
	1123:  "Pashto",
	1124:  "Filipino",
	1125:  "Divehi",
	1126:  "Edo",
	1127:  "Fulfulde - Nigeria",
	1128:  "Hausa - Nigeria",
	1129:  "Ibibio - Nigeria",
	1130:  "Yoruba",
	1131:  "Quecha - Bolivia",
	1132:  "Sepedi",
	1136:  "Igbo - Nigeria",
	1137:  "Kanuri - Nigeria",
	1138:  "Oromo",
	1139:  "Tigrigna - Ethiopia",
	1140:  "Guarani - Paraguay",
	1141:  "Hawaiian - United States",
	1142:  "Latin",
	1143:  "Somali",
	1144:  "Yi",
	1145:  "Papiamentu",
	1152:  "Uighur - China",
	1153:  "Maori - New Zealand",
	2049:  "Arabic - Iraq",
	2052:  "Chinese - People's Republic of China",
	2055:  "German - Switzerland",
	2057:  "English - United Kingdom",
	2058:  "Spanish - Mexico",
	2060:  "French - Belgium",
	2064:  "Italian - Switzerland",
	2067:  "Dutch - Belgium",
	2068:  "Norwegian (Nynorsk)",
	2070:  "Portuguese - Portugal",
	2072:  "Romanian - Moldava",
	2073:  "Russian - Moldava",
	2074:  "Serbian (Latin)",
	2077:  "Swedish - Finland",
	2080:  "Urdu - India",
	2092:  "Azeri (Cyrillic)",
	2108:  "Gaelic (Ireland)",
	2110:  "Malay - Brunei Darussalam",
	2115:  "Uzbek (Cyrillic)",
	2117:  "Bengali (Bangladesh)",
	2118:  "Punjabi (Pakistan)",
	2128:  "Mongolian (Mongolian)",
	2129:  "Tibetan - Bhutan",
	2137:  "Sindhi - Pakistan",
	2143:  "Tamazight (Latin)",
	2144:  "Kashmiri (Devanagari)",
	2145:  "Nepali - India",
	2155:  "Quecha - Ecuador",
	2163:  "Tigrigna - Eritrea",
	3073:  "Arabic - Egypt",
	3076:  "Chinese - Hong Kong SAR",
	3079:  "German - Austria",
	3081:  "English - Australia",
	3082:  "Spanish - Spain (Modern Sort)",
	3084:  "French - Canada",
	3098:  "Serbian (Cyrillic)",
	3179:  "Quecha - Peru",
	4097:  "Arabic - Libya",
	4100:  "Chinese - Singapore",
	4103:  "German - Luxembourg",
	4105:  "English - Canada",
	4106:  "Spanish - Guatemala",
	4108:  "French - Switzerland",
	4122:  "Croatian (Bosnia/Herzegovina)",
	5121:  "Arabic - Algeria",
	5124:  "Chinese - Macao SAR",
	5127:  "German - Liechtenstein",
	5129:  "English - New Zealand",
	5130:  "Spanish - Costa Rica",
	5132:  "French - Luxembourg",
	5146:  "Bosnian (Bosnia/Herzegovina)",
	6145:  "Arabic - Morocco",
	6153:  "English - Ireland",
	6154:  "Spanish - Panama",
	6156:  "French - Monaco",
	7169:  "Arabic - Tunisia",
	7177:  "English - South Africa",
	7178:  "Spanish - Dominican Republic",
	7180:  "French - West Indies",
	8193:  "Arabic - Oman",
	8201:  "English - Jamaica",
	8202:  "Spanish - Venezuela",
	8204:  "French - Reunion",
	9217:  "Arabic - Yemen",
	9225:  "English - Caribbean",
	9226:  "Spanish - Colombia",
	9228:  "French - Democratic Rep. of Congo",
	10241: "Arabic - Syria",
	10249: "English - Belize",
	10250: "Spanish - Peru",
	10252: "French - Senegal",
	11265: "Arabic - Jordan",
	11273: "English - Trinidad",
	11274: "Spanish - Argentina",
	11276: "French - Cameroon",
	12289: "Arabic - Lebanon",
	12297: "English - Zimbabwe",
	12298: "Spanish - Ecuador",
	12300: "French - Cote d'Ivoire",
	13313: "Arabic - Kuwait",
	13321: "English - Philippines",
	13322: "Spanish - Chile",
	13324: "French - Mali",
	14337: "Arabic - U.A.E.",
	14345: "English - Indonesia",
	14346: "Spanish - Uruguay",
	14348: "French - Morocco",
	15361: "Arabic - Bahrain",
	15369: "English - Hong Kong SAR",
	15370: "Spanish - Paraguay",
	15372: "French - Haiti",
	16385: "Arabic - Qatar",
	16393: "English - India",
	16394: "Spanish - Bolivia",
	17417: "English - Malaysia",
	17418: "Spanish - El Salvador",
	18441: "English - Singapore",
	18442: "Spanish - Honduras",
	19466: "Spanish - Nicaragua",
	20490: "Spanish - Puerto Rico",
	21514: "Spanish - United States",
	58378: "Spanish - Latin America",
	58380: "French - North Africa",
}

func languageName(language uint16) string {
	if found, ok := localeMap[language]; ok {
		return found
	}
	return "Unknown"
}
