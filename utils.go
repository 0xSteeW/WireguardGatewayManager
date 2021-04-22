package wgm

func assignIfNull(original *string, new string) {
	if (original == nil || *original == "") && new != "" {
		*original = new
	}
}
