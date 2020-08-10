// @program: hello-students
// @author: edte
// @create: 2020-08-10 19:36
package service

import (
	"math/rand"
	"time"
)

type studentData struct {
	allNumber     int
	maleNumber    int
	femaleNumber  int
	malePercent   string
	femalePercent string
	regions       map[string]int
}

var student = studentData{
	allNumber:     10000,
	maleNumber:    4000,
	femaleNumber:  6000,
	malePercent:   "40.00%",
	femalePercent: "60.00%",
	regions: map[string]int{
		"北京":  20,
		"天津":  20,
		"河北":  20,
		"山西":  23,
		"内蒙古": 20,
		"辽宁":  23,
		"吉林":  23,
		"黑龙江": 234,
		"上海":  99,
		"江苏":  82,
		"浙江":  99,
		"安徽":  88,
		"福建":  99,
		"江西":  99,
		"山东":  88,
		"河南":  1,
		"湖北":  99,
		"湖南":  998,
		"广东":  98,
		"广西":  8,
		"海南":  12,
		"重庆":  89,
		"四川":  12,
		"贵州":  23,
		"云南":  12,
		"西藏":  98,
		"陕西":  23,
		"甘肃":  98,
		"青海":  82,
		"宁夏":  32,
		"新疆":  23,
		"香港":  23,
		"澳门":  82,
		"台湾":  82,
	},
}

// GetGenderNumber 
func GetGenderNumber(c string) int {
	if c == "男" {
		return student.maleNumber
	}
	return student.femaleNumber
}

// GetGenderRatio
func GetGenderRatio(c string) string {
	if c == "男" {
		return student.malePercent
	}
	return student.femalePercent
}

// GetRegionNumber
func GetRegionNumber(region string) int {
	return student.regions[region]
}

var genderDesc = map[string][]string{
	"男": []string{"酷酷的男生", "英俊的男生", "豪爽的男生"},
	"女": []string{"漂亮的女生", "体贴的女生", "文静的女生"},
}

// GetRandomGenderDescription
func GetRandomGenderDescription(gender string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Intn(3)

	return genderDesc[gender][n]
}
