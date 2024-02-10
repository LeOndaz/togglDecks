package common

import (
	"errors"
	"strings"
)

func GetSuitFullName(firstLetter byte) (string, error) {
	for _, suit := range Suits {
		if strings.ToUpper(string(suit[0])) == strings.ToUpper(string(firstLetter)) {
			return suit, nil
		}
	}

	return "", errors.New("suit not found")
}

func GetRankFullName(firstLetter byte) (string, error) {
	for _, rank := range Ranks {
		if strings.ToUpper(string(rank[0])) == strings.ToUpper(string(firstLetter)) {
			return rank, nil
		}
	}

	return "", errors.New("rank not found")
}
