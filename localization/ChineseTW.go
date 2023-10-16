package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.ChineseTW] = map[string]string{
		"#admin":                              "行政",
		"#admin.allow-role":                   "允許角色",
		"#look":                               "看",
		"#move":                               "移動",
		"#review":                             "審查",
		"#admin.allow-role.value":             "價值",
		"#.subject":                           "目標",
		"#look.info":                          "資訊",
		"#look.review-list":                   "評論列表",
		"#admin.allow-role.Description":       "將分數顯示為角色。預設值：假",
		"#look.Description":                   "看到一些東西",
		"#move.Description":                   "將評論移至此頻道",
		"#review.Description":                 "用戶評論",
		"#admin.allow-role.value.Description": "設定值",
		"#.subject.Description":               "選擇目標",
		"#look.info.Description":              "查看使用者資訊",
		"#look.review-list.Description":       "查看用戶收到的評論列表",
		"#allow-role.NeedPermissions":         "機器人缺乏權限 - 管理角色",
		"#allow-role.InProgress":              "工作正在進行中",
		"#allow-role.proc.Title":              "修改選項",
		"#allow-role.proc.Description":        "允許角色",
		"#allow-role.proc.InProgress":         "論文集",
		"#allow-role.proc.Done":               "完全的",
		"#allow-role.Keep":                    "沒有更改設定",
		"#look.info.IsNone":                   "沒有評論",
		"#look.review-list.IsNone":            "沒有評論",
		"#look.review-list.menu.Title":        "對 %s 的評論",
		"#look.review-list.menu.Page":         "%d 頁",
		"#move.IsNone":                        "沒有為目標撰寫評論",
		"#move.Move":                          "搬來這",
		"#review.SelfReview":                  "你無法檢討自己",
		"#review.modal.Title":                 "評論 %s",
		"#review.lable.Score":                 "分數",
		"#review.lable.Title":                 "標題",
		"#review.lable.Content":               "細節",
		"$review.IsEdited":                    "此評論已編輯",
		"$review.NoAuthor":                    "已刪除的評論無法恢復，因為這裡不存在作者",
		"$review.DM":                          "已撰寫新評論",
    }
}
