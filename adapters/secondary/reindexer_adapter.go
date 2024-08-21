package secondary

// import (
// 	"SebStudy/infrastructure"
// 	"log"
// 	"sync"

// 	"github.com/restream/reindexer/v3"
// 	_ "github.com/restream/reindexer/v3/bindings/cproto"
// )

// const (
// 	namespaceTitle = "testdb"
// )

// type ReindexerAdapter struct {
// 	sync.RWMutex
// 	db *reindexer.Reindexer
// }

// func NewReindexerAdapter() *ReindexerAdapter {

// 	conn := reindexer.NewReindex("cproto://localhost:6534/" + namespaceTitle)

// 	if err := conn.Ping(); err != nil {
// 		log.Fatalf("failed to connect reindexer: %s\n", err)
// 		return nil
// 	}

// 	if err := conn.OpenNamespace(namespaceTitle, reindexer.DefaultNamespaceOptions(), infrastructure.Event[interface{}]{}); err != nil {
// 		log.Fatalf("failed to open namespace with err: %s", err)
// 		return nil
// 	}

// 	return &ReindexerAdapter{
// 		db: conn,
// 	}
// }

// func (r *ReindexerAdapter) Get(aggregateId string) (events []interface{}, err error) {
// 	// r.RLock()
// 	// defer r.RUnlock()
// 	query := r.db.Query(namespaceTitle).Where("aggregateId", reindexer.EQ, aggregateId)
// 	iterator := query.Exec()
// 	for iterator.Next() {
// 		ev := iterator.Object()
// 		events = append(events, ev)
// 	}

// 	return events, nil
// }

// func (r *ReindexerAdapter) Save(data interface{}) error {
// 	// r.Lock()
// 	// defer r.Unlock()
// 	// r.db.Delete()
// 	reindexerErr, err := r.db.Insert(namespaceTitle, data)
// 	if err != nil {
// 		return err
// 	} else if reindexerErr == reindexer.ErrCodeOK {
// 		return nil
// 	}

// 	return nil
// }
