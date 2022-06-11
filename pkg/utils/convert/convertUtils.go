package convert

import (
	"BAT-douyin/dao/dcomment"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"errors"
)

// ConvertVideoList 拿到所有视频（用于响应）
func ConvertVideoList(videos []*model.Video, u *model.User) ([]Res.BaseVideoRes, error) {
	videoList := make([]Res.BaseVideoRes, len(videos))
	for i := 0; i < len(videos); i++ {
		var author *model.User
		//查找到user
		author, exists := duser.GetById(videos[i].AuthorID)
		if !exists {
			return nil, errors.New("user not exists")
		}
		userRes := ConvertUserRes(author, u)
		videoList[i].Author = userRes
		videoList[i].Id = videos[i].ID
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].CoverUrl = videos[i].CoverUrl
		videoList[i].PlayUrl = videos[i].PlayUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].Title = videos[i].Title
		is := dvideo.IsFavoriteVideo(u, videos[i])
		videoList[i].IsFavorite = is
	}

	return videoList, nil
}

func ConvertUserRes(author, u *model.User) Res.BaseUserRes {
	var userRes Res.BaseUserRes

	userRes.Id = author.ID
	userRes.Name = author.UserName
	userRes.Followcount = author.FollowCount
	userRes.Followercount = author.FollowerCount
	userRes.Avatar = author.Avatar
	userRes.Signature = author.Signature
	userRes.BackgroundImage = author.BackgroundImage
	userRes.TotalFavorited = author.TotalFavorited
	userRes.FavoriteCount = author.FavoriteCount
	//is :=false
	//if u.ID!=0 {
	//	is = duser.IsFollowUser(u, author)
	//}
	is := duser.IsFollowUser(u, author)
	userRes.Isfollow = is
	return userRes

}

// op为2表示粉丝列表,1为关注列表
func ConvertUserListRes(users []model.FollowUser, u *model.User, op uint) ([]Res.BaseUserRes, error) {

	userRes := make([]Res.BaseUserRes, len(users))
	for i := 0; i < len(users); i++ {
		var taru *model.User
		var ok bool
		if op == 2 {
			taru, ok = duser.GetById(users[i].UserID)
		} else {
			taru, ok = duser.GetById(users[i].ToUserID)
		}
		if !ok {
			return nil, errors.New("user not exists")
		}

		userRes[i].Id = taru.ID
		userRes[i].Name = taru.UserName
		userRes[i].Followercount = taru.FollowerCount
		userRes[i].Followcount = taru.FollowCount
		userRes[i].Avatar = taru.Avatar
		userRes[i].Signature = taru.Signature
		userRes[i].BackgroundImage = taru.BackgroundImage
		var is bool
		if u.ID == 0 {
			is = false
		} else {
			is = duser.IsFollowUser(u, taru)
		}
		userRes[i].Isfollow = is
	}
	return userRes, nil
}
func GetCommentList(v *model.Video, u *model.User) ([]Res.BaseCommentListRes, error) {
	commentlist := dcomment.GetList(v)
	reslist := make([]Res.BaseCommentListRes, len(commentlist))

	for i := 0; i < len(commentlist); i++ {
		reslist[i].Id = commentlist[i].ID
		commentUser, ok := duser.GetById(commentlist[i].AuthorID)
		if !ok {
			return nil, errors.New("user not exists")
		}
		reslist[i].Author.Id = commentUser.ID
		reslist[i].Author.Name = commentUser.UserName
		reslist[i].Author.Followcount = commentUser.FollowCount
		reslist[i].Author.Followercount = commentUser.FollowerCount
		reslist[i].Author.Avatar = commentUser.Avatar
		reslist[i].Author.Signature = commentUser.Signature
		reslist[i].Author.BackgroundImage = commentUser.BackgroundImage
		reslist[i].Author.TotalFavorited = commentUser.TotalFavorited
		reslist[i].Author.FavoriteCount = commentUser.FavoriteCount
		is := false
		is = duser.IsFollowUser(u, commentUser)
		reslist[i].Author.Isfollow = is
		reslist[i].Content = commentlist[i].Content
		reslist[i].CreateDate = commentlist[i].CreatedAt.Format("01-02 15:04:05")
	}
	return reslist, nil
}
