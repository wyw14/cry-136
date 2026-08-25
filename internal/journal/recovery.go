package journal

import "github.com/wyw14/cry-136/internal/model"

type Recovery struct {
	Events []model.Event
	Revision uint64
}

func (r Recovery) Restore() map[string]string { return Replay(r.Events) }
func (r Recovery) Ready() bool { return len(r.Events) > 0 && r.Revision > 0 }
