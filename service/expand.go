// @program: hello-students
// @author: edte
// @create: 2020-08-11 16:34
package service

// ExpandForm
type ExpandForm struct {
	Name string `form:"name" binding:"required"`
}

type expandData struct {
	pictures     []string
	introduction string
}

var expand = map[string]expandData{
	"努力学习": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_llxx0.JPG",
			"http://cdn.redrock.team/hello-student_llxx1.jpg",
			"http://cdn.redrock.team/hello-student_llxx2.jpg",
		},
		introduction: "在重邮，你可以在新老图书馆的自习室浏览专业相关的书，可以在空闲的教室自习；同时你也可以利用专业书籍、网课视频、学霸笔记来提高自己的知识水平……",
	},
	"展现自我": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_zxzw0.png",
			"http://cdn.redrock.team/hello-student_zxzw1.jpg",
			"http://cdn.redrock.team/hello-student_zxzw2.jpg",
		},
		introduction: "在重邮，你总能找到一个属于自己的舞台，或是在新生辩论赛上舌战群雄，或是在演讲比赛中语惊四座，又或是在运动场上独领风骚……",
	},
	"发展兴趣": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_fxxq0.png",
			"http://cdn.redrock.team/hello-student_fxxq1.jpg",
			"http://cdn.redrock.team/hello-student_fxxq2.png",
		},
		introduction: "在重邮，你可以尽情发展自己的兴趣。无论你喜欢歌唱舞蹈、书法绘画，还是体育运动、志愿服务。你都可以找到对应的社团组织。重邮近百个学生社团，总有一个适合你！",
	},
	"拓展能力": {
		pictures: []string{
			"http://cdn.redrock.team/hello-student_tzll0.jpg",
			"http://cdn.redrock.team/hello-student_tzll1.jpg",
			"http://cdn.redrock.team/hello-student_tzll2.jpg",
		},
		introduction: "在重邮，你可以通过各种方式提升自己的能力，加入志愿者团队、参加科研竞赛、加入老师的实验室……在一次次付出中，收获能力、知识、成功……",
	},
}

// GetExpandPictures
func GetExpandPictures(name string) []string {
	return expand[name].pictures
}

// GetExpandIntroduction
func GetExpandIntroduction(name string) string {
	return expand[name].introduction
}
