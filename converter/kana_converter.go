package converter

func HankakuNumberToZenkakuNumber(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			if i >= '0' && i <= '9' {
				out <- i + 0xFEE0
			} else {
				out <- i
			}
		}
	}()
	return out
}

func ZenkakuNumberToHankakuNumber(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			if i >= '０' && i <= '９' {
				out <- i - 0xFEE0
			} else {
				out <- i
			}
		}
	}()
	return out
}

func HankakuEnglishToZenkakuEnglish(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			switch {
			case i >= 'a' && i <= 'z':
				out <- 'ａ' + i - 'a'
			case i >= 'A' && i <= 'Z':
				out <- 'Ａ' + i - 'A'
			default:
				out <- i
			}
		}
	}()
	return out
}

func ZenkakuEnglishToHankakuEnglish(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			switch {
			case i >= 'ａ' && i <= 'ｚ':
				out <- 'a' + i - 'ａ'
			case i >= 'Ａ' && i <= 'Ｚ':
				out <- 'A' + i - 'Ａ'
			default:
				out <- i
			}
		}
	}()
	return out
}

func ZenkakuSpaceToHankakuSpace(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for r := range in {
			if r == '　' {
				out <- ' '
			} else {
				out <- r
			}
		}
	}()
	return out
}

func HankakuSpaceToZenkakuSpace(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for r := range in {
			if r == ' ' {
				out <- '　'
			} else {
				out <- r
			}
		}
	}()
	return out
}

func ZenkakuKatakanaToHankakuKatakana(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for r := range in {
			switch r {
			case '、':
				out <- '､'
			case '。':
				out <- '｡'
			case '「':
				out <- '｢'
			case '」':
				out <- '｣'
			case '゛':
				out <- 'ﾞ'
			case '゜':
				out <- 'ﾟ'
			case 'ァ':
				out <- 'ｧ'
			case 'ア':
				out <- 'ｱ'
			case 'ィ':
				out <- 'ｨ'
			case 'イ':
				out <- 'ｲ'
			case 'ゥ':
				out <- 'ｩ'
			case 'ウ':
				out <- 'ｳ'
			case 'ェ':
				out <- 'ｪ'
			case 'エ':
				out <- 'ｴ'
			case 'ォ':
				out <- 'ｫ'
			case 'オ':
				out <- 'ｵ'
			case 'カ':
				out <- 'ｶ'
			case 'ガ':
				out <- 'ｶ'
				out <- 'ﾞ'
			case 'キ':
				out <- 'ｷ'
			case 'ギ':
				out <- 'ｷ'
				out <- 'ﾞ'
			case 'ク':
				out <- 'ｸ'
			case 'グ':
				out <- 'ｸ'
				out <- 'ﾞ'
			case 'ケ':
				out <- 'ｹ'
			case 'ゲ':
				out <- 'ｹ'
				out <- 'ﾞ'
			case 'コ':
				out <- 'ｺ'
			case 'ゴ':
				out <- 'ｺ'
				out <- 'ﾞ'
			case 'サ':
				out <- 'ｻ'
			case 'ザ':
				out <- 'ｻ'
				out <- 'ﾞ'
			case 'シ':
				out <- 'ｼ'
			case 'ジ':
				out <- 'ｼ'
				out <- 'ﾞ'
			case 'ス':
				out <- 'ｽ'
			case 'ズ':
				out <- 'ｽ'
				out <- 'ﾞ'
			case 'セ':
				out <- 'ｾ'
			case 'ゼ':
				out <- 'ｾ'
				out <- 'ﾞ'
			case 'ソ':
				out <- 'ｿ'
			case 'ゾ':
				out <- 'ｿ'
				out <- 'ﾞ'
			case 'タ':
				out <- 'ﾀ'
			case 'ダ':
				out <- 'ﾀ'
				out <- 'ﾞ'
			case 'チ':
				out <- 'ﾁ'
			case 'ヂ':
				out <- 'ﾁ'
				out <- 'ﾞ'
			case 'ッ':
				out <- 'ｯ'
			case 'ツ':
				out <- 'ﾂ'
			case 'ヅ':
				out <- 'ﾂ'
				out <- 'ﾞ'
			case 'テ':
				out <- 'ﾃ'
			case 'デ':
				out <- 'ﾃ'
				out <- 'ﾞ'
			case 'ト':
				out <- 'ﾄ'
			case 'ド':
				out <- 'ﾄ'
				out <- 'ﾞ'
			case 'ナ':
				out <- 'ﾅ'
			case 'ニ':
				out <- 'ﾆ'
			case 'ヌ':
				out <- 'ﾇ'
			case 'ネ':
				out <- 'ﾈ'
			case 'ノ':
				out <- 'ﾉ'
			case 'ハ':
				out <- 'ﾊ'
			case 'バ':
				out <- 'ﾊ'
				out <- 'ﾞ'
			case 'パ':
				out <- 'ﾊ'
				out <- 'ﾟ'
			case 'ヒ':
				out <- 'ﾋ'
			case 'ビ':
				out <- 'ﾋ'
				out <- 'ﾞ'
			case 'ピ':
				out <- 'ﾋ'
				out <- 'ﾟ'
			case 'フ':
				out <- 'ﾌ'
			case 'ブ':
				out <- 'ﾌ'
				out <- 'ﾞ'
			case 'プ':
				out <- 'ﾌ'
				out <- 'ﾟ'
			case 'ヘ':
				out <- 'ﾍ'
			case 'ベ':
				out <- 'ﾍ'
				out <- 'ﾞ'
			case 'ペ':
				out <- 'ﾍ'
				out <- 'ﾟ'
			case 'ホ':
				out <- 'ﾎ'
			case 'ボ':
				out <- 'ﾎ'
				out <- 'ﾞ'
			case 'ポ':
				out <- 'ﾎ'
				out <- 'ﾟ'
			case 'マ':
				out <- 'ﾏ'
			case 'ミ':
				out <- 'ﾐ'
			case 'ム':
				out <- 'ﾑ'
			case 'メ':
				out <- 'ﾒ'
			case 'モ':
				out <- 'ﾓ'
			case 'ャ':
				out <- 'ｬ'
			case 'ヤ':
				out <- 'ﾔ'
			case 'ュ':
				out <- 'ｭ'
			case 'ユ':
				out <- 'ﾕ'
			case 'ョ':
				out <- 'ｮ'
			case 'ヨ':
				out <- 'ﾖ'
			case 'ラ':
				out <- 'ﾗ'
			case 'リ':
				out <- 'ﾘ'
			case 'ル':
				out <- 'ﾙ'
			case 'レ':
				out <- 'ﾚ'
			case 'ロ':
				out <- 'ﾛ'
			case 'ヮ':
				out <- 'ﾜ'
			case 'ワ':
				out <- 'ﾜ'
			case 'ヰ':
				out <- 'ｲ'
			case 'ヱ':
				out <- 'ｴ'
			case 'ヲ':
				out <- 'ｦ'
			case 'ン':
				out <- 'ﾝ'
			case 'ヴ':
				out <- 'ｳ'
				out <- 'ﾞ'
			case '・':
				out <- '･'
			case 'ー':
				out <- 'ｰ'
			default:
				out <- r
			}
		}
	}()
	return out
}

func hankakuKatakanaToZenkakuKatakanaSimple(in rune) rune {
	switch in {
	case '｡':
		return '。'
	case '｢':
		return '「'
	case '｣':
		return '」'
	case '､':
		return '、'
	case '･':
		return '・'
	case 'ｦ':
		return 'ヲ'
	case 'ｧ':
		return 'ァ'
	case 'ｨ':
		return 'ィ'
	case 'ｩ':
		return 'ゥ'
	case 'ｪ':
		return 'ェ'
	case 'ｫ':
		return 'ォ'
	case 'ｬ':
		return 'ャ'
	case 'ｭ':
		return 'ュ'
	case 'ｮ':
		return 'ョ'
	case 'ｯ':
		return 'ッ'
	case 'ｰ':
		return 'ー'
	case 'ｱ':
		return 'ア'
	case 'ｲ':
		return 'イ'
	case 'ｳ':
		return 'ウ'
	case 'ｴ':
		return 'エ'
	case 'ｵ':
		return 'オ'
	case 'ｶ':
		return 'カ'
	case 'ｷ':
		return 'キ'
	case 'ｸ':
		return 'ク'
	case 'ｹ':
		return 'ケ'
	case 'ｺ':
		return 'コ'
	case 'ｻ':
		return 'サ'
	case 'ｼ':
		return 'シ'
	case 'ｽ':
		return 'ス'
	case 'ｾ':
		return 'セ'
	case 'ｿ':
		return 'ソ'
	case 'ﾀ':
		return 'タ'
	case 'ﾁ':
		return 'チ'
	case 'ﾂ':
		return 'ツ'
	case 'ﾃ':
		return 'テ'
	case 'ﾄ':
		return 'ト'
	case 'ﾅ':
		return 'ナ'
	case 'ﾆ':
		return 'ニ'
	case 'ﾇ':
		return 'ヌ'
	case 'ﾈ':
		return 'ネ'
	case 'ﾉ':
		return 'ノ'
	case 'ﾊ':
		return 'ハ'
	case 'ﾋ':
		return 'ヒ'
	case 'ﾌ':
		return 'フ'
	case 'ﾍ':
		return 'ヘ'
	case 'ﾎ':
		return 'ホ'
	case 'ﾏ':
		return 'マ'
	case 'ﾐ':
		return 'ミ'
	case 'ﾑ':
		return 'ム'
	case 'ﾒ':
		return 'メ'
	case 'ﾓ':
		return 'モ'
	case 'ﾔ':
		return 'ヤ'
	case 'ﾕ':
		return 'ユ'
	case 'ﾖ':
		return 'ヨ'
	case 'ﾗ':
		return 'ラ'
	case 'ﾘ':
		return 'リ'
	case 'ﾙ':
		return 'ル'
	case 'ﾚ':
		return 'レ'
	case 'ﾛ':
		return 'ロ'
	case 'ﾜ':
		return 'ワ'
	case 'ﾝ':
		return 'ン'
	case 'ﾞ':
		return '゛'
	case 'ﾟ':
		return '゜'
	default:
		return in
	}
}

func HankakuKatakanaToZenkakuKatakana(in <-chan rune, v bool) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		if v {
			var before *rune
			for r := range in {
				switch r {
				case 'ｦ', 'ｳ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾜ':
					if before != nil {
						out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
					}
					r := r
					before = &r
				case 'ﾞ':
					if before == nil {
						out <- '゛'
					} else {
						switch *before {
						case 'ｳ':
							out <- 'ヴ'
						case 'ｶ':
							out <- 'ガ'
						case 'ｷ':
							out <- 'ギ'
						case 'ｸ':
							out <- 'グ'
						case 'ｹ':
							out <- 'ゲ'
						case 'ｺ':
							out <- 'ゴ'
						case 'ｻ':
							out <- 'ザ'
						case 'ｼ':
							out <- 'ジ'
						case 'ｽ':
							out <- 'ズ'
						case 'ｾ':
							out <- 'ゼ'
						case 'ｿ':
							out <- 'ゾ'
						case 'ﾀ':
							out <- 'ダ'
						case 'ﾁ':
							out <- 'ヂ'
						case 'ﾂ':
							out <- 'ヅ'
						case 'ﾃ':
							out <- 'デ'
						case 'ﾄ':
							out <- 'ド'
						case 'ﾊ':
							out <- 'バ'
						case 'ﾋ':
							out <- 'ビ'
						case 'ﾌ':
							out <- 'ブ'
						case 'ﾍ':
							out <- 'ベ'
						case 'ﾎ':
							out <- 'ボ'
						default:
							out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
							out <- '゛'
						}
						before = nil
					}
				case 'ﾟ':
					if before == nil {
						out <- '゜'
					} else {
						switch *before {
						case 'ﾊ':
							out <- 'パ'
						case 'ﾋ':
							out <- 'ピ'
						case 'ﾌ':
							out <- 'プ'
						case 'ﾍ':
							out <- 'ペ'
						case 'ﾎ':
							out <- 'ポ'
						default:
							out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
							out <- '゜'
						}
						before = nil
					}
				default:
					if before != nil {
						out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
						before = nil
					}
					out <- hankakuKatakanaToZenkakuKatakanaSimple(r)
				}
			}
			if before != nil {
				out <- hankakuKatakanaToZenkakuKatakanaSimple(*before)
			}
		} else {
			for i := range in {
				out <- hankakuKatakanaToZenkakuKatakanaSimple(i)
			}
		}
	}()
	return out
}

func ZenkakuHiraganaToHankakuKatakana(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			switch i {
			case '、':
				out <- '､'
			case '。':
				out <- '｡'
			case '「':
				out <- '｢'
			case '」':
				out <- '｣'
			case '゛':
				out <- 'ﾞ'
			case '゜':
				out <- 'ﾟ'
			case 'ぁ':
				out <- 'ｧ'
			case 'あ':
				out <- 'ｱ'
			case 'ぃ':
				out <- 'ｨ'
			case 'い':
				out <- 'ｲ'
			case 'ぅ':
				out <- 'ｩ'
			case 'う':
				out <- 'ｳ'
			case 'ぇ':
				out <- 'ｪ'
			case 'え':
				out <- 'ｴ'
			case 'ぉ':
				out <- 'ｫ'
			case 'お':
				out <- 'ｵ'
			case 'か':
				out <- 'ｶ'
			case 'が':
				out <- 'ｶ'
				out <- 'ﾞ'
			case 'き':
				out <- 'ｷ'
			case 'ぎ':
				out <- 'ｷ'
				out <- 'ﾞ'
			case 'く':
				out <- 'ｸ'
			case 'ぐ':
				out <- 'ｸ'
				out <- 'ﾞ'
			case 'け':
				out <- 'ｹ'
			case 'げ':
				out <- 'ｹ'
				out <- 'ﾞ'
			case 'こ':
				out <- 'ｺ'
			case 'ご':
				out <- 'ｺ'
				out <- 'ﾞ'
			case 'さ':
				out <- 'ｻ'
			case 'ざ':
				out <- 'ｻ'
				out <- 'ﾞ'
			case 'し':
				out <- 'ｼ'
			case 'じ':
				out <- 'ｼ'
				out <- 'ﾞ'
			case 'す':
				out <- 'ｽ'
			case 'ず':
				out <- 'ｽ'
				out <- 'ﾞ'
			case 'せ':
				out <- 'ｾ'
			case 'ぜ':
				out <- 'ｾ'
				out <- 'ﾞ'
			case 'そ':
				out <- 'ｿ'
			case 'ぞ':
				out <- 'ｿ'
				out <- 'ﾞ'
			case 'た':
				out <- 'ﾀ'
			case 'だ':
				out <- 'ﾀ'
				out <- 'ﾞ'
			case 'ち':
				out <- 'ﾁ'
			case 'ぢ':
				out <- 'ﾁ'
				out <- 'ﾞ'
			case 'っ':
				out <- 'ｯ'
			case 'つ':
				out <- 'ﾂ'
			case 'づ':
				out <- 'ﾂ'
				out <- 'ﾞ'
			case 'て':
				out <- 'ﾃ'
			case 'で':
				out <- 'ﾃ'
				out <- 'ﾞ'
			case 'と':
				out <- 'ﾄ'
			case 'ど':
				out <- 'ﾄ'
				out <- 'ﾞ'
			case 'な':
				out <- 'ﾅ'
			case 'に':
				out <- 'ﾆ'
			case 'ぬ':
				out <- 'ﾇ'
			case 'ね':
				out <- 'ﾈ'
			case 'の':
				out <- 'ﾉ'
			case 'は':
				out <- 'ﾊ'
			case 'ば':
				out <- 'ﾊ'
				out <- 'ﾞ'
			case 'ぱ':
				out <- 'ﾊ'
				out <- 'ﾟ'
			case 'ひ':
				out <- 'ﾋ'
			case 'び':
				out <- 'ﾋ'
				out <- 'ﾞ'
			case 'ぴ':
				out <- 'ﾋ'
				out <- 'ﾟ'
			case 'ふ':
				out <- 'ﾌ'
			case 'ぶ':
				out <- 'ﾌ'
				out <- 'ﾞ'
			case 'ぷ':
				out <- 'ﾌ'
				out <- 'ﾟ'
			case 'へ':
				out <- 'ﾍ'
			case 'べ':
				out <- 'ﾍ'
				out <- 'ﾞ'
			case 'ぺ':
				out <- 'ﾍ'
				out <- 'ﾟ'
			case 'ほ':
				out <- 'ﾎ'
			case 'ぼ':
				out <- 'ﾎ'
				out <- 'ﾞ'
			case 'ぽ':
				out <- 'ﾎ'
				out <- 'ﾟ'
			case 'ま':
				out <- 'ﾏ'
			case 'み':
				out <- 'ﾐ'
			case 'む':
				out <- 'ﾑ'
			case 'め':
				out <- 'ﾒ'
			case 'も':
				out <- 'ﾓ'
			case 'ゃ':
				out <- 'ｬ'
			case 'や':
				out <- 'ﾔ'
			case 'ゅ':
				out <- 'ｭ'
			case 'ゆ':
				out <- 'ﾕ'
			case 'ょ':
				out <- 'ｮ'
			case 'よ':
				out <- 'ﾖ'
			case 'ら':
				out <- 'ﾗ'
			case 'り':
				out <- 'ﾘ'
			case 'る':
				out <- 'ﾙ'
			case 'れ':
				out <- 'ﾚ'
			case 'ろ':
				out <- 'ﾛ'
			case 'ゎ':
				out <- 'ﾜ'
			case 'わ':
				out <- 'ﾜ'
			case 'ゐ':
				out <- 'ｲ'
			case 'ゑ':
				out <- 'ｴ'
			case 'を':
				out <- 'ｦ'
			case 'ん':
				out <- 'ﾝ'
			case '・':
				out <- '･'
			case 'ー':
				out <- 'ｰ'
			default:
				out <- i
			}
		}
	}()
	return out
}

func hankakuKatakanaToZenkakuHiraganaSimple(in rune) rune {
	switch in {
	case '｡':
		return '。'
	case '｢':
		return '「'
	case '｣':
		return '」'
	case '､':
		return '、'
	case '･':
		return '・'
	case 'ｦ':
		return 'を'
	case 'ｧ':
		return 'ぁ'
	case 'ｨ':
		return 'ぃ'
	case 'ｩ':
		return 'ぅ'
	case 'ｪ':
		return 'ぇ'
	case 'ｫ':
		return 'ぉ'
	case 'ｬ':
		return 'ゃ'
	case 'ｭ':
		return 'ゅ'
	case 'ｮ':
		return 'ょ'
	case 'ｯ':
		return 'っ'
	case 'ｰ':
		return 'ー'
	case 'ｱ':
		return 'あ'
	case 'ｲ':
		return 'い'
	case 'ｳ':
		return 'う'
	case 'ｴ':
		return 'え'
	case 'ｵ':
		return 'お'
	case 'ｶ':
		return 'か'
	case 'ｷ':
		return 'き'
	case 'ｸ':
		return 'く'
	case 'ｹ':
		return 'け'
	case 'ｺ':
		return 'こ'
	case 'ｻ':
		return 'さ'
	case 'ｼ':
		return 'し'
	case 'ｽ':
		return 'す'
	case 'ｾ':
		return 'せ'
	case 'ｿ':
		return 'そ'
	case 'ﾀ':
		return 'た'
	case 'ﾁ':
		return 'ち'
	case 'ﾂ':
		return 'つ'
	case 'ﾃ':
		return 'て'
	case 'ﾄ':
		return 'と'
	case 'ﾅ':
		return 'な'
	case 'ﾆ':
		return 'に'
	case 'ﾇ':
		return 'ぬ'
	case 'ﾈ':
		return 'ね'
	case 'ﾉ':
		return 'の'
	case 'ﾊ':
		return 'は'
	case 'ﾋ':
		return 'ひ'
	case 'ﾌ':
		return 'ふ'
	case 'ﾍ':
		return 'へ'
	case 'ﾎ':
		return 'ほ'
	case 'ﾏ':
		return 'ま'
	case 'ﾐ':
		return 'み'
	case 'ﾑ':
		return 'む'
	case 'ﾒ':
		return 'め'
	case 'ﾓ':
		return 'も'
	case 'ﾔ':
		return 'や'
	case 'ﾕ':
		return 'ゆ'
	case 'ﾖ':
		return 'よ'
	case 'ﾗ':
		return 'ら'
	case 'ﾘ':
		return 'り'
	case 'ﾙ':
		return 'る'
	case 'ﾚ':
		return 'れ'
	case 'ﾛ':
		return 'ろ'
	case 'ﾜ':
		return 'わ'
	case 'ﾝ':
		return 'ん'
	case 'ﾞ':
		return '゛'
	case 'ﾟ':
		return '゜'
	default:
		return in
	}
}

func HankakuKatakanaToZenkakuHiragana(in <-chan rune, v bool) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		if v {
			var before *rune
			for r := range in {
				switch r {
				case 'ｳ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾜ':
					if before != nil {
						out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
					}
					r := r
					before = &r
				case 'ﾞ':
					if before == nil {
						out <- '゛'
					} else {
						switch *before {
						case 'ｶ':
							out <- 'が'
						case 'ｷ':
							out <- 'ぎ'
						case 'ｸ':
							out <- 'ぐ'
						case 'ｹ':
							out <- 'げ'
						case 'ｺ':
							out <- 'ご'
						case 'ｻ':
							out <- 'ざ'
						case 'ｼ':
							out <- 'じ'
						case 'ｽ':
							out <- 'ず'
						case 'ｾ':
							out <- 'ぜ'
						case 'ｿ':
							out <- 'ぞ'
						case 'ﾀ':
							out <- 'だ'
						case 'ﾁ':
							out <- 'ぢ'
						case 'ﾂ':
							out <- 'づ'
						case 'ﾃ':
							out <- 'で'
						case 'ﾄ':
							out <- 'ど'
						case 'ﾊ':
							out <- 'ば'
						case 'ﾋ':
							out <- 'び'
						case 'ﾌ':
							out <- 'ぶ'
						case 'ﾍ':
							out <- 'べ'
						case 'ﾎ':
							out <- 'ぼ'
						default:
							out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
							out <- '゛'
						}
						before = nil
					}
				case 'ﾟ':
					if before == nil {
						out <- '゜'
					} else {
						switch *before {
						case 'ﾊ':
							out <- 'ぱ'
						case 'ﾋ':
							out <- 'ぴ'
						case 'ﾌ':
							out <- 'ぷ'
						case 'ﾍ':
							out <- 'ぺ'
						case 'ﾎ':
							out <- 'ぽ'
						default:
							out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
							out <- '゜'
						}
						before = nil
					}
				default:
					if before != nil {
						out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
						before = nil
					}
					out <- hankakuKatakanaToZenkakuHiraganaSimple(r)
				}
			}
			if before != nil {
				out <- hankakuKatakanaToZenkakuHiraganaSimple(*before)
			}
		} else {
			for i := range in {
				out <- hankakuKatakanaToZenkakuHiraganaSimple(i)
			}
		}
	}()
	return out
}

func ZenkakuKatakanaToZenkakuHiragana(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			switch {
			case i >= 'ァ' && i <= 'ン', i == 'ヽ', i == 'ヾ':
				out <- 'ぁ' + i - 'ァ'
			default:
				out <- i
			}
		}
	}()
	return out
}

func ZenkakuHiraganaToZenkakuKatakana(in <-chan rune) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for i := range in {
			switch {
			case i >= 'ぁ' && i <= 'ん', i == 'ゝ', i == 'ゞ':
				out <- 'ァ' + i - 'ぁ'
			default:
				out <- i
			}
		}
	}()
	return out
}

func NewKanaConverters(mode string) ([]func(<-chan rune) <-chan rune, error) {
	options, err := NewKanaConverterOptions(mode)
	if err != nil {
		return nil, err
	}
	var converters []func(<-chan rune) <-chan rune
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
		if !options.optr {
			converters = append(converters, ZenkakuEnglishToHankakuEnglish)
		}
		if !options.optn {
			converters = append(converters, ZenkakuNumberToHankakuNumber)
		}
	}
	if options.optA {
		if !options.optR {
			converters = append(converters, HankakuEnglishToZenkakuEnglish)
		}
		if !options.optN {
			converters = append(converters, HankakuNumberToZenkakuNumber)
		}
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
			converters = append(converters, func(in <-chan rune) <-chan rune { return HankakuKatakanaToZenkakuHiragana(in, options.optV) })
		}
		converters = append(converters, ZenkakuKatakanaToHankakuKatakana)
		return converters, nil
	}
	if options.optK {
		if options.optc {
			converters = append(converters, ZenkakuKatakanaToZenkakuHiragana)
		}
		converters = append(converters, func(in <-chan rune) <-chan rune { return HankakuKatakanaToZenkakuKatakana(in, options.optV) })
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
		converters = append(converters, func(in <-chan rune) <-chan rune { return HankakuKatakanaToZenkakuHiragana(in, options.optV) })
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
