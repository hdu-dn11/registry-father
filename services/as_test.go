package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"registry-father/model"
	"strconv"
	"testing"
)

func TestGetASInfoList(t *testing.T) {
	mockList := make([]*model.ASInfo, 0)
	_ = os.RemoveAll("./data")
	_ = os.MkdirAll("./data", 0644)
	for i := 0; i < 10; i++ {
		as := &model.ASInfo{
			ASN:   uint32(1234567890 + i),
			Owner: strconv.Itoa(1234567890 + i),
			IPv4:  []string{""},
			IPv6:  []string{""},
		}
		err := SaveASInfo(as)
		if err != nil {
			t.Fatal(err)
		}
		mockList = append(mockList, as)
	}

	list, err := GetASInfoList()
	if err != nil {
		return
	}
	assert.Equal(t, mockList, list, "list should be equal")
	_ = os.RemoveAll("./data")
}

func TestSaveASInfo(t *testing.T) {
	as := &model.ASInfo{
		ASN:   uint32(1234567890),
		Owner: strconv.Itoa(1234567890),
		IPv4:  []string{""},
		IPv6:  []string{""},
	}
	for i := 0; i < 3; i++ {
		_ = SaveASInfo(as)
	}
	_ = os.RemoveAll("./data")
}
