package main

func getSsAccount() (ssAcount *SsAcount, err error) {
	ssAcount = &SsAcount{
		Server:     "45.76.214.188",
		ServerPort: 8392,
		Password:   "mock123",
		Method:     "aes-256-cfb",
	}
	return
}
