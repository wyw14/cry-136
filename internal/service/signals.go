package service

import (
	"github.com/wyw14/cry-136/internal/boron"
	"github.com/wyw14/cry-136/internal/coolant"
	"github.com/wyw14/cry-136/internal/core"
	"github.com/wyw14/cry-136/internal/cycle"
	"github.com/wyw14/cry-136/internal/irradiation"
	"github.com/wyw14/cry-136/internal/journal"
	"github.com/wyw14/cry-136/internal/model"
	"github.com/wyw14/cry-136/internal/neutron"
	"github.com/wyw14/cry-136/internal/permit"
	"github.com/wyw14/cry-136/internal/radiation"
	"github.com/wyw14/cry-136/internal/rod"
	"github.com/wyw14/cry-136/internal/scram"
)

type operationalSignal interface {
	Activate()
	Deactivate()
	SetValue(float64)
	Value() float64
	Active() bool
	Sequence() uint64
	Label() string
	Role() string
}

func collectOperationalEquipment() []model.Equipment {
	signals := []operationalSignal{
		model.NewTelemetryState("plant-telemetry"),
		model.NewAuditState("operator-audit"),
		cycle.NewCycleMetrics("cycle-progress"),
		core.NewCoreTelemetry("core-power"),
		core.NewCoreMetrics("core-margin"),
		core.NewCoreDiagnostics("core-diagnostics"),
		coolant.NewCoolantTelemetry("coolant-flow"),
		coolant.NewCoolantMetrics("coolant-margin"),
		coolant.NewCoolantDiagnostics("coolant-diagnostics"),
		neutron.NewNeutronTelemetry("neutron-flux"),
		neutron.NewNeutronMetrics("neutron-dwell"),
		neutron.NewNeutronDiagnostics("neutron-diagnostics"),
		irradiation.NewIrradiationTelemetry("irradiation-position"),
		irradiation.NewIrradiationMetrics("irradiation-lease"),
		permit.NewPermitTelemetry("permit-evidence"),
		permit.NewPermitMetrics("permit-margin"),
		radiation.NewCoordinatorState("radiation-coordination"),
		radiation.NewChannelState("radiation-channel"),
		radiation.NewRadiationState("radiation-state"),
		radiation.NewRadiationMetrics("radiation-dose"),
		rod.NewRodTelemetry("rod-position"),
		rod.NewRodMetrics("rod-motion"),
		rod.NewRodDiagnostics("rod-diagnostics"),
		scram.NewLatchState("scram-latch"),
		scram.NewIntegrationState("scram-integration"),
		scram.NewScramTelemetry("scram-telemetry"),
		scram.NewScramMetrics("scram-response"),
		boron.NewInjectionState("boron-injection"),
		boron.NewBoronTelemetry("boron-pressure"),
		boron.NewBoronMetrics("boron-delivery"),
		journal.NewIntegrationState("journal-integration"),
		journal.NewJournalTelemetry("journal-telemetry"),
		journal.NewJournalMetrics("journal-durability"),
		NewServiceMetrics("service-runtime"),
	}
	equipment := make([]model.Equipment, 0, len(signals)+1)
	for index, signal := range signals {
		signal.Deactivate()
		signal.Activate()
		signal.SetValue(float64(index + 1))
		state := "inactive"
		if signal.Active() {
			state = signal.Role()
		}
		if signal.Value() <= 0 {
			state = "unavailable"
		}
		equipment = append(equipment, model.Equipment{Name: signal.Label(), State: state, Revision: signal.Sequence()})
	}
	worker := cycle.NewCycleWorker("rod-withdrawal-supervisor")
	worker.Start()
	equipment = append(equipment, model.Equipment{Name: worker.Name(), State: "running", Revision: worker.Epoch()})
	return equipment
}
