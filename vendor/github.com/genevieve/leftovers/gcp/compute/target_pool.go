package compute

import "fmt"

type TargetPool struct {
	client targetPoolsClient
	name   string
	region string
}

func NewTargetPool(client targetPoolsClient, name, region string) TargetPool {
	return TargetPool{
		client: client,
		name:   name,
		region: region,
	}
}

func (t TargetPool) Delete() error {
	err := t.client.DeleteTargetPool(t.region, t.name)

	if err != nil {
		return fmt.Errorf("ERROR deleting target pool %s: %s", t.name, err)
	}

	return nil
}

func (t TargetPool) Name() string {
	return t.name
}
