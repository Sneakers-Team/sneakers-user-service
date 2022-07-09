package user

import "msqrd/user-service/pkg/store"

type GRPCServer struct {
	Store *store.Store
}
