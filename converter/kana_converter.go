package converter

type KanaConverterRune struct {
	Rune        rune
	IsConverted bool
}

func HankakuEnglishToZenkakuEnglish(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune >= 'a' && r.Rune <= 'z':
				out <- KanaConverterRune{Rune: 'ａ' + r.Rune - 'a', IsConverted: true}
			case r.Rune >= 'A' && r.Rune <= 'Z':
				out <- KanaConverterRune{Rune: 'Ａ' + r.Rune - 'A', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuEnglishToHankakuEnglish(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune >= 'ａ' && r.Rune <= 'ｚ':
				out <- KanaConverterRune{Rune: 'a' + r.Rune - 'ａ', IsConverted: true}
			case r.Rune >= 'Ａ' && r.Rune <= 'Ｚ':
				out <- KanaConverterRune{Rune: 'A' + r.Rune - 'Ａ', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func HankakuNumberToZenkakuNumber(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			if r.Rune >= '0' && r.Rune <= '9' {
				out <- KanaConverterRune{Rune: r.Rune + 0xFEE0, IsConverted: true}
			} else {
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuNumberToHankakuNumber(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			if r.Rune >= '０' && r.Rune <= '９' {
				out <- KanaConverterRune{Rune: r.Rune - 0xFEE0, IsConverted: true}
			} else {
				out <- r
			}
		}
	}()
	return out
}

func HankakuEnglishNumberToZenkakuEnglishNumber(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune == '\u0022', r.Rune == '\u0027', r.Rune == '\u005C', r.Rune == '\u007E':
				out <- r
			case r.Rune >= '\u0021' && r.Rune <= '\u007E':
				out <- KanaConverterRune{Rune: r.Rune + 0xFEE0, IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuEnglishNumberToHankakuEnglishNumber(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune == '\uFF02', r.Rune == '\uFF07', r.Rune == '\uFF3C', r.Rune == '\uFF5E':
				out <- r
			case r.Rune >= '\uFF01' && r.Rune <= '\uFF5E':
				out <- KanaConverterRune{Rune: r.Rune - 0xFEE0, IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuSpaceToHankakuSpace(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			if r.Rune == '　' {
				out <- KanaConverterRune{Rune: ' ', IsConverted: true}
			} else {
				out <- r
			}
		}
	}()
	return out
}

func HankakuSpaceToZenkakuSpace(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			if r.Rune == ' ' {
				out <- KanaConverterRune{Rune: '　', IsConverted: true}
			} else {
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuKatakanaToHankakuKatakana(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch r.Rune {
			case '、':
				out <- KanaConverterRune{Rune: '､', IsConverted: true}
			case '。':
				out <- KanaConverterRune{Rune: '｡', IsConverted: true}
			case '「':
				out <- KanaConverterRune{Rune: '｢', IsConverted: true}
			case '」':
				out <- KanaConverterRune{Rune: '｣', IsConverted: true}
			case '゛':
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case '゜':
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ァ':
				out <- KanaConverterRune{Rune: 'ｧ', IsConverted: true}
			case 'ア':
				out <- KanaConverterRune{Rune: 'ｱ', IsConverted: true}
			case 'ィ':
				out <- KanaConverterRune{Rune: 'ｨ', IsConverted: true}
			case 'イ':
				out <- KanaConverterRune{Rune: 'ｲ', IsConverted: true}
			case 'ゥ':
				out <- KanaConverterRune{Rune: 'ｩ', IsConverted: true}
			case 'ウ':
				out <- KanaConverterRune{Rune: 'ｳ', IsConverted: true}
			case 'ェ':
				out <- KanaConverterRune{Rune: 'ｪ', IsConverted: true}
			case 'エ':
				out <- KanaConverterRune{Rune: 'ｴ', IsConverted: true}
			case 'ォ':
				out <- KanaConverterRune{Rune: 'ｫ', IsConverted: true}
			case 'オ':
				out <- KanaConverterRune{Rune: 'ｵ', IsConverted: true}
			case 'カ':
				out <- KanaConverterRune{Rune: 'ｶ', IsConverted: true}
			case 'ガ':
				out <- KanaConverterRune{Rune: 'ｶ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'キ':
				out <- KanaConverterRune{Rune: 'ｷ', IsConverted: true}
			case 'ギ':
				out <- KanaConverterRune{Rune: 'ｷ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ク':
				out <- KanaConverterRune{Rune: 'ｸ', IsConverted: true}
			case 'グ':
				out <- KanaConverterRune{Rune: 'ｸ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ケ':
				out <- KanaConverterRune{Rune: 'ｹ', IsConverted: true}
			case 'ゲ':
				out <- KanaConverterRune{Rune: 'ｹ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'コ':
				out <- KanaConverterRune{Rune: 'ｺ', IsConverted: true}
			case 'ゴ':
				out <- KanaConverterRune{Rune: 'ｺ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'サ':
				out <- KanaConverterRune{Rune: 'ｻ', IsConverted: true}
			case 'ザ':
				out <- KanaConverterRune{Rune: 'ｻ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'シ':
				out <- KanaConverterRune{Rune: 'ｼ', IsConverted: true}
			case 'ジ':
				out <- KanaConverterRune{Rune: 'ｼ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ス':
				out <- KanaConverterRune{Rune: 'ｽ', IsConverted: true}
			case 'ズ':
				out <- KanaConverterRune{Rune: 'ｽ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'セ':
				out <- KanaConverterRune{Rune: 'ｾ', IsConverted: true}
			case 'ゼ':
				out <- KanaConverterRune{Rune: 'ｾ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ソ':
				out <- KanaConverterRune{Rune: 'ｿ', IsConverted: true}
			case 'ゾ':
				out <- KanaConverterRune{Rune: 'ｿ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'タ':
				out <- KanaConverterRune{Rune: 'ﾀ', IsConverted: true}
			case 'ダ':
				out <- KanaConverterRune{Rune: 'ﾀ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'チ':
				out <- KanaConverterRune{Rune: 'ﾁ', IsConverted: true}
			case 'ヂ':
				out <- KanaConverterRune{Rune: 'ﾁ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ッ':
				out <- KanaConverterRune{Rune: 'ｯ', IsConverted: true}
			case 'ツ':
				out <- KanaConverterRune{Rune: 'ﾂ', IsConverted: true}
			case 'ヅ':
				out <- KanaConverterRune{Rune: 'ﾂ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'テ':
				out <- KanaConverterRune{Rune: 'ﾃ', IsConverted: true}
			case 'デ':
				out <- KanaConverterRune{Rune: 'ﾃ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ト':
				out <- KanaConverterRune{Rune: 'ﾄ', IsConverted: true}
			case 'ド':
				out <- KanaConverterRune{Rune: 'ﾄ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ナ':
				out <- KanaConverterRune{Rune: 'ﾅ', IsConverted: true}
			case 'ニ':
				out <- KanaConverterRune{Rune: 'ﾆ', IsConverted: true}
			case 'ヌ':
				out <- KanaConverterRune{Rune: 'ﾇ', IsConverted: true}
			case 'ネ':
				out <- KanaConverterRune{Rune: 'ﾈ', IsConverted: true}
			case 'ノ':
				out <- KanaConverterRune{Rune: 'ﾉ', IsConverted: true}
			case 'ハ':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
			case 'バ':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'パ':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ヒ':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
			case 'ビ':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ピ':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'フ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
			case 'ブ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'プ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ヘ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
			case 'ベ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ペ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ホ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
			case 'ボ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ポ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'マ':
				out <- KanaConverterRune{Rune: 'ﾏ', IsConverted: true}
			case 'ミ':
				out <- KanaConverterRune{Rune: 'ﾐ', IsConverted: true}
			case 'ム':
				out <- KanaConverterRune{Rune: 'ﾑ', IsConverted: true}
			case 'メ':
				out <- KanaConverterRune{Rune: 'ﾒ', IsConverted: true}
			case 'モ':
				out <- KanaConverterRune{Rune: 'ﾓ', IsConverted: true}
			case 'ャ':
				out <- KanaConverterRune{Rune: 'ｬ', IsConverted: true}
			case 'ヤ':
				out <- KanaConverterRune{Rune: 'ﾔ', IsConverted: true}
			case 'ュ':
				out <- KanaConverterRune{Rune: 'ｭ', IsConverted: true}
			case 'ユ':
				out <- KanaConverterRune{Rune: 'ﾕ', IsConverted: true}
			case 'ョ':
				out <- KanaConverterRune{Rune: 'ｮ', IsConverted: true}
			case 'ヨ':
				out <- KanaConverterRune{Rune: 'ﾖ', IsConverted: true}
			case 'ラ':
				out <- KanaConverterRune{Rune: 'ﾗ', IsConverted: true}
			case 'リ':
				out <- KanaConverterRune{Rune: 'ﾘ', IsConverted: true}
			case 'ル':
				out <- KanaConverterRune{Rune: 'ﾙ', IsConverted: true}
			case 'レ':
				out <- KanaConverterRune{Rune: 'ﾚ', IsConverted: true}
			case 'ロ':
				out <- KanaConverterRune{Rune: 'ﾛ', IsConverted: true}
			case 'ヮ':
				out <- KanaConverterRune{Rune: 'ﾜ', IsConverted: true}
			case 'ワ':
				out <- KanaConverterRune{Rune: 'ﾜ', IsConverted: true}
			case 'ヰ':
				out <- KanaConverterRune{Rune: 'ｲ', IsConverted: true}
			case 'ヱ':
				out <- KanaConverterRune{Rune: 'ｴ', IsConverted: true}
			case 'ヲ':
				out <- KanaConverterRune{Rune: 'ｦ', IsConverted: true}
			case 'ン':
				out <- KanaConverterRune{Rune: 'ﾝ', IsConverted: true}
			case 'ヴ':
				out <- KanaConverterRune{Rune: 'ｳ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case '・':
				out <- KanaConverterRune{Rune: '･', IsConverted: true}
			case 'ー':
				out <- KanaConverterRune{Rune: 'ｰ', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func hankakuKatakanaToZenkakuKatakanaSimple(in rune) KanaConverterRune {
	switch in {
	case '｡':
		return KanaConverterRune{Rune: '。', IsConverted: true}
	case '｢':
		return KanaConverterRune{Rune: '「', IsConverted: true}
	case '｣':
		return KanaConverterRune{Rune: '」', IsConverted: true}
	case '､':
		return KanaConverterRune{Rune: '、', IsConverted: true}
	case '･':
		return KanaConverterRune{Rune: '・', IsConverted: true}
	case 'ｦ':
		return KanaConverterRune{Rune: 'ヲ', IsConverted: true}
	case 'ｧ':
		return KanaConverterRune{Rune: 'ァ', IsConverted: true}
	case 'ｨ':
		return KanaConverterRune{Rune: 'ィ', IsConverted: true}
	case 'ｩ':
		return KanaConverterRune{Rune: 'ゥ', IsConverted: true}
	case 'ｪ':
		return KanaConverterRune{Rune: 'ェ', IsConverted: true}
	case 'ｫ':
		return KanaConverterRune{Rune: 'ォ', IsConverted: true}
	case 'ｬ':
		return KanaConverterRune{Rune: 'ャ', IsConverted: true}
	case 'ｭ':
		return KanaConverterRune{Rune: 'ュ', IsConverted: true}
	case 'ｮ':
		return KanaConverterRune{Rune: 'ョ', IsConverted: true}
	case 'ｯ':
		return KanaConverterRune{Rune: 'ッ', IsConverted: true}
	case 'ｰ':
		return KanaConverterRune{Rune: 'ー', IsConverted: true}
	case 'ｱ':
		return KanaConverterRune{Rune: 'ア', IsConverted: true}
	case 'ｲ':
		return KanaConverterRune{Rune: 'イ', IsConverted: true}
	case 'ｳ':
		return KanaConverterRune{Rune: 'ウ', IsConverted: true}
	case 'ｴ':
		return KanaConverterRune{Rune: 'エ', IsConverted: true}
	case 'ｵ':
		return KanaConverterRune{Rune: 'オ', IsConverted: true}
	case 'ｶ':
		return KanaConverterRune{Rune: 'カ', IsConverted: true}
	case 'ｷ':
		return KanaConverterRune{Rune: 'キ', IsConverted: true}
	case 'ｸ':
		return KanaConverterRune{Rune: 'ク', IsConverted: true}
	case 'ｹ':
		return KanaConverterRune{Rune: 'ケ', IsConverted: true}
	case 'ｺ':
		return KanaConverterRune{Rune: 'コ', IsConverted: true}
	case 'ｻ':
		return KanaConverterRune{Rune: 'サ', IsConverted: true}
	case 'ｼ':
		return KanaConverterRune{Rune: 'シ', IsConverted: true}
	case 'ｽ':
		return KanaConverterRune{Rune: 'ス', IsConverted: true}
	case 'ｾ':
		return KanaConverterRune{Rune: 'セ', IsConverted: true}
	case 'ｿ':
		return KanaConverterRune{Rune: 'ソ', IsConverted: true}
	case 'ﾀ':
		return KanaConverterRune{Rune: 'タ', IsConverted: true}
	case 'ﾁ':
		return KanaConverterRune{Rune: 'チ', IsConverted: true}
	case 'ﾂ':
		return KanaConverterRune{Rune: 'ツ', IsConverted: true}
	case 'ﾃ':
		return KanaConverterRune{Rune: 'テ', IsConverted: true}
	case 'ﾄ':
		return KanaConverterRune{Rune: 'ト', IsConverted: true}
	case 'ﾅ':
		return KanaConverterRune{Rune: 'ナ', IsConverted: true}
	case 'ﾆ':
		return KanaConverterRune{Rune: 'ニ', IsConverted: true}
	case 'ﾇ':
		return KanaConverterRune{Rune: 'ヌ', IsConverted: true}
	case 'ﾈ':
		return KanaConverterRune{Rune: 'ネ', IsConverted: true}
	case 'ﾉ':
		return KanaConverterRune{Rune: 'ノ', IsConverted: true}
	case 'ﾊ':
		return KanaConverterRune{Rune: 'ハ', IsConverted: true}
	case 'ﾋ':
		return KanaConverterRune{Rune: 'ヒ', IsConverted: true}
	case 'ﾌ':
		return KanaConverterRune{Rune: 'フ', IsConverted: true}
	case 'ﾍ':
		return KanaConverterRune{Rune: 'ヘ', IsConverted: true}
	case 'ﾎ':
		return KanaConverterRune{Rune: 'ホ', IsConverted: true}
	case 'ﾏ':
		return KanaConverterRune{Rune: 'マ', IsConverted: true}
	case 'ﾐ':
		return KanaConverterRune{Rune: 'ミ', IsConverted: true}
	case 'ﾑ':
		return KanaConverterRune{Rune: 'ム', IsConverted: true}
	case 'ﾒ':
		return KanaConverterRune{Rune: 'メ', IsConverted: true}
	case 'ﾓ':
		return KanaConverterRune{Rune: 'モ', IsConverted: true}
	case 'ﾔ':
		return KanaConverterRune{Rune: 'ヤ', IsConverted: true}
	case 'ﾕ':
		return KanaConverterRune{Rune: 'ユ', IsConverted: true}
	case 'ﾖ':
		return KanaConverterRune{Rune: 'ヨ', IsConverted: true}
	case 'ﾗ':
		return KanaConverterRune{Rune: 'ラ', IsConverted: true}
	case 'ﾘ':
		return KanaConverterRune{Rune: 'リ', IsConverted: true}
	case 'ﾙ':
		return KanaConverterRune{Rune: 'ル', IsConverted: true}
	case 'ﾚ':
		return KanaConverterRune{Rune: 'レ', IsConverted: true}
	case 'ﾛ':
		return KanaConverterRune{Rune: 'ロ', IsConverted: true}
	case 'ﾜ':
		return KanaConverterRune{Rune: 'ワ', IsConverted: true}
	case 'ﾝ':
		return KanaConverterRune{Rune: 'ン', IsConverted: true}
	case 'ﾞ':
		return KanaConverterRune{Rune: '゛', IsConverted: true}
	case 'ﾟ':
		return KanaConverterRune{Rune: '゜', IsConverted: true}
	default:
		return KanaConverterRune{Rune: in}
	}
}

func HankakuKatakanaToZenkakuKatakana(in <-chan KanaConverterRune, v bool) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		if v {
			var before *rune
			for r := range in {
				if r.IsConverted {
					out <- r
					continue
				}
				switch r.Rune {
				case 'ｦ', 'ｳ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾜ':
					if before != nil {
						out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
					}
					r := r
					before = &r.Rune
				case 'ﾞ':
					if before == nil {
						out <- KanaConverterRune{Rune: '゛', IsConverted: true}
					} else {
						switch *before {
						case 'ｳ':
							out <- KanaConverterRune{Rune: 'ヴ', IsConverted: true}
						case 'ｶ':
							out <- KanaConverterRune{Rune: 'ガ', IsConverted: true}
						case 'ｷ':
							out <- KanaConverterRune{Rune: 'ギ', IsConverted: true}
						case 'ｸ':
							out <- KanaConverterRune{Rune: 'グ', IsConverted: true}
						case 'ｹ':
							out <- KanaConverterRune{Rune: 'ゲ', IsConverted: true}
						case 'ｺ':
							out <- KanaConverterRune{Rune: 'ゴ', IsConverted: true}
						case 'ｻ':
							out <- KanaConverterRune{Rune: 'ザ', IsConverted: true}
						case 'ｼ':
							out <- KanaConverterRune{Rune: 'ジ', IsConverted: true}
						case 'ｽ':
							out <- KanaConverterRune{Rune: 'ズ', IsConverted: true}
						case 'ｾ':
							out <- KanaConverterRune{Rune: 'ゼ', IsConverted: true}
						case 'ｿ':
							out <- KanaConverterRune{Rune: 'ゾ', IsConverted: true}
						case 'ﾀ':
							out <- KanaConverterRune{Rune: 'ダ', IsConverted: true}
						case 'ﾁ':
							out <- KanaConverterRune{Rune: 'ヂ', IsConverted: true}
						case 'ﾂ':
							out <- KanaConverterRune{Rune: 'ヅ', IsConverted: true}
						case 'ﾃ':
							out <- KanaConverterRune{Rune: 'デ', IsConverted: true}
						case 'ﾄ':
							out <- KanaConverterRune{Rune: 'ド', IsConverted: true}
						case 'ﾊ':
							out <- KanaConverterRune{Rune: 'バ', IsConverted: true}
						case 'ﾋ':
							out <- KanaConverterRune{Rune: 'ビ', IsConverted: true}
						case 'ﾌ':
							out <- KanaConverterRune{Rune: 'ブ', IsConverted: true}
						case 'ﾍ':
							out <- KanaConverterRune{Rune: 'ベ', IsConverted: true}
						case 'ﾎ':
							out <- KanaConverterRune{Rune: 'ボ', IsConverted: true}
						default:
							out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
							out <- KanaConverterRune{Rune: '゛', IsConverted: true}
						}
						before = nil
					}
				case 'ﾟ':
					if before == nil {
						out <- KanaConverterRune{Rune: '゜', IsConverted: true}
					} else {
						switch *before {
						case 'ﾊ':
							out <- KanaConverterRune{Rune: 'パ', IsConverted: true}
						case 'ﾋ':
							out <- KanaConverterRune{Rune: 'ピ', IsConverted: true}
						case 'ﾌ':
							out <- KanaConverterRune{Rune: 'プ', IsConverted: true}
						case 'ﾍ':
							out <- KanaConverterRune{Rune: 'ペ', IsConverted: true}
						case 'ﾎ':
							out <- KanaConverterRune{Rune: 'ポ', IsConverted: true}
						default:
							out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
							out <- KanaConverterRune{Rune: '゜', IsConverted: true}
						}
						before = nil
					}
				default:
					if before != nil {
						out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
						before = nil
					}
					out <- hankakuKatakanaToZenkakuKatakanaSimple(r.Rune)
				}
			}
			if before != nil {
				out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
			}
		} else {
			for r := range in {
				if r.IsConverted {
					out <- r
					continue
				}
				out <- hankakuKatakanaToZenkakuKatakanaSimple(r.Rune)
			}
		}
	}()
	return out
}

func ZenkakuHiraganaToHankakuKatakana(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch r.Rune {
			case '、':
				out <- KanaConverterRune{Rune: '､', IsConverted: true}
			case '。':
				out <- KanaConverterRune{Rune: '｡', IsConverted: true}
			case '「':
				out <- KanaConverterRune{Rune: '｢', IsConverted: true}
			case '」':
				out <- KanaConverterRune{Rune: '｣', IsConverted: true}
			case '゛':
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case '゜':
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ぁ':
				out <- KanaConverterRune{Rune: 'ｧ', IsConverted: true}
			case 'あ':
				out <- KanaConverterRune{Rune: 'ｱ', IsConverted: true}
			case 'ぃ':
				out <- KanaConverterRune{Rune: 'ｨ', IsConverted: true}
			case 'い':
				out <- KanaConverterRune{Rune: 'ｲ', IsConverted: true}
			case 'ぅ':
				out <- KanaConverterRune{Rune: 'ｩ', IsConverted: true}
			case 'う':
				out <- KanaConverterRune{Rune: 'ｳ', IsConverted: true}
			case 'ぇ':
				out <- KanaConverterRune{Rune: 'ｪ', IsConverted: true}
			case 'え':
				out <- KanaConverterRune{Rune: 'ｴ', IsConverted: true}
			case 'ぉ':
				out <- KanaConverterRune{Rune: 'ｫ', IsConverted: true}
			case 'お':
				out <- KanaConverterRune{Rune: 'ｵ', IsConverted: true}
			case 'か':
				out <- KanaConverterRune{Rune: 'ｶ', IsConverted: true}
			case 'が':
				out <- KanaConverterRune{Rune: 'ｶ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'き':
				out <- KanaConverterRune{Rune: 'ｷ', IsConverted: true}
			case 'ぎ':
				out <- KanaConverterRune{Rune: 'ｷ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'く':
				out <- KanaConverterRune{Rune: 'ｸ', IsConverted: true}
			case 'ぐ':
				out <- KanaConverterRune{Rune: 'ｸ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'け':
				out <- KanaConverterRune{Rune: 'ｹ', IsConverted: true}
			case 'げ':
				out <- KanaConverterRune{Rune: 'ｹ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'こ':
				out <- KanaConverterRune{Rune: 'ｺ', IsConverted: true}
			case 'ご':
				out <- KanaConverterRune{Rune: 'ｺ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'さ':
				out <- KanaConverterRune{Rune: 'ｻ', IsConverted: true}
			case 'ざ':
				out <- KanaConverterRune{Rune: 'ｻ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'し':
				out <- KanaConverterRune{Rune: 'ｼ', IsConverted: true}
			case 'じ':
				out <- KanaConverterRune{Rune: 'ｼ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'す':
				out <- KanaConverterRune{Rune: 'ｽ', IsConverted: true}
			case 'ず':
				out <- KanaConverterRune{Rune: 'ｽ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'せ':
				out <- KanaConverterRune{Rune: 'ｾ', IsConverted: true}
			case 'ぜ':
				out <- KanaConverterRune{Rune: 'ｾ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'そ':
				out <- KanaConverterRune{Rune: 'ｿ', IsConverted: true}
			case 'ぞ':
				out <- KanaConverterRune{Rune: 'ｿ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'た':
				out <- KanaConverterRune{Rune: 'ﾀ', IsConverted: true}
			case 'だ':
				out <- KanaConverterRune{Rune: 'ﾀ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ち':
				out <- KanaConverterRune{Rune: 'ﾁ', IsConverted: true}
			case 'ぢ':
				out <- KanaConverterRune{Rune: 'ﾁ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'っ':
				out <- KanaConverterRune{Rune: 'ｯ', IsConverted: true}
			case 'つ':
				out <- KanaConverterRune{Rune: 'ﾂ', IsConverted: true}
			case 'づ':
				out <- KanaConverterRune{Rune: 'ﾂ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'て':
				out <- KanaConverterRune{Rune: 'ﾃ', IsConverted: true}
			case 'で':
				out <- KanaConverterRune{Rune: 'ﾃ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'と':
				out <- KanaConverterRune{Rune: 'ﾄ', IsConverted: true}
			case 'ど':
				out <- KanaConverterRune{Rune: 'ﾄ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'な':
				out <- KanaConverterRune{Rune: 'ﾅ', IsConverted: true}
			case 'に':
				out <- KanaConverterRune{Rune: 'ﾆ', IsConverted: true}
			case 'ぬ':
				out <- KanaConverterRune{Rune: 'ﾇ', IsConverted: true}
			case 'ね':
				out <- KanaConverterRune{Rune: 'ﾈ', IsConverted: true}
			case 'の':
				out <- KanaConverterRune{Rune: 'ﾉ', IsConverted: true}
			case 'は':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
			case 'ば':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ぱ':
				out <- KanaConverterRune{Rune: 'ﾊ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ひ':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
			case 'び':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ぴ':
				out <- KanaConverterRune{Rune: 'ﾋ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ふ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
			case 'ぶ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ぷ':
				out <- KanaConverterRune{Rune: 'ﾌ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'へ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
			case 'べ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ぺ':
				out <- KanaConverterRune{Rune: 'ﾍ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ほ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
			case 'ぼ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾞ', IsConverted: true}
			case 'ぽ':
				out <- KanaConverterRune{Rune: 'ﾎ', IsConverted: true}
				out <- KanaConverterRune{Rune: 'ﾟ', IsConverted: true}
			case 'ま':
				out <- KanaConverterRune{Rune: 'ﾏ', IsConverted: true}
			case 'み':
				out <- KanaConverterRune{Rune: 'ﾐ', IsConverted: true}
			case 'む':
				out <- KanaConverterRune{Rune: 'ﾑ', IsConverted: true}
			case 'め':
				out <- KanaConverterRune{Rune: 'ﾒ', IsConverted: true}
			case 'も':
				out <- KanaConverterRune{Rune: 'ﾓ', IsConverted: true}
			case 'ゃ':
				out <- KanaConverterRune{Rune: 'ｬ', IsConverted: true}
			case 'や':
				out <- KanaConverterRune{Rune: 'ﾔ', IsConverted: true}
			case 'ゅ':
				out <- KanaConverterRune{Rune: 'ｭ', IsConverted: true}
			case 'ゆ':
				out <- KanaConverterRune{Rune: 'ﾕ', IsConverted: true}
			case 'ょ':
				out <- KanaConverterRune{Rune: 'ｮ', IsConverted: true}
			case 'よ':
				out <- KanaConverterRune{Rune: 'ﾖ', IsConverted: true}
			case 'ら':
				out <- KanaConverterRune{Rune: 'ﾗ', IsConverted: true}
			case 'り':
				out <- KanaConverterRune{Rune: 'ﾘ', IsConverted: true}
			case 'る':
				out <- KanaConverterRune{Rune: 'ﾙ', IsConverted: true}
			case 'れ':
				out <- KanaConverterRune{Rune: 'ﾚ', IsConverted: true}
			case 'ろ':
				out <- KanaConverterRune{Rune: 'ﾛ', IsConverted: true}
			case 'ゎ':
				out <- KanaConverterRune{Rune: 'ﾜ', IsConverted: true}
			case 'わ':
				out <- KanaConverterRune{Rune: 'ﾜ', IsConverted: true}
			case 'ゐ':
				out <- KanaConverterRune{Rune: 'ｲ', IsConverted: true}
			case 'ゑ':
				out <- KanaConverterRune{Rune: 'ｴ', IsConverted: true}
			case 'を':
				out <- KanaConverterRune{Rune: 'ｦ', IsConverted: true}
			case 'ん':
				out <- KanaConverterRune{Rune: 'ﾝ', IsConverted: true}
			case '・':
				out <- KanaConverterRune{Rune: '･', IsConverted: true}
			case 'ー':
				out <- KanaConverterRune{Rune: 'ｰ', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func hankakuKatakanaToZenkakuHiraganaSimple(in rune) KanaConverterRune {
	switch in {
	case '｡':
		return KanaConverterRune{Rune: '。', IsConverted: true}
	case '｢':
		return KanaConverterRune{Rune: '「', IsConverted: true}
	case '｣':
		return KanaConverterRune{Rune: '」', IsConverted: true}
	case '､':
		return KanaConverterRune{Rune: '、', IsConverted: true}
	case '･':
		return KanaConverterRune{Rune: '・', IsConverted: true}
	case 'ｦ':
		return KanaConverterRune{Rune: 'を', IsConverted: true}
	case 'ｧ':
		return KanaConverterRune{Rune: 'ぁ', IsConverted: true}
	case 'ｨ':
		return KanaConverterRune{Rune: 'ぃ', IsConverted: true}
	case 'ｩ':
		return KanaConverterRune{Rune: 'ぅ', IsConverted: true}
	case 'ｪ':
		return KanaConverterRune{Rune: 'ぇ', IsConverted: true}
	case 'ｫ':
		return KanaConverterRune{Rune: 'ぉ', IsConverted: true}
	case 'ｬ':
		return KanaConverterRune{Rune: 'ゃ', IsConverted: true}
	case 'ｭ':
		return KanaConverterRune{Rune: 'ゅ', IsConverted: true}
	case 'ｮ':
		return KanaConverterRune{Rune: 'ょ', IsConverted: true}
	case 'ｯ':
		return KanaConverterRune{Rune: 'っ', IsConverted: true}
	case 'ｰ':
		return KanaConverterRune{Rune: 'ー', IsConverted: true}
	case 'ｱ':
		return KanaConverterRune{Rune: 'あ', IsConverted: true}
	case 'ｲ':
		return KanaConverterRune{Rune: 'い', IsConverted: true}
	case 'ｳ':
		return KanaConverterRune{Rune: 'う', IsConverted: true}
	case 'ｴ':
		return KanaConverterRune{Rune: 'え', IsConverted: true}
	case 'ｵ':
		return KanaConverterRune{Rune: 'お', IsConverted: true}
	case 'ｶ':
		return KanaConverterRune{Rune: 'か', IsConverted: true}
	case 'ｷ':
		return KanaConverterRune{Rune: 'き', IsConverted: true}
	case 'ｸ':
		return KanaConverterRune{Rune: 'く', IsConverted: true}
	case 'ｹ':
		return KanaConverterRune{Rune: 'け', IsConverted: true}
	case 'ｺ':
		return KanaConverterRune{Rune: 'こ', IsConverted: true}
	case 'ｻ':
		return KanaConverterRune{Rune: 'さ', IsConverted: true}
	case 'ｼ':
		return KanaConverterRune{Rune: 'し', IsConverted: true}
	case 'ｽ':
		return KanaConverterRune{Rune: 'す', IsConverted: true}
	case 'ｾ':
		return KanaConverterRune{Rune: 'せ', IsConverted: true}
	case 'ｿ':
		return KanaConverterRune{Rune: 'そ', IsConverted: true}
	case 'ﾀ':
		return KanaConverterRune{Rune: 'た', IsConverted: true}
	case 'ﾁ':
		return KanaConverterRune{Rune: 'ち', IsConverted: true}
	case 'ﾂ':
		return KanaConverterRune{Rune: 'つ', IsConverted: true}
	case 'ﾃ':
		return KanaConverterRune{Rune: 'て', IsConverted: true}
	case 'ﾄ':
		return KanaConverterRune{Rune: 'と', IsConverted: true}
	case 'ﾅ':
		return KanaConverterRune{Rune: 'な', IsConverted: true}
	case 'ﾆ':
		return KanaConverterRune{Rune: 'に', IsConverted: true}
	case 'ﾇ':
		return KanaConverterRune{Rune: 'ぬ', IsConverted: true}
	case 'ﾈ':
		return KanaConverterRune{Rune: 'ね', IsConverted: true}
	case 'ﾉ':
		return KanaConverterRune{Rune: 'の', IsConverted: true}
	case 'ﾊ':
		return KanaConverterRune{Rune: 'は', IsConverted: true}
	case 'ﾋ':
		return KanaConverterRune{Rune: 'ひ', IsConverted: true}
	case 'ﾌ':
		return KanaConverterRune{Rune: 'ふ', IsConverted: true}
	case 'ﾍ':
		return KanaConverterRune{Rune: 'へ', IsConverted: true}
	case 'ﾎ':
		return KanaConverterRune{Rune: 'ほ', IsConverted: true}
	case 'ﾏ':
		return KanaConverterRune{Rune: 'ま', IsConverted: true}
	case 'ﾐ':
		return KanaConverterRune{Rune: 'み', IsConverted: true}
	case 'ﾑ':
		return KanaConverterRune{Rune: 'む', IsConverted: true}
	case 'ﾒ':
		return KanaConverterRune{Rune: 'め', IsConverted: true}
	case 'ﾓ':
		return KanaConverterRune{Rune: 'も', IsConverted: true}
	case 'ﾔ':
		return KanaConverterRune{Rune: 'や', IsConverted: true}
	case 'ﾕ':
		return KanaConverterRune{Rune: 'ゆ', IsConverted: true}
	case 'ﾖ':
		return KanaConverterRune{Rune: 'よ', IsConverted: true}
	case 'ﾗ':
		return KanaConverterRune{Rune: 'ら', IsConverted: true}
	case 'ﾘ':
		return KanaConverterRune{Rune: 'り', IsConverted: true}
	case 'ﾙ':
		return KanaConverterRune{Rune: 'る', IsConverted: true}
	case 'ﾚ':
		return KanaConverterRune{Rune: 'れ', IsConverted: true}
	case 'ﾛ':
		return KanaConverterRune{Rune: 'ろ', IsConverted: true}
	case 'ﾜ':
		return KanaConverterRune{Rune: 'わ', IsConverted: true}
	case 'ﾝ':
		return KanaConverterRune{Rune: 'ん', IsConverted: true}
	case 'ﾞ':
		return KanaConverterRune{Rune: '゛', IsConverted: true}
	case 'ﾟ':
		return KanaConverterRune{Rune: '゜', IsConverted: true}
	default:
		return KanaConverterRune{Rune: in}
	}
}

func HankakuKatakanaToZenkakuHiragana(in <-chan KanaConverterRune, v bool) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		if v {
			var before *rune
			for r := range in {
				if r.IsConverted {
					out <- r
					continue
				}
				switch r.Rune {
				case 'ｳ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾜ':
					if before != nil {
						out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
					}
					r := r
					before = &r.Rune
				case 'ﾞ':
					if before == nil {
						out <- KanaConverterRune{Rune: '゛', IsConverted: true}
					} else {
						switch *before {
						case 'ｶ':
							out <- KanaConverterRune{Rune: 'が', IsConverted: true}
						case 'ｷ':
							out <- KanaConverterRune{Rune: 'ぎ', IsConverted: true}
						case 'ｸ':
							out <- KanaConverterRune{Rune: 'ぐ', IsConverted: true}
						case 'ｹ':
							out <- KanaConverterRune{Rune: 'げ', IsConverted: true}
						case 'ｺ':
							out <- KanaConverterRune{Rune: 'ご', IsConverted: true}
						case 'ｻ':
							out <- KanaConverterRune{Rune: 'ざ', IsConverted: true}
						case 'ｼ':
							out <- KanaConverterRune{Rune: 'じ', IsConverted: true}
						case 'ｽ':
							out <- KanaConverterRune{Rune: 'ず', IsConverted: true}
						case 'ｾ':
							out <- KanaConverterRune{Rune: 'ぜ', IsConverted: true}
						case 'ｿ':
							out <- KanaConverterRune{Rune: 'ぞ', IsConverted: true}
						case 'ﾀ':
							out <- KanaConverterRune{Rune: 'だ', IsConverted: true}
						case 'ﾁ':
							out <- KanaConverterRune{Rune: 'ぢ', IsConverted: true}
						case 'ﾂ':
							out <- KanaConverterRune{Rune: 'づ', IsConverted: true}
						case 'ﾃ':
							out <- KanaConverterRune{Rune: 'で', IsConverted: true}
						case 'ﾄ':
							out <- KanaConverterRune{Rune: 'ど', IsConverted: true}
						case 'ﾊ':
							out <- KanaConverterRune{Rune: 'ば', IsConverted: true}
						case 'ﾋ':
							out <- KanaConverterRune{Rune: 'び', IsConverted: true}
						case 'ﾌ':
							out <- KanaConverterRune{Rune: 'ぶ', IsConverted: true}
						case 'ﾍ':
							out <- KanaConverterRune{Rune: 'べ', IsConverted: true}
						case 'ﾎ':
							out <- KanaConverterRune{Rune: 'ぼ', IsConverted: true}
						default:
							out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
							out <- KanaConverterRune{Rune: '゛', IsConverted: true}
						}
						before = nil
					}
				case 'ﾟ':
					if before == nil {
						out <- KanaConverterRune{Rune: '゜', IsConverted: true}
					} else {
						switch *before {
						case 'ﾊ':
							out <- KanaConverterRune{Rune: 'ぱ', IsConverted: true}
						case 'ﾋ':
							out <- KanaConverterRune{Rune: 'ぴ', IsConverted: true}
						case 'ﾌ':
							out <- KanaConverterRune{Rune: 'ぷ', IsConverted: true}
						case 'ﾍ':
							out <- KanaConverterRune{Rune: 'ぺ', IsConverted: true}
						case 'ﾎ':
							out <- KanaConverterRune{Rune: 'ぽ', IsConverted: true}
						default:
							out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
							out <- KanaConverterRune{Rune: '゜', IsConverted: true}
						}
						before = nil
					}
				default:
					if before != nil {
						out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
						before = nil
					}
					out <- hankakuKatakanaToZenkakuHiraganaSimple(r.Rune)
				}
			}
			if before != nil {
				out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
			}
		} else {
			for r := range in {
				if r.IsConverted {
					out <- r
					continue
				}
				out <- hankakuKatakanaToZenkakuHiraganaSimple(r.Rune)
			}
		}
	}()
	return out
}

func ZenkakuKatakanaToZenkakuHiragana(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune >= 'ァ' && r.Rune <= 'ン', r.Rune == 'ヽ', r.Rune == 'ヾ':
				out <- KanaConverterRune{Rune: 'ぁ' + r.Rune - 'ァ', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuHiraganaToZenkakuKatakana(in <-chan KanaConverterRune) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for r := range in {
			if r.IsConverted {
				out <- r
				continue
			}
			switch {
			case r.Rune >= 'ぁ' && r.Rune <= 'ん', r.Rune == 'ゝ', r.Rune == 'ゞ':
				out <- KanaConverterRune{Rune: 'ァ' + r.Rune - 'ぁ', IsConverted: true}
			default:
				out <- r
			}
		}
	}()
	return out
}

func NewKanaConverters(mode string) ([]func(<-chan KanaConverterRune) <-chan KanaConverterRune, error) {
	options, err := NewKanaConverterOptions(mode)
	if err != nil {
		return nil, err
	}
	var converters []func(<-chan KanaConverterRune) <-chan KanaConverterRune
	if options.optr {
		converters = append(converters, ZenkakuEnglishToHankakuEnglish)
	}
	if options.optR {
		converters = append(converters, HankakuEnglishToZenkakuEnglish)
	}
	if options.optn {
		converters = append(converters, ZenkakuNumberToHankakuNumber)
	}
	if options.optN {
		converters = append(converters, HankakuNumberToZenkakuNumber)
	}
	if options.opta {
		converters = append(converters, ZenkakuEnglishNumberToHankakuEnglishNumber)
	}
	if options.optA {
		converters = append(converters, HankakuEnglishNumberToZenkakuEnglishNumber)
	}
	if options.opts {
		converters = append(converters, ZenkakuSpaceToHankakuSpace)
	}
	if options.optS {
		converters = append(converters, HankakuSpaceToZenkakuSpace)
	}

	// kc, kC, KH, hc and hC are not combined
	if options.optk {
		if options.opth {
			converters = append(converters, ZenkakuHiraganaToHankakuKatakana)
		}
		if options.optH {
			converters = append(converters, func(in <-chan KanaConverterRune) <-chan KanaConverterRune {
				return HankakuKatakanaToZenkakuHiragana(in, options.optV)
			})
		}
		converters = append(converters, ZenkakuKatakanaToHankakuKatakana)
		return converters, nil
	}
	if options.optK {
		if options.optc {
			converters = append(converters, ZenkakuKatakanaToZenkakuHiragana)
		}
		converters = append(converters, func(in <-chan KanaConverterRune) <-chan KanaConverterRune {
			return HankakuKatakanaToZenkakuKatakana(in, options.optV)
		})
		if options.opth {
			converters = append(converters, ZenkakuHiraganaToHankakuKatakana)
		}
		if options.optC {
			converters = append(converters, ZenkakuHiraganaToZenkakuKatakana)
		}
		return converters, nil
	}
	if options.opth {
		converters = append(converters, ZenkakuHiraganaToHankakuKatakana)
		return converters, nil
	}
	if options.optH {
		if options.optC {
			converters = append(converters, ZenkakuHiraganaToZenkakuKatakana)
		}
		converters = append(converters, func(in <-chan KanaConverterRune) <-chan KanaConverterRune {
			return HankakuKatakanaToZenkakuHiragana(in, options.optV)
		})
		if options.optc {
			converters = append(converters, ZenkakuKatakanaToZenkakuHiragana)
		}
		return converters, nil
	}
	if options.optc {
		converters = append(converters, ZenkakuKatakanaToZenkakuHiragana)
	}
	if options.optC {
		converters = append(converters, ZenkakuHiraganaToZenkakuKatakana)
	}
	return converters, nil
}
