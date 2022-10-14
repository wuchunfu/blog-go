package service

import "myblog/dao"

const (
	About             = "about"
	ArticleLikeCount  = "article_like_count"
	ArticleUserLike   = "article_user_like:"
	ArticleViewsCount = "article_views_count"
	BlogViewsCount    = "blog_views_count"
	CommentLikeCount  = "comment_like_count"
	CommentUserLike   = "comment_user_like:"
	Delete            = "delete:"
	PageCover         = "page_cover"
	TalkLikeCount     = "talk_like_count"
	TalkUserLike      = "talk_user_like:"
	UniqueVisitor     = "unique_visitor"
	User              = "user:"
	UserArea          = "user_area"
	VisitorArea       = "visitor_area"
	WebsiteConfig     = "website_config"
)

var (
	articleDao      dao.ArticleDao
	categoryDao     dao.CategoryDao
	commentDao      dao.CommentDao
	friendLinkDao   dao.FriendLinkDao
	messageDao      dao.MessageDao
	menuDao         dao.MenuDao
	operationLogDao dao.OperationLogDao
	pageDao         dao.PageDao
	photoDao        dao.PhotoDao
	roleDao         dao.RoleDao
	resourceDao     dao.ResourceDao
	tagDao          dao.TagDao
	talkDao         dao.TalkDao
	userDao         dao.UserDao
	uniqueViewDao   dao.UniqueViewDao
	rdb             = dao.GetRDB()
)

var (
	blogInfoService BlogInfoService
	userService     UserService
	uniqueService   UniqueViewService
)
