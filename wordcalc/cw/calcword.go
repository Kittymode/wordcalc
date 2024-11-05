package cw

import (
	"strconv"
	"strings"
)

func CalcWord(word1 string, word2 string, operator string) string {
	switch operator {
	case " + ":
		word2 = word2[1 : len(word2)-1]
		word1 = word1[1 : len(word1)-1]
		if len(word1) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		if len(word2) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		return word1 + word2
	case " - ":
		word2 = word2[1 : len(word2)-1]
		word1 = word1[1 : len(word1)-1]
		if len(word1) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		if len(word2) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		result := strings.ReplaceAll(word1, word2, "")
		return result
	case " / ":
		word1 = word1[1 : len(word1)-1]
		if len(word1) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		word2num, err := strconv.Atoi(word2)
		if err != nil {
			panic("ошибка")
		}
		if word2num > 10 {
			panic("делитель не может быть больше 10 ")
		} else if word2num < 1 {
			panic("делитель не может быть меньше 1")
		}
		resultLen := len(word1) / word2num
		result := word1[:resultLen]
		return result
	case " * ":
		word1 = word1[1 : len(word1)-1]
		if len(word1) > 10 {
			panic("длинна строки должна быть не более десяти символов")
		}
		word2num, err := strconv.Atoi(word2)
		if err != nil {
			panic("ошибка")
		}
		if word2num > 10 {
			panic("множитель не может быть больше 10 ")
		} else if word2num < 1 {
			panic("множитель не может быть меньше 1")
		}

		result := strings.Repeat(word1, word2num)
		if len(result) > 40 {
			result = result[:40] + "..."
		} else if len(result) == 40 {
			result = result[:40]
		} else if len(result) < 40 {
			return result
		}
		return result
	default:
		panic("неизвестная операция")
	}
}
