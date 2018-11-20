package cmd

import cmd_contrib "github.com/nayotta/metathings/cmd/contrib"

func init_service_cmd_option(dst, src interface{}) {
	init_listen_optioner(dst, src)
	init_storage_optioner(dst, src)
	init_transport_credential_optioner(dst, src)
}

func init_listen_optioner(dst, src interface{}) {
	dst_lis, ok1 := dst.(cmd_contrib.ListenOptioner)
	src_lis, ok2 := src.(cmd_contrib.ListenOptioner)

	if ok1 && ok2 {
		if dst_lis.GetListen() == "" {
			dst_lis.SetListen(src_lis.GetListen())
		}
	}
}

func init_storage_optioner(dst, src interface{}) {
	dst_gst, ok1 := dst.(cmd_contrib.GetStorageOptioner)
	src_gst, ok2 := src.(cmd_contrib.GetStorageOptioner)

	if ok1 && ok2 {
		if dst_gst.GetStorage().GetDriver() == "" {
			dst_gst.GetStorage().SetDriver(src_gst.GetStorage().GetDriver())
		}

		if dst_gst.GetStorage().GetUri() == "" {
			dst_gst.GetStorage().SetDriver(src_gst.GetStorage().GetUri())
		}
	}
}

func init_transport_credential_optioner(dst, src interface{}) {
	dst_cred, ok1 := dst.(cmd_contrib.GetTransportCredentialOptioner)
	src_cred, ok2 := dst.(cmd_contrib.GetTransportCredentialOptioner)

	if ok1 && ok2 {
		if dst_cred.GetTransportCredential().GetCertFile() == "" {
			dst_cred.GetTransportCredential().SetCertFile(src_cred.GetTransportCredential().GetCertFile())
		}

		if dst_cred.GetTransportCredential().GetKeyFile() == "" {
			dst_cred.GetTransportCredential().SetKeyFile(src_cred.GetTransportCredential().GetKeyFile())
		}
	}
}
