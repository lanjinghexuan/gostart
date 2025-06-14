package handler

import (
	"fmt"
	"time"
)

// FormatCommentTime 格式化评论时间
func FormatCommentTime(commentTime time.Time) string {
	now := time.Now()
	duration := now.Sub(commentTime)

	// 确保 commentTime 是在 now 之前的时间
	if duration < 0 {
		return "未来时间"
	}

	switch {
	case duration < 5*time.Minute:
		return fmt.Sprintf("%d分钟前", int(duration.Minutes()))
	case duration < 30*time.Minute:
		return "半小时前"
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d小时前", int(duration.Hours()))
	case duration < 3*24*time.Hour:
		return fmt.Sprintf("%d天前", int(duration.Hours()/24))
	default:
		return commentTime.Format("2006-01-02 15:04")
	}
}
