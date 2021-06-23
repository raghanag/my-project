package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/raghanag/my-project/cmd/go-graphql/graph/generated"
	"github.com/raghanag/my-project/cmd/go-graphql/graph/model"
)

func (r *myQueryResolver) IngestData(ctx context.Context) (*model.IngestData, error) {
	status, err := r.Ingest.Create()

	if err != nil {
		return nil, err
	} else {
		return &model.IngestData{
			IngestionStatus: *status,
		}, nil
	}
}

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myQueryResolver struct{ *Resolver }
