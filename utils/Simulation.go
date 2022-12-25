package utils

import (
	"errors"
	"fmt"
)

// Y coord down, X -- right
type Simulation[T any] struct {
	simulationMap      [][]T
	minPoint, maxPoint IntVector
}

func (s *Simulation[T]) IsInSimulation(point IntVector) bool {
	return s.minPoint.X <= point.X && point.X <= s.maxPoint.X &&
		s.minPoint.Y <= point.Y && point.Y <= s.maxPoint.Y
}

func (s *Simulation[T]) TranslateCoordinates(in IntVector) (IntVector, error) {
	if !s.IsInSimulation(in) {
		return IntVector{}, errors.New(fmt.Sprintf("Point %v falls outside of Simulation", in))
	}
	out := in.Sub(s.minPoint)
	return out, nil
}

func (s *Simulation[T]) Set(coords IntVector, item T) error {
	simCoords, err := s.TranslateCoordinates(coords)
	if err != nil {
		return err
	}
	s.simulationMap[simCoords.Y][simCoords.X] = item
	return nil
}

func (s *Simulation[T]) Get(coords IntVector) (T, error) {
	simCoords, err := s.TranslateCoordinates(coords)
	if err != nil {
		return s.simulationMap[0][0], err
	}
	return s.simulationMap[simCoords.Y][simCoords.X], nil
}

func (s *Simulation[T]) Visualize(printItemFunc func(T)) {
	for _, row := range s.simulationMap {
		for _, item := range row {
			printItemFunc(item)
		}
		println()
	}
}

func InitSimulation[T any](minPoint, maxPoint IntVector, defaultSym T) Simulation[T] {
	horSize := maxPoint.X - minPoint.X + 1
	verSize := maxPoint.Y - minPoint.Y + 1

	caveMap := make([][]T, verSize)
	for i := range caveMap {
		caveMap[i] = make([]T, horSize)
		for j := range caveMap[i] {
			caveMap[i][j] = defaultSym
		}
	}

	return Simulation[T]{
		simulationMap: caveMap,
		minPoint:      minPoint,
		maxPoint:      maxPoint,
	}
}
