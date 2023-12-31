package converter_test

import (
	"testing"

	"github.com/ArmadaSuit/udf-go/converter"
)

func TestHankakuEnglishToZenkakuEnglish(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"},
			want: " !\"#$%&'()*+,-./0123456789:;<=>?@ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ[\\]^_`ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ{|}~",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuEnglishToZenkakuEnglish(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuEnglishToHankakuEnglish(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～"},
			want: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ABCDEFGHIJKLMNOPQRSTUVWXYZ［＼］＾＿｀abcdefghijklmnopqrstuvwxyz｛｜｝～",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuEnglishToHankakuEnglish(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestHankakuNumberToZenkakuNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"},
			want: " !\"#$%&'()*+,-./０１２３４５６７８９:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuNumberToZenkakuNumber(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuNumberToHankakuNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～"},
			want: " ！”＃＄％＆＇（）＊＋，－．／0123456789：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuNumberToHankakuNumber(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestHankakuEnglishNumberToZenkakuEnglishNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ．",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"},
			want: " ！\"＃＄％＆'（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［\\］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝~",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuEnglishNumberToZenkakuEnglishNumber(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuEnglishNumberToHankakuEnglishNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ．"},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ascii",
			args: args{in: " ！”＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～"},
			want: " !”#$%&＇()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[＼]^_`abcdefghijklmnopqrstuvwxyz{|}～",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuEnglishNumberToHankakuEnglishNumber(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestHankakuSpaceToZenkakuSpace(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The　quick　brown　fox　jumps　over　the　lazy　dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと ちりぬるを わかよたれそ つねならむ うゐのおくやま けふこえて あさきゆめみし ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "spaces",
			args: args{in: "     　　　　　"},
			want: "　　　　　　　　　　",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuSpaceToZenkakuSpace(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuSpaceToHankakuSpace(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The　quick　brown　fox　jumps　over　the　lazy　dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと ちりぬるを わかよたれそ つねならむ うゐのおくやま けふこえて あさきゆめみし ゑひもせす",
		},
		{
			name: "spaces",
			args: args{in: "     　　　　　"},
			want: "          ",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuSpaceToHankakuSpace(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuKatakanaToHankakuKatakana(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをんﾞﾟゝゞ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "ｲﾛﾊﾆﾎﾍﾄ　ﾁﾘﾇﾙｦ　ﾜｶﾖﾀﾚｿ　ﾂﾈﾅﾗﾑ　ｳｲﾉｵｸﾔﾏ　ｹﾌｺｴﾃ　ｱｻｷﾕﾒﾐｼ　ｴﾋﾓｾｽ",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜﾜｲｴｦﾝｳﾞヵヶﾞﾟヽヾ",
		},
		{
			name: "all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ"},
			want: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "､｡｢｣･ｰ",
		},
		{
			name: "japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ"},
			want: "､｡｢｣･ｰ",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuKatakanaToHankakuKatakana(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestHankakuKatakanaToZenkakuKatakana(t *testing.T) {
	type args struct {
		in string
		v  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ",
		},
		{
			name: "all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝｳﾞﾞﾟ"},
			want: "ァアィイゥウェエォオカカ゛キキ゛クク゛ケケ゛ココ゛ササ゛シシ゛スス゛セセ゛ソソ゛タタ゛チチ゛ッツツ゛テテ゛トト゛ナニヌネノハハ゛ハ゜ヒヒ゛ヒ゜フフ゛フ゜ヘヘ゛ヘ゜ホホ゛ホ゜マミムメモャヤュユョヨラリルレロワヲンウ゛゛゜",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "、。「」・ー",
		},
		{
			name: "japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ"},
			want: "、。「」・ー",
		},
		{
			name: "ligature version: english",
			args: args{in: "The quick brown fox jumps over the lazy dog.", v: true},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "ligature version: zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす", v: true},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ligature version: all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ", v: true},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ",
		},
		{
			name: "ligature version: zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス", v: true},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "ligature version: all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ", v: true},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ",
		},
		{
			name: "ligature version: all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝｳﾞﾞﾟ", v: true},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロワヲンヴ゛゜",
		},
		{
			name: "ligature version: japanese zenkaku marks",
			args: args{in: "、。「」・ー", v: true},
			want: "、。「」・ー",
		},
		{
			name: "ligature version: japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ", v: true},
			want: "、。「」・ー",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuKatakanaToZenkakuKatakana(converter.GenerateForKanaConverter(tt.args.in), tt.args.v)); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuHiraganaToHankakuKatakana(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "ｲﾛﾊﾆﾎﾍﾄ　ﾁﾘﾇﾙｦ　ﾜｶﾖﾀﾚｿ　ﾂﾈﾅﾗﾑ　ｳｲﾉｵｸﾔﾏ　ｹﾌｺｴﾃ　ｱｻｷﾕﾒﾐｼ　ｴﾋﾓｾｽ",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜﾜｲｴｦﾝﾞﾟゝゞ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶﾞﾟヽヾ",
		},
		{
			name: "all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ"},
			want: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "､｡｢｣･ｰ",
		},
		{
			name: "japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ"},
			want: "､｡｢｣･ｰ",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuHiraganaToHankakuKatakana(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestHankakuKatakanaToZenkakuHiragana(t *testing.T) {
	type args struct {
		in string
		v  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ",
		},
		{
			name: "all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ"},
			want: "ぁあぃいぅうぇえぉおかか゛きき゛くく゛けけ゛ここ゛ささ゛しし゛すす゛せせ゛そそ゛たた゛ちち゛っつつ゛てて゛とと゛なにぬねのはは゛は゜ひひ゛ひ゜ふふ゛ふ゜へへ゛へ゜ほほ゛ほ゜まみむめもゃやゅゆょよらりるれろわをん゛゜",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "、。「」・ー",
		},
		{
			name: "japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ"},
			want: "、。「」・ー",
		},
		{
			name: "ligature version: english",
			args: args{in: "The quick brown fox jumps over the lazy dog.", v: true},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "ligature version: zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす", v: true},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "ligature version: all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ", v: true},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ",
		},
		{
			name: "ligature version: zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス", v: true},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "ligature version: all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ", v: true},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ",
		},
		{
			name: "ligature version: all hankaku katakana",
			args: args{in: "ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜｦﾝﾞﾟ", v: true},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろわをん゛゜",
		},
		{
			name: "ligature version: japanese zenkaku marks",
			args: args{in: "、。「」・ー", v: true},
			want: "、。「」・ー",
		},
		{
			name: "ligature version: japanese hankaku marks",
			args: args{in: "､｡｢｣･ｰ", v: true},
			want: "、。「」・ー",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.HankakuKatakanaToZenkakuHiragana(converter.GenerateForKanaConverter(tt.args.in), tt.args.v)); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuKatakanaToZenkakuHiragana(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをんヴヵヶ゛゜ゝゞ",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "、。「」・ー",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuKatakanaToZenkakuHiragana(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestZenkakuHiraganaToZenkakuKatakana(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{in: "The quick brown fox jumps over the lazy dog."},
			want: "The quick brown fox jumps over the lazy dog.",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "all zenkaku hiragana",
			args: args{in: "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん゛゜ゝゞ"},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲン゛゜ヽヾ",
		},
		{
			name: "zenkaku katakana",
			args: args{in: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス"},
			want: "イロハニホヘト　チリヌルヲ　ワカヨタレソ　ツネナラム　ウヰノオクヤマ　ケフコエテ　アサキユメミシ　ヱヒモセス",
		},
		{
			name: "all zenkaku katakana",
			args: args{in: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ"},
			want: "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ゛゜ヽヾ",
		},
		{
			name: "japanese zenkaku marks",
			args: args{in: "、。「」・ー"},
			want: "、。「」・ー",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.StringForKanaConverter(converter.ZenkakuHiraganaToZenkakuKatakana(converter.GenerateForKanaConverter(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

func TestNewKanaConverters(t *testing.T) {
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

			converters, err := converter.NewKanaConverters(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKanaConverters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			in := converter.GenerateForKanaConverter(tt.args.in)
			for _, c := range converters {
				in = c(in)
			}
			if got := converter.StringForKanaConverter(in); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.mode, got, tt.want)
			}
		})
	}
}
