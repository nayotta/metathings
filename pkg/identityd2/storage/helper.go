package metathings_identityd2_storage

func wrap_token(tkn *Token) *Token {
	tkn.Domain = &Domain{Id: tkn.DomainId}
	tkn.Entity = &Entity{Id: tkn.EntityId}
	if tkn.CredentialId != nil {
		tkn.Credential = &Credential{Id: tkn.CredentialId}
	}

	return tkn
}
