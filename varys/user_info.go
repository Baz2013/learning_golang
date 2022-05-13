package varys

type UserInfo struct {
	Name string `json:"name"`
}

//模拟第三方库，模拟在开发的时候，这个库的方法还没有实现

func GetInfoByUID(uid int64) (user *UserInfo, err error) {

	return &UserInfo{Name: "gucb3"}, nil
}
