package main

import (
	"strings"
)

var (
	// This need to be a list and not a map so it is ordered. This ordering
	// allows us to not match "blue" in "aliceblue" since "aliceblue" is a
	// colour that will already be matched.
	// https://www.w3schools.com/colors/colors_names.asp
	webColours = [][2]string{
		[2]string{"lightgoldenrodyellow", "#fafad2"},
		[2]string{"mediumspringgreen", "#00fa9a"},
		[2]string{"mediumaquamarine", "#66cdaa"},
		[2]string{"mediumslateblue", "#7b68ee"},
		[2]string{"mediumvioletred", "#c71585"},
		[2]string{"mediumturquoise", "#48d1cc"},
		[2]string{"lightslategray", "#778899"},
		[2]string{"lightslategrey", "#778899"},
		[2]string{"mediumseagreen", "#3cb371"},
		[2]string{"blanchedalmond", "#ffebcd"},
		[2]string{"lightsteelblue", "#b0c4de"},
		[2]string{"cornflowerblue", "#6495ed"},
		[2]string{"darkolivegreen", "#556b2f"},
		[2]string{"darkgoldenrod", "#b8860b"},
		[2]string{"palegoldenrod", "#eee8aa"},
		[2]string{"paleturquoise", "#afeeee"},
		[2]string{"lavenderblush", "#fff0f5"},
		[2]string{"rebeccapurple", "#663399"},
		[2]string{"darkslategray", "#2f4f4f"},
		[2]string{"darkslategrey", "#2f4f4f"},
		[2]string{"lightseagreen", "#20b2aa"},
		[2]string{"palevioletred", "#db7093"},
		[2]string{"darkturquoise", "#00ced1"},
		[2]string{"darkslateblue", "#483d8b"},
		[2]string{"antiquewhite", "#faebd7"},
		[2]string{"darkseagreen", "#8fbc8f"},
		[2]string{"lightskyblue", "#87cefa"},
		[2]string{"mediumorchid", "#ba55d3"},
		[2]string{"lemonchiffon", "#fffacd"},
		[2]string{"mediumpurple", "#9370db"},
		[2]string{"midnightblue", "#191970"},
		[2]string{"greenyellow", "#adff2f"},
		[2]string{"darkmagenta", "#8b008b"},
		[2]string{"lightsalmon", "#ffa07a"},
		[2]string{"lightyellow", "#ffffe0"},
		[2]string{"deepskyblue", "#00bfff"},
		[2]string{"navajowhite", "#ffdead"},
		[2]string{"saddlebrown", "#8b4513"},
		[2]string{"springgreen", "#00ff7f"},
		[2]string{"forestgreen", "#228b22"},
		[2]string{"floralwhite", "#fffaf0"},
		[2]string{"yellowgreen", "#9acd32"},
		[2]string{"papayawhip", "#ffefd5"},
		[2]string{"aquamarine", "#7fffd4"},
		[2]string{"dodgerblue", "#1e90ff"},
		[2]string{"chartreuse", "#7fff00"},
		[2]string{"blueviolet", "#8a2be2"},
		[2]string{"darkviolet", "#9400d3"},
		[2]string{"darkorange", "#ff8c00"},
		[2]string{"lightgreen", "#90ee90"},
		[2]string{"ghostwhite", "#f8f8ff"},
		[2]string{"whitesmoke", "#f5f5f5"},
		[2]string{"darkorchid", "#9932cc"},
		[2]string{"mediumblue", "#0000cd"},
		[2]string{"powderblue", "#b0e0e6"},
		[2]string{"lightcoral", "#f08080"},
		[2]string{"darksalmon", "#e9967a"},
		[2]string{"sandybrown", "#f4a460"},
		[2]string{"indianred", "#cd5c5c"},
		[2]string{"royalblue", "#4169e1"},
		[2]string{"steelblue", "#4682b4"},
		[2]string{"aliceblue", "#f0f8ff"},
		[2]string{"slategrey", "#708090"},
		[2]string{"mistyrose", "#ffe4e1"},
		[2]string{"turquoise", "#40e0d0"},
		[2]string{"lawngreen", "#7cfc00"},
		[2]string{"mintcream", "#f5fffa"},
		[2]string{"lightblue", "#add8e6"},
		[2]string{"slategray", "#708090"},
		[2]string{"lightcyan", "#e0ffff"},
		[2]string{"goldenrod", "#daa520"},
		[2]string{"lightgray", "#d3d3d3"},
		[2]string{"lightgrey", "#d3d3d3"},
		[2]string{"gainsboro", "#dcdcdc"},
		[2]string{"olivedrab", "#6b8e23"},
		[2]string{"chocolate", "#d2691e"},
		[2]string{"darkgreen", "#006400"},
		[2]string{"peachpuff", "#ffdab9"},
		[2]string{"rosybrown", "#bc8f8f"},
		[2]string{"burlywood", "#deb887"},
		[2]string{"firebrick", "#b22222"},
		[2]string{"slateblue", "#6a5acd"},
		[2]string{"lightpink", "#ffb6c1"},
		[2]string{"limegreen", "#32cd32"},
		[2]string{"orangered", "#ff4500"},
		[2]string{"cadetblue", "#5f9ea0"},
		[2]string{"darkkhaki", "#bdb76b"},
		[2]string{"palegreen", "#98fb98"},
		[2]string{"honeydew", "#f0fff0"},
		[2]string{"seashell", "#fff5ee"},
		[2]string{"seagreen", "#2e8b57"},
		[2]string{"deeppink", "#ff1493"},
		[2]string{"cornsilk", "#fff8dc"},
		[2]string{"darkblue", "#00008b"},
		[2]string{"darkcyan", "#008b8b"},
		[2]string{"darkgray", "#a9a9a9"},
		[2]string{"darkgrey", "#a9a9a9"},
		[2]string{"moccasin", "#ffe4b5"},
		[2]string{"lavender", "#e6e6fa"},
		[2]string{"darkred", "#8b0000"},
		[2]string{"hotpink", "#ff69b4"},
		[2]string{"skyblue", "#87ceeb"},
		[2]string{"oldlace", "#fdf5e6"},
		[2]string{"thistle", "#d8bfd8"},
		[2]string{"fuchsia", "#ff00ff"},
		[2]string{"magenta", "#ff00ff"},
		[2]string{"dimgrey", "#696969"},
		[2]string{"crimson", "#dc143c"},
		[2]string{"dimgray", "#696969"},
		[2]string{"tomato", "#ff6347"},
		[2]string{"bisque", "#ffe4c4"},
		[2]string{"silver", "#c0c0c0"},
		[2]string{"orchid", "#da70d6"},
		[2]string{"orange", "#ffa500"},
		[2]string{"yellow", "#ffff00"},
		[2]string{"sienna", "#a0522d"},
		[2]string{"maroon", "#800000"},
		[2]string{"salmon", "#fa8072"},
		[2]string{"purple", "#800080"},
		[2]string{"indigo", "#4b0082"},
		[2]string{"violet", "#ee82ee"},
		[2]string{"green", "#008000"},
		[2]string{"beige", "#f5f5dc"},
		[2]string{"azure", "#f0ffff"},
		[2]string{"olive", "#808000"},
		[2]string{"ivory", "#fffff0"},
		[2]string{"coral", "#ff7f50"},
		[2]string{"wheat", "#f5deb3"},
		[2]string{"white", "#ffffff"},
		[2]string{"linen", "#faf0e6"},
		[2]string{"brown", "#a52a2a"},
		[2]string{"khaki", "#f0e68c"},
		[2]string{"black", "#000000"},
		[2]string{"cyan", "#00ffff"},
		[2]string{"blue", "#0000ff"},
		[2]string{"aqua", "#00ffff"},
		[2]string{"navy", "#000080"},
		[2]string{"peru", "#cd853f"},
		[2]string{"teal", "#008080"},
		[2]string{"grey", "#808080"},
		[2]string{"snow", "#fffafa"},
		[2]string{"gray", "#808080"},
		[2]string{"gold", "#ffd700"},
		[2]string{"plum", "#dda0dd"},
		[2]string{"pink", "#ffc0cb"},
		[2]string{"lime", "#00ff00"},
		[2]string{"red", "#ff0000"},
		[2]string{"tan", "#d2b48c"},
        [2]string{"white", "#FFFFFF"},
        [2]string{"aliceblue", "#F0F8FF"},
        [2]string{"antiquewhite", "#FAEBD7"},
        [2]string{"antiquewhite1", "#FFEFDB"},
        [2]string{"antiquewhite2", "#EEDFCC"},
        [2]string{"antiquewhite3", "#CDC0B0"},
        [2]string{"antiquewhite4", "#8B8378"},
        [2]string{"aquamarine", "#7FFFD4"},
        [2]string{"aquamarine1", "#7FFFD4"},
        [2]string{"aquamarine2", "#76EEC6"},
        [2]string{"aquamarine3", "#66CDAA"},
        [2]string{"aquamarine4", "#458B74"},
        [2]string{"azure", "#F0FFFF"},
        [2]string{"azure1", "#F0FFFF"},
        [2]string{"azure2", "#E0EEEE"},
        [2]string{"azure3", "#C1CDCD"},
        [2]string{"azure4", "#838B8B"},
        [2]string{"beige", "#F5F5DC"},
        [2]string{"bisque", "#FFE4C4"},
        [2]string{"bisque1", "#FFE4C4"},
        [2]string{"bisque2", "#EED5B7"},
        [2]string{"bisque3", "#CDB79E"},
        [2]string{"bisque4", "#8B7D6B"},
        [2]string{"black", "#000000"},
        [2]string{"blanchedalmond", "#FFEBCD"},
        [2]string{"blue", "#0000FF"},
        [2]string{"blue1", "#0000FF"},
        [2]string{"blue2", "#0000EE"},
        [2]string{"blue3", "#0000CD"},
        [2]string{"blue4", "#00008B"},
        [2]string{"blueviolet", "#8A2BE2"},
        [2]string{"brown", "#A52A2A"},
        [2]string{"brown1", "#FF4040"},
        [2]string{"brown2", "#EE3B3B"},
        [2]string{"brown3", "#CD3333"},
        [2]string{"brown4", "#8B2323"},
        [2]string{"burlywood", "#DEB887"},
        [2]string{"burlywood1", "#FFD39B"},
        [2]string{"burlywood2", "#EEC591"},
        [2]string{"burlywood3", "#CDAA7D"},
        [2]string{"burlywood4", "#8B7355"},
        [2]string{"cadetblue", "#5F9EA0"},
        [2]string{"cadetblue1", "#98F5FF"},
        [2]string{"cadetblue2", "#8EE5EE"},
        [2]string{"cadetblue3", "#7AC5CD"},
        [2]string{"cadetblue4", "#53868B"},
        [2]string{"chartreuse", "#7FFF00"},
        [2]string{"chartreuse1", "#7FFF00"},
        [2]string{"chartreuse2", "#76EE00"},
        [2]string{"chartreuse3", "#66CD00"},
        [2]string{"chartreuse4", "#458B00"},
        [2]string{"chocolate", "#D2691E"},
        [2]string{"chocolate1", "#FF7F24"},
        [2]string{"chocolate2", "#EE7621"},
        [2]string{"chocolate3", "#CD661D"},
        [2]string{"chocolate4", "#8B4513"},
        [2]string{"coral", "#FF7F50"},
        [2]string{"coral1", "#FF7256"},
        [2]string{"coral2", "#EE6A50"},
        [2]string{"coral3", "#CD5B45"},
        [2]string{"coral4", "#8B3E2F"},
        [2]string{"cornflowerblue", "#6495ED"},
        [2]string{"cornsilk", "#FFF8DC"},
        [2]string{"cornsilk1", "#FFF8DC"},
        [2]string{"cornsilk2", "#EEE8CD"},
        [2]string{"cornsilk3", "#CDC8B1"},
        [2]string{"cornsilk4", "#8B8878"},
        [2]string{"cyan", "#00FFFF"},
        [2]string{"cyan1", "#00FFFF"},
        [2]string{"cyan2", "#00EEEE"},
        [2]string{"cyan3", "#00CDCD"},
        [2]string{"cyan4", "#008B8B"},
        [2]string{"darkblue", "#00008B"},
        [2]string{"darkcyan", "#008B8B"},
        [2]string{"darkgoldenrod", "#B8860B"},
        [2]string{"darkgoldenrod1", "#FFB90F"},
        [2]string{"darkgoldenrod2", "#EEAD0E"},
        [2]string{"darkgoldenrod3", "#CD950C"},
        [2]string{"darkgoldenrod4", "#8B6508"},
        [2]string{"darkgray", "#A9A9A9"},
        [2]string{"darkgreen", "#006400"},
        [2]string{"darkgrey", "#A9A9A9"},
        [2]string{"darkkhaki", "#BDB76B"},
        [2]string{"darkmagenta", "#8B008B"},
        [2]string{"darkolivegreen", "#556B2F"},
        [2]string{"darkolivegreen1", "#CAFF70"},
        [2]string{"darkolivegreen2", "#BCEE68"},
        [2]string{"darkolivegreen3", "#A2CD5A"},
        [2]string{"darkolivegreen4", "#6E8B3D"},
        [2]string{"darkorange", "#FF8C00"},
        [2]string{"darkorange1", "#FF7F00"},
        [2]string{"darkorange2", "#EE7600"},
        [2]string{"darkorange3", "#CD6600"},
        [2]string{"darkorange4", "#8B4500"},
        [2]string{"darkorchid", "#9932CC"},
        [2]string{"darkorchid1", "#BF3EFF"},
        [2]string{"darkorchid2", "#B23AEE"},
        [2]string{"darkorchid3", "#9A32CD"},
        [2]string{"darkorchid4", "#68228B"},
        [2]string{"darkred", "#8B0000"},
        [2]string{"darksalmon", "#E9967A"},
        [2]string{"darkseagreen", "#8FBC8F"},
        [2]string{"darkseagreen1", "#C1FFC1"},
        [2]string{"darkseagreen2", "#B4EEB4"},
        [2]string{"darkseagreen3", "#9BCD9B"},
        [2]string{"darkseagreen4", "#698B69"},
        [2]string{"darkslateblue", "#483D8B"},
        [2]string{"darkslategray", "#2F4F4F"},
        [2]string{"darkslategray1", "#97FFFF"},
        [2]string{"darkslategray2", "#8DEEEE"},
        [2]string{"darkslategray3", "#79CDCD"},
        [2]string{"darkslategray4", "#528B8B"},
        [2]string{"darkslategrey", "#2F4F4F"},
        [2]string{"darkturquoise", "#00CED1"},
        [2]string{"darkviolet", "#9400D3"},
        [2]string{"deeppink", "#FF1493"},
        [2]string{"deeppink1", "#FF1493"},
        [2]string{"deeppink2", "#EE1289"},
        [2]string{"deeppink3", "#CD1076"},
        [2]string{"deeppink4", "#8B0A50"},
        [2]string{"deepskyblue", "#00BFFF"},
        [2]string{"deepskyblue1", "#00BFFF"},
        [2]string{"deepskyblue2", "#00B2EE"},
        [2]string{"deepskyblue3", "#009ACD"},
        [2]string{"deepskyblue4", "#00688B"},
        [2]string{"dimgray", "#696969"},
        [2]string{"dimgrey", "#696969"},
        [2]string{"dodgerblue", "#1E90FF"},
        [2]string{"dodgerblue1", "#1E90FF"},
        [2]string{"dodgerblue2", "#1C86EE"},
        [2]string{"dodgerblue3", "#1874CD"},
        [2]string{"dodgerblue4", "#104E8B"},
        [2]string{"firebrick", "#B22222"},
        [2]string{"firebrick1", "#FF3030"},
        [2]string{"firebrick2", "#EE2C2C"},
        [2]string{"firebrick3", "#CD2626"},
        [2]string{"firebrick4", "#8B1A1A"},
        [2]string{"floralwhite", "#FFFAF0"},
        [2]string{"forestgreen", "#228B22"},
        [2]string{"gainsboro", "#DCDCDC"},
        [2]string{"ghostwhite", "#F8F8FF"},
        [2]string{"gold", "#FFD700"},
        [2]string{"gold1", "#FFD700"},
        [2]string{"gold2", "#EEC900"},
        [2]string{"gold3", "#CDAD00"},
        [2]string{"gold4", "#8B7500"},
        [2]string{"goldenrod", "#DAA520"},
        [2]string{"goldenrod1", "#FFC125"},
        [2]string{"goldenrod2", "#EEB422"},
        [2]string{"goldenrod3", "#CD9B1D"},
        [2]string{"goldenrod4", "#8B6914"},
        [2]string{"gray", "#BEBEBE"},
        [2]string{"gray0", "#000000"},
        [2]string{"gray1", "#030303"},
        [2]string{"gray2", "#050505"},
        [2]string{"gray3", "#080808"},
        [2]string{"gray4", "#0A0A0A"},
        [2]string{"gray5", "#0D0D0D"},
        [2]string{"gray6", "#0F0F0F"},
        [2]string{"gray7", "#121212"},
        [2]string{"gray8", "#141414"},
        [2]string{"gray9", "#171717"},
        [2]string{"gray10", "#1A1A1A"},
        [2]string{"gray11", "#1C1C1C"},
        [2]string{"gray12", "#1F1F1F"},
        [2]string{"gray13", "#212121"},
        [2]string{"gray14", "#242424"},
        [2]string{"gray15", "#262626"},
        [2]string{"gray16", "#292929"},
        [2]string{"gray17", "#2B2B2B"},
        [2]string{"gray18", "#2E2E2E"},
        [2]string{"gray19", "#303030"},
        [2]string{"gray20", "#333333"},
        [2]string{"gray21", "#363636"},
        [2]string{"gray22", "#383838"},
        [2]string{"gray23", "#3B3B3B"},
        [2]string{"gray24", "#3D3D3D"},
        [2]string{"gray25", "#404040"},
        [2]string{"gray26", "#424242"},
        [2]string{"gray27", "#454545"},
        [2]string{"gray28", "#474747"},
        [2]string{"gray29", "#4A4A4A"},
        [2]string{"gray30", "#4D4D4D"},
        [2]string{"gray31", "#4F4F4F"},
        [2]string{"gray32", "#525252"},
        [2]string{"gray33", "#545454"},
        [2]string{"gray34", "#575757"},
        [2]string{"gray35", "#595959"},
        [2]string{"gray36", "#5C5C5C"},
        [2]string{"gray37", "#5E5E5E"},
        [2]string{"gray38", "#616161"},
        [2]string{"gray39", "#636363"},
        [2]string{"gray40", "#666666"},
        [2]string{"gray41", "#696969"},
        [2]string{"gray42", "#6B6B6B"},
        [2]string{"gray43", "#6E6E6E"},
        [2]string{"gray44", "#707070"},
        [2]string{"gray45", "#737373"},
        [2]string{"gray46", "#757575"},
        [2]string{"gray47", "#787878"},
        [2]string{"gray48", "#7A7A7A"},
        [2]string{"gray49", "#7D7D7D"},
        [2]string{"gray50", "#7F7F7F"},
        [2]string{"gray51", "#828282"},
        [2]string{"gray52", "#858585"},
        [2]string{"gray53", "#878787"},
        [2]string{"gray54", "#8A8A8A"},
        [2]string{"gray55", "#8C8C8C"},
        [2]string{"gray56", "#8F8F8F"},
        [2]string{"gray57", "#919191"},
        [2]string{"gray58", "#949494"},
        [2]string{"gray59", "#969696"},
        [2]string{"gray60", "#999999"},
        [2]string{"gray61", "#9C9C9C"},
        [2]string{"gray62", "#9E9E9E"},
        [2]string{"gray63", "#A1A1A1"},
        [2]string{"gray64", "#A3A3A3"},
        [2]string{"gray65", "#A6A6A6"},
        [2]string{"gray66", "#A8A8A8"},
        [2]string{"gray67", "#ABABAB"},
        [2]string{"gray68", "#ADADAD"},
        [2]string{"gray69", "#B0B0B0"},
        [2]string{"gray70", "#B3B3B3"},
        [2]string{"gray71", "#B5B5B5"},
        [2]string{"gray72", "#B8B8B8"},
        [2]string{"gray73", "#BABABA"},
        [2]string{"gray74", "#BDBDBD"},
        [2]string{"gray75", "#BFBFBF"},
        [2]string{"gray76", "#C2C2C2"},
        [2]string{"gray77", "#C4C4C4"},
        [2]string{"gray78", "#C7C7C7"},
        [2]string{"gray79", "#C9C9C9"},
        [2]string{"gray80", "#CCCCCC"},
        [2]string{"gray81", "#CFCFCF"},
        [2]string{"gray82", "#D1D1D1"},
        [2]string{"gray83", "#D4D4D4"},
        [2]string{"gray84", "#D6D6D6"},
        [2]string{"gray85", "#D9D9D9"},
        [2]string{"gray86", "#DBDBDB"},
        [2]string{"gray87", "#DEDEDE"},
        [2]string{"gray88", "#E0E0E0"},
        [2]string{"gray89", "#E3E3E3"},
        [2]string{"gray90", "#E5E5E5"},
        [2]string{"gray91", "#E8E8E8"},
        [2]string{"gray92", "#EBEBEB"},
        [2]string{"gray93", "#EDEDED"},
        [2]string{"gray94", "#F0F0F0"},
        [2]string{"gray95", "#F2F2F2"},
        [2]string{"gray96", "#F5F5F5"},
        [2]string{"gray97", "#F7F7F7"},
        [2]string{"gray98", "#FAFAFA"},
        [2]string{"gray99", "#FCFCFC"},
        [2]string{"gray100", "#FFFFFF"},
        [2]string{"green", "#00FF00"},
        [2]string{"green1", "#00FF00"},
        [2]string{"green2", "#00EE00"},
        [2]string{"green3", "#00CD00"},
        [2]string{"green4", "#008B00"},
        [2]string{"greenyellow", "#ADFF2F"},
        [2]string{"grey", "#BEBEBE"},
        [2]string{"grey0", "#000000"},
        [2]string{"grey1", "#030303"},
        [2]string{"grey2", "#050505"},
        [2]string{"grey3", "#080808"},
        [2]string{"grey4", "#0A0A0A"},
        [2]string{"grey5", "#0D0D0D"},
        [2]string{"grey6", "#0F0F0F"},
        [2]string{"grey7", "#121212"},
        [2]string{"grey8", "#141414"},
        [2]string{"grey9", "#171717"},
        [2]string{"grey10", "#1A1A1A"},
        [2]string{"grey11", "#1C1C1C"},
        [2]string{"grey12", "#1F1F1F"},
        [2]string{"grey13", "#212121"},
        [2]string{"grey14", "#242424"},
        [2]string{"grey15", "#262626"},
        [2]string{"grey16", "#292929"},
        [2]string{"grey17", "#2B2B2B"},
        [2]string{"grey18", "#2E2E2E"},
        [2]string{"grey19", "#303030"},
        [2]string{"grey20", "#333333"},
        [2]string{"grey21", "#363636"},
        [2]string{"grey22", "#383838"},
        [2]string{"grey23", "#3B3B3B"},
        [2]string{"grey24", "#3D3D3D"},
        [2]string{"grey25", "#404040"},
        [2]string{"grey26", "#424242"},
        [2]string{"grey27", "#454545"},
        [2]string{"grey28", "#474747"},
        [2]string{"grey29", "#4A4A4A"},
        [2]string{"grey30", "#4D4D4D"},
        [2]string{"grey31", "#4F4F4F"},
        [2]string{"grey32", "#525252"},
        [2]string{"grey33", "#545454"},
        [2]string{"grey34", "#575757"},
        [2]string{"grey35", "#595959"},
        [2]string{"grey36", "#5C5C5C"},
        [2]string{"grey37", "#5E5E5E"},
        [2]string{"grey38", "#616161"},
        [2]string{"grey39", "#636363"},
        [2]string{"grey40", "#666666"},
        [2]string{"grey41", "#696969"},
        [2]string{"grey42", "#6B6B6B"},
        [2]string{"grey43", "#6E6E6E"},
        [2]string{"grey44", "#707070"},
        [2]string{"grey45", "#737373"},
        [2]string{"grey46", "#757575"},
        [2]string{"grey47", "#787878"},
        [2]string{"grey48", "#7A7A7A"},
        [2]string{"grey49", "#7D7D7D"},
        [2]string{"grey50", "#7F7F7F"},
        [2]string{"grey51", "#828282"},
        [2]string{"grey52", "#858585"},
        [2]string{"grey53", "#878787"},
        [2]string{"grey54", "#8A8A8A"},
        [2]string{"grey55", "#8C8C8C"},
        [2]string{"grey56", "#8F8F8F"},
        [2]string{"grey57", "#919191"},
        [2]string{"grey58", "#949494"},
        [2]string{"grey59", "#969696"},
        [2]string{"grey60", "#999999"},
        [2]string{"grey61", "#9C9C9C"},
        [2]string{"grey62", "#9E9E9E"},
        [2]string{"grey63", "#A1A1A1"},
        [2]string{"grey64", "#A3A3A3"},
        [2]string{"grey65", "#A6A6A6"},
        [2]string{"grey66", "#A8A8A8"},
        [2]string{"grey67", "#ABABAB"},
        [2]string{"grey68", "#ADADAD"},
        [2]string{"grey69", "#B0B0B0"},
        [2]string{"grey70", "#B3B3B3"},
        [2]string{"grey71", "#B5B5B5"},
        [2]string{"grey72", "#B8B8B8"},
        [2]string{"grey73", "#BABABA"},
        [2]string{"grey74", "#BDBDBD"},
        [2]string{"grey75", "#BFBFBF"},
        [2]string{"grey76", "#C2C2C2"},
        [2]string{"grey77", "#C4C4C4"},
        [2]string{"grey78", "#C7C7C7"},
        [2]string{"grey79", "#C9C9C9"},
        [2]string{"grey80", "#CCCCCC"},
        [2]string{"grey81", "#CFCFCF"},
        [2]string{"grey82", "#D1D1D1"},
        [2]string{"grey83", "#D4D4D4"},
        [2]string{"grey84", "#D6D6D6"},
        [2]string{"grey85", "#D9D9D9"},
        [2]string{"grey86", "#DBDBDB"},
        [2]string{"grey87", "#DEDEDE"},
        [2]string{"grey88", "#E0E0E0"},
        [2]string{"grey89", "#E3E3E3"},
        [2]string{"grey90", "#E5E5E5"},
        [2]string{"grey91", "#E8E8E8"},
        [2]string{"grey92", "#EBEBEB"},
        [2]string{"grey93", "#EDEDED"},
        [2]string{"grey94", "#F0F0F0"},
        [2]string{"grey95", "#F2F2F2"},
        [2]string{"grey96", "#F5F5F5"},
        [2]string{"grey97", "#F7F7F7"},
        [2]string{"grey98", "#FAFAFA"},
        [2]string{"grey99", "#FCFCFC"},
        [2]string{"grey100", "#FFFFFF"},
        [2]string{"honeydew", "#F0FFF0"},
        [2]string{"honeydew1", "#F0FFF0"},
        [2]string{"honeydew2", "#E0EEE0"},
        [2]string{"honeydew3", "#C1CDC1"},
        [2]string{"honeydew4", "#838B83"},
        [2]string{"hotpink", "#FF69B4"},
        [2]string{"hotpink1", "#FF6EB4"},
        [2]string{"hotpink2", "#EE6AA7"},
        [2]string{"hotpink3", "#CD6090"},
        [2]string{"hotpink4", "#8B3A62"},
        [2]string{"indianred", "#CD5C5C"},
        [2]string{"indianred1", "#FF6A6A"},
        [2]string{"indianred2", "#EE6363"},
        [2]string{"indianred3", "#CD5555"},
        [2]string{"indianred4", "#8B3A3A"},
        [2]string{"ivory", "#FFFFF0"},
        [2]string{"ivory1", "#FFFFF0"},
        [2]string{"ivory2", "#EEEEE0"},
        [2]string{"ivory3", "#CDCDC1"},
        [2]string{"ivory4", "#8B8B83"},
        [2]string{"khaki", "#F0E68C"},
        [2]string{"khaki1", "#FFF68F"},
        [2]string{"khaki2", "#EEE685"},
        [2]string{"khaki3", "#CDC673"},
        [2]string{"khaki4", "#8B864E"},
        [2]string{"lavender", "#E6E6FA"},
        [2]string{"lavenderblush", "#FFF0F5"},
        [2]string{"lavenderblush1", "#FFF0F5"},
        [2]string{"lavenderblush2", "#EEE0E5"},
        [2]string{"lavenderblush3", "#CDC1C5"},
        [2]string{"lavenderblush4", "#8B8386"},
        [2]string{"lawngreen", "#7CFC00"},
        [2]string{"lemonchiffon", "#FFFACD"},
        [2]string{"lemonchiffon1", "#FFFACD"},
        [2]string{"lemonchiffon2", "#EEE9BF"},
        [2]string{"lemonchiffon3", "#CDC9A5"},
        [2]string{"lemonchiffon4", "#8B8970"},
        [2]string{"lightblue", "#ADD8E6"},
        [2]string{"lightblue1", "#BFEFFF"},
        [2]string{"lightblue2", "#B2DFEE"},
        [2]string{"lightblue3", "#9AC0CD"},
        [2]string{"lightblue4", "#68838B"},
        [2]string{"lightcoral", "#F08080"},
        [2]string{"lightcyan", "#E0FFFF"},
        [2]string{"lightcyan1", "#E0FFFF"},
        [2]string{"lightcyan2", "#D1EEEE"},
        [2]string{"lightcyan3", "#B4CDCD"},
        [2]string{"lightcyan4", "#7A8B8B"},
        [2]string{"lightgoldenrod", "#EEDD82"},
        [2]string{"lightgoldenrod1", "#FFEC8B"},
        [2]string{"lightgoldenrod2", "#EEDC82"},
        [2]string{"lightgoldenrod3", "#CDBE70"},
        [2]string{"lightgoldenrod4", "#8B814C"},
        [2]string{"lightgoldenrodyellow", "#FAFAD2"},
        [2]string{"lightgray", "#D3D3D3"},
        [2]string{"lightgreen", "#90EE90"},
        [2]string{"lightgrey", "#D3D3D3"},
        [2]string{"lightpink", "#FFB6C1"},
        [2]string{"lightpink1", "#FFAEB9"},
        [2]string{"lightpink2", "#EEA2AD"},
        [2]string{"lightpink3", "#CD8C95"},
        [2]string{"lightpink4", "#8B5F65"},
        [2]string{"lightsalmon", "#FFA07A"},
        [2]string{"lightsalmon1", "#FFA07A"},
        [2]string{"lightsalmon2", "#EE9572"},
        [2]string{"lightsalmon3", "#CD8162"},
        [2]string{"lightsalmon4", "#8B5742"},
        [2]string{"lightseagreen", "#20B2AA"},
        [2]string{"lightskyblue", "#87CEFA"},
        [2]string{"lightskyblue1", "#B0E2FF"},
        [2]string{"lightskyblue2", "#A4D3EE"},
        [2]string{"lightskyblue3", "#8DB6CD"},
        [2]string{"lightskyblue4", "#607B8B"},
        [2]string{"lightslateblue", "#8470FF"},
        [2]string{"lightslategray", "#778899"},
        [2]string{"lightslategrey", "#778899"},
        [2]string{"lightsteelblue", "#B0C4DE"},
        [2]string{"lightsteelblue1", "#CAE1FF"},
        [2]string{"lightsteelblue2", "#BCD2EE"},
        [2]string{"lightsteelblue3", "#A2B5CD"},
        [2]string{"lightsteelblue4", "#6E7B8B"},
        [2]string{"lightyellow", "#FFFFE0"},
        [2]string{"lightyellow1", "#FFFFE0"},
        [2]string{"lightyellow2", "#EEEED1"},
        [2]string{"lightyellow3", "#CDCDB4"},
        [2]string{"lightyellow4", "#8B8B7A"},
        [2]string{"limegreen", "#32CD32"},
        [2]string{"linen", "#FAF0E6"},
        [2]string{"magenta", "#FF00FF"},
        [2]string{"magenta1", "#FF00FF"},
        [2]string{"magenta2", "#EE00EE"},
        [2]string{"magenta3", "#CD00CD"},
        [2]string{"magenta4", "#8B008B"},
        [2]string{"maroon", "#B03060"},
        [2]string{"maroon1", "#FF34B3"},
        [2]string{"maroon2", "#EE30A7"},
        [2]string{"maroon3", "#CD2990"},
        [2]string{"maroon4", "#8B1C62"},
        [2]string{"mediumaquamarine", "#66CDAA"},
        [2]string{"mediumblue", "#0000CD"},
        [2]string{"mediumorchid", "#BA55D3"},
        [2]string{"mediumorchid1", "#E066FF"},
        [2]string{"mediumorchid2", "#D15FEE"},
        [2]string{"mediumorchid3", "#B452CD"},
        [2]string{"mediumorchid4", "#7A378B"},
        [2]string{"mediumpurple", "#9370DB"},
        [2]string{"mediumpurple1", "#AB82FF"},
        [2]string{"mediumpurple2", "#9F79EE"},
        [2]string{"mediumpurple3", "#8968CD"},
        [2]string{"mediumpurple4", "#5D478B"},
        [2]string{"mediumseagreen", "#3CB371"},
        [2]string{"mediumslateblue", "#7B68EE"},
        [2]string{"mediumspringgreen", "#00FA9A"},
        [2]string{"mediumturquoise", "#48D1CC"},
        [2]string{"mediumvioletred", "#C71585"},
        [2]string{"midnightblue", "#191970"},
        [2]string{"mintcream", "#F5FFFA"},
        [2]string{"mistyrose", "#FFE4E1"},
        [2]string{"mistyrose1", "#FFE4E1"},
        [2]string{"mistyrose2", "#EED5D2"},
        [2]string{"mistyrose3", "#CDB7B5"},
        [2]string{"mistyrose4", "#8B7D7B"},
        [2]string{"moccasin", "#FFE4B5"},
        [2]string{"navajowhite", "#FFDEAD"},
        [2]string{"navajowhite1", "#FFDEAD"},
        [2]string{"navajowhite2", "#EECFA1"},
        [2]string{"navajowhite3", "#CDB38B"},
        [2]string{"navajowhite4", "#8B795E"},
        [2]string{"navy", "#000080"},
        [2]string{"navyblue", "#000080"},
        [2]string{"oldlace", "#FDF5E6"},
        [2]string{"olivedrab", "#6B8E23"},
        [2]string{"olivedrab1", "#C0FF3E"},
        [2]string{"olivedrab2", "#B3EE3A"},
        [2]string{"olivedrab3", "#9ACD32"},
        [2]string{"olivedrab4", "#698B22"},
        [2]string{"orange", "#FFA500"},
        [2]string{"orange1", "#FFA500"},
        [2]string{"orange2", "#EE9A00"},
        [2]string{"orange3", "#CD8500"},
        [2]string{"orange4", "#8B5A00"},
        [2]string{"orangered", "#FF4500"},
        [2]string{"orangered1", "#FF4500"},
        [2]string{"orangered2", "#EE4000"},
        [2]string{"orangered3", "#CD3700"},
        [2]string{"orangered4", "#8B2500"},
        [2]string{"orchid", "#DA70D6"},
        [2]string{"orchid1", "#FF83FA"},
        [2]string{"orchid2", "#EE7AE9"},
        [2]string{"orchid3", "#CD69C9"},
        [2]string{"orchid4", "#8B4789"},
        [2]string{"palegoldenrod", "#EEE8AA"},
        [2]string{"palegreen", "#98FB98"},
        [2]string{"palegreen1", "#9AFF9A"},
        [2]string{"palegreen2", "#90EE90"},
        [2]string{"palegreen3", "#7CCD7C"},
        [2]string{"palegreen4", "#548B54"},
        [2]string{"paleturquoise", "#AFEEEE"},
        [2]string{"paleturquoise1", "#BBFFFF"},
        [2]string{"paleturquoise2", "#AEEEEE"},
        [2]string{"paleturquoise3", "#96CDCD"},
        [2]string{"paleturquoise4", "#668B8B"},
        [2]string{"palevioletred", "#DB7093"},
        [2]string{"palevioletred1", "#FF82AB"},
        [2]string{"palevioletred2", "#EE799F"},
        [2]string{"palevioletred3", "#CD6889"},
        [2]string{"palevioletred4", "#8B475D"},
        [2]string{"papayawhip", "#FFEFD5"},
        [2]string{"peachpuff", "#FFDAB9"},
        [2]string{"peachpuff1", "#FFDAB9"},
        [2]string{"peachpuff2", "#EECBAD"},
        [2]string{"peachpuff3", "#CDAF95"},
        [2]string{"peachpuff4", "#8B7765"},
        [2]string{"peru", "#CD853F"},
        [2]string{"pink", "#FFC0CB"},
        [2]string{"pink1", "#FFB5C5"},
        [2]string{"pink2", "#EEA9B8"},
        [2]string{"pink3", "#CD919E"},
        [2]string{"pink4", "#8B636C"},
        [2]string{"plum", "#DDA0DD"},
        [2]string{"plum1", "#FFBBFF"},
        [2]string{"plum2", "#EEAEEE"},
        [2]string{"plum3", "#CD96CD"},
        [2]string{"plum4", "#8B668B"},
        [2]string{"powderblue", "#B0E0E6"},
        [2]string{"purple", "#A020F0"},
        [2]string{"purple1", "#9B30FF"},
        [2]string{"purple2", "#912CEE"},
        [2]string{"purple3", "#7D26CD"},
        [2]string{"purple4", "#551A8B"},
        [2]string{"red", "#FF0000"},
        [2]string{"red1", "#FF0000"},
        [2]string{"red2", "#EE0000"},
        [2]string{"red3", "#CD0000"},
        [2]string{"red4", "#8B0000"},
        [2]string{"rosybrown", "#BC8F8F"},
        [2]string{"rosybrown1", "#FFC1C1"},
        [2]string{"rosybrown2", "#EEB4B4"},
        [2]string{"rosybrown3", "#CD9B9B"},
        [2]string{"rosybrown4", "#8B6969"},
        [2]string{"royalblue", "#4169E1"},
        [2]string{"royalblue1", "#4876FF"},
        [2]string{"royalblue2", "#436EEE"},
        [2]string{"royalblue3", "#3A5FCD"},
        [2]string{"royalblue4", "#27408B"},
        [2]string{"saddlebrown", "#8B4513"},
        [2]string{"salmon", "#FA8072"},
        [2]string{"salmon1", "#FF8C69"},
        [2]string{"salmon2", "#EE8262"},
        [2]string{"salmon3", "#CD7054"},
        [2]string{"salmon4", "#8B4C39"},
        [2]string{"sandybrown", "#F4A460"},
        [2]string{"seagreen", "#2E8B57"},
        [2]string{"seagreen1", "#54FF9F"},
        [2]string{"seagreen2", "#4EEE94"},
        [2]string{"seagreen3", "#43CD80"},
        [2]string{"seagreen4", "#2E8B57"},
        [2]string{"seashell", "#FFF5EE"},
        [2]string{"seashell1", "#FFF5EE"},
        [2]string{"seashell2", "#EEE5DE"},
        [2]string{"seashell3", "#CDC5BF"},
        [2]string{"seashell4", "#8B8682"},
        [2]string{"sienna", "#A0522D"},
        [2]string{"sienna1", "#FF8247"},
        [2]string{"sienna2", "#EE7942"},
        [2]string{"sienna3", "#CD6839"},
        [2]string{"sienna4", "#8B4726"},
        [2]string{"skyblue", "#87CEEB"},
        [2]string{"skyblue1", "#87CEFF"},
        [2]string{"skyblue2", "#7EC0EE"},
        [2]string{"skyblue3", "#6CA6CD"},
        [2]string{"skyblue4", "#4A708B"},
        [2]string{"slateblue", "#6A5ACD"},
        [2]string{"slateblue1", "#836FFF"},
        [2]string{"slateblue2", "#7A67EE"},
        [2]string{"slateblue3", "#6959CD"},
        [2]string{"slateblue4", "#473C8B"},
        [2]string{"slategray", "#708090"},
        [2]string{"slategray1", "#C6E2FF"},
        [2]string{"slategray2", "#B9D3EE"},
        [2]string{"slategray3", "#9FB6CD"},
        [2]string{"slategray4", "#6C7B8B"},
        [2]string{"slategrey", "#708090"},
        [2]string{"snow", "#FFFAFA"},
        [2]string{"snow1", "#FFFAFA"},
        [2]string{"snow2", "#EEE9E9"},
        [2]string{"snow3", "#CDC9C9"},
        [2]string{"snow4", "#8B8989"},
        [2]string{"springgreen", "#00FF7F"},
        [2]string{"springgreen1", "#00FF7F"},
        [2]string{"springgreen2", "#00EE76"},
        [2]string{"springgreen3", "#00CD66"},
        [2]string{"springgreen4", "#008B45"},
        [2]string{"steelblue", "#4682B4"},
        [2]string{"steelblue1", "#63B8FF"},
        [2]string{"steelblue2", "#5CACEE"},
        [2]string{"steelblue3", "#4F94CD"},
        [2]string{"steelblue4", "#36648B"},
        [2]string{"tan", "#D2B48C"},
        [2]string{"tan1", "#FFA54F"},
        [2]string{"tan2", "#EE9A49"},
        [2]string{"tan3", "#CD853F"},
        [2]string{"tan4", "#8B5A2B"},
        [2]string{"thistle", "#D8BFD8"},
        [2]string{"thistle1", "#FFE1FF"},
        [2]string{"thistle2", "#EED2EE"},
        [2]string{"thistle3", "#CDB5CD"},
        [2]string{"thistle4", "#8B7B8B"},
        [2]string{"tomato", "#FF6347"},
        [2]string{"tomato1", "#FF6347"},
        [2]string{"tomato2", "#EE5C42"},
        [2]string{"tomato3", "#CD4F39"},
        [2]string{"tomato4", "#8B3626"},
        [2]string{"turquoise", "#40E0D0"},
        [2]string{"turquoise1", "#00F5FF"},
        [2]string{"turquoise2", "#00E5EE"},
        [2]string{"turquoise3", "#00C5CD"},
        [2]string{"turquoise4", "#00868B"},
        [2]string{"violet", "#EE82EE"},
        [2]string{"violetred", "#D02090"},
        [2]string{"violetred1", "#FF3E96"},
        [2]string{"violetred2", "#EE3A8C"},
        [2]string{"violetred3", "#CD3278"},
        [2]string{"violetred4", "#8B2252"},
        [2]string{"wheat", "#F5DEB3"},
        [2]string{"wheat1", "#FFE7BA"},
        [2]string{"wheat2", "#EED8AE"},
        [2]string{"wheat3", "#CDBA96"},
        [2]string{"wheat4", "#8B7E66"},
        [2]string{"whitesmoke", "#F5F5F5"},
        [2]string{"yellow", "#FFFF00"},
        [2]string{"yellow1", "#FFFF00"},
        [2]string{"yellow2", "#EEEE00"},
        [2]string{"yellow3", "#CDCD00"},
        [2]string{"yellow4", "#8B8B00"},
        [2]string{"yellowgreen", "#9ACD32"},
	}
	webColoursDisabled = false
)

func parseWebColours(line string) colours {
	var clrs colours
	if webColoursDisabled {
		return clrs
	}

	used := make([]bool, len(line))
	line = strings.ToLower(line)
	for _, tuple := range webColours {
		curLine := line
		for len(curLine) > 0 {
			offset := len(line) - len(curLine)
			index := strings.Index(curLine, tuple[0])
			if index != -1 {
				if !used[offset+index] {
					if !checkBoundary || isWord(line, offset+index, offset+index+len(tuple[0])) {
						colour := &Colour{
							ColStart: offset + index + 1,
							ColEnd:   offset + index + len(tuple[0]),
							Hex:      tuple[1],
							Line:     line,
						}
						clrs = append(clrs, colour)
						for i := offset + index; i < offset+index+len(tuple[0]); i++ {
							used[i] = true
						}
					}
				}
				curLine = curLine[index+len(tuple[0]):]
			} else {
				break
			}
		}
	}
	return clrs
}
