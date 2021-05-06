package database

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	SavePoints(1, &[]float64{123, 123, 43434.1})
	SavePoints(2, &[]float64{1234, 123, 43434.1})
	SavePoints(3, &[]float64{2342, 34234, 43434.1})
	ps, _ := QueryPoints(2)
	fmt.Println(ps)
	DelPoints(0)
	ps, _ = QueryPoints(234234234)
	fmt.Println(ps)
	SavePoints(1, &[]float64{123, 123, 43434.1})
	SavePoints(2, &[]float64{1234, 123, 43434.1})
	SavePoints(3, &[]float64{2342, 34234, 43434.1})
	DelPoints(0)
	ps, _ = QueryPoints(0)
	fmt.Println(ps)

}
