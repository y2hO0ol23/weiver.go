package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Thai] = map[string]string{
		"#admin":                              "ผู้ดูแลระบบ",
		"#admin.allow-role":                   "อนุญาตบทบาท",
		"#look":                               "ดู",
		"#move":                               "เคลื่อนไหว",
		"#review":                             "ทบทวน",
		"#admin.allow-role.value":             "ค่า",
		"#.subject":                           "เป้า",
		"#look.info":                          "ข้อมูล",
		"#look.reviews":                       "ความคิดเห็น",
		"#admin.allow-role.Description":       "แสดงคะแนนตามบทบาท ค่าเริ่มต้น: เท็จ",
		"#look.Description":                   "เห็นอะไรบางอย่าง",
		"#move.Description":                   "ย้ายรีวิวมาที่ช่องนี้",
		"#review.Description":                 "รีวิวจากผู้ใช้",
		"#admin.allow-role.value.Description": "ตั้งค่า",
		"#.subject.Description":               "เลือกเป้าหมาย",
		"#look.info.Description":              "ดูข้อมูลผู้ใช้",
		"#look.reviews.Description":           "ดูบทวิจารณ์ที่หัวข้อได้รับ",
		"#allow-role.NeedPermissions":         "บอทขาดสิทธิ์ - จัดการบทบาท",
		"#allow-role.InProgress":              "อยู่ระหว่างดำเนินการ",
		"#allow-role.proc.Title":              "แก้ไขตัวเลือก",
		"#allow-role.proc.Description":        "อนุญาตบทบาท",
		"#allow-role.proc.InProgress":         "กำลังดำเนินการ",
		"#allow-role.proc.Done":               "สมบูรณ์",
		"#allow-role.Keep":                    "ไม่มีการเปลี่ยนแปลงการตั้งค่า",
		"#look.info.IsNone":                   "ไม่มีบทวิจารณ์",
		"#look.reviews.IsNone":                "ไม่มีบทวิจารณ์",
		"#look.reviews.menu.Title":            "บทวิจารณ์สำหรับ %s",
		"#look.reviews.menu.Page":             "%d หน้า",
		"#move.IsNone":                        "ไม่มีการเขียนบทวิจารณ์สำหรับเป้าหมาย",
		"#move.Move":                          "ย้ายมาที่นี่",
		"#review.SelfReview":                  "คุณไม่สามารถทบทวนตัวเองได้",
		"#review.modal.Title":                 "รีวิว %s",
		"#review.lable.Score":                 "คะแนน",
		"#review.lable.Title":                 "ชื่อ",
		"#review.lable.Content":               "รายละเอียด",
		"$review.IsEdited":                    "รีวิวนี้ได้รับการแก้ไขแล้ว",
		"$review.NoAuthor":                    "บทวิจารณ์ที่ถูกลบไม่สามารถกู้คืนได้เนื่องจากไม่มีผู้เขียนอยู่ที่นี่",
		"$review.DM":                          "มีการเขียนบทวิจารณ์ใหม่แล้ว",
    }
}
