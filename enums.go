package vinderman

type EFortCollectedState string

const (
	EFortCollectedStateUnknown        EFortCollectedState = "Unknown"
	EFortCollectedStateNew            EFortCollectedState = "New"
	EFortCollectedStateKnown          EFortCollectedState = "Known"
	EFortCollectedStateNewlyCollected EFortCollectedState = "NewlyCollected"
	EFortCollectedStateCollected      EFortCollectedState = "Collected"
	EFortCollectedStateNewBest        EFortCollectedState = "NewBest"
	EFortCollectedStateNewRecord      EFortCollectedState = "NewRecord"
	EFortCollectedStateNewLocation    EFortCollectedState = "NewLocation"
	EFortCollectedStateNewlyCompleted EFortCollectedState = "NewlyCOmpleted"
	EFortCollectedStateComplete       EFortCollectedState = "Complete"
)
