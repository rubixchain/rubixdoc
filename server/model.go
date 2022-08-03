package main

import "github.com/EnsurityTechnologies/ensweb"

const (
	UserTable     string = "UserTable"
	RoleTable     string = "RoleTable"
	UserRoleTable string = "UserRoleTable"
	TenantTable   string = "TenantTable"
)

func (s *Server) createModels() error {
	cfg := ensweb.EntityConfig{
		DefaultTenantName:    "rubixdoc",
		DefaultAdminName:     "admin",
		DefaultAdminPassword: "admin@123",
		TenantTableName:      TenantTable,
		UserTableName:        UserTable,
		RoleTableName:        RoleTable,
		UserRoleTableName:    UserRoleTable,
	}
	err := s.SetupEntity(cfg)
	if err != nil {
		s.log.Error("failed to setup entities")
		return err
	}

	return nil
}
