package component

import (
	"fmt"
	"github.com/JPereirax/GOIF/internal"
	"github.com/JPereirax/GOIF/types"
)

type DirectComponent struct{}

func init() {
	CreateComponent("direct").Processor(DirectComponent{}.run).End()
}

func (DirectComponent) run(ctx types.ContextRequest) (*types.HttpTransport, error) {
	for _, functions := range internal.ContextFunctions {
		if functions.Id == fmt.Sprintf("direct:%s", ctx.Component.Uri) {
			for _, processor := range functions.Processors {
				ctx.HttpTransport = processor(ctx.HttpTransport)
			}
			break
		}
	}

	return ctx.HttpTransport, nil
}