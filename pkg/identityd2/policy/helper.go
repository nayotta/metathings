package metathings_identityd2_policy

import (
	"fmt"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

func ConvertGroup(grp *storage.Group) string {
	return fmt.Sprintf("dom.%s.grp.%s", *grp.DomainId, *grp.Id)
}

func ConvertSubject(sub *storage.Entity) string {
	return fmt.Sprintf("sub.%s", *sub.Id)
}

func ConvertObject(obj *storage.Entity) string {
	return fmt.Sprintf("obj.%s", *obj.Id)
}

func ConvertEntity(ent *storage.Entity) string {
	return fmt.Sprintf("ent.%s", *ent.Id)
}

func ConvertRoleForObject(grp *storage.Group) string {
	return fmt.Sprintf("dom.%s.grp.%s.data", *grp.DomainId, *grp.Id)
}

func ConvertRoleForSubject(grp *storage.Group, rol *storage.Role) string {
	return fmt.Sprintf("dom.%s.grp.%s.rol.%s", *grp.DomainId, *grp.Id, *rol.Name)
}

func ConvertRolesForSubject(grp *storage.Group) []string {
	var ys []string

	for _, r := range grp.Roles {
		ys = append(ys, ConvertRoleForSubject(grp, r))
	}

	return ys
}

func ConvertRole(grp *storage.Group, rol *storage.Role) string {
	return fmt.Sprintf("dom.%s.grp.%s.rol.%s", *grp.DomainId, *grp.Id, *rol.Name)
}

func ConvertUngroupingRole(rol *storage.Role) string {
	return fmt.Sprintf("rol.%s", *rol.Name)
}
