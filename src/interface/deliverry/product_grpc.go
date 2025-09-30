package deliverry

import (
	"context"

	pb "github.com/faujiahmat/zentra-proto/protogen/product"
)

type ProductGrpc interface {
	FindManyByIds(ctx context.Context, productIds []uint32) ([]*pb.ProductCart, error)
}
