package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	client_native "github.com/haproxytech/client-native/v2"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/nameserver"
	"github.com/haproxytech/models/v2"
)

//CreateNameserverHandlerImpl implementation of the CreateNameserverHandler interface using client-native client
type CreateNameserverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//DeleteNameserverHandlerImpl implementation of the DeleteNameserverHandler interface using client-native client
type DeleteNameserverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//GetNameserverHandlerImpl implementation of the GetNameserverHandler interface using client-native client
type GetNameserverHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//GetNameserversHandlerImpl implementation of the GetNameserversHandler interface using client-native client
type GetNameserversHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//ReplaceNameserverHandlerImpl implementation of the ReplaceNameserverHandler interface using client-native client
type ReplaceNameserverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//Handle executing the request and returning a response
func (h *CreateNameserverHandlerImpl) Handle(params nameserver.CreateNameserverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return nameserver.NewCreateNameserverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.CreateNameserver(params.Resolver, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return nameserver.NewCreateNameserverDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return nameserver.NewCreateNameserverDefault(int(*e.Code)).WithPayload(e)
			}
			return nameserver.NewCreateNameserverCreated().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return nameserver.NewCreateNameserverAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return nameserver.NewCreateNameserverAccepted().WithPayload(params.Data)
}

//Handle executing the request and returning a response
func (h *DeleteNameserverHandlerImpl) Handle(params nameserver.DeleteNameserverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return nameserver.NewDeleteNameserverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.DeleteNameserver(params.Name, params.Resolver, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return nameserver.NewDeleteNameserverDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return nameserver.NewDeleteNameserverDefault(int(*e.Code)).WithPayload(e)
			}
			return nameserver.NewDeleteNameserverNoContent()
		}
		rID := h.ReloadAgent.Reload()
		return nameserver.NewDeleteNameserverAccepted().WithReloadID(rID)
	}
	return nameserver.NewDeleteNameserverAccepted()
}

//Handle executing the request and returning a response
func (h *GetNameserverHandlerImpl) Handle(params nameserver.GetNameserverParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, b, err := h.Client.Configuration.GetNameserver(params.Name, params.Resolver, t)
	if err != nil {
		e := misc.HandleError(err)
		return nameserver.NewGetNameserverDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return nameserver.NewGetNameserverOK().WithPayload(&nameserver.GetNameserverOKBody{Version: v, Data: b}).WithConfigurationVersion(v)
}

//Handle executing the request and returning a response
func (h *GetNameserversHandlerImpl) Handle(params nameserver.GetNameserversParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, bs, err := h.Client.Configuration.GetNameservers(params.Resolver, t)
	if err != nil {
		e := misc.HandleError(err)
		return nameserver.NewGetNameserversDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return nameserver.NewGetNameserversOK().WithPayload(&nameserver.GetNameserversOKBody{Version: v, Data: bs}).WithConfigurationVersion(v)
}

//Handle executing the request and returning a response
func (h *ReplaceNameserverHandlerImpl) Handle(params nameserver.ReplaceNameserverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return nameserver.NewReplaceNameserverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.EditNameserver(params.Name, params.Resolver, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return nameserver.NewReplaceNameserverDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return nameserver.NewReplaceNameserverDefault(int(*e.Code)).WithPayload(e)
			}
			return nameserver.NewReplaceNameserverOK().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return nameserver.NewReplaceNameserverAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return nameserver.NewReplaceNameserverAccepted().WithPayload(params.Data)
}