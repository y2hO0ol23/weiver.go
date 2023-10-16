package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.ChineseCN] = map[string]string{
		"#admin":                              "行政",
		"#admin.allow-role":                   "允许角色",
		"#look":                               "看",
		"#move":                               "移动",
		"#review":                             "审查",
		"#admin.allow-role.value":             "价值",
		"#.subject":                           "目标",
		"#look.info":                          "信息",
		"#look.review-list":                   "评论列表",
		"#admin.allow-role.Description":       "将分数显示为角色。默认值：假",
		"#look.Description":                   "看到一些东西",
		"#move.Description":                   "将评论移至此频道",
		"#review.Description":                 "用户评论",
		"#admin.allow-role.value.Description": "设定值",
		"#.subject.Description":               "选择目标",
		"#look.info.Description":              "查看用户信息",
		"#look.review-list.Description":       "查看用户收到的评论列表",
		"#allow-role.NeedPermissions":         "机器人缺乏权限 - 管理角色",
		"#allow-role.InProgress":              "工作正在进行中",
		"#allow-role.proc.Title":              "修改选项",
		"#allow-role.proc.Description":        "允许角色",
		"#allow-role.proc.InProgress":         "论文集",
		"#allow-role.proc.Done":               "完全的",
		"#allow-role.Keep":                    "没有更改设置",
		"#look.info.IsNone":                   "没有评论",
		"#look.review-list.IsNone":            "没有评论",
		"#look.review-list.menu.Title":        "对 %s 的评论",
		"#look.review-list.menu.Page":         "%d 页",
		"#move.IsNone":                        "没有为目标撰写评论",
		"#move.Move":                          "搬来这",
		"#review.SelfReview":                  "你无法检讨自己",
		"#review.modal.Title":                 "评论 %s",
		"#review.lable.Score":                 "分数",
		"#review.lable.Title":                 "标题",
		"#review.lable.Content":               "细节",
		"$review.IsEdited":                    "此评论已编辑",
		"$review.NoAuthor":                    "已删除的评论无法恢复，因为这里不存在作者",
		"$review.DM":                          "已撰写新评论",
    }
}
