// @program: hello-students
// @author: edte
// @create: 2020-08-11 16:06
package service

// CanteenForm
type CanteenForm struct {
	Name string
}

type canteenData struct {
	Pictures     []string
	Introduction string
}

var canteen = map[string]canteenData{
	// 中心食堂
	"zxst": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_zxst0.jpg",
			"http://cdn.redrock.team/hello-student_zxst1.jpg",
			"http://cdn.redrock.team/hello-student_zxst2.jpg",
			"http://cdn.redrock.team/hello-student_zxst3.jpg",
		},
		Introduction: "中心食堂与雨红莲广场相邻，位于学校的中心且菜式整体价格较低、种类繁多，小面深得同学们的喜爱",
	},
	// 大西北食堂
	"dxb": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_dxb0.jpg",
			"http://cdn.redrock.team/hello-student_dxb1.jpg",
			"http://cdn.redrock.team/hello-student_dxb2.jpg",
		},
		Introduction: "大西北食堂在中心食堂的背面，食堂空间较少，提供的菜式以清真食物为主。如果你想品尝西北美食，大西北是一个不错的选择。",
	},
	// 兴业苑食堂
	"yxyst": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_yxyst0.jpg",
			"http://cdn.redrock.team/hello-student_yxyst1.jpg",
			"http://cdn.redrock.team/hello-student_yxyst2.jpg",
		},
		Introduction: "兴业苑食堂与中心食堂一样，是老一派的食堂。菜式种类繁多，价格较低。食堂的小面、三鲜米线广受好评。",
	},
	// 红高粱食堂
	"hgl": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_hgl0.jpg",
			"http://cdn.redrock.team/hello-student_hgl1.png",
			"http://cdn.redrock.team/hello-student_hgl2.png",
		},
		Introduction: "红高粱是重邮地理位置最高的食堂，也是最好的食堂之一。食堂大而整洁。特色食物：石溪米线、肠粉",
	},
	// 千喜鹤食堂
	"qxh": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_qxh0.jpg",
			"http://cdn.redrock.team/hello-student_qxh1.png",
			"http://cdn.redrock.team/hello-student_qxh2.png\n\n",
		},
		Introduction: "千喜鹤食堂是知名网红食堂，食堂装修豪华而富有文艺气息，提供的美食种类很多且独特：蛋糕、奶茶、炸鸡、烧腊。",
	},
	// 延生食堂
	"ysst": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_ysst0.jpg",
			"http://cdn.redrock.team/hello-student_ysst1.jpg",
			"http://cdn.redrock.team/hello-student_ysst2.jpg",
			"http://cdn.redrock.team/hello-student_ysst3.jpg",
		},
		Introduction: "延生食堂在学校的口碑很好，食物很有特色：铁板炒饭、小炒、冒菜，深受同学们的好评。",
	},
	// 樱花食堂
	"yhst": {
		Pictures: []string{
			"http://cdn.redrock.team/hello-student_yhst0.jpg",
			"http://cdn.redrock.team/hello-student_yhst1.jpg",
			"http://cdn.redrock.team/hello-student_yhst2.jpg",
			"http://cdn.redrock.team/hello-student_yhst3.jpg",
		},
		Introduction: "樱花食堂是学校最新的食堂，食堂较小但光线很足。早餐午餐晚餐提供的菜式种类多而便宜。 ",
	},
}

// GetCanteenPictures
func GetCanteenPictures(name string) []string {
	return canteen[name].Pictures
}

// GetCanteenIntroduction
func GetCanteenIntroduction(name string) string {
	return canteen[name].Introduction
}
