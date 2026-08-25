package service

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wyw14/cry-136/internal/boron"
	"github.com/wyw14/cry-136/internal/coolant"
	"github.com/wyw14/cry-136/internal/core"
	"github.com/wyw14/cry-136/internal/cycle"
	"github.com/wyw14/cry-136/internal/irradiation"
	"github.com/wyw14/cry-136/internal/journal"
	"github.com/wyw14/cry-136/internal/model"
	"github.com/wyw14/cry-136/internal/neutron"
	"github.com/wyw14/cry-136/internal/permit"
	"github.com/wyw14/cry-136/internal/rod"
	"github.com/wyw14/cry-136/internal/scram"
)

type Runtime struct {
	mu          sync.RWMutex
	cycle       *cycle.Coordinator
	core        *core.Controller
	rods        *rod.Controller
	rodFlow     *rod.Coordinator
	coolant     *coolant.SwitchCoordinator
	neutron     *neutron.Coordinator
	recorder    *neutron.ShutdownRecorder
	irradiation *irradiation.Coordinator
	permit      *permit.Evaluator
	scram       *scram.Coordinator
	boron       *boron.Injector
	journal     *journal.Store
	operations  []model.Operation
	incidents   []model.Incident
	signals     []model.Equipment
	tripState   *cycle.TripState
}

func NewRuntime(dataDir string) *Runtime {
	rods := rod.NewController()
	runtime := &Runtime{cycle: cycle.NewCoordinator(), core: core.NewController(), rods: rods, rodFlow: rod.NewCoordinator(rods), coolant: coolant.NewSwitchCoordinator(), neutron: neutron.NewCoordinator(), recorder: neutron.NewShutdownRecorder(), irradiation: irradiation.NewCoordinator(), permit: permit.NewEvaluator(), scram: scram.NewCoordinator(), boron: boron.NewInjector(), journal: journal.NewStore(), signals: collectOperationalEquipment(), tripState: cycle.NewTripState()}
	runtime.permit.SetNeutron(true)
	runtime.permit.SetCoolant(true)
	runtime.permit.SetRod(true)
	runtime.permit.SetRadiation(true)
	runtime.core.AppendNeutron(0.1)
	return runtime
}

func (r *Runtime) StartCycle() model.CycleSnapshot {
	r.mu.Lock()
	defer r.mu.Unlock()
	snapshot := r.cycle.Start()
	r.tripState.BeginCycle()
	r.rodFlow.Retract(snapshot.Epoch)
	r.core.ApplyPower(1)
	r.core.ApplyCooling(1)
	r.journal.Append(model.NewEvent("cycle-start", snapshot.ID, string(snapshot.Phase)))
	return snapshot
}

func (r *Runtime) TriggerScram(reason string) model.Snapshot {
	r.mu.Lock()
	defer r.mu.Unlock()
	snapshot := r.cycle.Trip()
	r.tripState.Latch(snapshot.Epoch)
	r.scram.Trigger(reason)
	r.rodFlow.Insert(snapshot.Epoch)
	r.core.ApplyPower(0)
	// Commit the neutron decay proof before publishing or recovering. The
	// recorder samples the falling flux, then is flushed (marked durable) so
	// the decay curve is reliably persisted. Until DurableProof() holds, the
	// curve is an incomplete record and shutdown may not be published.
	r.recorder.RecordDecay(0)
	r.recorder.Durable()
	r.recorder.PublishState("shutdown")
	r.journal.Append(model.NewEvent("scram", snapshot.ID, reason))
	r.incidents = append(r.incidents, model.Incident{ID: uuid.New(), Severity: "critical", Message: reason, CreatedAt: time.Now().UTC()})
	return r.snapshotLocked()
}

func (r *Runtime) Snapshot() model.Snapshot {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.snapshotLocked()
}

func (r *Runtime) snapshotLocked() model.Snapshot {
	cycleSnapshot := r.cycle.Snapshot()
	equipment := append([]model.Equipment{r.rods.Equipment(), {Name: "coolant-train", State: r.coolant.Status(), Revision: uint64(r.coolant.Flow())}}, r.signals...)
	return model.Snapshot{Cycle: cycleSnapshot, Core: r.core.Snapshot(), Operations: append([]model.Operation(nil), r.operations...), Equipment: equipment, Interlocks: []model.Interlock{{Name: "scram", Engaged: r.scram.Latched(), Reason: r.scram.Reason()}, {Name: "power-permit", Engaged: !r.permit.Allowed(), Reason: r.permit.Summary()}}, Incidents: append([]model.Incident(nil), r.incidents...), UpdatedAt: time.Now().UTC()}
}
