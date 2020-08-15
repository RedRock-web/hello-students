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
	allNumber:     5100,
	maleNumber:    3600,
	femaleNumber:  1500,
	malePercent:   "70.59",
	femalePercent: "29.41",
	regions: map[string]int{
		"北京":  6,
		"天津":  8,
		"河北":  108,
		"山西":  85,
		"内蒙古": 10,
		"辽宁":  25,
		"吉林":  20,
		"黑龙江": 20,
		"上海":  8,
		"江苏":  50,
		"浙江":  50,
		"安徽":  139,
		"福建":  50,
		"江西":  92,
		"山东":  145,
		"河南":  128,
		"湖北":  150,
		"湖南":  120,
		"广东":  120,
		"广西":  55,
		"海南":  33,
		"重庆":  2862,
		"四川":  346,
		"贵州":  90,
		"云南":  105,
		"西藏":  5,
		"陕西":  105,
		"甘肃":  45,
		"青海":  15,
		"宁夏":  55,
		"新疆":  50,
	},
}

// GetGenderNumber 
func GetGenderNumber(c string) int {
	if c == "男" {
		return student.maleNumber
	}
	return student.femaleNumber
}

// GetAllNumber
func GetAllNumber() int {
	return student.allNumber
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
