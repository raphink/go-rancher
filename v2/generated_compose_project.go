package client

const (
	COMPOSE_PROJECT_TYPE = "composeProject"
)

type ComposeProject struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Binding *Binding `json:"binding,omitempty" yaml:"binding,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Environment map[string]interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	HealthState string `json:"healthState,omitempty" yaml:"health_state,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	PreviousEnvironment map[string]interface{} `json:"previousEnvironment,omitempty" yaml:"previous_environment,omitempty"`

	PreviousExternalId string `json:"previousExternalId,omitempty" yaml:"previous_external_id,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Templates map[string]interface{} `json:"templates,omitempty" yaml:"templates,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ComposeProjectCollection struct {
	Collection
	Data   []ComposeProject `json:"data,omitempty"`
	client *ComposeProjectClient
}

type ComposeProjectClient struct {
	rancherClient *RancherClient
}

type ComposeProjectOperations interface {
	List(opts *ListOpts) (*ComposeProjectCollection, error)
	Create(opts *ComposeProject) (*ComposeProject, error)
	Update(existing *ComposeProject, updates interface{}) (*ComposeProject, error)
	ById(id string) (*ComposeProject, error)
	Delete(container *ComposeProject) error

	ActionCancelupgrade(*ComposeProject) (*Stack, error)

	ActionCreate(*ComposeProject) (*Stack, error)

	ActionError(*ComposeProject) (*Stack, error)

	ActionFinishupgrade(*ComposeProject) (*Stack, error)

	ActionRemove(*ComposeProject) (*Stack, error)

	ActionRollback(*ComposeProject) (*Stack, error)
}

func newComposeProjectClient(rancherClient *RancherClient) *ComposeProjectClient {
	return &ComposeProjectClient{
		rancherClient: rancherClient,
	}
}

func (c *ComposeProjectClient) Create(container *ComposeProject) (*ComposeProject, error) {
	resp := &ComposeProject{}
	err := c.rancherClient.doCreate(COMPOSE_PROJECT_TYPE, container, resp)
	return resp, err
}

func (c *ComposeProjectClient) Update(existing *ComposeProject, updates interface{}) (*ComposeProject, error) {
	resp := &ComposeProject{}
	err := c.rancherClient.doUpdate(COMPOSE_PROJECT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ComposeProjectClient) List(opts *ListOpts) (*ComposeProjectCollection, error) {
	resp := &ComposeProjectCollection{}
	err := c.rancherClient.doList(COMPOSE_PROJECT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ComposeProjectCollection) Next() (*ComposeProjectCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ComposeProjectCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ComposeProjectClient) ById(id string) (*ComposeProject, error) {
	resp := &ComposeProject{}
	err := c.rancherClient.doById(COMPOSE_PROJECT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ComposeProjectClient) Delete(container *ComposeProject) error {
	return c.rancherClient.doResourceDelete(COMPOSE_PROJECT_TYPE, &container.Resource)
}

func (c *ComposeProjectClient) ActionCancelupgrade(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "cancelupgrade", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ComposeProjectClient) ActionCreate(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ComposeProjectClient) ActionError(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "error", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ComposeProjectClient) ActionFinishupgrade(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "finishupgrade", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ComposeProjectClient) ActionRemove(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ComposeProjectClient) ActionRollback(resource *ComposeProject) (*Stack, error) {

	resp := &Stack{}

	err := c.rancherClient.doAction(COMPOSE_PROJECT_TYPE, "rollback", &resource.Resource, nil, resp)

	return resp, err
}
