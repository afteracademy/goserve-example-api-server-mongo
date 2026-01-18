package startup

import (
	auth "github.com/afteracademy/goserve-example-api-server-mongo/api/auth/model"
	blog "github.com/afteracademy/goserve-example-api-server-mongo/api/blog/model"
	contact "github.com/afteracademy/goserve-example-api-server-mongo/api/contact/model"
	user "github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"github.com/afteracademy/goserve/v2/mongo"
)

func EnsureDbIndexes(db mongo.Database) {
	go mongo.Document[auth.Keystore](&auth.Keystore{}).EnsureIndexes(db)
	go mongo.Document[auth.ApiKey](&auth.ApiKey{}).EnsureIndexes(db)
	go mongo.Document[user.User](&user.User{}).EnsureIndexes(db)
	go mongo.Document[user.Role](&user.Role{}).EnsureIndexes(db)
	go mongo.Document[blog.Blog](&blog.Blog{}).EnsureIndexes(db)
	go mongo.Document[contact.Message](&contact.Message{}).EnsureIndexes(db)
}
