package htpasswd

import "testing"

func TestPasswordHashes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		Name		string
		Password	string
		Hash		string
		Match		bool
		Error		bool
	}{{Name: "md5 empty", Password: "", Hash: "$apr1$AQEmuiLe$2lPjK2hL6mnTakRWskgaQ1", Match: true}, {Name: "md5 password", Password: "password", Hash: "$apr1$6TMtuxUJ$0M76TkGjp0qVg/e7rfk22.", Match: true}, {Name: "md5 mismatch", Password: "mypassword", Hash: "$apr1$6TMtuxUJ$0M76TkGjp0qVg/e7rfk22.", Match: false}, {Name: "md5 missing salt", Password: "password", Hash: "$apr1$$0M76TkGjp0qVg/e7rfk22.", Match: false, Error: true}, {Name: "md5 with long password", Password: "passwordthatisreallyreallyreallyreallyreallyreallyreallylong", Hash: "$apr1$6VmuPCYl$OvuHDqaS59nsRov9HnsGc1", Match: true}, {Name: "crypt empty", Password: "", Hash: "lNO4S8u4F4oNo", Match: false, Error: true}, {Name: "crypt match", Password: "password", Hash: ".zs/E.NK2vwFs", Match: false, Error: true}, {Name: "crypt mismatch", Password: "mypassword", Hash: ".zs/E.NK2vwFs", Match: false, Error: true}, {Name: "crypt missing salt", Password: "password", Hash: "s", Match: false, Error: true}, {Name: "sha empty", Password: "", Hash: "{SHA}2jmj7l5rSw0yVb/vlWAYkK/YBwk=", Match: true}, {Name: "sha match", Password: "password", Hash: "{SHA}W6ph5Mm5Pz8GgiULbPgzG37mj9g=", Match: true}, {Name: "sha mismatch", Password: "mypassword", Hash: "{SHA}W6ph5Mm5Pz8GgiULbPgzG37mj9g=", Match: false}, {Name: "sha invalid", Password: "mypassword", Hash: "{SHA}", Match: false, Error: true}, {Name: "bcrypt strength 5 empty", Password: "", Hash: "$2y$05$Edf.Eeznh19sIYYcTc7YOeltcWjzFuvrcYp57lq78diiJr512GILm", Match: true, Error: false}, {Name: "bcrypt strength 5 match", Password: "password", Hash: "$2y$05$Vfd6hjeQXB6nTFTVMkoFE.CAItk2W8akuomafFBakd0n/mHqIzoUO", Match: true, Error: false}, {Name: "bcrypt strength 5 mismatch", Password: "mypassword", Hash: "$2y$05$Vfd6hjeQXB6nTFTVMkoFE.CAItk2W8akuomafFBakd0n/mHqIzoUO", Match: false, Error: false}, {Name: "bcrypt strength 10 empty", Password: "", Hash: "$2y$10$v0c.7wrYEv2AZnLsPXO57.48Qc5widamyKkmwrUolKwYW0Zw8zhJ.", Match: true, Error: false}, {Name: "bcrypt strength 10 match", Password: "password", Hash: "$2y$10$Fk32bQky/.91nbecGjFfPO1m97V12d.ickjAzpNF22NgMKs4qWDOK", Match: true, Error: false}, {Name: "bcrypt strength 10 mismatch", Password: "mypassword", Hash: "$2y$10$Fk32bQky/.91nbecGjFfPO1m97V12d.ickjAzpNF22NgMKs4qWDOK", Match: false, Error: false}, {Name: "bcrypt missing strength", Password: "password", Hash: "$2y$$Fk32bQky/.91nbecGjFfPO1m97V12d.ickjAzpNF22NgMKs4qWDOK", Match: false, Error: true}}
	for _, testCase := range testCases {
		match, err := testPassword(testCase.Password, testCase.Hash)
		if testCase.Error != (err != nil) {
			t.Errorf("%s: Expected error=%v, got %v", testCase.Name, testCase.Error, err)
		}
		if match != testCase.Match {
			t.Errorf("%s: Expected match=%v, got %v", testCase.Name, testCase.Match, match)
		}
	}
}
