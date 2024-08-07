package dockerNetwork

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/wharf/wharf/pkg/errors"
)

func GetAll(client *client.Client, ctx context.Context, ch chan *types.NetworkResource, errCh chan *errors.Error) {

	networks, err := client.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		errStruc := &errors.Error{
			Name: "Listing images",
			Err:  fmt.Errorf("error while docker images listing: %w", err),
		}
		errCh <- errStruc
		close(errCh)
		close(ch)
		return
	}
	close(errCh)

	for _, network := range networks {

		ch <- &network
	}

	close(ch)
}

func Prune(client *client.Client, ctx context.Context) (types.NetworksPruneReport, error) {
	pruneReport, err := client.NetworksPrune(ctx, filters.Args{})
	return pruneReport, err
}

func Remove(client *client.Client, ctx context.Context, id string) error {
	err := client.NetworkRemove(ctx, id)
	return err
}

func Disconnect(client *client.Client, ctx context.Context, networkId string, containerId string, force bool) error {
	err := client.NetworkDisconnect(ctx, networkId, containerId, force)
	return err
}

func Connect(client *client.Client, ctx context.Context, networkId string, containerId string) error {
	err := client.NetworkConnect(ctx, networkId, containerId, nil)
	return err
}

func Create(client *client.Client, ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	res, err := client.NetworkCreate(ctx, name, options)
	return res, err
}
