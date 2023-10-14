package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Korean] = map[string]string{
		// 슬래시 명령어
		"#allow-role":             "역할허용",
		"#allow-role.Description": "관리자 전용 - 기본:거짓",

		"#look":             "보기",
		"#look.Description": "무언가 보기",

		"#move":             "옮기기",
		"#move.Description": "이 채널로 리뷰 옮기기",

		"#review":             "리뷰",
		"#review.Description": "유저 리뷰하기",

		// 서브 명령어
		"#allow-role.value":             "값",
		"#allow-role.value.Description": "값 설정",

		"#.subject":             "대상",
		"#.subject.Description": "대상 선택",

		"#look.info":                    "정보",
		"#look.info.Description":        "유저 정보 보기",
		"#look.review-list":             "리뷰목록",
		"#look.review-list.Description": "유저가 받은 리뷰목록 보기",

		// 텍스트
		"#allow-role.InProgress":       "작업 진행 중",
		"#allow-role.proc.Title":       "옵션 수정",
		"#allow-role.proc.Description": "역할허용",
		"#allow-role.proc.InProgress":  "진행 중",
		"#allow-role.proc.Done":        "완료",
		"#allow-role.Keep":             "바뀐 설정이 없습니다",

		"#look.info.IsNone":            "리뷰 없음",
		"#look.review-list.IsNone":     "리뷰 없음",
		"#look.review-list.menu.Title": "%s의 리뷰",
		"#look.review-list.menu.Page":  "%d 페이지",

		"#move.IsNone": "대상에게 작성한 리뷰가 없습니다",
		"#move.Move":   "이곳으로 옮기기",

		"#review.SelfReview":    "자신을 리뷰할 수 없습니다",
		"#review.modal.Title":   "%s 리뷰하기",
		"#review.lable.Score":   "점수",
		"#review.lable.Title":   "제목",
		"#review.lable.Content": "내용",

		"$review.IsEdited": "해당 리뷰는 수정되었습니다.",
		"$review.NoAuthor": "삭제된 해당 리뷰의 글쓴이가 서버에 존재하지 않습니다",
		"$review.DM":       "새로운 리뷰가 작성되었습니다",
	}
}