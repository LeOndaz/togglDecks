package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"togglDecks/common"
)

func TestGetRankFullName(t *testing.T) {
	for _, rank := range common.Ranks {
		rankFullName, err := common.GetRankFullName(rank[0])

		assert.Nil(t, err)
		assert.Equal(t, rank, rankFullName, "expected rank=%s to have a fullName=%s", rank, rankFullName)
	}
}

func TestGetSuitFullName(t *testing.T) {
	for _, rank := range common.Ranks {
		rankFullName, err := common.GetRankFullName(rank[0])

		assert.Nil(t, err)
		assert.Equal(t, rank, rankFullName, "expected suit=%s to have a fullName=%s", rank, rankFullName)
	}
}
