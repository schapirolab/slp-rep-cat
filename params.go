// File generated by params.SaveGoCode

package main

import "github.com/emer/emergent/params"

var SavedParamsSets = params.Sets{
	{Name: "Base", Desc: "these are the best params", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: "Prjn", Desc: "keeping default params for generic prjns",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "true",
					"Prjn.Learn.Norm.On":     "true",
					"Prjn.Learn.WtBal.On":    "false",
				}},
			{Sel: ".PerPerPrjn", Desc: "Higer perception itself",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.02",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "false",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var": "0.5",
					"Prjn.CHL.Hebb": "0",
				}},
			{Sel: ".PerCTXPrjn", Desc: "Higher perception to MSP - low learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.0001",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.1",
					"Prjn.WtInit.Var":  "0.5",
					//"Prjn.CHL.MinusQ1" : "true",
				}},
			{Sel: ".CA1CTXPrjn", Desc: "Higher perception to MSP - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.0001",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
				}},
			{Sel: ".PerDGPrjn", Desc: "Higher perception to DG - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
					"Prjn.WtScale.Abs": "1",
				}},
			{Sel: ".PerCA3Prjn", Desc: "Higher perception to MSP - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
				}},
			{Sel: ".PerCA1Prjn", Desc: "Higher perception to MSP - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.05",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
				}},
			{Sel: "#CA3ToCA3", Desc: "CA3 recurrent - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
					"Prjn.CHL.SAvgCor": "1",
					"Prjn.WtScale.Rel": "2",
					"Prjn.WtInit.Mean": "0.1",
					"Prjn.WtInit.Var":  "0.5",
				}},
			{Sel: "#CodeNameToDG", Desc: "Higher WtScale",
				Params: params.Params{
					//"Prjn.Learn.Lrate":       "0.075",
					//"Prjn.Learn.Momentum.On": "false",
					//"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					//"Prjn.CHL.Hebb": "0",
					//"Prjn.CHL.SAvgCor": "1",
					"Prjn.WtScale.Rel": "5",
					//"Prjn.WtInit.Mean": "0.5",
					//"Prjn.WtInit.Var":  "0.25",
				}},
			{Sel: "#DGToCA3", Desc: "DG to CA3 - non-learning",
				Params: params.Params{
					"Prjn.CHL.SAvgCor":"1",
					"Prjn.Learn.Learn":"false",
					"Prjn.WtInit.Mean":"0.9",
					"Prjn.WtInit.Var":"0.01",
					"Prjn.WtScale.Rel":"10",
					//"Prjn.WtScale.Abs":"1",
					"Prjn.Learn.Momentum.On":"false",
					"Prjn.Learn.Norm.On":"false",
				}},
			//{Sel: "#pCA1ToCTX", Desc: "Higher perception to MSP - low learning rate",
			//	Params: params.Params{
			//		"Prjn.Learn.Lrate":       "0.0001",
			//		"Prjn.Learn.Momentum.On": "false",
			//		"Prjn.Learn.Norm.On":     "false",
			//		//"Prjn.Learn.WtBal.On":    "true",
			//		"Prjn.CHL.Hebb": "0",
			//		"Prjn.WtInit.Mean": "0.1",
			//		"Prjn.WtInit.Var":  "0.5",
			//		//"Prjn.CHL.MinusQ1" : "true",
			//	}},
			//{Sel: "#dCA1ToCTX", Desc: "Higher perception to MSP - low learning rate",
			//	Params: params.Params{
			//		"Prjn.Learn.Lrate":       "0.0001",
			//		"Prjn.Learn.Momentum.On": "false",
			//		"Prjn.Learn.Norm.On":     "false",
			//		//"Prjn.Learn.WtBal.On":    "true",
			//		"Prjn.CHL.Hebb": "0",
			//		"Prjn.WtInit.Mean": "0.1",
			//		"Prjn.WtInit.Var":  "0.5",
			//		//"Prjn.CHL.MinusQ1" : "true",
			//	}},
			{Sel: "#CTXTodCA1", Desc: "Higher perception to MSP - low learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.0001",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
					//"Prjn.WtInit.Mean": "0.1",
					//"Prjn.WtInit.Var":  "0.5",
					//"Prjn.CHL.MinusQ1" : "true",
				}},
			{Sel: ".Per", Desc: "All Per layers",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.1",
				}},
			{Sel: "#CodeName", Desc: "Codename only",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.6",
				}},
			{Sel: "#ClassName", Desc: "All Per layers",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "1.6",
				}},
			{Sel: "#CTX", Desc: "low inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.4",
				}},
			{Sel: "#dCA1", Desc: "low inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.2", // 2.1
				}},
			{Sel: "#pCA1", Desc: "low inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.2", // 2.5
				}},
			{Sel: "#DG", Desc: "very sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.8",
				}},
			{Sel: "#CA3", Desc: "very sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "5",
				}},
		},
	}},
	{Name: "NoCHL", Desc: "no learning in CHL main hip pathways -- for debugging auto-encoder", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "no learning",
				Params: params.Params{
					"Prjn.Learn.Lrate": "0",
				}},
		},
		"Sim": &params.Sheet{},
	}},
	{Name: "CHLWtBal", Desc: "CHL uses weight balance -- much better -- now in base", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{},
	}},
	{Name: "EncWtBal", Desc: "encoder uses weight balance -- worse", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".EcCa1Prjn", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{},
	}},
	{Name: "EncMom", Desc: "encoder uses momentum -- worse", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".EcCa1Prjn", Desc: "moment on",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "true",
					"Prjn.Learn.Norm.On":     "true",
				}},
		},
		"Sim": &params.Sheet{},
	}},
	{Name: "AllWtBal", Desc: "All use weight balance", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
			{Sel: ".EcCa1Prjn", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{},
	}},
}