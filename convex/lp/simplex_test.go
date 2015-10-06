package lp

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/gonum/matrix/mat64"
)

// x = [0.7287793210009457 -0.9371471942974932 -14.017213937483529]
// obj = -7.010518715722229

const convergenceTol = 1e-10

// This one fails because of floting point error in solve.
// a = &{{2 6 6 [0 0 0.44669228806208794 0 0 0 0 -0.23053805434391994 -0.8084250657828511 0 0 0]} 2 6}
// b = [0 -1.2291539097501485]
// c =  [0 2.3537032834646823 0 0 0 0]

func TestSimplex(t *testing.T) {
	// TODO(btracey): Add in Beale test
	for _, test := range []struct {
		A mat64.Matrix
		b []float64
		c []float64
		//initialBasic []int
		tol float64
	}{
		/*
			{
				// Basic feasible LP
				A: mat64.NewDense(2, 4, []float64{
					-1, 2, 1, 0,
					3, 1, 0, 1,
				}),
				b: []float64{4, 9},
				c: []float64{-1, -2, 0, 0},
				//initialBasic: nil,
				tol: 0,
			},
		*/
		/*
			{
				// Zero row that causes linear solver failure
				A: mat64.NewDense(3, 5, []float64{0.09917822373225804, 0, 0, -0.2588175087223661, -0.5935518220870567, 1.301111422556007, 0.12220247487326946, 0, 0, -1.9194869979254463, 0, 0, 0, 0, -0.8588221231396473}),
				b: []float64{0, 0, 0},
				c: []float64{0, 0.598992624019304, 0, 0, 0},
			},
		*/
		/*
			{
				// Case that causes linear solver failure
				A: mat64.NewDense(13, 26, []float64{-0.7001209024399848, -0.7027502615621812, -0, 0.7354444798695736, -0, 0.476457578966189, 0.7001209024399848, 0.7027502615621812, 0, -0.7354444798695736, 0, -0.476457578966189, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1.8446087238391438, -0, -0, 0.7705609478497938, -0, -0, -2.7311218710244463, 0, 0, -0.7705609478497938, 0, 0, 2.7311218710244463, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -0, 0.8332519091897401, 0.7762132098737671, -0, -0, -0.052470638647269585, 0, -0.8332519091897401, -0.7762132098737671, 0, 0, 0.052470638647269585, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.9898577208292023, 0.31653724289408824, -0, -0, 0.17797227766447388, 1.2702427184954932, -0.7998764021535656, -0.31653724289408824, 0, 0, -0.17797227766447388, -1.2702427184954932, 0.7998764021535656, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -0, -0, -0, -0, 0.4206278126213235, -0.7253374879437113, 0, 0, 0, 0, -0.4206278126213235, 0.7253374879437113, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, -1, -0, -0.7567988418466963, 0.3304567624749696, 0.8385927625193501, -0.0021606686026376387, -0, 0, 0.7567988418466963, -0.3304567624749696, -0.8385927625193501, 0.0021606686026376387, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, -2.230107839590404, -0.9897104202085316, -0, 0.24703471683023603, -0, -2.382860345431941, 0.6206871648345162, 0.9897104202085316, 0, -0.24703471683023603, 0, 2.382860345431941, -0.6206871648345162, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, -1, 1.4350469221322282, -0.9730343818431852, -0, 2.326429855201535, -0, -0.14347849887004038, -1.4350469221322282, 0.9730343818431852, 0, -2.326429855201535, 0, 0.14347849887004038, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, -0.7943912888763849, -0.13735037357335078, -0.5101161104860161, -0, -0, -1.4790634590370297, 0.050911195996747316, 0.13735037357335078, 0.5101161104860161, 0, 0, 1.4790634590370297, -0.050911195996747316, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, -0, -0.2515400440591492, 0.2058339272568599, -0, -0, -1.314023802253438, 0, 0.2515400440591492, -0.2058339272568599, 0, 0, 1.314023802253438, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, -1, 0.08279503413614919, -0.16669071891829756, -0, -0.6208413721884664, -0, -0.6348258970402827, -0.08279503413614919, 0.16669071891829756, 0, 0.6208413721884664, 0, 0.6348258970402827, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, -1, -0, -0, 0.49634739711260845, -0, -0, -0, 0, 0, -0.49634739711260845, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, -1, -0.2797437186631715, -0.8356683570259136, 1.8970426594969672, -0.4095711945594497, 0.45831284820623924, -0.6109615338552246, 0.2797437186631715, 0.8356683570259136, -1.8970426594969672, 0.4095711945594497, -0.45831284820623924, 0.6109615338552246, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, -1}),
				b: []float64{-0.8446087238391436, 0, 1.9898577208292023, 0, 0, -2.230107839590404, 0, 0.20560871112361512, 0, 0, 0, 0, 0},
				c: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
		*/
		/*
			{
				// Phase 1 of the above that panics
				A: mat64.NewDense(26, 52, []float64{0.7001209024399848, -0, -0, -0.31653724289408824, -0, -0, 0.9897104202085316, -1.4350469221322282, 0.13735037357335078, -0, -0.08279503413614919, -0, 0.2797437186631715, -0.7001209024399848, 0, 0, 0.31653724289408824, 0, 0, -0.9897104202085316, 1.4350469221322282, -0.13735037357335078, 0, 0.08279503413614919, 0, -0.2797437186631715, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.7027502615621812, -0, -0.8332519091897401, -0, -0, 0.7567988418466963, -0, 0.9730343818431852, 0.5101161104860161, 0.2515400440591492, 0.16669071891829756, -0, 0.8356683570259136, -0.7027502615621812, 0, 0.8332519091897401, 0, 0, -0.7567988418466963, 0, -0.9730343818431852, -0.5101161104860161, -0.2515400440591492, -0.16669071891829756, 0, -0.8356683570259136, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0.7705609478497938, -0.7762132098737671, -0, -0, -0.3304567624749696, -0.24703471683023603, -0, -0, -0.2058339272568599, -0, -0.49634739711260845, -1.8970426594969672, 0, 0.7705609478497938, 0.7762132098737671, 0, 0, 0.3304567624749696, 0.24703471683023603, 0, 0, 0.2058339272568599, 0, 0.49634739711260845, 1.8970426594969672, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0.7354444798695736, -0, -0, -0.17797227766447388, -0, -0.8385927625193501, -0, -2.326429855201535, -0, -0, 0.6208413721884664, -0, 0.4095711945594497, 0.7354444798695736, 0, 0, 0.17797227766447388, 0, 0.8385927625193501, 0, 2.326429855201535, 0, 0, -0.6208413721884664, 0, -0.4095711945594497, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -1.2702427184954932, -0.4206278126213235, 0.0021606686026376387, 2.382860345431941, -0, 1.4790634590370297, -0, -0, -0, -0.45831284820623924, 0, 0, 0, 1.2702427184954932, 0.4206278126213235, -0.0021606686026376387, -2.382860345431941, 0, -1.4790634590370297, 0, 0, 0, 0.45831284820623924, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0.476457578966189, 2.7311218710244463, 0.052470638647269585, 0.7998764021535656, 0.7253374879437113, -0, -0.6206871648345162, 0.14347849887004038, -0.050911195996747316, 1.314023802253438, 0.6348258970402827, -0, 0.6109615338552246, 0.476457578966189, -2.7311218710244463, -0.052470638647269585, -0.7998764021535656, -0.7253374879437113, 0, 0.6206871648345162, -0.14347849887004038, 0.050911195996747316, -1.314023802253438, -0.6348258970402827, 0, -0.6109615338552246, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0.7001209024399848, -0, -0, 0.31653724289408824, -0, -0, -0.9897104202085316, 1.4350469221322282, -0.13735037357335078, -0, 0.08279503413614919, -0, -0.2797437186631715, 0.7001209024399848, 0, 0, -0.31653724289408824, 0, 0, 0.9897104202085316, -1.4350469221322282, 0.13735037357335078, 0, -0.08279503413614919, 0, 0.2797437186631715, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0.7027502615621812, -0, 0.8332519091897401, -0, -0, -0.7567988418466963, -0, -0.9730343818431852, -0.5101161104860161, -0.2515400440591492, -0.16669071891829756, -0, -0.8356683570259136, 0.7027502615621812, 0, -0.8332519091897401, 0, 0, 0.7567988418466963, 0, 0.9730343818431852, 0.5101161104860161, 0.2515400440591492, 0.16669071891829756, 0, 0.8356683570259136, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, 0.7705609478497938, 0.7762132098737671, -0, -0, 0.3304567624749696, 0.24703471683023603, -0, -0, 0.2058339272568599, -0, 0.49634739711260845, 1.8970426594969672, 0, -0.7705609478497938, -0.7762132098737671, 0, 0, -0.3304567624749696, -0.24703471683023603, 0, 0, -0.2058339272568599, 0, -0.49634739711260845, -1.8970426594969672, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.7354444798695736, -0, -0, 0.17797227766447388, -0, 0.8385927625193501, -0, 2.326429855201535, -0, -0, -0.6208413721884664, -0, -0.4095711945594497, -0.7354444798695736, 0, 0, -0.17797227766447388, 0, -0.8385927625193501, 0, -2.326429855201535, 0, 0, 0.6208413721884664, 0, 0.4095711945594497, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, 1.2702427184954932, 0.4206278126213235, -0.0021606686026376387, -2.382860345431941, -0, -1.4790634590370297, -0, -0, -0, 0.45831284820623924, 0, 0, 0, -1.2702427184954932, -0.4206278126213235, 0.0021606686026376387, 2.382860345431941, 0, 1.4790634590370297, 0, 0, 0, -0.45831284820623924, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.476457578966189, -2.7311218710244463, -0.052470638647269585, -0.7998764021535656, -0.7253374879437113, -0, 0.6206871648345162, -0.14347849887004038, 0.050911195996747316, -1.314023802253438, -0.6348258970402827, -0, -0.6109615338552246, -0.476457578966189, 2.7311218710244463, 0.052470638647269585, 0.7998764021535656, 0.7253374879437113, 0, -0.6206871648345162, 0.14347849887004038, -0.050911195996747316, 1.314023802253438, 0.6348258970402827, 0, 0.6109615338552246, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -1, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -1, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -1, -0, -0, -0, -0, -0, -0, -0, -0, -0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -0, -1, -0, -0, -0, -0, -0, -0, -0, -0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -0, -0, -1, -0, -0, -0, -0, -0, -0, -0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -0, -0, -0, -1, -0, -0, -0, -0, -0, -0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, -0, -0, -0, -0, -0, -0, -0, -1, -0, -0, -0, -0, -0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, -0, -0, -0, -0, -0, -0, -0, -0, -1, -0, -0, -0, -0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -1, -0, -0, -0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -1, -0, -0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -1, -0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1.8446087238391438, 1, -0.9898577208292023, 1, 1, 2.230107839590404, 1, 0.7943912888763849, 1, 1, 1, 1, 1, -1.8446087238391438, -1, 0.9898577208292023, -1, -1, -2.230107839590404, -1, -0.7943912888763849, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
				b: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				c: []float64{-0.8446087238391436, 0, 1.9898577208292023, 0, 0, -2.230107839590404, 0, 0.20560871112361512, 0, 0, 0, 0, 0, 0.8446087238391436, -0, -1.9898577208292023, -0, -0, 2.230107839590404, -0, -0.20560871112361512, -0, -0, -0, -0, -0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		*/
		{
			// Dense case that panics.
			A: mat64.NewDense(6, 15, []float64{0.3279477313560112, 0.04126296122557327, 0.24121743535067522, -0.8676933623438741, -0.3279477313560112, -0.04126296122557327, -0.24121743535067522, 0.8676933623438741, 1, 0, 0, 0, 0, 0, 1.3702148909442915, 0.43713186538468607, 0.8613818492485417, -0.9298615442657688, -0.037784779008231184, -0.43713186538468607, -0.8613818492485417, 0.9298615442657688, 0.037784779008231184, 0, 1, 0, 0, 0, 0, 0.3478112701177931, -0, 0.748352668598051, -0.4294796840343912, -0, 0, -0.748352668598051, 0.4294796840343912, 0, 0, 0, 1, 0, 0, 0, -1, -0, 1.1913912184457485, 1.732132186658447, 0.4026384828544584, 0, -1.1913912184457485, -1.732132186658447, -0.4026384828544584, 0, 0, 0, 1, 0, 0, -0.4598555419763902, -0, 1.2088959976921831, -0.7297794575275871, 1.9835614149566971, 0, -1.2088959976921831, 0.7297794575275871, -1.9835614149566971, 0, 0, 0, 0, 1, 0, 0.5986809560819324, 0.19738159369304414, -1.0647198575836367, -0, -0.7264943883762761, -0.19738159369304414, 1.0647198575836367, 0, 0.7264943883762761, 0, 0, 0, 0, 0, 1, 0.36269644970561576}),
			b: []float64{2.3702148909442915, 1.3478112701177931, 0, -0.4598555419763902, 1.5986809560819324, 1.3626964497056158},
			c: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	} {
		//simplex(test.initialBasic, test.c, test.A, test.b, test.tol)
		testSimplex(t, test.c, test.A, test.b, convergenceTol)
	}

	// Try a bunch of random LPs
	nTest := 10000
	infeasible := 0
	unbounded := 0
	bounded := 0
	singular := 0
	bad := 0
	for i := 0; i < nTest; i++ {
		maxN := 6
		n := rand.Intn(maxN) + 2 // n must be at least two.
		m := rand.Intn(n-1) + 1  // m must be between 1 and n
		if m == 0 || n == 0 {
			continue
		}
		randValue := func() float64 {
			pZero := 0.1 // make sure there are zeros
			//var pZero float64
			v := rand.Float64()
			if v < pZero {
				return 0
			}
			return rand.NormFloat64()
		}
		a := mat64.NewDense(m, n, nil)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				a.Set(i, j, randValue())
			}
		}
		b := make([]float64, m)
		for i := range b {
			b[i] = randValue()
		}

		c := make([]float64, n)
		for i := range c {
			c[i] = randValue()
		}

		errPrimal := testSimplex(t, c, a, b, convergenceTol)

		primalInfeasible := errPrimal == ErrInfeasible
		primalUnbounded := errPrimal == ErrUnbounded
		primalBounded := errPrimal == nil
		primalASingular := errPrimal == ErrSingular

		switch {
		case primalInfeasible:
			infeasible++
		case primalUnbounded:
			unbounded++
		case primalBounded:
			bounded++
		case primalASingular:
			singular++
		default:
			bad++
		}
	}
	fmt.Println("Cases: ", nTest)
	fmt.Println("Bounded:", bounded)
	fmt.Println("Singular:", singular)
	fmt.Println("Unbounded:", unbounded)
	fmt.Println("Infeasible:", infeasible)
	fmt.Println("Bad:", bad)

}

func testSimplex(t *testing.T, c []float64, a mat64.Matrix, b []float64, convergenceTol float64) error {
	fmt.Println("Starting primal simplex")
	primalOpt, primalX, _, errPrimal := simplex(nil, c, a, b, convergenceTol)
	fmt.Println("Done primal simplex")

	if errPrimal == nil {
		// Check that it's feasible
		var bCheck mat64.Vector
		bCheck.MulVec(a, mat64.NewVector(len(primalX), primalX))
		//if !bCheck.EqualsApproxVec(mat64.NewVector(len(b), b), 1e-10) {
		if !mat64.EqualApprox(&bCheck, mat64.NewVector(len(b), b), 1e-10) {
			t.Errorf("No error in primal but solution infeasible")
		}
	}

	primalInfeasible := errPrimal == ErrInfeasible
	primalUnbounded := errPrimal == ErrUnbounded
	primalBounded := errPrimal == nil
	primalASingular := errPrimal == ErrSingular

	primalBad := !primalInfeasible && !primalUnbounded && !primalBounded && !primalASingular

	// If singular the problem is undefined and if bad something went wrong
	if errPrimal == ErrSingular || primalBad {
		if primalBad {
			fmt.Println("Primal bad: ", errPrimal)
		}
		return errPrimal
	}

	// Otherwise, compare the answer to its dual.

	// Construct and solve the dual LP.
	// Standard Frm:
	//  minimize c^T * x
	//    subject to  A * x = b, x >= 0
	// The dual of this problem is
	//  maximize -b^T * nu
	//   subject to A^T * nu + c >= 0
	// Which is
	//   minimize b^T * nu
	//   subject to -A^T * nu <= c

	negAT := &mat64.Dense{}
	negAT.Clone(a.T())
	negAT.Scale(-1, negAT)
	cNew, aNew, bNew := Convert(b, negAT, c, nil, nil)

	fmt.Println("Starting dual simplex")
	dualOpt, dualX, _, errDual := simplex(nil, cNew, aNew, bNew, convergenceTol)
	fmt.Println("Done dual simplex")
	if errDual == nil {
		// Check that the dual is feasible
		var bCheck mat64.Vector
		bCheck.MulVec(aNew, mat64.NewVector(len(dualX), dualX))
		//if !bCheck.EqualsApproxVec(mat64.NewVector(len(bNew), bNew), 1e-10) {
		if !mat64.EqualApprox(&bCheck, mat64.NewVector(len(bNew), bNew), 1e-10) {
			t.Errorf("No error in dual but solution infeasible")
		}
	}
	_ = primalOpt
	_ = dualOpt

	// Quick hack to deal with ErrZeroRow
	if errDual == ErrZeroColumn {
		return errPrimal
	}

	// If the primal problem is feasible, then the primal and the dual should
	// be the same answer. We have flopped the sign in the dual (minimizing
	// b^T *nu instead of maximizing -b^T*nu), so flip it back.
	if errPrimal == nil {
		if errDual != nil {
			t.Errorf("Primal feasible but dual errored: %s", errDual)
		}
		if math.Abs(primalOpt+dualOpt) > convergenceTol {
			t.Errorf("Primal and dual value mismatch. Primal %v, dual %v.", primalOpt, dualOpt)
		}
	}
	// If the primal problem is unbounded, then the dual is infeasible.
	if errPrimal == ErrUnbounded && errDual != ErrInfeasible {
		t.Errorf("Primal unbounded but dual not infeasible. ErrDual = %s", errDual)
	}

	// If the dual is unbounded, then the primal is infeasible.
	if errDual == ErrUnbounded && errPrimal != ErrInfeasible {
		t.Errorf("Dual unbounded but primal not infeasible. ErrDual = %s", errPrimal)
	}

	// It may be the case that both the primal and the dual are infeasible.
	if errPrimal == ErrInfeasible {
		if errDual != ErrUnbounded && errDual != ErrInfeasible {
			t.Errorf("Primal infeasible but dual not infeasible or unbounded: %s", errDual)
		}
	}

	return errPrimal
}
