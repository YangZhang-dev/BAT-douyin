package utils

import (
	"BAT-douyin/dao/database"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
)

//拿到所有视频（用于响应）
func ConvertVideoList(videos []model.Video) ([]Res.BaseVideoRes, error) {
	videoList := make([]Res.BaseVideoRes, len(videos))
	if len(videos) != 0 {
		for i := 0; i < len(videos); i++ {
			var user model.User
			//查找到user
			err := database.DB.First(&user, "id", videos[i].AuthorID).Error
			if err != nil {
				return nil, err
			}
			userRes := ConvertUserRes(user)
			videoList[i].Author = userRes

			videoList[i].Id = videos[i].ID
			videoList[i].CommentCount = videos[i].CommentCount
			videoList[i].CoverUrl = videos[i].CoverUrl
			videoList[i].PlayUrl = videos[i].PlayUrl
			videoList[i].FavoriteCount = videos[i].FavoriteCount
			videoList[i].Title = videos[i].Title

		}
	}
	return videoList, nil
}

func ConvertUserRes(user model.User) Res.BaseUserRes {
	var userRes Res.BaseUserRes

	userRes.Id = user.ID
	userRes.Name = user.UserName
	userRes.Followcount = user.FollowCount
	userRes.Followercount = user.FollowerCount
	userRes.Avatar = user.Avatar
	userRes.Signature = user.Signature
	userRes.BackgroundImage = user.BackgroundImage
	userRes.TotalFavorited = user.TotalFavorited
	userRes.FavoriteCount = user.FavoriteCount

	return userRes

}
