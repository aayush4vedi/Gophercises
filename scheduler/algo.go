package scheduler

type Algo uint8

const(
	_ Algo = iota
	SJF
	FCFS
	FPS
	EDF
)