package repository

import (
	"fmt"
	"gin-template/app/dto"

	env "gin-template/config"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mongo"
)

const (
	PageSizeLimit uint = 200
)

// go get -u github.com/upper/db/v4
// go get -u github.com/upper/db/v4/adapter/mongo
// https://upper.io/v4/getting-started/agnostic-db-api/
// https://github.com/upper/db
// https://upper.io/v4/adapter/mongo/
type Repository struct {
	conn       mongo.ConnectionURL
	collection string
}

func New(collection string) *Repository {
	r := &Repository{}
	r.conn = mongo.ConnectionURL{
		Host:     env.Get("MONGO_HOST"),
		Database: env.Get("MONGO_DB"),
		User:     env.Get("MONGO_USERNAME"),
		Password: env.Get("MONGO_PASSWORD"),
		Options: map[string]string{
			"authSource": "admin",
		},
	}
	r.collection = collection
	return r
}

func (s *Repository) Exec(callback func(collection db.Collection)) {
	sess, err := mongo.Open(s.conn)
	if err != nil {
		_ = fmt.Errorf("failed to open connection to mongodb %v", err)
		return
	}
	defer sess.Close()

	callback(sess.Collection(s.collection))
}

func (s *Repository) List(request *dto.ListRequest, out *dto.ListResult) (err error) {
	s.Exec(func(col db.Collection) {
		/// https://upper.io/v4/getting-started/agnostic-db-api/
		// apply filter
		cond := db.And(db.Cond{"_id": db.IsNotNull()})
		if request.QueryCriteria != nil {
			for k, v := range request.QueryCriteria {
				if k == "" || len(v) == 0 {
					continue
				}
				c := db.Cond{fmt.Sprintf("%s IN", k): v}
				cond = cond.And(c)
			}
		}

		r := col.Find(cond)
		out.TotalItems, _ = r.Count()

		// apply ordering
		if request.SortBy != nil {
			for k, v := range request.SortBy {
				if k == "" {
					continue
				}

				order := k
				if v == "desc" {
					order = fmt.Sprintf("-%s", k)
				}
				r = r.OrderBy(order)
			}
		}

		// apply paging
		r = r.Paginate(request.PageSize)
		out.TotalPages, _ = r.TotalPages()
		r = r.Page(request.PageIndex)

		// popup result
		err = r.All(out.Items)
		if err != nil {
			_ = fmt.Errorf("failed to query db, %v", err)
		}
	})
	return
}
