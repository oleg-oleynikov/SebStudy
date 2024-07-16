package secondary

import (
	"SebStudy/infrastructure"
	"log"
	"sync"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
)

type ReindexerAdapter struct {
	sync.RWMutex
	db *reindexer.Reindexer
}

func NewReindexerAdapter() *ReindexerAdapter {
	conn := reindexer.NewReindex("cproto://localhost:6534/testdb")

	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to connect reindexer: %s\n", err)
		return nil
	}

	namespaceTitle := "testdb"

	// if err := conn.OpenNamespace(namespaceTitle, reindexer.DefaultNamespaceOptions(), events.ResumeSended{}); err != nil {
	// 	log.Fatalf("failed to create namespace %s with error %s\n", namespaceTitle, err)
	// 	return nil
	// }

	if err := conn.RegisterNamespace(namespaceTitle, reindexer.DefaultNamespaceOptions(), infrastructure.EventMessage[interface{}]{}); err != nil {
		log.Fatalf("failed to register namespace with error: %s", err)
		return nil
	}

	if err := conn.OpenNamespace(namespaceTitle, reindexer.DefaultNamespaceOptions(), infrastructure.EventMessage[interface{}]{}); err != nil {
		log.Fatalf("failed to open namespace with err: %s", err)
		return nil
	}

	return &ReindexerAdapter{
		db: conn,
	}
}

func (r *ReindexerAdapter) Get(aggregateId string) (events []interface{}, err error) {
	r.RLock()
	defer r.RUnlock()
	query := r.db.Query("testdb").Where("aggregateId", reindexer.EQ, aggregateId)
	iterator := query.Exec()
	for iterator.Next() {
		ev := iterator.Object()
		events = append(events, ev)
	}

	return events, nil
}

func (r *ReindexerAdapter) Save(data interface{}) error {
	r.Lock()
	defer r.Unlock()
	reindexerErr, err := r.db.Insert("testdb", data)
	if err != nil {
		return err
	} else if reindexerErr == reindexer.ErrCodeOK {
		return nil
	}
	return nil
}
