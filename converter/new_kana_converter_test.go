package converter_test

import (
	"bytes"
	"testing"

	"github.com/ArmadaSuit/udf-go/converter"
	"golang.org/x/text/transform"
)

func TestNewKanaConverter(t *testing.T) {
	type args struct {
		in   string
		mode string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "hankaku english -> zenkaku english and zenkaku space -> hankaku space",
			args: args{in: "Ｔｈｅ　ｑｕｉｃｋ　ｂｒｏｗｎ　ｆｏｘ　ｊｕｍｐｓ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ.", mode: "rs"},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "hankaku english -> zenkaku english and hankaku space -> zenkaku space",
			args: args{in: "The quick brown fox jumps over the lazy dog.", mode: "RS"},
			want: "Ｔｈｅ　ｑｕｉｃｋ　ｂｒｏｗｎ　ｆｏｘ　ｊｕｍｐｓ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ.",
		},
		{
			name: "zenkaku english -> hankaku english and zenkaku number -> hankaku number: exclude marks",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～", mode: "rn"},
			want: " ！”＃＄％＆＇（）＊＋，－．／0123456789：；＜＝＞？＠ABCDEFGHIJKLMNOPQRSTUVWXYZ［＼］＾＿｀abcdefghijklmnopqrstuvwxyz｛｜｝～",
		},
		{
			name: "zenkaku english -> hankaku english and zenkaku number -> hankaku number: include marks",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～", mode: "a"},
			want: " !”#$%&＇()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[＼]^_`abcdefghijklmnopqrstuvwxyz{|}～",
		},
		{
			name: "zenkaku english -> hankaku english and zenkaku number -> hankaku number: include marks: verbose version",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～", mode: "arn"},
			want: " !”#$%&＇()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[＼]^_`abcdefghijklmnopqrstuvwxyz{|}～",
		},
		{
			name: "hankaku english -> zenkaku english and hankaku number -> zenkaku number: exclude marks",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", mode: "RN"},
			want: " !\"#$%&'()*+,-./０１２３４５６７８９:;<=>?@ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ[\\]^_`ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ{|}~",
		},
		{
			name: "hankaku english -> zenkaku english and hankaku number -> zenkaku number: include marks",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", mode: "A"},
			want: " ！\"＃＄％＆'（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［\\］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝~",
		},
		{
			name: "hankaku english -> zenkaku english and hankaku number -> zenkaku number: include marks: verbose version",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", mode: "ARN"},
			want: " ！\"＃＄％＆'（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［\\］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝~",
		},
		{
			name: "zenkaku english -> hankaku english and hankaku number -> zenkaku number",
			args: args{in: "1600 Ｐｅｎｎｓｙｌｖａｎｉａ Ａｖｅｎｕｅ", mode: "rN"},
			want: "１６００ Pennsylvania Avenue",
		},
		{
			name: "hankaku english -> zenkaku english and zenkaku number -> hanaku number",
			args: args{in: "１６００ Pennsylvania Avenue", mode: "Rn"},
			want: "1600 Ｐｅｎｎｓｙｌｖａｎｉａ Ａｖｅｎｕｅ",
		},
		{
			name: "zenkaku space -> hankaku space and zenkaku katakana -> hankaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス", mode: "sk"},
			want: "ｲﾛﾊﾆﾎﾍﾄ ﾁﾘﾇﾙｦ ﾜｶﾖﾀﾚｿ ﾂﾈﾅﾗﾑ ｳｲﾉｵｸﾔﾏ ｹﾌｺｴﾃ ｱｻｷﾕﾒﾐｼ ｴﾋﾓｾｽ",
		},
		{
			name: "hankaku space -> zenkaku space and zenkaku katakana -> hankaku katakana",
			args: args{in: "イロハニホヘト チリヌルヲ ワカヨタレソ ツネナラム ウヰノオクヤマ ケフコエテ アサキユメミシ ヱヒモセス", mode: "Sk"},
			want: "ｲﾛﾊﾆﾎﾍﾄ　ﾁﾘﾇﾙｦ　ﾜｶﾖﾀﾚｿ　ﾂﾈﾅﾗﾑ　ｳｲﾉｵｸﾔﾏ　ｹﾌｺｴﾃ　ｱｻｷﾕﾒﾐｼ　ｴﾋﾓｾｽ",
		},
		{
			name: "zenkaku space -> hankaku space and hankaku katakana -> zenkaku katakana",
			args: args{in: "ｱｷﾉﾀﾉ　ｶﾘﾎﾉｲﾎﾉ　ﾄﾏｦｱﾗﾐ　ﾜｶﾞｺﾛﾓﾃﾞﾊ　ﾂﾕﾆﾇﾚﾂﾂ", mode: "sK"},
			want: "アキノタノ カリホノイホノ トマヲアラミ ワカ゛コロモテ゛ハ ツユニヌレツツ",
		},
		{
			name: "hankaku space -> zenkaku space and hankaku katakana -> zenkaku katakana",
			args: args{in: "ｱｷﾉﾀﾉ ｶﾘﾎﾉｲﾎﾉ ﾄﾏｦｱﾗﾐ ﾜｶﾞｺﾛﾓﾃﾞﾊ ﾂﾕﾆﾇﾚﾂﾂ", mode: "SK"},
			want: "アキノタノ　カリホノイホノ　トマヲアラミ　ワカ゛コロモテ゛ハ　ツユニヌレツツ",
		},
		{
			name: "zenkaku space -> hankaku space and zenkaku katakana -> zenkaku hiragana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス", mode: "sc"},
			want: "いろはにほへと ちりぬるを わかよたれそ つねならむ うゐのおくやま けふこえて あさきゆめみし ゑひもせす",
		},
		{
			name: "hankaku space -> zenkaku space and zenkaku katakana -> zenkaku hiragana",
			args: args{in: "イロハニホヘト チリヌルヲ ワカヨタレソ ツネナラム ウヰノオクヤマ ケフコエテ アサキユメミシ ヱヒモセス", mode: "Sc"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "zenkaku space -> hankaku space and zenkaku hiragana -> zekaku katakana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす", mode: "sC"},
			want: "イロハニホヘト チリヌルヲ ワカヨタレソ ツネナラム ウヰノオクヤマ ケフコエテ アサキユメミシ ヱヒモセス",
		},
		{
			name: "hankaku space -> zenkaku space and zenkaku hiragana -> zekaku katakana",
			args: args{in: "いろはにほへと ちりぬるを わかよたれそ つねならむ うゐのおくやま けふこえて あさきゆめみし ゑひもせす", mode: "SC"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "zenkaku katakana and zenkaku hiragana -> hankaku katakana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "k"},
			want: "｢ﾎﾞｰﾙﾍﾟﾝの芯の太さは､0.7mmです｡｣",
		},
		{
			name: "hankaku katakana -> zenkaku katakana: ligature version",
			args: args{in: "｢ﾎﾞｰﾙﾍﾟﾝの芯の太さは､0.7mmです｡｣", mode: "KV"},
			want: "「ボールペンの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku katakana -> zenkaku katakana",
			args: args{in: "｢ﾎﾞｰﾙﾍﾟﾝの芯の太さは､0.7mmです｡｣", mode: "K"},
			want: "「ホ゛ールヘ゜ンの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and hankaku katakana -> zenkaku katakana",
			args: args{in: "｢ﾎﾞｰﾙﾍﾟﾝの芯の太さは､0.7mmです｡｣", mode: "AK"},
			want: "「ホ゛ールヘ゜ンの芯の太さは、０．７ｍｍです。」",
		},
		{
			name: "zenkaku hiragana -> hanaku katakana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "h"},
			want: "｢ボｰルペンﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana: ligature version",
			args: args{in: "｢ボールペンﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣", mode: "HV"},
			want: "「ボールペンの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana",
			args: args{in: "｢ボールペンﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣", mode: "H"},
			want: "「ボールペンの芯の太さは、0.7mmて゛す。」",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and hankaku katakana -> zenkaku hiragana",
			args: args{in: "｢ボールペンﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣", mode: "AH"},
			want: "「ボールペンの芯の太さは、０．７ｍｍて゛す。」",
		},
		{
			name: "zenkaku katakana -> zekaku hiragana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "c"},
			want: "「ぼーるぺんの芯の太さは、0.7mmです。」",
		},
		{
			name: "zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "C"},
			want: "「ボールペンノ芯ノ太サハ、0.7mmデス。」",
		},
		{
			name: "zenkaku katakana and zenkaku hiragana -> hankaku katakana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "kh"},
			want: "｢ﾎﾞｰﾙﾍﾟﾝﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ボールペンﾉ芯ﾉ太ｻﾊ、0.7mmﾃﾞｽ。」", mode: "kHV"},
			want: "｢ﾎﾞｰﾙﾍﾟﾝの芯の太さは､0.7mmです｡｣",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ﾎﾞｰﾙﾍﾟﾝの芯の太さは、0.7mmです。」", mode: "KhV"},
			want: "｢ボールペンﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣",
		},
		{
			name: "hankaku katakana -> zenkaku katakana: ligature version: single character",
			args: args{in: "ｳ", mode: "KV"},
			want: "ウ",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana: ligature version: single character",
			args: args{in: "ｳ", mode: "HV"},
			want: "う",
		},
		{
			name: "hankaku katakana -> zenkaku katakana: ligature version: will be voiced sound and will not be semi-voiced sound",
			args: args{in: "ｳﾟ", mode: "KV"},
			want: "ウ゜",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana: ligature version: will be voiced sound and will not be semi-voiced sound",
			args: args{in: "ｳﾟ", mode: "HV"},
			want: "う゜",
		},
		{
			name: "zenkaku katakana -> hankaku katakana and hanaku katakana -> zenkaku hiragana: common symbol in katakana and hiragana",
			args: args{in: "、。「」・ー゛゜､｡｢｣･ｰﾞﾟ", mode: "kH"},
			want: "､｡｢｣･ｰﾞﾟ、。「」・ー゛゜",
		},
		{
			name: "hankaku katakana -> zenkaku katakana and zenkaku katakana -> hankaku katakana: common symbol in katakana and hiragana",
			args: args{in: "、。「」・ー゛゜､｡｢｣･ｰﾞﾟ", mode: "Kh"},
			want: "､｡｢｣･ｰﾞﾟ、。「」・ー゛゜",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana and zenkaku hiragana -> zekaku katakana: common symbol in katakana and hiragana",
			args: args{in: "、。「」・ー゛゜､｡｢｣･ｰﾞﾟ", mode: "HC"},
			want: "、。「」・ー゛゜、。「」・ー゛゜",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ﾎﾞｰﾙﾍﾟﾝノ芯ノ太サハ、0.7mmデス。」", mode: "KcV"},
			want: "「ボールペンの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku english -> zenkakuenglish, hankaku number -> zenkaku number and zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ﾎﾞｰﾙﾍﾟﾝの芯の太さは、0.7mmです。」", mode: "KCV"},
			want: "「ボールペンノ芯ノ太サハ、0.7mmデス。」",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana and zenkaku katakana -> zekaku hiragana",
			args: args{in: "｢ﾎﾞｰﾙﾍﾟﾝノ芯ノ太サハ､0.7mmデス｡｣", mode: "HVc"},
			want: "「ぼーるぺんの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku katakana -> zenkaku hiragana and zenkaku hiragana -> zekaku katakana",
			args: args{in: "｢ぼーるぺんﾉ芯ﾉ太ｻﾊ､0.7mmﾃﾞｽ｡｣", mode: "HVC"},
			want: "「ボールペンの芯の太さは、0.7mmです。」",
		},
		{
			name: "hankaku english -> zenkaku english, hankaku number -> zenkaku number and zenkaku katakana -> zekaku hiragana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "Ac"},
			want: "「ぼーるぺんの芯の太さは、０．７ｍｍです。」",
		},
		{
			name: "hankaku english -> zenkakuenglish, hankaku number -> zenkaku number and zenkaku hiragana -> zekaku katakana",
			args: args{in: "「ボールペンの芯の太さは、0.7mmです。」", mode: "AC"},
			want: "「ボールペンノ芯ノ太サハ、０．７ｍｍデス。」",
		},
		{
			name:    "invalid option for english",
			args:    args{mode: "rR"},
			wantErr: true,
		},
		{
			name:    "invalid option for english: defferent order case",
			args:    args{mode: "Rr"},
			wantErr: true,
		},
		{
			name:    "invalid option for number",
			args:    args{mode: "nN"},
			wantErr: true,
		},
		{
			name:    "invalid option for number: defferent order case",
			args:    args{mode: "Nn"},
			wantErr: true,
		},
		{
			name:    "invalid option for english number",
			args:    args{mode: "aA"},
			wantErr: true,
		},
		{
			name:    "invalid option for english number: defferent order case",
			args:    args{mode: "Aa"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of english: case 1",
			args:    args{mode: "rA"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of english: case 1: defferent order case",
			args:    args{mode: "Ar"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of english: case 2",
			args:    args{mode: "Ra"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of english: case 2: defferent order case",
			args:    args{mode: "aR"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of number: case 1",
			args:    args{mode: "nA"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of number: case 1: defferent order case",
			args:    args{mode: "An"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of number: case 2",
			args:    args{mode: "Na"},
			wantErr: true,
		},
		{
			name:    "invalid option for cyclic conversion of number: case 2: defferent order case",
			args:    args{mode: "aN"},
			wantErr: true,
		},
		{
			name:    "invalid option for space",
			args:    args{mode: "sS"},
			wantErr: true,
		},
		{
			name:    "invalid option for space: defferent order case",
			args:    args{mode: "Ss"},
			wantErr: true,
		},
		{
			name:    "invalid option for katakana",
			args:    args{mode: "kK"},
			wantErr: true,
		},
		{
			name:    "invalid option for katakana: defferent order case",
			args:    args{mode: "Kk"},
			wantErr: true,
		},
		{
			name:    "invalid option for hankaku katakana",
			args:    args{mode: "KH"},
			wantErr: true,
		},
		{
			name:    "invalid option for hankaku katakana: defferent order case",
			args:    args{mode: "HK"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku katakana",
			args:    args{mode: "kc"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku katakana: defferent order case",
			args:    args{mode: "ck"},
			wantErr: true,
		},
		{
			name:    "invalid option for ambiguous conversion: case 1",
			args:    args{mode: "kC"},
			wantErr: true,
		},
		{
			name:    "invalid option for ambiguous conversion: case 1: defferent order case",
			args:    args{mode: "Ck"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku hiragana - hankaku katakana",
			args:    args{mode: "hH"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku hiragana - hankaku katakana: defferent order case",
			args:    args{mode: "Hh"},
			wantErr: true,
		},
		{
			name:    "invalid option for ambiguous conversion: case 2",
			args:    args{mode: "hc"},
			wantErr: true,
		},
		{
			name:    "invalid option for ambiguous conversion: case 2: defferent order case",
			args:    args{mode: "ch"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku hiragana",
			args:    args{mode: "hC"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku hiragana: defferent order case",
			args:    args{mode: "Ch"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku katakana - zenkaku katakana",
			args:    args{mode: "cC"},
			wantErr: true,
		},
		{
			name:    "invalid option for zenkaku katakana - zenkaku katakana: defferent order case",
			args:    args{mode: "Cc"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			tran, err := converter.NewKanaConverter(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKanaConverters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			var got bytes.Buffer
			w := transform.NewWriter(&got, tran)
			if _, err := w.Write([]byte(tt.args.in)); err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("NewKanaConverter() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if err := w.Close(); err != nil {
				t.Error(err)
			}

			if tt.want != got.String() {
				t.Errorf("%v is converted %v, want %v", tt.args.mode, got.String(), tt.want)
			}
		})
	}
}
