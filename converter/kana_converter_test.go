package converter_test

import (
	"testing"

	"github.com/ArmadaSuit/udf-go/converter"
)

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
			name: "numbers",
			args: args{in: "0123456789"},
			want: "０１２３４５６７８９",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.String(converter.HankakuNumberToZenkakuNumber(converter.Generate(tt.args.in))); got != tt.want {
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
			name: "numbers",
			args: args{in: "０１２３４５６７８９"},
			want: "0123456789",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.String(converter.ZenkakuNumberToHankakuNumber(converter.Generate(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}

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
			name: "all english alphabet",
			args: args{in: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"},
			want: "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "numbers",
			args: args{in: "0123456789"},
			want: "0123456789",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.String(converter.HankakuEnglishToZenkakuEnglish(converter.Generate(tt.args.in))); got != tt.want {
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
			name: "all english alphabet",
			args: args{in: "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ"},
			want: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "zenkaku hiragana",
			args: args{in: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす"},
			want: "いろはにほへと　ちりぬるを　わかよたれそ　つねならむ　うゐのおくやま　けふこえて　あさきゆめみし　ゑひもせす",
		},
		{
			name: "numbers",
			args: args{in: "0123456789"},
			want: "0123456789",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := converter.String(converter.ZenkakuEnglishToHankakuEnglish(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.HankakuSpaceToZenkakuSpace(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.ZenkakuSpaceToHankakuSpace(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.ZenkakuKatakanaToHankakuKatakana(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.HankakuKatakanaToZenkakuKatakana(converter.Generate(tt.args.in), tt.args.v)); got != tt.want {
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

			if got := converter.String(converter.ZenkakuHiraganaToHankakuKatakana(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.HankakuKatakanaToZenkakuHiragana(converter.Generate(tt.args.in), tt.args.v)); got != tt.want {
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

			if got := converter.String(converter.ZenkakuKatakanaToZenkakuHiragana(converter.Generate(tt.args.in))); got != tt.want {
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

			if got := converter.String(converter.ZenkakuHiraganaToZenkakuKatakana(converter.Generate(tt.args.in))); got != tt.want {
				t.Errorf("%v is converted %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}
