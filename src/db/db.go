package db

import (
	"github.com/Robert076/tips-microservice/tip"
)

func GetTipsFromDatabase() []tip.Tip {
	tips := []tip.Tip{
		{Id: 1, Text: "Containerize your apps using multi-stage builds to save storage space."},
		{Id: 2, Text: "Use the single-responsability principle."},
		{Id: 3, Text: "Do not nest if statements under any circumstances."},
		{Id: 4, Text: "Put your codebase in a source control provider like GitHub"},
	}
	return tips
}
