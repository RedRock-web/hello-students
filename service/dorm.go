// @program: hello-students
// @author: edte
// @create: 2020-08-10 23:14
package service

// DormForm
type DormForm struct {
	Name string
}

type dormData struct {
	pictures      []string
	configuration []string
}

var dorm = map[string]dormData{
	"知行苑": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_zxy0.jpg",
			"http://cdn.redrock.team/hello-student_zxy1.jpg",
			"http://cdn.redrock.team/hello-student_zxy2.jpg",
			"http://cdn.redrock.team/hello-student_zxy3.jpg"},
		configuration: []string{
			"公共洗衣机",
			"空调",
			"阳台(7,8舍)",
			"独立卫生间(7,8舍）",
			"四人间上下床(1,5,6舍)",
			"六人间上下床+上床下桌(7,8舍)"},
	},
	"宁静苑": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_ljy0.jpeg",
			"http://cdn.redrock.team/hello-student_ljy1.jpg",
			"http://cdn.redrock.team/hello-student_ljy2.jpg",
			"http://cdn.redrock.team/hello-student_ljy3.jpg",
		},
		configuration: []string{
			"电梯(10舍)",
			"公共洗衣机",
			"独立卫生间",
			"阳台",
			"空调",
			"四人间（上床下桌）或六人间（上下床+上床下桌）",
		},
	},
	"明理苑": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_mly0.jpeg",
			"http://cdn.redrock.team/hello-student_mly1.jpg",
			"http://cdn.redrock.team/hello-student_mly2.jpg",
			"http://cdn.redrock.team/hello-student_mly3.jpg",
		},
		configuration: []string{
			"空调",
			"阳台",
			"独立卫生间",
			"独立洗衣机",
			"六人间（上床下桌+上下床）",
		},
	},
	"兴业苑": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_yxy0.jpg",
			"http://cdn.redrock.team/hello-student_yxy1.jpg",
			"http://cdn.redrock.team/hello-student_yxy2.jpg",
			"http://cdn.redrock.team/hello-student_yxy3.png",
		},
		configuration: []string{
			"空调",
			"空调",
			"独立卫生间",
			"公共洗衣机",
			"四人间（部分为上床下桌）或六人间（上下床+上床下桌）",
		},
	},
}

// GetDormPictures
func GetDormPictures(name string) []string {
	return dorm[name].pictures
}

// GetDormConfiguration
func GetDormConfiguration(name string) []string {
	return dorm[name].configuration
}
